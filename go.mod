module github.com/MikeSpreitzer/ocm-pointers

go 1.19

require (
	open-cluster-management.io/api v0.11.0
	open-cluster-management.io/cert-policy-controller v0.86.0
	open-cluster-management.io/config-policy-controller v0.11.0
	open-cluster-management.io/governance-policy-propagator v0.11.0
	open-cluster-management.io/iam-policy-controller v0.86.0
	open-cluster-management.io/multicloud-operators-channel v0.11.0
	open-cluster-management.io/multicloud-operators-subscription v0.11.0
)

require (
	github.com/ghodss/yaml v1.0.1-0.20190212211648-25d852aebe32 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	k8s.io/api v0.27.1 // indirect
	k8s.io/apimachinery v0.27.1 // indirect
	k8s.io/klog/v2 v2.100.1 // indirect
	k8s.io/utils v0.0.0-20230505201702-9f6742963106 // indirect
	sigs.k8s.io/controller-runtime v0.14.6 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
)

replace (
	k8s.io/api => k8s.io/api v0.26.4
	k8s.io/apimachinery => k8s.io/apimachinery v0.26.4
	k8s.io/client-go => k8s.io/client-go v0.26.4
	open-cluster-management.io/cert-policy-controller => github.com/stolostron/cert-policy-controller v0.0.0-20230509165757-f0f4763c2a45
	open-cluster-management.io/iam-policy-controller => github.com/stolostron/iam-policy-controller v0.0.0-20230509200505-c33971547b94
)
