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

type BranchService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*users_service.UnimplementedBranchServiceServer
}

func NewBranchService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *BranchService {
	return &BranchService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *BranchService) Create(ctx context.Context, req *users_service.CreateBranch) (resp *users_service.Branch, err error) {

	i.log.Info("---CreateBranch------>", logger.Any("req", req))
	id, err := i.strg.Branch().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateBranch->Branch->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp, err = i.strg.Branch().GetById(ctx, id)
	if err != nil {
		i.log.Error("!!!GetByIdBranch->Branch->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return
}

func (i *BranchService) GetById(ctx context.Context, req *users_service.BranchPrimaryKey) (resp *users_service.Branch, err error) {
	i.log.Info("---GetBranchByID------>", logger.Any("req", req))

	resp, err = i.strg.Branch().GetById(ctx, req)
	if err != nil {
		i.log.Error("!!!GetBranchByID->Branch->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return resp, nil
}

func (i *BranchService) GetList(ctx context.Context, req *users_service.GetListBranchRequest) (resp *users_service.GetListBranchResponse, err error) {

	i.log.Info("---GetBranchs------>", logger.Any("req", req))

	resp, err = i.strg.Branch().GetList(ctx, req)
	if err != nil {
		i.log.Error("!!!GetBranchs->Branch->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *BranchService) Update(ctx context.Context, req *users_service.UpdateBranch) (resp *users_service.Branch, err error) {

	i.log.Info("---UpdateBranch------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Branch().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateBranch--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Branch().GetById(ctx, &users_service.BranchPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!GetTeacher->Teacher->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *BranchService) Delete(ctx context.Context, req *users_service.BranchPrimaryKey) (resp *users_service.Empty, err error) {

	i.log.Info("---DeleteBranch------>", logger.Any("req", req))

	err = i.strg.Branch().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteBranch->Branch->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &users_service.Empty{}, nil
}
