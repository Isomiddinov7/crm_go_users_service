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

type LessonService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*users_service.UnimplementedLessonServiceServer
}

func NewLessonService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *LessonService {
	return &LessonService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *LessonService) Create(ctx context.Context, req *users_service.CreateLesson) (resp *users_service.Lesson, err error) {

	i.log.Info("---CreateLesson------>", logger.Any("req", req))
	id, err := i.strg.Lesson().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateLesson->Lesson->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err != nil {
		i.log.Error("!!!GetByIdTask->task->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp, err = i.strg.Lesson().GetById(ctx, id)
	if err != nil {
		i.log.Error("!!!GetByIdLesson->Lesson->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return
}

func (i *LessonService) GetById(ctx context.Context, req *users_service.LessonPrimaryKey) (resp *users_service.Lesson, err error) {
	i.log.Info("---GetLessonByID------>", logger.Any("req", req))

	resp, err = i.strg.Lesson().GetById(ctx, req)
	if err != nil {
		i.log.Error("!!!GetLessonByID->Lesson->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return resp, nil
}

func (i *LessonService) GetList(ctx context.Context, req *users_service.GetListLessonRequest) (resp *users_service.GetListLessonResponse, err error) {

	i.log.Info("---GetLessons------>", logger.Any("req", req))

	resp, err = i.strg.Lesson().GetList(ctx, req)
	if err != nil {
		i.log.Error("!!!GetLessons->Lesson->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *LessonService) Update(ctx context.Context, req *users_service.UpdateLesson) (resp *users_service.Lesson, err error) {

	i.log.Info("---UpdateLesson------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Lesson().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateLesson--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Lesson().GetById(ctx, &users_service.LessonPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!GetTeacher->Teacher->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *LessonService) Delete(ctx context.Context, req *users_service.LessonPrimaryKey) (resp *users_service.Empty, err error) {

	i.log.Info("---DeleteLesson------>", logger.Any("req", req))

	err = i.strg.Lesson().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteLesson->Lesson->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &users_service.Empty{}, nil
}
