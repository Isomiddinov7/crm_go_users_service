// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: do_task.proto

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

type DoTaskPrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DoTaskPrimaryKey) Reset() {
	*x = DoTaskPrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_do_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoTaskPrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoTaskPrimaryKey) ProtoMessage() {}

func (x *DoTaskPrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_do_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoTaskPrimaryKey.ProtoReflect.Descriptor instead.
func (*DoTaskPrimaryKey) Descriptor() ([]byte, []int) {
	return file_do_task_proto_rawDescGZIP(), []int{0}
}

func (x *DoTaskPrimaryKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateDoTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentId string  `protobuf:"bytes,1,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	TaskId    string  `protobuf:"bytes,2,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	LessonId  string  `protobuf:"bytes,3,opt,name=lesson_id,json=lessonId,proto3" json:"lesson_id,omitempty"`
	Score     float64 `protobuf:"fixed64,4,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *CreateDoTask) Reset() {
	*x = CreateDoTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_do_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDoTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDoTask) ProtoMessage() {}

func (x *CreateDoTask) ProtoReflect() protoreflect.Message {
	mi := &file_do_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDoTask.ProtoReflect.Descriptor instead.
func (*CreateDoTask) Descriptor() ([]byte, []int) {
	return file_do_task_proto_rawDescGZIP(), []int{1}
}

func (x *CreateDoTask) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

func (x *CreateDoTask) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *CreateDoTask) GetLessonId() string {
	if x != nil {
		return x.LessonId
	}
	return ""
}

func (x *CreateDoTask) GetScore() float64 {
	if x != nil {
		return x.Score
	}
	return 0
}

type DoTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	StudentId string  `protobuf:"bytes,2,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	TaskId    string  `protobuf:"bytes,3,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	LessonId  string  `protobuf:"bytes,4,opt,name=lesson_id,json=lessonId,proto3" json:"lesson_id,omitempty"`
	CreatedAt string  `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string  `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Score     float64 `protobuf:"fixed64,7,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *DoTask) Reset() {
	*x = DoTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_do_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoTask) ProtoMessage() {}

func (x *DoTask) ProtoReflect() protoreflect.Message {
	mi := &file_do_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoTask.ProtoReflect.Descriptor instead.
func (*DoTask) Descriptor() ([]byte, []int) {
	return file_do_task_proto_rawDescGZIP(), []int{2}
}

func (x *DoTask) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DoTask) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

func (x *DoTask) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *DoTask) GetLessonId() string {
	if x != nil {
		return x.LessonId
	}
	return ""
}

func (x *DoTask) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *DoTask) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *DoTask) GetScore() float64 {
	if x != nil {
		return x.Score
	}
	return 0
}

type UpdateDoTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	StudentId string  `protobuf:"bytes,2,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	TaskId    string  `protobuf:"bytes,3,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	LessonId  string  `protobuf:"bytes,4,opt,name=lesson_id,json=lessonId,proto3" json:"lesson_id,omitempty"`
	Score     float64 `protobuf:"fixed64,5,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *UpdateDoTask) Reset() {
	*x = UpdateDoTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_do_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDoTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDoTask) ProtoMessage() {}

func (x *UpdateDoTask) ProtoReflect() protoreflect.Message {
	mi := &file_do_task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDoTask.ProtoReflect.Descriptor instead.
func (*UpdateDoTask) Descriptor() ([]byte, []int) {
	return file_do_task_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateDoTask) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateDoTask) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

func (x *UpdateDoTask) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *UpdateDoTask) GetLessonId() string {
	if x != nil {
		return x.LessonId
	}
	return ""
}

func (x *UpdateDoTask) GetScore() float64 {
	if x != nil {
		return x.Score
	}
	return 0
}

type GetListDoTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int64  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *GetListDoTaskRequest) Reset() {
	*x = GetListDoTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_do_task_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListDoTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListDoTaskRequest) ProtoMessage() {}

func (x *GetListDoTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_do_task_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListDoTaskRequest.ProtoReflect.Descriptor instead.
func (*GetListDoTaskRequest) Descriptor() ([]byte, []int) {
	return file_do_task_proto_rawDescGZIP(), []int{4}
}

func (x *GetListDoTaskRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetListDoTaskRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetListDoTaskRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type GetListDoTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count   int64     `protobuf:"varint,1,opt,name=Count,proto3" json:"Count,omitempty"`
	DoTasks []*DoTask `protobuf:"bytes,2,rep,name=DoTasks,proto3" json:"DoTasks,omitempty"`
}

func (x *GetListDoTaskResponse) Reset() {
	*x = GetListDoTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_do_task_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListDoTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListDoTaskResponse) ProtoMessage() {}

func (x *GetListDoTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_do_task_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListDoTaskResponse.ProtoReflect.Descriptor instead.
func (*GetListDoTaskResponse) Descriptor() ([]byte, []int) {
	return file_do_task_proto_rawDescGZIP(), []int{5}
}

func (x *GetListDoTaskResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetListDoTaskResponse) GetDoTasks() []*DoTask {
	if x != nil {
		return x.DoTasks
	}
	return nil
}

var File_do_task_proto protoreflect.FileDescriptor

var file_do_task_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x64, 0x6f, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0d, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x15,
	0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x10, 0x44, 0x6f, 0x54, 0x61, 0x73, 0x6b, 0x50,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x79, 0x0a, 0x0c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x44, 0x6f, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x22, 0xc1, 0x01, 0x0a, 0x06, 0x44, 0x6f, 0x54, 0x61, 0x73, 0x6b, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x73, 0x73, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x65, 0x73, 0x73,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x89, 0x01, 0x0a, 0x0c, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x44, 0x6f, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x22, 0x5c, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x44,
	0x6f, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x22, 0x5e, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x6f, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x2f, 0x0a, 0x07, 0x44, 0x6f, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x44, 0x6f, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x07, 0x44, 0x6f, 0x54, 0x61, 0x73,
	0x6b, 0x73, 0x32, 0xf1, 0x02, 0x0a, 0x0f, 0x54, 0x6f, 0x64, 0x6f, 0x54, 0x61, 0x73, 0x6b, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x54, 0x61, 0x73, 0x6b, 0x1a, 0x15, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x6f,
	0x54, 0x61, 0x73, 0x6b, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49,
	0x64, 0x12, 0x1f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x44, 0x6f, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b,
	0x65, 0x79, 0x1a, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x44, 0x6f, 0x54, 0x61, 0x73, 0x6b, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x23, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x6f,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x44, 0x6f, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x54, 0x61, 0x73, 0x6b, 0x1a, 0x15, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x6f, 0x54, 0x61, 0x73,
	0x6b, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1f, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x6f,
	0x54, 0x61, 0x73, 0x6b, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x14,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x18, 0x5a, 0x16, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_do_task_proto_rawDescOnce sync.Once
	file_do_task_proto_rawDescData = file_do_task_proto_rawDesc
)

func file_do_task_proto_rawDescGZIP() []byte {
	file_do_task_proto_rawDescOnce.Do(func() {
		file_do_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_do_task_proto_rawDescData)
	})
	return file_do_task_proto_rawDescData
}

var file_do_task_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_do_task_proto_goTypes = []interface{}{
	(*DoTaskPrimaryKey)(nil),      // 0: users_service.DoTaskPrimaryKey
	(*CreateDoTask)(nil),          // 1: users_service.CreateDoTask
	(*DoTask)(nil),                // 2: users_service.DoTask
	(*UpdateDoTask)(nil),          // 3: users_service.UpdateDoTask
	(*GetListDoTaskRequest)(nil),  // 4: users_service.GetListDoTaskRequest
	(*GetListDoTaskResponse)(nil), // 5: users_service.GetListDoTaskResponse
	(*Empty)(nil),                 // 6: users_service.Empty
}
var file_do_task_proto_depIdxs = []int32{
	2, // 0: users_service.GetListDoTaskResponse.DoTasks:type_name -> users_service.DoTask
	1, // 1: users_service.TodoTaskService.Create:input_type -> users_service.CreateDoTask
	0, // 2: users_service.TodoTaskService.GetById:input_type -> users_service.DoTaskPrimaryKey
	4, // 3: users_service.TodoTaskService.GetList:input_type -> users_service.GetListDoTaskRequest
	3, // 4: users_service.TodoTaskService.Update:input_type -> users_service.UpdateDoTask
	0, // 5: users_service.TodoTaskService.Delete:input_type -> users_service.DoTaskPrimaryKey
	2, // 6: users_service.TodoTaskService.Create:output_type -> users_service.DoTask
	2, // 7: users_service.TodoTaskService.GetById:output_type -> users_service.DoTask
	5, // 8: users_service.TodoTaskService.GetList:output_type -> users_service.GetListDoTaskResponse
	2, // 9: users_service.TodoTaskService.Update:output_type -> users_service.DoTask
	6, // 10: users_service.TodoTaskService.Delete:output_type -> users_service.Empty
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_do_task_proto_init() }
func file_do_task_proto_init() {
	if File_do_task_proto != nil {
		return
	}
	file_teacher_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_do_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoTaskPrimaryKey); i {
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
		file_do_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDoTask); i {
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
		file_do_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoTask); i {
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
		file_do_task_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDoTask); i {
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
		file_do_task_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListDoTaskRequest); i {
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
		file_do_task_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListDoTaskResponse); i {
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
			RawDescriptor: file_do_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_do_task_proto_goTypes,
		DependencyIndexes: file_do_task_proto_depIdxs,
		MessageInfos:      file_do_task_proto_msgTypes,
	}.Build()
	File_do_task_proto = out.File
	file_do_task_proto_rawDesc = nil
	file_do_task_proto_goTypes = nil
	file_do_task_proto_depIdxs = nil
}
