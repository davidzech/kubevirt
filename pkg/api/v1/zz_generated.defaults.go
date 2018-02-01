// +build !ignore_autogenerated

/*
Copyright 2018 The KubeVirt Authors.

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

// This file was autogenerated by defaulter-gen. Do not edit it manually!

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&OfflineVirtualMachine{}, func(obj interface{}) { SetObjectDefaults_OfflineVirtualMachine(obj.(*OfflineVirtualMachine)) })
	scheme.AddTypeDefaultingFunc(&OfflineVirtualMachineList{}, func(obj interface{}) { SetObjectDefaults_OfflineVirtualMachineList(obj.(*OfflineVirtualMachineList)) })
	scheme.AddTypeDefaultingFunc(&VirtualMachine{}, func(obj interface{}) { SetObjectDefaults_VirtualMachine(obj.(*VirtualMachine)) })
	scheme.AddTypeDefaultingFunc(&VirtualMachineList{}, func(obj interface{}) { SetObjectDefaults_VirtualMachineList(obj.(*VirtualMachineList)) })
	scheme.AddTypeDefaultingFunc(&VirtualMachinePreset{}, func(obj interface{}) { SetObjectDefaults_VirtualMachinePreset(obj.(*VirtualMachinePreset)) })
	scheme.AddTypeDefaultingFunc(&VirtualMachinePresetList{}, func(obj interface{}) { SetObjectDefaults_VirtualMachinePresetList(obj.(*VirtualMachinePresetList)) })
	scheme.AddTypeDefaultingFunc(&VirtualMachineReplicaSet{}, func(obj interface{}) { SetObjectDefaults_VirtualMachineReplicaSet(obj.(*VirtualMachineReplicaSet)) })
	scheme.AddTypeDefaultingFunc(&VirtualMachineReplicaSetList{}, func(obj interface{}) {
		SetObjectDefaults_VirtualMachineReplicaSetList(obj.(*VirtualMachineReplicaSetList))
	})
	return nil
}

func SetObjectDefaults_OfflineVirtualMachine(in *OfflineVirtualMachine) {
	if in.Spec.Template != nil {
		if in.Spec.Template.Domain.Firmware != nil {
			SetDefaults_Firmware(in.Spec.Template.Domain.Firmware)
		}
		if in.Spec.Template.Domain.Clock != nil {
			if in.Spec.Template.Domain.Clock.Timer != nil {
				if in.Spec.Template.Domain.Clock.Timer.HPET != nil {
					SetDefaults_HPETTimer(in.Spec.Template.Domain.Clock.Timer.HPET)
				}
				if in.Spec.Template.Domain.Clock.Timer.KVM != nil {
					SetDefaults_KVMTimer(in.Spec.Template.Domain.Clock.Timer.KVM)
				}
				if in.Spec.Template.Domain.Clock.Timer.PIT != nil {
					SetDefaults_PITTimer(in.Spec.Template.Domain.Clock.Timer.PIT)
				}
				if in.Spec.Template.Domain.Clock.Timer.RTC != nil {
					SetDefaults_RTCTimer(in.Spec.Template.Domain.Clock.Timer.RTC)
				}
				if in.Spec.Template.Domain.Clock.Timer.Hyperv != nil {
					SetDefaults_HypervTimer(in.Spec.Template.Domain.Clock.Timer.Hyperv)
				}
			}
		}
		if in.Spec.Template.Domain.Features != nil {
			SetDefaults_FeatureState(&in.Spec.Template.Domain.Features.ACPI)
			if in.Spec.Template.Domain.Features.APIC != nil {
				SetDefaults_FeatureState(in.Spec.Template.Domain.Features.APIC)
			}
			if in.Spec.Template.Domain.Features.Hyperv != nil {
				if in.Spec.Template.Domain.Features.Hyperv.Relaxed != nil {
					SetDefaults_FeatureState(in.Spec.Template.Domain.Features.Hyperv.Relaxed)
				}
				if in.Spec.Template.Domain.Features.Hyperv.VAPIC != nil {
					SetDefaults_FeatureState(in.Spec.Template.Domain.Features.Hyperv.VAPIC)
				}
				if in.Spec.Template.Domain.Features.Hyperv.Spinlocks != nil {
					SetDefaults_FeatureSpinlocks(in.Spec.Template.Domain.Features.Hyperv.Spinlocks)
				}
				if in.Spec.Template.Domain.Features.Hyperv.VPIndex != nil {
					SetDefaults_FeatureState(in.Spec.Template.Domain.Features.Hyperv.VPIndex)
				}
				if in.Spec.Template.Domain.Features.Hyperv.Runtime != nil {
					SetDefaults_FeatureState(in.Spec.Template.Domain.Features.Hyperv.Runtime)
				}
				if in.Spec.Template.Domain.Features.Hyperv.SyNIC != nil {
					SetDefaults_FeatureState(in.Spec.Template.Domain.Features.Hyperv.SyNIC)
				}
				if in.Spec.Template.Domain.Features.Hyperv.SyNICTimer != nil {
					SetDefaults_FeatureState(in.Spec.Template.Domain.Features.Hyperv.SyNICTimer)
				}
				if in.Spec.Template.Domain.Features.Hyperv.Reset != nil {
					SetDefaults_FeatureState(in.Spec.Template.Domain.Features.Hyperv.Reset)
				}
				if in.Spec.Template.Domain.Features.Hyperv.VendorID != nil {
					SetDefaults_FeatureVendorID(in.Spec.Template.Domain.Features.Hyperv.VendorID)
				}
			}
		}
		for i := range in.Spec.Template.Domain.Devices.Disks {
			a := &in.Spec.Template.Domain.Devices.Disks[i]
			SetDefaults_DiskDevice(&a.DiskDevice)
			if a.DiskDevice.Floppy != nil {
				SetDefaults_FloppyTarget(a.DiskDevice.Floppy)
			}
			if a.DiskDevice.CDRom != nil {
				SetDefaults_CDRomTarget(a.DiskDevice.CDRom)
			}
		}
		if in.Spec.Template.Domain.Devices.Watchdog != nil {
			SetDefaults_Watchdog(in.Spec.Template.Domain.Devices.Watchdog)
			if in.Spec.Template.Domain.Devices.Watchdog.WatchdogDevice.I6300ESB != nil {
				SetDefaults_I6300ESBWatchdog(in.Spec.Template.Domain.Devices.Watchdog.WatchdogDevice.I6300ESB)
			}
		}
	}
}

func SetObjectDefaults_OfflineVirtualMachineList(in *OfflineVirtualMachineList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_OfflineVirtualMachine(a)
	}
}

func SetObjectDefaults_VirtualMachine(in *VirtualMachine) {
	SetDefaults_VirtualMachine(in)
	if in.Spec.Domain.Firmware != nil {
		SetDefaults_Firmware(in.Spec.Domain.Firmware)
	}
	if in.Spec.Domain.Clock != nil {
		if in.Spec.Domain.Clock.Timer != nil {
			if in.Spec.Domain.Clock.Timer.HPET != nil {
				SetDefaults_HPETTimer(in.Spec.Domain.Clock.Timer.HPET)
			}
			if in.Spec.Domain.Clock.Timer.KVM != nil {
				SetDefaults_KVMTimer(in.Spec.Domain.Clock.Timer.KVM)
			}
			if in.Spec.Domain.Clock.Timer.PIT != nil {
				SetDefaults_PITTimer(in.Spec.Domain.Clock.Timer.PIT)
			}
			if in.Spec.Domain.Clock.Timer.RTC != nil {
				SetDefaults_RTCTimer(in.Spec.Domain.Clock.Timer.RTC)
			}
			if in.Spec.Domain.Clock.Timer.Hyperv != nil {
				SetDefaults_HypervTimer(in.Spec.Domain.Clock.Timer.Hyperv)
			}
		}
	}
	if in.Spec.Domain.Features != nil {
		SetDefaults_FeatureState(&in.Spec.Domain.Features.ACPI)
		if in.Spec.Domain.Features.APIC != nil {
			SetDefaults_FeatureAPIC(in.Spec.Domain.Features.APIC)
		}
		if in.Spec.Domain.Features.Hyperv != nil {
			if in.Spec.Domain.Features.Hyperv.Relaxed != nil {
				SetDefaults_FeatureState(in.Spec.Domain.Features.Hyperv.Relaxed)
			}
			if in.Spec.Domain.Features.Hyperv.VAPIC != nil {
				SetDefaults_FeatureState(in.Spec.Domain.Features.Hyperv.VAPIC)
			}
			if in.Spec.Domain.Features.Hyperv.Spinlocks != nil {
				SetDefaults_FeatureSpinlocks(in.Spec.Domain.Features.Hyperv.Spinlocks)
			}
			if in.Spec.Domain.Features.Hyperv.VPIndex != nil {
				SetDefaults_FeatureState(in.Spec.Domain.Features.Hyperv.VPIndex)
			}
			if in.Spec.Domain.Features.Hyperv.Runtime != nil {
				SetDefaults_FeatureState(in.Spec.Domain.Features.Hyperv.Runtime)
			}
			if in.Spec.Domain.Features.Hyperv.SyNIC != nil {
				SetDefaults_FeatureState(in.Spec.Domain.Features.Hyperv.SyNIC)
			}
			if in.Spec.Domain.Features.Hyperv.SyNICTimer != nil {
				SetDefaults_FeatureState(in.Spec.Domain.Features.Hyperv.SyNICTimer)
			}
			if in.Spec.Domain.Features.Hyperv.Reset != nil {
				SetDefaults_FeatureState(in.Spec.Domain.Features.Hyperv.Reset)
			}
			if in.Spec.Domain.Features.Hyperv.VendorID != nil {
				SetDefaults_FeatureVendorID(in.Spec.Domain.Features.Hyperv.VendorID)
			}
		}
	}
	for i := range in.Spec.Domain.Devices.Disks {
		a := &in.Spec.Domain.Devices.Disks[i]
		SetDefaults_DiskDevice(&a.DiskDevice)
		if a.DiskDevice.Floppy != nil {
			SetDefaults_FloppyTarget(a.DiskDevice.Floppy)
		}
		if a.DiskDevice.CDRom != nil {
			SetDefaults_CDRomTarget(a.DiskDevice.CDRom)
		}
	}
	if in.Spec.Domain.Devices.Watchdog != nil {
		SetDefaults_Watchdog(in.Spec.Domain.Devices.Watchdog)
		if in.Spec.Domain.Devices.Watchdog.WatchdogDevice.I6300ESB != nil {
			SetDefaults_I6300ESBWatchdog(in.Spec.Domain.Devices.Watchdog.WatchdogDevice.I6300ESB)
		}
	}
}

func SetObjectDefaults_VirtualMachineList(in *VirtualMachineList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_VirtualMachine(a)
	}
}

func SetObjectDefaults_VirtualMachinePreset(in *VirtualMachinePreset) {
	if in.Spec.Domain != nil {
		if in.Spec.Domain.Firmware != nil {
			SetDefaults_Firmware(in.Spec.Domain.Firmware)
		}
		if in.Spec.Domain.Clock != nil {
			if in.Spec.Domain.Clock.Timer != nil {
				if in.Spec.Domain.Clock.Timer.HPET != nil {
					SetDefaults_HPETTimer(in.Spec.Domain.Clock.Timer.HPET)
				}
				if in.Spec.Domain.Clock.Timer.KVM != nil {
					SetDefaults_KVMTimer(in.Spec.Domain.Clock.Timer.KVM)
				}
				if in.Spec.Domain.Clock.Timer.PIT != nil {
					SetDefaults_PITTimer(in.Spec.Domain.Clock.Timer.PIT)
				}
				if in.Spec.Domain.Clock.Timer.RTC != nil {
					SetDefaults_RTCTimer(in.Spec.Domain.Clock.Timer.RTC)
				}
				if in.Spec.Domain.Clock.Timer.Hyperv != nil {
					SetDefaults_HypervTimer(in.Spec.Domain.Clock.Timer.Hyperv)
				}
			}
		}
		if in.Spec.Domain.Features != nil {
			SetDefaults_FeatureState(&in.Spec.Domain.Features.ACPI)
			if in.Spec.Domain.Features.APIC != nil {
				SetDefaults_FeatureAPIC(in.Spec.Domain.Features.APIC)
			}
			if in.Spec.Domain.Features.Hyperv != nil {
				if in.Spec.Domain.Features.Hyperv.Relaxed != nil {
					SetDefaults_FeatureState(in.Spec.Domain.Features.Hyperv.Relaxed)
				}
				if in.Spec.Domain.Features.Hyperv.VAPIC != nil {
					SetDefaults_FeatureState(in.Spec.Domain.Features.Hyperv.VAPIC)
				}
				if in.Spec.Domain.Features.Hyperv.Spinlocks != nil {
					SetDefaults_FeatureSpinlocks(in.Spec.Domain.Features.Hyperv.Spinlocks)
				}
				if in.Spec.Domain.Features.Hyperv.VPIndex != nil {
					SetDefaults_FeatureState(in.Spec.Domain.Features.Hyperv.VPIndex)
				}
				if in.Spec.Domain.Features.Hyperv.Runtime != nil {
					SetDefaults_FeatureState(in.Spec.Domain.Features.Hyperv.Runtime)
				}
				if in.Spec.Domain.Features.Hyperv.SyNIC != nil {
					SetDefaults_FeatureState(in.Spec.Domain.Features.Hyperv.SyNIC)
				}
				if in.Spec.Domain.Features.Hyperv.SyNICTimer != nil {
					SetDefaults_FeatureState(in.Spec.Domain.Features.Hyperv.SyNICTimer)
				}
				if in.Spec.Domain.Features.Hyperv.Reset != nil {
					SetDefaults_FeatureState(in.Spec.Domain.Features.Hyperv.Reset)
				}
				if in.Spec.Domain.Features.Hyperv.VendorID != nil {
					SetDefaults_FeatureVendorID(in.Spec.Domain.Features.Hyperv.VendorID)
				}
			}
		}
		for i := range in.Spec.Domain.Devices.Disks {
			a := &in.Spec.Domain.Devices.Disks[i]
			SetDefaults_DiskDevice(&a.DiskDevice)
			if a.DiskDevice.Floppy != nil {
				SetDefaults_FloppyTarget(a.DiskDevice.Floppy)
			}
			if a.DiskDevice.CDRom != nil {
				SetDefaults_CDRomTarget(a.DiskDevice.CDRom)
			}
		}
		if in.Spec.Domain.Devices.Watchdog != nil {
			SetDefaults_Watchdog(in.Spec.Domain.Devices.Watchdog)
			if in.Spec.Domain.Devices.Watchdog.WatchdogDevice.I6300ESB != nil {
				SetDefaults_I6300ESBWatchdog(in.Spec.Domain.Devices.Watchdog.WatchdogDevice.I6300ESB)
			}
		}
	}
}

func SetObjectDefaults_VirtualMachinePresetList(in *VirtualMachinePresetList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_VirtualMachinePreset(a)
	}
}

func SetObjectDefaults_VirtualMachineReplicaSet(in *VirtualMachineReplicaSet) {
	if in.Spec.Template != nil {
		if in.Spec.Template.Spec.Domain.Firmware != nil {
			SetDefaults_Firmware(in.Spec.Template.Spec.Domain.Firmware)
		}
		if in.Spec.Template.Spec.Domain.Clock != nil {
			if in.Spec.Template.Spec.Domain.Clock.Timer != nil {
				if in.Spec.Template.Spec.Domain.Clock.Timer.HPET != nil {
					SetDefaults_HPETTimer(in.Spec.Template.Spec.Domain.Clock.Timer.HPET)
				}
				if in.Spec.Template.Spec.Domain.Clock.Timer.KVM != nil {
					SetDefaults_KVMTimer(in.Spec.Template.Spec.Domain.Clock.Timer.KVM)
				}
				if in.Spec.Template.Spec.Domain.Clock.Timer.PIT != nil {
					SetDefaults_PITTimer(in.Spec.Template.Spec.Domain.Clock.Timer.PIT)
				}
				if in.Spec.Template.Spec.Domain.Clock.Timer.RTC != nil {
					SetDefaults_RTCTimer(in.Spec.Template.Spec.Domain.Clock.Timer.RTC)
				}
				if in.Spec.Template.Spec.Domain.Clock.Timer.Hyperv != nil {
					SetDefaults_HypervTimer(in.Spec.Template.Spec.Domain.Clock.Timer.Hyperv)
				}
			}
		}
		if in.Spec.Template.Spec.Domain.Features != nil {
			SetDefaults_FeatureState(&in.Spec.Template.Spec.Domain.Features.ACPI)
			if in.Spec.Template.Spec.Domain.Features.APIC != nil {
				SetDefaults_FeatureAPIC(in.Spec.Template.Spec.Domain.Features.APIC)
			}
			if in.Spec.Template.Spec.Domain.Features.Hyperv != nil {
				if in.Spec.Template.Spec.Domain.Features.Hyperv.Relaxed != nil {
					SetDefaults_FeatureState(in.Spec.Template.Spec.Domain.Features.Hyperv.Relaxed)
				}
				if in.Spec.Template.Spec.Domain.Features.Hyperv.VAPIC != nil {
					SetDefaults_FeatureState(in.Spec.Template.Spec.Domain.Features.Hyperv.VAPIC)
				}
				if in.Spec.Template.Spec.Domain.Features.Hyperv.Spinlocks != nil {
					SetDefaults_FeatureSpinlocks(in.Spec.Template.Spec.Domain.Features.Hyperv.Spinlocks)
				}
				if in.Spec.Template.Spec.Domain.Features.Hyperv.VPIndex != nil {
					SetDefaults_FeatureState(in.Spec.Template.Spec.Domain.Features.Hyperv.VPIndex)
				}
				if in.Spec.Template.Spec.Domain.Features.Hyperv.Runtime != nil {
					SetDefaults_FeatureState(in.Spec.Template.Spec.Domain.Features.Hyperv.Runtime)
				}
				if in.Spec.Template.Spec.Domain.Features.Hyperv.SyNIC != nil {
					SetDefaults_FeatureState(in.Spec.Template.Spec.Domain.Features.Hyperv.SyNIC)
				}
				if in.Spec.Template.Spec.Domain.Features.Hyperv.SyNICTimer != nil {
					SetDefaults_FeatureState(in.Spec.Template.Spec.Domain.Features.Hyperv.SyNICTimer)
				}
				if in.Spec.Template.Spec.Domain.Features.Hyperv.Reset != nil {
					SetDefaults_FeatureState(in.Spec.Template.Spec.Domain.Features.Hyperv.Reset)
				}
				if in.Spec.Template.Spec.Domain.Features.Hyperv.VendorID != nil {
					SetDefaults_FeatureVendorID(in.Spec.Template.Spec.Domain.Features.Hyperv.VendorID)
				}
			}
		}
		for i := range in.Spec.Template.Spec.Domain.Devices.Disks {
			a := &in.Spec.Template.Spec.Domain.Devices.Disks[i]
			SetDefaults_DiskDevice(&a.DiskDevice)
			if a.DiskDevice.Floppy != nil {
				SetDefaults_FloppyTarget(a.DiskDevice.Floppy)
			}
			if a.DiskDevice.CDRom != nil {
				SetDefaults_CDRomTarget(a.DiskDevice.CDRom)
			}
		}
		if in.Spec.Template.Spec.Domain.Devices.Watchdog != nil {
			SetDefaults_Watchdog(in.Spec.Template.Spec.Domain.Devices.Watchdog)
			if in.Spec.Template.Spec.Domain.Devices.Watchdog.WatchdogDevice.I6300ESB != nil {
				SetDefaults_I6300ESBWatchdog(in.Spec.Template.Spec.Domain.Devices.Watchdog.WatchdogDevice.I6300ESB)
			}
		}
	}
}

func SetObjectDefaults_VirtualMachineReplicaSetList(in *VirtualMachineReplicaSetList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_VirtualMachineReplicaSet(a)
	}
}
