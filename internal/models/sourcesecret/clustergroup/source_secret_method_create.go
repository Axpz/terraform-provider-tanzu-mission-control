// Code generated by go-swagger; DO NOT EDIT.

package clustergroupsourcesecret

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1ClustergroupFluxcdSourceSecretRequest Request to create a SourceSecret.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.fluxcd.sourcesecret.CreateSourceSecretRequest
type VmwareTanzuManageV1alpha1ClustergroupFluxcdSourceSecretRequest struct {

	// SourceSecret to create.
	SourceSecret *VmwareTanzuManageV1alpha1ClustergroupFluxcdSourcesecretSourceSecret `json:"sourceSecret,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupFluxcdSourceSecretRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupFluxcdSourceSecretRequest) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupFluxcdSourceSecretRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VmwareTanzuManageV1alpha1ClustergroupFluxcdSourceSecretResponse Response from creating a SourceSecret.
//
// swagger:model vmware.tanzu.manage.v1alpha1.clustergroup.fluxcd.sourcesecret.CreateSourceSecretResponse
type VmwareTanzuManageV1alpha1ClustergroupFluxcdSourceSecretResponse struct {

	// SourceSecret created.
	SourceSecret *VmwareTanzuManageV1alpha1ClustergroupFluxcdSourcesecretSourceSecret `json:"sourceSecret,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupFluxcdSourceSecretResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClustergroupFluxcdSourceSecretResponse) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClustergroupFluxcdSourceSecretResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
