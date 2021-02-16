// Copyright 2018 The Cluster Monitoring Operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"github.com/imdario/mergo"
	"k8s.io/apimachinery/pkg/util/wait"
	"reflect"
	"strings"
	"time"

	"github.com/pkg/errors"
	appsv1 "k8s.io/api/apps/v1"
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const (
	deploymentCreateTimeout = 5 * time.Minute
	metadataPrefix          = "monitoring.openshift.io/"
)

type Client struct {
	version               string
	namespace             string
	userWorkloadNamespace string
	namespaceSelector     string
	kclient               kubernetes.Interface
	eclient               apiextensionsclient.Interface
}

func (c *Client) CreateOrUpdateDeployment(dep *appsv1.Deployment) error {
	existing, err := c.kclient.AppsV1().Deployments(dep.GetNamespace()).Get(dep.GetName(), metav1.GetOptions{})

	if apierrors.IsNotFound(err) {
		err = c.CreateDeployment(dep)
		return errors.Wrap(err, "creating Deployment object failed")
	}
	if err != nil {
		return errors.Wrap(err, "retrieving Deployment object failed")
	}
	if reflect.DeepEqual(dep.Spec, existing.Spec) {
		// Nothing to do, as the currently existing deployment is equivalent to the one that would be applied.
		return nil
	}

	required := dep.DeepCopy()
	mergeMetadata(&required.ObjectMeta, existing.ObjectMeta)

	err = c.UpdateDeployment(required)
	if err != nil {
		uErr, ok := err.(*apierrors.StatusError)
		if ok && uErr.ErrStatus.Code == 422 && uErr.ErrStatus.Reason == metav1.StatusReasonInvalid {
			// try to delete Deployment
			err = c.DeleteDeployment(existing)
			if err != nil {
				return errors.Wrap(err, "deleting Deployment object failed")
			}
			err = c.CreateDeployment(required)
			if err != nil {
				return errors.Wrap(err, "creating Deployment object failed after update failed")
			}
		}
		return errors.Wrap(err, "updating Deployment object failed")
	}
	return nil
}

func (c *Client) CreateDeployment(dep *appsv1.Deployment) error {
	d, err := c.kclient.AppsV1().Deployments(dep.GetNamespace()).Create(dep)
	if err != nil {
		return err
	}

	return c.WaitForDeploymentRollout(d)
}

func (c *Client) WaitForDeploymentRollout(dep *appsv1.Deployment) error {
	var lastErr error
	if err := wait.Poll(time.Second, deploymentCreateTimeout, func() (bool, error) {
		d, err := c.kclient.AppsV1().Deployments(dep.GetNamespace()).Get(dep.GetName(), metav1.GetOptions{})
		if err != nil {
			return false, err
		}
		if d.Generation > d.Status.ObservedGeneration {
			lastErr = errors.Errorf("current generation %d, observed generation %d",
				d.Generation, d.Status.ObservedGeneration)
			return false, nil
		}
		if d.Status.UpdatedReplicas != d.Status.Replicas {
			lastErr = errors.Errorf("expected %d replicas, got %d updated replicas",
				d.Status.Replicas, d.Status.UpdatedReplicas)
			return false, nil
		}
		if d.Status.UnavailableReplicas != 0 {
			lastErr = errors.Errorf("got %d unavailable replicas",
				d.Status.UnavailableReplicas)
			return false, nil
		}
		return true, nil
	}); err != nil {
		if err == wait.ErrWaitTimeout && lastErr != nil {
			err = lastErr
		}
		return errors.Wrapf(err, "waiting for DeploymentRollout of %s/%s", dep.GetNamespace(), dep.GetName())
	}
	return nil
}

// mergeMetadata merges labels and annotations from `existing` map into `required` one where `required` has precedence
// over `existing` keys and values. Additionally function performs filtering of labels and annotations from `exiting` map
// where keys starting from string defined in `metadataPrefix` are deleted. This prevents issues with preserving stale
// metadata defined by the operator
func mergeMetadata(required *metav1.ObjectMeta, existing metav1.ObjectMeta) {
	for k := range existing.Annotations {
		if strings.HasPrefix(k, metadataPrefix) {
			delete(existing.Annotations, k)
		}
	}

	for k := range existing.Labels {
		if strings.HasPrefix(k, metadataPrefix) {
			delete(existing.Labels, k)
		}
	}

	mergo.Merge(&required.Annotations, existing.Annotations)
	mergo.Merge(&required.Labels, existing.Labels)
}

func (c *Client) UpdateDeployment(dep *appsv1.Deployment) error {
	updated, err := c.kclient.AppsV1().Deployments(dep.GetNamespace()).Update(dep)
	if err != nil {
		return err
	}

	return c.WaitForDeploymentRollout(updated)
}

func (c *Client) DeleteDeployment(d *appsv1.Deployment) error {
	p := metav1.DeletePropagationForeground
	err := c.kclient.AppsV1().Deployments(d.GetNamespace()).Delete(d.GetName(), &metav1.DeleteOptions{PropagationPolicy: &p})
	if apierrors.IsNotFound(err) {
		return nil
	}

	return err
}
