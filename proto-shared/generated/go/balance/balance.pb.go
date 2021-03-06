// Code generated by protoc-gen-go. DO NOT EDIT.
// source: balance/balance.proto

package balance

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Balance struct {
	AcceptLanguage       string               `protobuf:"bytes,1,opt,name=accept_language,json=acceptLanguage,proto3" json:"accept_language,omitempty"`
	Locale               string               `protobuf:"bytes,2,opt,name=locale,proto3" json:"locale,omitempty"`
	Id                   string               `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Account              string               `protobuf:"bytes,4,opt,name=account,proto3" json:"account,omitempty"`
	Amount               int32                `protobuf:"varint,5,opt,name=amount,proto3" json:"amount,omitempty"`
	DateBalance          *timestamp.Timestamp `protobuf:"bytes,6,opt,name=date_balance,json=dateBalance,proto3" json:"date_balance,omitempty"`
	Description          string               `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Balance) Reset()         { *m = Balance{} }
func (m *Balance) String() string { return proto.CompactTextString(m) }
func (*Balance) ProtoMessage()    {}
func (*Balance) Descriptor() ([]byte, []int) {
	return fileDescriptor_15e3a49c407c385e, []int{0}
}

func (m *Balance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Balance.Unmarshal(m, b)
}
func (m *Balance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Balance.Marshal(b, m, deterministic)
}
func (m *Balance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Balance.Merge(m, src)
}
func (m *Balance) XXX_Size() int {
	return xxx_messageInfo_Balance.Size(m)
}
func (m *Balance) XXX_DiscardUnknown() {
	xxx_messageInfo_Balance.DiscardUnknown(m)
}

var xxx_messageInfo_Balance proto.InternalMessageInfo

func (m *Balance) GetAcceptLanguage() string {
	if m != nil {
		return m.AcceptLanguage
	}
	return ""
}

func (m *Balance) GetLocale() string {
	if m != nil {
		return m.Locale
	}
	return ""
}

func (m *Balance) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Balance) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *Balance) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Balance) GetDateBalance() *timestamp.Timestamp {
	if m != nil {
		return m.DateBalance
	}
	return nil
}

func (m *Balance) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type Setup struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Pid                  string   `protobuf:"bytes,2,opt,name=pid,proto3" json:"pid,omitempty"`
	IpAdress             string   `protobuf:"bytes,3,opt,name=ip_adress,json=ipAdress,proto3" json:"ip_adress,omitempty"`
	Host                 string   `protobuf:"bytes,4,opt,name=host,proto3" json:"host,omitempty"`
	Port                 string   `protobuf:"bytes,5,opt,name=port,proto3" json:"port,omitempty"`
	ErrorCode            string   `protobuf:"bytes,6,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	ResponseTime         int32    `protobuf:"varint,7,opt,name=response_time,json=responseTime,proto3" json:"response_time,omitempty"`
	IsRandonTime         bool     `protobuf:"varint,8,opt,name=is_randon_time,json=isRandonTime,proto3" json:"is_randon_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Setup) Reset()         { *m = Setup{} }
func (m *Setup) String() string { return proto.CompactTextString(m) }
func (*Setup) ProtoMessage()    {}
func (*Setup) Descriptor() ([]byte, []int) {
	return fileDescriptor_15e3a49c407c385e, []int{1}
}

func (m *Setup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Setup.Unmarshal(m, b)
}
func (m *Setup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Setup.Marshal(b, m, deterministic)
}
func (m *Setup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Setup.Merge(m, src)
}
func (m *Setup) XXX_Size() int {
	return xxx_messageInfo_Setup.Size(m)
}
func (m *Setup) XXX_DiscardUnknown() {
	xxx_messageInfo_Setup.DiscardUnknown(m)
}

var xxx_messageInfo_Setup proto.InternalMessageInfo

func (m *Setup) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Setup) GetPid() string {
	if m != nil {
		return m.Pid
	}
	return ""
}

func (m *Setup) GetIpAdress() string {
	if m != nil {
		return m.IpAdress
	}
	return ""
}

func (m *Setup) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Setup) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *Setup) GetErrorCode() string {
	if m != nil {
		return m.ErrorCode
	}
	return ""
}

func (m *Setup) GetResponseTime() int32 {
	if m != nil {
		return m.ResponseTime
	}
	return 0
}

func (m *Setup) GetIsRandonTime() bool {
	if m != nil {
		return m.IsRandonTime
	}
	return false
}

type SetupRequest struct {
	Setup                *Setup   `protobuf:"bytes,1,opt,name=setup,proto3" json:"setup,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetupRequest) Reset()         { *m = SetupRequest{} }
func (m *SetupRequest) String() string { return proto.CompactTextString(m) }
func (*SetupRequest) ProtoMessage()    {}
func (*SetupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15e3a49c407c385e, []int{2}
}

func (m *SetupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetupRequest.Unmarshal(m, b)
}
func (m *SetupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetupRequest.Marshal(b, m, deterministic)
}
func (m *SetupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetupRequest.Merge(m, src)
}
func (m *SetupRequest) XXX_Size() int {
	return xxx_messageInfo_SetupRequest.Size(m)
}
func (m *SetupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetupRequest proto.InternalMessageInfo

func (m *SetupRequest) GetSetup() *Setup {
	if m != nil {
		return m.Setup
	}
	return nil
}

type SetupResponse struct {
	Setup                *Setup   `protobuf:"bytes,1,opt,name=setup,proto3" json:"setup,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetupResponse) Reset()         { *m = SetupResponse{} }
func (m *SetupResponse) String() string { return proto.CompactTextString(m) }
func (*SetupResponse) ProtoMessage()    {}
func (*SetupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_15e3a49c407c385e, []int{3}
}

func (m *SetupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetupResponse.Unmarshal(m, b)
}
func (m *SetupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetupResponse.Marshal(b, m, deterministic)
}
func (m *SetupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetupResponse.Merge(m, src)
}
func (m *SetupResponse) XXX_Size() int {
	return xxx_messageInfo_SetupResponse.Size(m)
}
func (m *SetupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetupResponse proto.InternalMessageInfo

func (m *SetupResponse) GetSetup() *Setup {
	if m != nil {
		return m.Setup
	}
	return nil
}

type HealthRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthRequest) Reset()         { *m = HealthRequest{} }
func (m *HealthRequest) String() string { return proto.CompactTextString(m) }
func (*HealthRequest) ProtoMessage()    {}
func (*HealthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15e3a49c407c385e, []int{4}
}

func (m *HealthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthRequest.Unmarshal(m, b)
}
func (m *HealthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthRequest.Marshal(b, m, deterministic)
}
func (m *HealthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthRequest.Merge(m, src)
}
func (m *HealthRequest) XXX_Size() int {
	return xxx_messageInfo_HealthRequest.Size(m)
}
func (m *HealthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HealthRequest proto.InternalMessageInfo

type HealthResponse struct {
	Health               *Setup   `protobuf:"bytes,1,opt,name=health,proto3" json:"health,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthResponse) Reset()         { *m = HealthResponse{} }
func (m *HealthResponse) String() string { return proto.CompactTextString(m) }
func (*HealthResponse) ProtoMessage()    {}
func (*HealthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_15e3a49c407c385e, []int{5}
}

func (m *HealthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthResponse.Unmarshal(m, b)
}
func (m *HealthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthResponse.Marshal(b, m, deterministic)
}
func (m *HealthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthResponse.Merge(m, src)
}
func (m *HealthResponse) XXX_Size() int {
	return xxx_messageInfo_HealthResponse.Size(m)
}
func (m *HealthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthResponse proto.InternalMessageInfo

func (m *HealthResponse) GetHealth() *Setup {
	if m != nil {
		return m.Health
	}
	return nil
}

type StressRequest struct {
	Count                int32    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StressRequest) Reset()         { *m = StressRequest{} }
func (m *StressRequest) String() string { return proto.CompactTextString(m) }
func (*StressRequest) ProtoMessage()    {}
func (*StressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15e3a49c407c385e, []int{6}
}

func (m *StressRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StressRequest.Unmarshal(m, b)
}
func (m *StressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StressRequest.Marshal(b, m, deterministic)
}
func (m *StressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StressRequest.Merge(m, src)
}
func (m *StressRequest) XXX_Size() int {
	return xxx_messageInfo_StressRequest.Size(m)
}
func (m *StressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StressRequest proto.InternalMessageInfo

func (m *StressRequest) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type StressResponse struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StressResponse) Reset()         { *m = StressResponse{} }
func (m *StressResponse) String() string { return proto.CompactTextString(m) }
func (*StressResponse) ProtoMessage()    {}
func (*StressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_15e3a49c407c385e, []int{7}
}

func (m *StressResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StressResponse.Unmarshal(m, b)
}
func (m *StressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StressResponse.Marshal(b, m, deterministic)
}
func (m *StressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StressResponse.Merge(m, src)
}
func (m *StressResponse) XXX_Size() int {
	return xxx_messageInfo_StressResponse.Size(m)
}
func (m *StressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StressResponse proto.InternalMessageInfo

func (m *StressResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type ListBalanceRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBalanceRequest) Reset()         { *m = ListBalanceRequest{} }
func (m *ListBalanceRequest) String() string { return proto.CompactTextString(m) }
func (*ListBalanceRequest) ProtoMessage()    {}
func (*ListBalanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15e3a49c407c385e, []int{8}
}

func (m *ListBalanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBalanceRequest.Unmarshal(m, b)
}
func (m *ListBalanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBalanceRequest.Marshal(b, m, deterministic)
}
func (m *ListBalanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBalanceRequest.Merge(m, src)
}
func (m *ListBalanceRequest) XXX_Size() int {
	return xxx_messageInfo_ListBalanceRequest.Size(m)
}
func (m *ListBalanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBalanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListBalanceRequest proto.InternalMessageInfo

type ListBalanceResponse struct {
	Balance              []*Balance `protobuf:"bytes,1,rep,name=balance,proto3" json:"balance,omitempty"`
	Total                int32      `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListBalanceResponse) Reset()         { *m = ListBalanceResponse{} }
func (m *ListBalanceResponse) String() string { return proto.CompactTextString(m) }
func (*ListBalanceResponse) ProtoMessage()    {}
func (*ListBalanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_15e3a49c407c385e, []int{9}
}

func (m *ListBalanceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBalanceResponse.Unmarshal(m, b)
}
func (m *ListBalanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBalanceResponse.Marshal(b, m, deterministic)
}
func (m *ListBalanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBalanceResponse.Merge(m, src)
}
func (m *ListBalanceResponse) XXX_Size() int {
	return xxx_messageInfo_ListBalanceResponse.Size(m)
}
func (m *ListBalanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBalanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListBalanceResponse proto.InternalMessageInfo

func (m *ListBalanceResponse) GetBalance() []*Balance {
	if m != nil {
		return m.Balance
	}
	return nil
}

func (m *ListBalanceResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

type GetBalanceRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBalanceRequest) Reset()         { *m = GetBalanceRequest{} }
func (m *GetBalanceRequest) String() string { return proto.CompactTextString(m) }
func (*GetBalanceRequest) ProtoMessage()    {}
func (*GetBalanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15e3a49c407c385e, []int{10}
}

func (m *GetBalanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBalanceRequest.Unmarshal(m, b)
}
func (m *GetBalanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBalanceRequest.Marshal(b, m, deterministic)
}
func (m *GetBalanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBalanceRequest.Merge(m, src)
}
func (m *GetBalanceRequest) XXX_Size() int {
	return xxx_messageInfo_GetBalanceRequest.Size(m)
}
func (m *GetBalanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBalanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBalanceRequest proto.InternalMessageInfo

func (m *GetBalanceRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetBalanceResponse struct {
	Balance              *Balance `protobuf:"bytes,1,opt,name=balance,proto3" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBalanceResponse) Reset()         { *m = GetBalanceResponse{} }
func (m *GetBalanceResponse) String() string { return proto.CompactTextString(m) }
func (*GetBalanceResponse) ProtoMessage()    {}
func (*GetBalanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_15e3a49c407c385e, []int{11}
}

func (m *GetBalanceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBalanceResponse.Unmarshal(m, b)
}
func (m *GetBalanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBalanceResponse.Marshal(b, m, deterministic)
}
func (m *GetBalanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBalanceResponse.Merge(m, src)
}
func (m *GetBalanceResponse) XXX_Size() int {
	return xxx_messageInfo_GetBalanceResponse.Size(m)
}
func (m *GetBalanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBalanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBalanceResponse proto.InternalMessageInfo

func (m *GetBalanceResponse) GetBalance() *Balance {
	if m != nil {
		return m.Balance
	}
	return nil
}

type AddBalanceRequest struct {
	Balance              *Balance `protobuf:"bytes,1,opt,name=balance,proto3" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddBalanceRequest) Reset()         { *m = AddBalanceRequest{} }
func (m *AddBalanceRequest) String() string { return proto.CompactTextString(m) }
func (*AddBalanceRequest) ProtoMessage()    {}
func (*AddBalanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15e3a49c407c385e, []int{12}
}

func (m *AddBalanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddBalanceRequest.Unmarshal(m, b)
}
func (m *AddBalanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddBalanceRequest.Marshal(b, m, deterministic)
}
func (m *AddBalanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddBalanceRequest.Merge(m, src)
}
func (m *AddBalanceRequest) XXX_Size() int {
	return xxx_messageInfo_AddBalanceRequest.Size(m)
}
func (m *AddBalanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddBalanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddBalanceRequest proto.InternalMessageInfo

func (m *AddBalanceRequest) GetBalance() *Balance {
	if m != nil {
		return m.Balance
	}
	return nil
}

type AddBalanceResponse struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddBalanceResponse) Reset()         { *m = AddBalanceResponse{} }
func (m *AddBalanceResponse) String() string { return proto.CompactTextString(m) }
func (*AddBalanceResponse) ProtoMessage()    {}
func (*AddBalanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_15e3a49c407c385e, []int{13}
}

func (m *AddBalanceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddBalanceResponse.Unmarshal(m, b)
}
func (m *AddBalanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddBalanceResponse.Marshal(b, m, deterministic)
}
func (m *AddBalanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddBalanceResponse.Merge(m, src)
}
func (m *AddBalanceResponse) XXX_Size() int {
	return xxx_messageInfo_AddBalanceResponse.Size(m)
}
func (m *AddBalanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddBalanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddBalanceResponse proto.InternalMessageInfo

func (m *AddBalanceResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func init() {
	proto.RegisterType((*Balance)(nil), "balance.Balance")
	proto.RegisterType((*Setup)(nil), "balance.Setup")
	proto.RegisterType((*SetupRequest)(nil), "balance.SetupRequest")
	proto.RegisterType((*SetupResponse)(nil), "balance.SetupResponse")
	proto.RegisterType((*HealthRequest)(nil), "balance.HealthRequest")
	proto.RegisterType((*HealthResponse)(nil), "balance.HealthResponse")
	proto.RegisterType((*StressRequest)(nil), "balance.StressRequest")
	proto.RegisterType((*StressResponse)(nil), "balance.StressResponse")
	proto.RegisterType((*ListBalanceRequest)(nil), "balance.ListBalanceRequest")
	proto.RegisterType((*ListBalanceResponse)(nil), "balance.ListBalanceResponse")
	proto.RegisterType((*GetBalanceRequest)(nil), "balance.GetBalanceRequest")
	proto.RegisterType((*GetBalanceResponse)(nil), "balance.GetBalanceResponse")
	proto.RegisterType((*AddBalanceRequest)(nil), "balance.AddBalanceRequest")
	proto.RegisterType((*AddBalanceResponse)(nil), "balance.AddBalanceResponse")
}

func init() { proto.RegisterFile("balance/balance.proto", fileDescriptor_15e3a49c407c385e) }

var fileDescriptor_15e3a49c407c385e = []byte{
	// 634 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x5f, 0x6f, 0xd3, 0x3e,
	0x14, 0x5d, 0xb6, 0xa5, 0x7f, 0x6e, 0xdb, 0x6c, 0xf3, 0x6f, 0xbf, 0x29, 0xea, 0x40, 0x54, 0xde,
	0x80, 0x0a, 0xb1, 0x56, 0x2a, 0x03, 0xf1, 0x82, 0x60, 0xe3, 0x61, 0x08, 0xed, 0x29, 0x43, 0x42,
	0xe2, 0xa5, 0x72, 0x13, 0x93, 0x58, 0x4a, 0xe3, 0x60, 0x3b, 0x88, 0x0f, 0xcb, 0x57, 0x80, 0xcf,
	0x80, 0xe2, 0x3f, 0x5b, 0x58, 0xcb, 0x04, 0x4f, 0xf5, 0x3d, 0xf7, 0xe4, 0x9c, 0x7b, 0x6e, 0x9d,
	0xc0, 0xff, 0x0b, 0x92, 0x93, 0x22, 0xa6, 0x53, 0xfb, 0x3b, 0x29, 0x05, 0x57, 0x1c, 0xb5, 0x6d,
	0x39, 0x7c, 0x90, 0x72, 0x9e, 0xe6, 0x74, 0xaa, 0xe1, 0x45, 0xf5, 0x79, 0xaa, 0xd8, 0x92, 0x4a,
	0x45, 0x96, 0xa5, 0x61, 0xe2, 0x1f, 0x1e, 0xb4, 0xcf, 0x0d, 0x19, 0x3d, 0x86, 0x1d, 0x12, 0xc7,
	0xb4, 0x54, 0xf3, 0x9c, 0x14, 0x69, 0x45, 0x52, 0x1a, 0x7a, 0x23, 0x6f, 0xdc, 0x8d, 0x02, 0x03,
	0x5f, 0x5a, 0x14, 0x1d, 0x40, 0x2b, 0xe7, 0x31, 0xc9, 0x69, 0xb8, 0xa9, 0xfb, 0xb6, 0x42, 0x01,
	0x6c, 0xb2, 0x24, 0xdc, 0xd2, 0xd8, 0x26, 0x4b, 0x50, 0x08, 0x6d, 0x12, 0xc7, 0xbc, 0x2a, 0x54,
	0xb8, 0xad, 0x41, 0x57, 0xd6, 0x0a, 0x64, 0xa9, 0x1b, 0xfe, 0xc8, 0x1b, 0xfb, 0x91, 0xad, 0xd0,
	0x2b, 0xe8, 0x27, 0x44, 0xd1, 0xb9, 0x9d, 0x3f, 0x6c, 0x8d, 0xbc, 0x71, 0x6f, 0x36, 0x9c, 0x98,
	0x18, 0x13, 0x17, 0x63, 0xf2, 0xc1, 0xc5, 0x88, 0x7a, 0x35, 0xdf, 0x25, 0x18, 0x41, 0x2f, 0xa1,
	0x32, 0x16, 0xac, 0x54, 0x8c, 0x17, 0x61, 0x5b, 0x9b, 0x36, 0x21, 0xfc, 0xdd, 0x03, 0xff, 0x8a,
	0xaa, 0xaa, 0x44, 0x08, 0xb6, 0x0b, 0xb2, 0x74, 0x11, 0xf5, 0x19, 0xed, 0xc2, 0x56, 0xc9, 0x12,
	0x9b, 0xaa, 0x3e, 0xa2, 0x43, 0xe8, 0xb2, 0x72, 0x4e, 0x12, 0x41, 0xa5, 0xb4, 0xc9, 0x3a, 0xac,
	0x3c, 0xd3, 0x75, 0x2d, 0x91, 0x71, 0xe9, 0xc2, 0xe9, 0x73, 0x8d, 0x95, 0x5c, 0x98, 0x5c, 0xdd,
	0x48, 0x9f, 0xd1, 0x7d, 0x00, 0x2a, 0x04, 0x17, 0xf3, 0x98, 0x27, 0x26, 0x53, 0x37, 0xea, 0x6a,
	0xe4, 0x2d, 0x4f, 0x28, 0x3a, 0x82, 0x81, 0xa0, 0xb2, 0xe4, 0x85, 0xa4, 0xf3, 0xfa, 0xff, 0xd1,
	0x73, 0xfb, 0x51, 0xdf, 0x81, 0x75, 0x58, 0x74, 0x0c, 0x01, 0x93, 0x73, 0x41, 0x8a, 0x84, 0x17,
	0x86, 0xd5, 0x19, 0x79, 0xe3, 0x4e, 0xd4, 0x67, 0x32, 0xd2, 0x60, 0xcd, 0xc2, 0xa7, 0xd0, 0xd7,
	0xe9, 0x22, 0xfa, 0xa5, 0xa2, 0x52, 0xa1, 0x63, 0xf0, 0x65, 0x5d, 0xeb, 0x94, 0xbd, 0x59, 0x30,
	0x71, 0xf7, 0xc4, 0xb0, 0x4c, 0x13, 0x3f, 0x87, 0x81, 0x7d, 0xca, 0x18, 0xfe, 0xe5, 0x63, 0x3b,
	0x30, 0x78, 0x47, 0x49, 0xae, 0x32, 0xeb, 0x86, 0x5f, 0x42, 0xe0, 0x00, 0x2b, 0xf4, 0x08, 0x5a,
	0x99, 0x46, 0xfe, 0xa0, 0x64, 0xbb, 0xf8, 0x21, 0x0c, 0xae, 0x54, 0xbd, 0x53, 0x37, 0xf8, 0x3e,
	0xf8, 0xe6, 0xe2, 0x78, 0x7a, 0x17, 0xa6, 0xc0, 0x63, 0x08, 0x1c, 0xcd, 0x1a, 0x1c, 0x40, 0x4b,
	0x50, 0x59, 0xe5, 0x86, 0xd8, 0x89, 0x6c, 0x85, 0xf7, 0x01, 0x5d, 0x32, 0xa9, 0xec, 0xc5, 0x70,
	0x03, 0x7e, 0x84, 0xff, 0x7e, 0x43, 0xad, 0xc8, 0x13, 0x70, 0x2f, 0x4c, 0xe8, 0x8d, 0xb6, 0xc6,
	0xbd, 0xd9, 0xee, 0xf5, 0x98, 0x8e, 0xea, 0x08, 0xf5, 0x60, 0x8a, 0x2b, 0x92, 0xeb, 0x4b, 0xe2,
	0x47, 0xa6, 0xc0, 0x47, 0xb0, 0x77, 0x41, 0x6f, 0xb9, 0xd9, 0xd7, 0xc1, 0x73, 0xaf, 0x03, 0x7e,
	0x03, 0xa8, 0x49, 0x5a, 0x67, 0xee, 0xdd, 0x69, 0x8e, 0x5f, 0xc3, 0xde, 0x59, 0x92, 0xdc, 0xb2,
	0xf9, 0x17, 0x81, 0xa7, 0x80, 0x9a, 0x02, 0x77, 0x2f, 0x71, 0xf6, 0xd3, 0x83, 0xc0, 0x72, 0xaf,
	0xa8, 0xf8, 0xca, 0x62, 0x8a, 0xde, 0x43, 0xaf, 0xb1, 0x41, 0x74, 0x78, 0x6d, 0xb5, 0xba, 0xed,
	0xe1, 0xbd, 0xf5, 0x4d, 0x63, 0x8a, 0x37, 0xd0, 0x05, 0xc0, 0xcd, 0x3e, 0xd0, 0xf0, 0x9a, 0xbd,
	0xb2, 0xc9, 0xe1, 0xe1, 0xda, 0x5e, 0x53, 0xe8, 0x26, 0x55, 0x43, 0x68, 0x65, 0x57, 0x0d, 0xa1,
	0xd5, 0x35, 0xe0, 0x8d, 0xf3, 0x17, 0x9f, 0x4e, 0x53, 0xa6, 0xb2, 0x6a, 0x31, 0x89, 0xf9, 0x72,
	0x9a, 0xf2, 0x93, 0x8c, 0x7e, 0x23, 0xe6, 0xe3, 0x79, 0x22, 0x33, 0x22, 0x68, 0x32, 0x4d, 0x69,
	0x41, 0x05, 0x51, 0xf5, 0x89, 0xbb, 0xaf, 0xee, 0xa2, 0xa5, 0x29, 0xcf, 0x7e, 0x05, 0x00, 0x00,
	0xff, 0xff, 0xf4, 0x31, 0x96, 0x16, 0x8f, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BalanceServiceClient is the client API for BalanceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BalanceServiceClient interface {
	//Api Unary
	ListBalance(ctx context.Context, in *ListBalanceRequest, opts ...grpc.CallOption) (*ListBalanceResponse, error)
	//Api Unary
	GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error)
	//Api Unary
	AddBalance(ctx context.Context, in *AddBalanceRequest, opts ...grpc.CallOption) (*AddBalanceResponse, error)
}

type balanceServiceClient struct {
	cc *grpc.ClientConn
}

func NewBalanceServiceClient(cc *grpc.ClientConn) BalanceServiceClient {
	return &balanceServiceClient{cc}
}

func (c *balanceServiceClient) ListBalance(ctx context.Context, in *ListBalanceRequest, opts ...grpc.CallOption) (*ListBalanceResponse, error) {
	out := new(ListBalanceResponse)
	err := c.cc.Invoke(ctx, "/balance.BalanceService/ListBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *balanceServiceClient) GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error) {
	out := new(GetBalanceResponse)
	err := c.cc.Invoke(ctx, "/balance.BalanceService/GetBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *balanceServiceClient) AddBalance(ctx context.Context, in *AddBalanceRequest, opts ...grpc.CallOption) (*AddBalanceResponse, error) {
	out := new(AddBalanceResponse)
	err := c.cc.Invoke(ctx, "/balance.BalanceService/AddBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BalanceServiceServer is the server API for BalanceService service.
type BalanceServiceServer interface {
	//Api Unary
	ListBalance(context.Context, *ListBalanceRequest) (*ListBalanceResponse, error)
	//Api Unary
	GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error)
	//Api Unary
	AddBalance(context.Context, *AddBalanceRequest) (*AddBalanceResponse, error)
}

// UnimplementedBalanceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBalanceServiceServer struct {
}

func (*UnimplementedBalanceServiceServer) ListBalance(ctx context.Context, req *ListBalanceRequest) (*ListBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBalance not implemented")
}
func (*UnimplementedBalanceServiceServer) GetBalance(ctx context.Context, req *GetBalanceRequest) (*GetBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (*UnimplementedBalanceServiceServer) AddBalance(ctx context.Context, req *AddBalanceRequest) (*AddBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBalance not implemented")
}

func RegisterBalanceServiceServer(s *grpc.Server, srv BalanceServiceServer) {
	s.RegisterService(&_BalanceService_serviceDesc, srv)
}

func _BalanceService_ListBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalanceServiceServer).ListBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/balance.BalanceService/ListBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalanceServiceServer).ListBalance(ctx, req.(*ListBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BalanceService_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalanceServiceServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/balance.BalanceService/GetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalanceServiceServer).GetBalance(ctx, req.(*GetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BalanceService_AddBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalanceServiceServer).AddBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/balance.BalanceService/AddBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalanceServiceServer).AddBalance(ctx, req.(*AddBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BalanceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "balance.BalanceService",
	HandlerType: (*BalanceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListBalance",
			Handler:    _BalanceService_ListBalance_Handler,
		},
		{
			MethodName: "GetBalance",
			Handler:    _BalanceService_GetBalance_Handler,
		},
		{
			MethodName: "AddBalance",
			Handler:    _BalanceService_AddBalance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "balance/balance.proto",
}
