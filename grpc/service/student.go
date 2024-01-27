package service

import (
	"context"
	"crm_go_users_service/config"
	"crm_go_users_service/genproto/users_service"
	"crm_go_users_service/grpc/client"
	"crm_go_users_service/pkg/logger"
	"crm_go_users_service/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StudentService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*users_service.UnimplementedStudentServiceServer
}

func NewStudentService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *StudentService {
	return &StudentService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *StudentService) Create(ctx context.Context, req *users_service.CreateStudent) (resp *users_service.Student, err error) {

	i.log.Info("---CreateStudent------>", logger.Any("req", req))

	id, err := i.strg.Student().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateStudent->Student->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp, err = i.strg.Student().GetById(ctx, id)
	if err != nil {
		i.log.Error("!!!GetByIdStudent->Teacher->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return
}

func (i *StudentService) GetById(ctx context.Context, req *users_service.StudentPrimaryKey) (resp *users_service.Student, err error) {
	i.log.Info("---GetStudentByID------>", logger.Any("req", req))

	resp, err = i.strg.Student().GetById(ctx, req)
	if err != nil {
		i.log.Error("!!!GetStudentByID->Student->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return resp, nil
}

func (i *StudentService) GetList(ctx context.Context, req *users_service.GetListStudentRequest) (resp *users_service.GetListStudentResponse, err error) {

	i.log.Info("---GetStudents------>", logger.Any("req", req))

	resp, err = i.strg.Student().GetList(ctx, req)
	if err != nil {
		i.log.Error("!!!GetStudents->Student->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *StudentService) Update(ctx context.Context, req *users_service.UpdateStudent) (resp *users_service.Student, err error) {

	i.log.Info("---UpdateStudent------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Student().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateStudent--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Student().GetById(ctx, &users_service.StudentPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!GetTeacher->Teacher->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *StudentService) Delete(ctx context.Context, req *users_service.StudentPrimaryKey) (resp *users_service.Empty, err error) {

	i.log.Info("---DeleteStudent------>", logger.Any("req", req))

	err = i.strg.Student().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteStudent->Student->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &users_service.Empty{}, nil
}
