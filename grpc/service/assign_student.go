package service

import (
	"context"
	"crm_go_users_service/config"
	"crm_go_users_service/genproto/users_service"
	"crm_go_users_service/grpc/client"
	"crm_go_users_service/pkg/logger"
	"crm_go_users_service/storage"
	"errors"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AssignStudentService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*users_service.UnimplementedAssignStudentServiceServer
}

func NewAssignStudentService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *AssignStudentService {
	return &AssignStudentService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *AssignStudentService) Create(ctx context.Context, req *users_service.CreateAssignStudent) (resp *users_service.AssignStudent, err error) {

	i.log.Info("---CreateAssignStudent------>", logger.Any("req", req))
	event, err := i.strg.Event().GetById(ctx, &users_service.EventPrimaryKey{Id: req.EventId})
	if err != nil {
		i.log.Error("!!!CreateAssignStudent->GET_BY_ID_EVENT->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	temp := event.Day + " " + event.StartTime
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, temp)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	assign_time := time.Now()
	if t.Hour() > assign_time.Add(3*time.Hour).Hour() {
		id, err := i.strg.AssignStudent().Create(ctx, req)
		if err != nil {
			i.log.Error("!!!CreateAssignStudent->AssignStudent->Create--->", logger.Error(err))
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		resp, err = i.strg.AssignStudent().GetById(ctx, id)
		fmt.Println(resp)
		if err != nil {
			i.log.Error("!!!GetByIdAssignStudent->AssignStudent->Get--->", logger.Error(err))
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	} else {
		return nil, errors.New("do not assig event")
	}

	return
}

func (i *AssignStudentService) GetById(ctx context.Context, req *users_service.AssignStudentPrimaryKey) (resp *users_service.AssignStudent, err error) {
	i.log.Info("---GetAssignStudentByID------>", logger.Any("req", req))

	resp, err = i.strg.AssignStudent().GetById(ctx, req)
	if err != nil {
		i.log.Error("!!!GetAssignStudentByID->AssignStudent->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return resp, nil
}

func (i *AssignStudentService) GetList(ctx context.Context, req *users_service.GetListAssignStudentRequest) (resp *users_service.GetListAssignStudentResponse, err error) {

	i.log.Info("---GetAssignStudents------>", logger.Any("req", req))

	resp, err = i.strg.AssignStudent().GetList(ctx, req)
	if err != nil {
		i.log.Error("!!!GetAssignStudents->AssignStudent->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *AssignStudentService) Update(ctx context.Context, req *users_service.UpdateAssignStudent) (resp *users_service.AssignStudent, err error) {

	i.log.Info("---UpdateAssignStudent------>", logger.Any("req", req))

	rowsAffected, err := i.strg.AssignStudent().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateAssignStudent--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.AssignStudent().GetById(ctx, &users_service.AssignStudentPrimaryKey{Id: req.Id})
	if err != nil {
		i.log.Error("!!!GetTeacher->Teacher->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *AssignStudentService) Delete(ctx context.Context, req *users_service.AssignStudentPrimaryKey) (resp *users_service.Empty, err error) {

	i.log.Info("---DeleteAssignStudent------>", logger.Any("req", req))

	err = i.strg.AssignStudent().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteAssignStudent->AssignStudent->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &users_service.Empty{}, nil
}
