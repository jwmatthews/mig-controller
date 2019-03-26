// +build !ignore_autogenerated

/*
Copyright 2019 Red Hat Inc.

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
// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupStorageConfig) DeepCopyInto(out *BackupStorageConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupStorageConfig.
func (in *BackupStorageConfig) DeepCopy() *BackupStorageConfig {
	if in == nil {
		return nil
	}
	out := new(BackupStorageConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigAssetCollection) DeepCopyInto(out *MigAssetCollection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigAssetCollection.
func (in *MigAssetCollection) DeepCopy() *MigAssetCollection {
	if in == nil {
		return nil
	}
	out := new(MigAssetCollection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MigAssetCollection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigAssetCollectionList) DeepCopyInto(out *MigAssetCollectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MigAssetCollection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigAssetCollectionList.
func (in *MigAssetCollectionList) DeepCopy() *MigAssetCollectionList {
	if in == nil {
		return nil
	}
	out := new(MigAssetCollectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MigAssetCollectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigAssetCollectionSpec) DeepCopyInto(out *MigAssetCollectionSpec) {
	*out = *in
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigAssetCollectionSpec.
func (in *MigAssetCollectionSpec) DeepCopy() *MigAssetCollectionSpec {
	if in == nil {
		return nil
	}
	out := new(MigAssetCollectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigAssetCollectionStatus) DeepCopyInto(out *MigAssetCollectionStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigAssetCollectionStatus.
func (in *MigAssetCollectionStatus) DeepCopy() *MigAssetCollectionStatus {
	if in == nil {
		return nil
	}
	out := new(MigAssetCollectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigMigration) DeepCopyInto(out *MigMigration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigMigration.
func (in *MigMigration) DeepCopy() *MigMigration {
	if in == nil {
		return nil
	}
	out := new(MigMigration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MigMigration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigMigrationList) DeepCopyInto(out *MigMigrationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MigMigration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigMigrationList.
func (in *MigMigrationList) DeepCopy() *MigMigrationList {
	if in == nil {
		return nil
	}
	out := new(MigMigrationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MigMigrationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigMigrationSpec) DeepCopyInto(out *MigMigrationSpec) {
	*out = *in
	if in.MigStageRef != nil {
		in, out := &in.MigStageRef, &out.MigStageRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigMigrationSpec.
func (in *MigMigrationSpec) DeepCopy() *MigMigrationSpec {
	if in == nil {
		return nil
	}
	out := new(MigMigrationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigMigrationStatus) DeepCopyInto(out *MigMigrationStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigMigrationStatus.
func (in *MigMigrationStatus) DeepCopy() *MigMigrationStatus {
	if in == nil {
		return nil
	}
	out := new(MigMigrationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigPlan) DeepCopyInto(out *MigPlan) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigPlan.
func (in *MigPlan) DeepCopy() *MigPlan {
	if in == nil {
		return nil
	}
	out := new(MigPlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MigPlan) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigPlanList) DeepCopyInto(out *MigPlanList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MigPlan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigPlanList.
func (in *MigPlanList) DeepCopy() *MigPlanList {
	if in == nil {
		return nil
	}
	out := new(MigPlanList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MigPlanList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigPlanSpec) DeepCopyInto(out *MigPlanSpec) {
	*out = *in
	if in.SrcClusterRef != nil {
		in, out := &in.SrcClusterRef, &out.SrcClusterRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.DestClusterRef != nil {
		in, out := &in.DestClusterRef, &out.DestClusterRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.MigStorageRef != nil {
		in, out := &in.MigStorageRef, &out.MigStorageRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.MigAssetCollectionRef != nil {
		in, out := &in.MigAssetCollectionRef, &out.MigAssetCollectionRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigPlanSpec.
func (in *MigPlanSpec) DeepCopy() *MigPlanSpec {
	if in == nil {
		return nil
	}
	out := new(MigPlanSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigPlanStatus) DeepCopyInto(out *MigPlanStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigPlanStatus.
func (in *MigPlanStatus) DeepCopy() *MigPlanStatus {
	if in == nil {
		return nil
	}
	out := new(MigPlanStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigStage) DeepCopyInto(out *MigStage) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigStage.
func (in *MigStage) DeepCopy() *MigStage {
	if in == nil {
		return nil
	}
	out := new(MigStage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MigStage) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigStageList) DeepCopyInto(out *MigStageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MigStage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigStageList.
func (in *MigStageList) DeepCopy() *MigStageList {
	if in == nil {
		return nil
	}
	out := new(MigStageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MigStageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigStageSpec) DeepCopyInto(out *MigStageSpec) {
	*out = *in
	if in.MigPlanRef != nil {
		in, out := &in.MigPlanRef, &out.MigPlanRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigStageSpec.
func (in *MigStageSpec) DeepCopy() *MigStageSpec {
	if in == nil {
		return nil
	}
	out := new(MigStageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigStageStatus) DeepCopyInto(out *MigStageStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigStageStatus.
func (in *MigStageStatus) DeepCopy() *MigStageStatus {
	if in == nil {
		return nil
	}
	out := new(MigStageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigStorage) DeepCopyInto(out *MigStorage) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigStorage.
func (in *MigStorage) DeepCopy() *MigStorage {
	if in == nil {
		return nil
	}
	out := new(MigStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MigStorage) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigStorageList) DeepCopyInto(out *MigStorageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MigStorage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigStorageList.
func (in *MigStorageList) DeepCopy() *MigStorageList {
	if in == nil {
		return nil
	}
	out := new(MigStorageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MigStorageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigStorageSpec) DeepCopyInto(out *MigStorageSpec) {
	*out = *in
	out.VolumeSnapshotConfig = in.VolumeSnapshotConfig
	out.BackupStorageConfig = in.BackupStorageConfig
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigStorageSpec.
func (in *MigStorageSpec) DeepCopy() *MigStorageSpec {
	if in == nil {
		return nil
	}
	out := new(MigStorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigStorageStatus) DeepCopyInto(out *MigStorageStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigStorageStatus.
func (in *MigStorageStatus) DeepCopy() *MigStorageStatus {
	if in == nil {
		return nil
	}
	out := new(MigStorageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeSnapshotConfig) DeepCopyInto(out *VolumeSnapshotConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeSnapshotConfig.
func (in *VolumeSnapshotConfig) DeepCopy() *VolumeSnapshotConfig {
	if in == nil {
		return nil
	}
	out := new(VolumeSnapshotConfig)
	in.DeepCopyInto(out)
	return out
}
