/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/
// Code generated by go-swagger; DO NOT EDIT.

package credentialmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// VmwareTanzuManageV1alpha1AccountCredentialStatus Status of the credential.
//
// swagger:model vmware.tanzu.manage.v1alpha1.account.credential.Status
type VmwareTanzuManageV1alpha1AccountCredentialStatus struct {
	// Phase of the credential.
	Phase *VmwareTanzuManageV1alpha1AccountCredentialStatusPhase `json:"phase,omitempty"`

	// Additional information about the phase.
	PhaseInfo string `json:"phaseInfo,omitempty"`
}

// VmwareTanzuManageV1alpha1AccountCredentialStatusPhase The overall phase of a credential.
//
//   - PHASE_UNSPECIFIED: Unspecified phase.
//   - CREATED: The credential is created and can be used.
//   - VALIDATING: The credential's capabilities are being validated by the intended service.
//
// Credentials can be used even if they have not been validated- this phase
// is set by the intended service if it validates credentials.
//   - VALID: The credential satisfies the intended service's requirements.
//   - INVALID: The credential does not satisfy the intended service's requirements.
//
// Invalid credentials might require user action to fix their permissions- this information
// is provided by the intended service.
//   - ERROR: An error occurred while the credential was being created or validated.
//   - DELETING: The credential clean up has begun.
//   - DELETED: The credential clean up has completed and will be removed from TMC.
//
// swagger:model vmware.tanzu.manage.v1alpha1.account.credential.Status.Phase
type VmwareTanzuManageV1alpha1AccountCredentialStatusPhase string

func NewVmwareTanzuManageV1alpha1AccountCredentialStatusPhase(value VmwareTanzuManageV1alpha1AccountCredentialStatusPhase) *VmwareTanzuManageV1alpha1AccountCredentialStatusPhase {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmwareTanzuManageV1alpha1AccountCredentialStatusPhase.
func (m VmwareTanzuManageV1alpha1AccountCredentialStatusPhase) Pointer() *VmwareTanzuManageV1alpha1AccountCredentialStatusPhase {
	return &m
}

const (

	// VmwareTanzuManageV1alpha1AccountCredentialStatusPhasePHASEUNSPECIFIED captures enum value "PHASE_UNSPECIFIED"
	VmwareTanzuManageV1alpha1AccountCredentialStatusPhasePHASEUNSPECIFIED VmwareTanzuManageV1alpha1AccountCredentialStatusPhase = "PHASE_UNSPECIFIED"

	// VmwareTanzuManageV1alpha1AccountCredentialStatusPhaseCREATED captures enum value "CREATED"
	VmwareTanzuManageV1alpha1AccountCredentialStatusPhaseCREATED VmwareTanzuManageV1alpha1AccountCredentialStatusPhase = "CREATED"

	// VmwareTanzuManageV1alpha1AccountCredentialStatusPhaseVALIDATING captures enum value "VALIDATING"
	VmwareTanzuManageV1alpha1AccountCredentialStatusPhaseVALIDATING VmwareTanzuManageV1alpha1AccountCredentialStatusPhase = "VALIDATING"

	// VmwareTanzuManageV1alpha1AccountCredentialStatusPhaseVALID captures enum value "VALID"
	VmwareTanzuManageV1alpha1AccountCredentialStatusPhaseVALID VmwareTanzuManageV1alpha1AccountCredentialStatusPhase = "VALID"

	// VmwareTanzuManageV1alpha1AccountCredentialStatusPhaseINVALID captures enum value "INVALID"
	VmwareTanzuManageV1alpha1AccountCredentialStatusPhaseINVALID VmwareTanzuManageV1alpha1AccountCredentialStatusPhase = "INVALID"

	// VmwareTanzuManageV1alpha1AccountCredentialStatusPhaseERROR captures enum value "ERROR"
	VmwareTanzuManageV1alpha1AccountCredentialStatusPhaseERROR VmwareTanzuManageV1alpha1AccountCredentialStatusPhase = "ERROR"

	// VmwareTanzuManageV1alpha1AccountCredentialStatusPhaseDELETING captures enum value "DELETING"
	VmwareTanzuManageV1alpha1AccountCredentialStatusPhaseDELETING VmwareTanzuManageV1alpha1AccountCredentialStatusPhase = "DELETING"

	// VmwareTanzuManageV1alpha1AccountCredentialStatusPhaseDELETED captures enum value "DELETED"
	VmwareTanzuManageV1alpha1AccountCredentialStatusPhaseDELETED VmwareTanzuManageV1alpha1AccountCredentialStatusPhase = "DELETED"
)

// for schema
var vmwareTanzuManageV1alpha1AccountCredentialStatusPhaseEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1AccountCredentialStatusPhase
	if err := json.Unmarshal([]byte(`["PHASE_UNSPECIFIED","CREATED","VALIDATING","VALID","INVALID","ERROR","DELETING","DELETED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vmwareTanzuManageV1alpha1AccountCredentialStatusPhaseEnum = append(vmwareTanzuManageV1alpha1AccountCredentialStatusPhaseEnum, v)
	}
}