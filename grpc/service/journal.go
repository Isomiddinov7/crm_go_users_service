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

type JournalService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*users_service.UnimplementedJournalServiceServer
}

func NewJournalService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *JournalService {
	return &JournalService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *JournalService) Create(ctx context.Context, req *users_service.CreateJournal) (resp *users_service.Journal, err error) {

	i.log.Info("---CreateJournal------>", logger.Any("req", req))
	id, err := i.strg.Journal().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateJournal->Journal->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err != nil {
		i.log.Error("!!!GetByIdTask->task->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp, err = i.strg.Journal().GetById(ctx, id)
	if err != nil {
		i.log.Error("!!!GetByIdJournal->Journal->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return
}

func (i *JournalService) GetById(ctx context.Context, req *users_service.JournalPrimaryKey) (resp *users_service.Journal, err error) {
	i.log.Info("---GetJournalByID------>", logger.Any("req", req))

	resp, err = i.strg.Journal().GetById(ctx, req)
	if err != nil {
		i.log.Error("!!!GetJournalByID->Journal->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return resp, nil
}

func (i *JournalService) GetList(ctx context.Context, req *users_service.GetListJournalRequest) (resp *users_service.GetListJournalResponse, err error) {

	i.log.Info("---GetJournals------>", logger.Any("req", req))

	resp, err = i.strg.Journal().GetList(ctx, req)
	if err != nil {
		i.log.Error("!!!GetJournals->Journal->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *JournalService) Update(ctx context.Context, req *users_service.UpdateJournal) (resp *users_service.Journal, err error) {

	i.log.Info("---UpdateJournal------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Journal().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateJournal--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Journal().GetById(ctx, &users_service.JournalPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!GetTeacher->Teacher->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *JournalService) Delete(ctx context.Context, req *users_service.JournalPrimaryKey) (resp *users_service.Empty, err error) {

	i.log.Info("---DeleteJournal------>", logger.Any("req", req))

	err = i.strg.Journal().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteJournal->Journal->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &users_service.Empty{}, nil
}
