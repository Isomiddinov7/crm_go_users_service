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

type SupperAdminService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*users_service.UnimplementedSupperAdminServiceServer
}

func NewSupperAdminService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *SupperAdminService {
	return &SupperAdminService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *SupperAdminService) Create(ctx context.Context, req *users_service.CreateSupperAdmin) (resp *users_service.SupperAdmin, err error) {

	i.log.Info("---CreateSupperAdmin------>", logger.Any("req", req))

	id, err := i.strg.SuperAdmin().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateSupperAdmin->SupperAdmin->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp, err = i.strg.SuperAdmin().GetById(ctx, id)
	if err != nil {
		i.log.Error("!!!GetByIdSupperAdmin->Teacher->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	fmt.Println(resp)
	return
}

func (i *SupperAdminService) GetById(ctx context.Context, req *users_service.SupperAdminPrimaryKey) (resp *users_service.SupperAdmin, err error) {
	i.log.Info("---GetSupperAdminByID------>", logger.Any("req", req))

	resp, err = i.strg.SuperAdmin().GetById(ctx, req)
	if err != nil {
		i.log.Error("!!!GetSupperAdminByID->SupperAdmin->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return resp, nil
}

func (i *SupperAdminService) GetList(ctx context.Context, req *users_service.GetListSupperAdminRequest) (resp *users_service.GetListSupperAdminResponse, err error) {

	i.log.Info("---GetSupperAdmins------>", logger.Any("req", req))

	resp, err = i.strg.SuperAdmin().GetList(ctx, req)
	if err != nil {
		i.log.Error("!!!GetSupperAdmins->SupperAdmin->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *SupperAdminService) Update(ctx context.Context, req *users_service.UpdateSupperAdmin) (resp *users_service.SupperAdmin, err error) {

	i.log.Info("---UpdateSupperAdmin------>", logger.Any("req", req))

	rowsAffected, err := i.strg.SuperAdmin().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateSupperAdmin--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.SuperAdmin().GetById(ctx, &users_service.SupperAdminPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!GetSuperAdmin->SuperAdmin->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *SupperAdminService) Delete(ctx context.Context, req *users_service.SupperAdminPrimaryKey) (resp *users_service.Empty, err error) {

	i.log.Info("---DeleteSupperAdmin------>", logger.Any("req", req))

	err = i.strg.SuperAdmin().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteSupperAdmin->SupperAdmin->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &users_service.Empty{}, nil
}
