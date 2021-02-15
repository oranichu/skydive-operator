// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SkydiveAgents) DeepCopyInto(out *SkydiveAgents) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SkydiveAgents.
func (in *SkydiveAgents) DeepCopy() *SkydiveAgents {
	if in == nil {
		return nil
	}
	out := new(SkydiveAgents)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SkydiveAgents) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SkydiveAgentsList) DeepCopyInto(out *SkydiveAgentsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SkydiveAgents, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SkydiveAgentsList.
func (in *SkydiveAgentsList) DeepCopy() *SkydiveAgentsList {
	if in == nil {
		return nil
	}
	out := new(SkydiveAgentsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SkydiveAgentsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SkydiveAgentsSpec) DeepCopyInto(out *SkydiveAgentsSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SkydiveAgentsSpec.
func (in *SkydiveAgentsSpec) DeepCopy() *SkydiveAgentsSpec {
	if in == nil {
		return nil
	}
	out := new(SkydiveAgentsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SkydiveAgentsStatus) DeepCopyInto(out *SkydiveAgentsStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SkydiveAgentsStatus.
func (in *SkydiveAgentsStatus) DeepCopy() *SkydiveAgentsStatus {
	if in == nil {
		return nil
	}
	out := new(SkydiveAgentsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SkydiveAnalyzer) DeepCopyInto(out *SkydiveAnalyzer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SkydiveAnalyzer.
func (in *SkydiveAnalyzer) DeepCopy() *SkydiveAnalyzer {
	if in == nil {
		return nil
	}
	out := new(SkydiveAnalyzer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SkydiveAnalyzer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SkydiveAnalyzerList) DeepCopyInto(out *SkydiveAnalyzerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SkydiveAnalyzer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SkydiveAnalyzerList.
func (in *SkydiveAnalyzerList) DeepCopy() *SkydiveAnalyzerList {
	if in == nil {
		return nil
	}
	out := new(SkydiveAnalyzerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SkydiveAnalyzerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SkydiveAnalyzerSpec) DeepCopyInto(out *SkydiveAnalyzerSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SkydiveAnalyzerSpec.
func (in *SkydiveAnalyzerSpec) DeepCopy() *SkydiveAnalyzerSpec {
	if in == nil {
		return nil
	}
	out := new(SkydiveAnalyzerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SkydiveAnalyzerStatus) DeepCopyInto(out *SkydiveAnalyzerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SkydiveAnalyzerStatus.
func (in *SkydiveAnalyzerStatus) DeepCopy() *SkydiveAnalyzerStatus {
	if in == nil {
		return nil
	}
	out := new(SkydiveAnalyzerStatus)
	in.DeepCopyInto(out)
	return out
}
