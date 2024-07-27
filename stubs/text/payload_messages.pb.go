// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: text/payload_messages.proto

package text

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TextBaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *TextBaseResponse) Reset() {
	*x = TextBaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_text_payload_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextBaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextBaseResponse) ProtoMessage() {}

func (x *TextBaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_text_payload_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextBaseResponse.ProtoReflect.Descriptor instead.
func (*TextBaseResponse) Descriptor() ([]byte, []int) {
	return file_text_payload_messages_proto_rawDescGZIP(), []int{0}
}

func (x *TextBaseResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CreateTextRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActivityId string `protobuf:"bytes,1,opt,name=activity_id,proto3" json:"activity_id,omitempty"`
	Text       string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *CreateTextRequest) Reset() {
	*x = CreateTextRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_text_payload_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTextRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTextRequest) ProtoMessage() {}

func (x *CreateTextRequest) ProtoReflect() protoreflect.Message {
	mi := &file_text_payload_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTextRequest.ProtoReflect.Descriptor instead.
func (*CreateTextRequest) Descriptor() ([]byte, []int) {
	return file_text_payload_messages_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTextRequest) GetActivityId() string {
	if x != nil {
		return x.ActivityId
	}
	return ""
}

func (x *CreateTextRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type GetAllTextByActivityIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActivityId   string  `protobuf:"bytes,1,opt,name=activity_id,proto3" json:"activity_id,omitempty"`
	Search       *string `protobuf:"bytes,2,opt,name=search,proto3,oneof" json:"search,omitempty"`
	Page         *int32  `protobuf:"varint,3,opt,name=page,proto3,oneof" json:"page,omitempty"`
	Limit        *int32  `protobuf:"varint,4,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
	IsNewest     *bool   `protobuf:"varint,7,opt,name=is_newest,proto3,oneof" json:"is_newest,omitempty"`
	IsOldest     *bool   `protobuf:"varint,8,opt,name=is_oldest,proto3,oneof" json:"is_oldest,omitempty"`
	IsAscending  *bool   `protobuf:"varint,9,opt,name=is_ascending,proto3,oneof" json:"is_ascending,omitempty"`
	IsDescending *bool   `protobuf:"varint,10,opt,name=is_descending,proto3,oneof" json:"is_descending,omitempty"`
}

func (x *GetAllTextByActivityIDRequest) Reset() {
	*x = GetAllTextByActivityIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_text_payload_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTextByActivityIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTextByActivityIDRequest) ProtoMessage() {}

func (x *GetAllTextByActivityIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_text_payload_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTextByActivityIDRequest.ProtoReflect.Descriptor instead.
func (*GetAllTextByActivityIDRequest) Descriptor() ([]byte, []int) {
	return file_text_payload_messages_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllTextByActivityIDRequest) GetActivityId() string {
	if x != nil {
		return x.ActivityId
	}
	return ""
}

func (x *GetAllTextByActivityIDRequest) GetSearch() string {
	if x != nil && x.Search != nil {
		return *x.Search
	}
	return ""
}

func (x *GetAllTextByActivityIDRequest) GetPage() int32 {
	if x != nil && x.Page != nil {
		return *x.Page
	}
	return 0
}

func (x *GetAllTextByActivityIDRequest) GetLimit() int32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

func (x *GetAllTextByActivityIDRequest) GetIsNewest() bool {
	if x != nil && x.IsNewest != nil {
		return *x.IsNewest
	}
	return false
}

func (x *GetAllTextByActivityIDRequest) GetIsOldest() bool {
	if x != nil && x.IsOldest != nil {
		return *x.IsOldest
	}
	return false
}

func (x *GetAllTextByActivityIDRequest) GetIsAscending() bool {
	if x != nil && x.IsAscending != nil {
		return *x.IsAscending
	}
	return false
}

func (x *GetAllTextByActivityIDRequest) GetIsDescending() bool {
	if x != nil && x.IsDescending != nil {
		return *x.IsDescending
	}
	return false
}

type TextPaging struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentPage int32 `protobuf:"varint,1,opt,name=current_page,proto3" json:"current_page,omitempty"`
	TotalPage   int32 `protobuf:"varint,2,opt,name=total_page,proto3" json:"total_page,omitempty"`
	Count       int32 `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *TextPaging) Reset() {
	*x = TextPaging{}
	if protoimpl.UnsafeEnabled {
		mi := &file_text_payload_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextPaging) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextPaging) ProtoMessage() {}

func (x *TextPaging) ProtoReflect() protoreflect.Message {
	mi := &file_text_payload_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextPaging.ProtoReflect.Descriptor instead.
func (*TextPaging) Descriptor() ([]byte, []int) {
	return file_text_payload_messages_proto_rawDescGZIP(), []int{3}
}

func (x *TextPaging) GetCurrentPage() int32 {
	if x != nil {
		return x.CurrentPage
	}
	return 0
}

func (x *TextPaging) GetTotalPage() int32 {
	if x != nil {
		return x.TotalPage
	}
	return 0
}

func (x *TextPaging) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type GetAllTextByActivityIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Texts   []*GetTextByIDResponse `protobuf:"bytes,2,rep,name=texts,proto3" json:"texts,omitempty"`
	Paging  *TextPaging            `protobuf:"bytes,3,opt,name=paging,proto3" json:"paging,omitempty"`
}

func (x *GetAllTextByActivityIDResponse) Reset() {
	*x = GetAllTextByActivityIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_text_payload_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTextByActivityIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTextByActivityIDResponse) ProtoMessage() {}

func (x *GetAllTextByActivityIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_text_payload_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTextByActivityIDResponse.ProtoReflect.Descriptor instead.
func (*GetAllTextByActivityIDResponse) Descriptor() ([]byte, []int) {
	return file_text_payload_messages_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllTextByActivityIDResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllTextByActivityIDResponse) GetTexts() []*GetTextByIDResponse {
	if x != nil {
		return x.Texts
	}
	return nil
}

func (x *GetAllTextByActivityIDResponse) GetPaging() *TextPaging {
	if x != nil {
		return x.Paging
	}
	return nil
}

type GetTextByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTextByIDRequest) Reset() {
	*x = GetTextByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_text_payload_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTextByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTextByIDRequest) ProtoMessage() {}

func (x *GetTextByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_text_payload_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTextByIDRequest.ProtoReflect.Descriptor instead.
func (*GetTextByIDRequest) Descriptor() ([]byte, []int) {
	return file_text_payload_messages_proto_rawDescGZIP(), []int{5}
}

func (x *GetTextByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetTextByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ActivityId string                 `protobuf:"bytes,2,opt,name=activity_id,proto3" json:"activity_id,omitempty"`
	Text       string                 `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,proto3" json:"created_at,omitempty"`
	UpdatedAt  *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updated_at,proto3" json:"updated_at,omitempty"`
	DeletedAt  *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=deleted_at,proto3,oneof" json:"deleted_at,omitempty"`
}

func (x *GetTextByIDResponse) Reset() {
	*x = GetTextByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_text_payload_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTextByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTextByIDResponse) ProtoMessage() {}

func (x *GetTextByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_text_payload_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTextByIDResponse.ProtoReflect.Descriptor instead.
func (*GetTextByIDResponse) Descriptor() ([]byte, []int) {
	return file_text_payload_messages_proto_rawDescGZIP(), []int{6}
}

func (x *GetTextByIDResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetTextByIDResponse) GetActivityId() string {
	if x != nil {
		return x.ActivityId
	}
	return ""
}

func (x *GetTextByIDResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *GetTextByIDResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *GetTextByIDResponse) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *GetTextByIDResponse) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type UpdateTextByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Text *string `protobuf:"bytes,2,opt,name=text,proto3,oneof" json:"text,omitempty"`
}

func (x *UpdateTextByIDRequest) Reset() {
	*x = UpdateTextByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_text_payload_messages_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTextByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTextByIDRequest) ProtoMessage() {}

func (x *UpdateTextByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_text_payload_messages_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTextByIDRequest.ProtoReflect.Descriptor instead.
func (*UpdateTextByIDRequest) Descriptor() ([]byte, []int) {
	return file_text_payload_messages_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateTextByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateTextByIDRequest) GetText() string {
	if x != nil && x.Text != nil {
		return *x.Text
	}
	return ""
}

type DeleteTextByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteTextByIDRequest) Reset() {
	*x = DeleteTextByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_text_payload_messages_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTextByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTextByIDRequest) ProtoMessage() {}

func (x *DeleteTextByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_text_payload_messages_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTextByIDRequest.ProtoReflect.Descriptor instead.
func (*DeleteTextByIDRequest) Descriptor() ([]byte, []int) {
	return file_text_payload_messages_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteTextByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_text_payload_messages_proto protoreflect.FileDescriptor

var file_text_payload_messages_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x74, 0x65, 0x78, 0x74, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x10, 0x54, 0x65, 0x78, 0x74, 0x42, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x49, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x78,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x89,
	0x03, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x65, 0x78, 0x74, 0x42, 0x79, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f,
	0x69, 0x64, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x88, 0x01, 0x01, 0x12,
	0x17, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x6e, 0x65, 0x77, 0x65, 0x73, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x48, 0x03, 0x52, 0x09, 0x69, 0x73, 0x5f, 0x6e, 0x65, 0x77,
	0x65, 0x73, 0x74, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x6f, 0x6c, 0x64,
	0x65, 0x73, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x48, 0x04, 0x52, 0x09, 0x69, 0x73, 0x5f,
	0x6f, 0x6c, 0x64, 0x65, 0x73, 0x74, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0c, 0x69, 0x73, 0x5f,
	0x61, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x48,
	0x05, 0x52, 0x0c, 0x69, 0x73, 0x5f, 0x61, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x88,
	0x01, 0x01, 0x12, 0x29, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x65, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x48, 0x06, 0x52, 0x0d, 0x69, 0x73, 0x5f,
	0x64, 0x65, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a,
	0x07, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x70, 0x61, 0x67,
	0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x0c, 0x0a, 0x0a, 0x5f,
	0x69, 0x73, 0x5f, 0x6e, 0x65, 0x77, 0x65, 0x73, 0x74, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x69, 0x73,
	0x5f, 0x6f, 0x6c, 0x64, 0x65, 0x73, 0x74, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x69, 0x73, 0x5f, 0x61,
	0x73, 0x63, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x69, 0x73, 0x5f,
	0x64, 0x65, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x66, 0x0a, 0x0a, 0x54, 0x65,
	0x78, 0x74, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x97, 0x01, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x65, 0x78,
	0x74, 0x42, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x30, 0x0a, 0x05, 0x74, 0x65, 0x78, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x78, 0x74, 0x42, 0x79,
	0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x74, 0x65, 0x78, 0x74,
	0x73, 0x12, 0x29, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x50, 0x61,
	0x67, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x22, 0x24, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x54, 0x65, 0x78, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0xa3, 0x02, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x54, 0x65, 0x78, 0x74, 0x42, 0x79,
	0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x12, 0x3a, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x12, 0x3a, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x12, 0x3f, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x0a, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x22, 0x49, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x17, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x74,
	0x65, 0x78, 0x74, 0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x78,
	0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x42, 0x2d, 0x5a, 0x2b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x73,
	0x61, 0x74, 0x61, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2f, 0x73, 0x74, 0x75, 0x62, 0x73, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_text_payload_messages_proto_rawDescOnce sync.Once
	file_text_payload_messages_proto_rawDescData = file_text_payload_messages_proto_rawDesc
)

func file_text_payload_messages_proto_rawDescGZIP() []byte {
	file_text_payload_messages_proto_rawDescOnce.Do(func() {
		file_text_payload_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_text_payload_messages_proto_rawDescData)
	})
	return file_text_payload_messages_proto_rawDescData
}

var file_text_payload_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_text_payload_messages_proto_goTypes = []any{
	(*TextBaseResponse)(nil),               // 0: proto.TextBaseResponse
	(*CreateTextRequest)(nil),              // 1: proto.CreateTextRequest
	(*GetAllTextByActivityIDRequest)(nil),  // 2: proto.GetAllTextByActivityIDRequest
	(*TextPaging)(nil),                     // 3: proto.TextPaging
	(*GetAllTextByActivityIDResponse)(nil), // 4: proto.GetAllTextByActivityIDResponse
	(*GetTextByIDRequest)(nil),             // 5: proto.GetTextByIDRequest
	(*GetTextByIDResponse)(nil),            // 6: proto.GetTextByIDResponse
	(*UpdateTextByIDRequest)(nil),          // 7: proto.UpdateTextByIDRequest
	(*DeleteTextByIDRequest)(nil),          // 8: proto.DeleteTextByIDRequest
	(*timestamppb.Timestamp)(nil),          // 9: google.protobuf.Timestamp
}
var file_text_payload_messages_proto_depIdxs = []int32{
	6, // 0: proto.GetAllTextByActivityIDResponse.texts:type_name -> proto.GetTextByIDResponse
	3, // 1: proto.GetAllTextByActivityIDResponse.paging:type_name -> proto.TextPaging
	9, // 2: proto.GetTextByIDResponse.created_at:type_name -> google.protobuf.Timestamp
	9, // 3: proto.GetTextByIDResponse.updated_at:type_name -> google.protobuf.Timestamp
	9, // 4: proto.GetTextByIDResponse.deleted_at:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_text_payload_messages_proto_init() }
func file_text_payload_messages_proto_init() {
	if File_text_payload_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_text_payload_messages_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TextBaseResponse); i {
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
		file_text_payload_messages_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateTextRequest); i {
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
		file_text_payload_messages_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllTextByActivityIDRequest); i {
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
		file_text_payload_messages_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*TextPaging); i {
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
		file_text_payload_messages_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllTextByActivityIDResponse); i {
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
		file_text_payload_messages_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetTextByIDRequest); i {
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
		file_text_payload_messages_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetTextByIDResponse); i {
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
		file_text_payload_messages_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateTextByIDRequest); i {
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
		file_text_payload_messages_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteTextByIDRequest); i {
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
	file_text_payload_messages_proto_msgTypes[2].OneofWrappers = []any{}
	file_text_payload_messages_proto_msgTypes[6].OneofWrappers = []any{}
	file_text_payload_messages_proto_msgTypes[7].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_text_payload_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_text_payload_messages_proto_goTypes,
		DependencyIndexes: file_text_payload_messages_proto_depIdxs,
		MessageInfos:      file_text_payload_messages_proto_msgTypes,
	}.Build()
	File_text_payload_messages_proto = out.File
	file_text_payload_messages_proto_rawDesc = nil
	file_text_payload_messages_proto_goTypes = nil
	file_text_payload_messages_proto_depIdxs = nil
}
