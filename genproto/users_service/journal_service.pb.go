// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: journal_service.proto

package users_service

import (
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

type JournalPrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *JournalPrimaryKey) Reset() {
	*x = JournalPrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_journal_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JournalPrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JournalPrimaryKey) ProtoMessage() {}

func (x *JournalPrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_journal_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JournalPrimaryKey.ProtoReflect.Descriptor instead.
func (*JournalPrimaryKey) Descriptor() ([]byte, []int) {
	return file_journal_service_proto_rawDescGZIP(), []int{0}
}

func (x *JournalPrimaryKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Journal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	GroupId   string `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	From      string `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	To        string `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	StudentId string `protobuf:"bytes,5,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	CreatedAt string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Journal) Reset() {
	*x = Journal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_journal_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Journal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Journal) ProtoMessage() {}

func (x *Journal) ProtoReflect() protoreflect.Message {
	mi := &file_journal_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Journal.ProtoReflect.Descriptor instead.
func (*Journal) Descriptor() ([]byte, []int) {
	return file_journal_service_proto_rawDescGZIP(), []int{1}
}

func (x *Journal) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Journal) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *Journal) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Journal) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Journal) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

func (x *Journal) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Journal) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type CreateJournal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId   string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	From      string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To        string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	StudentId string `protobuf:"bytes,4,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
}

func (x *CreateJournal) Reset() {
	*x = CreateJournal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_journal_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateJournal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJournal) ProtoMessage() {}

func (x *CreateJournal) ProtoReflect() protoreflect.Message {
	mi := &file_journal_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJournal.ProtoReflect.Descriptor instead.
func (*CreateJournal) Descriptor() ([]byte, []int) {
	return file_journal_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateJournal) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *CreateJournal) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *CreateJournal) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *CreateJournal) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

type UpdateJournal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	GroupId   string `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	From      string `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	To        string `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	StudentId string `protobuf:"bytes,5,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
}

func (x *UpdateJournal) Reset() {
	*x = UpdateJournal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_journal_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateJournal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateJournal) ProtoMessage() {}

func (x *UpdateJournal) ProtoReflect() protoreflect.Message {
	mi := &file_journal_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateJournal.ProtoReflect.Descriptor instead.
func (*UpdateJournal) Descriptor() ([]byte, []int) {
	return file_journal_service_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateJournal) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateJournal) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *UpdateJournal) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *UpdateJournal) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *UpdateJournal) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

type GetListJournalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int64  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *GetListJournalRequest) Reset() {
	*x = GetListJournalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_journal_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListJournalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListJournalRequest) ProtoMessage() {}

func (x *GetListJournalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_journal_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListJournalRequest.ProtoReflect.Descriptor instead.
func (*GetListJournalRequest) Descriptor() ([]byte, []int) {
	return file_journal_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetListJournalRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetListJournalRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetListJournalRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type GetListJournalResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count    int64      `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Journals []*Journal `protobuf:"bytes,2,rep,name=journals,proto3" json:"journals,omitempty"`
}

func (x *GetListJournalResponse) Reset() {
	*x = GetListJournalResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_journal_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListJournalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListJournalResponse) ProtoMessage() {}

func (x *GetListJournalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_journal_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListJournalResponse.ProtoReflect.Descriptor instead.
func (*GetListJournalResponse) Descriptor() ([]byte, []int) {
	return file_journal_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetListJournalResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetListJournalResponse) GetJournals() []*Journal {
	if x != nil {
		return x.Journals
	}
	return nil
}

var File_journal_service_proto protoreflect.FileDescriptor

var file_journal_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x15, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a,
	0x11, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b,
	0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0xb5, 0x01, 0x0a, 0x07, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a,
	0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x6d, 0x0a, 0x0d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x7d, 0x0a, 0x0d, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x5d, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x62, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x08, 0x6a, 0x6f, 0x75, 0x72, 0x6e,
	0x61, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61,
	0x6c, 0x52, 0x08, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x32, 0xf9, 0x02, 0x0a, 0x0e,
	0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a,
	0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x1a, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x22, 0x00,
	0x12, 0x45, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x20, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4a, 0x6f, 0x75, 0x72,
	0x6e, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x16, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4a, 0x6f,
	0x75, 0x72, 0x6e, 0x61, 0x6c, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x24, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x40, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x1a, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61,
	0x6c, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x20, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4a, 0x6f,
	0x75, 0x72, 0x6e, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a,
	0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x18, 0x5a, 0x16, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_journal_service_proto_rawDescOnce sync.Once
	file_journal_service_proto_rawDescData = file_journal_service_proto_rawDesc
)

func file_journal_service_proto_rawDescGZIP() []byte {
	file_journal_service_proto_rawDescOnce.Do(func() {
		file_journal_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_journal_service_proto_rawDescData)
	})
	return file_journal_service_proto_rawDescData
}

var file_journal_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_journal_service_proto_goTypes = []interface{}{
	(*JournalPrimaryKey)(nil),      // 0: users_service.JournalPrimaryKey
	(*Journal)(nil),                // 1: users_service.Journal
	(*CreateJournal)(nil),          // 2: users_service.CreateJournal
	(*UpdateJournal)(nil),          // 3: users_service.UpdateJournal
	(*GetListJournalRequest)(nil),  // 4: users_service.GetListJournalRequest
	(*GetListJournalResponse)(nil), // 5: users_service.GetListJournalResponse
	(*Empty)(nil),                  // 6: users_service.Empty
}
var file_journal_service_proto_depIdxs = []int32{
	1, // 0: users_service.GetListJournalResponse.journals:type_name -> users_service.Journal
	2, // 1: users_service.JournalService.Create:input_type -> users_service.CreateJournal
	0, // 2: users_service.JournalService.GetById:input_type -> users_service.JournalPrimaryKey
	4, // 3: users_service.JournalService.GetList:input_type -> users_service.GetListJournalRequest
	3, // 4: users_service.JournalService.Update:input_type -> users_service.UpdateJournal
	0, // 5: users_service.JournalService.Delete:input_type -> users_service.JournalPrimaryKey
	1, // 6: users_service.JournalService.Create:output_type -> users_service.Journal
	1, // 7: users_service.JournalService.GetById:output_type -> users_service.Journal
	5, // 8: users_service.JournalService.GetList:output_type -> users_service.GetListJournalResponse
	1, // 9: users_service.JournalService.Update:output_type -> users_service.Journal
	6, // 10: users_service.JournalService.Delete:output_type -> users_service.Empty
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_journal_service_proto_init() }
func file_journal_service_proto_init() {
	if File_journal_service_proto != nil {
		return
	}
	file_teacher_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_journal_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JournalPrimaryKey); i {
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
		file_journal_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Journal); i {
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
		file_journal_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateJournal); i {
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
		file_journal_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateJournal); i {
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
		file_journal_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListJournalRequest); i {
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
		file_journal_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListJournalResponse); i {
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
			RawDescriptor: file_journal_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_journal_service_proto_goTypes,
		DependencyIndexes: file_journal_service_proto_depIdxs,
		MessageInfos:      file_journal_service_proto_msgTypes,
	}.Build()
	File_journal_service_proto = out.File
	file_journal_service_proto_rawDesc = nil
	file_journal_service_proto_goTypes = nil
	file_journal_service_proto_depIdxs = nil
}
