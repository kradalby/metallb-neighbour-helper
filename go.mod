module github.com/kradalby/metallb-neighbour-helper

go 1.14

require (
	github.com/gobuffalo/envy v1.9.0
	github.com/googleapis/gnostic v0.3.0 // indirect
	github.com/imdario/mergo v0.3.7 // indirect
	github.com/kradalby/opnsense-go v0.0.0-20210123082920-7015c8e12160
	github.com/mikioh/ipaddr v0.0.0-20190404000644-d465c8ab6721
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/api v0.18.6
	k8s.io/apimachinery v0.18.6
	k8s.io/client-go v0.18.6
	k8s.io/kubernetes v1.18.6

)

replace (
	// github.com/kradalby/opnsense-go => ../opnsense-go

	// Things that kubernetes needs, probably pretty stupid...
	go.uber.org/atomic => go.uber.org/atomic v0.0.0-20181018215023-8dc6146f7569
	go.uber.org/multierr => go.uber.org/multierr v0.0.0-20180122172545-ddea229ff1df
	k8s.io/api => k8s.io/api v0.18.6
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.18.6
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.7-rc.0
	k8s.io/apiserver => k8s.io/apiserver v0.18.6
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.18.6
	k8s.io/client-go => k8s.io/client-go v0.18.6
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.18.6
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.18.6
	k8s.io/code-generator => k8s.io/code-generator v0.18.7-rc.0
	k8s.io/component-base => k8s.io/component-base v0.18.6
	k8s.io/cri-api => k8s.io/cri-api v0.18.7-rc.0
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.18.6
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.18.6
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.18.6

	k8s.io/kube-proxy => k8s.io/kube-proxy v0.18.6
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.18.6
	k8s.io/kubelet => k8s.io/kubelet v0.18.6
	k8s.io/kubernetes => k8s.io/kubernetes v1.18.6
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.18.6
	k8s.io/metrics => k8s.io/metrics v0.18.6
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.18.6
)

replace k8s.io/kubectl => k8s.io/kubectl v0.22.2

replace k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.18.6

replace k8s.io/sample-controller => k8s.io/sample-controller v0.18.6
