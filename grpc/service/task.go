package service

import (
	"context"
	"crm_go_users_service/config"
	"crm_go_users_service/genproto/users_service"
	"crm_go_users_service/grpc/client"
	"crm_go_users_service/pkg/logger"
	"crm_go_users_service/storage"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TaskService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*users_service.UnimplementedTaskServiceServer
}

func NewTaskService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *TaskService {
	return &TaskService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *TaskService) Create(ctx context.Context, req *users_service.CreateTask) (resp *users_service.Task, err error) {

	i.log.Info("---CreateTask------>", logger.Any("req", req))

	id, err := i.strg.Task().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateTask->Task->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp, err = i.strg.Task().GetById(ctx, id)
	if err != nil {
		i.log.Error("!!!GetByIdTask->Teacher->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return
}

func (i *TaskService) GetById(ctx context.Context, req *users_service.TaskPrimaryKey) (resp *users_service.Task, err error) {
	i.log.Info("---GetTaskByID------>", logger.Any("req", req))

	resp, err = i.strg.Task().GetById(ctx, req)
	if err != nil {
		i.log.Error("!!!GetTaskByID->Task->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return resp, nil
}

func (i *TaskService) GetList(ctx context.Context, req *users_service.GetListTaskRequest) (resp *users_service.GetListTaskResponse, err error) {

	i.log.Info("---GetTasks------>", logger.Any("req", req))

	resp, err = i.strg.Task().GetList(ctx, req)
	if err != nil {
		i.log.Error("!!!GetTasks->Task->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *TaskService) Update(ctx context.Context, req *users_service.UpdateTask) (resp *users_service.Task, err error) {

	i.log.Info("---UpdateTask------>", logger.Any("req", req))
	task, err := i.strg.Task().GetById(ctx, &users_service.TaskPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!UpdateTask--GetById--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	deadlineTime, err := time.Parse(time.RFC3339, task.Deadline)
	if err != nil {
		i.log.Error("!!!ParseDeadline--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, "failed to parse deadline")
	}

	deadlineTime1, err := time.Parse(time.RFC3339, req.Deadline)
	if err != nil {
		i.log.Error("!!!ParseDeadline--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, "failed to parse deadline")
	}
	if deadlineTime1.YearDay()-deadlineTime.YearDay() < 3 {
		req.Score /= 2
		rowsAffected, err := i.strg.Task().Update(ctx, req)

		if err != nil {
			i.log.Error("!!!UpdateTask--->", logger.Error(err))
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		if rowsAffected <= 0 {
			return nil, status.Error(codes.InvalidArgument, "no rows were affected")
		}
	} else {
		req.Score = 0
		rowsAffected, err := i.strg.Task().Update(ctx, req)
		if err != nil {
			i.log.Error("!!!UpdateTask--->", logger.Error(err))
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		if rowsAffected <= 0 {
			return nil, status.Error(codes.InvalidArgument, "no rows were affected")
		}
	}

	resp, err = i.strg.Task().GetById(ctx, &users_service.TaskPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!GetTeacher->Teacher->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return resp, err
}

func (i *TaskService) Delete(ctx context.Context, req *users_service.TaskPrimaryKey) (resp *users_service.Empty, err error) {

	i.log.Info("---DeleteTask------>", logger.Any("req", req))

	err = i.strg.Task().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteTask->Task->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &users_service.Empty{}, nil
}
