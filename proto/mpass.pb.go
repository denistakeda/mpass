// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: proto/mpass.proto

package proto

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SignUpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *SignUpRequest) Reset() {
	*x = SignUpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mpass_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpRequest) ProtoMessage() {}

func (x *SignUpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mpass_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpRequest.ProtoReflect.Descriptor instead.
func (*SignUpRequest) Descriptor() ([]byte, []int) {
	return file_proto_mpass_proto_rawDescGZIP(), []int{0}
}

func (x *SignUpRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *SignUpRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SignUpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *SignUpResponse) Reset() {
	*x = SignUpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mpass_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpResponse) ProtoMessage() {}

func (x *SignUpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mpass_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpResponse.ProtoReflect.Descriptor instead.
func (*SignUpResponse) Descriptor() ([]byte, []int) {
	return file_proto_mpass_proto_rawDescGZIP(), []int{1}
}

func (x *SignUpResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type SignInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *SignInRequest) Reset() {
	*x = SignInRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mpass_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInRequest) ProtoMessage() {}

func (x *SignInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mpass_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInRequest.ProtoReflect.Descriptor instead.
func (*SignInRequest) Descriptor() ([]byte, []int) {
	return file_proto_mpass_proto_rawDescGZIP(), []int{2}
}

func (x *SignInRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *SignInRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SignInResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *SignInResponse) Reset() {
	*x = SignInResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mpass_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInResponse) ProtoMessage() {}

func (x *SignInResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mpass_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInResponse.ProtoReflect.Descriptor instead.
func (*SignInResponse) Descriptor() ([]byte, []int) {
	return file_proto_mpass_proto_rawDescGZIP(), []int{3}
}

func (x *SignInResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type AddRecordsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Records []*Record `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
}

func (x *AddRecordsRequest) Reset() {
	*x = AddRecordsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mpass_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRecordsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRecordsRequest) ProtoMessage() {}

func (x *AddRecordsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mpass_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRecordsRequest.ProtoReflect.Descriptor instead.
func (*AddRecordsRequest) Descriptor() ([]byte, []int) {
	return file_proto_mpass_proto_rawDescGZIP(), []int{4}
}

func (x *AddRecordsRequest) GetRecords() []*Record {
	if x != nil {
		return x.Records
	}
	return nil
}

type AllRecordsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Records []*Record `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
}

func (x *AllRecordsResponse) Reset() {
	*x = AllRecordsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mpass_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllRecordsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllRecordsResponse) ProtoMessage() {}

func (x *AllRecordsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mpass_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllRecordsResponse.ProtoReflect.Descriptor instead.
func (*AllRecordsResponse) Descriptor() ([]byte, []int) {
	return file_proto_mpass_proto_rawDescGZIP(), []int{5}
}

func (x *AllRecordsResponse) GetRecords() []*Record {
	if x != nil {
		return x.Records
	}
	return nil
}

type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	LastUpdateDate *timestamp.Timestamp `protobuf:"bytes,2,opt,name=lastUpdateDate,proto3" json:"lastUpdateDate,omitempty"`
	// Types that are assignable to Record:
	//
	//	*Record_LoginPasswordRecord
	//	*Record_TextRecord
	//	*Record_BinaryRecord
	//	*Record_BankCardRecord
	Record isRecord_Record `protobuf_oneof:"record"`
}

func (x *Record) Reset() {
	*x = Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mpass_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mpass_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_proto_mpass_proto_rawDescGZIP(), []int{6}
}

func (x *Record) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Record) GetLastUpdateDate() *timestamp.Timestamp {
	if x != nil {
		return x.LastUpdateDate
	}
	return nil
}

func (m *Record) GetRecord() isRecord_Record {
	if m != nil {
		return m.Record
	}
	return nil
}

func (x *Record) GetLoginPasswordRecord() *LoginPasswordRecord {
	if x, ok := x.GetRecord().(*Record_LoginPasswordRecord); ok {
		return x.LoginPasswordRecord
	}
	return nil
}

func (x *Record) GetTextRecord() *TextRecord {
	if x, ok := x.GetRecord().(*Record_TextRecord); ok {
		return x.TextRecord
	}
	return nil
}

func (x *Record) GetBinaryRecord() *BinaryRecord {
	if x, ok := x.GetRecord().(*Record_BinaryRecord); ok {
		return x.BinaryRecord
	}
	return nil
}

func (x *Record) GetBankCardRecord() *BankCardRecord {
	if x, ok := x.GetRecord().(*Record_BankCardRecord); ok {
		return x.BankCardRecord
	}
	return nil
}

type isRecord_Record interface {
	isRecord_Record()
}

type Record_LoginPasswordRecord struct {
	LoginPasswordRecord *LoginPasswordRecord `protobuf:"bytes,3,opt,name=loginPasswordRecord,proto3,oneof"`
}

type Record_TextRecord struct {
	TextRecord *TextRecord `protobuf:"bytes,4,opt,name=textRecord,proto3,oneof"`
}

type Record_BinaryRecord struct {
	BinaryRecord *BinaryRecord `protobuf:"bytes,5,opt,name=binaryRecord,proto3,oneof"`
}

type Record_BankCardRecord struct {
	BankCardRecord *BankCardRecord `protobuf:"bytes,6,opt,name=bankCardRecord,proto3,oneof"`
}

func (*Record_LoginPasswordRecord) isRecord_Record() {}

func (*Record_TextRecord) isRecord_Record() {}

func (*Record_BinaryRecord) isRecord_Record() {}

func (*Record_BankCardRecord) isRecord_Record() {}

type LoginPasswordRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginPasswordRecord) Reset() {
	*x = LoginPasswordRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mpass_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginPasswordRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginPasswordRecord) ProtoMessage() {}

func (x *LoginPasswordRecord) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mpass_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginPasswordRecord.ProtoReflect.Descriptor instead.
func (*LoginPasswordRecord) Descriptor() ([]byte, []int) {
	return file_proto_mpass_proto_rawDescGZIP(), []int{7}
}

func (x *LoginPasswordRecord) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *LoginPasswordRecord) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type TextRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *TextRecord) Reset() {
	*x = TextRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mpass_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextRecord) ProtoMessage() {}

func (x *TextRecord) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mpass_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextRecord.ProtoReflect.Descriptor instead.
func (*TextRecord) Descriptor() ([]byte, []int) {
	return file_proto_mpass_proto_rawDescGZIP(), []int{8}
}

func (x *TextRecord) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type BinaryRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Binary []byte `protobuf:"bytes,1,opt,name=binary,proto3" json:"binary,omitempty"`
}

func (x *BinaryRecord) Reset() {
	*x = BinaryRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mpass_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BinaryRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinaryRecord) ProtoMessage() {}

func (x *BinaryRecord) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mpass_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinaryRecord.ProtoReflect.Descriptor instead.
func (*BinaryRecord) Descriptor() ([]byte, []int) {
	return file_proto_mpass_proto_rawDescGZIP(), []int{9}
}

func (x *BinaryRecord) GetBinary() []byte {
	if x != nil {
		return x.Binary
	}
	return nil
}

type BankCardRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CardCode string `protobuf:"bytes,1,opt,name=card_code,json=cardCode,proto3" json:"card_code,omitempty"`
	Month    uint32 `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty"`
	Day      uint32 `protobuf:"varint,3,opt,name=day,proto3" json:"day,omitempty"`
	Code     uint32 `protobuf:"varint,4,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *BankCardRecord) Reset() {
	*x = BankCardRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mpass_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankCardRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankCardRecord) ProtoMessage() {}

func (x *BankCardRecord) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mpass_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankCardRecord.ProtoReflect.Descriptor instead.
func (*BankCardRecord) Descriptor() ([]byte, []int) {
	return file_proto_mpass_proto_rawDescGZIP(), []int{10}
}

func (x *BankCardRecord) GetCardCode() string {
	if x != nil {
		return x.CardCode
	}
	return ""
}

func (x *BankCardRecord) GetMonth() uint32 {
	if x != nil {
		return x.Month
	}
	return 0
}

func (x *BankCardRecord) GetDay() uint32 {
	if x != nil {
		return x.Day
	}
	return 0
}

func (x *BankCardRecord) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

var File_proto_mpass_proto protoreflect.FileDescriptor

var file_proto_mpass_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x70, 0x61, 0x73, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x26, 0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e,
	0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x41, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x22, 0x26, 0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x39, 0x0a, 0x11, 0x41,
	0x64, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x24, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x07, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x3a, 0x0a, 0x12, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x07,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x70, 0x62, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x22, 0xdb, 0x02, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x42, 0x0a,
	0x0e, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x4b, 0x0a, 0x13, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x48, 0x00, 0x52, 0x13, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x30,
	0x0a, 0x0a, 0x74, 0x65, 0x78, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x48, 0x00, 0x52, 0x0a, 0x74, 0x65, 0x78, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x12, 0x36, 0x0a, 0x0c, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x69, 0x6e, 0x61,
	0x72, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x48, 0x00, 0x52, 0x0c, 0x62, 0x69, 0x6e, 0x61,
	0x72, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x3c, 0x0a, 0x0e, 0x62, 0x61, 0x6e, 0x6b,
	0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x48, 0x00, 0x52, 0x0e, 0x62, 0x61, 0x6e, 0x6b, 0x43, 0x61, 0x72, 0x64,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x42, 0x08, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x22, 0x47, 0x0a, 0x13, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x20, 0x0a, 0x0a, 0x54, 0x65, 0x78,
	0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x26, 0x0a, 0x0c, 0x42,
	0x69, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x62,
	0x69, 0x6e, 0x61, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x62, 0x69, 0x6e,
	0x61, 0x72, 0x79, 0x22, 0x69, 0x0a, 0x0e, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x61, 0x72, 0x64, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x72, 0x64, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x64, 0x61, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x32, 0xeb,
	0x01, 0x0a, 0x0c, 0x4d, 0x70, 0x61, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x2f, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x53,
	0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70,
	0x62, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2f, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x70, 0x62, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12,
	0x15, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3c,
	0x0a, 0x0a, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x24, 0x5a, 0x22,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x6e, 0x69, 0x73,
	0x74, 0x61, 0x6b, 0x65, 0x64, 0x61, 0x2f, 0x6d, 0x70, 0x61, 0x73, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_mpass_proto_rawDescOnce sync.Once
	file_proto_mpass_proto_rawDescData = file_proto_mpass_proto_rawDesc
)

func file_proto_mpass_proto_rawDescGZIP() []byte {
	file_proto_mpass_proto_rawDescOnce.Do(func() {
		file_proto_mpass_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_mpass_proto_rawDescData)
	})
	return file_proto_mpass_proto_rawDescData
}

var file_proto_mpass_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_mpass_proto_goTypes = []interface{}{
	(*SignUpRequest)(nil),       // 0: pb.SignUpRequest
	(*SignUpResponse)(nil),      // 1: pb.SignUpResponse
	(*SignInRequest)(nil),       // 2: pb.SignInRequest
	(*SignInResponse)(nil),      // 3: pb.SignInResponse
	(*AddRecordsRequest)(nil),   // 4: pb.AddRecordsRequest
	(*AllRecordsResponse)(nil),  // 5: pb.AllRecordsResponse
	(*Record)(nil),              // 6: pb.Record
	(*LoginPasswordRecord)(nil), // 7: pb.LoginPasswordRecord
	(*TextRecord)(nil),          // 8: pb.TextRecord
	(*BinaryRecord)(nil),        // 9: pb.BinaryRecord
	(*BankCardRecord)(nil),      // 10: pb.BankCardRecord
	(*timestamp.Timestamp)(nil), // 11: google.protobuf.Timestamp
	(*empty.Empty)(nil),         // 12: google.protobuf.Empty
}
var file_proto_mpass_proto_depIdxs = []int32{
	6,  // 0: pb.AddRecordsRequest.records:type_name -> pb.Record
	6,  // 1: pb.AllRecordsResponse.records:type_name -> pb.Record
	11, // 2: pb.Record.lastUpdateDate:type_name -> google.protobuf.Timestamp
	7,  // 3: pb.Record.loginPasswordRecord:type_name -> pb.LoginPasswordRecord
	8,  // 4: pb.Record.textRecord:type_name -> pb.TextRecord
	9,  // 5: pb.Record.binaryRecord:type_name -> pb.BinaryRecord
	10, // 6: pb.Record.bankCardRecord:type_name -> pb.BankCardRecord
	0,  // 7: pb.MpassService.SignUp:input_type -> pb.SignUpRequest
	2,  // 8: pb.MpassService.SignIn:input_type -> pb.SignInRequest
	4,  // 9: pb.MpassService.AddRecords:input_type -> pb.AddRecordsRequest
	12, // 10: pb.MpassService.AllRecords:input_type -> google.protobuf.Empty
	1,  // 11: pb.MpassService.SignUp:output_type -> pb.SignUpResponse
	3,  // 12: pb.MpassService.SignIn:output_type -> pb.SignInResponse
	12, // 13: pb.MpassService.AddRecords:output_type -> google.protobuf.Empty
	5,  // 14: pb.MpassService.AllRecords:output_type -> pb.AllRecordsResponse
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_proto_mpass_proto_init() }
func file_proto_mpass_proto_init() {
	if File_proto_mpass_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_mpass_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_mpass_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_mpass_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_mpass_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_mpass_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRecordsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_mpass_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllRecordsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_mpass_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_mpass_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginPasswordRecord); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_mpass_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextRecord); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_mpass_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BinaryRecord); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_mpass_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankCardRecord); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_proto_mpass_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*Record_LoginPasswordRecord)(nil),
		(*Record_TextRecord)(nil),
		(*Record_BinaryRecord)(nil),
		(*Record_BankCardRecord)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_mpass_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_mpass_proto_goTypes,
		DependencyIndexes: file_proto_mpass_proto_depIdxs,
		MessageInfos:      file_proto_mpass_proto_msgTypes,
	}.Build()
	File_proto_mpass_proto = out.File
	file_proto_mpass_proto_rawDesc = nil
	file_proto_mpass_proto_goTypes = nil
	file_proto_mpass_proto_depIdxs = nil
}
