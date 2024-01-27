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

type ManagerService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*users_service.UnimplementedManagerServiceServer
}

func NewManagerService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *ManagerService {
	return &ManagerService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *ManagerService) Create(ctx context.Context, req *users_service.CreateManager) (resp *users_service.Manager, err error) {

	i.log.Info("---CreateManager------>", logger.Any("req", req))
	id, err := i.strg.Manager().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateManager->Manager->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err != nil {
		i.log.Error("!!!GetByIdTask->task->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp, err = i.strg.Manager().GetById(ctx, id)
	if err != nil {
		i.log.Error("!!!GetByIdManager->Manager->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return
}

func (i *ManagerService) GetById(ctx context.Context, req *users_service.ManagerPrimaryKey) (resp *users_service.Manager, err error) {
	i.log.Info("---GetManagerByID------>", logger.Any("req", req))

	resp, err = i.strg.Manager().GetById(ctx, req)
	if err != nil {
		i.log.Error("!!!GetManagerByID->Manager->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return resp, nil
}

func (i *ManagerService) GetList(ctx context.Context, req *users_service.GetListManagerRequest) (resp *users_service.GetListManagerResponse, err error) {

	i.log.Info("---GetManagers------>", logger.Any("req", req))

	resp, err = i.strg.Manager().GetList(ctx, req)
	if err != nil {
		i.log.Error("!!!GetManagers->Manager->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *ManagerService) Update(ctx context.Context, req *users_service.UpdateManager) (resp *users_service.Manager, err error) {

	i.log.Info("---UpdateManager------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Manager().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateManager--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Manager().GetById(ctx, &users_service.ManagerPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!GetTeacher->Teacher->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *ManagerService) Delete(ctx context.Context, req *users_service.ManagerPrimaryKey) (resp *users_service.Empty, err error) {

	i.log.Info("---DeleteManager------>", logger.Any("req", req))

	err = i.strg.Manager().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteManager->Manager->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &users_service.Empty{}, nil
}
