// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Account) DeepCopyInto(out *Account) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Account.
func (in *Account) DeepCopy() *Account {
	if in == nil {
		return nil
	}
	out := new(Account)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Account) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountClaim) DeepCopyInto(out *AccountClaim) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountClaim.
func (in *AccountClaim) DeepCopy() *AccountClaim {
	if in == nil {
		return nil
	}
	out := new(AccountClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccountClaim) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountClaimCondition) DeepCopyInto(out *AccountClaimCondition) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountClaimCondition.
func (in *AccountClaimCondition) DeepCopy() *AccountClaimCondition {
	if in == nil {
		return nil
	}
	out := new(AccountClaimCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountClaimList) DeepCopyInto(out *AccountClaimList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AccountClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountClaimList.
func (in *AccountClaimList) DeepCopy() *AccountClaimList {
	if in == nil {
		return nil
	}
	out := new(AccountClaimList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccountClaimList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountClaimSpec) DeepCopyInto(out *AccountClaimSpec) {
	*out = *in
	out.LegalEntity = in.LegalEntity
	out.AwsCredentialSecret = in.AwsCredentialSecret
	in.Aws.DeepCopyInto(&out.Aws)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountClaimSpec.
func (in *AccountClaimSpec) DeepCopy() *AccountClaimSpec {
	if in == nil {
		return nil
	}
	out := new(AccountClaimSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountClaimStatus) DeepCopyInto(out *AccountClaimStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]AccountClaimCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountClaimStatus.
func (in *AccountClaimStatus) DeepCopy() *AccountClaimStatus {
	if in == nil {
		return nil
	}
	out := new(AccountClaimStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountCondition) DeepCopyInto(out *AccountCondition) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountCondition.
func (in *AccountCondition) DeepCopy() *AccountCondition {
	if in == nil {
		return nil
	}
	out := new(AccountCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountList) DeepCopyInto(out *AccountList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Account, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountList.
func (in *AccountList) DeepCopy() *AccountList {
	if in == nil {
		return nil
	}
	out := new(AccountList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccountList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountPool) DeepCopyInto(out *AccountPool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountPool.
func (in *AccountPool) DeepCopy() *AccountPool {
	if in == nil {
		return nil
	}
	out := new(AccountPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccountPool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountPoolList) DeepCopyInto(out *AccountPoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AccountPool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountPoolList.
func (in *AccountPoolList) DeepCopy() *AccountPoolList {
	if in == nil {
		return nil
	}
	out := new(AccountPoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccountPoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountPoolSpec) DeepCopyInto(out *AccountPoolSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountPoolSpec.
func (in *AccountPoolSpec) DeepCopy() *AccountPoolSpec {
	if in == nil {
		return nil
	}
	out := new(AccountPoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountPoolStatus) DeepCopyInto(out *AccountPoolStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountPoolStatus.
func (in *AccountPoolStatus) DeepCopy() *AccountPoolStatus {
	if in == nil {
		return nil
	}
	out := new(AccountPoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountSpec) DeepCopyInto(out *AccountSpec) {
	*out = *in
	out.LegalEntity = in.LegalEntity
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountSpec.
func (in *AccountSpec) DeepCopy() *AccountSpec {
	if in == nil {
		return nil
	}
	out := new(AccountSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountStatus) DeepCopyInto(out *AccountStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]AccountCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountStatus.
func (in *AccountStatus) DeepCopy() *AccountStatus {
	if in == nil {
		return nil
	}
	out := new(AccountStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Aws) DeepCopyInto(out *Aws) {
	*out = *in
	if in.Regions != nil {
		in, out := &in.Regions, &out.Regions
		*out = make([]AwsRegions, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Aws.
func (in *Aws) DeepCopy() *Aws {
	if in == nil {
		return nil
	}
	out := new(Aws)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsCredentialSecret) DeepCopyInto(out *AwsCredentialSecret) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsCredentialSecret.
func (in *AwsCredentialSecret) DeepCopy() *AwsCredentialSecret {
	if in == nil {
		return nil
	}
	out := new(AwsCredentialSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsRegions) DeepCopyInto(out *AwsRegions) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsRegions.
func (in *AwsRegions) DeepCopy() *AwsRegions {
	if in == nil {
		return nil
	}
	out := new(AwsRegions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LegalEntity) DeepCopyInto(out *LegalEntity) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LegalEntity.
func (in *LegalEntity) DeepCopy() *LegalEntity {
	if in == nil {
		return nil
	}
	out := new(LegalEntity)
	in.DeepCopyInto(out)
	return out
}
