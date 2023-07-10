//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and reactivejob-operator-cop contributors
SPDX-License-Identifier: Apache-2.0
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReactiveJobOperator) DeepCopyInto(out *ReactiveJobOperator) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReactiveJobOperator.
func (in *ReactiveJobOperator) DeepCopy() *ReactiveJobOperator {
	if in == nil {
		return nil
	}
	out := new(ReactiveJobOperator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReactiveJobOperator) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReactiveJobOperatorList) DeepCopyInto(out *ReactiveJobOperatorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ReactiveJobOperator, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReactiveJobOperatorList.
func (in *ReactiveJobOperatorList) DeepCopy() *ReactiveJobOperatorList {
	if in == nil {
		return nil
	}
	out := new(ReactiveJobOperatorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReactiveJobOperatorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReactiveJobOperatorSpec) DeepCopyInto(out *ReactiveJobOperatorSpec) {
	*out = *in
	out.Spec = in.Spec
	out.Image = in.Image
	in.KubernetesProperties.DeepCopyInto(&out.KubernetesProperties)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReactiveJobOperatorSpec.
func (in *ReactiveJobOperatorSpec) DeepCopy() *ReactiveJobOperatorSpec {
	if in == nil {
		return nil
	}
	out := new(ReactiveJobOperatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReactiveJobOperatorStatus) DeepCopyInto(out *ReactiveJobOperatorStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReactiveJobOperatorStatus.
func (in *ReactiveJobOperatorStatus) DeepCopy() *ReactiveJobOperatorStatus {
	if in == nil {
		return nil
	}
	out := new(ReactiveJobOperatorStatus)
	in.DeepCopyInto(out)
	return out
}
