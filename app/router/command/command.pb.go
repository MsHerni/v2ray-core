package command

import (
	net "github.com/v2fly/v2ray-core/v5/common/net"
	_ "github.com/v2fly/v2ray-core/v5/common/protoext"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// RoutingContext is the context with information relative to routing process.
// It conforms to the structure of v2ray.core.features.routing.Context and
// v2ray.core.features.routing.Route.
type RoutingContext struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	InboundTag        string                 `protobuf:"bytes,1,opt,name=InboundTag,proto3" json:"InboundTag,omitempty"`
	Network           net.Network            `protobuf:"varint,2,opt,name=Network,proto3,enum=v2ray.core.common.net.Network" json:"Network,omitempty"`
	SourceIPs         [][]byte               `protobuf:"bytes,3,rep,name=SourceIPs,proto3" json:"SourceIPs,omitempty"`
	TargetIPs         [][]byte               `protobuf:"bytes,4,rep,name=TargetIPs,proto3" json:"TargetIPs,omitempty"`
	SourcePort        uint32                 `protobuf:"varint,5,opt,name=SourcePort,proto3" json:"SourcePort,omitempty"`
	TargetPort        uint32                 `protobuf:"varint,6,opt,name=TargetPort,proto3" json:"TargetPort,omitempty"`
	TargetDomain      string                 `protobuf:"bytes,7,opt,name=TargetDomain,proto3" json:"TargetDomain,omitempty"`
	Protocol          string                 `protobuf:"bytes,8,opt,name=Protocol,proto3" json:"Protocol,omitempty"`
	User              string                 `protobuf:"bytes,9,opt,name=User,proto3" json:"User,omitempty"`
	Level             uint32                 `protobuf:"varint,13,opt,name=Level,proto3" json:"Level,omitempty"`
	Attributes        map[string]string      `protobuf:"bytes,10,rep,name=Attributes,proto3" json:"Attributes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	OutboundGroupTags []string               `protobuf:"bytes,11,rep,name=OutboundGroupTags,proto3" json:"OutboundGroupTags,omitempty"`
	OutboundTag       string                 `protobuf:"bytes,12,opt,name=OutboundTag,proto3" json:"OutboundTag,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *RoutingContext) Reset() {
	*x = RoutingContext{}
	mi := &file_app_router_command_command_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoutingContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoutingContext) ProtoMessage() {}

func (x *RoutingContext) ProtoReflect() protoreflect.Message {
	mi := &file_app_router_command_command_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoutingContext.ProtoReflect.Descriptor instead.
func (*RoutingContext) Descriptor() ([]byte, []int) {
	return file_app_router_command_command_proto_rawDescGZIP(), []int{0}
}

func (x *RoutingContext) GetInboundTag() string {
	if x != nil {
		return x.InboundTag
	}
	return ""
}

func (x *RoutingContext) GetNetwork() net.Network {
	if x != nil {
		return x.Network
	}
	return net.Network(0)
}

func (x *RoutingContext) GetSourceIPs() [][]byte {
	if x != nil {
		return x.SourceIPs
	}
	return nil
}

func (x *RoutingContext) GetTargetIPs() [][]byte {
	if x != nil {
		return x.TargetIPs
	}
	return nil
}

func (x *RoutingContext) GetSourcePort() uint32 {
	if x != nil {
		return x.SourcePort
	}
	return 0
}

func (x *RoutingContext) GetTargetPort() uint32 {
	if x != nil {
		return x.TargetPort
	}
	return 0
}

func (x *RoutingContext) GetTargetDomain() string {
	if x != nil {
		return x.TargetDomain
	}
	return ""
}

func (x *RoutingContext) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *RoutingContext) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *RoutingContext) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *RoutingContext) GetAttributes() map[string]string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *RoutingContext) GetOutboundGroupTags() []string {
	if x != nil {
		return x.OutboundGroupTags
	}
	return nil
}

func (x *RoutingContext) GetOutboundTag() string {
	if x != nil {
		return x.OutboundTag
	}
	return ""
}

// SubscribeRoutingStatsRequest subscribes to routing statistics channel if
// opened by v2ray-core.
// * FieldSelectors selects a subset of fields in routing statistics to return.
// Valid selectors:
//   - inbound: Selects connection's inbound tag.
//   - network: Selects connection's network.
//   - ip: Equivalent as "ip_source" and "ip_target", selects both source and
//     target IP.
//   - port: Equivalent as "port_source" and "port_target", selects both source
//     and target port.
//   - domain: Selects target domain.
//   - protocol: Select connection's protocol.
//   - user: Select connection's inbound user email.
//   - attributes: Select connection's additional attributes.
//   - outbound: Equivalent as "outbound" and "outbound_group", select both
//     outbound tag and outbound group tags.
//
// * If FieldSelectors is left empty, all fields will be returned.
type SubscribeRoutingStatsRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	FieldSelectors []string               `protobuf:"bytes,1,rep,name=FieldSelectors,proto3" json:"FieldSelectors,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *SubscribeRoutingStatsRequest) Reset() {
	*x = SubscribeRoutingStatsRequest{}
	mi := &file_app_router_command_command_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubscribeRoutingStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRoutingStatsRequest) ProtoMessage() {}

func (x *SubscribeRoutingStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_router_command_command_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRoutingStatsRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRoutingStatsRequest) Descriptor() ([]byte, []int) {
	return file_app_router_command_command_proto_rawDescGZIP(), []int{1}
}

func (x *SubscribeRoutingStatsRequest) GetFieldSelectors() []string {
	if x != nil {
		return x.FieldSelectors
	}
	return nil
}

// TestRouteRequest manually tests a routing result according to the routing
// context message.
// * RoutingContext is the routing message without outbound information.
// * FieldSelectors selects the fields to return in the routing result. All
// fields are returned if left empty.
// * PublishResult broadcasts the routing result to routing statistics channel
// if set true.
type TestRouteRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	RoutingContext *RoutingContext        `protobuf:"bytes,1,opt,name=RoutingContext,proto3" json:"RoutingContext,omitempty"`
	FieldSelectors []string               `protobuf:"bytes,2,rep,name=FieldSelectors,proto3" json:"FieldSelectors,omitempty"`
	PublishResult  bool                   `protobuf:"varint,3,opt,name=PublishResult,proto3" json:"PublishResult,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *TestRouteRequest) Reset() {
	*x = TestRouteRequest{}
	mi := &file_app_router_command_command_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TestRouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestRouteRequest) ProtoMessage() {}

func (x *TestRouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_router_command_command_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestRouteRequest.ProtoReflect.Descriptor instead.
func (*TestRouteRequest) Descriptor() ([]byte, []int) {
	return file_app_router_command_command_proto_rawDescGZIP(), []int{2}
}

func (x *TestRouteRequest) GetRoutingContext() *RoutingContext {
	if x != nil {
		return x.RoutingContext
	}
	return nil
}

func (x *TestRouteRequest) GetFieldSelectors() []string {
	if x != nil {
		return x.FieldSelectors
	}
	return nil
}

func (x *TestRouteRequest) GetPublishResult() bool {
	if x != nil {
		return x.PublishResult
	}
	return false
}

type PrincipleTargetInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tag           []string               `protobuf:"bytes,1,rep,name=tag,proto3" json:"tag,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PrincipleTargetInfo) Reset() {
	*x = PrincipleTargetInfo{}
	mi := &file_app_router_command_command_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrincipleTargetInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrincipleTargetInfo) ProtoMessage() {}

func (x *PrincipleTargetInfo) ProtoReflect() protoreflect.Message {
	mi := &file_app_router_command_command_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrincipleTargetInfo.ProtoReflect.Descriptor instead.
func (*PrincipleTargetInfo) Descriptor() ([]byte, []int) {
	return file_app_router_command_command_proto_rawDescGZIP(), []int{3}
}

func (x *PrincipleTargetInfo) GetTag() []string {
	if x != nil {
		return x.Tag
	}
	return nil
}

type OverrideInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Target        string                 `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OverrideInfo) Reset() {
	*x = OverrideInfo{}
	mi := &file_app_router_command_command_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OverrideInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OverrideInfo) ProtoMessage() {}

func (x *OverrideInfo) ProtoReflect() protoreflect.Message {
	mi := &file_app_router_command_command_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OverrideInfo.ProtoReflect.Descriptor instead.
func (*OverrideInfo) Descriptor() ([]byte, []int) {
	return file_app_router_command_command_proto_rawDescGZIP(), []int{4}
}

func (x *OverrideInfo) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

type BalancerMsg struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Override        *OverrideInfo          `protobuf:"bytes,5,opt,name=override,proto3" json:"override,omitempty"`
	PrincipleTarget *PrincipleTargetInfo   `protobuf:"bytes,6,opt,name=principle_target,json=principleTarget,proto3" json:"principle_target,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *BalancerMsg) Reset() {
	*x = BalancerMsg{}
	mi := &file_app_router_command_command_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BalancerMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalancerMsg) ProtoMessage() {}

func (x *BalancerMsg) ProtoReflect() protoreflect.Message {
	mi := &file_app_router_command_command_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalancerMsg.ProtoReflect.Descriptor instead.
func (*BalancerMsg) Descriptor() ([]byte, []int) {
	return file_app_router_command_command_proto_rawDescGZIP(), []int{5}
}

func (x *BalancerMsg) GetOverride() *OverrideInfo {
	if x != nil {
		return x.Override
	}
	return nil
}

func (x *BalancerMsg) GetPrincipleTarget() *PrincipleTargetInfo {
	if x != nil {
		return x.PrincipleTarget
	}
	return nil
}

type GetBalancerInfoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tag           string                 `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBalancerInfoRequest) Reset() {
	*x = GetBalancerInfoRequest{}
	mi := &file_app_router_command_command_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBalancerInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalancerInfoRequest) ProtoMessage() {}

func (x *GetBalancerInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_router_command_command_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalancerInfoRequest.ProtoReflect.Descriptor instead.
func (*GetBalancerInfoRequest) Descriptor() ([]byte, []int) {
	return file_app_router_command_command_proto_rawDescGZIP(), []int{6}
}

func (x *GetBalancerInfoRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

type GetBalancerInfoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Balancer      *BalancerMsg           `protobuf:"bytes,1,opt,name=balancer,proto3" json:"balancer,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBalancerInfoResponse) Reset() {
	*x = GetBalancerInfoResponse{}
	mi := &file_app_router_command_command_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBalancerInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalancerInfoResponse) ProtoMessage() {}

func (x *GetBalancerInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_router_command_command_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalancerInfoResponse.ProtoReflect.Descriptor instead.
func (*GetBalancerInfoResponse) Descriptor() ([]byte, []int) {
	return file_app_router_command_command_proto_rawDescGZIP(), []int{7}
}

func (x *GetBalancerInfoResponse) GetBalancer() *BalancerMsg {
	if x != nil {
		return x.Balancer
	}
	return nil
}

type OverrideBalancerTargetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BalancerTag   string                 `protobuf:"bytes,1,opt,name=balancerTag,proto3" json:"balancerTag,omitempty"`
	Target        string                 `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OverrideBalancerTargetRequest) Reset() {
	*x = OverrideBalancerTargetRequest{}
	mi := &file_app_router_command_command_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OverrideBalancerTargetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OverrideBalancerTargetRequest) ProtoMessage() {}

func (x *OverrideBalancerTargetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_router_command_command_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OverrideBalancerTargetRequest.ProtoReflect.Descriptor instead.
func (*OverrideBalancerTargetRequest) Descriptor() ([]byte, []int) {
	return file_app_router_command_command_proto_rawDescGZIP(), []int{8}
}

func (x *OverrideBalancerTargetRequest) GetBalancerTag() string {
	if x != nil {
		return x.BalancerTag
	}
	return ""
}

func (x *OverrideBalancerTargetRequest) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

type OverrideBalancerTargetResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OverrideBalancerTargetResponse) Reset() {
	*x = OverrideBalancerTargetResponse{}
	mi := &file_app_router_command_command_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OverrideBalancerTargetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OverrideBalancerTargetResponse) ProtoMessage() {}

func (x *OverrideBalancerTargetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_router_command_command_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OverrideBalancerTargetResponse.ProtoReflect.Descriptor instead.
func (*OverrideBalancerTargetResponse) Descriptor() ([]byte, []int) {
	return file_app_router_command_command_proto_rawDescGZIP(), []int{9}
}

type Config struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Config) Reset() {
	*x = Config{}
	mi := &file_app_router_command_command_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_app_router_command_command_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_app_router_command_command_proto_rawDescGZIP(), []int{10}
}

var File_app_router_command_command_proto protoreflect.FileDescriptor

var file_app_router_command_command_proto_rawDesc = string([]byte{
	0x0a, 0x20, 0x61, 0x70, 0x70, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1d, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61,
	0x70, 0x70, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x1a, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x65,
	0x78, 0x74, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6e, 0x65, 0x74, 0x2f,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x04,
	0x0a, 0x0e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x12, 0x1e, 0x0a, 0x0a, 0x49, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x54, 0x61, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x49, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x54, 0x61, 0x67,
	0x12, 0x38, 0x0a, 0x07, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1e, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x52, 0x07, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x49, 0x50, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x09, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x50, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x49, 0x50, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x09, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x49, 0x50, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x50, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x50, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x5d, 0x0a, 0x0a, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x0a,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x0a, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12,
	0x2c, 0x0a, 0x11, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x54, 0x61, 0x67, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x4f, 0x75, 0x74, 0x62,
	0x6f, 0x75, 0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x61, 0x67, 0x73, 0x12, 0x20, 0x0a,
	0x0b, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x54, 0x61, 0x67, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x54, 0x61, 0x67, 0x1a,
	0x3d, 0x0a, 0x0f, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x46,
	0x0a, 0x1c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x69,
	0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26,
	0x0a, 0x0e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x22, 0xb7, 0x01, 0x0a, 0x10, 0x54, 0x65, 0x73, 0x74, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x55, 0x0a, 0x0e, 0x52,
	0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x61, 0x70, 0x70, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x52, 0x0e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0d, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x27, 0x0a, 0x13, 0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x6c, 0x65, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x22, 0x26, 0x0a, 0x0c, 0x4f, 0x76, 0x65,
	0x72, 0x72, 0x69, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x22, 0xb5, 0x01, 0x0a, 0x0b, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x4d, 0x73,
	0x67, 0x12, 0x47, 0x0a, 0x08, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x61, 0x70, 0x70, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x2e, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x08, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x12, 0x5d, 0x0a, 0x10, 0x70, 0x72,
	0x69, 0x6e, 0x63, 0x69, 0x70, 0x6c, 0x65, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x6c, 0x65, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69,
	0x70, 0x6c, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x2a, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x74, 0x61, 0x67, 0x22, 0x61, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x46, 0x0a, 0x08, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x61, 0x70, 0x70, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x52, 0x08,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x22, 0x59, 0x0a, 0x1d, 0x4f, 0x76, 0x65, 0x72,
	0x72, 0x69, 0x64, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x72, 0x54, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x54, 0x61, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x22, 0x20, 0x0a, 0x1e, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3a,
	0x19, 0x82, 0xb5, 0x18, 0x15, 0x0a, 0x0b, 0x67, 0x72, 0x70, 0x63, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x32, 0xa8, 0x04, 0x0a, 0x0e, 0x52,
	0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x87, 0x01,
	0x0a, 0x15, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x69,
	0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x3b, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x6d, 0x0a, 0x09, 0x54, 0x65, 0x73, 0x74, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x12, 0x2f, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x22, 0x00, 0x12, 0x82, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x35, 0x2e, 0x76, 0x32, 0x72,
	0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x36, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61,
	0x70, 0x70, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x97, 0x01, 0x0a, 0x16,
	0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x3c, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x3d, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x72, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x78, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x32, 0x72,
	0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x50, 0x01, 0x5a, 0x31, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x32, 0x66, 0x6c, 0x79, 0x2f, 0x76,
	0x32, 0x72, 0x61, 0x79, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x35, 0x2f, 0x61, 0x70, 0x70,
	0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0xaa,
	0x02, 0x1d, 0x56, 0x32, 0x52, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x70, 0x70,
	0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_app_router_command_command_proto_rawDescOnce sync.Once
	file_app_router_command_command_proto_rawDescData []byte
)

func file_app_router_command_command_proto_rawDescGZIP() []byte {
	file_app_router_command_command_proto_rawDescOnce.Do(func() {
		file_app_router_command_command_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_app_router_command_command_proto_rawDesc), len(file_app_router_command_command_proto_rawDesc)))
	})
	return file_app_router_command_command_proto_rawDescData
}

var file_app_router_command_command_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_app_router_command_command_proto_goTypes = []any{
	(*RoutingContext)(nil),                 // 0: v2ray.core.app.router.command.RoutingContext
	(*SubscribeRoutingStatsRequest)(nil),   // 1: v2ray.core.app.router.command.SubscribeRoutingStatsRequest
	(*TestRouteRequest)(nil),               // 2: v2ray.core.app.router.command.TestRouteRequest
	(*PrincipleTargetInfo)(nil),            // 3: v2ray.core.app.router.command.PrincipleTargetInfo
	(*OverrideInfo)(nil),                   // 4: v2ray.core.app.router.command.OverrideInfo
	(*BalancerMsg)(nil),                    // 5: v2ray.core.app.router.command.BalancerMsg
	(*GetBalancerInfoRequest)(nil),         // 6: v2ray.core.app.router.command.GetBalancerInfoRequest
	(*GetBalancerInfoResponse)(nil),        // 7: v2ray.core.app.router.command.GetBalancerInfoResponse
	(*OverrideBalancerTargetRequest)(nil),  // 8: v2ray.core.app.router.command.OverrideBalancerTargetRequest
	(*OverrideBalancerTargetResponse)(nil), // 9: v2ray.core.app.router.command.OverrideBalancerTargetResponse
	(*Config)(nil),                         // 10: v2ray.core.app.router.command.Config
	nil,                                    // 11: v2ray.core.app.router.command.RoutingContext.AttributesEntry
	(net.Network)(0),                       // 12: v2ray.core.common.net.Network
}
var file_app_router_command_command_proto_depIdxs = []int32{
	12, // 0: v2ray.core.app.router.command.RoutingContext.Network:type_name -> v2ray.core.common.net.Network
	11, // 1: v2ray.core.app.router.command.RoutingContext.Attributes:type_name -> v2ray.core.app.router.command.RoutingContext.AttributesEntry
	0,  // 2: v2ray.core.app.router.command.TestRouteRequest.RoutingContext:type_name -> v2ray.core.app.router.command.RoutingContext
	4,  // 3: v2ray.core.app.router.command.BalancerMsg.override:type_name -> v2ray.core.app.router.command.OverrideInfo
	3,  // 4: v2ray.core.app.router.command.BalancerMsg.principle_target:type_name -> v2ray.core.app.router.command.PrincipleTargetInfo
	5,  // 5: v2ray.core.app.router.command.GetBalancerInfoResponse.balancer:type_name -> v2ray.core.app.router.command.BalancerMsg
	1,  // 6: v2ray.core.app.router.command.RoutingService.SubscribeRoutingStats:input_type -> v2ray.core.app.router.command.SubscribeRoutingStatsRequest
	2,  // 7: v2ray.core.app.router.command.RoutingService.TestRoute:input_type -> v2ray.core.app.router.command.TestRouteRequest
	6,  // 8: v2ray.core.app.router.command.RoutingService.GetBalancerInfo:input_type -> v2ray.core.app.router.command.GetBalancerInfoRequest
	8,  // 9: v2ray.core.app.router.command.RoutingService.OverrideBalancerTarget:input_type -> v2ray.core.app.router.command.OverrideBalancerTargetRequest
	0,  // 10: v2ray.core.app.router.command.RoutingService.SubscribeRoutingStats:output_type -> v2ray.core.app.router.command.RoutingContext
	0,  // 11: v2ray.core.app.router.command.RoutingService.TestRoute:output_type -> v2ray.core.app.router.command.RoutingContext
	7,  // 12: v2ray.core.app.router.command.RoutingService.GetBalancerInfo:output_type -> v2ray.core.app.router.command.GetBalancerInfoResponse
	9,  // 13: v2ray.core.app.router.command.RoutingService.OverrideBalancerTarget:output_type -> v2ray.core.app.router.command.OverrideBalancerTargetResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_app_router_command_command_proto_init() }
func file_app_router_command_command_proto_init() {
	if File_app_router_command_command_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_app_router_command_command_proto_rawDesc), len(file_app_router_command_command_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_router_command_command_proto_goTypes,
		DependencyIndexes: file_app_router_command_command_proto_depIdxs,
		MessageInfos:      file_app_router_command_command_proto_msgTypes,
	}.Build()
	File_app_router_command_command_proto = out.File
	file_app_router_command_command_proto_goTypes = nil
	file_app_router_command_command_proto_depIdxs = nil
}
