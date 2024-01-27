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

type ScheduleService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*users_service.UnimplementedScheduleServiceServer
}

func NewScheduleService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *ScheduleService {
	return &ScheduleService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *ScheduleService) Create(ctx context.Context, req *users_service.CreateSchedule) (resp *users_service.Schedule, err error) {

	i.log.Info("---CreateSchedule------>", logger.Any("req", req))
	id, err := i.strg.Schedule().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateSchedule->Schedule->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err != nil {
		i.log.Error("!!!GetByIdTask->task->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp, err = i.strg.Schedule().GetById(ctx, id)
	if err != nil {
		i.log.Error("!!!GetByIdSchedule->Schedule->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return
}

func (i *ScheduleService) GetById(ctx context.Context, req *users_service.SchedulePrimaryKey) (resp *users_service.Schedule, err error) {
	i.log.Info("---GetScheduleByID------>", logger.Any("req", req))

	resp, err = i.strg.Schedule().GetById(ctx, req)
	if err != nil {
		i.log.Error("!!!GetScheduleByID->Schedule->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return resp, nil
}

func (i *ScheduleService) GetList(ctx context.Context, req *users_service.GetListScheduleRequest) (resp *users_service.GetListScheduleResponse, err error) {

	i.log.Info("---GetSchedules------>", logger.Any("req", req))

	resp, err = i.strg.Schedule().GetList(ctx, req)
	if err != nil {
		i.log.Error("!!!GetSchedules->Schedule->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *ScheduleService) Update(ctx context.Context, req *users_service.UpdateSchedule) (resp *users_service.Schedule, err error) {

	i.log.Info("---UpdateSchedule------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Schedule().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateSchedule--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Schedule().GetById(ctx, &users_service.SchedulePrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!GetTeacher->Teacher->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *ScheduleService) Delete(ctx context.Context, req *users_service.SchedulePrimaryKey) (resp *users_service.Empty, err error) {

	i.log.Info("---DeleteSchedule------>", logger.Any("req", req))

	err = i.strg.Schedule().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteSchedule->Schedule->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &users_service.Empty{}, nil
}
