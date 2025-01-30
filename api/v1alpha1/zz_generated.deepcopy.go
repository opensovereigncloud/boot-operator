//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPBootConfig) DeepCopyInto(out *HTTPBootConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPBootConfig.
func (in *HTTPBootConfig) DeepCopy() *HTTPBootConfig {
	if in == nil {
		return nil
	}
	out := new(HTTPBootConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HTTPBootConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPBootConfigList) DeepCopyInto(out *HTTPBootConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HTTPBootConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPBootConfigList.
func (in *HTTPBootConfigList) DeepCopy() *HTTPBootConfigList {
	if in == nil {
		return nil
	}
	out := new(HTTPBootConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HTTPBootConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPBootConfigSpec) DeepCopyInto(out *HTTPBootConfigSpec) {
	*out = *in
	if in.IgnitionSecretRef != nil {
		in, out := &in.IgnitionSecretRef, &out.IgnitionSecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.SystemIPs != nil {
		in, out := &in.SystemIPs, &out.SystemIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPBootConfigSpec.
func (in *HTTPBootConfigSpec) DeepCopy() *HTTPBootConfigSpec {
	if in == nil {
		return nil
	}
	out := new(HTTPBootConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPBootConfigStatus) DeepCopyInto(out *HTTPBootConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPBootConfigStatus.
func (in *HTTPBootConfigStatus) DeepCopy() *HTTPBootConfigStatus {
	if in == nil {
		return nil
	}
	out := new(HTTPBootConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPXEBootConfig) DeepCopyInto(out *IPXEBootConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPXEBootConfig.
func (in *IPXEBootConfig) DeepCopy() *IPXEBootConfig {
	if in == nil {
		return nil
	}
	out := new(IPXEBootConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IPXEBootConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPXEBootConfigList) DeepCopyInto(out *IPXEBootConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IPXEBootConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPXEBootConfigList.
func (in *IPXEBootConfigList) DeepCopy() *IPXEBootConfigList {
	if in == nil {
		return nil
	}
	out := new(IPXEBootConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IPXEBootConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPXEBootConfigSpec) DeepCopyInto(out *IPXEBootConfigSpec) {
	*out = *in
	if in.SystemIPs != nil {
		in, out := &in.SystemIPs, &out.SystemIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IgnitionSecretRef != nil {
		in, out := &in.IgnitionSecretRef, &out.IgnitionSecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.IPXEScriptSecretRef != nil {
		in, out := &in.IPXEScriptSecretRef, &out.IPXEScriptSecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPXEBootConfigSpec.
func (in *IPXEBootConfigSpec) DeepCopy() *IPXEBootConfigSpec {
	if in == nil {
		return nil
	}
	out := new(IPXEBootConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPXEBootConfigStatus) DeepCopyInto(out *IPXEBootConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPXEBootConfigStatus.
func (in *IPXEBootConfigStatus) DeepCopy() *IPXEBootConfigStatus {
	if in == nil {
		return nil
	}
	out := new(IPXEBootConfigStatus)
	in.DeepCopyInto(out)
	return out
}
