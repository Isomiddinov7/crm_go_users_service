package service

import (
	"context"
	"crm_go_users_service/config"
	"crm_go_users_service/genproto/users_service"
	"crm_go_users_service/grpc/client"
	"crm_go_users_service/pkg/logger"
	"crm_go_users_service/storage"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TeacherService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*users_service.UnimplementedTeacherServiceServer
}

func NewTeacherService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *TeacherService {
	return &TeacherService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *TeacherService) Create(ctx context.Context, req *users_service.CreateTeacher) (resp *users_service.Teacher, err error) {

	i.log.Info("---CreateTeacher------>", logger.Any("req", req))

	id, err := i.strg.Teacher().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateTeacher->Teacher->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp, err = i.strg.Teacher().GetById(ctx, id)
	if err != nil {
		i.log.Error("!!!GetByIdTeacher->Teacher->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	fmt.Println(resp)
	return
}

func (i *TeacherService) GetById(ctx context.Context, req *users_service.TeacherPrimaryKey) (resp *users_service.Teacher, err error) {
	i.log.Info("---GetTeacherByID------>", logger.Any("req", req))

	resp, err = i.strg.Teacher().GetById(ctx, req)
	if err != nil {
		i.log.Error("!!!GetTeacherByID->Teacher->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return resp, nil
}

func (i *TeacherService) GetList(ctx context.Context, req *users_service.GetListTeacherRequest) (resp *users_service.GetListTeacherResponse, err error) {

	i.log.Info("---GetTeachers------>", logger.Any("req", req))

	resp, err = i.strg.Teacher().GetList(ctx, req)
	if err != nil {
		i.log.Error("!!!GetTeachers->Teacher->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *TeacherService) Update(ctx context.Context, req *users_service.UpdateTeacher) (resp *users_service.Teacher, err error) {

	i.log.Info("---UpdateTeacher------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Teacher().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateTeacher--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Teacher().GetById(ctx, &users_service.TeacherPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!GetTeacher->Teacher->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *TeacherService) Delete(ctx context.Context, req *users_service.TeacherPrimaryKey) (resp *users_service.Empty, err error) {

	i.log.Info("---DeleteTeacher------>", logger.Any("req", req))

	err = i.strg.Teacher().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteTeacher->Teacher->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &users_service.Empty{}, nil
}
