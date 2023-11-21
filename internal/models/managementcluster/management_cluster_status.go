/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/
// Code generated by go-swagger; DO NOT EDIT.

package managementclustermodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	clustermodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/cluster"
	statusmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/status"
)

// VmwareTanzuManageV1alpha1ManagementclusterStatus The ManagementCluster status.
//
// swagger:model vmware.tanzu.manage.v1alpha1.managementcluster.Status
type VmwareTanzuManageV1alpha1ManagementclusterStatus struct {

	// CPU allocation of a cluster.
	AllocatedCPU *clustermodel.VmwareTanzuManageV1alpha1CommonClusterResourceAllocation `json:"allocatedCpu,omitempty"`

	// Memory allocation of a cluster.
	AllocatedMemory *clustermodel.VmwareTanzuManageV1alpha1CommonClusterResourceAllocation `json:"allocatedMemory,omitempty"`

	// Conditions of the resource.
	Conditions map[string]statusmodel.VmwareTanzuCoreV1alpha1StatusCondition `json:"conditions,omitempty"`

	// List of the extensions running on the ManagementCluster Management Cluster.
	Extensions []string `json:"extensions"`

	// Health of a resource.
	Health *clustermodel.VmwareTanzuManageV1alpha1CommonClusterHealth `json:"health,omitempty"`

	// Health details of the ManagementCluster.
	HealthDetails *clustermodel.VmwareTanzuManageV1alpha1CommonClusterHealthInfo `json:"healthDetails,omitempty"`

	// Infrastructure provider.
	InfrastructureProvider *clustermodel.VmwareTanzuManageV1alpha1CommonClusterInfrastructureProvider `json:"infrastructureProvider,omitempty"`

	// Kubernetes Server Git Version.
	KubeServerVersion string `json:"kubeServerVersion,omitempty"`

	// Kubernetes Provider which is discovered after registration.
	KubernetesProvider *clustermodel.VmwareTanzuManageV1alpha1CommonClusterKubernetesProvider `json:"kubernetesProvider,omitempty"`

	// Timestamp when metadata was last updated on this cluster.
	// Format: date-time
	LastUpdate strfmt.DateTime `json:"lastUpdate,omitempty"`

	// Phase of the resource.
	Phase *VmwareTanzuManageV1alpha1ManagementclusterPhase `json:"phase,omitempty"`

	// Region.
	Region string `json:"region,omitempty"`

	// URL to fetch the TMC registration YAML.
	// If the management cluster is registered with proxy, Get on this
	// URL would need user token with sufficient permission to read the
	// proxy set during the registration set. In all other cases, this
	// URL can be fetched without user token.
	RegistrationURL string `json:"registrationUrl,omitempty"`
}

// MarshalBinary interface implementation
func (m *VmwareTanzuManageV1alpha1ManagementclusterStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VmwareTanzuManageV1alpha1ManagementclusterStatus) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ManagementclusterStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}