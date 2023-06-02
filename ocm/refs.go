package ocm

import (
	addonV1a1 "open-cluster-management.io/api/addon/v1alpha1"
	clusterV1 "open-cluster-management.io/api/cluster/v1"
	clusterV1a1 "open-cluster-management.io/api/cluster/v1alpha1"
	clusterV1b1 "open-cluster-management.io/api/cluster/v1beta1"
	clusterV1b2 "open-cluster-management.io/api/cluster/v1beta2"
	operatorV1 "open-cluster-management.io/api/operator/v1"
	workV1 "open-cluster-management.io/api/work/v1"
	policyV1 "open-cluster-management.io/config-policy-controller/api/v1"
	propolicyV1 "open-cluster-management.io/governance-policy-propagator/api/v1"
	propolicyV1b1 "open-cluster-management.io/governance-policy-propagator/api/v1beta1"
	chanappsV1 "open-cluster-management.io/multicloud-operators-channel/pkg/apis/apps/v1"
	ansibleV1a1 "open-cluster-management.io/multicloud-operators-subscription/pkg/apis/apps/ansible/v1alpha1"
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

// ManagedClusterSet in cluster.open-cluster-management.io/v1beta2 but NOT v1
// ManagedClusterSet defines a group of ManagedClusters that user's workload can run on.
// type OCMClusterV1a1_ManagedClusterSet = clusterV1a1.ManagedClusterSet
type OCMClusterV1b2_ManagedClusterSet = clusterV1b2.ManagedClusterSet
type OCMClusterV1b1_ManagedClusterSet = clusterV1b1.ManagedClusterSet

// ManagedClusterSetBinding in cluster.open-cluster-management.io/v1beta2
// ManagedClusterSetBinding appears in a namespace and conveys limited privileges to use
// from within that workspace.
type OCMClusterV1b2_ManagedClusterSetBinding = clusterV1b2.ManagedClusterSetBinding
type OCMClusterV1b1_ManagedClusterSetBinding = clusterV1b1.ManagedClusterSetBinding

// Placement in cluster.open-cluster-management.io/v1beta1 and NOT later
// Placement defines a rule to select a set of ManagedClusters from the ManagedClusterSets bound
// to the placement namespace.
type OCMClusterV1b1_Placement = clusterV1b1.Placement

// AddOnPlacementScore in cluster.open-cluster-management.io/v1alpha1
// AddOnPlacementScore represents a bundle of scores of one managed cluster, which could be used by placement.
type OCMClusterV1a1_AddOnPlacementScore = clusterV1a1.AddOnPlacementScore

// PlacementDecision in cluster.open-cluster-management.io/v1beta1 and NOT later
// PlacementDecision indicates a decision from a placement
// type OCMClusterV1a1_PlacementDecision = clusterV1a1.PlacementDecision removed in v0.9
type OCMClusterV1b1_PlacementDecision = clusterV1b1.PlacementDecision

// ClusterClaim in cluster.open-cluster-management.io/v1alpha1
// ClusterClaim represents cluster information that a managed cluster claims
type OCMClusterV1a1_ClusterClaim = clusterV1a1.ClusterClaim

// ClusterManagementAddOn in addon.open-cluster-management.io/v1alpha1
// ClusterManagementAddOn represents the registration of an add-on to the cluster manager.
type OCMAddonV1a1_ClusterManagementAddOn = addonV1a1.ClusterManagementAddOn

// AddOnDeploymentConfig in addon.open-cluster-management.io
type OCMAddonV1a1_AddOnDeploymentConfig = addonV1a1.AddOnDeploymentConfig

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

// Deployable was in apps.open-cluster-management.io/v1 up through v0.8.0
// Gone from v0.9.0 onward, users expected to migrate to ManifestWork.
// type OCMAppsV1_Deployable = deployableV1.Deployable

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
// A Manifest work exists in the hub's mailbox namespace for the managed cluster.
// There is nothing that binds a destination-agnostic ManifestWork with a Placement.
type OCMWorkV1_ManifestWork = workV1.ManifestWork

// AnsibleJob in tower.ansible.com/v1alpha1
type AnsibleTowerV1a1_AnsibleJob = ansibleV1a1.AnsibleJob

//////////////////////////////// Governance, Risk, Compliance ////////////////////////////////

// Policy in policy.open-cluster-management.io/v1
// A Policy contains a slice of PolicyTemplate.
type OCMPolicyV1_Policy = propolicyV1.Policy

// PolicySet in policy.open-cluster-management.io/v1beta1
// A PolicySet refers to some policies.
type OCMPolicyV1b1_PolicySet = propolicyV1b1.PolicySet

// PolicyAutomation in policy.open-cluster-management.io/v1beta1
type OCMPolicyV1b1_PolicyAutomation = propolicyV1b1.PolicyAutomation

// PlacementBinding in policy.open-cluster-management.io/v1
// Binds a slice of (Policy or PolicySet) with a Placement or PlacementRule
type OCMPolicyV1_PlacementBinding = propolicyV1.PlacementBinding

// ConfigurationPolicy in policy.open-cluster-management.io/v1
type OCMPolicyV1_ConfigurationPolicy = policyV1.ConfigurationPolicy
