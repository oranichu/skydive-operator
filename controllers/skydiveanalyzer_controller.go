/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"github.com/pkg/errors"
	"k8s.io/client-go/tools/clientcmd"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	kclient "skydive/pkg/kclient"
	"skydive/pkg/manifests"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	skydivev1beta1 "skydive/api/v1beta1"
)

const (
	namespaceName  = "default"
	kubeconfigPath = "/Users/orannahoum/.kube/config"
)

var (
	SkydiveAnalyzerDeployment = "skydive-analyzer/deployment.yaml"
)

// SkydiveAnalyzerReconciler reconciles a SkydiveAnalyzer object
type SkydiveAnalyzerReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme

	kclient.KClient
}

// +kubebuilder:rbac:groups=skydive.example.com,resources=skydiveanalyzers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=skydive.example.com,resources=skydiveanalyzers/status,verbs=get;update;patch
func (r *SkydiveAnalyzerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {

	log := r.Log.WithValues("skydiveanalyzer", req.NamespacedName)

	//if err := r.cleanupOwnedResources(ctx, log, &myKind); err != nil {
	//	log.Error(err, "failed to clean up old Deployment resources for this MyKind")
	//	return ctrl.Result{}, err
	//}

	log.Info("initializing skydive-analyzer Deployment...")
	assets := manifests.NewAssets("assets")

	dep, err := kclient.NewDeployment(assets.MustNewAssetReader(SkydiveAnalyzerDeployment))
	if err != nil {
		log.Error(err, "initializing skydive-analyzer Deployment failed")
		return ctrl.Result{}, errors.Wrap(err, "initializing skydive-analyzer Deployment failed")
	}

	log.Info("initializing skydive-analyzer task completed successfully")

	log.Info("Building configuration")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		log.Error(err, "Configuration build has failed")
		return ctrl.Result{}, err
	}

	log.Info("Creating kubernetes client")
	kclient_instance, err := kclient.New(config, "", namespaceName, "")
	if err != nil {
		log.Error(err, "Kubernets client build failed")
		return ctrl.Result{}, err
	}

	log.Info("Creating deployment")
	err = kclient_instance.CreateOrUpdateDeployment(dep)
	if err != nil {
		log.Info("Deployment creation has failed")
		return ctrl.Result{}, err
	}

	log.Info("reconciling skydive-analyzer task completed successfully")
	return ctrl.Result{}, nil

}

func (r *SkydiveAnalyzerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&skydivev1beta1.SkydiveAnalyzer{}).
		Complete(r)
}
