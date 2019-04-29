// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/mdb/redis/v1/cluster.proto

package redis // import "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/redis/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import config "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/redis/v1/config"
import timeofday "google.golang.org/genproto/googleapis/type/timeofday"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Cluster_Environment int32

const (
	Cluster_ENVIRONMENT_UNSPECIFIED Cluster_Environment = 0
	// Stable environment with a conservative update policy:
	// only hotfixes are applied during regular maintenance.
	Cluster_PRODUCTION Cluster_Environment = 1
	// Environment with more aggressive update policy: new versions
	// are rolled out irrespective of backward compatibility.
	Cluster_PRESTABLE Cluster_Environment = 2
)

var Cluster_Environment_name = map[int32]string{
	0: "ENVIRONMENT_UNSPECIFIED",
	1: "PRODUCTION",
	2: "PRESTABLE",
}
var Cluster_Environment_value = map[string]int32{
	"ENVIRONMENT_UNSPECIFIED": 0,
	"PRODUCTION":              1,
	"PRESTABLE":               2,
}

func (x Cluster_Environment) String() string {
	return proto.EnumName(Cluster_Environment_name, int32(x))
}
func (Cluster_Environment) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cluster_6d995a7e51dda886, []int{0, 0}
}

type Cluster_Health int32

const (
	// Cluster is in unknown state (we have no data)
	Cluster_HEALTH_UNKNOWN Cluster_Health = 0
	// Cluster is alive and well (all hosts are alive)
	Cluster_ALIVE Cluster_Health = 1
	// Cluster is inoperable (it cannot perform any of its essential functions)
	Cluster_DEAD Cluster_Health = 2
	// Cluster is partially alive (it can perform some of its essential functions)
	Cluster_DEGRADED Cluster_Health = 3
)

var Cluster_Health_name = map[int32]string{
	0: "HEALTH_UNKNOWN",
	1: "ALIVE",
	2: "DEAD",
	3: "DEGRADED",
}
var Cluster_Health_value = map[string]int32{
	"HEALTH_UNKNOWN": 0,
	"ALIVE":          1,
	"DEAD":           2,
	"DEGRADED":       3,
}

func (x Cluster_Health) String() string {
	return proto.EnumName(Cluster_Health_name, int32(x))
}
func (Cluster_Health) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cluster_6d995a7e51dda886, []int{0, 1}
}

type Cluster_Status int32

const (
	// Cluster status is unknown
	Cluster_STATUS_UNKNOWN Cluster_Status = 0
	// Cluster is being created
	Cluster_CREATING Cluster_Status = 1
	// Cluster is running
	Cluster_RUNNING Cluster_Status = 2
	// Cluster failed
	Cluster_ERROR Cluster_Status = 3
	// Cluster is being updated.
	Cluster_UPDATING Cluster_Status = 4
	// Cluster is stopping.
	Cluster_STOPPING Cluster_Status = 5
	// Cluster stopped.
	Cluster_STOPPED Cluster_Status = 6
	// Cluster is starting.
	Cluster_STARTING Cluster_Status = 7
)

var Cluster_Status_name = map[int32]string{
	0: "STATUS_UNKNOWN",
	1: "CREATING",
	2: "RUNNING",
	3: "ERROR",
	4: "UPDATING",
	5: "STOPPING",
	6: "STOPPED",
	7: "STARTING",
}
var Cluster_Status_value = map[string]int32{
	"STATUS_UNKNOWN": 0,
	"CREATING":       1,
	"RUNNING":        2,
	"ERROR":          3,
	"UPDATING":       4,
	"STOPPING":       5,
	"STOPPED":        6,
	"STARTING":       7,
}

func (x Cluster_Status) String() string {
	return proto.EnumName(Cluster_Status_name, int32(x))
}
func (Cluster_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cluster_6d995a7e51dda886, []int{0, 2}
}

type Host_Role int32

const (
	// Role of the host in the cluster is unknown.
	Host_ROLE_UNKNOWN Host_Role = 0
	// Host is the master Redis server in the cluster.
	Host_MASTER Host_Role = 1
	// Host is a replica (standby) Redis server in the cluster.
	Host_REPLICA Host_Role = 2
)

var Host_Role_name = map[int32]string{
	0: "ROLE_UNKNOWN",
	1: "MASTER",
	2: "REPLICA",
}
var Host_Role_value = map[string]int32{
	"ROLE_UNKNOWN": 0,
	"MASTER":       1,
	"REPLICA":      2,
}

func (x Host_Role) String() string {
	return proto.EnumName(Host_Role_name, int32(x))
}
func (Host_Role) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cluster_6d995a7e51dda886, []int{3, 0}
}

type Host_Health int32

const (
	// Health of the host is unknown.
	Host_HEALTH_UNKNOWN Host_Health = 0
	// The host is performing all its functions normally.
	Host_ALIVE Host_Health = 1
	// The host is inoperable, and cannot perform any of its essential functions.
	Host_DEAD Host_Health = 2
	// The host is degraded, and can perform only some of its essential functions.
	Host_DEGRADED Host_Health = 3
)

var Host_Health_name = map[int32]string{
	0: "HEALTH_UNKNOWN",
	1: "ALIVE",
	2: "DEAD",
	3: "DEGRADED",
}
var Host_Health_value = map[string]int32{
	"HEALTH_UNKNOWN": 0,
	"ALIVE":          1,
	"DEAD":           2,
	"DEGRADED":       3,
}

func (x Host_Health) String() string {
	return proto.EnumName(Host_Health_name, int32(x))
}
func (Host_Health) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cluster_6d995a7e51dda886, []int{3, 1}
}

type Service_Type int32

const (
	Service_TYPE_UNSPECIFIED Service_Type = 0
	// The host is a Redis server.
	Service_REDIS Service_Type = 1
	// The host provides a Sentinel-only service (a quorum node).
	Service_ARBITER Service_Type = 2
)

var Service_Type_name = map[int32]string{
	0: "TYPE_UNSPECIFIED",
	1: "REDIS",
	2: "ARBITER",
}
var Service_Type_value = map[string]int32{
	"TYPE_UNSPECIFIED": 0,
	"REDIS":            1,
	"ARBITER":          2,
}

func (x Service_Type) String() string {
	return proto.EnumName(Service_Type_name, int32(x))
}
func (Service_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cluster_6d995a7e51dda886, []int{4, 0}
}

type Service_Health int32

const (
	// Health of the server is unknown.
	Service_HEALTH_UNKNOWN Service_Health = 0
	// The server is working normally.
	Service_ALIVE Service_Health = 1
	// The server is dead or unresponsive.
	Service_DEAD Service_Health = 2
)

var Service_Health_name = map[int32]string{
	0: "HEALTH_UNKNOWN",
	1: "ALIVE",
	2: "DEAD",
}
var Service_Health_value = map[string]int32{
	"HEALTH_UNKNOWN": 0,
	"ALIVE":          1,
	"DEAD":           2,
}

func (x Service_Health) String() string {
	return proto.EnumName(Service_Health_name, int32(x))
}
func (Service_Health) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cluster_6d995a7e51dda886, []int{4, 1}
}

// Description of a Redis cluster. For more information, see
// the Managed Service for Redis [documentation](/docs/managed-redis/concepts/).
type Cluster struct {
	// ID of the Redis cluster.
	// This ID is assigned by MDB at creation time.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the folder that the Redis cluster belongs to.
	FolderId string `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Creation timestamp in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Name of the Redis cluster.
	// The name is unique within the folder. 3-63 characters long.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the Redis cluster. 0-256 characters long.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Custom labels for the Redis cluster as `key:value` pairs.
	// Maximum 64 per cluster.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Deployment environment of the Redis cluster.
	Environment Cluster_Environment `protobuf:"varint,7,opt,name=environment,proto3,enum=yandex.cloud.mdb.redis.v1.Cluster_Environment" json:"environment,omitempty"`
	// Description of monitoring systems relevant to the Redis cluster.
	Monitoring []*Monitoring `protobuf:"bytes,8,rep,name=monitoring,proto3" json:"monitoring,omitempty"`
	// Configuration of the Redis cluster.
	Config    *ClusterConfig `protobuf:"bytes,9,opt,name=config,proto3" json:"config,omitempty"`
	NetworkId string         `protobuf:"bytes,10,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	// Aggregated cluster health
	Health Cluster_Health `protobuf:"varint,11,opt,name=health,proto3,enum=yandex.cloud.mdb.redis.v1.Cluster_Health" json:"health,omitempty"`
	// Cluster status
	Status               Cluster_Status `protobuf:"varint,12,opt,name=status,proto3,enum=yandex.cloud.mdb.redis.v1.Cluster_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Cluster) Reset()         { *m = Cluster{} }
func (m *Cluster) String() string { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()    {}
func (*Cluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_cluster_6d995a7e51dda886, []int{0}
}
func (m *Cluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cluster.Unmarshal(m, b)
}
func (m *Cluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cluster.Marshal(b, m, deterministic)
}
func (dst *Cluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cluster.Merge(dst, src)
}
func (m *Cluster) XXX_Size() int {
	return xxx_messageInfo_Cluster.Size(m)
}
func (m *Cluster) XXX_DiscardUnknown() {
	xxx_messageInfo_Cluster.DiscardUnknown(m)
}

var xxx_messageInfo_Cluster proto.InternalMessageInfo

func (m *Cluster) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Cluster) GetFolderId() string {
	if m != nil {
		return m.FolderId
	}
	return ""
}

func (m *Cluster) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Cluster) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Cluster) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Cluster) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Cluster) GetEnvironment() Cluster_Environment {
	if m != nil {
		return m.Environment
	}
	return Cluster_ENVIRONMENT_UNSPECIFIED
}

func (m *Cluster) GetMonitoring() []*Monitoring {
	if m != nil {
		return m.Monitoring
	}
	return nil
}

func (m *Cluster) GetConfig() *ClusterConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *Cluster) GetNetworkId() string {
	if m != nil {
		return m.NetworkId
	}
	return ""
}

func (m *Cluster) GetHealth() Cluster_Health {
	if m != nil {
		return m.Health
	}
	return Cluster_HEALTH_UNKNOWN
}

func (m *Cluster) GetStatus() Cluster_Status {
	if m != nil {
		return m.Status
	}
	return Cluster_STATUS_UNKNOWN
}

type Monitoring struct {
	// Name of the monitoring system.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the monitoring system.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Link to the monitoring system charts for the Redis cluster.
	Link                 string   `protobuf:"bytes,3,opt,name=link,proto3" json:"link,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Monitoring) Reset()         { *m = Monitoring{} }
func (m *Monitoring) String() string { return proto.CompactTextString(m) }
func (*Monitoring) ProtoMessage()    {}
func (*Monitoring) Descriptor() ([]byte, []int) {
	return fileDescriptor_cluster_6d995a7e51dda886, []int{1}
}
func (m *Monitoring) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Monitoring.Unmarshal(m, b)
}
func (m *Monitoring) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Monitoring.Marshal(b, m, deterministic)
}
func (dst *Monitoring) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Monitoring.Merge(dst, src)
}
func (m *Monitoring) XXX_Size() int {
	return xxx_messageInfo_Monitoring.Size(m)
}
func (m *Monitoring) XXX_DiscardUnknown() {
	xxx_messageInfo_Monitoring.DiscardUnknown(m)
}

var xxx_messageInfo_Monitoring proto.InternalMessageInfo

func (m *Monitoring) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Monitoring) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Monitoring) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

type ClusterConfig struct {
	// Version of Redis server software.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// Configuration for Redis servers in the cluster.
	//
	// Types that are valid to be assigned to RedisConfig:
	//	*ClusterConfig_RedisConfig_5_0
	RedisConfig isClusterConfig_RedisConfig `protobuf_oneof:"redis_config"`
	// Resources allocated to Redis hosts.
	Resources *Resources `protobuf:"bytes,3,opt,name=resources,proto3" json:"resources,omitempty"`
	// Start time for the daily backup in UTC timezone
	BackupWindowStart    *timeofday.TimeOfDay `protobuf:"bytes,4,opt,name=backup_window_start,json=backupWindowStart,proto3" json:"backup_window_start,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ClusterConfig) Reset()         { *m = ClusterConfig{} }
func (m *ClusterConfig) String() string { return proto.CompactTextString(m) }
func (*ClusterConfig) ProtoMessage()    {}
func (*ClusterConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_cluster_6d995a7e51dda886, []int{2}
}
func (m *ClusterConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterConfig.Unmarshal(m, b)
}
func (m *ClusterConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterConfig.Marshal(b, m, deterministic)
}
func (dst *ClusterConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterConfig.Merge(dst, src)
}
func (m *ClusterConfig) XXX_Size() int {
	return xxx_messageInfo_ClusterConfig.Size(m)
}
func (m *ClusterConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterConfig proto.InternalMessageInfo

func (m *ClusterConfig) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type isClusterConfig_RedisConfig interface {
	isClusterConfig_RedisConfig()
}

type ClusterConfig_RedisConfig_5_0 struct {
	RedisConfig_5_0 *config.RedisConfigSet5_0 `protobuf:"bytes,2,opt,name=redis_config_5_0,json=redisConfig50,proto3,oneof"`
}

func (*ClusterConfig_RedisConfig_5_0) isClusterConfig_RedisConfig() {}

func (m *ClusterConfig) GetRedisConfig() isClusterConfig_RedisConfig {
	if m != nil {
		return m.RedisConfig
	}
	return nil
}

func (m *ClusterConfig) GetRedisConfig_5_0() *config.RedisConfigSet5_0 {
	if x, ok := m.GetRedisConfig().(*ClusterConfig_RedisConfig_5_0); ok {
		return x.RedisConfig_5_0
	}
	return nil
}

func (m *ClusterConfig) GetResources() *Resources {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *ClusterConfig) GetBackupWindowStart() *timeofday.TimeOfDay {
	if m != nil {
		return m.BackupWindowStart
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ClusterConfig) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ClusterConfig_OneofMarshaler, _ClusterConfig_OneofUnmarshaler, _ClusterConfig_OneofSizer, []interface{}{
		(*ClusterConfig_RedisConfig_5_0)(nil),
	}
}

func _ClusterConfig_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ClusterConfig)
	// redis_config
	switch x := m.RedisConfig.(type) {
	case *ClusterConfig_RedisConfig_5_0:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RedisConfig_5_0); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ClusterConfig.RedisConfig has unexpected type %T", x)
	}
	return nil
}

func _ClusterConfig_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ClusterConfig)
	switch tag {
	case 2: // redis_config.redis_config_5_0
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(config.RedisConfigSet5_0)
		err := b.DecodeMessage(msg)
		m.RedisConfig = &ClusterConfig_RedisConfig_5_0{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ClusterConfig_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ClusterConfig)
	// redis_config
	switch x := m.RedisConfig.(type) {
	case *ClusterConfig_RedisConfig_5_0:
		s := proto.Size(x.RedisConfig_5_0)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Host struct {
	// Name of the Redis host. The host name is assigned by MDB at creation time, and cannot be changed.
	// 1-63 characters long.
	//
	// The name is unique across all existing MDB hosts in Yandex.Cloud, as it defines the FQDN of the host.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// ID of the Redis host. The ID is assigned by MDB at creation time.
	ClusterId string `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// ID of the availability zone where the Redis host resides.
	ZoneId string `protobuf:"bytes,3,opt,name=zone_id,json=zoneId,proto3" json:"zone_id,omitempty"`
	// ID of the subnet that the host belongs to.
	SubnetId string `protobuf:"bytes,4,opt,name=subnet_id,json=subnetId,proto3" json:"subnet_id,omitempty"`
	// Resources allocated to the Redis host.
	Resources *Resources `protobuf:"bytes,5,opt,name=resources,proto3" json:"resources,omitempty"`
	// Role of the host in the cluster.
	Role Host_Role `protobuf:"varint,6,opt,name=role,proto3,enum=yandex.cloud.mdb.redis.v1.Host_Role" json:"role,omitempty"`
	// Status code of the aggregated health of the host.
	Health Host_Health `protobuf:"varint,7,opt,name=health,proto3,enum=yandex.cloud.mdb.redis.v1.Host_Health" json:"health,omitempty"`
	// Services provided by the host.
	Services             []*Service `protobuf:"bytes,8,rep,name=services,proto3" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Host) Reset()         { *m = Host{} }
func (m *Host) String() string { return proto.CompactTextString(m) }
func (*Host) ProtoMessage()    {}
func (*Host) Descriptor() ([]byte, []int) {
	return fileDescriptor_cluster_6d995a7e51dda886, []int{3}
}
func (m *Host) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Host.Unmarshal(m, b)
}
func (m *Host) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Host.Marshal(b, m, deterministic)
}
func (dst *Host) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Host.Merge(dst, src)
}
func (m *Host) XXX_Size() int {
	return xxx_messageInfo_Host.Size(m)
}
func (m *Host) XXX_DiscardUnknown() {
	xxx_messageInfo_Host.DiscardUnknown(m)
}

var xxx_messageInfo_Host proto.InternalMessageInfo

func (m *Host) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Host) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *Host) GetZoneId() string {
	if m != nil {
		return m.ZoneId
	}
	return ""
}

func (m *Host) GetSubnetId() string {
	if m != nil {
		return m.SubnetId
	}
	return ""
}

func (m *Host) GetResources() *Resources {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *Host) GetRole() Host_Role {
	if m != nil {
		return m.Role
	}
	return Host_ROLE_UNKNOWN
}

func (m *Host) GetHealth() Host_Health {
	if m != nil {
		return m.Health
	}
	return Host_HEALTH_UNKNOWN
}

func (m *Host) GetServices() []*Service {
	if m != nil {
		return m.Services
	}
	return nil
}

type Service struct {
	// Type of the service provided by the host.
	Type Service_Type `protobuf:"varint,1,opt,name=type,proto3,enum=yandex.cloud.mdb.redis.v1.Service_Type" json:"type,omitempty"`
	// Status code of server availability.
	Health               Service_Health `protobuf:"varint,2,opt,name=health,proto3,enum=yandex.cloud.mdb.redis.v1.Service_Health" json:"health,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_cluster_6d995a7e51dda886, []int{4}
}
func (m *Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service.Unmarshal(m, b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service.Marshal(b, m, deterministic)
}
func (dst *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(dst, src)
}
func (m *Service) XXX_Size() int {
	return xxx_messageInfo_Service.Size(m)
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

func (m *Service) GetType() Service_Type {
	if m != nil {
		return m.Type
	}
	return Service_TYPE_UNSPECIFIED
}

func (m *Service) GetHealth() Service_Health {
	if m != nil {
		return m.Health
	}
	return Service_HEALTH_UNKNOWN
}

type Resources struct {
	// ID of the preset for computational resources available to a host (CPU, memory etc.).
	// All available presets are listed in the [documentation](/docs/managed-redis/concepts/instance-types).
	ResourcePresetId string `protobuf:"bytes,1,opt,name=resource_preset_id,json=resourcePresetId,proto3" json:"resource_preset_id,omitempty"`
	// Volume of the storage available to a host, in bytes.
	DiskSize             int64    `protobuf:"varint,2,opt,name=disk_size,json=diskSize,proto3" json:"disk_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resources) Reset()         { *m = Resources{} }
func (m *Resources) String() string { return proto.CompactTextString(m) }
func (*Resources) ProtoMessage()    {}
func (*Resources) Descriptor() ([]byte, []int) {
	return fileDescriptor_cluster_6d995a7e51dda886, []int{5}
}
func (m *Resources) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resources.Unmarshal(m, b)
}
func (m *Resources) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resources.Marshal(b, m, deterministic)
}
func (dst *Resources) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resources.Merge(dst, src)
}
func (m *Resources) XXX_Size() int {
	return xxx_messageInfo_Resources.Size(m)
}
func (m *Resources) XXX_DiscardUnknown() {
	xxx_messageInfo_Resources.DiscardUnknown(m)
}

var xxx_messageInfo_Resources proto.InternalMessageInfo

func (m *Resources) GetResourcePresetId() string {
	if m != nil {
		return m.ResourcePresetId
	}
	return ""
}

func (m *Resources) GetDiskSize() int64 {
	if m != nil {
		return m.DiskSize
	}
	return 0
}

func init() {
	proto.RegisterType((*Cluster)(nil), "yandex.cloud.mdb.redis.v1.Cluster")
	proto.RegisterMapType((map[string]string)(nil), "yandex.cloud.mdb.redis.v1.Cluster.LabelsEntry")
	proto.RegisterType((*Monitoring)(nil), "yandex.cloud.mdb.redis.v1.Monitoring")
	proto.RegisterType((*ClusterConfig)(nil), "yandex.cloud.mdb.redis.v1.ClusterConfig")
	proto.RegisterType((*Host)(nil), "yandex.cloud.mdb.redis.v1.Host")
	proto.RegisterType((*Service)(nil), "yandex.cloud.mdb.redis.v1.Service")
	proto.RegisterType((*Resources)(nil), "yandex.cloud.mdb.redis.v1.Resources")
	proto.RegisterEnum("yandex.cloud.mdb.redis.v1.Cluster_Environment", Cluster_Environment_name, Cluster_Environment_value)
	proto.RegisterEnum("yandex.cloud.mdb.redis.v1.Cluster_Health", Cluster_Health_name, Cluster_Health_value)
	proto.RegisterEnum("yandex.cloud.mdb.redis.v1.Cluster_Status", Cluster_Status_name, Cluster_Status_value)
	proto.RegisterEnum("yandex.cloud.mdb.redis.v1.Host_Role", Host_Role_name, Host_Role_value)
	proto.RegisterEnum("yandex.cloud.mdb.redis.v1.Host_Health", Host_Health_name, Host_Health_value)
	proto.RegisterEnum("yandex.cloud.mdb.redis.v1.Service_Type", Service_Type_name, Service_Type_value)
	proto.RegisterEnum("yandex.cloud.mdb.redis.v1.Service_Health", Service_Health_name, Service_Health_value)
}

func init() {
	proto.RegisterFile("yandex/cloud/mdb/redis/v1/cluster.proto", fileDescriptor_cluster_6d995a7e51dda886)
}

var fileDescriptor_cluster_6d995a7e51dda886 = []byte{
	// 1063 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xff, 0x6e, 0xe3, 0xc4,
	0x13, 0xbf, 0xfc, 0x8e, 0x27, 0x6d, 0xe4, 0xef, 0x7e, 0x4f, 0x9c, 0x69, 0x75, 0xa2, 0xb2, 0x80,
	0x2b, 0x12, 0xe7, 0xb4, 0x3d, 0x2a, 0xdd, 0x71, 0xd2, 0x81, 0x9b, 0x6c, 0xaf, 0x16, 0xa9, 0x13,
	0xad, 0x9d, 0x9e, 0xe0, 0x0f, 0x2c, 0x27, 0xde, 0xa6, 0x56, 0x1d, 0x3b, 0xb2, 0x37, 0x29, 0xe9,
	0xfb, 0xf0, 0x00, 0x3c, 0x12, 0xaf, 0xc0, 0x13, 0xa0, 0xdd, 0x75, 0xda, 0xf4, 0xa0, 0x69, 0x41,
	0xfc, 0xe7, 0x99, 0xf9, 0x7c, 0xc6, 0xb3, 0xb3, 0x33, 0x1f, 0x2d, 0xbc, 0x58, 0xf8, 0x71, 0x40,
	0x7f, 0x69, 0x8d, 0xa2, 0x64, 0x16, 0xb4, 0x26, 0xc1, 0xb0, 0x95, 0xd2, 0x20, 0xcc, 0x5a, 0xf3,
	0xfd, 0xd6, 0x28, 0x9a, 0x65, 0x8c, 0xa6, 0xc6, 0x34, 0x4d, 0x58, 0x82, 0x3e, 0x95, 0x40, 0x43,
	0x00, 0x8d, 0x49, 0x30, 0x34, 0x04, 0xd0, 0x98, 0xef, 0x6f, 0x7d, 0x36, 0x4e, 0x92, 0x71, 0x44,
	0x5b, 0x02, 0x38, 0x9c, 0x9d, 0xb7, 0x58, 0x38, 0xa1, 0x19, 0xf3, 0x27, 0x53, 0xc9, 0xdd, 0xda,
	0xce, 0x01, 0x6c, 0x31, 0xa5, 0x22, 0x98, 0x9c, 0x07, 0xfe, 0x22, 0x0f, 0xb6, 0xd6, 0x54, 0x90,
	0xc4, 0xe7, 0xe1, 0x58, 0xda, 0x87, 0xde, 0x9e, 0x24, 0xe8, 0xbf, 0xd5, 0xa0, 0xd6, 0x96, 0xb5,
	0xa1, 0x26, 0x14, 0xc3, 0x40, 0x2b, 0xec, 0x14, 0x76, 0x15, 0x52, 0x0c, 0x03, 0xb4, 0x0d, 0xca,
	0x79, 0x12, 0x05, 0x34, 0xf5, 0xc2, 0x40, 0x2b, 0x0a, 0x77, 0x5d, 0x3a, 0xac, 0x00, 0xbd, 0x01,
	0x18, 0xa5, 0xd4, 0x67, 0x34, 0xf0, 0x7c, 0xa6, 0x95, 0x76, 0x0a, 0xbb, 0x8d, 0x83, 0x2d, 0x43,
	0xd6, 0x66, 0x2c, 0x8b, 0x37, 0xdc, 0x65, 0xf1, 0x44, 0xc9, 0xd1, 0x26, 0x43, 0x08, 0xca, 0xb1,
	0x3f, 0xa1, 0x5a, 0x59, 0xa4, 0x14, 0xdf, 0x68, 0x07, 0x1a, 0x01, 0xcd, 0x46, 0x69, 0x38, 0x65,
	0x61, 0x12, 0x6b, 0x15, 0x11, 0x5a, 0x75, 0xa1, 0x63, 0xa8, 0x46, 0xfe, 0x90, 0x46, 0x99, 0x56,
	0xdd, 0x29, 0xed, 0x36, 0x0e, 0x0c, 0xe3, 0xde, 0x26, 0x1a, 0xf9, 0x89, 0x8c, 0xae, 0x20, 0xe0,
	0x98, 0xa5, 0x0b, 0x92, 0xb3, 0x51, 0x1f, 0x1a, 0x34, 0x9e, 0x87, 0x69, 0x12, 0x4f, 0x68, 0xcc,
	0xb4, 0xda, 0x4e, 0x61, 0xb7, 0xf9, 0xa8, 0x64, 0xf8, 0x96, 0x45, 0x56, 0x53, 0x20, 0x0c, 0x30,
	0x49, 0xe2, 0x90, 0x25, 0x69, 0x18, 0x8f, 0xb5, 0xba, 0xa8, 0xee, 0x8b, 0x35, 0x09, 0x4f, 0x6f,
	0xc0, 0x64, 0x85, 0x88, 0xbe, 0x87, 0xaa, 0xbc, 0x23, 0x4d, 0x11, 0xdd, 0xdc, 0x7d, 0xb8, 0xa6,
	0xb6, 0xc0, 0x93, 0x9c, 0x87, 0x9e, 0x03, 0xc4, 0x94, 0x5d, 0x25, 0xe9, 0x25, 0xbf, 0x31, 0x10,
	0x3d, 0x54, 0x72, 0x8f, 0x15, 0x20, 0x13, 0xaa, 0x17, 0xd4, 0x8f, 0xd8, 0x85, 0xd6, 0x10, 0x87,
	0xfe, 0xea, 0x11, 0x87, 0x3e, 0x11, 0x04, 0x92, 0x13, 0x79, 0x8a, 0x8c, 0xf9, 0x6c, 0x96, 0x69,
	0x1b, 0x8f, 0x4e, 0xe1, 0x08, 0x02, 0xc9, 0x89, 0x5b, 0x6f, 0xa0, 0xb1, 0x72, 0x2d, 0x48, 0x85,
	0xd2, 0x25, 0x5d, 0xe4, 0x53, 0xc7, 0x3f, 0xd1, 0x53, 0xa8, 0xcc, 0xfd, 0x68, 0x46, 0xf3, 0x91,
	0x93, 0xc6, 0xb7, 0xc5, 0xd7, 0x05, 0xdd, 0x82, 0xc6, 0xca, 0x25, 0xa0, 0x6d, 0x78, 0x86, 0xed,
	0x33, 0x8b, 0xf4, 0xec, 0x53, 0x6c, 0xbb, 0xde, 0xc0, 0x76, 0xfa, 0xb8, 0x6d, 0x1d, 0x5b, 0xb8,
	0xa3, 0x3e, 0x41, 0x4d, 0x80, 0x3e, 0xe9, 0x75, 0x06, 0x6d, 0xd7, 0xea, 0xd9, 0x6a, 0x01, 0x6d,
	0x82, 0xd2, 0x27, 0xd8, 0x71, 0xcd, 0xa3, 0x2e, 0x56, 0x8b, 0xfa, 0x77, 0x50, 0x95, 0x47, 0x43,
	0x08, 0x9a, 0x27, 0xd8, 0xec, 0xba, 0x27, 0xde, 0xc0, 0xfe, 0xc1, 0xee, 0x7d, 0xb0, 0xd5, 0x27,
	0x48, 0x81, 0x8a, 0xd9, 0xb5, 0xce, 0xb0, 0x5a, 0x40, 0x75, 0x28, 0x77, 0xb0, 0xd9, 0x51, 0x8b,
	0x68, 0x03, 0xea, 0x1d, 0xfc, 0x9e, 0x98, 0x1d, 0xdc, 0x51, 0x4b, 0xfa, 0x02, 0xaa, 0xf2, 0x60,
	0x3c, 0x81, 0xe3, 0x9a, 0xee, 0xc0, 0x59, 0x49, 0xb0, 0x01, 0xf5, 0x36, 0xc1, 0xa6, 0x6b, 0xd9,
	0xef, 0xd5, 0x02, 0x6a, 0x40, 0x8d, 0x0c, 0x6c, 0x9b, 0x1b, 0x45, 0x9e, 0x1b, 0x13, 0xd2, 0x23,
	0x6a, 0x89, 0xa3, 0x06, 0xfd, 0x8e, 0x44, 0x95, 0xb9, 0xe5, 0xb8, 0xbd, 0x7e, 0x9f, 0x5b, 0x15,
	0xce, 0x11, 0x16, 0xee, 0xa8, 0x55, 0x19, 0x32, 0x89, 0x00, 0xd6, 0xf4, 0x33, 0x80, 0xdb, 0x11,
	0xba, 0xd9, 0xa6, 0xc2, 0xfd, 0xdb, 0x54, 0xfc, 0xeb, 0x36, 0x21, 0x28, 0x47, 0x61, 0x7c, 0x29,
	0x16, 0x57, 0x21, 0xe2, 0x5b, 0xff, 0xb5, 0x08, 0x9b, 0x77, 0x06, 0x0b, 0x69, 0x50, 0x9b, 0xd3,
	0x34, 0xe3, 0x39, 0x64, 0xfa, 0xa5, 0x89, 0x7e, 0x06, 0x55, 0x5c, 0xb4, 0x27, 0x47, 0xcf, 0x3b,
	0xf4, 0xf6, 0xc4, 0x6f, 0x1a, 0x07, 0xaf, 0xd6, 0x8c, 0x84, 0x04, 0x1b, 0x84, 0xdb, 0xf2, 0x17,
	0x0e, 0x65, 0x87, 0xde, 0xde, 0xc9, 0x13, 0xb2, 0x99, 0xde, 0x3a, 0x0f, 0xf7, 0xd0, 0x11, 0x28,
	0x29, 0xcd, 0x92, 0x59, 0x3a, 0xa2, 0x59, 0xae, 0x2e, 0x9f, 0xaf, 0x49, 0x4c, 0x96, 0x58, 0x72,
	0x4b, 0x43, 0xc7, 0xf0, 0xff, 0xa1, 0x3f, 0xba, 0x9c, 0x4d, 0xbd, 0xab, 0x30, 0x0e, 0x92, 0x2b,
	0x2f, 0x63, 0x7e, 0xca, 0x84, 0xec, 0x34, 0x0e, 0x3e, 0x59, 0x6a, 0x15, 0xd7, 0x51, 0xa1, 0x53,
	0xbd, 0xf3, 0x8e, 0xbf, 0x20, 0xff, 0x93, 0x94, 0x0f, 0x82, 0xe1, 0x70, 0xc2, 0x51, 0x13, 0x36,
	0x56, 0xcf, 0xaa, 0xff, 0x5e, 0x82, 0xf2, 0x49, 0x92, 0xb1, 0xbf, 0x6d, 0xfd, 0x73, 0x80, 0x5c,
	0xeb, 0x6f, 0x55, 0x53, 0xc9, 0x3d, 0x56, 0x80, 0x9e, 0x41, 0xed, 0x3a, 0x89, 0x29, 0x8f, 0xc9,
	0xd6, 0x57, 0xb9, 0x69, 0x09, 0xb1, 0xcd, 0x66, 0xc3, 0x98, 0x32, 0x1e, 0x92, 0xca, 0x58, 0x97,
	0x0e, 0x2b, 0xb8, 0xdb, 0x8d, 0xca, 0xbf, 0xeb, 0xc6, 0x6b, 0x28, 0xa7, 0x49, 0x44, 0xb5, 0xaa,
	0x58, 0xdc, 0x75, 0x74, 0x7e, 0x36, 0x83, 0x24, 0x11, 0x25, 0x82, 0x81, 0xde, 0xdd, 0xe8, 0x86,
	0x14, 0xcb, 0x2f, 0x1f, 0xe2, 0x7e, 0x24, 0x1a, 0xef, 0xa0, 0x9e, 0xd1, 0x74, 0x1e, 0xf2, 0xe2,
	0xa5, 0x3a, 0xea, 0x6b, 0x32, 0x38, 0x12, 0x4a, 0x6e, 0x38, 0xfa, 0x3e, 0x94, 0x79, 0x35, 0x48,
	0x85, 0x0d, 0xd2, 0xeb, 0xe2, 0x95, 0x35, 0x03, 0xa8, 0x9e, 0x9a, 0x8e, 0x8b, 0x49, 0xbe, 0x64,
	0xb8, 0xdf, 0xb5, 0xda, 0xe6, 0x7f, 0xb1, 0xde, 0x7f, 0x14, 0xa0, 0x96, 0x57, 0x82, 0xde, 0x42,
	0x99, 0x0f, 0x89, 0xb8, 0xe6, 0xe6, 0xc1, 0x8b, 0x87, 0x6b, 0x37, 0xdc, 0xc5, 0x94, 0x12, 0x41,
	0x5a, 0x11, 0xdd, 0xe2, 0x83, 0x8a, 0xb9, 0xa4, 0xdf, 0xed, 0x9f, 0xfe, 0x0d, 0x94, 0x79, 0x42,
	0xf4, 0x14, 0x54, 0xf7, 0xc7, 0x3e, 0xfe, 0x48, 0xe8, 0x14, 0xa8, 0x10, 0xdc, 0xb1, 0x1c, 0xd9,
	0x02, 0x93, 0x1c, 0x59, 0xbc, 0x1f, 0x45, 0x7d, 0xff, 0x1f, 0xb7, 0x40, 0x3f, 0x03, 0xe5, 0x66,
	0x74, 0xd0, 0xd7, 0x80, 0x96, 0xc3, 0xe3, 0x4d, 0x53, 0x9a, 0xc9, 0xc9, 0x94, 0xa3, 0xae, 0x2e,
	0x23, 0x7d, 0x11, 0x90, 0xe3, 0x1b, 0x84, 0xd9, 0xa5, 0x97, 0x85, 0xd7, 0x52, 0xb8, 0x4b, 0xa4,
	0xce, 0x1d, 0x4e, 0x78, 0x4d, 0x8f, 0xf0, 0x4f, 0xed, 0x71, 0xc8, 0x2e, 0x66, 0x43, 0x63, 0x94,
	0x4c, 0xf2, 0x27, 0xca, 0x4b, 0xf9, 0x44, 0x19, 0x27, 0x2f, 0xc7, 0x34, 0x16, 0xef, 0x85, 0xfb,
	0xdf, 0x2e, 0x6f, 0xc5, 0xc7, 0xb0, 0x2a, 0x60, 0xaf, 0xfe, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x30,
	0x85, 0x6d, 0x7c, 0x67, 0x09, 0x00, 0x00,
}