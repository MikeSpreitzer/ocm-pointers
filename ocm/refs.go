package ocm

import (
	addonV1a1 "open-cluster-management.io/api/addon/v1alpha1"
	clusterV1 "open-cluster-management.io/api/cluster/v1"
	clusterV1a1 "open-cluster-management.io/api/cluster/v1alpha1"
	clusterV1b1 "open-cluster-management.io/api/cluster/v1beta1"
	operatorV1 "open-cluster-management.io/api/operator/v1"
	workV1 "open-cluster-management.io/api/work/v1"
	policyV1 "open-cluster-management.io/config-policy-controller/api/v1"
	propolicyV1 "open-cluster-management.io/governance-policy-propagator/api/v1"
	chanappsV1 "open-cluster-management.io/multicloud-operators-channel/pkg/apis/apps/v1"
	ansibleV1a1 "open-cluster-management.io/multicloud-operators-subscription/pkg/apis/apps/ansible/v1alpha1"
	deployableV1 "open-cluster-management.io/multicloud-operators-subscription/pkg/apis/apps/deployable/v1"
	helmreleaseV1 "open-cluster-management.io/multicloud-operators-subscription/pkg/apis/apps/helmrelease/v1"
	placementruleV1 "open-cluster-management.io/multicloud-operators-subscription/pkg/apis/apps/placementrule/v1"
	appsV1 "open-cluster-management.io/multicloud-operators-subscription/pkg/apis/apps/v1"
	appsV1a1 "open-cluster-management.io/multicloud-operators-subscription/pkg/apis/apps/v1alpha1"
)

////////////////////////////////////////////////////////////////
// This file has pointers to the API objct types in the xxx.open-cluster-management.io kube API groups.

//////////////////////////////// Clusters & Machinery ////////////////////////////////

// ManagedCluster in cluster.open-cluster-management.io/v1
// ManagedCluster represents the desired state and current status of managed
// cluster. ManagedCluster is a cluster scoped resource. The name is the cluster
// UID.
type OCMClusterV1_ManagedCluster = clusterV1.ManagedCluster

// ManagedClusterSet in cluster.open-cluster-management.io/v1beta1
// ManagedClusterSet defines a group of ManagedClusters that user's workload can run on.
type OCMClusterV1b1_ManagedClusterSet = clusterV1b1.ManagedClusterSet

// Placement in cluster.open-cluster-management.io/v1beta1
// Placement defines a rule to select a set of ManagedClusters from the ManagedClusterSets bound
// to the placement namespace.
type OCMClusterV1b1_Placement = clusterV1b1.Placement

// AddOnPlacementScore in cluster.open-cluster-management.io/v1alpha1
// AddOnPlacementScore represents a bundle of scores of one managed cluster, which could be used by placement.
type OCMClusterV1a1_AddOnPlacementScore = clusterV1a1.AddOnPlacementScore

// ClusterClaim in cluster.open-cluster-management.io/v1alpha1
// ClusterClaim represents cluster information that a managed cluster claims
type OCMClusterV1a1_ClusterClaim = clusterV1a1.ClusterClaim

// ManagedClusterSet in cluster.open-cluster-management.io/v1alpha1
// ManagedClusterSet defines a group of ManagedClusters that user's workload can run on.
type OCMClusterV1a1_ManagedClusterSet = clusterV1a1.ManagedClusterSet

// PlacementDecision in cluster.open-cluster-management.io/v1alpha1
// PlacementDecision indicates a decision from a placement
type OCMClusterV1a1_PlacementDecision = clusterV1a1.PlacementDecision

// Placement in cluster.open-cluster-management.io/v1alpha1
// Placement defines a rule to select a set of ManagedClusters from the ManagedClusterSets bound
// to the placement namespace.
type OCMClusterV1a1_Placement = clusterV1a1.Placement

// ClusterManagementAddOn in addon.open-cluster-management.io/v1alpha1
// ClusterManagementAddOn represents the registration of an add-on to the cluster manager.
type OCMAddonV1a1_ClusterManagementAddOn = addonV1a1.ClusterManagementAddOn

// ManagedClusterAddOn in addon.open-cluster-management.io/v1alpha1
// ManagedClusterAddOn is the Custom Resource object which holds the current state
// of an add-on. This object is used by add-on operators to convey their state.
type OCMAddonV1a1_ManagedClusterAddOn = addonV1a1.ManagedClusterAddOn

// ClusterManager in operator.open-cluster-management.io/v1
// ClusterManager configures the controllers on the hub that govern registration and work distribution for attached Klusterlets.
type OCMOperatorV1_ClusterManager = operatorV1.ClusterManager

// Klusterlet in operator.open-cluster-management.io/v1
// Klusterlet represents controllers to install the resources for a managed cluster.
type OCMOperatorV1_Klusterlet = operatorV1.Klusterlet

//////////////////////////////// Apps ////////////////////////////////

// Channel in apps.open-cluster-management.io/v1
type OCMAppsV1_Channel = chanappsV1.Channel

// Subscription in apps.open-cluster-management.io/v1
type OCMAppsV1_Subscription = appsV1.Subscription

// Deployable in apps.open-cluster-management.io/v1
type OCMAppsV1_Deployable = deployableV1.Deployable

// HelmRelease in apps.open-cluster-management.io/v1
type OCMAppsV1_HelmRelease = helmreleaseV1.HelmRelease

// PlacementRule in apps.open-cluster-management.io/v1
type OCMAppsV1_PlacementRule = placementruleV1.PlacementRule

// SubscriptionReport in apps.open-cluster-management.io/v1alpha1
type OCMAppsV1a1_SubscriptionReport = appsV1a1.SubscriptionReport

// SubscriptionStatus in apps.open-cluster-management.io/v1alpha1
type OCMAppsV1a1_SubscriptionStatus = appsV1a1.SubscriptionStatus

// AppliedManifestWork in work.open-cluster-management.io/v1
// AppliedManifestWork represents an applied manifestwork on managed cluster that is placed
// on a managed cluster. An AppliedManifestWork links to a manifestwork on a hub recording resources
// deployed in the managed cluster.
type OCMWorkV1_AppliedManifestWork = workV1.AppliedManifestWork

// ManifestWork in work.open-cluster-management.io/v1
// ManifestWork represents a manifests workload that hub wants to deploy on the managed cluster.
// A manifest workload is defined as a set of Kubernetes resources.
type OCMWorkV1_ManifestWork = workV1.ManifestWork

// AnsibleJob in tower.ansible.com/v1alpha1
type AnsibleTowerV1a1_AnsibleJob = ansibleV1a1.AnsibleJob

//////////////////////////////// Governance, Risk, Compliance ////////////////////////////////

// Policy in policy.open-cluster-management.io/v1
type OCMPolicyV1_Policy = propolicyV1.Policy

// PlacementBinding in policy.open-cluster-management.io/v1
type OCMPolicyV1_PlacementBinding = propolicyV1.PlacementBinding

// ConfigurationPolicy in policy.open-cluster-management.io/v1
type OCMPolicyV1_ConfigurationPolicy = policyV1.ConfigurationPolicy
