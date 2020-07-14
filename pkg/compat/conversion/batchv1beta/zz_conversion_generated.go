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
// Code generated by conversion-gen-linux. DO NOT EDIT.

package batchv1beta

import (
	unsafe "unsafe"

	v1beta1 "k8s.io/api/batch/v1beta1"
	v2alpha1 "k8s.io/api/batch/v2alpha1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*v1beta1.CronJob)(nil), (*v2alpha1.CronJob)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_CronJob_To_v2alpha1_CronJob(a.(*v1beta1.CronJob), b.(*v2alpha1.CronJob), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v2alpha1.CronJob)(nil), (*v1beta1.CronJob)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2alpha1_CronJob_To_v1beta1_CronJob(a.(*v2alpha1.CronJob), b.(*v1beta1.CronJob), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.CronJobList)(nil), (*v2alpha1.CronJobList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_CronJobList_To_v2alpha1_CronJobList(a.(*v1beta1.CronJobList), b.(*v2alpha1.CronJobList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v2alpha1.CronJobList)(nil), (*v1beta1.CronJobList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2alpha1_CronJobList_To_v1beta1_CronJobList(a.(*v2alpha1.CronJobList), b.(*v1beta1.CronJobList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.CronJobSpec)(nil), (*v2alpha1.CronJobSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_CronJobSpec_To_v2alpha1_CronJobSpec(a.(*v1beta1.CronJobSpec), b.(*v2alpha1.CronJobSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v2alpha1.CronJobSpec)(nil), (*v1beta1.CronJobSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2alpha1_CronJobSpec_To_v1beta1_CronJobSpec(a.(*v2alpha1.CronJobSpec), b.(*v1beta1.CronJobSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.CronJobStatus)(nil), (*v2alpha1.CronJobStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_CronJobStatus_To_v2alpha1_CronJobStatus(a.(*v1beta1.CronJobStatus), b.(*v2alpha1.CronJobStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v2alpha1.CronJobStatus)(nil), (*v1beta1.CronJobStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2alpha1_CronJobStatus_To_v1beta1_CronJobStatus(a.(*v2alpha1.CronJobStatus), b.(*v1beta1.CronJobStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.JobTemplate)(nil), (*v2alpha1.JobTemplate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_JobTemplate_To_v2alpha1_JobTemplate(a.(*v1beta1.JobTemplate), b.(*v2alpha1.JobTemplate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v2alpha1.JobTemplate)(nil), (*v1beta1.JobTemplate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2alpha1_JobTemplate_To_v1beta1_JobTemplate(a.(*v2alpha1.JobTemplate), b.(*v1beta1.JobTemplate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.JobTemplateSpec)(nil), (*v2alpha1.JobTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_JobTemplateSpec_To_v2alpha1_JobTemplateSpec(a.(*v1beta1.JobTemplateSpec), b.(*v2alpha1.JobTemplateSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v2alpha1.JobTemplateSpec)(nil), (*v1beta1.JobTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v2alpha1_JobTemplateSpec_To_v1beta1_JobTemplateSpec(a.(*v2alpha1.JobTemplateSpec), b.(*v1beta1.JobTemplateSpec), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1beta1_CronJob_To_v2alpha1_CronJob(in *v1beta1.CronJob, out *v2alpha1.CronJob, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_CronJobSpec_To_v2alpha1_CronJobSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_CronJobStatus_To_v2alpha1_CronJobStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_CronJob_To_v2alpha1_CronJob is an autogenerated conversion function.
func Convert_v1beta1_CronJob_To_v2alpha1_CronJob(in *v1beta1.CronJob, out *v2alpha1.CronJob, s conversion.Scope) error {
	return autoConvert_v1beta1_CronJob_To_v2alpha1_CronJob(in, out, s)
}

func autoConvert_v2alpha1_CronJob_To_v1beta1_CronJob(in *v2alpha1.CronJob, out *v1beta1.CronJob, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v2alpha1_CronJobSpec_To_v1beta1_CronJobSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v2alpha1_CronJobStatus_To_v1beta1_CronJobStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v2alpha1_CronJob_To_v1beta1_CronJob is an autogenerated conversion function.
func Convert_v2alpha1_CronJob_To_v1beta1_CronJob(in *v2alpha1.CronJob, out *v1beta1.CronJob, s conversion.Scope) error {
	return autoConvert_v2alpha1_CronJob_To_v1beta1_CronJob(in, out, s)
}

func autoConvert_v1beta1_CronJobList_To_v2alpha1_CronJobList(in *v1beta1.CronJobList, out *v2alpha1.CronJobList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v2alpha1.CronJob)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1beta1_CronJobList_To_v2alpha1_CronJobList is an autogenerated conversion function.
func Convert_v1beta1_CronJobList_To_v2alpha1_CronJobList(in *v1beta1.CronJobList, out *v2alpha1.CronJobList, s conversion.Scope) error {
	return autoConvert_v1beta1_CronJobList_To_v2alpha1_CronJobList(in, out, s)
}

func autoConvert_v2alpha1_CronJobList_To_v1beta1_CronJobList(in *v2alpha1.CronJobList, out *v1beta1.CronJobList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1beta1.CronJob)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v2alpha1_CronJobList_To_v1beta1_CronJobList is an autogenerated conversion function.
func Convert_v2alpha1_CronJobList_To_v1beta1_CronJobList(in *v2alpha1.CronJobList, out *v1beta1.CronJobList, s conversion.Scope) error {
	return autoConvert_v2alpha1_CronJobList_To_v1beta1_CronJobList(in, out, s)
}

func autoConvert_v1beta1_CronJobSpec_To_v2alpha1_CronJobSpec(in *v1beta1.CronJobSpec, out *v2alpha1.CronJobSpec, s conversion.Scope) error {
	out.Schedule = in.Schedule
	out.StartingDeadlineSeconds = (*int64)(unsafe.Pointer(in.StartingDeadlineSeconds))
	out.ConcurrencyPolicy = v2alpha1.ConcurrencyPolicy(in.ConcurrencyPolicy)
	out.Suspend = (*bool)(unsafe.Pointer(in.Suspend))
	if err := Convert_v1beta1_JobTemplateSpec_To_v2alpha1_JobTemplateSpec(&in.JobTemplate, &out.JobTemplate, s); err != nil {
		return err
	}
	out.SuccessfulJobsHistoryLimit = (*int32)(unsafe.Pointer(in.SuccessfulJobsHistoryLimit))
	out.FailedJobsHistoryLimit = (*int32)(unsafe.Pointer(in.FailedJobsHistoryLimit))
	return nil
}

// Convert_v1beta1_CronJobSpec_To_v2alpha1_CronJobSpec is an autogenerated conversion function.
func Convert_v1beta1_CronJobSpec_To_v2alpha1_CronJobSpec(in *v1beta1.CronJobSpec, out *v2alpha1.CronJobSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_CronJobSpec_To_v2alpha1_CronJobSpec(in, out, s)
}

func autoConvert_v2alpha1_CronJobSpec_To_v1beta1_CronJobSpec(in *v2alpha1.CronJobSpec, out *v1beta1.CronJobSpec, s conversion.Scope) error {
	out.Schedule = in.Schedule
	out.StartingDeadlineSeconds = (*int64)(unsafe.Pointer(in.StartingDeadlineSeconds))
	out.ConcurrencyPolicy = v1beta1.ConcurrencyPolicy(in.ConcurrencyPolicy)
	out.Suspend = (*bool)(unsafe.Pointer(in.Suspend))
	if err := Convert_v2alpha1_JobTemplateSpec_To_v1beta1_JobTemplateSpec(&in.JobTemplate, &out.JobTemplate, s); err != nil {
		return err
	}
	out.SuccessfulJobsHistoryLimit = (*int32)(unsafe.Pointer(in.SuccessfulJobsHistoryLimit))
	out.FailedJobsHistoryLimit = (*int32)(unsafe.Pointer(in.FailedJobsHistoryLimit))
	return nil
}

// Convert_v2alpha1_CronJobSpec_To_v1beta1_CronJobSpec is an autogenerated conversion function.
func Convert_v2alpha1_CronJobSpec_To_v1beta1_CronJobSpec(in *v2alpha1.CronJobSpec, out *v1beta1.CronJobSpec, s conversion.Scope) error {
	return autoConvert_v2alpha1_CronJobSpec_To_v1beta1_CronJobSpec(in, out, s)
}

func autoConvert_v1beta1_CronJobStatus_To_v2alpha1_CronJobStatus(in *v1beta1.CronJobStatus, out *v2alpha1.CronJobStatus, s conversion.Scope) error {
	out.Active = *(*[]v1.ObjectReference)(unsafe.Pointer(&in.Active))
	out.LastScheduleTime = (*metav1.Time)(unsafe.Pointer(in.LastScheduleTime))
	return nil
}

// Convert_v1beta1_CronJobStatus_To_v2alpha1_CronJobStatus is an autogenerated conversion function.
func Convert_v1beta1_CronJobStatus_To_v2alpha1_CronJobStatus(in *v1beta1.CronJobStatus, out *v2alpha1.CronJobStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_CronJobStatus_To_v2alpha1_CronJobStatus(in, out, s)
}

func autoConvert_v2alpha1_CronJobStatus_To_v1beta1_CronJobStatus(in *v2alpha1.CronJobStatus, out *v1beta1.CronJobStatus, s conversion.Scope) error {
	out.Active = *(*[]v1.ObjectReference)(unsafe.Pointer(&in.Active))
	out.LastScheduleTime = (*metav1.Time)(unsafe.Pointer(in.LastScheduleTime))
	return nil
}

// Convert_v2alpha1_CronJobStatus_To_v1beta1_CronJobStatus is an autogenerated conversion function.
func Convert_v2alpha1_CronJobStatus_To_v1beta1_CronJobStatus(in *v2alpha1.CronJobStatus, out *v1beta1.CronJobStatus, s conversion.Scope) error {
	return autoConvert_v2alpha1_CronJobStatus_To_v1beta1_CronJobStatus(in, out, s)
}

func autoConvert_v1beta1_JobTemplate_To_v2alpha1_JobTemplate(in *v1beta1.JobTemplate, out *v2alpha1.JobTemplate, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_JobTemplateSpec_To_v2alpha1_JobTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_JobTemplate_To_v2alpha1_JobTemplate is an autogenerated conversion function.
func Convert_v1beta1_JobTemplate_To_v2alpha1_JobTemplate(in *v1beta1.JobTemplate, out *v2alpha1.JobTemplate, s conversion.Scope) error {
	return autoConvert_v1beta1_JobTemplate_To_v2alpha1_JobTemplate(in, out, s)
}

func autoConvert_v2alpha1_JobTemplate_To_v1beta1_JobTemplate(in *v2alpha1.JobTemplate, out *v1beta1.JobTemplate, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v2alpha1_JobTemplateSpec_To_v1beta1_JobTemplateSpec(&in.Template, &out.Template, s); err != nil {
		return err
	}
	return nil
}

// Convert_v2alpha1_JobTemplate_To_v1beta1_JobTemplate is an autogenerated conversion function.
func Convert_v2alpha1_JobTemplate_To_v1beta1_JobTemplate(in *v2alpha1.JobTemplate, out *v1beta1.JobTemplate, s conversion.Scope) error {
	return autoConvert_v2alpha1_JobTemplate_To_v1beta1_JobTemplate(in, out, s)
}

func autoConvert_v1beta1_JobTemplateSpec_To_v2alpha1_JobTemplateSpec(in *v1beta1.JobTemplateSpec, out *v2alpha1.JobTemplateSpec, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Spec = in.Spec
	return nil
}

// Convert_v1beta1_JobTemplateSpec_To_v2alpha1_JobTemplateSpec is an autogenerated conversion function.
func Convert_v1beta1_JobTemplateSpec_To_v2alpha1_JobTemplateSpec(in *v1beta1.JobTemplateSpec, out *v2alpha1.JobTemplateSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_JobTemplateSpec_To_v2alpha1_JobTemplateSpec(in, out, s)
}

func autoConvert_v2alpha1_JobTemplateSpec_To_v1beta1_JobTemplateSpec(in *v2alpha1.JobTemplateSpec, out *v1beta1.JobTemplateSpec, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Spec = in.Spec
	return nil
}

// Convert_v2alpha1_JobTemplateSpec_To_v1beta1_JobTemplateSpec is an autogenerated conversion function.
func Convert_v2alpha1_JobTemplateSpec_To_v1beta1_JobTemplateSpec(in *v2alpha1.JobTemplateSpec, out *v1beta1.JobTemplateSpec, s conversion.Scope) error {
	return autoConvert_v2alpha1_JobTemplateSpec_To_v1beta1_JobTemplateSpec(in, out, s)
}