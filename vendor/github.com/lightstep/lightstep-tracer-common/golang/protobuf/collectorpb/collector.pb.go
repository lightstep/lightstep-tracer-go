// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/lightstep/lightstep-tracer-common/tmpgen/collector.proto

/*
Package collectorpb is a generated protocol buffer package.

It is generated from these files:
	github.com/lightstep/lightstep-tracer-common/tmpgen/collector.proto

It has these top-level messages:
	SpanContext
	KeyValue
	Log
	Reference
	Span
	Reporter
	MetricsSample
	InternalMetrics
	Auth
	ReportRequest
	Command
	ReportResponse
*/
package collectorpb // import "github.com/lightstep/lightstep-tracer-common/golang/protobuf/collectorpb"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type Reference_Relationship int32

const (
	Reference_CHILD_OF     Reference_Relationship = 0
	Reference_FOLLOWS_FROM Reference_Relationship = 1
)

var Reference_Relationship_name = map[int32]string{
	0: "CHILD_OF",
	1: "FOLLOWS_FROM",
}
var Reference_Relationship_value = map[string]int32{
	"CHILD_OF":     0,
	"FOLLOWS_FROM": 1,
}

func (x Reference_Relationship) String() string {
	return proto.EnumName(Reference_Relationship_name, int32(x))
}
func (Reference_Relationship) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

type SpanContext struct {
	TraceId uint64            `protobuf:"varint,1,opt,name=trace_id,json=traceId" json:"trace_id,omitempty"`
	SpanId  uint64            `protobuf:"varint,2,opt,name=span_id,json=spanId" json:"span_id,omitempty"`
	Baggage map[string]string `protobuf:"bytes,3,rep,name=baggage" json:"baggage,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *SpanContext) Reset()                    { *m = SpanContext{} }
func (m *SpanContext) String() string            { return proto.CompactTextString(m) }
func (*SpanContext) ProtoMessage()               {}
func (*SpanContext) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SpanContext) GetTraceId() uint64 {
	if m != nil {
		return m.TraceId
	}
	return 0
}

func (m *SpanContext) GetSpanId() uint64 {
	if m != nil {
		return m.SpanId
	}
	return 0
}

func (m *SpanContext) GetBaggage() map[string]string {
	if m != nil {
		return m.Baggage
	}
	return nil
}

// Represent both tags and log fields.
type KeyValue struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*KeyValue_StringValue
	//	*KeyValue_IntValue
	//	*KeyValue_DoubleValue
	//	*KeyValue_BoolValue
	//	*KeyValue_JsonValue
	Value isKeyValue_Value `protobuf_oneof:"value"`
}

func (m *KeyValue) Reset()                    { *m = KeyValue{} }
func (m *KeyValue) String() string            { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()               {}
func (*KeyValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isKeyValue_Value interface {
	isKeyValue_Value()
}

type KeyValue_StringValue struct {
	StringValue string `protobuf:"bytes,2,opt,name=string_value,json=stringValue,oneof"`
}
type KeyValue_IntValue struct {
	IntValue int64 `protobuf:"varint,3,opt,name=int_value,json=intValue,oneof"`
}
type KeyValue_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,4,opt,name=double_value,json=doubleValue,oneof"`
}
type KeyValue_BoolValue struct {
	BoolValue bool `protobuf:"varint,5,opt,name=bool_value,json=boolValue,oneof"`
}
type KeyValue_JsonValue struct {
	JsonValue string `protobuf:"bytes,6,opt,name=json_value,json=jsonValue,oneof"`
}

func (*KeyValue_StringValue) isKeyValue_Value() {}
func (*KeyValue_IntValue) isKeyValue_Value()    {}
func (*KeyValue_DoubleValue) isKeyValue_Value() {}
func (*KeyValue_BoolValue) isKeyValue_Value()   {}
func (*KeyValue_JsonValue) isKeyValue_Value()   {}

func (m *KeyValue) GetValue() isKeyValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *KeyValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyValue) GetStringValue() string {
	if x, ok := m.GetValue().(*KeyValue_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *KeyValue) GetIntValue() int64 {
	if x, ok := m.GetValue().(*KeyValue_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (m *KeyValue) GetDoubleValue() float64 {
	if x, ok := m.GetValue().(*KeyValue_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (m *KeyValue) GetBoolValue() bool {
	if x, ok := m.GetValue().(*KeyValue_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (m *KeyValue) GetJsonValue() string {
	if x, ok := m.GetValue().(*KeyValue_JsonValue); ok {
		return x.JsonValue
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*KeyValue) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _KeyValue_OneofMarshaler, _KeyValue_OneofUnmarshaler, _KeyValue_OneofSizer, []interface{}{
		(*KeyValue_StringValue)(nil),
		(*KeyValue_IntValue)(nil),
		(*KeyValue_DoubleValue)(nil),
		(*KeyValue_BoolValue)(nil),
		(*KeyValue_JsonValue)(nil),
	}
}

func _KeyValue_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*KeyValue)
	// value
	switch x := m.Value.(type) {
	case *KeyValue_StringValue:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.StringValue)
	case *KeyValue_IntValue:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.IntValue))
	case *KeyValue_DoubleValue:
		b.EncodeVarint(4<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.DoubleValue))
	case *KeyValue_BoolValue:
		t := uint64(0)
		if x.BoolValue {
			t = 1
		}
		b.EncodeVarint(5<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *KeyValue_JsonValue:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.JsonValue)
	case nil:
	default:
		return fmt.Errorf("KeyValue.Value has unexpected type %T", x)
	}
	return nil
}

func _KeyValue_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*KeyValue)
	switch tag {
	case 2: // value.string_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &KeyValue_StringValue{x}
		return true, err
	case 3: // value.int_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &KeyValue_IntValue{int64(x)}
		return true, err
	case 4: // value.double_value
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Value = &KeyValue_DoubleValue{math.Float64frombits(x)}
		return true, err
	case 5: // value.bool_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &KeyValue_BoolValue{x != 0}
		return true, err
	case 6: // value.json_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &KeyValue_JsonValue{x}
		return true, err
	default:
		return false, nil
	}
}

func _KeyValue_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*KeyValue)
	// value
	switch x := m.Value.(type) {
	case *KeyValue_StringValue:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.StringValue)))
		n += len(x.StringValue)
	case *KeyValue_IntValue:
		n += proto.SizeVarint(3<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.IntValue))
	case *KeyValue_DoubleValue:
		n += proto.SizeVarint(4<<3 | proto.WireFixed64)
		n += 8
	case *KeyValue_BoolValue:
		n += proto.SizeVarint(5<<3 | proto.WireVarint)
		n += 1
	case *KeyValue_JsonValue:
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.JsonValue)))
		n += len(x.JsonValue)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Log struct {
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Fields    []*KeyValue                `protobuf:"bytes,2,rep,name=fields" json:"fields,omitempty"`
}

func (m *Log) Reset()                    { *m = Log{} }
func (m *Log) String() string            { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()               {}
func (*Log) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Log) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Log) GetFields() []*KeyValue {
	if m != nil {
		return m.Fields
	}
	return nil
}

type Reference struct {
	Relationship Reference_Relationship `protobuf:"varint,1,opt,name=relationship,enum=lightstep.collector.Reference_Relationship" json:"relationship,omitempty"`
	SpanContext  *SpanContext           `protobuf:"bytes,2,opt,name=span_context,json=spanContext" json:"span_context,omitempty"`
}

func (m *Reference) Reset()                    { *m = Reference{} }
func (m *Reference) String() string            { return proto.CompactTextString(m) }
func (*Reference) ProtoMessage()               {}
func (*Reference) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Reference) GetRelationship() Reference_Relationship {
	if m != nil {
		return m.Relationship
	}
	return Reference_CHILD_OF
}

func (m *Reference) GetSpanContext() *SpanContext {
	if m != nil {
		return m.SpanContext
	}
	return nil
}

type Span struct {
	SpanContext    *SpanContext               `protobuf:"bytes,1,opt,name=span_context,json=spanContext" json:"span_context,omitempty"`
	OperationName  string                     `protobuf:"bytes,2,opt,name=operation_name,json=operationName" json:"operation_name,omitempty"`
	References     []*Reference               `protobuf:"bytes,3,rep,name=references" json:"references,omitempty"`
	StartTimestamp *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=start_timestamp,json=startTimestamp" json:"start_timestamp,omitempty"`
	DurationMicros uint64                     `protobuf:"varint,5,opt,name=duration_micros,json=durationMicros" json:"duration_micros,omitempty"`
	Tags           []*KeyValue                `protobuf:"bytes,6,rep,name=tags" json:"tags,omitempty"`
	Logs           []*Log                     `protobuf:"bytes,7,rep,name=logs" json:"logs,omitempty"`
}

func (m *Span) Reset()                    { *m = Span{} }
func (m *Span) String() string            { return proto.CompactTextString(m) }
func (*Span) ProtoMessage()               {}
func (*Span) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Span) GetSpanContext() *SpanContext {
	if m != nil {
		return m.SpanContext
	}
	return nil
}

func (m *Span) GetOperationName() string {
	if m != nil {
		return m.OperationName
	}
	return ""
}

func (m *Span) GetReferences() []*Reference {
	if m != nil {
		return m.References
	}
	return nil
}

func (m *Span) GetStartTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.StartTimestamp
	}
	return nil
}

func (m *Span) GetDurationMicros() uint64 {
	if m != nil {
		return m.DurationMicros
	}
	return 0
}

func (m *Span) GetTags() []*KeyValue {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Span) GetLogs() []*Log {
	if m != nil {
		return m.Logs
	}
	return nil
}

type Reporter struct {
	ReporterId uint64      `protobuf:"varint,1,opt,name=reporter_id,json=reporterId" json:"reporter_id,omitempty"`
	Tags       []*KeyValue `protobuf:"bytes,4,rep,name=tags" json:"tags,omitempty"`
}

func (m *Reporter) Reset()                    { *m = Reporter{} }
func (m *Reporter) String() string            { return proto.CompactTextString(m) }
func (*Reporter) ProtoMessage()               {}
func (*Reporter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Reporter) GetReporterId() uint64 {
	if m != nil {
		return m.ReporterId
	}
	return 0
}

func (m *Reporter) GetTags() []*KeyValue {
	if m != nil {
		return m.Tags
	}
	return nil
}

type MetricsSample struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*MetricsSample_IntValue
	//	*MetricsSample_DoubleValue
	Value isMetricsSample_Value `protobuf_oneof:"value"`
}

func (m *MetricsSample) Reset()                    { *m = MetricsSample{} }
func (m *MetricsSample) String() string            { return proto.CompactTextString(m) }
func (*MetricsSample) ProtoMessage()               {}
func (*MetricsSample) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type isMetricsSample_Value interface {
	isMetricsSample_Value()
}

type MetricsSample_IntValue struct {
	IntValue int64 `protobuf:"varint,2,opt,name=int_value,json=intValue,oneof"`
}
type MetricsSample_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,3,opt,name=double_value,json=doubleValue,oneof"`
}

func (*MetricsSample_IntValue) isMetricsSample_Value()    {}
func (*MetricsSample_DoubleValue) isMetricsSample_Value() {}

func (m *MetricsSample) GetValue() isMetricsSample_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *MetricsSample) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MetricsSample) GetIntValue() int64 {
	if x, ok := m.GetValue().(*MetricsSample_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (m *MetricsSample) GetDoubleValue() float64 {
	if x, ok := m.GetValue().(*MetricsSample_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*MetricsSample) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _MetricsSample_OneofMarshaler, _MetricsSample_OneofUnmarshaler, _MetricsSample_OneofSizer, []interface{}{
		(*MetricsSample_IntValue)(nil),
		(*MetricsSample_DoubleValue)(nil),
	}
}

func _MetricsSample_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*MetricsSample)
	// value
	switch x := m.Value.(type) {
	case *MetricsSample_IntValue:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.IntValue))
	case *MetricsSample_DoubleValue:
		b.EncodeVarint(3<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.DoubleValue))
	case nil:
	default:
		return fmt.Errorf("MetricsSample.Value has unexpected type %T", x)
	}
	return nil
}

func _MetricsSample_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*MetricsSample)
	switch tag {
	case 2: // value.int_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &MetricsSample_IntValue{int64(x)}
		return true, err
	case 3: // value.double_value
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Value = &MetricsSample_DoubleValue{math.Float64frombits(x)}
		return true, err
	default:
		return false, nil
	}
}

func _MetricsSample_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*MetricsSample)
	// value
	switch x := m.Value.(type) {
	case *MetricsSample_IntValue:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.IntValue))
	case *MetricsSample_DoubleValue:
		n += proto.SizeVarint(3<<3 | proto.WireFixed64)
		n += 8
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type InternalMetrics struct {
	StartTimestamp *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=start_timestamp,json=startTimestamp" json:"start_timestamp,omitempty"`
	DurationMicros uint64                     `protobuf:"varint,2,opt,name=duration_micros,json=durationMicros" json:"duration_micros,omitempty"`
	Logs           []*Log                     `protobuf:"bytes,3,rep,name=logs" json:"logs,omitempty"`
	Counts         []*MetricsSample           `protobuf:"bytes,4,rep,name=counts" json:"counts,omitempty"`
	Gauges         []*MetricsSample           `protobuf:"bytes,5,rep,name=gauges" json:"gauges,omitempty"`
}

func (m *InternalMetrics) Reset()                    { *m = InternalMetrics{} }
func (m *InternalMetrics) String() string            { return proto.CompactTextString(m) }
func (*InternalMetrics) ProtoMessage()               {}
func (*InternalMetrics) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *InternalMetrics) GetStartTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.StartTimestamp
	}
	return nil
}

func (m *InternalMetrics) GetDurationMicros() uint64 {
	if m != nil {
		return m.DurationMicros
	}
	return 0
}

func (m *InternalMetrics) GetLogs() []*Log {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *InternalMetrics) GetCounts() []*MetricsSample {
	if m != nil {
		return m.Counts
	}
	return nil
}

func (m *InternalMetrics) GetGauges() []*MetricsSample {
	if m != nil {
		return m.Gauges
	}
	return nil
}

type Auth struct {
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
}

func (m *Auth) Reset()                    { *m = Auth{} }
func (m *Auth) String() string            { return proto.CompactTextString(m) }
func (*Auth) ProtoMessage()               {}
func (*Auth) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Auth) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type ReportRequest struct {
	Reporter              *Reporter        `protobuf:"bytes,1,opt,name=reporter" json:"reporter,omitempty"`
	Auth                  *Auth            `protobuf:"bytes,2,opt,name=auth" json:"auth,omitempty"`
	Spans                 []*Span          `protobuf:"bytes,3,rep,name=spans" json:"spans,omitempty"`
	TimestampOffsetMicros int64            `protobuf:"varint,5,opt,name=timestamp_offset_micros,json=timestampOffsetMicros" json:"timestamp_offset_micros,omitempty"`
	InternalMetrics       *InternalMetrics `protobuf:"bytes,6,opt,name=internal_metrics,json=internalMetrics" json:"internal_metrics,omitempty"`
}

func (m *ReportRequest) Reset()                    { *m = ReportRequest{} }
func (m *ReportRequest) String() string            { return proto.CompactTextString(m) }
func (*ReportRequest) ProtoMessage()               {}
func (*ReportRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ReportRequest) GetReporter() *Reporter {
	if m != nil {
		return m.Reporter
	}
	return nil
}

func (m *ReportRequest) GetAuth() *Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *ReportRequest) GetSpans() []*Span {
	if m != nil {
		return m.Spans
	}
	return nil
}

func (m *ReportRequest) GetTimestampOffsetMicros() int64 {
	if m != nil {
		return m.TimestampOffsetMicros
	}
	return 0
}

func (m *ReportRequest) GetInternalMetrics() *InternalMetrics {
	if m != nil {
		return m.InternalMetrics
	}
	return nil
}

type Command struct {
	Disable bool `protobuf:"varint,1,opt,name=disable" json:"disable,omitempty"`
	DevMode bool `protobuf:"varint,2,opt,name=dev_mode,json=devMode" json:"dev_mode,omitempty"`
}

func (m *Command) Reset()                    { *m = Command{} }
func (m *Command) String() string            { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()               {}
func (*Command) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Command) GetDisable() bool {
	if m != nil {
		return m.Disable
	}
	return false
}

func (m *Command) GetDevMode() bool {
	if m != nil {
		return m.DevMode
	}
	return false
}

type ReportResponse struct {
	Commands          []*Command                 `protobuf:"bytes,1,rep,name=commands" json:"commands,omitempty"`
	ReceiveTimestamp  *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=receive_timestamp,json=receiveTimestamp" json:"receive_timestamp,omitempty"`
	TransmitTimestamp *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=transmit_timestamp,json=transmitTimestamp" json:"transmit_timestamp,omitempty"`
	Errors            []string                   `protobuf:"bytes,4,rep,name=errors" json:"errors,omitempty"`
	Warnings          []string                   `protobuf:"bytes,5,rep,name=warnings" json:"warnings,omitempty"`
	Infos             []string                   `protobuf:"bytes,6,rep,name=infos" json:"infos,omitempty"`
}

func (m *ReportResponse) Reset()                    { *m = ReportResponse{} }
func (m *ReportResponse) String() string            { return proto.CompactTextString(m) }
func (*ReportResponse) ProtoMessage()               {}
func (*ReportResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ReportResponse) GetCommands() []*Command {
	if m != nil {
		return m.Commands
	}
	return nil
}

func (m *ReportResponse) GetReceiveTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.ReceiveTimestamp
	}
	return nil
}

func (m *ReportResponse) GetTransmitTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.TransmitTimestamp
	}
	return nil
}

func (m *ReportResponse) GetErrors() []string {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *ReportResponse) GetWarnings() []string {
	if m != nil {
		return m.Warnings
	}
	return nil
}

func (m *ReportResponse) GetInfos() []string {
	if m != nil {
		return m.Infos
	}
	return nil
}

func init() {
	proto.RegisterType((*SpanContext)(nil), "lightstep.collector.SpanContext")
	proto.RegisterType((*KeyValue)(nil), "lightstep.collector.KeyValue")
	proto.RegisterType((*Log)(nil), "lightstep.collector.Log")
	proto.RegisterType((*Reference)(nil), "lightstep.collector.Reference")
	proto.RegisterType((*Span)(nil), "lightstep.collector.Span")
	proto.RegisterType((*Reporter)(nil), "lightstep.collector.Reporter")
	proto.RegisterType((*MetricsSample)(nil), "lightstep.collector.MetricsSample")
	proto.RegisterType((*InternalMetrics)(nil), "lightstep.collector.InternalMetrics")
	proto.RegisterType((*Auth)(nil), "lightstep.collector.Auth")
	proto.RegisterType((*ReportRequest)(nil), "lightstep.collector.ReportRequest")
	proto.RegisterType((*Command)(nil), "lightstep.collector.Command")
	proto.RegisterType((*ReportResponse)(nil), "lightstep.collector.ReportResponse")
	proto.RegisterEnum("lightstep.collector.Reference_Relationship", Reference_Relationship_name, Reference_Relationship_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CollectorService service

type CollectorServiceClient interface {
	Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportResponse, error)
}

type collectorServiceClient struct {
	cc *grpc.ClientConn
}

func NewCollectorServiceClient(cc *grpc.ClientConn) CollectorServiceClient {
	return &collectorServiceClient{cc}
}

func (c *collectorServiceClient) Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportResponse, error) {
	out := new(ReportResponse)
	err := grpc.Invoke(ctx, "/lightstep.collector.CollectorService/Report", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CollectorService service

type CollectorServiceServer interface {
	Report(context.Context, *ReportRequest) (*ReportResponse, error)
}

func RegisterCollectorServiceServer(s *grpc.Server, srv CollectorServiceServer) {
	s.RegisterService(&_CollectorService_serviceDesc, srv)
}

func _CollectorService_Report_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectorServiceServer).Report(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lightstep.collector.CollectorService/Report",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectorServiceServer).Report(ctx, req.(*ReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CollectorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lightstep.collector.CollectorService",
	HandlerType: (*CollectorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Report",
			Handler:    _CollectorService_Report_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/lightstep/lightstep-tracer-common/tmpgen/collector.proto",
}

func init() {
	proto.RegisterFile("github.com/lightstep/lightstep-tracer-common/tmpgen/collector.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 1116 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0xce, 0xae, 0x5d, 0xff, 0x1c, 0xbb, 0x89, 0x3b, 0xfc, 0x39, 0xa6, 0xa5, 0xee, 0x06, 0x44,
	0xa0, 0x64, 0x4d, 0x8d, 0x90, 0x4a, 0x2e, 0x2a, 0x11, 0x43, 0x1a, 0x0b, 0x07, 0x57, 0x93, 0x0a,
	0xa4, 0xde, 0x58, 0xe3, 0xdd, 0xf1, 0x66, 0xe9, 0xee, 0xcc, 0x32, 0x33, 0x36, 0xe4, 0x0e, 0x21,
	0xf1, 0x04, 0x88, 0x17, 0xe0, 0x0a, 0x9e, 0x80, 0xb7, 0x40, 0x42, 0xf0, 0x08, 0x3c, 0x08, 0xda,
	0x99, 0xdd, 0xb5, 0x53, 0x9c, 0x1f, 0x21, 0xee, 0x76, 0xce, 0x7c, 0xdf, 0x9c, 0x33, 0xe7, 0xfb,
	0xce, 0xd8, 0x30, 0x08, 0x42, 0x75, 0x3a, 0x9f, 0xba, 0x1e, 0x8f, 0x7b, 0x51, 0x18, 0x9c, 0x2a,
	0xa9, 0x68, 0xb2, 0xfc, 0xda, 0x53, 0x82, 0x78, 0x54, 0xec, 0x79, 0x3c, 0x8e, 0x39, 0xeb, 0xa9,
	0x38, 0x09, 0x28, 0xeb, 0x79, 0x3c, 0x8a, 0xa8, 0xa7, 0xb8, 0x70, 0x13, 0xc1, 0x15, 0x47, 0x2f,
	0x15, 0x78, 0xb7, 0xd8, 0xea, 0xdc, 0x0d, 0x38, 0x0f, 0x22, 0xda, 0xd3, 0x90, 0xe9, 0x7c, 0xd6,
	0x53, 0x61, 0x4c, 0xa5, 0x22, 0x71, 0x62, 0x58, 0x9d, 0xdb, 0x19, 0x80, 0x24, 0x61, 0x8f, 0x30,
	0xc6, 0x15, 0x51, 0x21, 0x67, 0xd2, 0xec, 0x3a, 0xbf, 0x5b, 0xd0, 0x38, 0x49, 0x08, 0x1b, 0x70,
	0xa6, 0xe8, 0xb7, 0x0a, 0xdd, 0x81, 0x9a, 0xae, 0x65, 0x12, 0xfa, 0x6d, 0xab, 0x6b, 0xed, 0x96,
	0x0f, 0xec, 0xf7, 0x2d, 0x5c, 0xd5, 0xb1, 0xa1, 0x8f, 0x5e, 0x87, 0xaa, 0x4c, 0x08, 0x4b, 0x77,
	0xed, 0x62, 0xb7, 0x92, 0x86, 0x86, 0x3e, 0x7a, 0x0c, 0xd5, 0x29, 0x09, 0x02, 0x12, 0xd0, 0x76,
	0xa9, 0x5b, 0xda, 0x6d, 0xf4, 0xf7, 0xdc, 0x35, 0x15, 0xbb, 0x2b, 0xe9, 0xdc, 0x03, 0x83, 0xff,
	0x94, 0x29, 0x71, 0x86, 0x73, 0x76, 0x67, 0x1f, 0x9a, 0xab, 0x1b, 0xa8, 0x05, 0xa5, 0xe7, 0xf4,
	0x4c, 0xd7, 0x53, 0xc7, 0xe9, 0x27, 0x7a, 0x19, 0x6e, 0x2c, 0x48, 0x34, 0xa7, 0xba, 0x8a, 0x3a,
	0x36, 0x8b, 0x7d, 0xfb, 0xa1, 0xe5, 0xfc, 0x65, 0x41, 0xed, 0x33, 0x7a, 0xf6, 0x45, 0x1a, 0x58,
	0x43, 0xdc, 0x81, 0xa6, 0x54, 0x22, 0x64, 0xc1, 0x64, 0x85, 0x7f, 0xb4, 0x81, 0x1b, 0x26, 0x6a,
	0x68, 0xf7, 0xa0, 0x1e, 0x32, 0x95, 0x21, 0x4a, 0x5d, 0x6b, 0xb7, 0x94, 0xde, 0xf3, 0x68, 0x03,
	0xd7, 0x42, 0xa6, 0x0c, 0x64, 0x07, 0x9a, 0x3e, 0x9f, 0x4f, 0x23, 0x9a, 0xa1, 0xca, 0x5d, 0x6b,
	0x37, 0x45, 0x34, 0x4c, 0xd4, 0x80, 0xee, 0x02, 0x4c, 0x39, 0x8f, 0x32, 0xc8, 0x8d, 0xae, 0xb5,
	0x5b, 0x3b, 0xda, 0xc0, 0xf5, 0x34, 0x56, 0x00, 0xbe, 0x92, 0x9c, 0x65, 0x80, 0x4a, 0x56, 0x4b,
	0x3d, 0x8d, 0x69, 0xc0, 0x41, 0x35, 0xbb, 0xa7, 0xb3, 0x80, 0xd2, 0x88, 0x07, 0xe8, 0x21, 0xd4,
	0x0b, 0x7d, 0xf5, 0xb5, 0x1a, 0xfd, 0x8e, 0x6b, 0x04, 0x76, 0x73, 0x07, 0xb8, 0x4f, 0x73, 0x04,
	0x5e, 0x82, 0xd1, 0x87, 0x50, 0x99, 0x85, 0x34, 0xf2, 0x65, 0xdb, 0xd6, 0xda, 0xdc, 0x59, 0xab,
	0x4d, 0xde, 0x39, 0x9c, 0x81, 0x9d, 0x3f, 0x2c, 0xa8, 0x63, 0x3a, 0xa3, 0x82, 0x32, 0x8f, 0xa2,
	0x31, 0x34, 0x05, 0x8d, 0x8c, 0x81, 0x4e, 0x43, 0x53, 0xc1, 0x66, 0xff, 0xfe, 0xda, 0xa3, 0x0a,
	0x96, 0x8b, 0x57, 0x28, 0xf8, 0xdc, 0x01, 0x68, 0x00, 0x4d, 0xed, 0x27, 0xcf, 0xf8, 0x41, 0xcb,
	0xd1, 0xe8, 0x77, 0xaf, 0xf2, 0x0d, 0x6e, 0xc8, 0xe5, 0xc2, 0x71, 0xa1, 0xb9, 0x9a, 0x02, 0x35,
	0xa1, 0x36, 0x38, 0x1a, 0x8e, 0x3e, 0x99, 0x8c, 0x0f, 0x5b, 0x1b, 0xa8, 0x05, 0xcd, 0xc3, 0xf1,
	0x68, 0x34, 0xfe, 0xf2, 0x64, 0x72, 0x88, 0xc7, 0xc7, 0x2d, 0xcb, 0xf9, 0xa1, 0x04, 0xe5, 0xf4,
	0xb0, 0x7f, 0x65, 0xb7, 0xfe, 0x43, 0x76, 0xf4, 0x16, 0x6c, 0xf2, 0x84, 0x0a, 0x9d, 0x7e, 0xc2,
	0x48, 0x9c, 0x7b, 0xf2, 0x66, 0x11, 0xfd, 0x9c, 0xc4, 0x14, 0x3d, 0x02, 0x10, 0x79, 0x47, 0x64,
	0x36, 0x1f, 0x6f, 0x5c, 0xde, 0x38, 0xbc, 0xc2, 0x40, 0x03, 0xd8, 0x92, 0x8a, 0x08, 0x35, 0x59,
	0xea, 0x5f, 0xbe, 0x52, 0xff, 0x4d, 0x4d, 0x29, 0xd6, 0xe8, 0x3e, 0x6c, 0xf9, 0xf3, 0xac, 0xd4,
	0x38, 0xf4, 0x04, 0x97, 0xda, 0x95, 0x66, 0x8c, 0x37, 0xf3, 0xad, 0x63, 0xbd, 0x83, 0x1e, 0x40,
	0x59, 0x91, 0x40, 0xb6, 0x2b, 0xd7, 0xf1, 0x8b, 0x86, 0xa2, 0xf7, 0xa0, 0x1c, 0xf1, 0x40, 0xb6,
	0xab, 0x9a, 0xd2, 0x5e, 0x4b, 0x19, 0xf1, 0x00, 0x6b, 0x94, 0x33, 0x85, 0x1a, 0xa6, 0x09, 0x17,
	0x8a, 0x0a, 0xb4, 0x03, 0x0d, 0x91, 0x7d, 0x9f, 0x7f, 0x7a, 0x20, 0x0f, 0x0f, 0xfd, 0xa2, 0xa2,
	0xf2, 0xb5, 0x2b, 0x72, 0x14, 0xdc, 0x3c, 0xa6, 0x4a, 0x84, 0x9e, 0x3c, 0x21, 0x71, 0x12, 0x51,
	0x84, 0xa0, 0xac, 0x45, 0x32, 0x6f, 0x82, 0xfe, 0x3e, 0x3f, 0xef, 0xf6, 0xb5, 0xe6, 0xbd, 0xb4,
	0x66, 0xde, 0x97, 0xd3, 0xfa, 0xab, 0x0d, 0x5b, 0x43, 0xa6, 0xa8, 0x60, 0x24, 0xca, 0xd2, 0xaf,
	0x13, 0xd0, 0xfa, 0x3f, 0x04, 0xb4, 0x2f, 0x14, 0x30, 0x57, 0xa3, 0x74, 0x1d, 0x35, 0xd0, 0x3e,
	0x54, 0x3c, 0x3e, 0x67, 0x2a, 0x6f, 0xaf, 0xb3, 0x16, 0x7f, 0xae, 0x99, 0x38, 0x63, 0xa4, 0xdc,
	0x80, 0xcc, 0x03, 0x9a, 0xda, 0xe9, 0xda, 0x5c, 0xc3, 0x70, 0xde, 0x81, 0xf2, 0xc7, 0x73, 0x75,
	0x8a, 0xee, 0x41, 0x93, 0x78, 0x1e, 0x95, 0x72, 0xa2, 0xf8, 0x73, 0xca, 0x32, 0x81, 0x1a, 0x26,
	0xf6, 0x34, 0x0d, 0x39, 0xbf, 0xd9, 0x70, 0xd3, 0x38, 0x06, 0xd3, 0xaf, 0xe7, 0x54, 0x2a, 0xf4,
	0x11, 0xd4, 0x72, 0x7f, 0x64, 0xdd, 0xbc, 0x73, 0xc1, 0x4c, 0x19, 0x10, 0x2e, 0xe0, 0x68, 0x0f,
	0xca, 0x64, 0xae, 0x4e, 0xb3, 0x27, 0x67, 0x7b, 0x2d, 0x2d, 0x2d, 0x0c, 0x6b, 0x18, 0xea, 0xc1,
	0x8d, 0x74, 0xea, 0xf3, 0x6e, 0x6e, 0x5f, 0xf8, 0x48, 0x60, 0x83, 0x43, 0xfb, 0xf0, 0x5a, 0xa1,
	0xf4, 0x84, 0xcf, 0x66, 0x92, 0xaa, 0xd5, 0x99, 0xd3, 0x16, 0xc3, 0xaf, 0x14, 0x90, 0xb1, 0x46,
	0x64, 0xca, 0x8d, 0xa1, 0x15, 0x66, 0xf6, 0x99, 0xc4, 0xa6, 0x6b, 0xfa, 0xd7, 0xa1, 0xd1, 0x7f,
	0x73, 0x6d, 0xde, 0x17, 0xbc, 0x86, 0xb7, 0xc2, 0xf3, 0x01, 0xe7, 0x11, 0x54, 0x07, 0x3c, 0x8e,
	0x09, 0xf3, 0x51, 0x1b, 0xaa, 0x7e, 0x28, 0xc9, 0x34, 0x32, 0x33, 0x50, 0xc3, 0xf9, 0x12, 0x6d,
	0x43, 0xcd, 0xa7, 0x8b, 0x49, 0xcc, 0x7d, 0x33, 0x05, 0xe9, 0x16, 0x5d, 0x1c, 0x73, 0x9f, 0x3a,
	0xbf, 0xd8, 0xb0, 0x99, 0x77, 0x5e, 0x26, 0x9c, 0x49, 0x8a, 0x1e, 0x42, 0xcd, 0x33, 0x47, 0xca,
	0xb6, 0xa5, 0x7b, 0x72, 0x7b, 0x6d, 0x6d, 0x59, 0x5e, 0x5c, 0xa0, 0xd1, 0x63, 0xb8, 0x25, 0xa8,
	0x47, 0xc3, 0x05, 0x5d, 0x99, 0x05, 0xfb, 0xca, 0x59, 0x68, 0x65, 0xa4, 0xe5, 0x34, 0x0c, 0x01,
	0x29, 0x41, 0x98, 0x8c, 0xc3, 0xd5, 0xa9, 0x2a, 0x5d, 0x79, 0xd2, 0xad, 0x9c, 0xb5, 0x3c, 0xea,
	0x55, 0xa8, 0x50, 0x21, 0xb8, 0x30, 0xee, 0xaf, 0xe3, 0x6c, 0x85, 0x3a, 0x50, 0xfb, 0x86, 0x08,
	0x16, 0xb2, 0xc0, 0x78, 0xbb, 0x8e, 0x8b, 0x75, 0xfa, 0x27, 0x24, 0x64, 0x33, 0x6e, 0x5e, 0xc8,
	0x3a, 0x36, 0x8b, 0xfe, 0x4f, 0x16, 0xb4, 0x06, 0xf9, 0xed, 0x4f, 0xa8, 0x58, 0x84, 0x1e, 0x45,
	0xdf, 0x59, 0x50, 0x31, 0xfd, 0x43, 0xce, 0x25, 0x06, 0xcd, 0x6c, 0xdd, 0xd9, 0xb9, 0x14, 0x63,
	0x04, 0x70, 0xf6, 0xbe, 0xff, 0xf3, 0xef, 0x1f, 0xed, 0xb7, 0x9d, 0x2d, 0xfd, 0xd7, 0x6e, 0xd1,
	0xef, 0x19, 0x6b, 0xcb, 0x7d, 0xeb, 0xdd, 0x67, 0xb7, 0xd0, 0x8b, 0xd1, 0x83, 0x07, 0xb0, 0xed,
	0xf1, 0x78, 0xe5, 0x60, 0xf3, 0x9f, 0xd3, 0x0d, 0x44, 0xe2, 0x3d, 0xb1, 0x9e, 0x35, 0x8a, 0x3c,
	0xc9, 0xf4, 0x67, 0xbb, 0x3c, 0x3a, 0x79, 0x72, 0x30, 0xad, 0xe8, 0xde, 0x7d, 0xf0, 0x4f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xfb, 0x9f, 0x66, 0xc4, 0xbe, 0x0a, 0x00, 0x00,
}