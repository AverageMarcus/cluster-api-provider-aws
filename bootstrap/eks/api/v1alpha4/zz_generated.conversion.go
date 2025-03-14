//go:build !ignore_autogenerated_conversions
// +build !ignore_autogenerated_conversions

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha4

import (
	unsafe "unsafe"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1beta1 "sigs.k8s.io/cluster-api-provider-aws/bootstrap/eks/api/v1beta1"
	apiv1alpha4 "sigs.k8s.io/cluster-api/api/v1alpha4"
	apiv1beta1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*EKSConfig)(nil), (*v1beta1.EKSConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_EKSConfig_To_v1beta1_EKSConfig(a.(*EKSConfig), b.(*v1beta1.EKSConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.EKSConfig)(nil), (*EKSConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_EKSConfig_To_v1alpha4_EKSConfig(a.(*v1beta1.EKSConfig), b.(*EKSConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*EKSConfigList)(nil), (*v1beta1.EKSConfigList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_EKSConfigList_To_v1beta1_EKSConfigList(a.(*EKSConfigList), b.(*v1beta1.EKSConfigList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.EKSConfigList)(nil), (*EKSConfigList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_EKSConfigList_To_v1alpha4_EKSConfigList(a.(*v1beta1.EKSConfigList), b.(*EKSConfigList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*EKSConfigSpec)(nil), (*v1beta1.EKSConfigSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_EKSConfigSpec_To_v1beta1_EKSConfigSpec(a.(*EKSConfigSpec), b.(*v1beta1.EKSConfigSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.EKSConfigSpec)(nil), (*EKSConfigSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_EKSConfigSpec_To_v1alpha4_EKSConfigSpec(a.(*v1beta1.EKSConfigSpec), b.(*EKSConfigSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*EKSConfigStatus)(nil), (*v1beta1.EKSConfigStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_EKSConfigStatus_To_v1beta1_EKSConfigStatus(a.(*EKSConfigStatus), b.(*v1beta1.EKSConfigStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.EKSConfigStatus)(nil), (*EKSConfigStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_EKSConfigStatus_To_v1alpha4_EKSConfigStatus(a.(*v1beta1.EKSConfigStatus), b.(*EKSConfigStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*EKSConfigTemplate)(nil), (*v1beta1.EKSConfigTemplate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_EKSConfigTemplate_To_v1beta1_EKSConfigTemplate(a.(*EKSConfigTemplate), b.(*v1beta1.EKSConfigTemplate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.EKSConfigTemplate)(nil), (*EKSConfigTemplate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_EKSConfigTemplate_To_v1alpha4_EKSConfigTemplate(a.(*v1beta1.EKSConfigTemplate), b.(*EKSConfigTemplate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*EKSConfigTemplateList)(nil), (*v1beta1.EKSConfigTemplateList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_EKSConfigTemplateList_To_v1beta1_EKSConfigTemplateList(a.(*EKSConfigTemplateList), b.(*v1beta1.EKSConfigTemplateList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.EKSConfigTemplateList)(nil), (*EKSConfigTemplateList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_EKSConfigTemplateList_To_v1alpha4_EKSConfigTemplateList(a.(*v1beta1.EKSConfigTemplateList), b.(*EKSConfigTemplateList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*EKSConfigTemplateResource)(nil), (*v1beta1.EKSConfigTemplateResource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_EKSConfigTemplateResource_To_v1beta1_EKSConfigTemplateResource(a.(*EKSConfigTemplateResource), b.(*v1beta1.EKSConfigTemplateResource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.EKSConfigTemplateResource)(nil), (*EKSConfigTemplateResource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_EKSConfigTemplateResource_To_v1alpha4_EKSConfigTemplateResource(a.(*v1beta1.EKSConfigTemplateResource), b.(*EKSConfigTemplateResource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*EKSConfigTemplateSpec)(nil), (*v1beta1.EKSConfigTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_EKSConfigTemplateSpec_To_v1beta1_EKSConfigTemplateSpec(a.(*EKSConfigTemplateSpec), b.(*v1beta1.EKSConfigTemplateSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.EKSConfigTemplateSpec)(nil), (*EKSConfigTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_EKSConfigTemplateSpec_To_v1alpha4_EKSConfigTemplateSpec(a.(*v1beta1.EKSConfigTemplateSpec), b.(*EKSConfigTemplateSpec), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha4_EKSConfig_To_v1beta1_EKSConfig(in *EKSConfig, out *v1beta1.EKSConfig, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha4_EKSConfigSpec_To_v1beta1_EKSConfigSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha4_EKSConfigStatus_To_v1beta1_EKSConfigStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha4_EKSConfig_To_v1beta1_EKSConfig is an autogenerated conversion function.
func Convert_v1alpha4_EKSConfig_To_v1beta1_EKSConfig(in *EKSConfig, out *v1beta1.EKSConfig, s conversion.Scope) error {
	return autoConvert_v1alpha4_EKSConfig_To_v1beta1_EKSConfig(in, out, s)
}

func autoConvert_v1beta1_EKSConfig_To_v1alpha4_EKSConfig(in *v1beta1.EKSConfig, out *EKSConfig, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_EKSConfigSpec_To_v1alpha4_EKSConfigSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_EKSConfigStatus_To_v1alpha4_EKSConfigStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_EKSConfig_To_v1alpha4_EKSConfig is an autogenerated conversion function.
func Convert_v1beta1_EKSConfig_To_v1alpha4_EKSConfig(in *v1beta1.EKSConfig, out *EKSConfig, s conversion.Scope) error {
	return autoConvert_v1beta1_EKSConfig_To_v1alpha4_EKSConfig(in, out, s)
}

func autoConvert_v1alpha4_EKSConfigList_To_v1beta1_EKSConfigList(in *EKSConfigList, out *v1beta1.EKSConfigList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1beta1.EKSConfig, len(*in))
		for i := range *in {
			if err := Convert_v1alpha4_EKSConfig_To_v1beta1_EKSConfig(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha4_EKSConfigList_To_v1beta1_EKSConfigList is an autogenerated conversion function.
func Convert_v1alpha4_EKSConfigList_To_v1beta1_EKSConfigList(in *EKSConfigList, out *v1beta1.EKSConfigList, s conversion.Scope) error {
	return autoConvert_v1alpha4_EKSConfigList_To_v1beta1_EKSConfigList(in, out, s)
}

func autoConvert_v1beta1_EKSConfigList_To_v1alpha4_EKSConfigList(in *v1beta1.EKSConfigList, out *EKSConfigList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EKSConfig, len(*in))
		for i := range *in {
			if err := Convert_v1beta1_EKSConfig_To_v1alpha4_EKSConfig(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1beta1_EKSConfigList_To_v1alpha4_EKSConfigList is an autogenerated conversion function.
func Convert_v1beta1_EKSConfigList_To_v1alpha4_EKSConfigList(in *v1beta1.EKSConfigList, out *EKSConfigList, s conversion.Scope) error {
	return autoConvert_v1beta1_EKSConfigList_To_v1alpha4_EKSConfigList(in, out, s)
}

func autoConvert_v1alpha4_EKSConfigSpec_To_v1beta1_EKSConfigSpec(in *EKSConfigSpec, out *v1beta1.EKSConfigSpec, s conversion.Scope) error {
	out.KubeletExtraArgs = *(*map[string]string)(unsafe.Pointer(&in.KubeletExtraArgs))
	return nil
}

// Convert_v1alpha4_EKSConfigSpec_To_v1beta1_EKSConfigSpec is an autogenerated conversion function.
func Convert_v1alpha4_EKSConfigSpec_To_v1beta1_EKSConfigSpec(in *EKSConfigSpec, out *v1beta1.EKSConfigSpec, s conversion.Scope) error {
	return autoConvert_v1alpha4_EKSConfigSpec_To_v1beta1_EKSConfigSpec(in, out, s)
}

func autoConvert_v1beta1_EKSConfigSpec_To_v1alpha4_EKSConfigSpec(in *v1beta1.EKSConfigSpec, out *EKSConfigSpec, s conversion.Scope) error {
	out.KubeletExtraArgs = *(*map[string]string)(unsafe.Pointer(&in.KubeletExtraArgs))
	return nil
}

// Convert_v1beta1_EKSConfigSpec_To_v1alpha4_EKSConfigSpec is an autogenerated conversion function.
func Convert_v1beta1_EKSConfigSpec_To_v1alpha4_EKSConfigSpec(in *v1beta1.EKSConfigSpec, out *EKSConfigSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_EKSConfigSpec_To_v1alpha4_EKSConfigSpec(in, out, s)
}

func autoConvert_v1alpha4_EKSConfigStatus_To_v1beta1_EKSConfigStatus(in *EKSConfigStatus, out *v1beta1.EKSConfigStatus, s conversion.Scope) error {
	out.Ready = in.Ready
	out.DataSecretName = (*string)(unsafe.Pointer(in.DataSecretName))
	out.FailureReason = in.FailureReason
	out.FailureMessage = in.FailureMessage
	out.ObservedGeneration = in.ObservedGeneration
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(apiv1beta1.Conditions, len(*in))
		for i := range *in {
			if err := apiv1alpha4.Convert_v1alpha4_Condition_To_v1beta1_Condition(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Conditions = nil
	}
	return nil
}

// Convert_v1alpha4_EKSConfigStatus_To_v1beta1_EKSConfigStatus is an autogenerated conversion function.
func Convert_v1alpha4_EKSConfigStatus_To_v1beta1_EKSConfigStatus(in *EKSConfigStatus, out *v1beta1.EKSConfigStatus, s conversion.Scope) error {
	return autoConvert_v1alpha4_EKSConfigStatus_To_v1beta1_EKSConfigStatus(in, out, s)
}

func autoConvert_v1beta1_EKSConfigStatus_To_v1alpha4_EKSConfigStatus(in *v1beta1.EKSConfigStatus, out *EKSConfigStatus, s conversion.Scope) error {
	out.Ready = in.Ready
	out.DataSecretName = (*string)(unsafe.Pointer(in.DataSecretName))
	out.FailureReason = in.FailureReason
	out.FailureMessage = in.FailureMessage
	out.ObservedGeneration = in.ObservedGeneration
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(apiv1alpha4.Conditions, len(*in))
		for i := range *in {
			if err := apiv1alpha4.Convert_v1beta1_Condition_To_v1alpha4_Condition(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Conditions = nil
	}
	return nil
}

// Convert_v1beta1_EKSConfigStatus_To_v1alpha4_EKSConfigStatus is an autogenerated conversion function.
func Convert_v1beta1_EKSConfigStatus_To_v1alpha4_EKSConfigStatus(in *v1beta1.EKSConfigStatus, out *EKSConfigStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_EKSConfigStatus_To_v1alpha4_EKSConfigStatus(in, out, s)
}

func autoConvert_v1alpha4_EKSConfigTemplate_To_v1beta1_EKSConfigTemplate(in *EKSConfigTemplate, out *v1beta1.EKSConfigTemplate, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha4_EKSConfigTemplateSpec_To_v1beta1_EKSConfigTemplateSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha4_EKSConfigTemplate_To_v1beta1_EKSConfigTemplate is an autogenerated conversion function.
func Convert_v1alpha4_EKSConfigTemplate_To_v1beta1_EKSConfigTemplate(in *EKSConfigTemplate, out *v1beta1.EKSConfigTemplate, s conversion.Scope) error {
	return autoConvert_v1alpha4_EKSConfigTemplate_To_v1beta1_EKSConfigTemplate(in, out, s)
}

func autoConvert_v1beta1_EKSConfigTemplate_To_v1alpha4_EKSConfigTemplate(in *v1beta1.EKSConfigTemplate, out *EKSConfigTemplate, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_EKSConfigTemplateSpec_To_v1alpha4_EKSConfigTemplateSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_EKSConfigTemplate_To_v1alpha4_EKSConfigTemplate is an autogenerated conversion function.
func Convert_v1beta1_EKSConfigTemplate_To_v1alpha4_EKSConfigTemplate(in *v1beta1.EKSConfigTemplate, out *EKSConfigTemplate, s conversion.Scope) error {
	return autoConvert_v1beta1_EKSConfigTemplate_To_v1alpha4_EKSConfigTemplate(in, out, s)
}

func autoConvert_v1alpha4_EKSConfigTemplateList_To_v1beta1_EKSConfigTemplateList(in *EKSConfigTemplateList, out *v1beta1.EKSConfigTemplateList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1beta1.EKSConfigTemplate)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha4_EKSConfigTemplateList_To_v1beta1_EKSConfigTemplateList is an autogenerated conversion function.
func Convert_v1alpha4_EKSConfigTemplateList_To_v1beta1_EKSConfigTemplateList(in *EKSConfigTemplateList, out *v1beta1.EKSConfigTemplateList, s conversion.Scope) error {
	return autoConvert_v1alpha4_EKSConfigTemplateList_To_v1beta1_EKSConfigTemplateList(in, out, s)
}

func autoConvert_v1beta1_EKSConfigTemplateList_To_v1alpha4_EKSConfigTemplateList(in *v1beta1.EKSConfigTemplateList, out *EKSConfigTemplateList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]EKSConfigTemplate)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1beta1_EKSConfigTemplateList_To_v1alpha4_EKSConfigTemplateList is an autogenerated conversion function.
func Convert_v1beta1_EKSConfigTemplateList_To_v1alpha4_EKSConfigTemplateList(in *v1beta1.EKSConfigTemplateList, out *EKSConfigTemplateList, s conversion.Scope) error {
	return autoConvert_v1beta1_EKSConfigTemplateList_To_v1alpha4_EKSConfigTemplateList(in, out, s)
}

func autoConvert_v1alpha4_EKSConfigTemplateResource_To_v1beta1_EKSConfigTemplateResource(in *EKSConfigTemplateResource, out *v1beta1.EKSConfigTemplateResource, s conversion.Scope) error {
	if err := Convert_v1alpha4_EKSConfigSpec_To_v1beta1_EKSConfigSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha4_EKSConfigTemplateResource_To_v1beta1_EKSConfigTemplateResource is an autogenerated conversion function.
func Convert_v1alpha4_EKSConfigTemplateResource_To_v1beta1_EKSConfigTemplateResource(in *EKSConfigTemplateResource, out *v1beta1.EKSConfigTemplateResource, s conversion.Scope) error {
	return autoConvert_v1alpha4_EKSConfigTemplateResource_To_v1beta1_EKSConfigTemplateResource(in, out, s)
}

func autoConvert_v1beta1_EKSConfigTemplateResource_To_v1alpha4_EKSConfigTemplateResource(in *v1beta1.EKSConfigTemplateResource, out *EKSConfigTemplateResource, s conversion.Scope) error {
	if err := Convert_v1beta1_EKSConfigSpec_To_v1alpha4_EKSConfigSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_EKSConfigTemplateResource_To_v1alpha4_EKSConfigTemplateResource is an autogenerated conversion function.
func Convert_v1beta1_EKSConfigTemplateResource_To_v1alpha4_EKSConfigTemplateResource(in *v1beta1.EKSConfigTemplateResource, out *EKSConfigTemplateResource, s conversion.Scope) error {
	return autoConvert_v1beta1_EKSConfigTemplateResource_To_v1alpha4_EKSConfigTemplateResource(in, out, s)
}

func autoConvert_v1alpha4_EKSConfigTemplateSpec_To_v1beta1_EKSConfigTemplateSpec(in *EKSConfigTemplateSpec, out *v1beta1.EKSConfigTemplateSpec, s conversion.Scope) error {
	if err := Convert_v1alpha4_EKSConfigTemplateResource_To_v1beta1_EKSConfigTemplateResource(&in.Template, &out.Template, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha4_EKSConfigTemplateSpec_To_v1beta1_EKSConfigTemplateSpec is an autogenerated conversion function.
func Convert_v1alpha4_EKSConfigTemplateSpec_To_v1beta1_EKSConfigTemplateSpec(in *EKSConfigTemplateSpec, out *v1beta1.EKSConfigTemplateSpec, s conversion.Scope) error {
	return autoConvert_v1alpha4_EKSConfigTemplateSpec_To_v1beta1_EKSConfigTemplateSpec(in, out, s)
}

func autoConvert_v1beta1_EKSConfigTemplateSpec_To_v1alpha4_EKSConfigTemplateSpec(in *v1beta1.EKSConfigTemplateSpec, out *EKSConfigTemplateSpec, s conversion.Scope) error {
	if err := Convert_v1beta1_EKSConfigTemplateResource_To_v1alpha4_EKSConfigTemplateResource(&in.Template, &out.Template, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_EKSConfigTemplateSpec_To_v1alpha4_EKSConfigTemplateSpec is an autogenerated conversion function.
func Convert_v1beta1_EKSConfigTemplateSpec_To_v1alpha4_EKSConfigTemplateSpec(in *v1beta1.EKSConfigTemplateSpec, out *EKSConfigTemplateSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_EKSConfigTemplateSpec_To_v1alpha4_EKSConfigTemplateSpec(in, out, s)
}
