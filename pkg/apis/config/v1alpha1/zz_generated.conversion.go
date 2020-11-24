// +build !ignore_autogenerated

/*
Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

package v1alpha1

import (
	unsafe "unsafe"

	healthcheckconfig "github.com/gardener/gardener/extensions/pkg/controller/healthcheck/config"
	configv1alpha1 "github.com/gardener/gardener/extensions/pkg/controller/healthcheck/config/v1alpha1"
	config "github.com/metal-stack/gardener-extension-provider-metal/pkg/apis/config"
	resource "k8s.io/apimachinery/pkg/api/resource"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*AccountingExporterClientConfiguration)(nil), (*config.AccountingExporterClientConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_AccountingExporterClientConfiguration_To_config_AccountingExporterClientConfiguration(a.(*AccountingExporterClientConfiguration), b.(*config.AccountingExporterClientConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.AccountingExporterClientConfiguration)(nil), (*AccountingExporterClientConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_AccountingExporterClientConfiguration_To_v1alpha1_AccountingExporterClientConfiguration(a.(*config.AccountingExporterClientConfiguration), b.(*AccountingExporterClientConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AccountingExporterConfiguration)(nil), (*config.AccountingExporterConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_AccountingExporterConfiguration_To_config_AccountingExporterConfiguration(a.(*AccountingExporterConfiguration), b.(*config.AccountingExporterConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.AccountingExporterConfiguration)(nil), (*AccountingExporterConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_AccountingExporterConfiguration_To_v1alpha1_AccountingExporterConfiguration(a.(*config.AccountingExporterConfiguration), b.(*AccountingExporterConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AccountingExporterNetworkTrafficConfiguration)(nil), (*config.AccountingExporterNetworkTrafficConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_AccountingExporterNetworkTrafficConfiguration_To_config_AccountingExporterNetworkTrafficConfiguration(a.(*AccountingExporterNetworkTrafficConfiguration), b.(*config.AccountingExporterNetworkTrafficConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.AccountingExporterNetworkTrafficConfiguration)(nil), (*AccountingExporterNetworkTrafficConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_AccountingExporterNetworkTrafficConfiguration_To_v1alpha1_AccountingExporterNetworkTrafficConfiguration(a.(*config.AccountingExporterNetworkTrafficConfiguration), b.(*AccountingExporterNetworkTrafficConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Auth)(nil), (*config.Auth)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Auth_To_config_Auth(a.(*Auth), b.(*config.Auth), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.Auth)(nil), (*Auth)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_Auth_To_v1alpha1_Auth(a.(*config.Auth), b.(*Auth), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ControllerConfiguration)(nil), (*config.ControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ControllerConfiguration_To_config_ControllerConfiguration(a.(*ControllerConfiguration), b.(*config.ControllerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ControllerConfiguration)(nil), (*ControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ControllerConfiguration_To_v1alpha1_ControllerConfiguration(a.(*config.ControllerConfiguration), b.(*ControllerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ETCD)(nil), (*config.ETCD)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ETCD_To_config_ETCD(a.(*ETCD), b.(*config.ETCD), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ETCD)(nil), (*ETCD)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ETCD_To_v1alpha1_ETCD(a.(*config.ETCD), b.(*ETCD), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ETCDBackup)(nil), (*config.ETCDBackup)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ETCDBackup_To_config_ETCDBackup(a.(*ETCDBackup), b.(*config.ETCDBackup), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ETCDBackup)(nil), (*ETCDBackup)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ETCDBackup_To_v1alpha1_ETCDBackup(a.(*config.ETCDBackup), b.(*ETCDBackup), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ETCDStorage)(nil), (*config.ETCDStorage)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ETCDStorage_To_config_ETCDStorage(a.(*ETCDStorage), b.(*config.ETCDStorage), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ETCDStorage)(nil), (*ETCDStorage)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ETCDStorage_To_v1alpha1_ETCDStorage(a.(*config.ETCDStorage), b.(*ETCDStorage), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MachineImage)(nil), (*config.MachineImage)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_MachineImage_To_config_MachineImage(a.(*MachineImage), b.(*config.MachineImage), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.MachineImage)(nil), (*MachineImage)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_MachineImage_To_v1alpha1_MachineImage(a.(*config.MachineImage), b.(*MachineImage), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*SplunkAudit)(nil), (*config.SplunkAudit)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_SplunkAudit_To_config_SplunkAudit(a.(*SplunkAudit), b.(*config.SplunkAudit), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.SplunkAudit)(nil), (*SplunkAudit)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_SplunkAudit_To_v1alpha1_SplunkAudit(a.(*config.SplunkAudit), b.(*SplunkAudit), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_AccountingExporterClientConfiguration_To_config_AccountingExporterClientConfiguration(in *AccountingExporterClientConfiguration, out *config.AccountingExporterClientConfiguration, s conversion.Scope) error {
	out.Hostname = in.Hostname
	out.Port = in.Port
	out.CA = in.CA
	out.Cert = in.Cert
	out.CertKey = in.CertKey
	return nil
}

// Convert_v1alpha1_AccountingExporterClientConfiguration_To_config_AccountingExporterClientConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_AccountingExporterClientConfiguration_To_config_AccountingExporterClientConfiguration(in *AccountingExporterClientConfiguration, out *config.AccountingExporterClientConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_AccountingExporterClientConfiguration_To_config_AccountingExporterClientConfiguration(in, out, s)
}

func autoConvert_config_AccountingExporterClientConfiguration_To_v1alpha1_AccountingExporterClientConfiguration(in *config.AccountingExporterClientConfiguration, out *AccountingExporterClientConfiguration, s conversion.Scope) error {
	out.Hostname = in.Hostname
	out.Port = in.Port
	out.CA = in.CA
	out.Cert = in.Cert
	out.CertKey = in.CertKey
	return nil
}

// Convert_config_AccountingExporterClientConfiguration_To_v1alpha1_AccountingExporterClientConfiguration is an autogenerated conversion function.
func Convert_config_AccountingExporterClientConfiguration_To_v1alpha1_AccountingExporterClientConfiguration(in *config.AccountingExporterClientConfiguration, out *AccountingExporterClientConfiguration, s conversion.Scope) error {
	return autoConvert_config_AccountingExporterClientConfiguration_To_v1alpha1_AccountingExporterClientConfiguration(in, out, s)
}

func autoConvert_v1alpha1_AccountingExporterConfiguration_To_config_AccountingExporterConfiguration(in *AccountingExporterConfiguration, out *config.AccountingExporterConfiguration, s conversion.Scope) error {
	out.Enabled = in.Enabled
	if err := Convert_v1alpha1_AccountingExporterNetworkTrafficConfiguration_To_config_AccountingExporterNetworkTrafficConfiguration(&in.NetworkTraffic, &out.NetworkTraffic, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_AccountingExporterClientConfiguration_To_config_AccountingExporterClientConfiguration(&in.Client, &out.Client, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_AccountingExporterConfiguration_To_config_AccountingExporterConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_AccountingExporterConfiguration_To_config_AccountingExporterConfiguration(in *AccountingExporterConfiguration, out *config.AccountingExporterConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_AccountingExporterConfiguration_To_config_AccountingExporterConfiguration(in, out, s)
}

func autoConvert_config_AccountingExporterConfiguration_To_v1alpha1_AccountingExporterConfiguration(in *config.AccountingExporterConfiguration, out *AccountingExporterConfiguration, s conversion.Scope) error {
	out.Enabled = in.Enabled
	if err := Convert_config_AccountingExporterNetworkTrafficConfiguration_To_v1alpha1_AccountingExporterNetworkTrafficConfiguration(&in.NetworkTraffic, &out.NetworkTraffic, s); err != nil {
		return err
	}
	if err := Convert_config_AccountingExporterClientConfiguration_To_v1alpha1_AccountingExporterClientConfiguration(&in.Client, &out.Client, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_AccountingExporterConfiguration_To_v1alpha1_AccountingExporterConfiguration is an autogenerated conversion function.
func Convert_config_AccountingExporterConfiguration_To_v1alpha1_AccountingExporterConfiguration(in *config.AccountingExporterConfiguration, out *AccountingExporterConfiguration, s conversion.Scope) error {
	return autoConvert_config_AccountingExporterConfiguration_To_v1alpha1_AccountingExporterConfiguration(in, out, s)
}

func autoConvert_v1alpha1_AccountingExporterNetworkTrafficConfiguration_To_config_AccountingExporterNetworkTrafficConfiguration(in *AccountingExporterNetworkTrafficConfiguration, out *config.AccountingExporterNetworkTrafficConfiguration, s conversion.Scope) error {
	out.Enabled = in.Enabled
	out.InternalNetworks = *(*[]string)(unsafe.Pointer(&in.InternalNetworks))
	return nil
}

// Convert_v1alpha1_AccountingExporterNetworkTrafficConfiguration_To_config_AccountingExporterNetworkTrafficConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_AccountingExporterNetworkTrafficConfiguration_To_config_AccountingExporterNetworkTrafficConfiguration(in *AccountingExporterNetworkTrafficConfiguration, out *config.AccountingExporterNetworkTrafficConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_AccountingExporterNetworkTrafficConfiguration_To_config_AccountingExporterNetworkTrafficConfiguration(in, out, s)
}

func autoConvert_config_AccountingExporterNetworkTrafficConfiguration_To_v1alpha1_AccountingExporterNetworkTrafficConfiguration(in *config.AccountingExporterNetworkTrafficConfiguration, out *AccountingExporterNetworkTrafficConfiguration, s conversion.Scope) error {
	out.Enabled = in.Enabled
	out.InternalNetworks = *(*[]string)(unsafe.Pointer(&in.InternalNetworks))
	return nil
}

// Convert_config_AccountingExporterNetworkTrafficConfiguration_To_v1alpha1_AccountingExporterNetworkTrafficConfiguration is an autogenerated conversion function.
func Convert_config_AccountingExporterNetworkTrafficConfiguration_To_v1alpha1_AccountingExporterNetworkTrafficConfiguration(in *config.AccountingExporterNetworkTrafficConfiguration, out *AccountingExporterNetworkTrafficConfiguration, s conversion.Scope) error {
	return autoConvert_config_AccountingExporterNetworkTrafficConfiguration_To_v1alpha1_AccountingExporterNetworkTrafficConfiguration(in, out, s)
}

func autoConvert_v1alpha1_Auth_To_config_Auth(in *Auth, out *config.Auth, s conversion.Scope) error {
	out.Enabled = in.Enabled
	out.ProviderTenant = in.ProviderTenant
	return nil
}

// Convert_v1alpha1_Auth_To_config_Auth is an autogenerated conversion function.
func Convert_v1alpha1_Auth_To_config_Auth(in *Auth, out *config.Auth, s conversion.Scope) error {
	return autoConvert_v1alpha1_Auth_To_config_Auth(in, out, s)
}

func autoConvert_config_Auth_To_v1alpha1_Auth(in *config.Auth, out *Auth, s conversion.Scope) error {
	out.Enabled = in.Enabled
	out.ProviderTenant = in.ProviderTenant
	return nil
}

// Convert_config_Auth_To_v1alpha1_Auth is an autogenerated conversion function.
func Convert_config_Auth_To_v1alpha1_Auth(in *config.Auth, out *Auth, s conversion.Scope) error {
	return autoConvert_config_Auth_To_v1alpha1_Auth(in, out, s)
}

func autoConvert_v1alpha1_ControllerConfiguration_To_config_ControllerConfiguration(in *ControllerConfiguration, out *config.ControllerConfiguration, s conversion.Scope) error {
	out.MachineImages = *(*[]config.MachineImage)(unsafe.Pointer(&in.MachineImages))
	if err := Convert_v1alpha1_ETCD_To_config_ETCD(&in.ETCD, &out.ETCD, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_Auth_To_config_Auth(&in.Auth, &out.Auth, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_SplunkAudit_To_config_SplunkAudit(&in.SplunkAudit, &out.SplunkAudit, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_AccountingExporterConfiguration_To_config_AccountingExporterConfiguration(&in.AccountingExporter, &out.AccountingExporter, s); err != nil {
		return err
	}
	out.HealthCheckConfig = (*healthcheckconfig.HealthCheckConfig)(unsafe.Pointer(in.HealthCheckConfig))
	return nil
}

// Convert_v1alpha1_ControllerConfiguration_To_config_ControllerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_ControllerConfiguration_To_config_ControllerConfiguration(in *ControllerConfiguration, out *config.ControllerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_ControllerConfiguration_To_config_ControllerConfiguration(in, out, s)
}

func autoConvert_config_ControllerConfiguration_To_v1alpha1_ControllerConfiguration(in *config.ControllerConfiguration, out *ControllerConfiguration, s conversion.Scope) error {
	out.MachineImages = *(*[]MachineImage)(unsafe.Pointer(&in.MachineImages))
	if err := Convert_config_ETCD_To_v1alpha1_ETCD(&in.ETCD, &out.ETCD, s); err != nil {
		return err
	}
	if err := Convert_config_Auth_To_v1alpha1_Auth(&in.Auth, &out.Auth, s); err != nil {
		return err
	}
	if err := Convert_config_SplunkAudit_To_v1alpha1_SplunkAudit(&in.SplunkAudit, &out.SplunkAudit, s); err != nil {
		return err
	}
	if err := Convert_config_AccountingExporterConfiguration_To_v1alpha1_AccountingExporterConfiguration(&in.AccountingExporter, &out.AccountingExporter, s); err != nil {
		return err
	}
	out.HealthCheckConfig = (*configv1alpha1.HealthCheckConfig)(unsafe.Pointer(in.HealthCheckConfig))
	return nil
}

// Convert_config_ControllerConfiguration_To_v1alpha1_ControllerConfiguration is an autogenerated conversion function.
func Convert_config_ControllerConfiguration_To_v1alpha1_ControllerConfiguration(in *config.ControllerConfiguration, out *ControllerConfiguration, s conversion.Scope) error {
	return autoConvert_config_ControllerConfiguration_To_v1alpha1_ControllerConfiguration(in, out, s)
}

func autoConvert_v1alpha1_ETCD_To_config_ETCD(in *ETCD, out *config.ETCD, s conversion.Scope) error {
	if err := Convert_v1alpha1_ETCDStorage_To_config_ETCDStorage(&in.Storage, &out.Storage, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_ETCDBackup_To_config_ETCDBackup(&in.Backup, &out.Backup, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_ETCD_To_config_ETCD is an autogenerated conversion function.
func Convert_v1alpha1_ETCD_To_config_ETCD(in *ETCD, out *config.ETCD, s conversion.Scope) error {
	return autoConvert_v1alpha1_ETCD_To_config_ETCD(in, out, s)
}

func autoConvert_config_ETCD_To_v1alpha1_ETCD(in *config.ETCD, out *ETCD, s conversion.Scope) error {
	if err := Convert_config_ETCDStorage_To_v1alpha1_ETCDStorage(&in.Storage, &out.Storage, s); err != nil {
		return err
	}
	if err := Convert_config_ETCDBackup_To_v1alpha1_ETCDBackup(&in.Backup, &out.Backup, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_ETCD_To_v1alpha1_ETCD is an autogenerated conversion function.
func Convert_config_ETCD_To_v1alpha1_ETCD(in *config.ETCD, out *ETCD, s conversion.Scope) error {
	return autoConvert_config_ETCD_To_v1alpha1_ETCD(in, out, s)
}

func autoConvert_v1alpha1_ETCDBackup_To_config_ETCDBackup(in *ETCDBackup, out *config.ETCDBackup, s conversion.Scope) error {
	out.Schedule = (*string)(unsafe.Pointer(in.Schedule))
	out.DeltaSnapshotPeriod = (*string)(unsafe.Pointer(in.DeltaSnapshotPeriod))
	return nil
}

// Convert_v1alpha1_ETCDBackup_To_config_ETCDBackup is an autogenerated conversion function.
func Convert_v1alpha1_ETCDBackup_To_config_ETCDBackup(in *ETCDBackup, out *config.ETCDBackup, s conversion.Scope) error {
	return autoConvert_v1alpha1_ETCDBackup_To_config_ETCDBackup(in, out, s)
}

func autoConvert_config_ETCDBackup_To_v1alpha1_ETCDBackup(in *config.ETCDBackup, out *ETCDBackup, s conversion.Scope) error {
	out.Schedule = (*string)(unsafe.Pointer(in.Schedule))
	out.DeltaSnapshotPeriod = (*string)(unsafe.Pointer(in.DeltaSnapshotPeriod))
	return nil
}

// Convert_config_ETCDBackup_To_v1alpha1_ETCDBackup is an autogenerated conversion function.
func Convert_config_ETCDBackup_To_v1alpha1_ETCDBackup(in *config.ETCDBackup, out *ETCDBackup, s conversion.Scope) error {
	return autoConvert_config_ETCDBackup_To_v1alpha1_ETCDBackup(in, out, s)
}

func autoConvert_v1alpha1_ETCDStorage_To_config_ETCDStorage(in *ETCDStorage, out *config.ETCDStorage, s conversion.Scope) error {
	out.ClassName = (*string)(unsafe.Pointer(in.ClassName))
	out.Capacity = (*resource.Quantity)(unsafe.Pointer(in.Capacity))
	return nil
}

// Convert_v1alpha1_ETCDStorage_To_config_ETCDStorage is an autogenerated conversion function.
func Convert_v1alpha1_ETCDStorage_To_config_ETCDStorage(in *ETCDStorage, out *config.ETCDStorage, s conversion.Scope) error {
	return autoConvert_v1alpha1_ETCDStorage_To_config_ETCDStorage(in, out, s)
}

func autoConvert_config_ETCDStorage_To_v1alpha1_ETCDStorage(in *config.ETCDStorage, out *ETCDStorage, s conversion.Scope) error {
	out.ClassName = (*string)(unsafe.Pointer(in.ClassName))
	out.Capacity = (*resource.Quantity)(unsafe.Pointer(in.Capacity))
	return nil
}

// Convert_config_ETCDStorage_To_v1alpha1_ETCDStorage is an autogenerated conversion function.
func Convert_config_ETCDStorage_To_v1alpha1_ETCDStorage(in *config.ETCDStorage, out *ETCDStorage, s conversion.Scope) error {
	return autoConvert_config_ETCDStorage_To_v1alpha1_ETCDStorage(in, out, s)
}

func autoConvert_v1alpha1_MachineImage_To_config_MachineImage(in *MachineImage, out *config.MachineImage, s conversion.Scope) error {
	out.Name = in.Name
	out.Version = in.Version
	out.Image = in.Image
	return nil
}

// Convert_v1alpha1_MachineImage_To_config_MachineImage is an autogenerated conversion function.
func Convert_v1alpha1_MachineImage_To_config_MachineImage(in *MachineImage, out *config.MachineImage, s conversion.Scope) error {
	return autoConvert_v1alpha1_MachineImage_To_config_MachineImage(in, out, s)
}

func autoConvert_config_MachineImage_To_v1alpha1_MachineImage(in *config.MachineImage, out *MachineImage, s conversion.Scope) error {
	out.Name = in.Name
	out.Version = in.Version
	out.Image = in.Image
	return nil
}

// Convert_config_MachineImage_To_v1alpha1_MachineImage is an autogenerated conversion function.
func Convert_config_MachineImage_To_v1alpha1_MachineImage(in *config.MachineImage, out *MachineImage, s conversion.Scope) error {
	return autoConvert_config_MachineImage_To_v1alpha1_MachineImage(in, out, s)
}

func autoConvert_v1alpha1_SplunkAudit_To_config_SplunkAudit(in *SplunkAudit, out *config.SplunkAudit, s conversion.Scope) error {
	out.Enabled = in.Enabled
	out.HecURL = in.HecURL
	out.HecToken = in.HecToken
	return nil
}

// Convert_v1alpha1_SplunkAudit_To_config_SplunkAudit is an autogenerated conversion function.
func Convert_v1alpha1_SplunkAudit_To_config_SplunkAudit(in *SplunkAudit, out *config.SplunkAudit, s conversion.Scope) error {
	return autoConvert_v1alpha1_SplunkAudit_To_config_SplunkAudit(in, out, s)
}

func autoConvert_config_SplunkAudit_To_v1alpha1_SplunkAudit(in *config.SplunkAudit, out *SplunkAudit, s conversion.Scope) error {
	out.Enabled = in.Enabled
	out.HecURL = in.HecURL
	out.HecToken = in.HecToken
	return nil
}

// Convert_config_SplunkAudit_To_v1alpha1_SplunkAudit is an autogenerated conversion function.
func Convert_config_SplunkAudit_To_v1alpha1_SplunkAudit(in *config.SplunkAudit, out *SplunkAudit, s conversion.Scope) error {
	return autoConvert_config_SplunkAudit_To_v1alpha1_SplunkAudit(in, out, s)
}
