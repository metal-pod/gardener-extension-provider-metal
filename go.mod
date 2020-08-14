module github.com/metal-stack/gardener-extension-provider-metal

go 1.14

require (
	contrib.go.opencensus.io/exporter/ocagent v0.4.12 // indirect
	git.apache.org/thrift.git v0.12.0 // indirect
	github.com/NYTimes/gziphandler v1.1.1 // indirect
	github.com/ajeddeloh/go-json v0.0.0-20170920214419-6a2fe990e083 // indirect
	github.com/ajeddeloh/yaml v0.0.0-00010101000000-000000000000 // indirect
	github.com/aliyun/alibaba-cloud-sdk-go v0.0.0-20190723075400-e63e3f9dd712 // indirect
	github.com/appscode/jsonpatch v0.0.0-20190108182946-7c0e3b262f30 // indirect
	github.com/coreos/container-linux-config-transpiler v0.9.0
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf
	github.com/coreos/ignition v0.35.0 // indirect
	github.com/docker/spdystream v0.0.0-20181023171402-6480d4af844c // indirect
	github.com/emicklei/go-restful v2.12.0+incompatible // indirect
	github.com/emicklei/go-restful-openapi v1.4.1 // indirect
	github.com/gardener/etcd-backup-restore v0.10.0 // indirect
	github.com/gardener/etcd-druid v0.3.0
	github.com/gardener/gardener v1.6.6
	github.com/gardener/machine-controller-manager v0.31.0
	github.com/go-ini/ini v1.57.0 // indirect
	github.com/go-logr/logr v0.1.0
	github.com/gobuffalo/packr/v2 v2.7.1
	github.com/golang/mock v1.4.3
	github.com/google/go-cmp v0.4.1
	github.com/google/uuid v1.1.1
	github.com/googleapis/gax-go v2.0.0+incompatible // indirect
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79 // indirect
	github.com/jetstack/cert-manager v0.6.2 // indirect
	github.com/metal-stack/firewall-controller v0.1.8
	github.com/metal-stack/metal-go v0.8.3
	github.com/metal-stack/metal-lib v0.5.0
	github.com/onsi/ginkgo v1.14.0
	github.com/onsi/gomega v1.10.1
	github.com/pkg/errors v0.9.1
	github.com/pkg/profile v1.2.1 // indirect
	github.com/prometheus/prometheus v2.5.0+incompatible // indirect
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.6.1 // indirect
	github.com/ugorji/go v1.1.7 // indirect
	github.com/vincent-petithory/dataurl v0.0.0-20191104211930-d1553a71de50 // indirect
	golang.org/x/build v0.0.0-20200701200223-dd5ac55b812e // indirect
	k8s.io/api v0.18.4
	k8s.io/apiextensions-apiserver v0.18.4
	k8s.io/apimachinery v0.18.4
	k8s.io/apiserver v0.17.11
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.io/component-base v0.18.0
	k8s.io/kubelet v0.17.11
	sigs.k8s.io/controller-runtime v0.6.0
)

replace (
	github.com/ajeddeloh/yaml => github.com/ajeddeloh/yaml v0.0.0-20170912190910-6b94386aeefd // indirect
	github.com/gardener/machine-controller-manager => github.com/metal-stack/machine-controller-manager v0.26.1-0.20200306160031-6be3c8ee6d66
	k8s.io/api => k8s.io/api v0.17.11 // 1.16.8
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.17.11 // 1.16.8
	k8s.io/apimachinery => k8s.io/apimachinery v0.17.11 // 1.16.8
	k8s.io/apiserver => k8s.io/apiserver v0.17.11 // 1.16.8
	k8s.io/client-go => k8s.io/client-go v0.17.11 // 1.16.8
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.17.11 // 1.16.8
	k8s.io/code-generator => k8s.io/code-generator v0.17.11 // 1.16.8
	k8s.io/component-base => k8s.io/component-base v0.17.11 // 1.16.8
	k8s.io/helm => k8s.io/helm v2.13.1+incompatible
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.17.11 // 1.16.8
	sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.4.0
)
