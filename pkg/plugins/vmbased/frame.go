// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package vmbased

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"openpitrix.io/openpitrix/pkg/client"
	clusterclient "openpitrix.io/openpitrix/pkg/client/cluster"
	runtimeclient "openpitrix.io/openpitrix/pkg/client/runtime"
	"openpitrix.io/openpitrix/pkg/constants"
	"openpitrix.io/openpitrix/pkg/devkit/app"
	"openpitrix.io/openpitrix/pkg/logger"
	"openpitrix.io/openpitrix/pkg/models"
	"openpitrix.io/openpitrix/pkg/pb"
	"openpitrix.io/openpitrix/pkg/pb/types"
	"openpitrix.io/openpitrix/pkg/pi"
	"openpitrix.io/openpitrix/pkg/util/jsonutil"
	"openpitrix.io/openpitrix/pkg/util/pbutil"
	"openpitrix.io/openpitrix/pkg/util/sshutil"
)

type Frame struct {
	Job            *models.Job
	ClusterWrapper *models.ClusterWrapper
	Runtime        *runtimeclient.Runtime
}

func (f *Frame) startConfdServiceLayer(nodeIds []string, failureAllowed bool) *models.TaskLayer {
	taskLayer := new(models.TaskLayer)
	for _, nodeId := range nodeIds {
		clusterNode := f.ClusterWrapper.ClusterNodes[nodeId]
		meta := &models.Meta{
			FrontgateId: f.ClusterWrapper.Cluster.FrontgateId,
			Timeout:     TimeoutStartConfd,
			NodeId:      clusterNode.NodeId,
			DroneIp:     clusterNode.PrivateIp,
		}
		directive, err := meta.ToString()
		if err != nil {
			return nil
		}
		startConfdTask := &models.Task{
			JobId:          f.Job.JobId,
			Owner:          f.Job.Owner,
			TaskAction:     ActionStartConfd,
			Target:         constants.TargetPilot,
			NodeId:         nodeId,
			Directive:      string(directive),
			FailureAllowed: failureAllowed,
		}
		taskLayer.Tasks = append(taskLayer.Tasks, startConfdTask)
	}
	if len(taskLayer.Tasks) > 0 {
		return taskLayer
	} else {
		return nil
	}
}

func (f *Frame) stopConfdServiceLayer(nodeIds []string, failtureAllowed bool) *models.TaskLayer {
	taskLayer := new(models.TaskLayer)
	for _, nodeId := range nodeIds {
		clusterNode := f.ClusterWrapper.ClusterNodes[nodeId]
		meta := &models.Meta{
			FrontgateId: f.ClusterWrapper.Cluster.FrontgateId,
			Timeout:     TimeoutStopConfd,
			NodeId:      clusterNode.NodeId,
			DroneIp:     clusterNode.PrivateIp,
		}
		directive, err := meta.ToString()
		if err != nil {
			return nil
		}
		stopConfdTask := &models.Task{
			JobId:          f.Job.JobId,
			Owner:          f.Job.Owner,
			TaskAction:     ActionStopConfd,
			Target:         constants.TargetPilot,
			NodeId:         nodeId,
			Directive:      string(directive),
			FailureAllowed: failtureAllowed,
		}
		taskLayer.Tasks = append(taskLayer.Tasks, stopConfdTask)
	}
	if len(taskLayer.Tasks) > 0 {
		return taskLayer
	} else {
		return nil
	}
}

// Put the nodes into two groups
func (f *Frame) getPreAndPostStartGroupNodes(nodeIds []string) ([]string, []string) {
	var preGroupNodes, postGroupNodes []string
	for _, nodeId := range nodeIds {
		role := f.ClusterWrapper.ClusterNodes[nodeId].Role
		serviceStr := f.ClusterWrapper.ClusterCommons[role].InitService
		if serviceStr != "" {
			service := app.Service{}
			err := json.Unmarshal([]byte(serviceStr), &service)
			if err != nil {
				logger.Error("Unmarshal cluster [%s] init service failed: %+v",
					f.ClusterWrapper.Cluster.ClusterId, err)
				return nil, nil
			}
			postStartService := false
			if service.PostStartService != nil {
				postStartService = *service.PostStartService
			}
			if postStartService {
				postGroupNodes = append(postGroupNodes, nodeId)
			} else {
				preGroupNodes = append(preGroupNodes, nodeId)
			}
		}
	}
	return preGroupNodes, postGroupNodes
}

// Put the nodes into two groups
func (f *Frame) getPreAndPostStopGroupNodes(nodeIds []string) ([]string, []string) {
	var preGroupNodes, postGroupNodes []string
	for _, nodeId := range nodeIds {
		role := f.ClusterWrapper.ClusterNodes[nodeId].Role
		serviceStr := f.ClusterWrapper.ClusterCommons[role].DestroyService
		if serviceStr != "" {
			service := app.Service{}
			err := json.Unmarshal([]byte(serviceStr), &service)
			if err != nil {
				logger.Error("Unmarshal cluster [%s] init service failed: %+v",
					f.ClusterWrapper.Cluster.ClusterId, err)
				return nil, nil
			}
			postStopService := false
			if service.PostStopService != nil {
				postStopService = *service.PostStopService
			}
			if postStopService {
				postGroupNodes = append(postGroupNodes, nodeId)
			} else {
				preGroupNodes = append(preGroupNodes, nodeId)
			}
		}
	}
	return preGroupNodes, postGroupNodes
}

func (f *Frame) deregisterCmdLayer(nodeIds []string, failureAllowed bool) *models.TaskLayer {
	taskLayer := new(models.TaskLayer)
	for _, nodeId := range nodeIds {
		ip := f.ClusterWrapper.ClusterNodes[nodeId].PrivateIp
		cnodes := GetCmdCnodes(
			ip,
			nil,
		)
		meta := &models.Meta{
			FrontgateId: f.ClusterWrapper.Cluster.FrontgateId,
			NodeId:      nodeId,
			DroneIp:     ip,
			Timeout:     TimeoutDeregister,
			Cnodes:      jsonutil.ToString(cnodes),
		}
		directive, err := meta.ToString()
		if err != nil {
			return nil
		}
		deregisterCmdTask := &models.Task{
			JobId:          f.Job.JobId,
			Owner:          f.Job.Owner,
			TaskAction:     ActionDeregesterCmd,
			Target:         constants.TargetPilot,
			NodeId:         nodeId,
			Directive:      string(directive),
			FailureAllowed: failureAllowed,
		}
		taskLayer.Tasks = append(taskLayer.Tasks, deregisterCmdTask)
	}

	if len(taskLayer.Tasks) > 0 {
		return taskLayer
	} else {
		return nil
	}
}

func (f *Frame) registerCmdLayer(nodeIds []string, serviceName string, failureAllowed bool) *models.TaskLayer {
	taskLayer := new(models.TaskLayer)
	for _, nodeId := range nodeIds {
		role := f.ClusterWrapper.ClusterNodes[nodeId].Role
		serviceStr := f.ClusterWrapper.GetCommonAttribute(role, serviceName)
		if serviceStr != nil {
			service := app.Service{}
			err := json.Unmarshal([]byte(serviceStr.(string)), &service)
			if err != nil {
				logger.Error("Unmarshal cluster [%s] service [%s] failed: %+v",
					f.ClusterWrapper.Cluster.ClusterId, serviceName, err)
				return nil
			}
			timeout := constants.DefaultServiceTimeout
			if service.Timeout != nil {
				timeout = int(*service.Timeout)
			}
			if service.Cmd == "" {
				continue
			}
			ip := f.ClusterWrapper.ClusterNodes[nodeId].PrivateIp
			cnodes := &models.Cmd{
				Cmd:     service.Cmd,
				Timeout: timeout,
			}
			meta := &models.Meta{
				FrontgateId: f.ClusterWrapper.Cluster.FrontgateId,
				NodeId:      nodeId,
				DroneIp:     ip,
				Timeout:     timeout,
				Cnodes:      jsonutil.ToString(cnodes),
			}
			directive, err := meta.ToString()
			if err != nil {
				return nil
			}
			registerCmdTask := &models.Task{
				JobId:          f.Job.JobId,
				Owner:          f.Job.Owner,
				TaskAction:     ActionRegisterCmd,
				Target:         constants.TargetPilot,
				NodeId:         nodeId,
				Directive:      string(directive),
				FailureAllowed: failureAllowed,
			}
			taskLayer.Tasks = append(taskLayer.Tasks, registerCmdTask)
		}
	}

	if len(taskLayer.Tasks) > 0 {
		return taskLayer
	} else {
		return nil
	}
}

func (f *Frame) constructServiceTasks(serviceName, cmdName string, nodeIds []string,
	serviceParams map[string]interface{}, failureAllowed bool) *models.TaskLayer {
	headTaskLayer := new(models.TaskLayer)
	if len(nodeIds) == 0 {
		return nil
	}

	roleNodeIds := make(map[string][]string)
	nodeIdRole := make(map[string]string)
	for _, nodeId := range nodeIds {
		clusterNode, exist := f.ClusterWrapper.ClusterNodes[nodeId]
		if !exist {
			logger.Error("ClusterConf [%s] node [%s] not exist", f.ClusterWrapper.Cluster.ClusterId, nodeId)
			continue
		}
		role := clusterNode.Role
		service := f.ClusterWrapper.GetCommonAttribute(role, serviceName)
		if service == nil {
			continue
		}

		agentInstalled := f.ClusterWrapper.GetCommonAttribute(role, "AgentInstalled")
		if agentInstalled == nil {
			continue
		}

		if service.(string) == "" || !agentInstalled.(bool) {
			continue
		}
		roleNodeIds[role] = append(roleNodeIds[role], nodeId)
		nodeIdRole[nodeId] = role
	}

	filterNodes := make(map[string]string)
	roleService := make(map[string]app.Service)
	for role, nodes := range roleNodeIds {
		serviceStr := f.ClusterWrapper.GetCommonAttribute(role, serviceName)
		if serviceStr == nil {
			return nil
		}
		service := app.Service{}
		err := json.Unmarshal([]byte(serviceStr.(string)), &service)
		if err != nil {
			logger.Error("Unmarshal cluster [%s] service [%s] failed: %+v",
				f.ClusterWrapper.Cluster.ClusterId, serviceName, err)
			return nil
		}
		roleService[role] = service
		execNodeNums := len(nodes)
		if service.NodesToExecuteOn != nil {
			execNodeNums = int(*service.NodesToExecuteOn)
		}
		if execNodeNums < len(nodes) && strings.HasSuffix(role, constants.ReplicaRoleSuffix) {
			// when the given nodes_to_execute_on is less than the length of the nodes, then ignore the replicas
			for _, nodeId := range nodes {
				filterNodes[nodeId] = ""
			}
			continue
		}
		num := execNodeNums
		for num < len(nodes) {
			filterNodes[nodes[num-1]] = ""
			num++
		}
	}

	orderNodeIds := make(map[int][]string)
	for nodeId, role := range nodeIdRole {
		_, exist := filterNodes[nodeId]
		if exist {
			continue
		}
		service := roleService[role]
		order := 0
		if service.Order != nil {
			order = int(*service.Order)
		}
		orderNodeIds[order] = append(orderNodeIds[order], nodeId)
	}

	var orders []int
	for order := range orderNodeIds {
		orders = append(orders, order)
	}

	sort.Ints(orders)

	for _, order := range orders {
		nodeIds := orderNodeIds[order]
		taskLayer := f.registerCmdLayer(nodeIds, serviceName, failureAllowed)
		headTaskLayer.Leaf().Child = taskLayer
	}
	return headTaskLayer.Child
}

func (f *Frame) initServiceLayer(nodeIds []string, failureAllowed bool) *models.TaskLayer {
	return f.constructServiceTasks("InitService", constants.ServiceCmdName, nodeIds, nil, failureAllowed)
}

func (f *Frame) startServiceLayer(nodeIds []string, failureAllowed bool) *models.TaskLayer {
	return f.constructServiceTasks("StartService", constants.ServiceCmdName, nodeIds, nil, failureAllowed)
}

func (f *Frame) stopServiceLayer(nodeIds []string, failureAllowed bool) *models.TaskLayer {
	return f.constructServiceTasks("StopService", constants.ServiceCmdName, nodeIds, nil, failureAllowed)
}

func (f *Frame) scaleOutPreCheckServiceLayer(nodeIds []string, failureAllowed bool) *models.TaskLayer {
	return f.constructServiceTasks("ScaleOutService", constants.ServiceCmdName, nodeIds, nil, failureAllowed)
}

func (f *Frame) scaleInPreCheckServiceLayer(nodeIds []string, failureAllowed bool) *models.TaskLayer {
	return f.constructServiceTasks("ScaleInService", constants.ServiceCmdName, nodeIds, nil, failureAllowed)
}

func (f *Frame) scaleOutServiceLayer(nodeIds []string, failureAllowed bool) *models.TaskLayer {
	return f.constructServiceTasks("ScaleOutService", constants.ServicePreCheckName, nodeIds, nil, failureAllowed)
}

func (f *Frame) scaleInServiceLayer(nodeIds []string, failureAllowed bool) *models.TaskLayer {
	return f.constructServiceTasks("ScaleInService", constants.ServicePreCheckName, nodeIds, nil, failureAllowed)
}

func (f *Frame) destroyServiceLayer(nodeIds []string, failureAllowed bool) *models.TaskLayer {
	return f.constructServiceTasks("DestroyService", constants.ServiceCmdName, nodeIds, nil, failureAllowed)
}

func (f *Frame) initAndStartServiceLayer(nodeIds []string, failureAllowed bool) *models.TaskLayer {
	headTaskLayer := new(models.TaskLayer)

	preStartNodes, postStartNodes := f.getPreAndPostStartGroupNodes(nodeIds)

	// Init service before start service
	headTaskLayer.Leaf().Child = f.initServiceLayer(preStartNodes, failureAllowed)

	// TODO: custom metadata
	headTaskLayer.Leaf().Child = f.startServiceLayer(nodeIds, failureAllowed)

	// Init service after start service
	headTaskLayer.Leaf().Child = f.initServiceLayer(postStartNodes, failureAllowed)

	return headTaskLayer.Child
}

func (f *Frame) destroyAndStopServiceLayer(nodeIds []string, extraLayer *models.TaskLayer, failureAllowed bool) *models.TaskLayer {
	headTaskLayer := new(models.TaskLayer)

	preStopNodes, postStopNodes := f.getPreAndPostStopGroupNodes(nodeIds)

	// Destroy service before stop service
	headTaskLayer.Leaf().Child = f.destroyServiceLayer(preStopNodes, failureAllowed)

	if extraLayer != nil {
		headTaskLayer.Leaf().Child = extraLayer
	}

	headTaskLayer.Leaf().Child = f.stopServiceLayer(nodeIds, failureAllowed)

	// Destroy service after stop service
	headTaskLayer.Leaf().Child = f.destroyServiceLayer(postStopNodes, failureAllowed)

	return headTaskLayer.Child
}

func (f *Frame) createVolumesLayer(nodeIds []string, failureAllowed bool) *models.TaskLayer {
	taskLayer := new(models.TaskLayer)
	for _, nodeId := range nodeIds {
		clusterNode := f.ClusterWrapper.ClusterNodes[nodeId]
		role := clusterNode.Role
		if strings.HasSuffix(role, constants.ReplicaRoleSuffix) {
			role = string([]byte(role)[:len(role)-len(constants.ReplicaRoleSuffix)])
		}
		clusterRole, exist := f.ClusterWrapper.ClusterRoles[role]
		if !exist {
			logger.Error("No such role [%s] in cluster role [%s]. ",
				role, f.ClusterWrapper.Cluster.ClusterId)
			return nil
		}

		size := clusterRole.StorageSize
		if size > 0 {
			mountPoints := strings.Split(clusterRole.MountPoint, ",")
			eachSize := int(size) / len(mountPoints)

			volume := &models.Volume{
				Name:      clusterNode.ClusterId + "_" + nodeId,
				Size:      eachSize,
				Zone:      f.Runtime.Zone,
				RuntimeId: f.Runtime.RuntimeId,
			}
			volumeTaskDirective, err := volume.ToString()
			if err != nil {
				return nil
			}

			createVolumesTask := &models.Task{
				JobId:          f.Job.JobId,
				Owner:          f.Job.Owner,
				TaskAction:     ActionCreateVolumes,
				Target:         f.Runtime.Provider,
				NodeId:         nodeId,
				Directive:      volumeTaskDirective,
				FailureAllowed: failureAllowed,
			}
			for range mountPoints {
				taskLayer.Tasks = append(taskLayer.Tasks, createVolumesTask)
			}
		}
	}
	return taskLayer
}

func (f *Frame) detachVolumesLayer(nodeIds []string, failureAllowed bool) *models.TaskLayer {
	taskLayer := new(models.TaskLayer)
	for _, nodeId := range nodeIds {
		clusterNode := f.ClusterWrapper.ClusterNodes[nodeId]
		volume := &models.Volume{
			Name:       clusterNode.ClusterId + "_" + nodeId,
			Zone:       f.Runtime.Zone,
			RuntimeId:  f.Runtime.RuntimeId,
			VolumeId:   clusterNode.VolumeId,
			InstanceId: clusterNode.InstanceId,
		}
		directive, err := volume.ToString()
		if err != nil {
			return nil
		}
		detachVolumesTask := &models.Task{
			JobId:          f.Job.JobId,
			Owner:          f.Job.Owner,
			TaskAction:     ActionDetachVolumes,
			Target:         f.Runtime.Provider,
			NodeId:         nodeId,
			Directive:      directive,
			FailureAllowed: failureAllowed,
		}
		taskLayer.Tasks = append(taskLayer.Tasks, detachVolumesTask)
	}
	return taskLayer
}

func (f *Frame) attachVolumesLayer(failureAllowed bool) *models.TaskLayer {
	taskLayer := new(models.TaskLayer)
	for nodeId, clusterNode := range f.ClusterWrapper.ClusterNodes {
		volume := &models.Volume{
			Name:       clusterNode.ClusterId + "_" + nodeId,
			Zone:       f.Runtime.Zone,
			RuntimeId:  f.Runtime.RuntimeId,
			VolumeId:   clusterNode.VolumeId,
			InstanceId: clusterNode.InstanceId,
		}
		directive, err := volume.ToString()
		if err != nil {
			return nil
		}
		attachVolumesTask := &models.Task{
			JobId:          f.Job.JobId,
			Owner:          f.Job.Owner,
			TaskAction:     ActionAttachVolumes,
			Target:         f.Runtime.Provider,
			NodeId:         nodeId,
			Directive:      directive,
			FailureAllowed: failureAllowed,
		}
		taskLayer.Tasks = append(taskLayer.Tasks, attachVolumesTask)
	}
	return taskLayer
}

func (f *Frame) deleteVolumesLayer(nodeIds []string, failureAllowed bool) *models.TaskLayer {
	taskLayer := new(models.TaskLayer)
	for _, nodeId := range nodeIds {
		clusterNode := f.ClusterWrapper.ClusterNodes[nodeId]
		volume := &models.Volume{
			Name:       clusterNode.ClusterId + "_" + nodeId,
			Zone:       f.Runtime.Zone,
			RuntimeId:  f.Runtime.RuntimeId,
			VolumeId:   clusterNode.VolumeId,
			InstanceId: clusterNode.InstanceId,
		}
		directive, err := volume.ToString()
		if err != nil {
			return nil
		}
		deleteVolumesTask := &models.Task{
			JobId:          f.Job.JobId,
			Owner:          f.Job.Owner,
			TaskAction:     ActionDeleteVolumes,
			Target:         f.Runtime.Provider,
			NodeId:         nodeId,
			Directive:      directive,
			FailureAllowed: failureAllowed,
		}
		taskLayer.Tasks = append(taskLayer.Tasks, deleteVolumesTask)
	}
	return taskLayer
}

func (f *Frame) formatAndMountVolumeLayer(nodeIds []string, failureAllowed bool) *models.TaskLayer {
	taskLayer := new(models.TaskLayer)

	for _, nodeId := range nodeIds {
		clusterNode := f.ClusterWrapper.ClusterNodes[nodeId]
		// cmd will be assigned when the task is handling
		meta := &models.Meta{
			FrontgateId: f.ClusterWrapper.Cluster.FrontgateId,
			Timeout:     TimeoutFormatAndMountVolume,
			NodeId:      clusterNode.NodeId,
			DroneIp:     clusterNode.PrivateIp,
		}
		directive, err := meta.ToString()
		if err != nil {
			return nil
		}
		formatVolumeTask := &models.Task{
			JobId:          f.Job.JobId,
			Owner:          f.Job.Owner,
			TaskAction:     ActionFormatAndMountVolume,
			Target:         constants.TargetPilot,
			NodeId:         nodeId,
			Directive:      string(directive),
			FailureAllowed: failureAllowed,
		}
		taskLayer.Tasks = append(taskLayer.Tasks, formatVolumeTask)
	}
	if len(taskLayer.Tasks) > 0 {
		return taskLayer
	} else {
		return nil
	}
}

func (f *Frame) sshKeygenLayer(failureAllowed bool) *models.TaskLayer {
	taskLayer := new(models.TaskLayer)
	ctx := client.GetSystemUserContext()
	clusterClient, err := clusterclient.NewClient(ctx)
	if err != nil {
		logger.Error("New ssh key gen task layer failed: %+v", err)
		return nil
	}

	for nodeId, clusterNode := range f.ClusterWrapper.ClusterNodes {
		role := clusterNode.Role
		clusterCommon := f.ClusterWrapper.ClusterCommons[role]
		keyType := clusterCommon.Passphraseless
		if keyType != "" {
			private, public, err := sshutil.MakeSSHKeyPair(keyType)
			if err != nil {
				logger.Error("Generate ssh key [%s] in cluster node [%s] failed",
					clusterCommon.Passphraseless, nodeId)
				return nil
			}
			_, err = clusterClient.ModifyClusterNode(ctx, &pb.ModifyClusterNodeRequest{
				ClusterNode: &pb.ClusterNode{
					NodeId: pbutil.ToProtoString(nodeId),
					PubKey: pbutil.ToProtoString(public),
				},
			})
			cmd := fmt.Sprintf("mkdir -p /root/.ssh/;chmod 700 /root/.ssh/;"+
				"echo \"%s\" > /root/.ssh/id_%s;echo \"%s\" > /root/.ssh/id_%s.pub;"+
				"chown 600 /root/.ssh/id_%s;chown 644 /root/.ssh/id_%s.pub",
				private, keyType, public, keyType, keyType, keyType)
			ip := clusterNode.PrivateIp
			cnodes := &models.Cmd{
				Cmd:     cmd,
				Timeout: TimeoutSshKeygen,
			}
			meta := &models.Meta{
				FrontgateId: f.ClusterWrapper.Cluster.FrontgateId,
				Timeout:     TimeoutSshKeygen,
				NodeId:      clusterNode.NodeId,
				DroneIp:     ip,
				Cnodes:      jsonutil.ToString(cnodes),
			}
			directive, err := meta.ToString()
			if err != nil {
				return nil
			}
			formatVolumeTask := &models.Task{
				JobId:          f.Job.JobId,
				Owner:          f.Job.Owner,
				TaskAction:     ActionRegisterCmd,
				Target:         constants.TargetPilot,
				NodeId:         nodeId,
				Directive:      string(directive),
				FailureAllowed: failureAllowed,
			}
			taskLayer.Tasks = append(taskLayer.Tasks, formatVolumeTask)
		}
	}
	if len(taskLayer.Tasks) > 0 {
		return taskLayer
	} else {
		return nil
	}
}

func (f *Frame) umountVolumeLayer(nodeIds []string, failureAllowed bool) *models.TaskLayer {
	taskLayer := new(models.TaskLayer)

	for _, nodeId := range nodeIds {
		clusterNode := f.ClusterWrapper.ClusterNodes[nodeId]
		clusterRole := f.ClusterWrapper.ClusterRoles[clusterNode.Role]
		cmd := UmountVolumeCmd(clusterRole.MountPoint)
		ip := clusterNode.PrivateIp
		cnodes := &models.Cmd{
			Cmd:     cmd,
			Timeout: TimeoutUmountVolume,
		}
		meta := &models.Meta{
			FrontgateId: f.ClusterWrapper.Cluster.FrontgateId,
			Timeout:     TimeoutUmountVolume,
			NodeId:      clusterNode.NodeId,
			DroneIp:     ip,
			Cnodes:      jsonutil.ToString(cnodes),
		}
		directive, err := meta.ToString()
		if err != nil {
			return nil
		}
		umountVolumeTask := &models.Task{
			JobId:          f.Job.JobId,
			Owner:          f.Job.Owner,
			TaskAction:     ActionRegisterCmd,
			Target:         constants.TargetPilot,
			NodeId:         nodeId,
			Directive:      string(directive),
			FailureAllowed: failureAllowed,
		}
		taskLayer.Tasks = append(taskLayer.Tasks, umountVolumeTask)
	}
	if len(taskLayer.Tasks) > 0 {
		return taskLayer
	} else {
		return nil
	}
}

/*
cat /opt/openpitrix/conf/drone.conf
IMAGE="mysql:5.7"
MOUNT_POINT="/data"
FILE_NAME="drone.conf"
FILE_CONF={\\"id\\":\\"cln-abcdefgh\\",\\"listen_port\\":9112}
*/
func (f *Frame) getUserDataValue(nodeId string) string {
	var result string
	clusterNode := f.ClusterWrapper.ClusterNodes[nodeId]
	role := clusterNode.Role
	if strings.HasSuffix(role, constants.ReplicaRoleSuffix) {
		role = string([]byte(role)[:len(role)-len(constants.ReplicaRoleSuffix)])
	}
	clusterRole, _ := f.ClusterWrapper.ClusterRoles[role]
	clusterCommon, _ := f.ClusterWrapper.ClusterCommons[role]
	mountPoint := clusterRole.MountPoint
	// Empty string can not be a parameter
	if len(mountPoint) == 0 {
		mountPoint = "#"
	}
	imageId := clusterCommon.ImageId

	droneConf := make(map[string]interface{})
	droneConf["id"] = nodeId
	droneConf["listen_port"] = constants.DroneServicePort
	droneConfStr := strings.Replace(jsonutil.ToString(droneConf), "\"", "\\\\\"", -1)

	result += fmt.Sprintf("IMAGE=\"%s\"\n", imageId)
	result += fmt.Sprintf("MOUNT_POINT=\"%s\"\n", mountPoint)
	result += fmt.Sprintf("FILE_NAME=\"%s\"\n", DroneConfFile)
	result += fmt.Sprintf("FILE_CONF=%s\n", droneConfStr)

	return result
}

func (f *Frame) getUserDataPath() string {
	return MetadataConfPath + OpenPitrixConfFile
}

func (f *Frame) getConfig(nodeId string) string {
	clusterNode := f.ClusterWrapper.ClusterNodes[nodeId]

	droneEndpoint := &pbtypes.DroneEndpoint{
		FrontgateId: f.ClusterWrapper.Cluster.FrontgateId,
		DroneIp:     clusterNode.PrivateIp,
		DronePort:   constants.DroneServicePort,
	}
	droneConfig := &pbtypes.DroneConfig{
		Id:             nodeId,
		Host:           clusterNode.PrivateIp,
		ListenPort:     constants.DroneServicePort,
		CmdInfoLogPath: ConfdCmdLogPath,
		ConfdSelfHost:  clusterNode.PrivateIp,
		LogLevel:       MetadataLogLevel,
	}
	config := &pbtypes.SetDroneConfigRequest{
		Endpoint: droneEndpoint,
		Config:   droneConfig,
	}
	return jsonutil.ToString(config)
}

func (f *Frame) setDroneConfigLayer(nodeIds []string, failureAllowed bool) *models.TaskLayer {
	var tasks []*models.Task
	for _, nodeId := range nodeIds {
		task := &models.Task{
			JobId:          f.Job.JobId,
			Owner:          f.Job.Owner,
			TaskAction:     ActionSetDroneConfig,
			Target:         constants.TargetPilot,
			NodeId:         nodeId,
			Directive:      f.getConfig(nodeId),
			FailureAllowed: failureAllowed,
		}
		tasks = append(tasks, task)
	}
	return &models.TaskLayer{
		Tasks: tasks,
	}
}

func (f *Frame) runInstancesLayer(nodeIds []string, failureAllowed bool) *models.TaskLayer {
	taskLayer := new(models.TaskLayer)
	apiServer := f.Runtime.RuntimeUrl
	zone := f.Runtime.Zone
	globalPi := pi.Global()
	imageId := ""
	var err error
	if globalPi != nil {
		imageId, err = globalPi.GlobalConfig().GetRuntimeImageId(apiServer, zone)
		if err != nil {
			return nil
		}
	}
	for _, nodeId := range nodeIds {
		clusterNode := f.ClusterWrapper.ClusterNodes[nodeId]
		role := clusterNode.Role
		if strings.HasSuffix(role, constants.ReplicaRoleSuffix) {
			role = string([]byte(role)[:len(role)-len(constants.ReplicaRoleSuffix)])
		}
		clusterRole, exist := f.ClusterWrapper.ClusterRoles[role]
		if !exist {
			logger.Error("No such role [%s] in cluster role [%s]. ",
				role, f.ClusterWrapper.Cluster.ClusterId)
			return nil
		}
		instance := &models.Instance{
			Name:          clusterNode.ClusterId + "_" + nodeId,
			NodeId:        nodeId,
			ImageId:       imageId,
			Cpu:           int(clusterRole.Cpu),
			Memory:        int(clusterRole.Memory),
			Gpu:           int(clusterRole.Gpu),
			Subnet:        clusterNode.SubnetId,
			RuntimeId:     f.Runtime.RuntimeId,
			Zone:          f.Runtime.Zone,
			LoginPasswd:   DefaultLoginPasswd,
			UserdataPath:  f.getUserDataPath(),
			UserDataValue: f.getUserDataValue(nodeId),
		}
		instanceTaskDirective, err := instance.ToString()
		if err != nil {
			return nil
		}
		runInstanceTask := &models.Task{
			JobId:          f.Job.JobId,
			Owner:          f.Job.Owner,
			TaskAction:     ActionRunInstances,
			Target:         f.Runtime.Provider,
			NodeId:         nodeId,
			Directive:      instanceTaskDirective,
			FailureAllowed: failureAllowed,
		}
		taskLayer.Tasks = append(taskLayer.Tasks, runInstanceTask)
	}

	if len(taskLayer.Tasks) > 0 {
		return taskLayer
	} else {
		return nil
	}
}

func (f *Frame) stopInstancesLayer(nodeIds []string, failureAllowed bool) *models.TaskLayer {
	taskLayer := new(models.TaskLayer)
	for _, nodeId := range nodeIds {
		clusterNode := f.ClusterWrapper.ClusterNodes[nodeId]
		instance := &models.Instance{
			Name:       clusterNode.ClusterId + "_" + nodeId,
			NodeId:     nodeId,
			InstanceId: clusterNode.InstanceId,
			RuntimeId:  f.Runtime.RuntimeId,
			Zone:       f.Runtime.Zone,
		}
		instanceTaskDirective, err := instance.ToString()
		if err != nil {
			return nil
		}
		stopInstanceTask := &models.Task{
			JobId:          f.Job.JobId,
			Owner:          f.Job.Owner,
			TaskAction:     ActionStopInstances,
			Target:         f.Runtime.Provider,
			NodeId:         nodeId,
			Directive:      instanceTaskDirective,
			FailureAllowed: failureAllowed,
		}
		taskLayer.Tasks = append(taskLayer.Tasks, stopInstanceTask)
	}

	if len(taskLayer.Tasks) > 0 {
		return taskLayer
	} else {
		return nil
	}
}

func (f *Frame) deleteInstancesLayer(nodeIds []string, failureAllowed bool) *models.TaskLayer {
	taskLayer := new(models.TaskLayer)
	for _, nodeId := range nodeIds {
		clusterNode := f.ClusterWrapper.ClusterNodes[nodeId]
		instance := &models.Instance{
			Name:       clusterNode.ClusterId + "_" + nodeId,
			NodeId:     nodeId,
			InstanceId: clusterNode.InstanceId,
			RuntimeId:  f.Runtime.RuntimeId,
			Zone:       f.Runtime.Zone,
		}
		instanceTaskDirective, err := instance.ToString()
		if err != nil {
			return nil
		}
		deleteInstanceTask := &models.Task{
			JobId:          f.Job.JobId,
			Owner:          f.Job.Owner,
			TaskAction:     ActionTerminateInstances,
			Target:         f.Runtime.Provider,
			NodeId:         nodeId,
			Directive:      instanceTaskDirective,
			FailureAllowed: false,
		}
		taskLayer.Tasks = append(taskLayer.Tasks, deleteInstanceTask)
	}

	if len(taskLayer.Tasks) > 0 {
		return taskLayer
	} else {
		return nil
	}
}

func (f *Frame) startInstancesLayer(failureAllowed bool) *models.TaskLayer {
	taskLayer := new(models.TaskLayer)
	for nodeId, clusterNode := range f.ClusterWrapper.ClusterNodes {
		instance := &models.Instance{
			Name:       clusterNode.ClusterId + "_" + nodeId,
			NodeId:     nodeId,
			InstanceId: clusterNode.InstanceId,
			RuntimeId:  f.Runtime.RuntimeId,
			Zone:       f.Runtime.Zone,
		}
		instanceTaskDirective, err := instance.ToString()
		if err != nil {
			return nil
		}
		startInstanceTask := &models.Task{
			JobId:          f.Job.JobId,
			Owner:          f.Job.Owner,
			TaskAction:     ActionStartInstances,
			Target:         f.Runtime.Provider,
			NodeId:         nodeId,
			Directive:      instanceTaskDirective,
			FailureAllowed: failureAllowed,
		}
		taskLayer.Tasks = append(taskLayer.Tasks, startInstanceTask)
	}

	if len(taskLayer.Tasks) > 0 {
		return taskLayer
	} else {
		return nil
	}
}

func (f *Frame) waitFrontgateLayer(failureAllowed bool) *models.TaskLayer {
	meta := &models.Meta{
		FrontgateId: f.ClusterWrapper.Cluster.FrontgateId,
	}
	directive, err := meta.ToString()
	if err != nil {
		return nil
	}
	// Wait frontgate available
	waitFrontgateTask := &models.Task{
		JobId:          f.Job.JobId,
		Owner:          f.Job.Owner,
		TaskAction:     ActionWaitFrontgateAvailable,
		Target:         f.Runtime.Provider,
		NodeId:         f.ClusterWrapper.Cluster.ClusterId,
		Directive:      string(directive),
		FailureAllowed: failureAllowed,
	}
	return &models.TaskLayer{
		Tasks: []*models.Task{waitFrontgateTask},
	}
}

func (f *Frame) registerMetadataLayer(failureAllowed bool) *models.TaskLayer {
	// When the task is handled by task controller, the cnodes will be filled in,
	meta := &models.Meta{
		FrontgateId: f.ClusterWrapper.Cluster.FrontgateId,
		Timeout:     TimeoutRegister,
		ClusterId:   f.ClusterWrapper.Cluster.ClusterId,
		Cnodes:      "",
	}
	directive, err := meta.ToString()
	if err != nil {
		return nil
	}
	registerMetadataTask := &models.Task{
		JobId:          f.Job.JobId,
		Owner:          f.Job.Owner,
		TaskAction:     ActionRegisterMetadata,
		Target:         constants.TargetPilot,
		NodeId:         f.ClusterWrapper.Cluster.ClusterId,
		Directive:      string(directive),
		FailureAllowed: failureAllowed,
	}
	return &models.TaskLayer{
		Tasks: []*models.Task{registerMetadataTask},
	}
}

func (f *Frame) registerNodesMetadataLayer(nodeIds []string, failureAllowed bool) *models.TaskLayer {
	metadata := &MetadataV1{
		ClusterWrapper: f.ClusterWrapper,
	}
	cnodes := jsonutil.ToString(metadata.GetClusterNodeCnodes(nodeIds))
	meta := &models.Meta{
		FrontgateId: f.ClusterWrapper.Cluster.FrontgateId,
		Timeout:     TimeoutRegister,
		ClusterId:   f.ClusterWrapper.Cluster.ClusterId,
		Cnodes:      cnodes,
	}
	directive, err := meta.ToString()
	if err != nil {
		return nil
	}
	task := &models.Task{
		JobId:          f.Job.JobId,
		Owner:          f.Job.Owner,
		TaskAction:     ActionRegisterMetadata,
		Target:         constants.TargetPilot,
		NodeId:         f.ClusterWrapper.Cluster.ClusterId,
		Directive:      string(directive),
		FailureAllowed: failureAllowed,
	}
	return &models.TaskLayer{
		Tasks: []*models.Task{task},
	}
}

func (f *Frame) registerScalingNodesMetadataLayer(nodeIds []string, path string, failureAllowed bool) *models.TaskLayer {
	clusterId := f.ClusterWrapper.Cluster.ClusterId
	metadata := &MetadataV1{
		ClusterWrapper: f.ClusterWrapper,
	}
	scalingCnodes := metadata.GetScalingCnodes(nodeIds, path)
	if scalingCnodes == nil {
		logger.Info("No new nodes for cluster [%s] is registered", clusterId)
		return nil
	}
	meta := &models.Meta{
		FrontgateId: f.ClusterWrapper.Cluster.FrontgateId,
		Timeout:     TimeoutRegister,
		ClusterId:   f.ClusterWrapper.Cluster.ClusterId,
		Cnodes:      jsonutil.ToString(scalingCnodes),
	}
	directive, err := meta.ToString()
	if err != nil {
		return nil
	}
	task := &models.Task{
		JobId:          f.Job.JobId,
		Owner:          f.Job.Owner,
		TaskAction:     ActionRegisterMetadata,
		Target:         constants.TargetPilot,
		NodeId:         f.ClusterWrapper.Cluster.ClusterId,
		Directive:      string(directive),
		FailureAllowed: failureAllowed,
	}
	return &models.TaskLayer{
		Tasks: []*models.Task{task},
	}
}

func (f *Frame) deregisterNodesMetadataLayer(nodeIds []string, failureAllowed bool) *models.TaskLayer {
	metadata := &MetadataV1{
		ClusterWrapper: f.ClusterWrapper,
	}
	cnodes := jsonutil.ToString(metadata.GetEmptyClusterNodeCnodes(nodeIds))
	meta := &models.Meta{
		FrontgateId: f.ClusterWrapper.Cluster.FrontgateId,
		Timeout:     TimeoutDeregister,
		ClusterId:   f.ClusterWrapper.Cluster.ClusterId,
		Cnodes:      cnodes,
	}
	directive, err := meta.ToString()
	if err != nil {
		return nil
	}
	task := &models.Task{
		JobId:          f.Job.JobId,
		Owner:          f.Job.Owner,
		TaskAction:     ActionDeregisterMetadata,
		Target:         constants.TargetPilot,
		NodeId:         f.ClusterWrapper.Cluster.ClusterId,
		Directive:      string(directive),
		FailureAllowed: failureAllowed,
	}
	return &models.TaskLayer{
		Tasks: []*models.Task{task},
	}
}

func (f *Frame) deregisterScalingNodesMetadataLayer(path string, failureAllowed bool) *models.TaskLayer {
	clusterId := f.ClusterWrapper.Cluster.ClusterId
	cnodes := map[string]interface{}{
		RegisterClustersRootPath: map[string]interface{}{
			clusterId: map[string]interface{}{
				path: "",
			},
		},
	}
	meta := &models.Meta{
		FrontgateId: f.ClusterWrapper.Cluster.FrontgateId,
		Timeout:     TimeoutDeregister,
		ClusterId:   f.ClusterWrapper.Cluster.ClusterId,
		Cnodes:      jsonutil.ToString(cnodes),
	}
	directive, err := meta.ToString()
	if err != nil {
		return nil
	}
	task := &models.Task{
		JobId:          f.Job.JobId,
		Owner:          f.Job.Owner,
		TaskAction:     ActionDeregisterMetadata,
		Target:         constants.TargetPilot,
		NodeId:         f.ClusterWrapper.Cluster.ClusterId,
		Directive:      string(directive),
		FailureAllowed: failureAllowed,
	}
	return &models.TaskLayer{
		Tasks: []*models.Task{task},
	}
}

func (f *Frame) deregisterMetadataLayer(failureAllowed bool) *models.TaskLayer {
	metadata := &MetadataV1{
		ClusterWrapper: f.ClusterWrapper,
	}
	cnodes := metadata.GetEmptyClusterCnodes()
	meta := &models.Meta{
		FrontgateId: f.ClusterWrapper.Cluster.FrontgateId,
		Timeout:     TimeoutDeregister,
		ClusterId:   f.ClusterWrapper.Cluster.ClusterId,
		Cnodes:      jsonutil.ToString(cnodes),
	}
	directive, err := meta.ToString()
	if err != nil {
		return nil
	}
	deregisterMetadataTask := &models.Task{
		JobId:          f.Job.JobId,
		Owner:          f.Job.Owner,
		TaskAction:     ActionDeregisterMetadata,
		Target:         constants.TargetPilot,
		NodeId:         f.ClusterWrapper.Cluster.ClusterId,
		Directive:      directive,
		FailureAllowed: failureAllowed,
	}
	return &models.TaskLayer{
		Tasks: []*models.Task{deregisterMetadataTask},
	}
}

func (f *Frame) CreateClusterLayer() *models.TaskLayer {
	var nodeIds []string
	for nodeId := range f.ClusterWrapper.ClusterNodes {
		nodeIds = append(nodeIds, nodeId)
	}
	headTaskLayer := new(models.TaskLayer)

	headTaskLayer.
		Append(f.createVolumesLayer(nodeIds, false)).        // create volume
		Append(f.runInstancesLayer(nodeIds, false)).         // run instance and attach volume to instance
		Append(f.formatAndMountVolumeLayer(nodeIds, false)). // format and mount volume to instance
		Append(f.waitFrontgateLayer(false)).                 // wait frontgate cluster to be active
		Append(f.sshKeygenLayer(false)).                     // generate ssh key
		Append(f.registerMetadataLayer(false)).              // register cluster metadata
		Append(f.setDroneConfigLayer(nodeIds, false)).       // set drone config
		Append(f.startConfdServiceLayer(nodeIds, false)).    // start confd service
		Append(f.initAndStartServiceLayer(nodeIds, false)).  // register init and start cmd to exec
		Append(f.deregisterCmdLayer(nodeIds, true))          // deregister cmd

	return headTaskLayer.Child
}

func (f *Frame) StopClusterLayer() *models.TaskLayer {
	var nodeIds []string
	for nodeId := range f.ClusterWrapper.ClusterNodes {
		nodeIds = append(nodeIds, nodeId)
	}
	headTaskLayer := new(models.TaskLayer)

	headTaskLayer.
		Append(f.waitFrontgateLayer(true)).             // wait frontgate cluster to be active
		Append(f.stopServiceLayer(nodeIds, true)).      // register stop cmd to exec
		Append(f.stopConfdServiceLayer(nodeIds, true)). // stop confd service
		Append(f.umountVolumeLayer(nodeIds, true)).     // umount volume from instance
		Append(f.detachVolumesLayer(nodeIds, false)).   // detach volume from instance
		Append(f.stopInstancesLayer(nodeIds, false)).   // stop instance
		Append(f.deregisterMetadataLayer(true))         // deregister cluster

	return headTaskLayer.Child
}

func (f *Frame) StartClusterLayer() *models.TaskLayer {
	var nodeIds []string
	for nodeId := range f.ClusterWrapper.ClusterNodes {
		nodeIds = append(nodeIds, nodeId)
	}
	headTaskLayer := new(models.TaskLayer)

	headTaskLayer.
		Append(f.startInstancesLayer(false)).             // start instance
		Append(f.attachVolumesLayer(false)).              // attach volume to instance, will auto mount
		Append(f.waitFrontgateLayer(false)).              // wait frontgate cluster to be active
		Append(f.registerMetadataLayer(false)).           // register cluster metadata
		Append(f.startConfdServiceLayer(nodeIds, false)). // start confd service
		Append(f.startServiceLayer(nodeIds, false)).      // register start cmd to exec
		Append(f.deregisterCmdLayer(nodeIds, true))       // deregister cmd

	return headTaskLayer.Child
}

func (f *Frame) DeleteClusterLayer() *models.TaskLayer {
	var nodeIds []string
	for nodeId := range f.ClusterWrapper.ClusterNodes {
		nodeIds = append(nodeIds, nodeId)
	}
	headTaskLayer := new(models.TaskLayer)

	headTaskLayer.
		Append(f.destroyAndStopServiceLayer(nodeIds, nil, true)). // register destroy and stop cmd to exec
		Append(f.stopConfdServiceLayer(nodeIds, true)).           // stop confd service
		Append(f.umountVolumeLayer(nodeIds, true)).               // umount volume from instance
		Append(f.detachVolumesLayer(nodeIds, false)).             // detach volume from instance
		Append(f.deleteInstancesLayer(nodeIds, false)).           // delete instance
		Append(f.deleteVolumesLayer(nodeIds, false)).             // delete volume
		Append(f.deregisterMetadataLayer(true))                   // deregister cluster

	return headTaskLayer.Child
}

func (f *Frame) AddClusterNodesLayer() *models.TaskLayer {
	var addNodeIds, nonAddNodeIds []string
	for nodeId, node := range f.ClusterWrapper.ClusterNodes {
		if node.Status == constants.StatusPending {
			addNodeIds = append(addNodeIds, nodeId)
		} else {
			nonAddNodeIds = append(nonAddNodeIds, nodeId)
		}
	}
	headTaskLayer := new(models.TaskLayer)

	headTaskLayer.
		Append(f.scaleOutPreCheckServiceLayer(nonAddNodeIds, false)).                       // register scale out pre check to exec
		Append(f.createVolumesLayer(addNodeIds, false)).                                    // create volume
		Append(f.runInstancesLayer(addNodeIds, false)).                                     // run instance and attach volume to instance
		Append(f.formatAndMountVolumeLayer(addNodeIds, false)).                             // format and mount volume to instance
		Append(f.registerNodesMetadataLayer(addNodeIds, false)).                            // register cluster nodes metadata
		Append(f.registerScalingNodesMetadataLayer(addNodeIds, RegisterNodeAdding, false)). // register adding hosts metadata
		Append(f.startConfdServiceLayer(addNodeIds, false)).                                // start confd service
		Append(f.initAndStartServiceLayer(addNodeIds, false)).                              // register init and start cmd to exec
		Append(f.scaleOutServiceLayer(nonAddNodeIds, false)).                               // register scale out cmd to exec
		Append(f.deregisterScalingNodesMetadataLayer(RegisterNodeAdding, true))             // deregister adding host metadata
	return headTaskLayer.Child
}

func (f *Frame) DeleteClusterNodesLayer() *models.TaskLayer {
	var deleteNodeIds, nonDeleteNodeIds []string
	for nodeId, node := range f.ClusterWrapper.ClusterNodes {
		if node.Status == constants.StatusDeleting {
			deleteNodeIds = append(deleteNodeIds, nodeId)
		} else {
			nonDeleteNodeIds = append(nonDeleteNodeIds, nodeId)
		}
	}
	headTaskLayer := new(models.TaskLayer)

	headTaskLayer.
		Append(f.registerScalingNodesMetadataLayer(deleteNodeIds, RegisterNodeDeleting, false)).                    // register scale in node metadata
		Append(f.scaleInPreCheckServiceLayer(nonDeleteNodeIds, false)).                                             // register scale in pre check to exec
		Append(f.destroyAndStopServiceLayer(deleteNodeIds, f.scaleInServiceLayer(nonDeleteNodeIds, false), false)). // register destroy, scale in and stop cmd to exec
		Append(f.stopConfdServiceLayer(deleteNodeIds, false)).                                                      // stop confd service
		Append(f.umountVolumeLayer(deleteNodeIds, false)).                                                          // umount volume from instance
		Append(f.detachVolumesLayer(deleteNodeIds, false)).                                                         // detach volume from instance
		Append(f.deleteInstancesLayer(deleteNodeIds, false)).                                                       // delete instance
		Append(f.deleteVolumesLayer(deleteNodeIds, false)).                                                         // delete volume
		Append(f.deregisterNodesMetadataLayer(deleteNodeIds, false)).                                               // deregister deleting cluster nodes metadata
		Append(f.deregisterScalingNodesMetadataLayer(RegisterNodeDeleting, false))                                  // deregister deleting nodes metadata
	return headTaskLayer.Child
}
