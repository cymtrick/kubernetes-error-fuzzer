// Code generated by protoc-gen-gogo.
// source: containerizer.proto
// DO NOT EDIT!

package mesosproto

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// *
// Encodes the launch command sent to the external containerizer
// program.
type Launch struct {
	ContainerId      *ContainerID  `protobuf:"bytes,1,req,name=container_id" json:"container_id,omitempty"`
	TaskInfo         *TaskInfo     `protobuf:"bytes,2,opt,name=task_info" json:"task_info,omitempty"`
	ExecutorInfo     *ExecutorInfo `protobuf:"bytes,3,opt,name=executor_info" json:"executor_info,omitempty"`
	Directory        *string       `protobuf:"bytes,4,opt,name=directory" json:"directory,omitempty"`
	User             *string       `protobuf:"bytes,5,opt,name=user" json:"user,omitempty"`
	SlaveId          *SlaveID      `protobuf:"bytes,6,opt,name=slave_id" json:"slave_id,omitempty"`
	SlavePid         *string       `protobuf:"bytes,7,opt,name=slave_pid" json:"slave_pid,omitempty"`
	Checkpoint       *bool         `protobuf:"varint,8,opt,name=checkpoint" json:"checkpoint,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Launch) Reset()         { *m = Launch{} }
func (m *Launch) String() string { return proto.CompactTextString(m) }
func (*Launch) ProtoMessage()    {}

func (m *Launch) GetContainerId() *ContainerID {
	if m != nil {
		return m.ContainerId
	}
	return nil
}

func (m *Launch) GetTaskInfo() *TaskInfo {
	if m != nil {
		return m.TaskInfo
	}
	return nil
}

func (m *Launch) GetExecutorInfo() *ExecutorInfo {
	if m != nil {
		return m.ExecutorInfo
	}
	return nil
}

func (m *Launch) GetDirectory() string {
	if m != nil && m.Directory != nil {
		return *m.Directory
	}
	return ""
}

func (m *Launch) GetUser() string {
	if m != nil && m.User != nil {
		return *m.User
	}
	return ""
}

func (m *Launch) GetSlaveId() *SlaveID {
	if m != nil {
		return m.SlaveId
	}
	return nil
}

func (m *Launch) GetSlavePid() string {
	if m != nil && m.SlavePid != nil {
		return *m.SlavePid
	}
	return ""
}

func (m *Launch) GetCheckpoint() bool {
	if m != nil && m.Checkpoint != nil {
		return *m.Checkpoint
	}
	return false
}

// *
// Encodes the update command sent to the external containerizer
// program.
type Update struct {
	ContainerId      *ContainerID `protobuf:"bytes,1,req,name=container_id" json:"container_id,omitempty"`
	Resources        []*Resource  `protobuf:"bytes,2,rep,name=resources" json:"resources,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Update) Reset()         { *m = Update{} }
func (m *Update) String() string { return proto.CompactTextString(m) }
func (*Update) ProtoMessage()    {}

func (m *Update) GetContainerId() *ContainerID {
	if m != nil {
		return m.ContainerId
	}
	return nil
}

func (m *Update) GetResources() []*Resource {
	if m != nil {
		return m.Resources
	}
	return nil
}

// *
// Encodes the wait command sent to the external containerizer
// program.
type Wait struct {
	ContainerId      *ContainerID `protobuf:"bytes,1,req,name=container_id" json:"container_id,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Wait) Reset()         { *m = Wait{} }
func (m *Wait) String() string { return proto.CompactTextString(m) }
func (*Wait) ProtoMessage()    {}

func (m *Wait) GetContainerId() *ContainerID {
	if m != nil {
		return m.ContainerId
	}
	return nil
}

// *
// Encodes the destroy command sent to the external containerizer
// program.
type Destroy struct {
	ContainerId      *ContainerID `protobuf:"bytes,1,req,name=container_id" json:"container_id,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Destroy) Reset()         { *m = Destroy{} }
func (m *Destroy) String() string { return proto.CompactTextString(m) }
func (*Destroy) ProtoMessage()    {}

func (m *Destroy) GetContainerId() *ContainerID {
	if m != nil {
		return m.ContainerId
	}
	return nil
}

// *
// Encodes the usage command sent to the external containerizer
// program.
type Usage struct {
	ContainerId      *ContainerID `protobuf:"bytes,1,req,name=container_id" json:"container_id,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Usage) Reset()         { *m = Usage{} }
func (m *Usage) String() string { return proto.CompactTextString(m) }
func (*Usage) ProtoMessage()    {}

func (m *Usage) GetContainerId() *ContainerID {
	if m != nil {
		return m.ContainerId
	}
	return nil
}

// *
// Information about a container termination, returned by the
// containerizer to the slave.
type Termination struct {
	// A container may be killed if it exceeds its resources; this will
	// be indicated by killed=true and described by the message string.
	// TODO(jaybuff): As part of MESOS-2035 we should remove killed and
	// replace it with a TaskStatus::Reason.
	Killed  *bool   `protobuf:"varint,1,req,name=killed" json:"killed,omitempty"`
	Message *string `protobuf:"bytes,2,req,name=message" json:"message,omitempty"`
	// Exit status of the process.
	Status           *int32 `protobuf:"varint,3,opt,name=status" json:"status,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Termination) Reset()         { *m = Termination{} }
func (m *Termination) String() string { return proto.CompactTextString(m) }
func (*Termination) ProtoMessage()    {}

func (m *Termination) GetKilled() bool {
	if m != nil && m.Killed != nil {
		return *m.Killed
	}
	return false
}

func (m *Termination) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *Termination) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

// *
// Information on all active containers returned by the containerizer
// to the slave.
type Containers struct {
	Containers       []*ContainerID `protobuf:"bytes,1,rep,name=containers" json:"containers,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Containers) Reset()         { *m = Containers{} }
func (m *Containers) String() string { return proto.CompactTextString(m) }
func (*Containers) ProtoMessage()    {}

func (m *Containers) GetContainers() []*ContainerID {
	if m != nil {
		return m.Containers
	}
	return nil
}
