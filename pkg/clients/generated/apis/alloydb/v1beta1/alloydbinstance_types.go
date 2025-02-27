// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type InstanceMachineConfig struct {
	/* The number of CPU's in the VM instance. */
	// +optional
	CpuCount *int `json:"cpuCount,omitempty"`
}

type InstanceReadPoolConfig struct {
	/* Read capacity, i.e. number of nodes in a read pool instance. */
	// +optional
	NodeCount *int `json:"nodeCount,omitempty"`
}

type AlloyDBInstanceSpec struct {
	/* Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels. */
	// +optional
	Annotations map[string]string `json:"annotations,omitempty"`

	/* 'Availability type of an Instance. Defaults to REGIONAL for both primary and read instances.
	Note that primary and read instances can have different availability types.
	Only READ_POOL instance supports ZONAL type. Users can't specify the zone for READ_POOL instance.
	Zone is automatically chosen from the list of zones in the region specified.
	Read pool of size 1 can only have zonal availability. Read pools with node count of 2 or more
	can have regional availability (nodes are present in 2 or more zones in a region).' Possible values: ["AVAILABILITY_TYPE_UNSPECIFIED", "ZONAL", "REGIONAL"]. */
	// +optional
	AvailabilityType *string `json:"availabilityType,omitempty"`

	ClusterRef v1alpha1.ResourceRef `json:"clusterRef"`

	/* Database flags. Set at instance level. * They are copied from primary instance on read instance creation. * Read instances can set new or override existing flags that are relevant for reads, e.g. for enabling columnar cache on a read instance. Flags set on read instance may or may not be present on primary. */
	// +optional
	DatabaseFlags map[string]string `json:"databaseFlags,omitempty"`

	/* User-settable and human-readable display name for the Instance. */
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

	/* The Compute Engine zone that the instance should serve from, per https://cloud.google.com/compute/docs/regions-zones This can ONLY be specified for ZONAL instances. If present for a REGIONAL instance, an error will be thrown. If this is absent for a ZONAL instance, instance is created in a random zone with available capacity. */
	// +optional
	GceZone *string `json:"gceZone,omitempty"`

	/* Immutable. The type of the instance. If the instance type is READ_POOL, provide the associated PRIMARY instance in the 'depends_on' meta-data attribute. Possible values: ["PRIMARY", "READ_POOL"]. */
	InstanceType string `json:"instanceType"`

	/* Configurations for the machines that host the underlying database engine. */
	// +optional
	MachineConfig *InstanceMachineConfig `json:"machineConfig,omitempty"`

	/* Read pool specific config. If the instance type is READ_POOL, this configuration must be provided. */
	// +optional
	ReadPoolConfig *InstanceReadPoolConfig `json:"readPoolConfig,omitempty"`

	/* Immutable. Optional. The instanceId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type AlloyDBInstanceStatus struct {
	/* Conditions represent the latest available observations of the
	   AlloyDBInstance's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Time the Instance was created in UTC. */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	/* The IP address for the Instance. This is the connection endpoint for an end-user application. */
	// +optional
	IpAddress *string `json:"ipAddress,omitempty"`

	/* The name of the instance resource. */
	// +optional
	Name *string `json:"name,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* Set to true if the current state of Instance does not match the user's intended state, and the service is actively updating the resource to reconcile them. This can happen due to user-triggered updates or system actions like failover or maintenance. */
	// +optional
	Reconciling *bool `json:"reconciling,omitempty"`

	/* The current state of the alloydb instance. */
	// +optional
	State *string `json:"state,omitempty"`

	/* The system-generated UID of the resource. */
	// +optional
	Uid *string `json:"uid,omitempty"`

	/* Time the Instance was updated in UTC. */
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AlloyDBInstance is the Schema for the alloydb API
// +k8s:openapi-gen=true
type AlloyDBInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AlloyDBInstanceSpec   `json:"spec,omitempty"`
	Status AlloyDBInstanceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AlloyDBInstanceList contains a list of AlloyDBInstance
type AlloyDBInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlloyDBInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AlloyDBInstance{}, &AlloyDBInstanceList{})
}
