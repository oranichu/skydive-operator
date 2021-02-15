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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SkydiveAnalyzerSpec defines the desired state of SkydiveAnalyzer
type SkydiveAnalyzerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +optional
	// +kubebuilder:default=1
	Replicas *int32 `json:"replicas,omitempty"`
}

// SkydiveAnalyzerStatus defines the observed state of SkydiveAnalyzer
type SkydiveAnalyzerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SkydiveAnalyzer is the Schema for the skydiveanalyzers API
type SkydiveAnalyzer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SkydiveAnalyzerSpec   `json:"spec,omitempty"`
	Status SkydiveAnalyzerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SkydiveAnalyzerList contains a list of SkydiveAnalyzer
type SkydiveAnalyzerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SkydiveAnalyzer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SkydiveAnalyzer{}, &SkydiveAnalyzerList{})
}
