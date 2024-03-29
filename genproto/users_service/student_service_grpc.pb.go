// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: student_service.proto

package users_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	StudentService_Create_FullMethodName  = "/users_service.StudentService/Create"
	StudentService_GetById_FullMethodName = "/users_service.StudentService/GetById"
	StudentService_GetList_FullMethodName = "/users_service.StudentService/GetList"
	StudentService_Update_FullMethodName  = "/users_service.StudentService/Update"
	StudentService_Delete_FullMethodName  = "/users_service.StudentService/Delete"
)

// StudentServiceClient is the client API for StudentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StudentServiceClient interface {
	Create(ctx context.Context, in *CreateStudent, opts ...grpc.CallOption) (*Student, error)
	GetById(ctx context.Context, in *StudentPrimaryKey, opts ...grpc.CallOption) (*Student, error)
	GetList(ctx context.Context, in *GetListStudentRequest, opts ...grpc.CallOption) (*GetListStudentResponse, error)
	Update(ctx context.Context, in *UpdateStudent, opts ...grpc.CallOption) (*Student, error)
	Delete(ctx context.Context, in *StudentPrimaryKey, opts ...grpc.CallOption) (*Empty, error)
}

type studentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStudentServiceClient(cc grpc.ClientConnInterface) StudentServiceClient {
	return &studentServiceClient{cc}
}

func (c *studentServiceClient) Create(ctx context.Context, in *CreateStudent, opts ...grpc.CallOption) (*Student, error) {
	out := new(Student)
	err := c.cc.Invoke(ctx, StudentService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) GetById(ctx context.Context, in *StudentPrimaryKey, opts ...grpc.CallOption) (*Student, error) {
	out := new(Student)
	err := c.cc.Invoke(ctx, StudentService_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) GetList(ctx context.Context, in *GetListStudentRequest, opts ...grpc.CallOption) (*GetListStudentResponse, error) {
	out := new(GetListStudentResponse)
	err := c.cc.Invoke(ctx, StudentService_GetList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) Update(ctx context.Context, in *UpdateStudent, opts ...grpc.CallOption) (*Student, error) {
	out := new(Student)
	err := c.cc.Invoke(ctx, StudentService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) Delete(ctx context.Context, in *StudentPrimaryKey, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, StudentService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudentServiceServer is the server API for StudentService service.
// All implementations must embed UnimplementedStudentServiceServer
// for forward compatibility
type StudentServiceServer interface {
	Create(context.Context, *CreateStudent) (*Student, error)
	GetById(context.Context, *StudentPrimaryKey) (*Student, error)
	GetList(context.Context, *GetListStudentRequest) (*GetListStudentResponse, error)
	Update(context.Context, *UpdateStudent) (*Student, error)
	Delete(context.Context, *StudentPrimaryKey) (*Empty, error)
	mustEmbedUnimplementedStudentServiceServer()
}

// UnimplementedStudentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStudentServiceServer struct {
}

func (UnimplementedStudentServiceServer) Create(context.Context, *CreateStudent) (*Student, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedStudentServiceServer) GetById(context.Context, *StudentPrimaryKey) (*Student, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedStudentServiceServer) GetList(context.Context, *GetListStudentRequest) (*GetListStudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedStudentServiceServer) Update(context.Context, *UpdateStudent) (*Student, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedStudentServiceServer) Delete(context.Context, *StudentPrimaryKey) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedStudentServiceServer) mustEmbedUnimplementedStudentServiceServer() {}

// UnsafeStudentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StudentServiceServer will
// result in compilation errors.
type UnsafeStudentServiceServer interface {
	mustEmbedUnimplementedStudentServiceServer()
}

func RegisterStudentServiceServer(s grpc.ServiceRegistrar, srv StudentServiceServer) {
	s.RegisterService(&StudentService_ServiceDesc, srv)
}

func _StudentService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStudent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).Create(ctx, req.(*CreateStudent))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StudentPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).GetById(ctx, req.(*StudentPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_GetList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).GetList(ctx, req.(*GetListStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStudent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).Update(ctx, req.(*UpdateStudent))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StudentPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).Delete(ctx, req.(*StudentPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

// StudentService_ServiceDesc is the grpc.ServiceDesc for StudentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StudentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "users_service.StudentService",
	HandlerType: (*StudentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _StudentService_Create_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _StudentService_GetById_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _StudentService_GetList_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _StudentService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _StudentService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "student_service.proto",
}
