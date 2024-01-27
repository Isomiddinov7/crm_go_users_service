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

type GroupService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*users_service.UnimplementedGroupServiceServer
}

func NewGroupService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *GroupService {
	return &GroupService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *GroupService) Create(ctx context.Context, req *users_service.CreateGroup) (resp *users_service.Group, err error) {

	i.log.Info("---CreateGroup------>", logger.Any("req", req))
	id, err := i.strg.Group().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateGroup->Group->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err != nil {
		i.log.Error("!!!GetByIdTask->task->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp, err = i.strg.Group().GetById(ctx, id)
	if err != nil {
		i.log.Error("!!!GetByIdGroup->Group->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return
}

func (i *GroupService) GetById(ctx context.Context, req *users_service.GroupPrimaryKey) (resp *users_service.Group, err error) {
	i.log.Info("---GetGroupByID------>", logger.Any("req", req))

	resp, err = i.strg.Group().GetById(ctx, req)
	if err != nil {
		i.log.Error("!!!GetGroupByID->Group->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return resp, nil
}

func (i *GroupService) GetList(ctx context.Context, req *users_service.GetListGroupRequest) (resp *users_service.GetListGroupResponse, err error) {

	i.log.Info("---GetGroups------>", logger.Any("req", req))

	resp, err = i.strg.Group().GetList(ctx, req)
	if err != nil {
		i.log.Error("!!!GetGroups->Group->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *GroupService) Update(ctx context.Context, req *users_service.UpdateGroup) (resp *users_service.Group, err error) {

	i.log.Info("---UpdateGroup------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Group().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateGroup--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Group().GetById(ctx, &users_service.GroupPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!GetTeacher->Teacher->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *GroupService) Delete(ctx context.Context, req *users_service.GroupPrimaryKey) (resp *users_service.Empty, err error) {

	i.log.Info("---DeleteGroup------>", logger.Any("req", req))

	err = i.strg.Group().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteGroup->Group->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &users_service.Empty{}, nil
}
