/*
Copyright © 2021 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package cluster

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/authctx"
	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/helper"
	clustermodel "gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/models/cluster"
	"gitlab.eng.vmware.com/olympus/terraform-provider-tanzu/internal/resources/common"
)

func DataSourceTMCCluster() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTMCClusterRead,
		Schema:      clusterSchema,
	}
}

func dataSourceTMCClusterRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := m.(authctx.TanzuContext)

	// Warning or errors can be collected in a slice type
	var (
		diags diag.Diagnostics
		resp  *clustermodel.VmwareTanzuManageV1alpha1ClusterGetClusterResponse
		err   error
	)

	getClusterResourceRetryableFn := func() (retry bool, err error) {
		resp, err = config.TMCConnection.ClusterResourceService.ManageV1alpha1ClusterResourceServiceGet(constructFullname(d))
		if err != nil {
			return true, errors.Wrapf(err, "Unable to get tanzu TMC cluster entry, name : %s", d.Get(clusterNameKey))
		}

		d.SetId(resp.Cluster.Meta.UID)

		if *(resp.Cluster.Status.Phase) != clustermodel.VmwareTanzuManageV1alpha1ClusterPhaseREADY {
			return true, nil
		}

		return false, nil
	}

	switch value, _ := d.Get(waitKey).(bool); value {
	case true:
		_, err = helper.Retry(getClusterResourceRetryableFn, 10*time.Second, 18)
	case false:
		_, err = getClusterResourceRetryableFn()
	}

	if err != nil || resp == nil {
		return diag.FromErr(errors.Wrapf(err, "Unable to get tanzu TMC cluster entry, name : %s", d.Get(clusterNameKey)))
	}

	// always run
	d.SetId(resp.Cluster.Meta.UID)

	status := map[string]interface{}{
		"type":                  resp.Cluster.Status.Type,
		"phase":                 resp.Cluster.Status.Phase,
		"health":                resp.Cluster.Status.Health,
		"k8s_version":           resp.Cluster.Status.KubeServerVersion,
		"node_count":            resp.Cluster.Status.NodeCount,
		"k8s_provider_type":     resp.Cluster.Status.KubernetesProvider.Type,
		"k8s_provider_version":  resp.Cluster.Status.KubernetesProvider.Version,
		"infra_provider":        resp.Cluster.Status.InfrastructureProvider,
		"infra_provider_region": resp.Cluster.Status.InfrastructureProviderRegion,
	}

	if resp.Cluster.FullName.ManagementClusterName == "attached" && resp.Cluster.Status.InstallerLink != "" {
		status["installer_link"] = resp.Cluster.Status.InstallerLink
		status["execution_cmd"] = fmt.Sprintf("kubectl create -f '%s'", resp.Cluster.Status.InstallerLink)
	}

	if err := d.Set(statusKey, status); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set(common.MetaKey, common.FlattenMeta(resp.Cluster.Meta)); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set(specKey, flattenSpec(resp.Cluster.Spec)); err != nil {
		return diag.FromErr(err)
	}

	return diags
}