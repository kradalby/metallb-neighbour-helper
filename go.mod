module github.com/kradalby/metallb-neighbour-helper

go 1.23.0

require (
	github.com/gobuffalo/envy v1.9.0
	github.com/kradalby/opnsense-go v0.0.0-20210123082920-7015c8e12160
	github.com/mikioh/ipaddr v0.0.0-20190404000644-d465c8ab6721
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/api v0.18.6
	k8s.io/apimachinery v0.18.6
	k8s.io/client-go v0.18.6
	k8s.io/kubernetes v1.32.10
)

require (
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/googleapis/gnostic v0.3.0 // indirect
	github.com/imdario/mergo v0.3.7 // indirect
	github.com/joho/godotenv v1.3.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/crypto v0.28.0 // indirect
	golang.org/x/mod v0.21.0 // indirect
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/oauth2 v0.23.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/term v0.25.0 // indirect
	golang.org/x/text v0.19.0 // indirect
	golang.org/x/time v0.7.0 // indirect
	google.golang.org/appengine v1.5.0 // indirect
	google.golang.org/protobuf v1.35.1 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	k8s.io/klog v1.0.0 // indirect
	k8s.io/kube-openapi v0.0.0-20241105132330-32ad38e42d3f // indirect
	k8s.io/utils v0.0.0-20241104100929-3ea5e8cea738 // indirect
	sigs.k8s.io/structured-merge-diff/v3 v3.0.0 // indirect
	sigs.k8s.io/yaml v1.4.0 // indirect
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
	k8s.io/kubernetes => k8s.io/kubernetes v1.32.10
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.18.6
	k8s.io/metrics => k8s.io/metrics v0.18.6
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.18.6
)

replace k8s.io/kubectl => k8s.io/kubectl v0.18.6

replace k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.18.6

replace k8s.io/sample-controller => k8s.io/sample-controller v0.18.6
