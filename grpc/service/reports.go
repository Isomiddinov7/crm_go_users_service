package service

import (
	"context"
	"crm_go_users_service/config"
	"crm_go_users_service/genproto/users_service"
	"crm_go_users_service/grpc/client"
	"crm_go_users_service/pkg/logger"
	"crm_go_users_service/storage"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ReportService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*users_service.UnimplementedReportServiceServer
}

func NewReportService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *ReportService {
	return &ReportService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *ReportService) GetTeacher(ctx context.Context, req *users_service.ReportRequest) (resp *users_service.TeacherList, err error) {

	i.log.Info("---Get Teacher Report------>", logger.Any("req", req))

	resp, err = i.strg.Reports().GetTeacher(ctx, req)
	if err != nil {
		i.log.Error("!!!GetTeacher->Report->GET--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	fmt.Println(resp)

	return
}

func (i *ReportService) GetStudent(ctx context.Context, req *users_service.ReportRequest) (resp *users_service.StudentList, err error) {

	i.log.Info("---Get Student Report------>", logger.Any("req", req))

	resp, err = i.strg.Reports().GetStudent(ctx, req)
	if err != nil {
		i.log.Error("!!!GetStudent->Report->GET--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	fmt.Println(resp)

	return
}

func (i *ReportService) GetAdministrator(ctx context.Context, req *users_service.ReportRequest) (resp *users_service.AdminList, err error) {

	i.log.Info("---Get Admin Report------>", logger.Any("req", req))

	resp, err = i.strg.Reports().GetAdministrator(ctx, req)
	if err != nil {
		i.log.Error("!!!GetAdmin->Report->GET--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	fmt.Println(resp)

	return
}

func (i *ReportService) GetSupportTeacher(ctx context.Context, req *users_service.ReportRequest) (resp *users_service.SupportTeacherList, err error) {

	i.log.Info("---Get SupportTeacher Report------>", logger.Any("req", req))

	resp, err = i.strg.Reports().GetSupportTeacher(ctx, req)
	if err != nil {
		i.log.Error("!!!GetSupportTeacher->Report->GET--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	fmt.Println(resp)

	return
}

func (i *ReportService) GetStudentGroupByTeacher(ctx context.Context, req *users_service.GroupId) (resp *users_service.StudentReportByTeacher, err error) {

	i.log.Info("---Get GetStudentGroupByTeacher Report------>", logger.Any("req", req))

	studentlist, err := i.strg.Reports().GetStudentGroup(ctx, &users_service.GroupId{GroupId: req.GroupId})
	if err != nil {
		i.log.Error("!!!TeacherReportStudentList->Report->GET--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	teacherId, err := i.strg.Reports().GetTeacherGroup(ctx, &users_service.GroupId{GroupId: req.GroupId})
	fmt.Println(teacherId.Id)
	if err != nil {
		i.log.Error("!!!GetTeacherGroup->Report->GET--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	teacher, err := i.strg.Teacher().GetById(ctx, &users_service.TeacherPrimaryKey{Id: teacherId.Id})
	if err != nil {
		i.log.Error("!!!TeacherPrimaryKey->Report->GET--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	count := 0
	for i := range studentlist.StudentList {
		resp.StudentList = append(resp.StudentList, studentlist.StudentList[i])
		count++
	}

	resp.Count = int64(count)
	resp.TeacherName = teacher.FullName
	return
}
