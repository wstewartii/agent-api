// Code generated by protoc-gen-go. DO NOT EDIT.
// source: agent_api.proto

/*
Package kolide_agent is a generated protocol buffer package.

It is generated from these files:
	agent_api.proto

It has these top-level messages:
	AgentApiRequest
	AgentApiResponse
	EnrollmentRequest
	EnrollmentResponse
	ConfigResponse
	LogCollection
	QueryCollection
	ResultCollection
	HealthCheckResponse
*/
package kolide_agent

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// AGENT is added as a new log type, for adding new
// logging capabilities from kolide/agent
type LogCollection_LogType int32

const (
	LogCollection_RESULT LogCollection_LogType = 0
	LogCollection_STATUS LogCollection_LogType = 1
	LogCollection_AGENT  LogCollection_LogType = 2
)

var LogCollection_LogType_name = map[int32]string{
	0: "RESULT",
	1: "STATUS",
	2: "AGENT",
}
var LogCollection_LogType_value = map[string]int32{
	"RESULT": 0,
	"STATUS": 1,
	"AGENT":  2,
}

func (x LogCollection_LogType) String() string {
	return proto.EnumName(LogCollection_LogType_name, int32(x))
}
func (LogCollection_LogType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

type HealthCheckResponse_ServingStatus int32

const (
	HealthCheckResponse_UNKNOWN     HealthCheckResponse_ServingStatus = 0
	HealthCheckResponse_SERVING     HealthCheckResponse_ServingStatus = 1
	HealthCheckResponse_NOT_SERVING HealthCheckResponse_ServingStatus = 2
)

var HealthCheckResponse_ServingStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "SERVING",
	2: "NOT_SERVING",
}
var HealthCheckResponse_ServingStatus_value = map[string]int32{
	"UNKNOWN":     0,
	"SERVING":     1,
	"NOT_SERVING": 2,
}

func (x HealthCheckResponse_ServingStatus) String() string {
	return proto.EnumName(HealthCheckResponse_ServingStatus_name, int32(x))
}
func (HealthCheckResponse_ServingStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{8, 0}
}

type AgentApiRequest struct {
	NodeKey string `protobuf:"bytes,1,opt,name=node_key,json=nodeKey" json:"node_key,omitempty"`
}

func (m *AgentApiRequest) Reset()                    { *m = AgentApiRequest{} }
func (m *AgentApiRequest) String() string            { return proto.CompactTextString(m) }
func (*AgentApiRequest) ProtoMessage()               {}
func (*AgentApiRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AgentApiRequest) GetNodeKey() string {
	if m != nil {
		return m.NodeKey
	}
	return ""
}

type AgentApiResponse struct {
	Message     string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	ErrorCode   string `protobuf:"bytes,2,opt,name=error_code,json=errorCode" json:"error_code,omitempty"`
	NodeInvalid bool   `protobuf:"varint,3,opt,name=node_invalid,json=nodeInvalid" json:"node_invalid,omitempty"`
}

func (m *AgentApiResponse) Reset()                    { *m = AgentApiResponse{} }
func (m *AgentApiResponse) String() string            { return proto.CompactTextString(m) }
func (*AgentApiResponse) ProtoMessage()               {}
func (*AgentApiResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AgentApiResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *AgentApiResponse) GetErrorCode() string {
	if m != nil {
		return m.ErrorCode
	}
	return ""
}

func (m *AgentApiResponse) GetNodeInvalid() bool {
	if m != nil {
		return m.NodeInvalid
	}
	return false
}

type EnrollmentRequest struct {
	EnrollSecret   string `protobuf:"bytes,1,opt,name=enroll_secret,json=enrollSecret" json:"enroll_secret,omitempty"`
	HostIdentifier string `protobuf:"bytes,2,opt,name=host_identifier,json=hostIdentifier" json:"host_identifier,omitempty"`
}

func (m *EnrollmentRequest) Reset()                    { *m = EnrollmentRequest{} }
func (m *EnrollmentRequest) String() string            { return proto.CompactTextString(m) }
func (*EnrollmentRequest) ProtoMessage()               {}
func (*EnrollmentRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *EnrollmentRequest) GetEnrollSecret() string {
	if m != nil {
		return m.EnrollSecret
	}
	return ""
}

func (m *EnrollmentRequest) GetHostIdentifier() string {
	if m != nil {
		return m.HostIdentifier
	}
	return ""
}

type EnrollmentResponse struct {
	NodeKey     string `protobuf:"bytes,1,opt,name=node_key,json=nodeKey" json:"node_key,omitempty"`
	NodeInvalid bool   `protobuf:"varint,2,opt,name=node_invalid,json=nodeInvalid" json:"node_invalid,omitempty"`
	ErrorCode   string `protobuf:"bytes,3,opt,name=error_code,json=errorCode" json:"error_code,omitempty"`
}

func (m *EnrollmentResponse) Reset()                    { *m = EnrollmentResponse{} }
func (m *EnrollmentResponse) String() string            { return proto.CompactTextString(m) }
func (*EnrollmentResponse) ProtoMessage()               {}
func (*EnrollmentResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *EnrollmentResponse) GetNodeKey() string {
	if m != nil {
		return m.NodeKey
	}
	return ""
}

func (m *EnrollmentResponse) GetNodeInvalid() bool {
	if m != nil {
		return m.NodeInvalid
	}
	return false
}

func (m *EnrollmentResponse) GetErrorCode() string {
	if m != nil {
		return m.ErrorCode
	}
	return ""
}

// kolide/cloud will be generating well-structured json already, so forward
// it along instead of de/re/de-serializing it as a protobuf too
// this might make sense to convert to full proto later
type ConfigResponse struct {
	ConfigJsonBlob string `protobuf:"bytes,1,opt,name=config_json_blob,json=configJsonBlob" json:"config_json_blob,omitempty"`
	NodeInvalid    bool   `protobuf:"varint,2,opt,name=node_invalid,json=nodeInvalid" json:"node_invalid,omitempty"`
	ErrorCode      string `protobuf:"bytes,3,opt,name=error_code,json=errorCode" json:"error_code,omitempty"`
}

func (m *ConfigResponse) Reset()                    { *m = ConfigResponse{} }
func (m *ConfigResponse) String() string            { return proto.CompactTextString(m) }
func (*ConfigResponse) ProtoMessage()               {}
func (*ConfigResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ConfigResponse) GetConfigJsonBlob() string {
	if m != nil {
		return m.ConfigJsonBlob
	}
	return ""
}

func (m *ConfigResponse) GetNodeInvalid() bool {
	if m != nil {
		return m.NodeInvalid
	}
	return false
}

func (m *ConfigResponse) GetErrorCode() string {
	if m != nil {
		return m.ErrorCode
	}
	return ""
}

type LogCollection struct {
	NodeKey   string                `protobuf:"bytes,1,opt,name=node_key,json=nodeKey" json:"node_key,omitempty"`
	LogType   LogCollection_LogType `protobuf:"varint,2,opt,name=log_type,json=logType,enum=kolide.agent.LogCollection_LogType" json:"log_type,omitempty"`
	Logs      []*LogCollection_Log  `protobuf:"bytes,3,rep,name=logs" json:"logs,omitempty"`
	ErrorCode string                `protobuf:"bytes,4,opt,name=error_code,json=errorCode" json:"error_code,omitempty"`
}

func (m *LogCollection) Reset()                    { *m = LogCollection{} }
func (m *LogCollection) String() string            { return proto.CompactTextString(m) }
func (*LogCollection) ProtoMessage()               {}
func (*LogCollection) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *LogCollection) GetNodeKey() string {
	if m != nil {
		return m.NodeKey
	}
	return ""
}

func (m *LogCollection) GetLogType() LogCollection_LogType {
	if m != nil {
		return m.LogType
	}
	return LogCollection_RESULT
}

func (m *LogCollection) GetLogs() []*LogCollection_Log {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *LogCollection) GetErrorCode() string {
	if m != nil {
		return m.ErrorCode
	}
	return ""
}

type LogCollection_Log struct {
	Data string `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
}

func (m *LogCollection_Log) Reset()                    { *m = LogCollection_Log{} }
func (m *LogCollection_Log) String() string            { return proto.CompactTextString(m) }
func (*LogCollection_Log) ProtoMessage()               {}
func (*LogCollection_Log) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

func (m *LogCollection_Log) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

// a query collection contains many queries
type QueryCollection struct {
	Queries     []*QueryCollection_Query `protobuf:"bytes,1,rep,name=queries" json:"queries,omitempty"`
	NodeInvalid bool                     `protobuf:"varint,2,opt,name=node_invalid,json=nodeInvalid" json:"node_invalid,omitempty"`
	ErrorCode   string                   `protobuf:"bytes,3,opt,name=error_code,json=errorCode" json:"error_code,omitempty"`
}

func (m *QueryCollection) Reset()                    { *m = QueryCollection{} }
func (m *QueryCollection) String() string            { return proto.CompactTextString(m) }
func (*QueryCollection) ProtoMessage()               {}
func (*QueryCollection) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *QueryCollection) GetQueries() []*QueryCollection_Query {
	if m != nil {
		return m.Queries
	}
	return nil
}

func (m *QueryCollection) GetNodeInvalid() bool {
	if m != nil {
		return m.NodeInvalid
	}
	return false
}

func (m *QueryCollection) GetErrorCode() string {
	if m != nil {
		return m.ErrorCode
	}
	return ""
}

type QueryCollection_Query struct {
	Id    string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Query string `protobuf:"bytes,2,opt,name=query" json:"query,omitempty"`
}

func (m *QueryCollection_Query) Reset()                    { *m = QueryCollection_Query{} }
func (m *QueryCollection_Query) String() string            { return proto.CompactTextString(m) }
func (*QueryCollection_Query) ProtoMessage()               {}
func (*QueryCollection_Query) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 0} }

func (m *QueryCollection_Query) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *QueryCollection_Query) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

// a result collection contains many results
type ResultCollection struct {
	NodeKey   string                     `protobuf:"bytes,1,opt,name=node_key,json=nodeKey" json:"node_key,omitempty"`
	Results   []*ResultCollection_Result `protobuf:"bytes,2,rep,name=results" json:"results,omitempty"`
	ErrorCode string                     `protobuf:"bytes,3,opt,name=error_code,json=errorCode" json:"error_code,omitempty"`
}

func (m *ResultCollection) Reset()                    { *m = ResultCollection{} }
func (m *ResultCollection) String() string            { return proto.CompactTextString(m) }
func (*ResultCollection) ProtoMessage()               {}
func (*ResultCollection) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ResultCollection) GetNodeKey() string {
	if m != nil {
		return m.NodeKey
	}
	return ""
}

func (m *ResultCollection) GetResults() []*ResultCollection_Result {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *ResultCollection) GetErrorCode() string {
	if m != nil {
		return m.ErrorCode
	}
	return ""
}

// status is moved here instead of appearing as a map of id[status]
// on the ResultCollection, as it does in the osq docs
type ResultCollection_Result struct {
	Id     string                               `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Rows   []*ResultCollection_Result_ResultRow `protobuf:"bytes,2,rep,name=rows" json:"rows,omitempty"`
	Status int32                                `protobuf:"varint,3,opt,name=status" json:"status,omitempty"`
}

func (m *ResultCollection_Result) Reset()                    { *m = ResultCollection_Result{} }
func (m *ResultCollection_Result) String() string            { return proto.CompactTextString(m) }
func (*ResultCollection_Result) ProtoMessage()               {}
func (*ResultCollection_Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7, 0} }

func (m *ResultCollection_Result) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ResultCollection_Result) GetRows() []*ResultCollection_Result_ResultRow {
	if m != nil {
		return m.Rows
	}
	return nil
}

func (m *ResultCollection_Result) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type ResultCollection_Result_ResultRow struct {
	Columns []*ResultCollection_Result_ResultRow_Column `protobuf:"bytes,1,rep,name=columns" json:"columns,omitempty"`
}

func (m *ResultCollection_Result_ResultRow) Reset()         { *m = ResultCollection_Result_ResultRow{} }
func (m *ResultCollection_Result_ResultRow) String() string { return proto.CompactTextString(m) }
func (*ResultCollection_Result_ResultRow) ProtoMessage()    {}
func (*ResultCollection_Result_ResultRow) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{7, 0, 0}
}

func (m *ResultCollection_Result_ResultRow) GetColumns() []*ResultCollection_Result_ResultRow_Column {
	if m != nil {
		return m.Columns
	}
	return nil
}

type ResultCollection_Result_ResultRow_Column struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *ResultCollection_Result_ResultRow_Column) Reset() {
	*m = ResultCollection_Result_ResultRow_Column{}
}
func (m *ResultCollection_Result_ResultRow_Column) String() string { return proto.CompactTextString(m) }
func (*ResultCollection_Result_ResultRow_Column) ProtoMessage()    {}
func (*ResultCollection_Result_ResultRow_Column) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{7, 0, 0, 0}
}

func (m *ResultCollection_Result_ResultRow_Column) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ResultCollection_Result_ResultRow_Column) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type HealthCheckResponse struct {
	Status HealthCheckResponse_ServingStatus `protobuf:"varint,1,opt,name=status,enum=kolide.agent.HealthCheckResponse_ServingStatus" json:"status,omitempty"`
}

func (m *HealthCheckResponse) Reset()                    { *m = HealthCheckResponse{} }
func (m *HealthCheckResponse) String() string            { return proto.CompactTextString(m) }
func (*HealthCheckResponse) ProtoMessage()               {}
func (*HealthCheckResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *HealthCheckResponse) GetStatus() HealthCheckResponse_ServingStatus {
	if m != nil {
		return m.Status
	}
	return HealthCheckResponse_UNKNOWN
}

func init() {
	proto.RegisterType((*AgentApiRequest)(nil), "kolide.agent.AgentApiRequest")
	proto.RegisterType((*AgentApiResponse)(nil), "kolide.agent.AgentApiResponse")
	proto.RegisterType((*EnrollmentRequest)(nil), "kolide.agent.EnrollmentRequest")
	proto.RegisterType((*EnrollmentResponse)(nil), "kolide.agent.EnrollmentResponse")
	proto.RegisterType((*ConfigResponse)(nil), "kolide.agent.ConfigResponse")
	proto.RegisterType((*LogCollection)(nil), "kolide.agent.LogCollection")
	proto.RegisterType((*LogCollection_Log)(nil), "kolide.agent.LogCollection.Log")
	proto.RegisterType((*QueryCollection)(nil), "kolide.agent.QueryCollection")
	proto.RegisterType((*QueryCollection_Query)(nil), "kolide.agent.QueryCollection.Query")
	proto.RegisterType((*ResultCollection)(nil), "kolide.agent.ResultCollection")
	proto.RegisterType((*ResultCollection_Result)(nil), "kolide.agent.ResultCollection.Result")
	proto.RegisterType((*ResultCollection_Result_ResultRow)(nil), "kolide.agent.ResultCollection.Result.ResultRow")
	proto.RegisterType((*ResultCollection_Result_ResultRow_Column)(nil), "kolide.agent.ResultCollection.Result.ResultRow.Column")
	proto.RegisterType((*HealthCheckResponse)(nil), "kolide.agent.HealthCheckResponse")
	proto.RegisterEnum("kolide.agent.LogCollection_LogType", LogCollection_LogType_name, LogCollection_LogType_value)
	proto.RegisterEnum("kolide.agent.HealthCheckResponse_ServingStatus", HealthCheckResponse_ServingStatus_name, HealthCheckResponse_ServingStatus_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Api service

type ApiClient interface {
	// Attempt to enroll a host with kolide/cloud
	RequestEnrollment(ctx context.Context, in *EnrollmentRequest, opts ...grpc.CallOption) (*EnrollmentResponse, error)
	// request an updated configuration from kolide/cloud
	// a generic request object is sent
	RequestConfig(ctx context.Context, in *AgentApiRequest, opts ...grpc.CallOption) (*ConfigResponse, error)
	// request/pull Dist queries from kolide/cloud
	// a generic request object is sent
	RequestQueries(ctx context.Context, in *AgentApiRequest, opts ...grpc.CallOption) (*QueryCollection, error)
	// publish logs from osqueryd to kolide/cloud
	// a generic response object is returned
	PublishLogs(ctx context.Context, in *LogCollection, opts ...grpc.CallOption) (*AgentApiResponse, error)
	// publish results from Dist queries to kolide/cloud
	// a generic response object is returned
	PublishResults(ctx context.Context, in *ResultCollection, opts ...grpc.CallOption) (*AgentApiResponse, error)
	// check the health of the GRPC server
	// a value indicating healthiness is returned
	// if you don't hear from this endpoint assume the worst
	CheckHealth(ctx context.Context, in *AgentApiRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type apiClient struct {
	cc *grpc.ClientConn
}

func NewApiClient(cc *grpc.ClientConn) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) RequestEnrollment(ctx context.Context, in *EnrollmentRequest, opts ...grpc.CallOption) (*EnrollmentResponse, error) {
	out := new(EnrollmentResponse)
	err := grpc.Invoke(ctx, "/kolide.agent.Api/RequestEnrollment", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) RequestConfig(ctx context.Context, in *AgentApiRequest, opts ...grpc.CallOption) (*ConfigResponse, error) {
	out := new(ConfigResponse)
	err := grpc.Invoke(ctx, "/kolide.agent.Api/RequestConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) RequestQueries(ctx context.Context, in *AgentApiRequest, opts ...grpc.CallOption) (*QueryCollection, error) {
	out := new(QueryCollection)
	err := grpc.Invoke(ctx, "/kolide.agent.Api/RequestQueries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) PublishLogs(ctx context.Context, in *LogCollection, opts ...grpc.CallOption) (*AgentApiResponse, error) {
	out := new(AgentApiResponse)
	err := grpc.Invoke(ctx, "/kolide.agent.Api/PublishLogs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) PublishResults(ctx context.Context, in *ResultCollection, opts ...grpc.CallOption) (*AgentApiResponse, error) {
	out := new(AgentApiResponse)
	err := grpc.Invoke(ctx, "/kolide.agent.Api/PublishResults", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) CheckHealth(ctx context.Context, in *AgentApiRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := grpc.Invoke(ctx, "/kolide.agent.Api/CheckHealth", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Api service

type ApiServer interface {
	// Attempt to enroll a host with kolide/cloud
	RequestEnrollment(context.Context, *EnrollmentRequest) (*EnrollmentResponse, error)
	// request an updated configuration from kolide/cloud
	// a generic request object is sent
	RequestConfig(context.Context, *AgentApiRequest) (*ConfigResponse, error)
	// request/pull Dist queries from kolide/cloud
	// a generic request object is sent
	RequestQueries(context.Context, *AgentApiRequest) (*QueryCollection, error)
	// publish logs from osqueryd to kolide/cloud
	// a generic response object is returned
	PublishLogs(context.Context, *LogCollection) (*AgentApiResponse, error)
	// publish results from Dist queries to kolide/cloud
	// a generic response object is returned
	PublishResults(context.Context, *ResultCollection) (*AgentApiResponse, error)
	// check the health of the GRPC server
	// a value indicating healthiness is returned
	// if you don't hear from this endpoint assume the worst
	CheckHealth(context.Context, *AgentApiRequest) (*HealthCheckResponse, error)
}

func RegisterApiServer(s *grpc.Server, srv ApiServer) {
	s.RegisterService(&_Api_serviceDesc, srv)
}

func _Api_RequestEnrollment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnrollmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).RequestEnrollment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kolide.agent.Api/RequestEnrollment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).RequestEnrollment(ctx, req.(*EnrollmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_RequestConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentApiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).RequestConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kolide.agent.Api/RequestConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).RequestConfig(ctx, req.(*AgentApiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_RequestQueries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentApiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).RequestQueries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kolide.agent.Api/RequestQueries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).RequestQueries(ctx, req.(*AgentApiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_PublishLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogCollection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).PublishLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kolide.agent.Api/PublishLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).PublishLogs(ctx, req.(*LogCollection))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_PublishResults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResultCollection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).PublishResults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kolide.agent.Api/PublishResults",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).PublishResults(ctx, req.(*ResultCollection))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_CheckHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentApiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).CheckHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kolide.agent.Api/CheckHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).CheckHealth(ctx, req.(*AgentApiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Api_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kolide.agent.Api",
	HandlerType: (*ApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestEnrollment",
			Handler:    _Api_RequestEnrollment_Handler,
		},
		{
			MethodName: "RequestConfig",
			Handler:    _Api_RequestConfig_Handler,
		},
		{
			MethodName: "RequestQueries",
			Handler:    _Api_RequestQueries_Handler,
		},
		{
			MethodName: "PublishLogs",
			Handler:    _Api_PublishLogs_Handler,
		},
		{
			MethodName: "PublishResults",
			Handler:    _Api_PublishResults_Handler,
		},
		{
			MethodName: "CheckHealth",
			Handler:    _Api_CheckHealth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agent_api.proto",
}

func init() { proto.RegisterFile("agent_api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 791 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x41, 0x6f, 0xea, 0x46,
	0x10, 0xc6, 0x36, 0xe0, 0x30, 0x04, 0xf0, 0xdb, 0x56, 0x15, 0x8f, 0x36, 0x29, 0x71, 0x54, 0x95,
	0x43, 0x4a, 0x25, 0x22, 0xf5, 0x50, 0xa9, 0xad, 0x28, 0x42, 0x69, 0x1a, 0x44, 0x12, 0x43, 0xda,
	0xde, 0x2c, 0x83, 0x37, 0x66, 0x9b, 0xc5, 0x4b, 0xbc, 0x76, 0x22, 0xa4, 0xde, 0xab, 0x1e, 0x7b,
	0x6c, 0xff, 0x4d, 0xff, 0x54, 0xcf, 0x95, 0xd7, 0x6b, 0x12, 0x4c, 0x03, 0xa9, 0xfa, 0x6e, 0x3b,
	0x9f, 0x67, 0xe6, 0x9b, 0x6f, 0x77, 0x66, 0x00, 0x6a, 0x8e, 0x87, 0xfd, 0xd0, 0x76, 0x16, 0xa4,
	0xbd, 0x08, 0x58, 0xc8, 0xd0, 0xfe, 0x1d, 0xa3, 0xc4, 0xc5, 0x6d, 0x81, 0x9b, 0x27, 0x50, 0xeb,
	0xc6, 0x87, 0xee, 0x82, 0x58, 0xf8, 0x3e, 0xc2, 0x3c, 0x44, 0x6f, 0x61, 0xcf, 0x67, 0x2e, 0xb6,
	0xef, 0xf0, 0xb2, 0xae, 0x34, 0x95, 0x56, 0xc9, 0xd2, 0x63, 0xfb, 0x02, 0x2f, 0x4d, 0x1f, 0x8c,
	0x27, 0x6f, 0xbe, 0x60, 0x3e, 0xc7, 0xa8, 0x0e, 0xfa, 0x1c, 0x73, 0xee, 0x78, 0x38, 0xf5, 0x96,
	0x26, 0x3a, 0x00, 0xc0, 0x41, 0xc0, 0x02, 0x7b, 0xca, 0x5c, 0x5c, 0x57, 0xc5, 0xc7, 0x92, 0x40,
	0x7a, 0xcc, 0xc5, 0xe8, 0x08, 0xf6, 0x05, 0x0f, 0xf1, 0x1f, 0x1c, 0x4a, 0xdc, 0xba, 0xd6, 0x54,
	0x5a, 0x7b, 0x56, 0x39, 0xc6, 0xce, 0x13, 0xc8, 0x74, 0xe0, 0x4d, 0xdf, 0x0f, 0x18, 0xa5, 0x73,
	0xec, 0x87, 0x69, 0x7d, 0xc7, 0x50, 0xc1, 0x02, 0xb4, 0x39, 0x9e, 0x06, 0x38, 0x94, 0xb4, 0xfb,
	0x09, 0x38, 0x12, 0x18, 0xfa, 0x14, 0x6a, 0x33, 0xc6, 0x43, 0x9b, 0xb8, 0xd8, 0x0f, 0xc9, 0x2d,
	0xc1, 0x81, 0x2c, 0xa0, 0x1a, 0xc3, 0xe7, 0x2b, 0xd4, 0xbc, 0x07, 0xf4, 0x9c, 0x42, 0x8a, 0x7a,
	0xf9, 0x0e, 0x36, 0xca, 0x56, 0x37, 0xca, 0xce, 0x08, 0xd7, 0x32, 0xc2, 0xcd, 0x5f, 0xa0, 0xda,
	0x63, 0xfe, 0x2d, 0xf1, 0x56, 0x74, 0x2d, 0x30, 0xa6, 0x02, 0xb1, 0x7f, 0xe6, 0xcc, 0xb7, 0x27,
	0x94, 0x4d, 0x24, 0x6d, 0x35, 0xc1, 0xbf, 0xe7, 0xcc, 0xff, 0x96, 0xb2, 0xc9, 0x3b, 0x60, 0xff,
	0x55, 0x85, 0xca, 0x80, 0x79, 0x3d, 0x46, 0x29, 0x9e, 0x86, 0x84, 0xf9, 0xdb, 0xc4, 0x7e, 0x0d,
	0x7b, 0x94, 0x79, 0x76, 0xb8, 0x5c, 0x24, 0x0f, 0x58, 0xed, 0x1c, 0xb7, 0x9f, 0xf7, 0x4f, 0x7b,
	0x2d, 0x53, 0x6c, 0x8d, 0x97, 0x0b, 0x6c, 0xe9, 0x34, 0x39, 0xa0, 0x53, 0xc8, 0x53, 0xe6, 0xf1,
	0xba, 0xd6, 0xd4, 0x5a, 0xe5, 0xce, 0xc7, 0x3b, 0x62, 0x2d, 0xe1, 0x9c, 0x11, 0x90, 0xcf, 0x08,
	0x68, 0xbc, 0x05, 0x6d, 0xc0, 0x3c, 0x84, 0x20, 0xef, 0x3a, 0xa1, 0x23, 0x2b, 0x16, 0x67, 0xf3,
	0x04, 0x74, 0x59, 0x02, 0x02, 0x28, 0x5a, 0xfd, 0xd1, 0xcd, 0x60, 0x6c, 0xe4, 0xe2, 0xf3, 0x68,
	0xdc, 0x1d, 0xdf, 0x8c, 0x0c, 0x05, 0x95, 0xa0, 0xd0, 0x3d, 0xeb, 0x0f, 0xc7, 0x86, 0x6a, 0xfe,
	0xa5, 0x40, 0xed, 0x3a, 0xc2, 0xc1, 0xf2, 0xd9, 0x5d, 0x7c, 0x05, 0xfa, 0x7d, 0x84, 0x03, 0x82,
	0x79, 0x5d, 0x11, 0x35, 0x67, 0xf4, 0x66, 0xfc, 0x13, 0xdb, 0x4a, 0x63, 0xfe, 0xff, 0xf3, 0x34,
	0x3e, 0x83, 0x82, 0xc8, 0x89, 0xaa, 0xa0, 0x12, 0x57, 0xaa, 0x53, 0x89, 0x8b, 0xde, 0x87, 0x42,
	0xcc, 0xb2, 0x94, 0x7d, 0x9c, 0x18, 0xe6, 0x1f, 0x1a, 0x18, 0x16, 0xe6, 0x11, 0x0d, 0x5f, 0xf7,
	0xa0, 0xdf, 0x80, 0x1e, 0x08, 0x77, 0x5e, 0x57, 0x85, 0xbe, 0x4f, 0xd6, 0xf5, 0x65, 0x73, 0x49,
	0xc0, 0x4a, 0xa3, 0x76, 0x95, 0xff, 0x9b, 0x0a, 0xc5, 0x24, 0x64, 0x43, 0x40, 0x0f, 0xf2, 0x01,
	0x7b, 0x4c, 0x79, 0x3f, 0x7f, 0x15, 0x6f, 0x4a, 0xcf, 0x1e, 0x2d, 0x11, 0x8c, 0x3e, 0x80, 0x22,
	0x0f, 0x9d, 0x30, 0xe2, 0x82, 0xba, 0x60, 0x49, 0xab, 0xf1, 0xbb, 0x02, 0xa5, 0x95, 0x2f, 0xba,
	0x02, 0x7d, 0xca, 0x68, 0x34, 0xf7, 0xd3, 0x57, 0xfc, 0xe2, 0x3f, 0xb2, 0xb5, 0x7b, 0x22, 0xdc,
	0x4a, 0xd3, 0x34, 0x3a, 0x50, 0x4c, 0xa0, 0xb8, 0xef, 0x7c, 0x67, 0x9e, 0x2e, 0x3b, 0x71, 0x8e,
	0xdf, 0xe6, 0xc1, 0xa1, 0x51, 0xba, 0xe4, 0x12, 0xc3, 0xfc, 0x53, 0x81, 0xf7, 0xbe, 0xc3, 0x0e,
	0x0d, 0x67, 0xbd, 0x19, 0x9e, 0xde, 0xad, 0xa6, 0xfd, 0x6c, 0xa5, 0x41, 0x11, 0x23, 0x95, 0xb9,
	0x8a, 0x7f, 0x09, 0x69, 0x8f, 0x70, 0xf0, 0x40, 0x7c, 0x6f, 0x24, 0xc2, 0x52, 0xd1, 0xe6, 0x97,
	0x50, 0x59, 0xfb, 0x80, 0xca, 0xa0, 0xdf, 0x0c, 0x2f, 0x86, 0x97, 0x3f, 0x0e, 0x8d, 0x5c, 0x6c,
	0x8c, 0xfa, 0xd6, 0x0f, 0xe7, 0xc3, 0x33, 0x43, 0x41, 0x35, 0x28, 0x0f, 0x2f, 0xc7, 0x76, 0x0a,
	0xa8, 0x9d, 0xbf, 0x35, 0xd0, 0xba, 0x0b, 0x82, 0x7e, 0x82, 0x37, 0x72, 0xb1, 0x3e, 0xad, 0x41,
	0x94, 0x19, 0xd4, 0x8d, 0x1d, 0xdc, 0x68, 0xbe, 0xec, 0x90, 0x54, 0x6c, 0xe6, 0xd0, 0x10, 0x2a,
	0xd2, 0x3d, 0xd9, 0x76, 0xe8, 0x60, 0x3d, 0x28, 0xf3, 0xbb, 0xd3, 0xf8, 0x68, 0xfd, 0xf3, 0xfa,
	0x8a, 0x34, 0x73, 0xe8, 0x0a, 0xaa, 0xd2, 0xf5, 0x5a, 0x4e, 0xdb, 0x8e, 0x84, 0x07, 0x5b, 0x47,
	0xd7, 0xcc, 0xa1, 0x01, 0x94, 0xaf, 0xa2, 0x09, 0x25, 0x7c, 0x36, 0x88, 0xf7, 0xce, 0x87, 0x5b,
	0xd6, 0x53, 0xe3, 0xf0, 0x25, 0xae, 0x55, 0x7d, 0x16, 0x54, 0x65, 0x36, 0x4b, 0xce, 0xca, 0xe1,
	0xf6, 0xae, 0x7b, 0x45, 0xce, 0x6b, 0x28, 0x8b, 0x46, 0x48, 0x7a, 0x62, 0x97, 0xe0, 0xa3, 0x9d,
	0x8d, 0x64, 0xe6, 0x26, 0x45, 0xf1, 0x37, 0xe0, 0xf4, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe6,
	0x24, 0x1d, 0x7b, 0x19, 0x08, 0x00, 0x00,
}
