module github.com/Shuanglu/containerTcpdump

go 1.15

replace k8s.io/api => k8s.io/api v0.20.1

replace k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.20.1

replace k8s.io/apimachinery => k8s.io/apimachinery v0.21.0-alpha.0

replace k8s.io/apiserver => k8s.io/apiserver v0.20.1

replace k8s.io/cli-runtime => k8s.io/cli-runtime v0.20.1

replace k8s.io/client-go => k8s.io/client-go v0.20.1

replace k8s.io/cloud-provider => k8s.io/cloud-provider v0.20.1

replace k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.20.1

replace k8s.io/code-generator => k8s.io/code-generator v0.20.2-rc.0

replace k8s.io/component-base => k8s.io/component-base v0.20.1

replace k8s.io/component-helpers => k8s.io/component-helpers v0.20.1

replace k8s.io/controller-manager => k8s.io/controller-manager v0.20.1

replace k8s.io/cri-api => k8s.io/cri-api v0.20.2-rc.0

replace k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.20.1

replace k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.20.1

replace k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.20.1

replace k8s.io/kube-proxy => k8s.io/kube-proxy v0.20.1

replace k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.20.1

replace k8s.io/kubectl => k8s.io/kubectl v0.20.1

replace k8s.io/kubelet => k8s.io/kubelet v0.20.1

replace k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.20.1

replace k8s.io/metrics => k8s.io/metrics v0.20.1

replace k8s.io/mount-utils => k8s.io/mount-utils v0.20.2-rc.0

replace k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.20.1

replace k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.20.1

replace k8s.io/sample-controller => k8s.io/sample-controller v0.20.1

require (
	github.com/Microsoft/go-winio v0.4.15 // indirect
	github.com/containerd/containerd v1.4.1 // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v20.10.1+incompatible
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/google/gopacket v1.1.19
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/moby/term v0.0.0-20200312100748-672ec06f55cd // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/pflag v1.0.5
	github.com/thedevsaddam/gojsonq/v2 v2.5.2
	github.com/vishvananda/netns v0.0.0-20200728191858-db3c7e526aae
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	google.golang.org/grpc v1.34.0
	k8s.io/cri-api v0.0.0-00010101000000-000000000000
)
