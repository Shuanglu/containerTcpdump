package containerruntime

//pb "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"

type ContainerRuntime interface {
	//RuntimeClientInit(addr string) (pb.RuntimeServiceClient, error)
	GetPodSandboxId(UID string) (string, error)
	GetPodSandboxStatusInfo(id string) (map[string]string, error)
	GetPodSandboxNetworkNamespace(podSandboxStatusInfo map[string]string) (string, error)
}
