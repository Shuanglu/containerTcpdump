package containerdtcpdump

/*
import (
	"fmt"
	"os/exec"
	"strings"

	//"github.com/kubernetes-sigs/cri-tools/"

	//"github.com/kubernetes-sigs/cri-tools/cmd/crictl/sandbox/main"

	//"github.com/golang/protobuf/jsonpb"

	log "github.com/sirupsen/logrus"
)


type ContainerdRuntime struct {
	//Name string `json:"Name"`
}


func getRuntimeClientConnection() (*grpc.ClientConn, error) {
	logrus.Debug("get runtime connection")
	// If no EP set then use the default endpoint types
	RuntimeEndpoint := "unix:///run/containerd/containerd.sock"
	return getConnection(RuntimeEndpoint)
}

func getConnection(endPoint string) (*grpc.ClientConn, error) {
	var conn *grpc.ClientConn
	Timeout := 5 * time.Second
	log.Info(fmt.Sprintf("connect using endpoint '%s' with '%s' timeout", endPoint, Timeout))
	addr, dialer, err := util.GetAddressAndDialer(endPoint)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to initialize runtime dailer due to %q", err))
	}
	conn, err = grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(Timeout), grpc.WithContextDialer(dialer))
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to connect to the endpoint due to %q", err))
	} else {
		log.Info(fmt.Sprintf("connected successfully using endpoint: %q", endPoint))
	}
	return conn, nil
}

func protobufObjectToJSON(obj proto.Message) (string, error) {
	jsonpbMarshaler := jsonpb.Marshaler{EmitDefaults: true, Indent: "  "}
	marshaledJSON, err := jsonpbMarshaler.MarshalToString(obj)
	if err != nil {
		return "", err
	}
	return marshaledJSON, nil
}

func getRuntimeClient() (pb.RuntimeServiceClient, *grpc.ClientConn, error) {
	// Set up a connection to the server.
	conn, err := getRuntimeClientConnection()
	if err != nil {
		return nil, nil, errors.Wrap(err, "connect")
	}
	runtimeClient := pb.NewRuntimeServiceClient(conn)
	return runtimeClient, conn, nil
}


func (r *ContainerdRuntime) RetrieveNetns(UID string) (string, error) {
	/*runtimeClient, _, err := getRuntimeClient()
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to initialzie the runtime client due to %q", err))
	}
	filter := &pb.PodSandboxFilter{}
	label := make(map[string]string)
	label["io.kubernetes.pod.uid"] = UID
	filter.LabelSelector = label
	st := &pb.PodSandboxStateValue{}
	st.State = pb.PodSandboxState_SANDBOX_READY
	filter.State = st
	request := &pb.ListPodSandboxRequest{
		Filter: filter,
	}
	r, err := runtimeClient.ListPodSandbox(context.Background(), request)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to list the sandbox of pod %q due to %q", UID, err))
		return "", err
	}
	podId := r.Items[0].Id
	resp, err := runtimeClient.PodSandboxStatus(context.Background(), &pb.PodSandboxStatusRequest{
		PodSandboxId: podId,
		Verbose:      true,
	})
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to get the sandbox detail of pod %q due to %q", UID, err))
		return "", err
	}
	keys := []string{}
	for k := range resp.Info {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		var res interface{}
		if k == "runtimeSpec" {
			json.Unmarshal([]byte(resp.Info[k]), &res)
			v := resp.Info[k]
			if k == "linux" {
				json.Unmarshal([]byte(v[k]), &res)
				if k == "namespaces" {
					json.Unmarshal([]byte(v[k]), &res)
					resp = v[k]

				}
			}
		}
	}

	for _, namespace := range resp.Status.Linux.Namespaces.Options {

	}
	command := "crictl -r unix:///run/containerd/containerd.sock pods -o json --label 'io.kubernetes.pod.uid=" + UID + "' | jq .items[0].id"
	log.Info(command)
	cmd := exec.Command("/bin/bash", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Warn(fmt.Sprintf("Failed to get the pod id of pod %q due to %q", UID, err))
		return "", err
	}
	podID := string(output)
	podID = strings.TrimSuffix(podID, "\n")
	podID = strings.TrimSuffix(podID, "\"")
	podID = strings.TrimPrefix(podID, "\"")
	command = "crictl -r unix:///run/containerd/containerd.sock inspectp \"" + podID + "\" | jq '.info.runtimeSpec.linux.namespaces |.[] | select(.type|contains(\"network\"))| .path'"
	log.Info(command)
	cmd = exec.Command("/bin/bash", "-c", command)
	output, err = cmd.CombinedOutput()
	netnsPath := string(output)
	if err != nil {
		log.Warn(fmt.Sprintf("Failed to get the network namespace of pod %q due to %q", UID, err))
		return "", err
	}
	return netnsPath, nil
}

/*
func (r *ContainerdRuntime) RetrievePID(UID string) (string, error) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	//command := "nsenter --target 1 --mount -- /bin/bash -c 'crictl inspect " + tContainerID + "|jq .info.pid'"
	log.Info(command)
	cmd := exec.Command("/bin/bash", "-c", command)
	output, err := cmd.CombinedOutput()
	//log.Info(string(output))
	//pattern := regexp.MustCompile("\"pid\": ([0-9]+)")
	//tempContainerPid := pattern.FindStringSubmatch(string(output))[0]
	//log.Info(tempContainerPid)
	//pattern = regexp.MustCompile("([0-9]+)")
	//containerPid := pattern.FindStringSubmatch(tempContainerPid)[0]
	//log.Info(containerPid)
	containerPid := string(output)
	if err != nil {
		log.Warn(fmt.Sprintf("Command execution returned error %q", err))
	}
	return containerPid, err
}
*/

import (
	"errors"
	"fmt"
	"time"

	"github.com/Shuanglu/containerTcpdump/pkg/containerruntime"
	//"github.com/containerd/containerd/log"
	log "github.com/sirupsen/logrus"
	gojsonq "github.com/thedevsaddam/gojsonq/v2"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

const (
	//dockerep       = "unix:///var/run/dockershim.sock"
	containerdep = "unix:///run/containerd/containerd.sock"
	//crioep         = "unix:///run/crio/crio.sock"
	defaultTimeout = 2 * time.Second
)

type ContainerdRuntimeClient struct {
	runtimeClient pb.RuntimeServiceClient
}

func ContainerdRuntimeClientInit(addr string) (containerruntime.ContainerRuntime, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(defaultTimeout))
	if err != nil {
		return nil, err
	}
	return &ContainerdRuntimeClient{
		runtimeClient: pb.NewRuntimeServiceClient(conn),
	}, nil
}

func (rc *ContainerdRuntimeClient) GetPodSandboxId(UID string) (string, error) {
	var podSandboxID string
	listPodSandboxRequest := pb.ListPodSandboxRequest{
		Filter: &pb.PodSandboxFilter{
			State: &pb.PodSandboxStateValue{
				State: pb.PodSandboxState_SANDBOX_READY,
			},
		},
	}
	podSandboxList, err := rc.runtimeClient.ListPodSandbox(context.Background(), &listPodSandboxRequest)
	if err != nil {
		//hostname, _ := os.Hostname()
		return "", err
	}
	for _, podSandbox := range podSandboxList.GetItems() {
		log.Info(fmt.Sprintf("ID: %s", podSandbox))
		if podSandbox.Metadata.Uid == UID {
			podSandboxID = podSandbox.Id
			log.Info(fmt.Sprintf("Sandbox: %s", podSandbox))
			break
		}
	}
	if podSandboxID == "" {
		err := errors.New("No matching Sandbox ID found")
		return podSandboxID, err
	}
	return podSandboxID, nil
}

func (rc *ContainerdRuntimeClient) GetPodSandboxStatusInfo(id string) (interface{}, error) {
	podSandboxStatusRequest := pb.PodSandboxStatusRequest{
		PodSandboxId: id,
		Verbose:      true,
	}
	podSandboxStatusResponse, err := rc.runtimeClient.PodSandboxStatus(context.Background(), &podSandboxStatusRequest)
	if err != nil {
		return nil, err
	}
	return podSandboxStatusResponse.GetInfo(), nil
}

func (rc *ContainerdRuntimeClient) GetPodSandboxNetworkNamespace(podSandboxStatusInfo interface{}) (string, error) {
	var netNamespacePath string
	info := podSandboxStatusInfo.(map[string]string)
	namespaces := gojsonq.New().FromString(info["info"]).Find("runtimeSpec.linux.namespaces")
	for _, namespace := range namespaces.([]interface{}) {
		namespaceMap := namespace.(map[string]interface{})
		if namespaceMap["type"].(string) == "network" {
			netNamespacePath = namespaceMap["path"].(string)
			break
		}
	}
	return netNamespacePath, nil
}
