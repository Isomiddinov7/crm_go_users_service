// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: group_service.proto

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

type GroupPrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GroupPrimaryKey) Reset() {
	*x = GroupPrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupPrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupPrimaryKey) ProtoMessage() {}

func (x *GroupPrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_group_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupPrimaryKey.ProtoReflect.Descriptor instead.
func (*GroupPrimaryKey) Descriptor() ([]byte, []int) {
	return file_group_service_proto_rawDescGZIP(), []int{0}
}

func (x *GroupPrimaryKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UniqueID         string `protobuf:"bytes,2,opt,name=uniqueID,proto3" json:"uniqueID,omitempty"`
	BranchId         string `protobuf:"bytes,3,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	Type             string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	TeacherId        string `protobuf:"bytes,5,opt,name=teacher_id,json=teacherId,proto3" json:"teacher_id,omitempty"`
	SupportTeacherId string `protobuf:"bytes,6,opt,name=support_teacher_id,json=supportTeacherId,proto3" json:"support_teacher_id,omitempty"`
	CreatedAt        string `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt        string `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_group_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_group_service_proto_rawDescGZIP(), []int{1}
}

func (x *Group) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Group) GetUniqueID() string {
	if x != nil {
		return x.UniqueID
	}
	return ""
}

func (x *Group) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *Group) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Group) GetTeacherId() string {
	if x != nil {
		return x.TeacherId
	}
	return ""
}

func (x *Group) GetSupportTeacherId() string {
	if x != nil {
		return x.SupportTeacherId
	}
	return ""
}

func (x *Group) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Group) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type CreateGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UniqueID         string `protobuf:"bytes,1,opt,name=uniqueID,proto3" json:"uniqueID,omitempty"`
	BranchId         string `protobuf:"bytes,2,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	Type             string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	TeacherId        string `protobuf:"bytes,4,opt,name=teacher_id,json=teacherId,proto3" json:"teacher_id,omitempty"`
	SupportTeacherId string `protobuf:"bytes,5,opt,name=support_teacher_id,json=supportTeacherId,proto3" json:"support_teacher_id,omitempty"`
}

func (x *CreateGroup) Reset() {
	*x = CreateGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroup) ProtoMessage() {}

func (x *CreateGroup) ProtoReflect() protoreflect.Message {
	mi := &file_group_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroup.ProtoReflect.Descriptor instead.
func (*CreateGroup) Descriptor() ([]byte, []int) {
	return file_group_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateGroup) GetUniqueID() string {
	if x != nil {
		return x.UniqueID
	}
	return ""
}

func (x *CreateGroup) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *CreateGroup) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreateGroup) GetTeacherId() string {
	if x != nil {
		return x.TeacherId
	}
	return ""
}

func (x *CreateGroup) GetSupportTeacherId() string {
	if x != nil {
		return x.SupportTeacherId
	}
	return ""
}

type UpdateGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UniqueID         string `protobuf:"bytes,2,opt,name=uniqueID,proto3" json:"uniqueID,omitempty"`
	BranchId         string `protobuf:"bytes,3,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	Type             string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	TeacherId        string `protobuf:"bytes,5,opt,name=teacher_id,json=teacherId,proto3" json:"teacher_id,omitempty"`
	SupportTeacherId string `protobuf:"bytes,6,opt,name=support_teacher_id,json=supportTeacherId,proto3" json:"support_teacher_id,omitempty"`
}

func (x *UpdateGroup) Reset() {
	*x = UpdateGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGroup) ProtoMessage() {}

func (x *UpdateGroup) ProtoReflect() protoreflect.Message {
	mi := &file_group_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGroup.ProtoReflect.Descriptor instead.
func (*UpdateGroup) Descriptor() ([]byte, []int) {
	return file_group_service_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateGroup) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateGroup) GetUniqueID() string {
	if x != nil {
		return x.UniqueID
	}
	return ""
}

func (x *UpdateGroup) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *UpdateGroup) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *UpdateGroup) GetTeacherId() string {
	if x != nil {
		return x.TeacherId
	}
	return ""
}

func (x *UpdateGroup) GetSupportTeacherId() string {
	if x != nil {
		return x.SupportTeacherId
	}
	return ""
}

type GetListGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int64  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *GetListGroupRequest) Reset() {
	*x = GetListGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListGroupRequest) ProtoMessage() {}

func (x *GetListGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_group_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListGroupRequest.ProtoReflect.Descriptor instead.
func (*GetListGroupRequest) Descriptor() ([]byte, []int) {
	return file_group_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetListGroupRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetListGroupRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetListGroupRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type GetListGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count  int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Groups []*Group `protobuf:"bytes,2,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (x *GetListGroupResponse) Reset() {
	*x = GetListGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListGroupResponse) ProtoMessage() {}

func (x *GetListGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_group_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListGroupResponse.ProtoReflect.Descriptor instead.
func (*GetListGroupResponse) Descriptor() ([]byte, []int) {
	return file_group_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetListGroupResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetListGroupResponse) GetGroups() []*Group {
	if x != nil {
		return x.Groups
	}
	return nil
}

var File_group_service_proto protoreflect.FileDescriptor

var file_group_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x1a, 0x15, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a, 0x0f, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xef,
	0x01, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x6e, 0x69, 0x71,
	0x75, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x6e, 0x69, 0x71,
	0x75, 0x65, 0x49, 0x44, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x5f,
	0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0xa7, 0x01, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x44, 0x12, 0x1b, 0x0a, 0x09,
	0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12,
	0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x64, 0x22, 0xb7, 0x01, 0x0a, 0x0b, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x6e,
	0x69, 0x71, 0x75, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x6e,
	0x69, 0x71, 0x75, 0x65, 0x49, 0x44, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x5f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x5b, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x22, 0x5a, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x2c, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x32, 0xe5, 0x02,
	0x0a, 0x0c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x1a, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x00, 0x12,
	0x54, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x1a, 0x14, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1e, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x14, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x18, 0x5a, 0x16, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_group_service_proto_rawDescOnce sync.Once
	file_group_service_proto_rawDescData = file_group_service_proto_rawDesc
)

func file_group_service_proto_rawDescGZIP() []byte {
	file_group_service_proto_rawDescOnce.Do(func() {
		file_group_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_group_service_proto_rawDescData)
	})
	return file_group_service_proto_rawDescData
}

var file_group_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_group_service_proto_goTypes = []interface{}{
	(*GroupPrimaryKey)(nil),      // 0: users_service.GroupPrimaryKey
	(*Group)(nil),                // 1: users_service.Group
	(*CreateGroup)(nil),          // 2: users_service.CreateGroup
	(*UpdateGroup)(nil),          // 3: users_service.UpdateGroup
	(*GetListGroupRequest)(nil),  // 4: users_service.GetListGroupRequest
	(*GetListGroupResponse)(nil), // 5: users_service.GetListGroupResponse
	(*Empty)(nil),                // 6: users_service.Empty
}
var file_group_service_proto_depIdxs = []int32{
	1, // 0: users_service.GetListGroupResponse.groups:type_name -> users_service.Group
	2, // 1: users_service.GroupService.Create:input_type -> users_service.CreateGroup
	0, // 2: users_service.GroupService.GetById:input_type -> users_service.GroupPrimaryKey
	4, // 3: users_service.GroupService.GetList:input_type -> users_service.GetListGroupRequest
	3, // 4: users_service.GroupService.Update:input_type -> users_service.UpdateGroup
	0, // 5: users_service.GroupService.Delete:input_type -> users_service.GroupPrimaryKey
	1, // 6: users_service.GroupService.Create:output_type -> users_service.Group
	1, // 7: users_service.GroupService.GetById:output_type -> users_service.Group
	5, // 8: users_service.GroupService.GetList:output_type -> users_service.GetListGroupResponse
	1, // 9: users_service.GroupService.Update:output_type -> users_service.Group
	6, // 10: users_service.GroupService.Delete:output_type -> users_service.Empty
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_group_service_proto_init() }
func file_group_service_proto_init() {
	if File_group_service_proto != nil {
		return
	}
	file_teacher_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_group_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupPrimaryKey); i {
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
		file_group_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
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
		file_group_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroup); i {
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
		file_group_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGroup); i {
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
		file_group_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListGroupRequest); i {
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
		file_group_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListGroupResponse); i {
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
			RawDescriptor: file_group_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_group_service_proto_goTypes,
		DependencyIndexes: file_group_service_proto_depIdxs,
		MessageInfos:      file_group_service_proto_msgTypes,
	}.Build()
	File_group_service_proto = out.File
	file_group_service_proto_rawDesc = nil
	file_group_service_proto_goTypes = nil
	file_group_service_proto_depIdxs = nil
}
