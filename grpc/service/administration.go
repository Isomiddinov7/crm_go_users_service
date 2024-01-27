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

type AdministrationService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*users_service.UnimplementedAdministrationServiceServer
}

func NewAdministrationService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *AdministrationService {
	return &AdministrationService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *AdministrationService) Create(ctx context.Context, req *users_service.CreateAdministration) (resp *users_service.Administration, err error) {

	i.log.Info("---CreateAdministration------>", logger.Any("req", req))

	id, err := i.strg.Administrator().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateAdministration->Administration->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp, err = i.strg.Administrator().GetById(ctx, id)
	if err != nil {
		i.log.Error("!!!GetByIdAdministration->Teacher->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	fmt.Println(resp)
	return
}

func (i *AdministrationService) GetById(ctx context.Context, req *users_service.AdministrationPrimaryKey) (resp *users_service.Administration, err error) {
	i.log.Info("---GetAdministrationByID------>", logger.Any("req", req))

	resp, err = i.strg.Administrator().GetById(ctx, req)
	if err != nil {
		i.log.Error("!!!GetAdministrationByID->Administration->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return resp, nil
}

func (i *AdministrationService) GetList(ctx context.Context, req *users_service.GetListAdministrationRequest) (resp *users_service.GetListAdministrationResponse, err error) {

	i.log.Info("---GetAdministrations------>", logger.Any("req", req))

	resp, err = i.strg.Administrator().GetList(ctx, req)
	if err != nil {
		i.log.Error("!!!GetAdministrations->Administration->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *AdministrationService) Update(ctx context.Context, req *users_service.UpdateAdministration) (resp *users_service.Administration, err error) {

	i.log.Info("---UpdateAdministration------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Administrator().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateAdministration--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Administrator().GetById(ctx, &users_service.AdministrationPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!GetTeacher->Teacher->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *AdministrationService) Delete(ctx context.Context, req *users_service.AdministrationPrimaryKey) (resp *users_service.Empty, err error) {

	i.log.Info("---DeleteAdministration------>", logger.Any("req", req))

	err = i.strg.Administrator().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteAdministration->Administration->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &users_service.Empty{}, nil
}
