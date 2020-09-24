module github.com/metal-stack/gardener-extension-provider-metal

go 1.15

require (
	github.com/ahmetb/gen-crd-api-reference-docs v0.1.5
	github.com/ajeddeloh/go-json v0.0.0-20170920214419-6a2fe990e083 // indirect
	github.com/ajeddeloh/yaml v0.0.0-00010101000000-000000000000 // indirect
	github.com/coreos/container-linux-config-transpiler v0.9.0
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf
	github.com/coreos/ignition v0.35.0 // indirect
	github.com/docker/spdystream v0.0.0-20181023171402-6480d4af844c // indirect
	github.com/emicklei/go-restful v2.12.0+incompatible // indirect
	github.com/gardener/etcd-druid v0.3.0
	github.com/gardener/gardener v1.8.2
	github.com/gardener/machine-controller-manager v0.34.0
	github.com/go-logr/logr v0.1.0
	github.com/gobuffalo/packr/v2 v2.8.0
	github.com/golang/mock v1.4.4
	github.com/google/go-cmp v0.5.2
	github.com/google/uuid v1.1.2
	github.com/imdario/mergo v0.3.8
	github.com/metal-stack/firewall-controller v0.1.8
	github.com/metal-stack/machine-controller-manager-provider-metal v0.1.3
	github.com/metal-stack/metal-go v0.9.1-0.20200924065251-2bea3718d7d7
	github.com/metal-stack/metal-lib v0.6.1
	github.com/onsi/ginkgo v1.14.1
	github.com/onsi/gomega v1.10.2
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5
	github.com/vincent-petithory/dataurl v0.0.0-20191104211930-d1553a71de50 // indirect
	go4.org v0.0.0-20180809161055-417644f6feb5 // indirect
	k8s.io/api v0.18.4
	k8s.io/apiextensions-apiserver v0.18.4
	k8s.io/apimachinery v0.18.4
	k8s.io/apiserver v0.17.11
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.io/code-generator v0.17.11
	k8s.io/component-base v0.17.11
	k8s.io/kubelet v0.17.11
	sigs.k8s.io/controller-runtime v0.6.0
)

replace (
	github.com/ajeddeloh/yaml => github.com/ajeddeloh/yaml v0.0.0-20170912190910-6b94386aeefd // indirect
	k8s.io/api => k8s.io/api v0.17.11
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.17.11
	k8s.io/apimachinery => k8s.io/apimachinery v0.17.11
	k8s.io/apiserver => k8s.io/apiserver v0.17.11
	k8s.io/client-go => k8s.io/client-go v0.17.11
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.17.11
	k8s.io/code-generator => k8s.io/code-generator v0.17.11
	k8s.io/component-base => k8s.io/component-base v0.17.11
	k8s.io/helm => k8s.io/helm v2.13.1+incompatible
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.17.11
	sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.4.0
)
