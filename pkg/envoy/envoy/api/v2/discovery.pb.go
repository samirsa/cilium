// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/discovery.proto

package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import core "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/core"
import _ "github.com/gogo/protobuf/gogoproto"
import any "github.com/golang/protobuf/ptypes/any"
import status "google.golang.org/genproto/googleapis/rpc/status"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A DiscoveryRequest requests a set of versioned resources of the same type for
// a given Envoy node on some API.
type DiscoveryRequest struct {
	// The version_info provided in the request messages will be the version_info
	// received with the most recent successfully processed response or empty on
	// the first request. It is expected that no new request is sent after a
	// response is received until the Envoy instance is ready to ACK/NACK the new
	// configuration. ACK/NACK takes place by returning the new API config version
	// as applied or the previous API config version respectively. Each type_url
	// (see below) has an independent version associated with it.
	VersionInfo string `protobuf:"bytes,1,opt,name=version_info,json=versionInfo,proto3" json:"version_info,omitempty"`
	// The node making the request.
	Node *core.Node `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty"`
	// List of resources to subscribe to, e.g. list of cluster names or a route
	// configuration name. If this is empty, all resources for the API are
	// returned. LDS/CDS expect empty resource_names, since this is global
	// discovery for the Envoy instance. The LDS and CDS responses will then imply
	// a number of resources that need to be fetched via EDS/RDS, which will be
	// explicitly enumerated in resource_names.
	ResourceNames []string `protobuf:"bytes,3,rep,name=resource_names,json=resourceNames,proto3" json:"resource_names,omitempty"`
	// Type of the resource that is being requested, e.g.
	// "type.googleapis.com/envoy.api.v2.ClusterLoadAssignment". This is implicit
	// in requests made via singleton xDS APIs such as CDS, LDS, etc. but is
	// required for ADS.
	TypeUrl string `protobuf:"bytes,4,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	// nonce corresponding to DiscoveryResponse being ACK/NACKed. See above
	// discussion on version_info and the DiscoveryResponse nonce comment. This
	// may be empty if no nonce is available, e.g. at startup or for non-stream
	// xDS implementations.
	ResponseNonce string `protobuf:"bytes,5,opt,name=response_nonce,json=responseNonce,proto3" json:"response_nonce,omitempty"`
	// This is populated when the previous :ref:`DiscoveryResponse <envoy_api_msg_DiscoveryResponse>`
	// failed to update configuration. The *message* field in *error_details* provides the Envoy
	// internal exception related to the failure. It is only intended for consumption during manual
	// debugging, the string provided is not guaranteed to be stable across Envoy versions.
	ErrorDetail          *status.Status `protobuf:"bytes,6,opt,name=error_detail,json=errorDetail,proto3" json:"error_detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DiscoveryRequest) Reset()         { *m = DiscoveryRequest{} }
func (m *DiscoveryRequest) String() string { return proto.CompactTextString(m) }
func (*DiscoveryRequest) ProtoMessage()    {}
func (*DiscoveryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_discovery_1d479fb5ebf463ee, []int{0}
}
func (m *DiscoveryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoveryRequest.Unmarshal(m, b)
}
func (m *DiscoveryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoveryRequest.Marshal(b, m, deterministic)
}
func (dst *DiscoveryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoveryRequest.Merge(dst, src)
}
func (m *DiscoveryRequest) XXX_Size() int {
	return xxx_messageInfo_DiscoveryRequest.Size(m)
}
func (m *DiscoveryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoveryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoveryRequest proto.InternalMessageInfo

func (m *DiscoveryRequest) GetVersionInfo() string {
	if m != nil {
		return m.VersionInfo
	}
	return ""
}

func (m *DiscoveryRequest) GetNode() *core.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *DiscoveryRequest) GetResourceNames() []string {
	if m != nil {
		return m.ResourceNames
	}
	return nil
}

func (m *DiscoveryRequest) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}

func (m *DiscoveryRequest) GetResponseNonce() string {
	if m != nil {
		return m.ResponseNonce
	}
	return ""
}

func (m *DiscoveryRequest) GetErrorDetail() *status.Status {
	if m != nil {
		return m.ErrorDetail
	}
	return nil
}

type DiscoveryResponse struct {
	// The version of the response data.
	VersionInfo string `protobuf:"bytes,1,opt,name=version_info,json=versionInfo,proto3" json:"version_info,omitempty"`
	// The response resources. These resources are typed and depend on the API being called.
	Resources []*any.Any `protobuf:"bytes,2,rep,name=resources,proto3" json:"resources,omitempty"`
	// [#not-implemented-hide:]
	// Canary is used to support two Envoy command line flags:
	//
	// * --terminate-on-canary-transition-failure. When set, Envoy is able to
	//   terminate if it detects that configuration is stuck at canary. Consider
	//   this example sequence of updates:
	//   - Management server applies a canary config successfully.
	//   - Management server rolls back to a production config.
	//   - Envoy rejects the new production config.
	//   Since there is no sensible way to continue receiving configuration
	//   updates, Envoy will then terminate and apply production config from a
	//   clean slate.
	// * --dry-run-canary. When set, a canary response will never be applied, only
	//   validated via a dry run.
	Canary bool `protobuf:"varint,3,opt,name=canary,proto3" json:"canary,omitempty"`
	// Type URL for resources. This must be consistent with the type_url in the
	// Any messages for resources if resources is non-empty. This effectively
	// identifies the xDS API when muxing over ADS.
	TypeUrl string `protobuf:"bytes,4,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	// For gRPC based subscriptions, the nonce provides a way to explicitly ack a
	// specific DiscoveryResponse in a following DiscoveryRequest. Additional
	// messages may have been sent by Envoy to the management server for the
	// previous version on the stream prior to this DiscoveryResponse, that were
	// unprocessed at response send time. The nonce allows the management server
	// to ignore any further DiscoveryRequests for the previous version until a
	// DiscoveryRequest bearing the nonce. The nonce is optional and is not
	// required for non-stream based xDS implementations.
	Nonce                string   `protobuf:"bytes,5,opt,name=nonce,proto3" json:"nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiscoveryResponse) Reset()         { *m = DiscoveryResponse{} }
func (m *DiscoveryResponse) String() string { return proto.CompactTextString(m) }
func (*DiscoveryResponse) ProtoMessage()    {}
func (*DiscoveryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_discovery_1d479fb5ebf463ee, []int{1}
}
func (m *DiscoveryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoveryResponse.Unmarshal(m, b)
}
func (m *DiscoveryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoveryResponse.Marshal(b, m, deterministic)
}
func (dst *DiscoveryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoveryResponse.Merge(dst, src)
}
func (m *DiscoveryResponse) XXX_Size() int {
	return xxx_messageInfo_DiscoveryResponse.Size(m)
}
func (m *DiscoveryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoveryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoveryResponse proto.InternalMessageInfo

func (m *DiscoveryResponse) GetVersionInfo() string {
	if m != nil {
		return m.VersionInfo
	}
	return ""
}

func (m *DiscoveryResponse) GetResources() []*any.Any {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *DiscoveryResponse) GetCanary() bool {
	if m != nil {
		return m.Canary
	}
	return false
}

func (m *DiscoveryResponse) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}

func (m *DiscoveryResponse) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

// IncrementalDiscoveryRequest and IncrementalDiscoveryResponse are used in a
// new gRPC endpoint for Incremental xDS. The feature is not supported for REST
// management servers.
//
// With Incremental xDS, the IncrementalDiscoveryResponses do not need to
// include a full snapshot of the tracked resources. Instead
// IncrementalDiscoveryResponses are a diff to the state of a xDS client.
// In Incremental XDS there are per resource versions which allows to track
// state at the resource granularity.
// An xDS Incremental session is always in the context of a gRPC bidirectional
// stream. This allows the xDS server to keep track of the state of xDS clients
// connected to it.
//
// In Incremental xDS the nonce field is required and used to pair
// IncrementalDiscoveryResponse to a IncrementalDiscoveryRequest ACK or NACK.
// Optionaly, a response message level system_version_info is present for
// debugging purposes only.
//
// IncrementalDiscoveryRequest can be sent in 3 situations:
//   1. Initial message in a xDS bidirectional gRPC stream.
//   2. As a ACK or NACK response to a previous IncrementalDiscoveryResponse.
//      In this case the response_nonce is set to the nonce value in the Response.
//      ACK or NACK is determined by the absence or presence of error_detail.
//   3. Spontaneous IncrementalDiscoveryRequest from the client.
//      This can be done to dynamically add or remove elements from the tracked
//      resource_names set. In this case response_nonce must be omitted.
type IncrementalDiscoveryRequest struct {
	// The node making the request.
	Node *core.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	// Type of the resource that is being requested, e.g.
	// "type.googleapis.com/envoy.api.v2.ClusterLoadAssignment". This is implicit
	// in requests made via singleton xDS APIs such as CDS, LDS, etc. but is
	// required for ADS.
	TypeUrl string `protobuf:"bytes,2,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	// IncrementalDiscoveryRequests allow the client to add or remove individual
	// resources to the set of tracked resources in the context of a stream.
	// All resource names in the resource_names_subscribe list are added to the
	// set of tracked resources and all resource names in the resource_names_unsubscribe
	// list are removed from the set of tracked resources.
	// Unlike in non incremental xDS, an empty resource_names_subscribe or
	// resource_names_unsubscribe list simply means that no resources are to be
	// added or removed to the resource list.
	// The xDS server must send updates for all tracked resources but can also
	// send updates for resources the client has not subscribed to. This behavior
	// is similar to non incremental xDS.
	// These two fields can be set for all types of IncrementalDiscoveryRequests
	// (initial, ACK/NACK or spontaneous).
	//
	// A list of Resource names to add to the list of tracked resources.
	ResourceNamesSubscribe []string `protobuf:"bytes,3,rep,name=resource_names_subscribe,json=resourceNamesSubscribe,proto3" json:"resource_names_subscribe,omitempty"`
	// A list of Resource names to remove from the list of tracked resources.
	ResourceNamesUnsubscribe []string `protobuf:"bytes,4,rep,name=resource_names_unsubscribe,json=resourceNamesUnsubscribe,proto3" json:"resource_names_unsubscribe,omitempty"`
	// This map must be populated when the IncrementalDiscoveryRequest is the
	// first in a stream. The keys are the resources names of the xDS resources
	// known to the xDS client. The values in the map are the associated resource
	// level version info.
	InitialResourceVersions map[string]string `protobuf:"bytes,5,rep,name=initial_resource_versions,json=initialResourceVersions,proto3" json:"initial_resource_versions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// When the IncrementalDiscoveryRequest is a ACK or NACK message in response
	// to a previous IncrementalDiscoveryResponse, the response_nonce must be the
	// nonce in the IncrementalDiscoveryResponse.
	// Otherwise response_nonce must be omitted.
	ResponseNonce string `protobuf:"bytes,6,opt,name=response_nonce,json=responseNonce,proto3" json:"response_nonce,omitempty"`
	// This is populated when the previous :ref:`DiscoveryResponse <envoy_api_msg_DiscoveryResponse>`
	// failed to update configuration. The *message* field in *error_details*
	// provides the Envoy internal exception related to the failure.
	ErrorDetail          *status.Status `protobuf:"bytes,7,opt,name=error_detail,json=errorDetail,proto3" json:"error_detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *IncrementalDiscoveryRequest) Reset()         { *m = IncrementalDiscoveryRequest{} }
func (m *IncrementalDiscoveryRequest) String() string { return proto.CompactTextString(m) }
func (*IncrementalDiscoveryRequest) ProtoMessage()    {}
func (*IncrementalDiscoveryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_discovery_1d479fb5ebf463ee, []int{2}
}
func (m *IncrementalDiscoveryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IncrementalDiscoveryRequest.Unmarshal(m, b)
}
func (m *IncrementalDiscoveryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IncrementalDiscoveryRequest.Marshal(b, m, deterministic)
}
func (dst *IncrementalDiscoveryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncrementalDiscoveryRequest.Merge(dst, src)
}
func (m *IncrementalDiscoveryRequest) XXX_Size() int {
	return xxx_messageInfo_IncrementalDiscoveryRequest.Size(m)
}
func (m *IncrementalDiscoveryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IncrementalDiscoveryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IncrementalDiscoveryRequest proto.InternalMessageInfo

func (m *IncrementalDiscoveryRequest) GetNode() *core.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *IncrementalDiscoveryRequest) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}

func (m *IncrementalDiscoveryRequest) GetResourceNamesSubscribe() []string {
	if m != nil {
		return m.ResourceNamesSubscribe
	}
	return nil
}

func (m *IncrementalDiscoveryRequest) GetResourceNamesUnsubscribe() []string {
	if m != nil {
		return m.ResourceNamesUnsubscribe
	}
	return nil
}

func (m *IncrementalDiscoveryRequest) GetInitialResourceVersions() map[string]string {
	if m != nil {
		return m.InitialResourceVersions
	}
	return nil
}

func (m *IncrementalDiscoveryRequest) GetResponseNonce() string {
	if m != nil {
		return m.ResponseNonce
	}
	return ""
}

func (m *IncrementalDiscoveryRequest) GetErrorDetail() *status.Status {
	if m != nil {
		return m.ErrorDetail
	}
	return nil
}

type IncrementalDiscoveryResponse struct {
	// The version of the response data (used for debugging).
	SystemVersionInfo string `protobuf:"bytes,1,opt,name=system_version_info,json=systemVersionInfo,proto3" json:"system_version_info,omitempty"`
	// The response resources. These are typed resources that match the type url
	// in the IncrementalDiscoveryRequest.
	Resources []*Resource `protobuf:"bytes,2,rep,name=resources,proto3" json:"resources,omitempty"`
	// Resources names of resources that have be deleted and to be removed from the xDS Client.
	// Removed resources for missing resources can be ignored.
	RemovedResources []string `protobuf:"bytes,6,rep,name=removed_resources,json=removedResources,proto3" json:"removed_resources,omitempty"`
	// The nonce provides a way for IncrementalDiscoveryRequests to uniquely
	// reference a IncrementalDiscoveryResponse. The nonce is required.
	Nonce                string   `protobuf:"bytes,5,opt,name=nonce,proto3" json:"nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IncrementalDiscoveryResponse) Reset()         { *m = IncrementalDiscoveryResponse{} }
func (m *IncrementalDiscoveryResponse) String() string { return proto.CompactTextString(m) }
func (*IncrementalDiscoveryResponse) ProtoMessage()    {}
func (*IncrementalDiscoveryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_discovery_1d479fb5ebf463ee, []int{3}
}
func (m *IncrementalDiscoveryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IncrementalDiscoveryResponse.Unmarshal(m, b)
}
func (m *IncrementalDiscoveryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IncrementalDiscoveryResponse.Marshal(b, m, deterministic)
}
func (dst *IncrementalDiscoveryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncrementalDiscoveryResponse.Merge(dst, src)
}
func (m *IncrementalDiscoveryResponse) XXX_Size() int {
	return xxx_messageInfo_IncrementalDiscoveryResponse.Size(m)
}
func (m *IncrementalDiscoveryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IncrementalDiscoveryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IncrementalDiscoveryResponse proto.InternalMessageInfo

func (m *IncrementalDiscoveryResponse) GetSystemVersionInfo() string {
	if m != nil {
		return m.SystemVersionInfo
	}
	return ""
}

func (m *IncrementalDiscoveryResponse) GetResources() []*Resource {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *IncrementalDiscoveryResponse) GetRemovedResources() []string {
	if m != nil {
		return m.RemovedResources
	}
	return nil
}

func (m *IncrementalDiscoveryResponse) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

type Resource struct {
	// The resource level version. It allows xDS to track the state of individual
	// resources.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// The resource being tracked.
	Resource             *any.Any `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_discovery_1d479fb5ebf463ee, []int{4}
}
func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (dst *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(dst, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Resource) GetResource() *any.Any {
	if m != nil {
		return m.Resource
	}
	return nil
}

func init() {
	proto.RegisterType((*DiscoveryRequest)(nil), "envoy.api.v2.DiscoveryRequest")
	proto.RegisterType((*DiscoveryResponse)(nil), "envoy.api.v2.DiscoveryResponse")
	proto.RegisterType((*IncrementalDiscoveryRequest)(nil), "envoy.api.v2.IncrementalDiscoveryRequest")
	proto.RegisterMapType((map[string]string)(nil), "envoy.api.v2.IncrementalDiscoveryRequest.InitialResourceVersionsEntry")
	proto.RegisterType((*IncrementalDiscoveryResponse)(nil), "envoy.api.v2.IncrementalDiscoveryResponse")
	proto.RegisterType((*Resource)(nil), "envoy.api.v2.Resource")
}

func init() {
	proto.RegisterFile("envoy/api/v2/discovery.proto", fileDescriptor_discovery_1d479fb5ebf463ee)
}

var fileDescriptor_discovery_1d479fb5ebf463ee = []byte{
	// 618 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xc1, 0x6a, 0xdb, 0x40,
	0x10, 0xad, 0x64, 0xc7, 0x71, 0xd6, 0x49, 0x49, 0xb6, 0x21, 0xd9, 0xb8, 0xa1, 0x75, 0x0d, 0x05,
	0x43, 0x60, 0x55, 0x5c, 0x0a, 0x21, 0xf4, 0xd2, 0x90, 0x16, 0xd2, 0x43, 0x0e, 0x0a, 0xc9, 0xa1,
	0x17, 0xb1, 0x96, 0x27, 0x46, 0x54, 0xde, 0x55, 0x77, 0x25, 0x81, 0xae, 0x39, 0xf7, 0x43, 0xfa,
	0x0b, 0xfd, 0x83, 0x5e, 0xfa, 0x0b, 0x3d, 0xf4, 0x4b, 0x8a, 0x56, 0xbb, 0xb1, 0xd5, 0x08, 0xe3,
	0xdb, 0xce, 0xcc, 0xdb, 0x37, 0xb3, 0xf3, 0x9e, 0x84, 0x8e, 0x81, 0xe7, 0xa2, 0xf0, 0x58, 0x12,
	0x79, 0xf9, 0xd8, 0x9b, 0x46, 0x2a, 0x14, 0x39, 0xc8, 0x82, 0x26, 0x52, 0xa4, 0x02, 0x6f, 0xeb,
	0x2a, 0x65, 0x49, 0x44, 0xf3, 0x71, 0xbf, 0x8e, 0x0d, 0x85, 0x04, 0x6f, 0xc2, 0x14, 0x54, 0xd8,
	0xfe, 0xd1, 0x4c, 0x88, 0x59, 0x0c, 0x9e, 0x8e, 0x26, 0xd9, 0x9d, 0xc7, 0xb8, 0xa1, 0xe9, 0x1f,
	0x9a, 0x92, 0x4c, 0x42, 0x4f, 0xa5, 0x2c, 0xcd, 0x94, 0x29, 0xec, 0xcf, 0xc4, 0x4c, 0xe8, 0xa3,
	0x57, 0x9e, 0xaa, 0xec, 0xf0, 0xde, 0x45, 0xbb, 0x17, 0x76, 0x12, 0x1f, 0xbe, 0x65, 0xa0, 0x52,
	0xfc, 0x0a, 0x6d, 0xe7, 0x20, 0x55, 0x24, 0x78, 0x10, 0xf1, 0x3b, 0x41, 0x9c, 0x81, 0x33, 0xda,
	0xf2, 0x7b, 0x26, 0x77, 0xc9, 0xef, 0x04, 0x3e, 0x41, 0x6d, 0x2e, 0xa6, 0x40, 0xdc, 0x81, 0x33,
	0xea, 0x8d, 0x0f, 0xe9, 0xf2, 0xf0, 0xb4, 0x1c, 0x97, 0x5e, 0x89, 0x29, 0xf8, 0x1a, 0x84, 0x5f,
	0xa3, 0xa7, 0x12, 0x94, 0xc8, 0x64, 0x08, 0x01, 0x67, 0x73, 0x50, 0xa4, 0x35, 0x68, 0x8d, 0xb6,
	0xfc, 0x1d, 0x9b, 0xbd, 0x2a, 0x93, 0xf8, 0x08, 0x75, 0xd3, 0x22, 0x81, 0x20, 0x93, 0x31, 0x69,
	0xeb, 0x96, 0x9b, 0x65, 0x7c, 0x23, 0x63, 0xc3, 0x90, 0x08, 0xae, 0x20, 0xe0, 0x82, 0x87, 0x40,
	0x36, 0x34, 0x60, 0xc7, 0x66, 0xaf, 0xca, 0x24, 0x7e, 0x87, 0xb6, 0x41, 0x4a, 0x21, 0x83, 0x29,
	0xa4, 0x2c, 0x8a, 0x49, 0x47, 0x4f, 0x87, 0x69, 0xb5, 0x13, 0x2a, 0x93, 0x90, 0x5e, 0xeb, 0x9d,
	0xf8, 0x3d, 0x8d, 0xbb, 0xd0, 0xb0, 0xe1, 0x4f, 0x07, 0xed, 0x2d, 0x2d, 0xa1, 0x62, 0x5c, 0x67,
	0x0b, 0xa7, 0x68, 0xcb, 0x3e, 0x41, 0x11, 0x77, 0xd0, 0x1a, 0xf5, 0xc6, 0xfb, 0xb6, 0x99, 0xd5,
	0x86, 0x7e, 0xe0, 0xc5, 0x79, 0xfb, 0xd7, 0x9f, 0x97, 0x4f, 0xfc, 0x05, 0x18, 0x1f, 0xa0, 0x4e,
	0xc8, 0x38, 0x93, 0x05, 0x69, 0x0d, 0x9c, 0x51, 0xd7, 0x37, 0xd1, 0xaa, 0x1d, 0xec, 0xa3, 0x8d,
	0xe5, 0xa7, 0x57, 0xc1, 0xf0, 0x7b, 0x1b, 0x3d, 0xbf, 0xe4, 0xa1, 0x84, 0x39, 0xf0, 0x94, 0xc5,
	0x8f, 0xb4, 0xb4, 0x42, 0x39, 0xeb, 0x08, 0xb5, 0xdc, 0xdd, 0xad, 0x77, 0x3f, 0x45, 0xa4, 0xae,
	0x61, 0xa0, 0xb2, 0x89, 0x0a, 0x65, 0x34, 0x01, 0xa3, 0xe6, 0x41, 0x4d, 0xcd, 0x6b, 0x5b, 0xc5,
	0xef, 0x51, 0xff, 0xbf, 0x9b, 0x19, 0x5f, 0xdc, 0x6d, 0xeb, 0xbb, 0xa4, 0x76, 0xf7, 0x66, 0x51,
	0xc7, 0xf7, 0x0e, 0x3a, 0x8a, 0x78, 0x94, 0x46, 0x2c, 0x0e, 0x1e, 0x68, 0x8c, 0x06, 0x8a, 0x6c,
	0xe8, 0x9d, 0x7f, 0xaa, 0xbf, 0x6a, 0xc5, 0x3a, 0xe8, 0x65, 0x45, 0xe5, 0x1b, 0xa6, 0x5b, 0x43,
	0xf4, 0x91, 0xa7, 0xb2, 0xf0, 0x0f, 0xa3, 0xe6, 0x6a, 0x83, 0xfd, 0x3a, 0xeb, 0xd8, 0x6f, 0x73,
	0x2d, 0xfb, 0xf5, 0x3f, 0xa3, 0xe3, 0x55, 0x63, 0xe1, 0x5d, 0xd4, 0xfa, 0x0a, 0x85, 0xf1, 0x5f,
	0x79, 0x2c, 0xad, 0x90, 0xb3, 0x38, 0x03, 0x23, 0x52, 0x15, 0x9c, 0xb9, 0xa7, 0xce, 0xf0, 0xb7,
	0x53, 0x92, 0x35, 0xbd, 0xdf, 0xb8, 0x9a, 0xa2, 0x67, 0xaa, 0x50, 0x29, 0xcc, 0x83, 0x06, 0x73,
	0xef, 0x55, 0xa5, 0xdb, 0x25, 0x8b, 0x9f, 0x3d, 0xb6, 0xf8, 0x41, 0x7d, 0xdd, 0x76, 0xe8, 0xc7,
	0x26, 0x3f, 0x41, 0x7b, 0x12, 0xe6, 0x22, 0x87, 0x69, 0xb0, 0xe0, 0xe8, 0x68, 0xc1, 0x77, 0x4d,
	0xc1, 0x7f, 0x00, 0x37, 0xdb, 0xfb, 0x16, 0x75, 0x2d, 0x04, 0x13, 0xb4, 0x69, 0x66, 0x36, 0xe3,
	0xda, 0x10, 0xbf, 0x41, 0x5d, 0xdb, 0xc0, 0xfc, 0x91, 0x1a, 0x3f, 0x43, 0xff, 0x01, 0x75, 0xde,
	0xfd, 0xf1, 0xf7, 0x85, 0xf3, 0xc5, 0xcd, 0xc7, 0x93, 0x8e, 0x46, 0xbc, 0xfd, 0x17, 0x00, 0x00,
	0xff, 0xff, 0x85, 0xc4, 0x84, 0xb3, 0x9e, 0x05, 0x00, 0x00,
}