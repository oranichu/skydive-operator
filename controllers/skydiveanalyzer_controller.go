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
	"github.com/pkg/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"skydive/pkg/client"
	"skydive/pkg/manifests"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	skydivev1beta1 "skydive/api/v1beta1"
)

// SkydiveAnalyzerReconciler reconciles a SkydiveAnalyzer object
type SkydiveAnalyzerReconciler struct {
	Log    logr.Logger
	Scheme *runtime.Scheme

	client  *client.Client
	factory *manifests.Factory
}

// +kubebuilder:rbac:groups=skydive.example.com,resources=skydiveanalyzers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=skydive.example.com,resources=skydiveanalyzers/status,verbs=get;update;patch

func (r *SkydiveAnalyzerReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("skydiveanalyzer", req.NamespacedName)

	//if err := r.cleanupOwnedResources(ctx, log, &myKind); err != nil {
	//	log.Error(err, "failed to clean up old Deployment resources for this MyKind")
	//	return ctrl.Result{}, err
	//}

	log.Info("initializing skydive-analyzer Deployment...")

	dep, err := r.factory.SkydiveAnalyzerDeployment()
	if err != nil {
		log.Error(err, "initializing skydive-analyzer Deployment failed")
		return ctrl.Result{}, errors.Wrap(err, "initializing skydive-analyzer Deployment failed")
	}

	log.Info("initializing skydive-analyzer task completed successfully")
	log.Info("reconciling skydive-analyzer Deployment")

	err = r.client.CreateOrUpdateDeployment(dep)
	if err != nil {
		return ctrl.Result{}, errors.Wrap(err, "reconciling skydive-analyzer Deployment failed")
	}

	log.Info("reconciling skydive-analyzer task completed successfully")

	return ctrl.Result{}, nil

}

//func buildDeployment(skydive_analyzer skydivev1beta1.SkydiveAnalyzer) *apps.Deployment {
//	deployment := apps.Deployment{
//		ObjectMeta: metav1.ObjectMeta{
//			Name:      "skydive-analyzer",
//			Namespace: "skydive",
//			OwnerReferences: []metav1.OwnerReference{
//				*metav1.NewControllerRef(
//					&skydive_analyzer,
//					skydivev1beta1.GroupVersion.WithKind("SkydiveAnalyzer"))},
//		},
//		Spec: apps.DeploymentSpec{
//			Replicas: skydive_analyzer.Spec.Replicas,
//			Selector: &metav1.LabelSelector{
//				MatchLabels: map[string]string{
//					"app":  "skydive",
//					"tier": "analyzer",
//				},
//			},
//			Template: core.PodTemplateSpec{
//				ObjectMeta: metav1.ObjectMeta{
//					Labels: map[string]string{
//						"app":  "skydive",
//						"tier": "analyzer"},
//				},
//				Spec: core.PodSpec{
//					Containers: []core.Container{
//						{
//							Name:            "skydive-analyzer",
//							Image:           "skydive/skydive",
//							ImagePullPolicy: core.PullAlways,
//							Args: []string{
//								"analyzer",
//								"--listen=0.0.0.0:8082",
//							},
//							Env: []core.EnvVar{
//								{
//									Name:  "SKYDIVE_ANALYZER_FLOW_BACKEND",
//									Value: "elasticsearch",
//								},
//								{
//									Name:  "SKYDIVE_ANALYZER_TOPOLOGY_BACKEND",
//									Value: "elasticsearch",
//								},
//								{
//									Name:  "SKYDIVE_ANALYZER_TOPOLOGY_PROBES",
//									Value: "k8s ovn",
//								},
//								{
//									Name:  "SKYDIVE_ANALYZER_TOPOLOGY_K8S_PROBES",
//									Value: "cluster namespace node pod container service deployment",
//								},
//								{
//									Name:  "SKYDIVE_UI",
//									Value: "'{\"theme\":\"light\",\"k8s_enabled\":\"true\",\"topology\": {\"favorites\":{\"infrastructure\":\"G.V().Has(\\\"Manager\\\", Without(\\\"k8s\\\"))\",\"kubernetes\":\"G.V().Has(\\\"Manager\\\",\\\"k8s\\\")\"},\"default_filter\":\"infrastructure\"}}'",
//								},
//								{
//									Name:  "SKYDIVE_ANALYZER_TOPOLOGY_FABRIC",
//									Value: "'TOR1->*[Type=host]/eth0 TOR1->*[Type=host]/eth1 TOR1->*[Type=host]/ens1 TOR1->*[Type=host]/ens2 TOR1->*[Type=host]/ens3'",
//								},
//								{
//									Name:  "SKYDIVE_ANALYZER_STARTUP_CAPTURE_GREMLIN",
//									Value: "'G.V().Has(\"Type\" , \"device\", \"Name\", Regex(\"eth0|ens1|ens2|ens3\"))'",
//								},
//								{
//									Name:  "SKYDIVE_ETCD_LISTEN",
//									Value: "0.0.0.0:12379",
//								},
//								{
//									Name:  "SKYDIVE_LOGGING_LEVEL",
//									Value: "${SKYDIVE_LOGGING_LEVEL}", // TODO: what ?
//								},
//							},
//							Ports: []core.ContainerPort{
//								{
//									ContainerPort: 8082,
//									Protocol:      core.ProtocolTCP,
//								},
//								{
//									ContainerPort: 8082,
//									Protocol:      core.ProtocolUDP,
//								},
//								{
//									ContainerPort: 12379,
//									Protocol:      core.ProtocolTCP,
//								},
//								{
//									ContainerPort: 12380,
//									Protocol:      core.ProtocolTCP,
//								},
//							},
//							LivenessProbe: &core.Probe{
//								Handler: core.Handler{
//									TCPSocket: &core.TCPSocketAction{
//										Port: intstr.IntOrString{
//											IntVal: 8082,
//										},
//									},
//								},
//								InitialDelaySeconds: 30,
//								TimeoutSeconds:      5,
//								PeriodSeconds:       10,
//								SuccessThreshold:    1,
//								FailureThreshold:    3,
//							},
//							ReadinessProbe: &core.Probe{
//								Handler: core.Handler{
//									TCPSocket: &core.TCPSocketAction{
//										Port: intstr.IntOrString{
//											IntVal: 8082,
//										},
//									},
//								},
//								InitialDelaySeconds: 30,
//								TimeoutSeconds:      5,
//								PeriodSeconds:       10,
//								SuccessThreshold:    1,
//								FailureThreshold:    1,
//							},
//							SecurityContext: &core.SecurityContext{
//								Privileged: pointer.BoolPtr(true),
//							},
//						},
//					},
//				},
//			},
//		},
//	}
//	return &deployment
//}

func (r *SkydiveAnalyzerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&skydivev1beta1.SkydiveAnalyzer{}).
		Complete(r)
}
