package dockertcpdump

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	"github.com/Shuanglu/containerTcpdump/pkg/containerruntime"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	dockerClient "github.com/docker/docker/client"
	log "github.com/sirupsen/logrus"
)

type DockerRuntime struct {
	//Name string `json:"Name"`
}

/*
func (r *DockerRuntime) RetrievePID(tContainerID string) (int, error) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	//cmd := exec.Command("/bin/bash", "-c", "nsenter -t 1 -m docker inspect $(docker ps |grep '"+tContainerName+"'|awk '{print $1}')|grep 'Pid'")
	//cmd.SysProcAttr = &syscall.SysProcAttr{Cloneflags: syscall.CLONE_NEWPID |
	//	syscall.CLONE_NEWNS,
	//containerName := fmt.Sprintf("\"%s\"", tContainerName)
	//command := "nsenter --target 1 --mount -- /bin/bash -c \"docker ps |grep '" + tContainerName + "'\""
	//cmd := exec.Command("/bin/bash", "-c", command)
	//output, err := cmd.CombinedOutput()
	//log.Info(command)
	//if err != nil {
	//	log.Warn(fmt.Sprintf("Command execution returned error %q", err))
	//}
	//tContainerInfo := string(output)
	//tContainerID := strings.Split(tContainerInfo, " ")[0]
	//log.Info(tContainerID)
	//| awk "{print $1}")|grep "Pid"
	// + tContainerName + '" | awk "{print $1}")|grep "Pid"'
	//command := "nsenter --target 1 --mount -- /bin/bash -c \"docker inspect " + tContainerID + " --format='{{json .State.Pid}}'\""
	//log.Info(command)
	//cmd := exec.Command("/bin/bash", "-c", command)
	//output, err := cmd.CombinedOutput()
	//log.Info(string(output))
	//pattern := regexp.MustCompile("\"Pid\": ([0-9]+)")
	//tempContainerPid := pattern.FindStringSubmatch(string(output))[0]
	//log.Info(tempContainerPid)
	//pattern = regexp.MustCompile("([0-9]+)")
	//containerPid := pattern.FindStringSubmatch(tempContainerPid)[0]
	//log.Info(containerPid)
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	containerDetail, err := cli.ContainerInspect(context.TODO(), tContainerID)
	if err != nil {
		log.Warn(fmt.Sprintf("Failed to initialize docker client"))
		containerPid := 0
		return containerPid, err
	}
	containerPid := containerDetail.ContainerJSONBase.State.Pid

	return containerPid, nil
}


func (r *DockerRuntime) RetrieveNetns(UID string) (string, error) {
	command := "docker ps |grep 'k8s_POD_.*" + UID + "'|awk '{print $1}'"
	log.Info(command)
	cmd := exec.Command("/bin/bash", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Warn(fmt.Sprintf("Failed to get the container id of pod %q due to %q", UID, err))
		return "", err
	}
	containerId := string(output)
	containerId = strings.TrimSuffix(containerId, "\n")
	command = " docker inspect " + containerId + " |jq .[0].NetworkSettings.SandboxKey"
	log.Info(command)
	cmd = exec.Command("/bin/bash", "-c", command)
	output, err = cmd.CombinedOutput()
	netnsPath := string(output)
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to initialize docker client for the pod %q due to %q", UID, err))
		return "", err
	}
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to get containers for the pod %q due to %q", UID, err))
		return "", err
	}
	containerId := "0"
	//containerName :=
	//containerNamePattern := regexp.MustCompile(containerName)
	for _, container := range containers {
		match, err := regexp.Match("k8s_POD_.*"+UID, []byte(container.Names[0]))
		if err != nil {
			log.Fatal(fmt.Sprintf("Failed to match the container for the pod %q", UID))
		}
		if match {
			containerId = container.ID
		}
	}
	containerJson, err := cli.ContainerInspect(context.Background(), containerId)
	if err != nil {
		log.Warn(fmt.Sprintf("Failed to get the container detail of pod %q due to %q", UID, err))
		return "", err
	}
	netnsPath := containerJson.NetworkSettings.SandboxKey
	return netnsPath, nil
}
*/

type DockerRuntimeClient struct {
	runtimeClient *dockerClient.Client
}

func DockerRuntimeClientInit(addr string) (containerruntime.ContainerRuntime, error) {
	NewDockerClient, err := dockerClient.NewClientWithOpts(dockerClient.FromEnv, dockerClient.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	return &DockerRuntimeClient{
		runtimeClient: NewDockerClient,
	}, nil
}

func (rc *DockerRuntimeClient) GetPodSandboxId(UID string) (string, error) {
	var podSandboxID string
	containers, err := rc.runtimeClient.ContainerList(context.Background(), types.ContainerListOptions{
		Filters: filters.NewArgs(
			filters.KeyValuePair{
				Key:   "status",
				Value: "running",
			},
		),
	})
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to get containers for the pod %q due to %q", UID, err))
		return "", err
	}

	for _, container := range containers {
		match, err := regexp.Match("k8s_POD_.*"+UID, []byte(container.Names[0]))
		if err != nil {
			log.Fatal(fmt.Sprintf("Failed to match the container for the pod %q", UID))
		}
		if match {
			podSandboxID = container.ID
		}
	}
	if podSandboxID == "" {
		err := errors.New("No matching Sandbox ID found")
		return podSandboxID, err
	}
	return podSandboxID, nil
}

func (rc *DockerRuntimeClient) GetPodSandboxStatusInfo(id string) (interface{}, error) {
	podSandboxStatusInfo, err := rc.runtimeClient.ContainerInspect(context.Background(), id)
	if err != nil {
		log.Warn(fmt.Sprintf("Failed to get the container detail of pod %q due to %q", id, err))
		return "", err
	}
	return podSandboxStatusInfo, nil
}

func (rc *DockerRuntimeClient) GetPodSandboxNetworkNamespace(podSandboxStatusInfo interface{}) (string, error) {
	info := podSandboxStatusInfo.(types.ContainerJSON)
	netnsPath := info.NetworkSettings.SandboxKey
	return netnsPath, nil
}
