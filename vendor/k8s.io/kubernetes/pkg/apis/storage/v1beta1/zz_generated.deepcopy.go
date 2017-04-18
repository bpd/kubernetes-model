// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1beta1

import (
	v1 "k8s.io/kubernetes/pkg/api/v1"
	conversion "k8s.io/kubernetes/pkg/conversion"
	runtime "k8s.io/kubernetes/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_StorageClass, InType: reflect.TypeOf(&StorageClass{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_StorageClassList, InType: reflect.TypeOf(&StorageClassList{})},
	)
}

func DeepCopy_v1beta1_StorageClass(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*StorageClass)
		out := out.(*StorageClass)
		out.TypeMeta = in.TypeMeta
		if err := v1.DeepCopy_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		out.Provisioner = in.Provisioner
		if in.Parameters != nil {
			in, out := &in.Parameters, &out.Parameters
			*out = make(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		} else {
			out.Parameters = nil
		}
		return nil
	}
}

func DeepCopy_v1beta1_StorageClassList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*StorageClassList)
		out := out.(*StorageClassList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]StorageClass, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_StorageClass(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}