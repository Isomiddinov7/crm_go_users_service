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

type RoleService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*users_service.UnimplementedRoleServiceServer
}

func NewRoleService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *RoleService {
	return &RoleService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *RoleService) Create(ctx context.Context, req *users_service.CreateRole) (resp *users_service.Role, err error) {

	i.log.Info("---CreateRole------>", logger.Any("req", req))
	id, err := i.strg.Role().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateRole->Role->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp, err = i.strg.Role().GetById(ctx, id)
	if err != nil {
		i.log.Error("!!!GetByIdRole->Role->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return
}

func (i *RoleService) GetById(ctx context.Context, req *users_service.RolePrimaryKey) (resp *users_service.Role, err error) {
	i.log.Info("---GetRoleByID------>", logger.Any("req", req))

	resp, err = i.strg.Role().GetById(ctx, req)
	if err != nil {
		i.log.Error("!!!GetRoleByID->Role->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return resp, nil
}

func (i *RoleService) GetList(ctx context.Context, req *users_service.GetListRoleRequest) (resp *users_service.GetListRoleResponse, err error) {

	i.log.Info("---GetRoles------>", logger.Any("req", req))

	resp, err = i.strg.Role().GetList(ctx, req)
	if err != nil {
		i.log.Error("!!!GetRoles->Role->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *RoleService) Update(ctx context.Context, req *users_service.UpdateRole) (resp *users_service.Role, err error) {

	i.log.Info("---UpdateRole------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Role().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateRole--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Role().GetById(ctx, &users_service.RolePrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!GetTeacher->Teacher->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *RoleService) Delete(ctx context.Context, req *users_service.RolePrimaryKey) (resp *users_service.Empty, err error) {

	i.log.Info("---DeleteRole------>", logger.Any("req", req))

	err = i.strg.Role().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteRole->Role->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &users_service.Empty{}, nil
}
