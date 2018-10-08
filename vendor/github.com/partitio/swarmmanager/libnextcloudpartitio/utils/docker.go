package utils

import (
	"fmt"
	"github.com/docker/docker/api"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"os"
	"path/filepath"
	"time"
)

var StackLabel = "com.docker.stack.namespace"
var ServiceNameLabel = "com.docker.swarm.service.name"

func NetworkExist(c client.APIClient, name string) (bool, error) {
	filter := filters.NewArgs()
	filter.Add("name", name)
	networks, err := c.NetworkList(context.Background(), types.NetworkListOptions{
		Filters: filter,
	})
	if err != nil {
		return false, err
	}
	return len(networks) != 0, nil
}

func WaitOnService(ctx context.Context, cli client.APIClient, serviceID string, replicas uint64, timeOut float64) error {

	taskFilter := filters.NewArgs()
	taskFilter.Add("service", serviceID)
	taskFilter.Add("_up-to-date", "true")

	getUpToDateTasks := func() ([]swarm.Task, error) {
		return cli.TaskList(ctx, types.TaskListOptions{Filters: taskFilter})
	}

	begin := time.Now()

	for {
		tasks, err := getUpToDateTasks()
		if err != nil {
			return err
		}
		if serviceReady(tasks, int(replicas)) {
			return nil
		}
		duration := time.Now().Sub(begin)
		if duration.Seconds() > timeOut {
			return fmt.Errorf("operation timed out")
		}

		time.Sleep(200 * time.Millisecond)
	}

	return nil
}

func serviceReady(tasks []swarm.Task, replicas int) bool {
	var count = 0
	for _, task := range tasks {
		if task.Status.State == swarm.TaskStateRunning {
			count++
		}
	}
	return count == replicas
}

func RexRayMount(name string, target string, size int) mount.Mount {
	return mount.Mount{
		Type:   mount.TypeVolume,
		Source: name,
		Target: target,
		VolumeOptions: &mount.VolumeOptions{
			DriverConfig: &mount.Driver{
				Name:    "rexray/rbd",
				Options: map[string]string{"size": fmt.Sprintf("%v", size)},
			},
		},
	}
}

func FirstRunningTask(ctx context.Context, cli client.APIClient, serviceID string) (*swarm.Task, error) {
	f := filters.NewArgs()
	f.Add("service", serviceID)
	f.Add("_up-to-date", "true")
	tasks, err := cli.TaskList(ctx, types.TaskListOptions{Filters: f})
	if err != nil {
		return nil, err
	}
	for i, task := range tasks {
		if task.Status.State == swarm.TaskStateRunning {
			return &tasks[i], nil
		}
	}
	return nil, fmt.Errorf("could not find running task for service : %s", serviceID)
}

func NodeByID(ctx context.Context, cli client.APIClient, nodeID string) (*swarm.Node, error) {
	f := filters.NewArgs()
	f.Add("id", nodeID)
	nodes, err := cli.NodeList(ctx, types.NodeListOptions{Filters: f})
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, fmt.Errorf("no node found for id : %s", nodeID)
	}

	return &nodes[0], nil
}

func NodeClient(cli client.APIClient, node *swarm.Node) (*client.Client, error) {
	hostname := node.Description.Hostname
	current, err := cli.Info(context.Background())
	if err != nil {
		return nil, err
	}
	if hostname == current.Name {
		c, err := client.NewClientWithOpts(client.FromEnv)
		if err != nil {
			return nil, err
		}
		return c, nil
	}
	h := fmt.Sprintf("tcp://%s:2376", node.Status.Addr)
	home := os.Getenv("HOME")
	machineCertPath := filepath.Join(home, ".docker/machine/machines", hostname)
	if _, err := os.Stat(machineCertPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("%s : cannot find docker-machine folder", hostname)
	}
	caPath := filepath.Join(machineCertPath, "ca.pem")
	certPath := filepath.Join(machineCertPath, "cert.pem")
	keyPath := filepath.Join(machineCertPath, "key.pem")
	version := os.Getenv("DOCKER_API_VERSION")
	if version == "" {
		version = api.DefaultVersion
	}
	nodeClient, err := client.NewClientWithOpts(client.WithHost(h),
		client.WithTLSClientConfig(caPath, certPath, keyPath), client.WithVersion(version))
	if err != nil {
		return nil, err
	}
	return nodeClient, nil
}

func ScaleService(cli client.APIClient, ctx context.Context, serviceID string, scale uint64) error {
	service, _, err := cli.ServiceInspectWithRaw(ctx, serviceID, types.ServiceInspectOptions{})
	if err != nil {
		return err
	}
	serviceMode := &service.Spec.Mode
	serviceMode.Replicated.Replicas = &scale
	_, e := cli.ServiceUpdate(ctx, serviceID, service.Version, service.Spec, types.ServiceUpdateOptions{})
	if err != nil {
		return e
	}
	return nil
}

func UpdateServiceConstraint(cli client.APIClient, ctx context.Context, serviceID string, constraints []string) error {
	service, _, err := cli.ServiceInspectWithRaw(ctx, serviceID, types.ServiceInspectOptions{})
	if err != nil {
		return err
	}
	service.Spec.TaskTemplate.Placement.Constraints = constraints
	_, e := cli.ServiceUpdate(ctx, serviceID, service.Version, service.Spec, types.ServiceUpdateOptions{})
	if err != nil {
		return e
	}
	return nil
}

func RemoveVolume(cli client.APIClient, volume string, force bool, timeOut float64) error {
	begin := time.Now()
	for {
		d := time.Now().Sub(begin)
		err := cli.VolumeRemove(context.Background(), volume, force)
		if err == nil {
			return nil

		}
		if d.Seconds() > timeOut {
			return fmt.Errorf("removing volume %s timed out", volume)
		}
		time.Sleep(200 * time.Millisecond)
	}
	return nil
}
