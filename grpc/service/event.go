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

type EventService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*users_service.UnimplementedEventServiceServer
}

func NewEventService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *EventService {
	return &EventService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *EventService) Create(ctx context.Context, req *users_service.CreateEvent) (resp *users_service.Event, err error) {

	i.log.Info("---CreateEvent------>", logger.Any("req", req))
	id, err := i.strg.Event().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateEvent->Event->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err != nil {
		i.log.Error("!!!GetByIdTask->task->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp, err = i.strg.Event().GetById(ctx, id)
	if err != nil {
		i.log.Error("!!!GetByIdEvent->Event->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return
}

func (i *EventService) GetById(ctx context.Context, req *users_service.EventPrimaryKey) (resp *users_service.Event, err error) {
	i.log.Info("---GetEventByID------>", logger.Any("req", req))

	resp, err = i.strg.Event().GetById(ctx, req)
	if err != nil {
		i.log.Error("!!!GetEventByID->Event->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return resp, nil
}

func (i *EventService) GetList(ctx context.Context, req *users_service.GetListEventRequest) (resp *users_service.GetListEventResponse, err error) {

	i.log.Info("---GetEvents------>", logger.Any("req", req))

	resp, err = i.strg.Event().GetList(ctx, req)
	if err != nil {
		i.log.Error("!!!GetEvents->Event->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *EventService) Update(ctx context.Context, req *users_service.UpdateEvent) (resp *users_service.Event, err error) {

	i.log.Info("---UpdateEvent------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Event().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateEvent--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Event().GetById(ctx, &users_service.EventPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!GetTeacher->Teacher->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *EventService) Delete(ctx context.Context, req *users_service.EventPrimaryKey) (resp *users_service.Empty, err error) {

	i.log.Info("---DeleteEvent------>", logger.Any("req", req))

	err = i.strg.Event().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteEvent->Event->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &users_service.Empty{}, nil
}
