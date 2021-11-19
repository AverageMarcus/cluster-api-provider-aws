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

package v1alpha2

import (
	time "time"
	unsafe "unsafe"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1beta1 "sigs.k8s.io/cluster-api-provider-aws/api/v1beta1"
	apiv1alpha2 "sigs.k8s.io/cluster-api/api/v1alpha2"
	apiv1beta1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*AWSCluster)(nil), (*v1beta1.AWSCluster)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_AWSCluster_To_v1beta1_AWSCluster(a.(*AWSCluster), b.(*v1beta1.AWSCluster), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.AWSCluster)(nil), (*AWSCluster)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_AWSCluster_To_v1alpha2_AWSCluster(a.(*v1beta1.AWSCluster), b.(*AWSCluster), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AWSClusterList)(nil), (*v1beta1.AWSClusterList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_AWSClusterList_To_v1beta1_AWSClusterList(a.(*AWSClusterList), b.(*v1beta1.AWSClusterList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.AWSClusterList)(nil), (*AWSClusterList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_AWSClusterList_To_v1alpha2_AWSClusterList(a.(*v1beta1.AWSClusterList), b.(*AWSClusterList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AWSClusterSpec)(nil), (*v1beta1.AWSClusterSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_AWSClusterSpec_To_v1beta1_AWSClusterSpec(a.(*AWSClusterSpec), b.(*v1beta1.AWSClusterSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.AWSClusterSpec)(nil), (*AWSClusterSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_AWSClusterSpec_To_v1alpha2_AWSClusterSpec(a.(*v1beta1.AWSClusterSpec), b.(*AWSClusterSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AWSClusterStatus)(nil), (*v1beta1.AWSClusterStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_AWSClusterStatus_To_v1beta1_AWSClusterStatus(a.(*AWSClusterStatus), b.(*v1beta1.AWSClusterStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.AWSClusterStatus)(nil), (*AWSClusterStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_AWSClusterStatus_To_v1alpha2_AWSClusterStatus(a.(*v1beta1.AWSClusterStatus), b.(*AWSClusterStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AWSLoadBalancerSpec)(nil), (*v1beta1.AWSLoadBalancerSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_AWSLoadBalancerSpec_To_v1beta1_AWSLoadBalancerSpec(a.(*AWSLoadBalancerSpec), b.(*v1beta1.AWSLoadBalancerSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.AWSLoadBalancerSpec)(nil), (*AWSLoadBalancerSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_AWSLoadBalancerSpec_To_v1alpha2_AWSLoadBalancerSpec(a.(*v1beta1.AWSLoadBalancerSpec), b.(*AWSLoadBalancerSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AWSMachine)(nil), (*v1beta1.AWSMachine)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_AWSMachine_To_v1beta1_AWSMachine(a.(*AWSMachine), b.(*v1beta1.AWSMachine), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.AWSMachine)(nil), (*AWSMachine)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_AWSMachine_To_v1alpha2_AWSMachine(a.(*v1beta1.AWSMachine), b.(*AWSMachine), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AWSMachineList)(nil), (*v1beta1.AWSMachineList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_AWSMachineList_To_v1beta1_AWSMachineList(a.(*AWSMachineList), b.(*v1beta1.AWSMachineList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.AWSMachineList)(nil), (*AWSMachineList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_AWSMachineList_To_v1alpha2_AWSMachineList(a.(*v1beta1.AWSMachineList), b.(*AWSMachineList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AWSMachineSpec)(nil), (*v1beta1.AWSMachineSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_AWSMachineSpec_To_v1beta1_AWSMachineSpec(a.(*AWSMachineSpec), b.(*v1beta1.AWSMachineSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.AWSMachineSpec)(nil), (*AWSMachineSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_AWSMachineSpec_To_v1alpha2_AWSMachineSpec(a.(*v1beta1.AWSMachineSpec), b.(*AWSMachineSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AWSMachineStatus)(nil), (*v1beta1.AWSMachineStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_AWSMachineStatus_To_v1beta1_AWSMachineStatus(a.(*AWSMachineStatus), b.(*v1beta1.AWSMachineStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.AWSMachineStatus)(nil), (*AWSMachineStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_AWSMachineStatus_To_v1alpha2_AWSMachineStatus(a.(*v1beta1.AWSMachineStatus), b.(*AWSMachineStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AWSMachineTemplate)(nil), (*v1beta1.AWSMachineTemplate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_AWSMachineTemplate_To_v1beta1_AWSMachineTemplate(a.(*AWSMachineTemplate), b.(*v1beta1.AWSMachineTemplate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.AWSMachineTemplate)(nil), (*AWSMachineTemplate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_AWSMachineTemplate_To_v1alpha2_AWSMachineTemplate(a.(*v1beta1.AWSMachineTemplate), b.(*AWSMachineTemplate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AWSMachineTemplateList)(nil), (*v1beta1.AWSMachineTemplateList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_AWSMachineTemplateList_To_v1beta1_AWSMachineTemplateList(a.(*AWSMachineTemplateList), b.(*v1beta1.AWSMachineTemplateList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.AWSMachineTemplateList)(nil), (*AWSMachineTemplateList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_AWSMachineTemplateList_To_v1alpha2_AWSMachineTemplateList(a.(*v1beta1.AWSMachineTemplateList), b.(*AWSMachineTemplateList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AWSMachineTemplateResource)(nil), (*v1beta1.AWSMachineTemplateResource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_AWSMachineTemplateResource_To_v1beta1_AWSMachineTemplateResource(a.(*AWSMachineTemplateResource), b.(*v1beta1.AWSMachineTemplateResource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.AWSMachineTemplateResource)(nil), (*AWSMachineTemplateResource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return autoConvert_v1beta1_AWSMachineTemplateResource_To_v1alpha2_AWSMachineTemplateResource(a.(*v1beta1.AWSMachineTemplateResource), b.(*AWSMachineTemplateResource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AWSMachineTemplateSpec)(nil), (*v1beta1.AWSMachineTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_AWSMachineTemplateSpec_To_v1beta1_AWSMachineTemplateSpec(a.(*AWSMachineTemplateSpec), b.(*v1beta1.AWSMachineTemplateSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.AWSMachineTemplateSpec)(nil), (*AWSMachineTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_AWSMachineTemplateSpec_To_v1alpha2_AWSMachineTemplateSpec(a.(*v1beta1.AWSMachineTemplateSpec), b.(*AWSMachineTemplateSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AWSResourceReference)(nil), (*v1beta1.AWSResourceReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_AWSResourceReference_To_v1beta1_AWSResourceReference(a.(*AWSResourceReference), b.(*v1beta1.AWSResourceReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.AWSResourceReference)(nil), (*AWSResourceReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_AWSResourceReference_To_v1alpha2_AWSResourceReference(a.(*v1beta1.AWSResourceReference), b.(*AWSResourceReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*BuildParams)(nil), (*v1beta1.BuildParams)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_BuildParams_To_v1beta1_BuildParams(a.(*BuildParams), b.(*v1beta1.BuildParams), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.BuildParams)(nil), (*BuildParams)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_BuildParams_To_v1alpha2_BuildParams(a.(*v1beta1.BuildParams), b.(*BuildParams), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClassicELB)(nil), (*v1beta1.ClassicELB)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_ClassicELB_To_v1beta1_ClassicELB(a.(*ClassicELB), b.(*v1beta1.ClassicELB), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.ClassicELB)(nil), (*ClassicELB)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ClassicELB_To_v1alpha2_ClassicELB(a.(*v1beta1.ClassicELB), b.(*ClassicELB), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClassicELBAttributes)(nil), (*v1beta1.ClassicELBAttributes)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_ClassicELBAttributes_To_v1beta1_ClassicELBAttributes(a.(*ClassicELBAttributes), b.(*v1beta1.ClassicELBAttributes), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.ClassicELBAttributes)(nil), (*ClassicELBAttributes)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ClassicELBAttributes_To_v1alpha2_ClassicELBAttributes(a.(*v1beta1.ClassicELBAttributes), b.(*ClassicELBAttributes), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClassicELBHealthCheck)(nil), (*v1beta1.ClassicELBHealthCheck)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_ClassicELBHealthCheck_To_v1beta1_ClassicELBHealthCheck(a.(*ClassicELBHealthCheck), b.(*v1beta1.ClassicELBHealthCheck), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.ClassicELBHealthCheck)(nil), (*ClassicELBHealthCheck)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ClassicELBHealthCheck_To_v1alpha2_ClassicELBHealthCheck(a.(*v1beta1.ClassicELBHealthCheck), b.(*ClassicELBHealthCheck), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClassicELBListener)(nil), (*v1beta1.ClassicELBListener)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_ClassicELBListener_To_v1beta1_ClassicELBListener(a.(*ClassicELBListener), b.(*v1beta1.ClassicELBListener), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.ClassicELBListener)(nil), (*ClassicELBListener)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ClassicELBListener_To_v1alpha2_ClassicELBListener(a.(*v1beta1.ClassicELBListener), b.(*ClassicELBListener), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*CloudInit)(nil), (*v1beta1.CloudInit)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_CloudInit_To_v1beta1_CloudInit(a.(*CloudInit), b.(*v1beta1.CloudInit), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.CloudInit)(nil), (*CloudInit)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_CloudInit_To_v1alpha2_CloudInit(a.(*v1beta1.CloudInit), b.(*CloudInit), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Filter)(nil), (*v1beta1.Filter)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_Filter_To_v1beta1_Filter(a.(*Filter), b.(*v1beta1.Filter), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.Filter)(nil), (*Filter)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_Filter_To_v1alpha2_Filter(a.(*v1beta1.Filter), b.(*Filter), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*IngressRule)(nil), (*v1beta1.IngressRule)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_IngressRule_To_v1beta1_IngressRule(a.(*IngressRule), b.(*v1beta1.IngressRule), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.IngressRule)(nil), (*IngressRule)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_IngressRule_To_v1alpha2_IngressRule(a.(*v1beta1.IngressRule), b.(*IngressRule), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Instance)(nil), (*v1beta1.Instance)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_Instance_To_v1beta1_Instance(a.(*Instance), b.(*v1beta1.Instance), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.Instance)(nil), (*Instance)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_Instance_To_v1alpha2_Instance(a.(*v1beta1.Instance), b.(*Instance), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*NetworkSpec)(nil), (*v1beta1.NetworkSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_NetworkSpec_To_v1beta1_NetworkSpec(a.(*NetworkSpec), b.(*v1beta1.NetworkSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.NetworkSpec)(nil), (*NetworkSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_NetworkSpec_To_v1alpha2_NetworkSpec(a.(*v1beta1.NetworkSpec), b.(*NetworkSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RouteTable)(nil), (*v1beta1.RouteTable)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_RouteTable_To_v1beta1_RouteTable(a.(*RouteTable), b.(*v1beta1.RouteTable), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.RouteTable)(nil), (*RouteTable)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_RouteTable_To_v1alpha2_RouteTable(a.(*v1beta1.RouteTable), b.(*RouteTable), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*SecurityGroup)(nil), (*v1beta1.SecurityGroup)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_SecurityGroup_To_v1beta1_SecurityGroup(a.(*SecurityGroup), b.(*v1beta1.SecurityGroup), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.SecurityGroup)(nil), (*SecurityGroup)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_SecurityGroup_To_v1alpha2_SecurityGroup(a.(*v1beta1.SecurityGroup), b.(*SecurityGroup), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*SubnetSpec)(nil), (*v1beta1.SubnetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_SubnetSpec_To_v1beta1_SubnetSpec(a.(*SubnetSpec), b.(*v1beta1.SubnetSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.SubnetSpec)(nil), (*SubnetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_SubnetSpec_To_v1alpha2_SubnetSpec(a.(*v1beta1.SubnetSpec), b.(*SubnetSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VPCSpec)(nil), (*v1beta1.VPCSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_VPCSpec_To_v1beta1_VPCSpec(a.(*VPCSpec), b.(*v1beta1.VPCSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.VPCSpec)(nil), (*VPCSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_VPCSpec_To_v1alpha2_VPCSpec(a.(*v1beta1.VPCSpec), b.(*VPCSpec), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha2_AWSCluster_To_v1beta1_AWSCluster(in *AWSCluster, out *v1beta1.AWSCluster, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := autoConvert_v1alpha2_AWSClusterSpec_To_v1beta1_AWSClusterSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := autoConvert_v1alpha2_AWSClusterStatus_To_v1beta1_AWSClusterStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha2_AWSCluster_To_v1beta1_AWSCluster is an autogenerated conversion function.
func Convert_v1alpha2_AWSCluster_To_v1beta1_AWSCluster(in *AWSCluster, out *v1beta1.AWSCluster, s conversion.Scope) error {
	return autoConvert_v1alpha2_AWSCluster_To_v1beta1_AWSCluster(in, out, s)
}

func autoConvert_v1beta1_AWSCluster_To_v1alpha2_AWSCluster(in *v1beta1.AWSCluster, out *AWSCluster, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := autoConvert_v1beta1_AWSClusterSpec_To_v1alpha2_AWSClusterSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := autoConvert_v1beta1_AWSClusterStatus_To_v1alpha2_AWSClusterStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_AWSCluster_To_v1alpha2_AWSCluster is an autogenerated conversion function.
func Convert_v1beta1_AWSCluster_To_v1alpha2_AWSCluster(in *v1beta1.AWSCluster, out *AWSCluster, s conversion.Scope) error {
	return autoConvert_v1beta1_AWSCluster_To_v1alpha2_AWSCluster(in, out, s)
}

func autoConvert_v1alpha2_AWSClusterList_To_v1beta1_AWSClusterList(in *AWSClusterList, out *v1beta1.AWSClusterList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1beta1.AWSCluster, len(*in))
		for i := range *in {
			if err := Convert_v1alpha2_AWSCluster_To_v1beta1_AWSCluster(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha2_AWSClusterList_To_v1beta1_AWSClusterList is an autogenerated conversion function.
func Convert_v1alpha2_AWSClusterList_To_v1beta1_AWSClusterList(in *AWSClusterList, out *v1beta1.AWSClusterList, s conversion.Scope) error {
	return autoConvert_v1alpha2_AWSClusterList_To_v1beta1_AWSClusterList(in, out, s)
}

func autoConvert_v1beta1_AWSClusterList_To_v1alpha2_AWSClusterList(in *v1beta1.AWSClusterList, out *AWSClusterList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AWSCluster, len(*in))
		for i := range *in {
			if err := Convert_v1beta1_AWSCluster_To_v1alpha2_AWSCluster(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1beta1_AWSClusterList_To_v1alpha2_AWSClusterList is an autogenerated conversion function.
func Convert_v1beta1_AWSClusterList_To_v1alpha2_AWSClusterList(in *v1beta1.AWSClusterList, out *AWSClusterList, s conversion.Scope) error {
	return autoConvert_v1beta1_AWSClusterList_To_v1alpha2_AWSClusterList(in, out, s)
}

func autoConvert_v1alpha2_AWSClusterSpec_To_v1beta1_AWSClusterSpec(in *AWSClusterSpec, out *v1beta1.AWSClusterSpec, s conversion.Scope) error {
	if err := Convert_v1alpha2_NetworkSpec_To_v1beta1_NetworkSpec(&in.NetworkSpec, &out.NetworkSpec, s); err != nil {
		return err
	}
	out.Region = in.Region
	if err := v1.Convert_string_To_Pointer_string(&in.SSHKeyName, &out.SSHKeyName, s); err != nil {
		return err
	}
	out.AdditionalTags = *(*v1beta1.Tags)(unsafe.Pointer(&in.AdditionalTags))
	if in.ControlPlaneLoadBalancer != nil {
		in, out := &in.ControlPlaneLoadBalancer, &out.ControlPlaneLoadBalancer
		*out = new(v1beta1.AWSLoadBalancerSpec)
		if err := Convert_v1alpha2_AWSLoadBalancerSpec_To_v1beta1_AWSLoadBalancerSpec(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.ControlPlaneLoadBalancer = nil
	}
	out.Bastion = v1beta1.Bastion{
		Enabled: !in.DisableBastionHost,
	}
	return nil
}

func autoConvert_v1beta1_AWSClusterSpec_To_v1alpha2_AWSClusterSpec(in *v1beta1.AWSClusterSpec, out *AWSClusterSpec, s conversion.Scope) error {
	if err := autoConvert_v1beta1_NetworkSpec_To_v1alpha2_NetworkSpec(&in.NetworkSpec, &out.NetworkSpec, s); err != nil {
		return err
	}
	out.Region = in.Region
	if err := v1.Convert_Pointer_string_To_string(&in.SSHKeyName, &out.SSHKeyName, s); err != nil {
		return err
	}
	// WARNING: in.ControlPlaneEndpoint requires manual conversion: does not exist in peer-type
	out.AdditionalTags = *(*Tags)(unsafe.Pointer(&in.AdditionalTags))
	if in.ControlPlaneLoadBalancer != nil {
		in, out := &in.ControlPlaneLoadBalancer, &out.ControlPlaneLoadBalancer
		*out = new(AWSLoadBalancerSpec)
		if err := autoConvert_v1beta1_AWSLoadBalancerSpec_To_v1alpha2_AWSLoadBalancerSpec(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.ControlPlaneLoadBalancer = nil
	}
	// WARNING: in.ImageLookupFormat requires manual conversion: does not exist in peer-type
	// WARNING: in.ImageLookupOrg requires manual conversion: does not exist in peer-type
	// WARNING: in.ImageLookupBaseOS requires manual conversion: does not exist in peer-type
	// WARNING: in.Bastion requires manual conversion: does not exist in peer-type
	// WARNING: in.IdentityRef requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1alpha2_AWSClusterStatus_To_v1beta1_AWSClusterStatus(in *AWSClusterStatus, out *v1beta1.AWSClusterStatus, s conversion.Scope) error {
	autoConvert_v1alpha2_Instance_To_v1beta1_Instance(&in.Bastion, out.Bastion, s)
	out.Ready = in.Ready
	// WARNING: in.APIEndpoints requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1beta1_AWSClusterStatus_To_v1alpha2_AWSClusterStatus(in *v1beta1.AWSClusterStatus, out *AWSClusterStatus, s conversion.Scope) error {
	out.Ready = in.Ready
	// WARNING: in.FailureDomains requires manual conversion: does not exist in peer-type
	// WARNING: in.Bastion requires manual conversion: inconvertible types (*sigs.k8s.io/cluster-api-provider-aws/api/v1beta1.Instance vs ./api/v1alpha2.Instance)
	// WARNING: in.Conditions requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1alpha2_AWSLoadBalancerSpec_To_v1beta1_AWSLoadBalancerSpec(in *AWSLoadBalancerSpec, out *v1beta1.AWSLoadBalancerSpec, s conversion.Scope) error {
	out.Scheme = (*v1beta1.ClassicELBScheme)(unsafe.Pointer(in.Scheme))
	return nil
}

// Convert_v1alpha2_AWSLoadBalancerSpec_To_v1beta1_AWSLoadBalancerSpec is an autogenerated conversion function.
func Convert_v1alpha2_AWSLoadBalancerSpec_To_v1beta1_AWSLoadBalancerSpec(in *AWSLoadBalancerSpec, out *v1beta1.AWSLoadBalancerSpec, s conversion.Scope) error {
	return autoConvert_v1alpha2_AWSLoadBalancerSpec_To_v1beta1_AWSLoadBalancerSpec(in, out, s)
}

func autoConvert_v1beta1_AWSLoadBalancerSpec_To_v1alpha2_AWSLoadBalancerSpec(in *v1beta1.AWSLoadBalancerSpec, out *AWSLoadBalancerSpec, s conversion.Scope) error {
	// WARNING: in.Name requires manual conversion: does not exist in peer-type
	out.Scheme = (*ClassicELBScheme)(unsafe.Pointer(in.Scheme))
	// WARNING: in.CrossZoneLoadBalancing requires manual conversion: does not exist in peer-type
	// WARNING: in.Subnets requires manual conversion: does not exist in peer-type
	// WARNING: in.AdditionalSecurityGroups requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1alpha2_AWSMachine_To_v1beta1_AWSMachine(in *AWSMachine, out *v1beta1.AWSMachine, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := autoConvert_v1alpha2_AWSMachineSpec_To_v1beta1_AWSMachineSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := autoConvert_v1alpha2_AWSMachineStatus_To_v1beta1_AWSMachineStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha2_AWSMachine_To_v1beta1_AWSMachine is an autogenerated conversion function.
func Convert_v1alpha2_AWSMachine_To_v1beta1_AWSMachine(in *AWSMachine, out *v1beta1.AWSMachine, s conversion.Scope) error {
	return autoConvert_v1alpha2_AWSMachine_To_v1beta1_AWSMachine(in, out, s)
}

func autoConvert_v1beta1_AWSMachine_To_v1alpha2_AWSMachine(in *v1beta1.AWSMachine, out *AWSMachine, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := autoConvert_v1beta1_AWSMachineSpec_To_v1alpha2_AWSMachineSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := autoConvert_v1beta1_AWSMachineStatus_To_v1alpha2_AWSMachineStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_AWSMachine_To_v1alpha2_AWSMachine is an autogenerated conversion function.
func Convert_v1beta1_AWSMachine_To_v1alpha2_AWSMachine(in *v1beta1.AWSMachine, out *AWSMachine, s conversion.Scope) error {
	return autoConvert_v1beta1_AWSMachine_To_v1alpha2_AWSMachine(in, out, s)
}

func autoConvert_v1alpha2_AWSMachineList_To_v1beta1_AWSMachineList(in *AWSMachineList, out *v1beta1.AWSMachineList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1beta1.AWSMachine, len(*in))
		for i := range *in {
			if err := Convert_v1alpha2_AWSMachine_To_v1beta1_AWSMachine(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha2_AWSMachineList_To_v1beta1_AWSMachineList is an autogenerated conversion function.
func Convert_v1alpha2_AWSMachineList_To_v1beta1_AWSMachineList(in *AWSMachineList, out *v1beta1.AWSMachineList, s conversion.Scope) error {
	return autoConvert_v1alpha2_AWSMachineList_To_v1beta1_AWSMachineList(in, out, s)
}

func autoConvert_v1beta1_AWSMachineList_To_v1alpha2_AWSMachineList(in *v1beta1.AWSMachineList, out *AWSMachineList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AWSMachine, len(*in))
		for i := range *in {
			if err := Convert_v1beta1_AWSMachine_To_v1alpha2_AWSMachine(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1beta1_AWSMachineList_To_v1alpha2_AWSMachineList is an autogenerated conversion function.
func Convert_v1beta1_AWSMachineList_To_v1alpha2_AWSMachineList(in *v1beta1.AWSMachineList, out *AWSMachineList, s conversion.Scope) error {
	return autoConvert_v1beta1_AWSMachineList_To_v1alpha2_AWSMachineList(in, out, s)
}

func autoConvert_v1alpha2_AWSMachineSpec_To_v1beta1_AWSMachineSpec(in *AWSMachineSpec, out *v1beta1.AWSMachineSpec, s conversion.Scope) error {
	out.ProviderID = (*string)(unsafe.Pointer(in.ProviderID))
	out.AMI = v1beta1.AMIReference{
		ID: in.AMI.ID,
	}
	out.ImageLookupOrg = in.ImageLookupOrg
	out.InstanceType = in.InstanceType
	out.AdditionalTags = *(*v1beta1.Tags)(unsafe.Pointer(&in.AdditionalTags))
	out.IAMInstanceProfile = in.IAMInstanceProfile
	out.PublicIP = (*bool)(unsafe.Pointer(in.PublicIP))
	out.AdditionalSecurityGroups = *(*[]v1beta1.AWSResourceReference)(unsafe.Pointer(&in.AdditionalSecurityGroups))
	out.FailureDomain = in.AvailabilityZone
	out.Subnet = (*v1beta1.AWSResourceReference)(unsafe.Pointer(in.Subnet))
	if err := v1.Convert_string_To_Pointer_string(&in.SSHKeyName, &out.SSHKeyName, s); err != nil {
		return err
	}
	out.RootVolume = &v1beta1.Volume{
		Size: in.RootDeviceSize,
	}
	out.NetworkInterfaces = *(*[]string)(unsafe.Pointer(&in.NetworkInterfaces))
	autoConvert_v1alpha2_CloudInit_To_v1beta1_CloudInit(in.CloudInit, &out.CloudInit, s)
	return nil
}

func autoConvert_v1beta1_AWSMachineSpec_To_v1alpha2_AWSMachineSpec(in *v1beta1.AWSMachineSpec, out *AWSMachineSpec, s conversion.Scope) error {
	out.ProviderID = (*string)(unsafe.Pointer(in.ProviderID))
	// WARNING: in.InstanceID requires manual conversion: does not exist in peer-type
	// WARNING: in.ImageLookupFormat requires manual conversion: does not exist in peer-type
	out.ImageLookupOrg = in.ImageLookupOrg
	// WARNING: in.ImageLookupBaseOS requires manual conversion: does not exist in peer-type
	out.InstanceType = in.InstanceType
	out.AdditionalTags = *(*Tags)(unsafe.Pointer(&in.AdditionalTags))
	out.IAMInstanceProfile = in.IAMInstanceProfile
	out.PublicIP = (*bool)(unsafe.Pointer(in.PublicIP))
	out.AdditionalSecurityGroups = *(*[]AWSResourceReference)(unsafe.Pointer(&in.AdditionalSecurityGroups))
	// WARNING: in.FailureDomain requires manual conversion: does not exist in peer-type
	out.Subnet = (*AWSResourceReference)(unsafe.Pointer(in.Subnet))
	if err := v1.Convert_Pointer_string_To_string(&in.SSHKeyName, &out.SSHKeyName, s); err != nil {
		return err
	}
	// WARNING: in.RootVolume requires manual conversion: does not exist in peer-type
	// WARNING: in.NonRootVolumes requires manual conversion: does not exist in peer-type
	out.NetworkInterfaces = *(*[]string)(unsafe.Pointer(&in.NetworkInterfaces))
	// WARNING: in.UncompressedUserData requires manual conversion: does not exist in peer-type
	// WARNING: in.CloudInit requires manual conversion: inconvertible types (sigs.k8s.io/cluster-api-provider-aws/api/v1beta1.CloudInit vs *./api/v1alpha2.CloudInit)
	// WARNING: in.SpotMarketOptions requires manual conversion: does not exist in peer-type
	// WARNING: in.Tenancy requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1alpha2_AWSMachineStatus_To_v1beta1_AWSMachineStatus(in *AWSMachineStatus, out *v1beta1.AWSMachineStatus, s conversion.Scope) error {
	out.Ready = in.Ready
	out.Addresses = *(*[]apiv1beta1.MachineAddress)(unsafe.Pointer(&in.Addresses))
	out.InstanceState = (*v1beta1.InstanceState)(unsafe.Pointer(in.InstanceState))
	out.FailureMessage = in.ErrorMessage
	out.FailureReason = in.ErrorReason
	return nil
}

func autoConvert_v1beta1_AWSMachineStatus_To_v1alpha2_AWSMachineStatus(in *v1beta1.AWSMachineStatus, out *AWSMachineStatus, s conversion.Scope) error {
	out.Ready = in.Ready
	// WARNING: in.Interruptible requires manual conversion: does not exist in peer-type
	out.Addresses = *(*[]apiv1alpha2.MachineAddress)(unsafe.Pointer(&in.Addresses))
	out.InstanceState = (*InstanceState)(unsafe.Pointer(in.InstanceState))
	// WARNING: in.FailureReason requires manual conversion: does not exist in peer-type
	// WARNING: in.FailureMessage requires manual conversion: does not exist in peer-type
	// WARNING: in.Conditions requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1alpha2_AWSMachineTemplate_To_v1beta1_AWSMachineTemplate(in *AWSMachineTemplate, out *v1beta1.AWSMachineTemplate, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha2_AWSMachineTemplateSpec_To_v1beta1_AWSMachineTemplateSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha2_AWSMachineTemplate_To_v1beta1_AWSMachineTemplate is an autogenerated conversion function.
func Convert_v1alpha2_AWSMachineTemplate_To_v1beta1_AWSMachineTemplate(in *AWSMachineTemplate, out *v1beta1.AWSMachineTemplate, s conversion.Scope) error {
	return autoConvert_v1alpha2_AWSMachineTemplate_To_v1beta1_AWSMachineTemplate(in, out, s)
}

func autoConvert_v1beta1_AWSMachineTemplate_To_v1alpha2_AWSMachineTemplate(in *v1beta1.AWSMachineTemplate, out *AWSMachineTemplate, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_AWSMachineTemplateSpec_To_v1alpha2_AWSMachineTemplateSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_AWSMachineTemplate_To_v1alpha2_AWSMachineTemplate is an autogenerated conversion function.
func Convert_v1beta1_AWSMachineTemplate_To_v1alpha2_AWSMachineTemplate(in *v1beta1.AWSMachineTemplate, out *AWSMachineTemplate, s conversion.Scope) error {
	return autoConvert_v1beta1_AWSMachineTemplate_To_v1alpha2_AWSMachineTemplate(in, out, s)
}

func autoConvert_v1alpha2_AWSMachineTemplateList_To_v1beta1_AWSMachineTemplateList(in *AWSMachineTemplateList, out *v1beta1.AWSMachineTemplateList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1beta1.AWSMachineTemplate, len(*in))
		for i := range *in {
			if err := Convert_v1alpha2_AWSMachineTemplate_To_v1beta1_AWSMachineTemplate(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha2_AWSMachineTemplateList_To_v1beta1_AWSMachineTemplateList is an autogenerated conversion function.
func Convert_v1alpha2_AWSMachineTemplateList_To_v1beta1_AWSMachineTemplateList(in *AWSMachineTemplateList, out *v1beta1.AWSMachineTemplateList, s conversion.Scope) error {
	return autoConvert_v1alpha2_AWSMachineTemplateList_To_v1beta1_AWSMachineTemplateList(in, out, s)
}

func autoConvert_v1beta1_AWSMachineTemplateList_To_v1alpha2_AWSMachineTemplateList(in *v1beta1.AWSMachineTemplateList, out *AWSMachineTemplateList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AWSMachineTemplate, len(*in))
		for i := range *in {
			if err := Convert_v1beta1_AWSMachineTemplate_To_v1alpha2_AWSMachineTemplate(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1beta1_AWSMachineTemplateList_To_v1alpha2_AWSMachineTemplateList is an autogenerated conversion function.
func Convert_v1beta1_AWSMachineTemplateList_To_v1alpha2_AWSMachineTemplateList(in *v1beta1.AWSMachineTemplateList, out *AWSMachineTemplateList, s conversion.Scope) error {
	return autoConvert_v1beta1_AWSMachineTemplateList_To_v1alpha2_AWSMachineTemplateList(in, out, s)
}

func autoConvert_v1alpha2_AWSMachineTemplateResource_To_v1beta1_AWSMachineTemplateResource(in *AWSMachineTemplateResource, out *v1beta1.AWSMachineTemplateResource, s conversion.Scope) error {
	if err := autoConvert_v1alpha2_AWSMachineSpec_To_v1beta1_AWSMachineSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha2_AWSMachineTemplateResource_To_v1beta1_AWSMachineTemplateResource is an autogenerated conversion function.
func Convert_v1alpha2_AWSMachineTemplateResource_To_v1beta1_AWSMachineTemplateResource(in *AWSMachineTemplateResource, out *v1beta1.AWSMachineTemplateResource, s conversion.Scope) error {
	return autoConvert_v1alpha2_AWSMachineTemplateResource_To_v1beta1_AWSMachineTemplateResource(in, out, s)
}

func autoConvert_v1beta1_AWSMachineTemplateResource_To_v1alpha2_AWSMachineTemplateResource(in *v1beta1.AWSMachineTemplateResource, out *AWSMachineTemplateResource, s conversion.Scope) error {
	// WARNING: in.ObjectMeta requires manual conversion: does not exist in peer-type
	if err := autoConvert_v1beta1_AWSMachineSpec_To_v1alpha2_AWSMachineSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha2_AWSMachineTemplateSpec_To_v1beta1_AWSMachineTemplateSpec(in *AWSMachineTemplateSpec, out *v1beta1.AWSMachineTemplateSpec, s conversion.Scope) error {
	if err := Convert_v1alpha2_AWSMachineTemplateResource_To_v1beta1_AWSMachineTemplateResource(&in.Template, &out.Template, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha2_AWSMachineTemplateSpec_To_v1beta1_AWSMachineTemplateSpec is an autogenerated conversion function.
func Convert_v1alpha2_AWSMachineTemplateSpec_To_v1beta1_AWSMachineTemplateSpec(in *AWSMachineTemplateSpec, out *v1beta1.AWSMachineTemplateSpec, s conversion.Scope) error {
	return autoConvert_v1alpha2_AWSMachineTemplateSpec_To_v1beta1_AWSMachineTemplateSpec(in, out, s)
}

func autoConvert_v1beta1_AWSMachineTemplateSpec_To_v1alpha2_AWSMachineTemplateSpec(in *v1beta1.AWSMachineTemplateSpec, out *AWSMachineTemplateSpec, s conversion.Scope) error {
	if err := autoConvert_v1beta1_AWSMachineTemplateResource_To_v1alpha2_AWSMachineTemplateResource(&in.Template, &out.Template, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_AWSMachineTemplateSpec_To_v1alpha2_AWSMachineTemplateSpec is an autogenerated conversion function.
func Convert_v1beta1_AWSMachineTemplateSpec_To_v1alpha2_AWSMachineTemplateSpec(in *v1beta1.AWSMachineTemplateSpec, out *AWSMachineTemplateSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_AWSMachineTemplateSpec_To_v1alpha2_AWSMachineTemplateSpec(in, out, s)
}

func autoConvert_v1alpha2_AWSResourceReference_To_v1beta1_AWSResourceReference(in *AWSResourceReference, out *v1beta1.AWSResourceReference, s conversion.Scope) error {
	out.ID = (*string)(unsafe.Pointer(in.ID))
	out.ARN = (*string)(unsafe.Pointer(in.ARN))
	out.Filters = *(*[]v1beta1.Filter)(unsafe.Pointer(&in.Filters))
	return nil
}

// Convert_v1alpha2_AWSResourceReference_To_v1beta1_AWSResourceReference is an autogenerated conversion function.
func Convert_v1alpha2_AWSResourceReference_To_v1beta1_AWSResourceReference(in *AWSResourceReference, out *v1beta1.AWSResourceReference, s conversion.Scope) error {
	return autoConvert_v1alpha2_AWSResourceReference_To_v1beta1_AWSResourceReference(in, out, s)
}

func autoConvert_v1beta1_AWSResourceReference_To_v1alpha2_AWSResourceReference(in *v1beta1.AWSResourceReference, out *AWSResourceReference, s conversion.Scope) error {
	out.ID = (*string)(unsafe.Pointer(in.ID))
	out.ARN = (*string)(unsafe.Pointer(in.ARN))
	out.Filters = *(*[]Filter)(unsafe.Pointer(&in.Filters))
	return nil
}

// Convert_v1beta1_AWSResourceReference_To_v1alpha2_AWSResourceReference is an autogenerated conversion function.
func Convert_v1beta1_AWSResourceReference_To_v1alpha2_AWSResourceReference(in *v1beta1.AWSResourceReference, out *AWSResourceReference, s conversion.Scope) error {
	return autoConvert_v1beta1_AWSResourceReference_To_v1alpha2_AWSResourceReference(in, out, s)
}

func autoConvert_v1alpha2_BuildParams_To_v1beta1_BuildParams(in *BuildParams, out *v1beta1.BuildParams, s conversion.Scope) error {
	out.Lifecycle = v1beta1.ResourceLifecycle(in.Lifecycle)
	out.ClusterName = in.ClusterName
	out.ResourceID = in.ResourceID
	out.Name = (*string)(unsafe.Pointer(in.Name))
	out.Role = (*string)(unsafe.Pointer(in.Role))
	out.Additional = *(*v1beta1.Tags)(unsafe.Pointer(&in.Additional))
	return nil
}

// Convert_v1alpha2_BuildParams_To_v1beta1_BuildParams is an autogenerated conversion function.
func Convert_v1alpha2_BuildParams_To_v1beta1_BuildParams(in *BuildParams, out *v1beta1.BuildParams, s conversion.Scope) error {
	return autoConvert_v1alpha2_BuildParams_To_v1beta1_BuildParams(in, out, s)
}

func autoConvert_v1beta1_BuildParams_To_v1alpha2_BuildParams(in *v1beta1.BuildParams, out *BuildParams, s conversion.Scope) error {
	out.Lifecycle = ResourceLifecycle(in.Lifecycle)
	out.ClusterName = in.ClusterName
	out.ResourceID = in.ResourceID
	out.Name = (*string)(unsafe.Pointer(in.Name))
	out.Role = (*string)(unsafe.Pointer(in.Role))
	out.Additional = *(*Tags)(unsafe.Pointer(&in.Additional))
	return nil
}

// Convert_v1beta1_BuildParams_To_v1alpha2_BuildParams is an autogenerated conversion function.
func Convert_v1beta1_BuildParams_To_v1alpha2_BuildParams(in *v1beta1.BuildParams, out *BuildParams, s conversion.Scope) error {
	return autoConvert_v1beta1_BuildParams_To_v1alpha2_BuildParams(in, out, s)
}

func autoConvert_v1alpha2_ClassicELB_To_v1beta1_ClassicELB(in *ClassicELB, out *v1beta1.ClassicELB, s conversion.Scope) error {
	out.Name = in.Name
	out.DNSName = in.DNSName
	out.Scheme = v1beta1.ClassicELBScheme(in.Scheme)
	out.SubnetIDs = *(*[]string)(unsafe.Pointer(&in.SubnetIDs))
	out.SecurityGroupIDs = *(*[]string)(unsafe.Pointer(&in.SecurityGroupIDs))
	if in.Listeners != nil {
		in, out := &in.Listeners, &out.Listeners
		*out = make([]v1beta1.ClassicELBListener, len(*in))
		for i, inELB := range *in {
			o := *out
			autoConvert_v1alpha2_ClassicELBListener_To_v1beta1_ClassicELBListener(inELB, &o[i], s)
		}
	} else {
		out.Listeners = nil
	}
	out.HealthCheck = (*v1beta1.ClassicELBHealthCheck)(unsafe.Pointer(in.HealthCheck))
	if err := Convert_v1alpha2_ClassicELBAttributes_To_v1beta1_ClassicELBAttributes(&in.Attributes, &out.Attributes, s); err != nil {
		return err
	}
	out.Tags = *(*map[string]string)(unsafe.Pointer(&in.Tags))
	return nil
}

// Convert_v1alpha2_ClassicELB_To_v1beta1_ClassicELB is an autogenerated conversion function.
func Convert_v1alpha2_ClassicELB_To_v1beta1_ClassicELB(in *ClassicELB, out *v1beta1.ClassicELB, s conversion.Scope) error {
	return autoConvert_v1alpha2_ClassicELB_To_v1beta1_ClassicELB(in, out, s)
}

func autoConvert_v1beta1_ClassicELB_To_v1alpha2_ClassicELB(in *v1beta1.ClassicELB, out *ClassicELB, s conversion.Scope) error {
	out.Name = in.Name
	out.DNSName = in.DNSName
	out.Scheme = ClassicELBScheme(in.Scheme)
	// WARNING: in.AvailabilityZones requires manual conversion: does not exist in peer-type
	out.SubnetIDs = *(*[]string)(unsafe.Pointer(&in.SubnetIDs))
	out.SecurityGroupIDs = *(*[]string)(unsafe.Pointer(&in.SecurityGroupIDs))
	if in.Listeners != nil {
		in, out := &in.Listeners, &out.Listeners
		*out = make([]*ClassicELBListener, len(*in))
		for i, inELB := range *in {
			o := *out
			autoConvert_v1beta1_ClassicELBListener_To_v1alpha2_ClassicELBListener(&inELB, o[i], s)
		}
	} else {
		out.Listeners = nil
	}
	out.HealthCheck = (*ClassicELBHealthCheck)(unsafe.Pointer(in.HealthCheck))
	if err := autoConvert_v1beta1_ClassicELBAttributes_To_v1alpha2_ClassicELBAttributes(&in.Attributes, &out.Attributes, s); err != nil {
		return err
	}
	out.Tags = *(*map[string]string)(unsafe.Pointer(&in.Tags))
	return nil
}

func autoConvert_v1alpha2_ClassicELBAttributes_To_v1beta1_ClassicELBAttributes(in *ClassicELBAttributes, out *v1beta1.ClassicELBAttributes, s conversion.Scope) error {
	out.IdleTimeout = time.Duration(in.IdleTimeout)
	return nil
}

// Convert_v1alpha2_ClassicELBAttributes_To_v1beta1_ClassicELBAttributes is an autogenerated conversion function.
func Convert_v1alpha2_ClassicELBAttributes_To_v1beta1_ClassicELBAttributes(in *ClassicELBAttributes, out *v1beta1.ClassicELBAttributes, s conversion.Scope) error {
	return autoConvert_v1alpha2_ClassicELBAttributes_To_v1beta1_ClassicELBAttributes(in, out, s)
}

func autoConvert_v1beta1_ClassicELBAttributes_To_v1alpha2_ClassicELBAttributes(in *v1beta1.ClassicELBAttributes, out *ClassicELBAttributes, s conversion.Scope) error {
	out.IdleTimeout = time.Duration(in.IdleTimeout)
	// WARNING: in.CrossZoneLoadBalancing requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1alpha2_ClassicELBHealthCheck_To_v1beta1_ClassicELBHealthCheck(in *ClassicELBHealthCheck, out *v1beta1.ClassicELBHealthCheck, s conversion.Scope) error {
	out.Target = in.Target
	out.Interval = time.Duration(in.Interval)
	out.Timeout = time.Duration(in.Timeout)
	out.HealthyThreshold = in.HealthyThreshold
	out.UnhealthyThreshold = in.UnhealthyThreshold
	return nil
}

// Convert_v1alpha2_ClassicELBHealthCheck_To_v1beta1_ClassicELBHealthCheck is an autogenerated conversion function.
func Convert_v1alpha2_ClassicELBHealthCheck_To_v1beta1_ClassicELBHealthCheck(in *ClassicELBHealthCheck, out *v1beta1.ClassicELBHealthCheck, s conversion.Scope) error {
	return autoConvert_v1alpha2_ClassicELBHealthCheck_To_v1beta1_ClassicELBHealthCheck(in, out, s)
}

func autoConvert_v1beta1_ClassicELBHealthCheck_To_v1alpha2_ClassicELBHealthCheck(in *v1beta1.ClassicELBHealthCheck, out *ClassicELBHealthCheck, s conversion.Scope) error {
	out.Target = in.Target
	out.Interval = time.Duration(in.Interval)
	out.Timeout = time.Duration(in.Timeout)
	out.HealthyThreshold = in.HealthyThreshold
	out.UnhealthyThreshold = in.UnhealthyThreshold
	return nil
}

// Convert_v1beta1_ClassicELBHealthCheck_To_v1alpha2_ClassicELBHealthCheck is an autogenerated conversion function.
func Convert_v1beta1_ClassicELBHealthCheck_To_v1alpha2_ClassicELBHealthCheck(in *v1beta1.ClassicELBHealthCheck, out *ClassicELBHealthCheck, s conversion.Scope) error {
	return autoConvert_v1beta1_ClassicELBHealthCheck_To_v1alpha2_ClassicELBHealthCheck(in, out, s)
}

func autoConvert_v1alpha2_ClassicELBListener_To_v1beta1_ClassicELBListener(in *ClassicELBListener, out *v1beta1.ClassicELBListener, s conversion.Scope) error {
	out.Protocol = v1beta1.ClassicELBProtocol(in.Protocol)
	out.Port = in.Port
	out.InstanceProtocol = v1beta1.ClassicELBProtocol(in.InstanceProtocol)
	out.InstancePort = in.InstancePort
	return nil
}

// Convert_v1alpha2_ClassicELBListener_To_v1beta1_ClassicELBListener is an autogenerated conversion function.
func Convert_v1alpha2_ClassicELBListener_To_v1beta1_ClassicELBListener(in *ClassicELBListener, out *v1beta1.ClassicELBListener, s conversion.Scope) error {
	return autoConvert_v1alpha2_ClassicELBListener_To_v1beta1_ClassicELBListener(in, out, s)
}

func autoConvert_v1beta1_ClassicELBListener_To_v1alpha2_ClassicELBListener(in *v1beta1.ClassicELBListener, out *ClassicELBListener, s conversion.Scope) error {
	out.Protocol = ClassicELBProtocol(in.Protocol)
	out.Port = in.Port
	out.InstanceProtocol = ClassicELBProtocol(in.InstanceProtocol)
	out.InstancePort = in.InstancePort
	return nil
}

// Convert_v1beta1_ClassicELBListener_To_v1alpha2_ClassicELBListener is an autogenerated conversion function.
func Convert_v1beta1_ClassicELBListener_To_v1alpha2_ClassicELBListener(in *v1beta1.ClassicELBListener, out *ClassicELBListener, s conversion.Scope) error {
	return autoConvert_v1beta1_ClassicELBListener_To_v1alpha2_ClassicELBListener(in, out, s)
}

func autoConvert_v1alpha2_CloudInit_To_v1beta1_CloudInit(in *CloudInit, out *v1beta1.CloudInit, s conversion.Scope) error {
	out.SecretCount = in.SecretCount
	out.SecretPrefix = in.SecretPrefix
	out.InsecureSkipSecretsManager = !in.EnableSecureSecretsManager
	if in.EnableSecureSecretsManager {
		out.SecureSecretsBackend = v1beta1.SecretBackendSecretsManager
	}
	return nil
}

func autoConvert_v1beta1_CloudInit_To_v1alpha2_CloudInit(in *v1beta1.CloudInit, out *CloudInit, s conversion.Scope) error {
	// WARNING: in.InsecureSkipSecretsManager requires manual conversion: does not exist in peer-type
	out.SecretCount = in.SecretCount
	out.SecretPrefix = in.SecretPrefix
	// WARNING: in.SecureSecretsBackend requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1alpha2_Filter_To_v1beta1_Filter(in *Filter, out *v1beta1.Filter, s conversion.Scope) error {
	out.Name = in.Name
	out.Values = *(*[]string)(unsafe.Pointer(&in.Values))
	return nil
}

// Convert_v1alpha2_Filter_To_v1beta1_Filter is an autogenerated conversion function.
func Convert_v1alpha2_Filter_To_v1beta1_Filter(in *Filter, out *v1beta1.Filter, s conversion.Scope) error {
	return autoConvert_v1alpha2_Filter_To_v1beta1_Filter(in, out, s)
}

func autoConvert_v1beta1_Filter_To_v1alpha2_Filter(in *v1beta1.Filter, out *Filter, s conversion.Scope) error {
	out.Name = in.Name
	out.Values = *(*[]string)(unsafe.Pointer(&in.Values))
	return nil
}

// Convert_v1beta1_Filter_To_v1alpha2_Filter is an autogenerated conversion function.
func Convert_v1beta1_Filter_To_v1alpha2_Filter(in *v1beta1.Filter, out *Filter, s conversion.Scope) error {
	return autoConvert_v1beta1_Filter_To_v1alpha2_Filter(in, out, s)
}

func autoConvert_v1alpha2_IngressRule_To_v1beta1_IngressRule(in *IngressRule, out *v1beta1.IngressRule, s conversion.Scope) error {
	out.Description = in.Description
	out.Protocol = v1beta1.SecurityGroupProtocol(in.Protocol)
	out.FromPort = in.FromPort
	out.ToPort = in.ToPort
	out.CidrBlocks = *(*[]string)(unsafe.Pointer(&in.CidrBlocks))
	out.SourceSecurityGroupIDs = *(*[]string)(unsafe.Pointer(&in.SourceSecurityGroupIDs))
	return nil
}

// Convert_v1alpha2_IngressRule_To_v1beta1_IngressRule is an autogenerated conversion function.
func Convert_v1alpha2_IngressRule_To_v1beta1_IngressRule(in *IngressRule, out *v1beta1.IngressRule, s conversion.Scope) error {
	return autoConvert_v1alpha2_IngressRule_To_v1beta1_IngressRule(in, out, s)
}

func autoConvert_v1beta1_IngressRule_To_v1alpha2_IngressRule(in *v1beta1.IngressRule, out *IngressRule, s conversion.Scope) error {
	out.Description = in.Description
	out.Protocol = SecurityGroupProtocol(in.Protocol)
	out.FromPort = in.FromPort
	out.ToPort = in.ToPort
	out.CidrBlocks = *(*[]string)(unsafe.Pointer(&in.CidrBlocks))
	out.SourceSecurityGroupIDs = *(*[]string)(unsafe.Pointer(&in.SourceSecurityGroupIDs))
	return nil
}

// Convert_v1beta1_IngressRule_To_v1alpha2_IngressRule is an autogenerated conversion function.
func Convert_v1beta1_IngressRule_To_v1alpha2_IngressRule(in *v1beta1.IngressRule, out *IngressRule, s conversion.Scope) error {
	return autoConvert_v1beta1_IngressRule_To_v1alpha2_IngressRule(in, out, s)
}

func autoConvert_v1alpha2_Instance_To_v1beta1_Instance(in *Instance, out *v1beta1.Instance, s conversion.Scope) error {
	out.ID = in.ID
	out.State = v1beta1.InstanceState(in.State)
	out.Type = in.Type
	out.SubnetID = in.SubnetID
	out.ImageID = in.ImageID
	out.SSHKeyName = (*string)(unsafe.Pointer(in.SSHKeyName))
	out.SecurityGroupIDs = *(*[]string)(unsafe.Pointer(&in.SecurityGroupIDs))
	out.UserData = (*string)(unsafe.Pointer(in.UserData))
	out.IAMProfile = in.IAMProfile
	out.Addresses = *(*[]apiv1beta1.MachineAddress)(unsafe.Pointer(&in.Addresses))
	out.PrivateIP = (*string)(unsafe.Pointer(in.PrivateIP))
	out.PublicIP = (*string)(unsafe.Pointer(in.PublicIP))
	out.ENASupport = (*bool)(unsafe.Pointer(in.ENASupport))
	out.EBSOptimized = (*bool)(unsafe.Pointer(in.EBSOptimized))
	out.RootVolume = &v1beta1.Volume{
		Size: in.RootDeviceSize,
	}
	out.NetworkInterfaces = *(*[]string)(unsafe.Pointer(&in.NetworkInterfaces))
	out.Tags = *(*map[string]string)(unsafe.Pointer(&in.Tags))
	return nil
}

func autoConvert_v1beta1_Instance_To_v1alpha2_Instance(in *v1beta1.Instance, out *Instance, s conversion.Scope) error {
	out.ID = in.ID
	out.State = InstanceState(in.State)
	out.Type = in.Type
	out.SubnetID = in.SubnetID
	out.ImageID = in.ImageID
	out.SSHKeyName = (*string)(unsafe.Pointer(in.SSHKeyName))
	out.SecurityGroupIDs = *(*[]string)(unsafe.Pointer(&in.SecurityGroupIDs))
	out.UserData = (*string)(unsafe.Pointer(in.UserData))
	out.IAMProfile = in.IAMProfile
	out.Addresses = *(*[]apiv1alpha2.MachineAddress)(unsafe.Pointer(&in.Addresses))
	out.PrivateIP = (*string)(unsafe.Pointer(in.PrivateIP))
	out.PublicIP = (*string)(unsafe.Pointer(in.PublicIP))
	out.ENASupport = (*bool)(unsafe.Pointer(in.ENASupport))
	out.EBSOptimized = (*bool)(unsafe.Pointer(in.EBSOptimized))
	// WARNING: in.RootVolume requires manual conversion: does not exist in peer-type
	// WARNING: in.NonRootVolumes requires manual conversion: does not exist in peer-type
	out.NetworkInterfaces = *(*[]string)(unsafe.Pointer(&in.NetworkInterfaces))
	out.Tags = *(*map[string]string)(unsafe.Pointer(&in.Tags))
	// WARNING: in.AvailabilityZone requires manual conversion: does not exist in peer-type
	// WARNING: in.SpotMarketOptions requires manual conversion: does not exist in peer-type
	// WARNING: in.Tenancy requires manual conversion: does not exist in peer-type
	// WARNING: in.VolumeIDs requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1alpha2_NetworkSpec_To_v1beta1_NetworkSpec(in *NetworkSpec, out *v1beta1.NetworkSpec, s conversion.Scope) error {
	if err := Convert_v1alpha2_VPCSpec_To_v1beta1_VPCSpec(&in.VPC, &out.VPC, s); err != nil {
		return err
	}
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make(v1beta1.Subnets, len(*in))
		for i, inSub := range *in {
			o := *out
			autoConvert_v1alpha2_SubnetSpec_To_v1beta1_SubnetSpec(inSub, &o[i], s)
		}
	} else {
		out.Subnets = nil
	}
	return nil
}

// Convert_v1alpha2_NetworkSpec_To_v1beta1_NetworkSpec is an autogenerated conversion function.
func Convert_v1alpha2_NetworkSpec_To_v1beta1_NetworkSpec(in *NetworkSpec, out *v1beta1.NetworkSpec, s conversion.Scope) error {
	return autoConvert_v1alpha2_NetworkSpec_To_v1beta1_NetworkSpec(in, out, s)
}

func autoConvert_v1beta1_NetworkSpec_To_v1alpha2_NetworkSpec(in *v1beta1.NetworkSpec, out *NetworkSpec, s conversion.Scope) error {
	if err := Convert_v1beta1_VPCSpec_To_v1alpha2_VPCSpec(&in.VPC, &out.VPC, s); err != nil {
		return err
	}
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make(Subnets, len(*in))
		for i, inSub := range *in {
			o := *out
			autoConvert_v1beta1_SubnetSpec_To_v1alpha2_SubnetSpec(&inSub, o[i], s)
		}
	} else {
		out.Subnets = nil
	}
	// WARNING: in.CNI requires manual conversion: does not exist in peer-type
	// WARNING: in.SecurityGroupOverrides requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1alpha2_RouteTable_To_v1beta1_RouteTable(in *RouteTable, out *v1beta1.RouteTable, s conversion.Scope) error {
	out.ID = in.ID
	return nil
}

// Convert_v1alpha2_RouteTable_To_v1beta1_RouteTable is an autogenerated conversion function.
func Convert_v1alpha2_RouteTable_To_v1beta1_RouteTable(in *RouteTable, out *v1beta1.RouteTable, s conversion.Scope) error {
	return autoConvert_v1alpha2_RouteTable_To_v1beta1_RouteTable(in, out, s)
}

func autoConvert_v1beta1_RouteTable_To_v1alpha2_RouteTable(in *v1beta1.RouteTable, out *RouteTable, s conversion.Scope) error {
	out.ID = in.ID
	return nil
}

// Convert_v1beta1_RouteTable_To_v1alpha2_RouteTable is an autogenerated conversion function.
func Convert_v1beta1_RouteTable_To_v1alpha2_RouteTable(in *v1beta1.RouteTable, out *RouteTable, s conversion.Scope) error {
	return autoConvert_v1beta1_RouteTable_To_v1alpha2_RouteTable(in, out, s)
}

func autoConvert_v1alpha2_SecurityGroup_To_v1beta1_SecurityGroup(in *SecurityGroup, out *v1beta1.SecurityGroup, s conversion.Scope) error {
	out.ID = in.ID
	out.Name = in.Name
	if in.IngressRules != nil {
		in, out := &in.IngressRules, &out.IngressRules
		*out = make(v1beta1.IngressRules, len(*in))
		for i, inIngress := range *in {
			o := *out
			autoConvert_v1alpha2_IngressRule_To_v1beta1_IngressRule(inIngress, &o[i], s)
		}
	} else {
		out.IngressRules = nil
	}
	out.Tags = *(*v1beta1.Tags)(unsafe.Pointer(&in.Tags))
	return nil
}

// Convert_v1alpha2_SecurityGroup_To_v1beta1_SecurityGroup is an autogenerated conversion function.
func Convert_v1alpha2_SecurityGroup_To_v1beta1_SecurityGroup(in *SecurityGroup, out *v1beta1.SecurityGroup, s conversion.Scope) error {
	return autoConvert_v1alpha2_SecurityGroup_To_v1beta1_SecurityGroup(in, out, s)
}

func autoConvert_v1beta1_SecurityGroup_To_v1alpha2_SecurityGroup(in *v1beta1.SecurityGroup, out *SecurityGroup, s conversion.Scope) error {
	out.ID = in.ID
	out.Name = in.Name
	if in.IngressRules != nil {
		in, out := &in.IngressRules, &out.IngressRules
		*out = make(IngressRules, len(*in))
		for i, inIngress := range *in {
			o := *out
			autoConvert_v1beta1_IngressRule_To_v1alpha2_IngressRule(&inIngress, o[i], s)
		}
	} else {
		out.IngressRules = nil
	}
	out.Tags = *(*Tags)(unsafe.Pointer(&in.Tags))
	return nil
}

// Convert_v1beta1_SecurityGroup_To_v1alpha2_SecurityGroup is an autogenerated conversion function.
func Convert_v1beta1_SecurityGroup_To_v1alpha2_SecurityGroup(in *v1beta1.SecurityGroup, out *SecurityGroup, s conversion.Scope) error {
	return autoConvert_v1beta1_SecurityGroup_To_v1alpha2_SecurityGroup(in, out, s)
}

func autoConvert_v1alpha2_SubnetSpec_To_v1beta1_SubnetSpec(in *SubnetSpec, out *v1beta1.SubnetSpec, s conversion.Scope) error {
	out.ID = in.ID
	out.CidrBlock = in.CidrBlock
	out.AvailabilityZone = in.AvailabilityZone
	out.IsPublic = in.IsPublic
	out.RouteTableID = (*string)(unsafe.Pointer(in.RouteTableID))
	out.NatGatewayID = (*string)(unsafe.Pointer(in.NatGatewayID))
	out.Tags = *(*v1beta1.Tags)(unsafe.Pointer(&in.Tags))
	return nil
}

// Convert_v1alpha2_SubnetSpec_To_v1beta1_SubnetSpec is an autogenerated conversion function.
func Convert_v1alpha2_SubnetSpec_To_v1beta1_SubnetSpec(in *SubnetSpec, out *v1beta1.SubnetSpec, s conversion.Scope) error {
	return autoConvert_v1alpha2_SubnetSpec_To_v1beta1_SubnetSpec(in, out, s)
}

func autoConvert_v1beta1_SubnetSpec_To_v1alpha2_SubnetSpec(in *v1beta1.SubnetSpec, out *SubnetSpec, s conversion.Scope) error {
	out.ID = in.ID
	out.CidrBlock = in.CidrBlock
	out.AvailabilityZone = in.AvailabilityZone
	out.IsPublic = in.IsPublic
	out.RouteTableID = (*string)(unsafe.Pointer(in.RouteTableID))
	out.NatGatewayID = (*string)(unsafe.Pointer(in.NatGatewayID))
	out.Tags = *(*Tags)(unsafe.Pointer(&in.Tags))
	return nil
}

// Convert_v1beta1_SubnetSpec_To_v1alpha2_SubnetSpec is an autogenerated conversion function.
func Convert_v1beta1_SubnetSpec_To_v1alpha2_SubnetSpec(in *v1beta1.SubnetSpec, out *SubnetSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_SubnetSpec_To_v1alpha2_SubnetSpec(in, out, s)
}

func autoConvert_v1alpha2_VPCSpec_To_v1beta1_VPCSpec(in *VPCSpec, out *v1beta1.VPCSpec, s conversion.Scope) error {
	out.ID = in.ID
	out.CidrBlock = in.CidrBlock
	out.InternetGatewayID = (*string)(unsafe.Pointer(in.InternetGatewayID))
	out.Tags = *(*v1beta1.Tags)(unsafe.Pointer(&in.Tags))
	return nil
}

// Convert_v1alpha2_VPCSpec_To_v1beta1_VPCSpec is an autogenerated conversion function.
func Convert_v1alpha2_VPCSpec_To_v1beta1_VPCSpec(in *VPCSpec, out *v1beta1.VPCSpec, s conversion.Scope) error {
	return autoConvert_v1alpha2_VPCSpec_To_v1beta1_VPCSpec(in, out, s)
}

func autoConvert_v1beta1_VPCSpec_To_v1alpha2_VPCSpec(in *v1beta1.VPCSpec, out *VPCSpec, s conversion.Scope) error {
	out.ID = in.ID
	out.CidrBlock = in.CidrBlock
	out.InternetGatewayID = (*string)(unsafe.Pointer(in.InternetGatewayID))
	out.Tags = *(*Tags)(unsafe.Pointer(&in.Tags))
	// WARNING: in.AvailabilityZoneUsageLimit requires manual conversion: does not exist in peer-type
	// WARNING: in.AvailabilityZoneSelection requires manual conversion: does not exist in peer-type
	return nil
}
