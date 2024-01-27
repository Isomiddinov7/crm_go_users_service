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

type SupportTeacherService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*users_service.UnimplementedSupportTeacherServiceServer
}

func NewSupportTeacherService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *SupportTeacherService {
	return &SupportTeacherService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *SupportTeacherService) Create(ctx context.Context, req *users_service.CreateSupportTeacher) (resp *users_service.SupportTeacher, err error) {

	i.log.Info("---CreateSupportTeacher------>", logger.Any("req", req))

	id, err := i.strg.SupportTeacher().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateSupportTeacher->SupportTeacher->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp, err = i.strg.SupportTeacher().GetById(ctx, id)
	if err != nil {
		i.log.Error("!!!GetByIdSupportTeacher->Teacher->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	fmt.Println(resp)
	return
}

func (i *SupportTeacherService) GetById(ctx context.Context, req *users_service.SupportTeacherPrimaryKey) (resp *users_service.SupportTeacher, err error) {
	i.log.Info("---GetSupportTeacherByID------>", logger.Any("req", req))

	resp, err = i.strg.SupportTeacher().GetById(ctx, req)
	if err != nil {
		i.log.Error("!!!GetSupportTeacherByID->SupportTeacher->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return resp, nil
}

func (i *SupportTeacherService) GetList(ctx context.Context, req *users_service.GetListSupportTeacherRequest) (resp *users_service.GetListSupportTeacherResponse, err error) {

	i.log.Info("---GetSupportTeachers------>", logger.Any("req", req))

	resp, err = i.strg.SupportTeacher().GetList(ctx, req)
	if err != nil {
		i.log.Error("!!!GetSupportTeachers->SupportTeacher->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *SupportTeacherService) Update(ctx context.Context, req *users_service.UpdateSupportTeacher) (resp *users_service.SupportTeacher, err error) {

	i.log.Info("---UpdateSupportTeacher------>", logger.Any("req", req))

	rowsAffected, err := i.strg.SupportTeacher().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateSupportTeacher--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.SupportTeacher().GetById(ctx, &users_service.SupportTeacherPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!GetTeacher->Teacher->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *SupportTeacherService) Delete(ctx context.Context, req *users_service.SupportTeacherPrimaryKey) (resp *users_service.Empty, err error) {

	i.log.Info("---DeleteSupportTeacher------>", logger.Any("req", req))

	err = i.strg.SupportTeacher().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteSupportTeacher->SupportTeacher->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &users_service.Empty{}, nil
}
