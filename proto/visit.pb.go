// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.3
// source: proto/visit.proto

package com_alfa_grpc_visit

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Visit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BuildingId   string `protobuf:"bytes,2,opt,name=buildingId,proto3" json:"buildingId,omitempty"`
	UserId       string `protobuf:"bytes,3,opt,name=userId,proto3" json:"userId,omitempty"`
	CheckinTime  string `protobuf:"bytes,4,opt,name=checkinTime,proto3" json:"checkinTime,omitempty"`
	CheckoutTime string `protobuf:"bytes,5,opt,name=checkoutTime,proto3" json:"checkoutTime,omitempty"`
}

func (x *Visit) Reset() {
	*x = Visit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_visit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Visit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Visit) ProtoMessage() {}

func (x *Visit) ProtoReflect() protoreflect.Message {
	mi := &file_proto_visit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Visit.ProtoReflect.Descriptor instead.
func (*Visit) Descriptor() ([]byte, []int) {
	return file_proto_visit_proto_rawDescGZIP(), []int{0}
}

func (x *Visit) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Visit) GetBuildingId() string {
	if x != nil {
		return x.BuildingId
	}
	return ""
}

func (x *Visit) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Visit) GetCheckinTime() string {
	if x != nil {
		return x.CheckinTime
	}
	return ""
}

func (x *Visit) GetCheckoutTime() string {
	if x != nil {
		return x.CheckoutTime
	}
	return ""
}

type VisitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BuildingId string `protobuf:"bytes,1,opt,name=buildingId,proto3" json:"buildingId,omitempty"`
	UserId     string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *VisitRequest) Reset() {
	*x = VisitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_visit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VisitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VisitRequest) ProtoMessage() {}

func (x *VisitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_visit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VisitRequest.ProtoReflect.Descriptor instead.
func (*VisitRequest) Descriptor() ([]byte, []int) {
	return file_proto_visit_proto_rawDescGZIP(), []int{1}
}

func (x *VisitRequest) GetBuildingId() string {
	if x != nil {
		return x.BuildingId
	}
	return ""
}

func (x *VisitRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type VisitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *VisitResponse) Reset() {
	*x = VisitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_visit_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VisitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VisitResponse) ProtoMessage() {}

func (x *VisitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_visit_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VisitResponse.ProtoReflect.Descriptor instead.
func (*VisitResponse) Descriptor() ([]byte, []int) {
	return file_proto_visit_proto_rawDescGZIP(), []int{2}
}

func (x *VisitResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type VisitListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *VisitListRequest) Reset() {
	*x = VisitListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_visit_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VisitListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VisitListRequest) ProtoMessage() {}

func (x *VisitListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_visit_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VisitListRequest.ProtoReflect.Descriptor instead.
func (*VisitListRequest) Descriptor() ([]byte, []int) {
	return file_proto_visit_proto_rawDescGZIP(), []int{3}
}

func (x *VisitListRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type VisitList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Visit []*Visit `protobuf:"bytes,1,rep,name=visit,proto3" json:"visit,omitempty"`
}

func (x *VisitList) Reset() {
	*x = VisitList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_visit_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VisitList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VisitList) ProtoMessage() {}

func (x *VisitList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_visit_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VisitList.ProtoReflect.Descriptor instead.
func (*VisitList) Descriptor() ([]byte, []int) {
	return file_proto_visit_proto_rawDescGZIP(), []int{4}
}

func (x *VisitList) GetVisit() []*Visit {
	if x != nil {
		return x.Visit
	}
	return nil
}

type CheckoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *CheckoutRequest) Reset() {
	*x = CheckoutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_visit_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckoutRequest) ProtoMessage() {}

func (x *CheckoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_visit_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckoutRequest.ProtoReflect.Descriptor instead.
func (*CheckoutRequest) Descriptor() ([]byte, []int) {
	return file_proto_visit_proto_rawDescGZIP(), []int{5}
}

func (x *CheckoutRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_proto_visit_proto protoreflect.FileDescriptor

var file_proto_visit_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x69, 0x73, 0x69, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6c, 0x66, 0x61, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x74, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x01, 0x0a, 0x05, 0x56, 0x69, 0x73, 0x69, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x6f, 0x75, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x46, 0x0a,
	0x0c, 0x56, 0x69, 0x73, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x25, 0x0a, 0x0d, 0x56, 0x69, 0x73, 0x69, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x28, 0x0a, 0x10,
	0x56, 0x69, 0x73, 0x69, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3d, 0x0a, 0x09, 0x56, 0x69, 0x73, 0x69, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x05, 0x76, 0x69, 0x73, 0x69, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6c, 0x66, 0x61, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x74, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x74, 0x52, 0x05,
	0x76, 0x69, 0x73, 0x69, 0x74, 0x22, 0x27, 0x0a, 0x0f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xfb,
	0x01, 0x0a, 0x0c, 0x56, 0x69, 0x73, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x4e, 0x0a, 0x03, 0x61, 0x64, 0x64, 0x12, 0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6c, 0x66,
	0x61, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x74, 0x2e, 0x56, 0x69, 0x73,
	0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x61, 0x6c, 0x66, 0x61, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x74, 0x2e,
	0x56, 0x69, 0x73, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x4f, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x25, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6c,
	0x66, 0x61, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x74, 0x2e, 0x56, 0x69,
	0x73, 0x69, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6c, 0x66, 0x61, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x76,
	0x69, 0x73, 0x69, 0x74, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00,
	0x12, 0x4a, 0x0a, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x12, 0x24, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x61, 0x6c, 0x66, 0x61, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x69, 0x73,
	0x69, 0x74, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x02, 0x50, 0x01,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_visit_proto_rawDescOnce sync.Once
	file_proto_visit_proto_rawDescData = file_proto_visit_proto_rawDesc
)

func file_proto_visit_proto_rawDescGZIP() []byte {
	file_proto_visit_proto_rawDescOnce.Do(func() {
		file_proto_visit_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_visit_proto_rawDescData)
	})
	return file_proto_visit_proto_rawDescData
}

var file_proto_visit_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_visit_proto_goTypes = []interface{}{
	(*Visit)(nil),            // 0: com.alfa.grpc.visit.Visit
	(*VisitRequest)(nil),     // 1: com.alfa.grpc.visit.VisitRequest
	(*VisitResponse)(nil),    // 2: com.alfa.grpc.visit.VisitResponse
	(*VisitListRequest)(nil), // 3: com.alfa.grpc.visit.VisitListRequest
	(*VisitList)(nil),        // 4: com.alfa.grpc.visit.VisitList
	(*CheckoutRequest)(nil),  // 5: com.alfa.grpc.visit.CheckoutRequest
	(*empty.Empty)(nil),      // 6: google.protobuf.Empty
}
var file_proto_visit_proto_depIdxs = []int32{
	0, // 0: com.alfa.grpc.visit.VisitList.visit:type_name -> com.alfa.grpc.visit.Visit
	1, // 1: com.alfa.grpc.visit.VisitService.add:input_type -> com.alfa.grpc.visit.VisitRequest
	3, // 2: com.alfa.grpc.visit.VisitService.list:input_type -> com.alfa.grpc.visit.VisitListRequest
	5, // 3: com.alfa.grpc.visit.VisitService.checkout:input_type -> com.alfa.grpc.visit.CheckoutRequest
	2, // 4: com.alfa.grpc.visit.VisitService.add:output_type -> com.alfa.grpc.visit.VisitResponse
	4, // 5: com.alfa.grpc.visit.VisitService.list:output_type -> com.alfa.grpc.visit.VisitList
	6, // 6: com.alfa.grpc.visit.VisitService.checkout:output_type -> google.protobuf.Empty
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_visit_proto_init() }
func file_proto_visit_proto_init() {
	if File_proto_visit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_visit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Visit); i {
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
		file_proto_visit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VisitRequest); i {
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
		file_proto_visit_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VisitResponse); i {
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
		file_proto_visit_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VisitListRequest); i {
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
		file_proto_visit_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VisitList); i {
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
		file_proto_visit_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckoutRequest); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_visit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_visit_proto_goTypes,
		DependencyIndexes: file_proto_visit_proto_depIdxs,
		MessageInfos:      file_proto_visit_proto_msgTypes,
	}.Build()
	File_proto_visit_proto = out.File
	file_proto_visit_proto_rawDesc = nil
	file_proto_visit_proto_goTypes = nil
	file_proto_visit_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// VisitServiceClient is the client API for VisitService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VisitServiceClient interface {
	Add(ctx context.Context, in *VisitRequest, opts ...grpc.CallOption) (*VisitResponse, error)
	List(ctx context.Context, in *VisitListRequest, opts ...grpc.CallOption) (*VisitList, error)
	Checkout(ctx context.Context, in *CheckoutRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type visitServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVisitServiceClient(cc grpc.ClientConnInterface) VisitServiceClient {
	return &visitServiceClient{cc}
}

func (c *visitServiceClient) Add(ctx context.Context, in *VisitRequest, opts ...grpc.CallOption) (*VisitResponse, error) {
	out := new(VisitResponse)
	err := c.cc.Invoke(ctx, "/com.alfa.grpc.visit.VisitService/add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *visitServiceClient) List(ctx context.Context, in *VisitListRequest, opts ...grpc.CallOption) (*VisitList, error) {
	out := new(VisitList)
	err := c.cc.Invoke(ctx, "/com.alfa.grpc.visit.VisitService/list", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *visitServiceClient) Checkout(ctx context.Context, in *CheckoutRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/com.alfa.grpc.visit.VisitService/checkout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VisitServiceServer is the server API for VisitService service.
type VisitServiceServer interface {
	Add(context.Context, *VisitRequest) (*VisitResponse, error)
	List(context.Context, *VisitListRequest) (*VisitList, error)
	Checkout(context.Context, *CheckoutRequest) (*empty.Empty, error)
}

// UnimplementedVisitServiceServer can be embedded to have forward compatible implementations.
type UnimplementedVisitServiceServer struct {
}

func (*UnimplementedVisitServiceServer) Add(context.Context, *VisitRequest) (*VisitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedVisitServiceServer) List(context.Context, *VisitListRequest) (*VisitList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedVisitServiceServer) Checkout(context.Context, *CheckoutRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Checkout not implemented")
}

func RegisterVisitServiceServer(s *grpc.Server, srv VisitServiceServer) {
	s.RegisterService(&_VisitService_serviceDesc, srv)
}

func _VisitService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VisitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VisitServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.alfa.grpc.visit.VisitService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VisitServiceServer).Add(ctx, req.(*VisitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VisitService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VisitListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VisitServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.alfa.grpc.visit.VisitService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VisitServiceServer).List(ctx, req.(*VisitListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VisitService_Checkout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VisitServiceServer).Checkout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.alfa.grpc.visit.VisitService/Checkout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VisitServiceServer).Checkout(ctx, req.(*CheckoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VisitService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.alfa.grpc.visit.VisitService",
	HandlerType: (*VisitServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "add",
			Handler:    _VisitService_Add_Handler,
		},
		{
			MethodName: "list",
			Handler:    _VisitService_List_Handler,
		},
		{
			MethodName: "checkout",
			Handler:    _VisitService_Checkout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/visit.proto",
}
