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

type ScoreService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*users_service.UnimplementedScoreServiceServer
}

func NewScoreService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *ScoreService {
	return &ScoreService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *ScoreService) Create(ctx context.Context, req *users_service.CreateScore) (resp *users_service.Score, err error) {

	i.log.Info("---CreateScore------>", logger.Any("req", req))
	id, err := i.strg.Score().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateScore->Score->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err != nil {
		i.log.Error("!!!GetByIdTask->task->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp, err = i.strg.Score().GetById(ctx, id)
	if err != nil {
		i.log.Error("!!!GetByIdScore->Score->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return
}

func (i *ScoreService) GetById(ctx context.Context, req *users_service.ScorePrimaryKey) (resp *users_service.Score, err error) {
	i.log.Info("---GetScoreByID------>", logger.Any("req", req))

	resp, err = i.strg.Score().GetById(ctx, req)
	if err != nil {
		i.log.Error("!!!GetScoreByID->Score->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return resp, nil
}

func (i *ScoreService) GetList(ctx context.Context, req *users_service.GetListScoreRequest) (resp *users_service.GetListScoreResponse, err error) {

	i.log.Info("---GetScores------>", logger.Any("req", req))

	resp, err = i.strg.Score().GetList(ctx, req)
	if err != nil {
		i.log.Error("!!!GetScores->Score->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *ScoreService) Update(ctx context.Context, req *users_service.UpdateScore) (resp *users_service.Score, err error) {

	i.log.Info("---UpdateScore------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Score().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateScore--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Score().GetById(ctx, &users_service.ScorePrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!GetTeacher->Teacher->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *ScoreService) Delete(ctx context.Context, req *users_service.ScorePrimaryKey) (resp *users_service.Empty, err error) {

	i.log.Info("---DeleteScore------>", logger.Any("req", req))

	err = i.strg.Score().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteScore->Score->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &users_service.Empty{}, nil
}
