// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/sb/e2ap.proto

package sb

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Criticality int32

const (
	Criticality_REJECT Criticality = 0
	Criticality_IGNORE Criticality = 1
	Criticality_NOTIFY Criticality = 2
)

var Criticality_name = map[int32]string{
	0: "REJECT",
	1: "IGNORE",
	2: "NOTIFY",
}

var Criticality_value = map[string]int32{
	"REJECT": 0,
	"IGNORE": 1,
	"NOTIFY": 2,
}

func (x Criticality) String() string {
	return proto.EnumName(Criticality_name, int32(x))
}

func (Criticality) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e888e9637db470e9, []int{0}
}

type Presence int32

const (
	Presence_OPTIONAL    Presence = 0
	Presence_CONDITIONAL Presence = 1
	Presence_MANDATORY   Presence = 2
)

var Presence_name = map[int32]string{
	0: "OPTIONAL",
	1: "CONDITIONAL",
	2: "MANDATORY",
}

var Presence_value = map[string]int32{
	"OPTIONAL":    0,
	"CONDITIONAL": 1,
	"MANDATORY":   2,
}

func (x Presence) String() string {
	return proto.EnumName(Presence_name, int32(x))
}

func (Presence) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e888e9637db470e9, []int{1}
}

type TriggeringMessage int32

const (
	TriggeringMessage_INITIATING_MESSAGE    TriggeringMessage = 0
	TriggeringMessage_SUCCESSFUL_OUTCOME    TriggeringMessage = 1
	TriggeringMessage_UNSUCCESSFULL_OUTCOME TriggeringMessage = 2
)

var TriggeringMessage_name = map[int32]string{
	0: "INITIATING_MESSAGE",
	1: "SUCCESSFUL_OUTCOME",
	2: "UNSUCCESSFULL_OUTCOME",
}

var TriggeringMessage_value = map[string]int32{
	"INITIATING_MESSAGE":    0,
	"SUCCESSFUL_OUTCOME":    1,
	"UNSUCCESSFULL_OUTCOME": 2,
}

func (x TriggeringMessage) String() string {
	return proto.EnumName(TriggeringMessage_name, int32(x))
}

func (TriggeringMessage) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e888e9637db470e9, []int{2}
}

type ProcedureCode int32

const (
	ProcedureCode_PC_INVALID              ProcedureCode = 0
	ProcedureCode_E2_SETUP                ProcedureCode = 1
	ProcedureCode_ERROR_INDICATION        ProcedureCode = 2
	ProcedureCode_RESET                   ProcedureCode = 3
	ProcedureCode_RIC_CONTROL             ProcedureCode = 4
	ProcedureCode_RIC_INDICATION          ProcedureCode = 5
	ProcedureCode_RIC_SERVICE_QUERY       ProcedureCode = 6
	ProcedureCode_RIC_SERVICE_UPDATE      ProcedureCode = 7
	ProcedureCode_RIC_SUBSCRIPTION        ProcedureCode = 8
	ProcedureCode_RIC_SUBSCRIPTION_DELETE ProcedureCode = 9
)

var ProcedureCode_name = map[int32]string{
	0: "PC_INVALID",
	1: "E2_SETUP",
	2: "ERROR_INDICATION",
	3: "RESET",
	4: "RIC_CONTROL",
	5: "RIC_INDICATION",
	6: "RIC_SERVICE_QUERY",
	7: "RIC_SERVICE_UPDATE",
	8: "RIC_SUBSCRIPTION",
	9: "RIC_SUBSCRIPTION_DELETE",
}

var ProcedureCode_value = map[string]int32{
	"PC_INVALID":              0,
	"E2_SETUP":                1,
	"ERROR_INDICATION":        2,
	"RESET":                   3,
	"RIC_CONTROL":             4,
	"RIC_INDICATION":          5,
	"RIC_SERVICE_QUERY":       6,
	"RIC_SERVICE_UPDATE":      7,
	"RIC_SUBSCRIPTION":        8,
	"RIC_SUBSCRIPTION_DELETE": 9,
}

func (x ProcedureCode) String() string {
	return proto.EnumName(ProcedureCode_name, int32(x))
}

func (ProcedureCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e888e9637db470e9, []int{3}
}

type ProtocolIEId int32

const (
	ProtocolIEId_UNDEFINED                    ProtocolIEId = 0
	ProtocolIEId_CAUSE                        ProtocolIEId = 1
	ProtocolIEId_CRITICALITY_DIAGNOSTICS      ProtocolIEId = 2
	ProtocolIEId_GLOBAL_E2_NODE_ID            ProtocolIEId = 3
	ProtocolIEId_GLOBAL_RIC_ID                ProtocolIEId = 4
	ProtocolIEId_RAN_FUNCTION_ID              ProtocolIEId = 5
	ProtocolIEId_RAN_FUNCTION_ID_ITEM         ProtocolIEId = 6
	ProtocolIEId_RAN_FUNCTION_IE_CAUSE_ITEM   ProtocolIEId = 7
	ProtocolIEId_RAN_FUNCTION_ITEM            ProtocolIEId = 8
	ProtocolIEId_RAN_FUNCTIONS_ACCEPTED       ProtocolIEId = 9
	ProtocolIEId_RAN_FUNCTIONS_ADDED          ProtocolIEId = 10
	ProtocolIEId_RAN_FUNCTIONS_DELETED        ProtocolIEId = 11
	ProtocolIEId_RAN_FUNCTIONS_MODIFIED       ProtocolIEId = 12
	ProtocolIEId_RAN_FUNCTIONS_REJECTED       ProtocolIEId = 13
	ProtocolIEId_RIC_ACTION_ADMITTED_ITEM     ProtocolIEId = 14
	ProtocolIEId_RIC_ACTION_ID                ProtocolIEId = 15
	ProtocolIEId_RIC_ACTION_NOT_ADMITTED_ITEM ProtocolIEId = 16
	ProtocolIEId_RIC_ACTIONS_ADMITTED         ProtocolIEId = 17
	ProtocolIEId_RIC_ACTIONS_NOT_ADMITTED     ProtocolIEId = 18
	ProtocolIEId_RIC_ACTION_TO_BE_SETUP_ITEM  ProtocolIEId = 19
	ProtocolIEId_RIC_CALL_PROCESS_ID          ProtocolIEId = 20
	ProtocolIEId_RIC_CONTROL_ACK_REQUEST      ProtocolIEId = 21
	ProtocolIEId_RIC_CONTROL_HEADER           ProtocolIEId = 22
	ProtocolIEId_RIC_CONTROL_MESSAGE          ProtocolIEId = 23
	ProtocolIEId_RIC_CONTROL_STATUS           ProtocolIEId = 24
	ProtocolIEId_RIC_INDICATION_HEADER        ProtocolIEId = 25
	ProtocolIEId_RIC_INDICATION_MESSAGE       ProtocolIEId = 26
	ProtocolIEId_RIC_INDICATION_SN            ProtocolIEId = 27
	ProtocolIEId_RIC_INDICATION_TYPE          ProtocolIEId = 28
	ProtocolIEId_RIC_REQUEST_ID               ProtocolIEId = 29
	ProtocolIEId_RIC_SUBSCRIPTION_DETAILS     ProtocolIEId = 30
	ProtocolIEId_TIME_TO_WAIT                 ProtocolIEId = 31
	ProtocolIEId_RIC_CONTROL_OUTCOME          ProtocolIEId = 32
)

var ProtocolIEId_name = map[int32]string{
	0:  "UNDEFINED",
	1:  "CAUSE",
	2:  "CRITICALITY_DIAGNOSTICS",
	3:  "GLOBAL_E2_NODE_ID",
	4:  "GLOBAL_RIC_ID",
	5:  "RAN_FUNCTION_ID",
	6:  "RAN_FUNCTION_ID_ITEM",
	7:  "RAN_FUNCTION_IE_CAUSE_ITEM",
	8:  "RAN_FUNCTION_ITEM",
	9:  "RAN_FUNCTIONS_ACCEPTED",
	10: "RAN_FUNCTIONS_ADDED",
	11: "RAN_FUNCTIONS_DELETED",
	12: "RAN_FUNCTIONS_MODIFIED",
	13: "RAN_FUNCTIONS_REJECTED",
	14: "RIC_ACTION_ADMITTED_ITEM",
	15: "RIC_ACTION_ID",
	16: "RIC_ACTION_NOT_ADMITTED_ITEM",
	17: "RIC_ACTIONS_ADMITTED",
	18: "RIC_ACTIONS_NOT_ADMITTED",
	19: "RIC_ACTION_TO_BE_SETUP_ITEM",
	20: "RIC_CALL_PROCESS_ID",
	21: "RIC_CONTROL_ACK_REQUEST",
	22: "RIC_CONTROL_HEADER",
	23: "RIC_CONTROL_MESSAGE",
	24: "RIC_CONTROL_STATUS",
	25: "RIC_INDICATION_HEADER",
	26: "RIC_INDICATION_MESSAGE",
	27: "RIC_INDICATION_SN",
	28: "RIC_INDICATION_TYPE",
	29: "RIC_REQUEST_ID",
	30: "RIC_SUBSCRIPTION_DETAILS",
	31: "TIME_TO_WAIT",
	32: "RIC_CONTROL_OUTCOME",
}

var ProtocolIEId_value = map[string]int32{
	"UNDEFINED":                    0,
	"CAUSE":                        1,
	"CRITICALITY_DIAGNOSTICS":      2,
	"GLOBAL_E2_NODE_ID":            3,
	"GLOBAL_RIC_ID":                4,
	"RAN_FUNCTION_ID":              5,
	"RAN_FUNCTION_ID_ITEM":         6,
	"RAN_FUNCTION_IE_CAUSE_ITEM":   7,
	"RAN_FUNCTION_ITEM":            8,
	"RAN_FUNCTIONS_ACCEPTED":       9,
	"RAN_FUNCTIONS_ADDED":          10,
	"RAN_FUNCTIONS_DELETED":        11,
	"RAN_FUNCTIONS_MODIFIED":       12,
	"RAN_FUNCTIONS_REJECTED":       13,
	"RIC_ACTION_ADMITTED_ITEM":     14,
	"RIC_ACTION_ID":                15,
	"RIC_ACTION_NOT_ADMITTED_ITEM": 16,
	"RIC_ACTIONS_ADMITTED":         17,
	"RIC_ACTIONS_NOT_ADMITTED":     18,
	"RIC_ACTION_TO_BE_SETUP_ITEM":  19,
	"RIC_CALL_PROCESS_ID":          20,
	"RIC_CONTROL_ACK_REQUEST":      21,
	"RIC_CONTROL_HEADER":           22,
	"RIC_CONTROL_MESSAGE":          23,
	"RIC_CONTROL_STATUS":           24,
	"RIC_INDICATION_HEADER":        25,
	"RIC_INDICATION_MESSAGE":       26,
	"RIC_INDICATION_SN":            27,
	"RIC_INDICATION_TYPE":          28,
	"RIC_REQUEST_ID":               29,
	"RIC_SUBSCRIPTION_DETAILS":     30,
	"TIME_TO_WAIT":                 31,
	"RIC_CONTROL_OUTCOME":          32,
}

func (x ProtocolIEId) String() string {
	return proto.EnumName(ProtocolIEId_name, int32(x))
}

func (ProtocolIEId) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e888e9637db470e9, []int{4}
}

type E2ApProtocolIE struct {
	Id          ProtocolIEId `protobuf:"varint,1,opt,name=id,proto3,enum=interface.e2ap.ProtocolIEId" json:"id,omitempty"`
	Criticality Criticality  `protobuf:"varint,2,opt,name=criticality,proto3,enum=interface.e2ap.Criticality" json:"criticality,omitempty"`
	// value
	Presence Presence `protobuf:"varint,3,opt,name=presence,proto3,enum=interface.e2ap.Presence" json:"presence,omitempty"`
}

func (m *E2ApProtocolIE) Reset()         { *m = E2ApProtocolIE{} }
func (m *E2ApProtocolIE) String() string { return proto.CompactTextString(m) }
func (*E2ApProtocolIE) ProtoMessage()    {}
func (*E2ApProtocolIE) Descriptor() ([]byte, []int) {
	return fileDescriptor_e888e9637db470e9, []int{0}
}
func (m *E2ApProtocolIE) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *E2ApProtocolIE) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_E2ApProtocolIE.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *E2ApProtocolIE) XXX_Merge(src proto.Message) {
	xxx_messageInfo_E2ApProtocolIE.Merge(m, src)
}
func (m *E2ApProtocolIE) XXX_Size() int {
	return m.Size()
}
func (m *E2ApProtocolIE) XXX_DiscardUnknown() {
	xxx_messageInfo_E2ApProtocolIE.DiscardUnknown(m)
}

var xxx_messageInfo_E2ApProtocolIE proto.InternalMessageInfo

func (m *E2ApProtocolIE) GetId() ProtocolIEId {
	if m != nil {
		return m.Id
	}
	return ProtocolIEId_UNDEFINED
}

func (m *E2ApProtocolIE) GetCriticality() Criticality {
	if m != nil {
		return m.Criticality
	}
	return Criticality_REJECT
}

func (m *E2ApProtocolIE) GetPresence() Presence {
	if m != nil {
		return m.Presence
	}
	return Presence_OPTIONAL
}

type E2ApProtocolIEsPair struct {
	Id               ProtocolIEId `protobuf:"varint,1,opt,name=id,proto3,enum=interface.e2ap.ProtocolIEId" json:"id,omitempty"`
	FirstCriticality Criticality  `protobuf:"varint,2,opt,name=firstCriticality,proto3,enum=interface.e2ap.Criticality" json:"firstCriticality,omitempty"`
	// firstValue = 3;
	SecondCriticality Criticality `protobuf:"varint,4,opt,name=secondCriticality,proto3,enum=interface.e2ap.Criticality" json:"secondCriticality,omitempty"`
	// secondValue = 5;
	Presence Presence `protobuf:"varint,6,opt,name=presence,proto3,enum=interface.e2ap.Presence" json:"presence,omitempty"`
}

func (m *E2ApProtocolIEsPair) Reset()         { *m = E2ApProtocolIEsPair{} }
func (m *E2ApProtocolIEsPair) String() string { return proto.CompactTextString(m) }
func (*E2ApProtocolIEsPair) ProtoMessage()    {}
func (*E2ApProtocolIEsPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_e888e9637db470e9, []int{1}
}
func (m *E2ApProtocolIEsPair) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *E2ApProtocolIEsPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_E2ApProtocolIEsPair.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *E2ApProtocolIEsPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_E2ApProtocolIEsPair.Merge(m, src)
}
func (m *E2ApProtocolIEsPair) XXX_Size() int {
	return m.Size()
}
func (m *E2ApProtocolIEsPair) XXX_DiscardUnknown() {
	xxx_messageInfo_E2ApProtocolIEsPair.DiscardUnknown(m)
}

var xxx_messageInfo_E2ApProtocolIEsPair proto.InternalMessageInfo

func (m *E2ApProtocolIEsPair) GetId() ProtocolIEId {
	if m != nil {
		return m.Id
	}
	return ProtocolIEId_UNDEFINED
}

func (m *E2ApProtocolIEsPair) GetFirstCriticality() Criticality {
	if m != nil {
		return m.FirstCriticality
	}
	return Criticality_REJECT
}

func (m *E2ApProtocolIEsPair) GetSecondCriticality() Criticality {
	if m != nil {
		return m.SecondCriticality
	}
	return Criticality_REJECT
}

func (m *E2ApProtocolIEsPair) GetPresence() Presence {
	if m != nil {
		return m.Presence
	}
	return Presence_OPTIONAL
}

type ProtocolIEContainer struct {
}

func (m *ProtocolIEContainer) Reset()         { *m = ProtocolIEContainer{} }
func (m *ProtocolIEContainer) String() string { return proto.CompactTextString(m) }
func (*ProtocolIEContainer) ProtoMessage()    {}
func (*ProtocolIEContainer) Descriptor() ([]byte, []int) {
	return fileDescriptor_e888e9637db470e9, []int{2}
}
func (m *ProtocolIEContainer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProtocolIEContainer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProtocolIEContainer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProtocolIEContainer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtocolIEContainer.Merge(m, src)
}
func (m *ProtocolIEContainer) XXX_Size() int {
	return m.Size()
}
func (m *ProtocolIEContainer) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtocolIEContainer.DiscardUnknown(m)
}

var xxx_messageInfo_ProtocolIEContainer proto.InternalMessageInfo

type ProtocolIESingleContainer struct {
}

func (m *ProtocolIESingleContainer) Reset()         { *m = ProtocolIESingleContainer{} }
func (m *ProtocolIESingleContainer) String() string { return proto.CompactTextString(m) }
func (*ProtocolIESingleContainer) ProtoMessage()    {}
func (*ProtocolIESingleContainer) Descriptor() ([]byte, []int) {
	return fileDescriptor_e888e9637db470e9, []int{3}
}
func (m *ProtocolIESingleContainer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProtocolIESingleContainer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProtocolIESingleContainer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProtocolIESingleContainer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtocolIESingleContainer.Merge(m, src)
}
func (m *ProtocolIESingleContainer) XXX_Size() int {
	return m.Size()
}
func (m *ProtocolIESingleContainer) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtocolIESingleContainer.DiscardUnknown(m)
}

var xxx_messageInfo_ProtocolIESingleContainer proto.InternalMessageInfo

type ProtocolIEField struct {
}

func (m *ProtocolIEField) Reset()         { *m = ProtocolIEField{} }
func (m *ProtocolIEField) String() string { return proto.CompactTextString(m) }
func (*ProtocolIEField) ProtoMessage()    {}
func (*ProtocolIEField) Descriptor() ([]byte, []int) {
	return fileDescriptor_e888e9637db470e9, []int{4}
}
func (m *ProtocolIEField) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProtocolIEField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProtocolIEField.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProtocolIEField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtocolIEField.Merge(m, src)
}
func (m *ProtocolIEField) XXX_Size() int {
	return m.Size()
}
func (m *ProtocolIEField) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtocolIEField.DiscardUnknown(m)
}

var xxx_messageInfo_ProtocolIEField proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("interface.e2ap.Criticality", Criticality_name, Criticality_value)
	proto.RegisterEnum("interface.e2ap.Presence", Presence_name, Presence_value)
	proto.RegisterEnum("interface.e2ap.TriggeringMessage", TriggeringMessage_name, TriggeringMessage_value)
	proto.RegisterEnum("interface.e2ap.ProcedureCode", ProcedureCode_name, ProcedureCode_value)
	proto.RegisterEnum("interface.e2ap.ProtocolIEId", ProtocolIEId_name, ProtocolIEId_value)
	proto.RegisterType((*E2ApProtocolIE)(nil), "interface.e2ap.E2apProtocolIE")
	proto.RegisterType((*E2ApProtocolIEsPair)(nil), "interface.e2ap.E2apProtocolIEsPair")
	proto.RegisterType((*ProtocolIEContainer)(nil), "interface.e2ap.ProtocolIEContainer")
	proto.RegisterType((*ProtocolIESingleContainer)(nil), "interface.e2ap.ProtocolIESingleContainer")
	proto.RegisterType((*ProtocolIEField)(nil), "interface.e2ap.ProtocolIEField")
}

func init() { proto.RegisterFile("api/sb/e2ap.proto", fileDescriptor_e888e9637db470e9) }

var fileDescriptor_e888e9637db470e9 = []byte{
	// 892 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x4d, 0x6f, 0x22, 0x47,
	0x10, 0x65, 0xb0, 0xcd, 0x9a, 0xb2, 0x8d, 0x9b, 0xc6, 0x1f, 0xd8, 0x38, 0xac, 0xe5, 0x53, 0x64,
	0x45, 0x5e, 0x85, 0xe4, 0x90, 0x4b, 0x0e, 0xed, 0xee, 0x86, 0x74, 0x32, 0xcc, 0xcc, 0x76, 0xf7,
	0x6c, 0xe4, 0x4b, 0x5a, 0x2c, 0xcc, 0xa2, 0x91, 0x2c, 0x40, 0x03, 0x39, 0xe4, 0x47, 0x44, 0xca,
	0x7f, 0xc9, 0x9f, 0xc8, 0x21, 0x87, 0xbd, 0x44, 0xca, 0x31, 0xb2, 0xff, 0x48, 0xd4, 0x03, 0x63,
	0x66, 0xe0, 0x90, 0xec, 0x6d, 0xf4, 0x5e, 0x55, 0xf5, 0x7b, 0xcf, 0x55, 0x06, 0xea, 0x83, 0x59,
	0xfc, 0x66, 0xfe, 0xfe, 0x4d, 0xd4, 0x19, 0xcc, 0xee, 0x66, 0xc9, 0x74, 0x31, 0xc5, 0xb5, 0x78,
	0xb2, 0x88, 0x92, 0x0f, 0x83, 0x61, 0x74, 0x67, 0xd1, 0x9b, 0xdf, 0x1d, 0xa8, 0xf1, 0xce, 0x60,
	0x16, 0x58, 0x76, 0x38, 0x7d, 0x14, 0x1c, 0x7f, 0x01, 0xe5, 0x78, 0xd4, 0x74, 0xae, 0x9d, 0xcf,
	0x6b, 0x9d, 0xab, 0xbb, 0x62, 0xfd, 0xdd, 0xba, 0x4e, 0x8c, 0x64, 0x39, 0x1e, 0xe1, 0x6f, 0xe1,
	0x60, 0x98, 0xc4, 0x8b, 0x78, 0x38, 0x78, 0x8c, 0x17, 0xbf, 0x34, 0xcb, 0x69, 0x5b, 0x6b, 0xb3,
	0x8d, 0xae, 0x4b, 0x64, 0xbe, 0x1e, 0x7f, 0x0d, 0xfb, 0xb3, 0x24, 0x9a, 0x47, 0x93, 0x61, 0xd4,
	0xdc, 0x49, 0x7b, 0x9b, 0xdb, 0x4f, 0x2e, 0x79, 0xf9, 0x52, 0x79, 0xf3, 0x6b, 0x19, 0x1a, 0x45,
	0xd5, 0xf3, 0x60, 0x10, 0x27, 0x9f, 0x28, 0xbd, 0x07, 0xe8, 0x43, 0x9c, 0xcc, 0x17, 0xf4, 0xd3,
	0xf4, 0x6f, 0x35, 0x61, 0x01, 0xf5, 0x79, 0x34, 0x9c, 0x4e, 0x46, 0xf9, 0x49, 0xbb, 0xff, 0x3d,
	0x69, 0xbb, 0xab, 0x90, 0x47, 0xe5, 0x7f, 0xe7, 0x71, 0x0a, 0x8d, 0xb5, 0x3b, 0x3a, 0x9d, 0x2c,
	0x06, 0xf1, 0x24, 0x4a, 0x6e, 0x5a, 0x70, 0xb1, 0x86, 0x55, 0x3c, 0x19, 0x3f, 0x46, 0x6b, 0xb2,
	0x0e, 0xc7, 0x6b, 0xb2, 0x1b, 0x47, 0x8f, 0xa3, 0xdb, 0x2f, 0xe1, 0x20, 0xaf, 0x05, 0xa0, 0x22,
	0xf9, 0xf7, 0x9c, 0x6a, 0x54, 0xb2, 0xdf, 0xa2, 0xe7, 0xf9, 0x92, 0x23, 0xc7, 0x7e, 0x7b, 0xbe,
	0x16, 0xdd, 0x07, 0x54, 0xbe, 0xfd, 0x06, 0xf6, 0x33, 0x3d, 0xf8, 0x10, 0xf6, 0xfd, 0x40, 0x0b,
	0xdf, 0x23, 0x2e, 0x2a, 0xe1, 0x63, 0x38, 0xa0, 0xbe, 0xc7, 0xc4, 0x0a, 0x70, 0xf0, 0x11, 0x54,
	0xfb, 0xc4, 0x63, 0x44, 0xfb, 0xd2, 0x76, 0xfe, 0x04, 0x75, 0x9d, 0xc4, 0xe3, 0x71, 0x94, 0xc4,
	0x93, 0x71, 0x3f, 0x9a, 0xcf, 0x07, 0xe3, 0x08, 0x9f, 0x01, 0x16, 0x9e, 0xd0, 0x82, 0x68, 0xe1,
	0xf5, 0x4c, 0x9f, 0x2b, 0x45, 0x7a, 0x1c, 0x95, 0x2c, 0xae, 0x42, 0x4a, 0xb9, 0x52, 0xdd, 0xd0,
	0x35, 0x7e, 0xa8, 0xa9, 0xdf, 0xb7, 0x52, 0x2e, 0xe0, 0x34, 0xf4, 0xd6, 0xcc, 0x9a, 0x2a, 0xdf,
	0xfe, 0xe5, 0xc0, 0x51, 0x90, 0x4c, 0x87, 0xd1, 0xe8, 0xe7, 0x24, 0xa2, 0xd3, 0x51, 0x84, 0x6b,
	0x00, 0x01, 0x35, 0xc2, 0x7b, 0x47, 0x5c, 0xc1, 0x50, 0xc9, 0xea, 0xe5, 0x1d, 0xa3, 0xb8, 0x0e,
	0x03, 0xe4, 0xe0, 0x13, 0x40, 0x5c, 0x4a, 0x5f, 0x1a, 0xe1, 0x31, 0x41, 0x89, 0x95, 0x8d, 0xca,
	0xb8, 0x0a, 0x7b, 0x92, 0x2b, 0xae, 0xd1, 0x8e, 0x35, 0x24, 0x05, 0x35, 0xd4, 0xf7, 0xb4, 0xf4,
	0x5d, 0xb4, 0x8b, 0x31, 0xd4, 0x2c, 0x90, 0xab, 0xdf, 0xc3, 0xa7, 0x50, 0xb7, 0x98, 0xe2, 0xf2,
	0x9d, 0xa0, 0xdc, 0xbc, 0x0d, 0xb9, 0x7c, 0x40, 0x15, 0xab, 0x3f, 0x0f, 0x87, 0x01, 0x23, 0x9a,
	0xa3, 0x57, 0xf6, 0xd1, 0x14, 0x0f, 0xef, 0x15, 0x95, 0x22, 0x0d, 0x0f, 0xed, 0xe3, 0x16, 0x9c,
	0x6f, 0xa2, 0x86, 0x71, 0x97, 0x6b, 0x8e, 0xaa, 0xb7, 0x7f, 0x56, 0xe0, 0x30, 0xbf, 0xca, 0x36,
	0xd7, 0xd0, 0x63, 0xbc, 0x2b, 0x3c, 0x6e, 0x5d, 0x55, 0x61, 0x8f, 0x92, 0x50, 0xd9, 0x74, 0x5a,
	0x70, 0x4e, 0xa5, 0xd0, 0x82, 0x12, 0x57, 0xe8, 0x07, 0xc3, 0x04, 0xe9, 0x79, 0xbe, 0xd2, 0x82,
	0x2a, 0x54, 0xb6, 0x4a, 0x7b, 0xae, 0x7f, 0x4f, 0x5c, 0xc3, 0x3b, 0xc6, 0xf3, 0x19, 0x37, 0x82,
	0xa1, 0x1d, 0x5c, 0x87, 0xa3, 0x15, 0x9c, 0x7a, 0x63, 0x68, 0x17, 0x37, 0xe0, 0x58, 0x12, 0xcf,
	0x74, 0x43, 0x8f, 0xa6, 0x52, 0x04, 0x43, 0x7b, 0xb8, 0x09, 0x27, 0x1b, 0xa0, 0x11, 0x9a, 0xf7,
	0x51, 0x05, 0xb7, 0xe1, 0xb2, 0xc8, 0x70, 0x93, 0x0a, 0x5a, 0xf2, 0xaf, 0xd2, 0x88, 0x0a, 0xbc,
	0x85, 0xf7, 0xf1, 0x25, 0x9c, 0xe5, 0x61, 0x65, 0x08, 0xa5, 0x3c, 0xd0, 0x9c, 0xa1, 0x2a, 0x3e,
	0x87, 0xc6, 0x06, 0xc7, 0x18, 0x67, 0x08, 0xec, 0xdf, 0xbf, 0x48, 0x2c, 0x63, 0x62, 0xe8, 0x60,
	0x7b, 0x5e, 0xdf, 0x67, 0xa2, 0x2b, 0x38, 0x43, 0x87, 0xdb, 0xdc, 0x72, 0xcf, 0x39, 0x43, 0x47,
	0xf8, 0x0a, 0x9a, 0xd6, 0x39, 0x59, 0x8a, 0x23, 0xac, 0x2f, 0xb4, 0xe6, 0x2b, 0x73, 0x35, 0x1b,
	0x4f, 0x8e, 0x15, 0x0c, 0x1d, 0xe3, 0x6b, 0xb8, 0xca, 0x41, 0x9e, 0xaf, 0x37, 0x9a, 0x50, 0x9a,
	0xd5, 0x4b, 0x85, 0x7a, 0xa1, 0x51, 0xbd, 0xf8, 0x98, 0x2a, 0x34, 0x23, 0x8c, 0x5f, 0x43, 0x2b,
	0x37, 0x59, 0xfb, 0xe6, 0x9e, 0x2f, 0xd7, 0x75, 0x39, 0xb8, 0x91, 0xe6, 0x62, 0x57, 0x92, 0xb8,
	0xae, 0x09, 0xa4, 0x6f, 0xaf, 0xc0, 0x6a, 0x3a, 0xc9, 0x36, 0x68, 0xb5, 0xab, 0x86, 0xd0, 0x1f,
	0x8c, 0xe4, 0x6f, 0x43, 0xae, 0x34, 0x3a, 0xcd, 0x96, 0x31, 0x23, 0xbf, 0xe3, 0x84, 0x71, 0x89,
	0xce, 0x5e, 0xa6, 0xad, 0xf0, 0xec, 0xfa, 0xce, 0x37, 0x1b, 0x94, 0x26, 0x3a, 0x54, 0xa8, 0x99,
	0xa6, 0x5f, 0x38, 0x80, 0x6c, 0xd6, 0x45, 0x9a, 0x70, 0x91, 0xca, 0xc6, 0x5d, 0x66, 0x37, 0x92,
	0xe3, 0x94, 0x87, 0x5a, 0xd9, 0xf3, 0x39, 0x58, 0x3f, 0x04, 0x1c, 0x5d, 0x65, 0x77, 0xb6, 0x32,
	0x60, 0x0d, 0x7e, 0x96, 0x05, 0xb7, 0x71, 0x22, 0x9a, 0x08, 0x57, 0xa1, 0x36, 0x46, 0x70, 0xa8,
	0x45, 0x9f, 0xdb, 0xc8, 0x7e, 0x24, 0x42, 0xa3, 0xd7, 0x9b, 0xde, 0xb2, 0x7f, 0x13, 0xd7, 0xf7,
	0xcd, 0x3f, 0x9e, 0xda, 0xce, 0xc7, 0xa7, 0xb6, 0xf3, 0xcf, 0x53, 0xdb, 0xf9, 0xed, 0xb9, 0x5d,
	0xfa, 0xf8, 0xdc, 0x2e, 0xfd, 0xfd, 0xdc, 0x2e, 0xbd, 0xaf, 0xa4, 0xbf, 0x98, 0x5f, 0xfd, 0x1b,
	0x00, 0x00, 0xff, 0xff, 0x8b, 0x69, 0x5f, 0x3d, 0x46, 0x07, 0x00, 0x00,
}

func (m *E2ApProtocolIE) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *E2ApProtocolIE) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *E2ApProtocolIE) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Presence != 0 {
		i = encodeVarintE2Ap(dAtA, i, uint64(m.Presence))
		i--
		dAtA[i] = 0x18
	}
	if m.Criticality != 0 {
		i = encodeVarintE2Ap(dAtA, i, uint64(m.Criticality))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintE2Ap(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *E2ApProtocolIEsPair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *E2ApProtocolIEsPair) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *E2ApProtocolIEsPair) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Presence != 0 {
		i = encodeVarintE2Ap(dAtA, i, uint64(m.Presence))
		i--
		dAtA[i] = 0x30
	}
	if m.SecondCriticality != 0 {
		i = encodeVarintE2Ap(dAtA, i, uint64(m.SecondCriticality))
		i--
		dAtA[i] = 0x20
	}
	if m.FirstCriticality != 0 {
		i = encodeVarintE2Ap(dAtA, i, uint64(m.FirstCriticality))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintE2Ap(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ProtocolIEContainer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProtocolIEContainer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProtocolIEContainer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ProtocolIESingleContainer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProtocolIESingleContainer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProtocolIESingleContainer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ProtocolIEField) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProtocolIEField) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProtocolIEField) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintE2Ap(dAtA []byte, offset int, v uint64) int {
	offset -= sovE2Ap(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *E2ApProtocolIE) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovE2Ap(uint64(m.Id))
	}
	if m.Criticality != 0 {
		n += 1 + sovE2Ap(uint64(m.Criticality))
	}
	if m.Presence != 0 {
		n += 1 + sovE2Ap(uint64(m.Presence))
	}
	return n
}

func (m *E2ApProtocolIEsPair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovE2Ap(uint64(m.Id))
	}
	if m.FirstCriticality != 0 {
		n += 1 + sovE2Ap(uint64(m.FirstCriticality))
	}
	if m.SecondCriticality != 0 {
		n += 1 + sovE2Ap(uint64(m.SecondCriticality))
	}
	if m.Presence != 0 {
		n += 1 + sovE2Ap(uint64(m.Presence))
	}
	return n
}

func (m *ProtocolIEContainer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ProtocolIESingleContainer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ProtocolIEField) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovE2Ap(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozE2Ap(x uint64) (n int) {
	return sovE2Ap(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *E2ApProtocolIE) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowE2Ap
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: E2apProtocolIE: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: E2apProtocolIE: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowE2Ap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= ProtocolIEId(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Criticality", wireType)
			}
			m.Criticality = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowE2Ap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Criticality |= Criticality(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Presence", wireType)
			}
			m.Presence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowE2Ap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Presence |= Presence(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipE2Ap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthE2Ap
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthE2Ap
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *E2ApProtocolIEsPair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowE2Ap
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: E2apProtocolIEsPair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: E2apProtocolIEsPair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowE2Ap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= ProtocolIEId(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FirstCriticality", wireType)
			}
			m.FirstCriticality = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowE2Ap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FirstCriticality |= Criticality(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SecondCriticality", wireType)
			}
			m.SecondCriticality = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowE2Ap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SecondCriticality |= Criticality(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Presence", wireType)
			}
			m.Presence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowE2Ap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Presence |= Presence(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipE2Ap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthE2Ap
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthE2Ap
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ProtocolIEContainer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowE2Ap
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProtocolIEContainer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProtocolIEContainer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipE2Ap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthE2Ap
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthE2Ap
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ProtocolIESingleContainer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowE2Ap
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProtocolIESingleContainer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProtocolIESingleContainer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipE2Ap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthE2Ap
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthE2Ap
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ProtocolIEField) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowE2Ap
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProtocolIEField: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProtocolIEField: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipE2Ap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthE2Ap
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthE2Ap
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipE2Ap(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowE2Ap
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowE2Ap
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowE2Ap
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthE2Ap
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupE2Ap
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthE2Ap
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthE2Ap        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowE2Ap          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupE2Ap = fmt.Errorf("proto: unexpected end of group")
)