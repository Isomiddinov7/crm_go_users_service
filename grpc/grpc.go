package grpc

import (
	"crm_go_users_service/config"
	"crm_go_users_service/genproto/users_service"
	"crm_go_users_service/grpc/client"
	"crm_go_users_service/grpc/service"
	"crm_go_users_service/pkg/logger"
	"crm_go_users_service/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServiceManagerI) (grpcServer *grpc.Server) {

	grpcServer = grpc.NewServer()
	users_service.RegisterSupperAdminServiceServer(grpcServer, service.NewSupperAdminService(cfg, log, strg, srvc))
	users_service.RegisterTeacherServiceServer(grpcServer, service.NewTeacherService(cfg, log, strg, srvc))
	users_service.RegisterSupportTeacherServiceServer(grpcServer, service.NewSupportTeacherService(cfg, log, strg, srvc))
	users_service.RegisterAdministrationServiceServer(grpcServer, service.NewAdministrationService(cfg, log, strg, srvc))
	users_service.RegisterStudentServiceServer(grpcServer, service.NewStudentService(cfg, log, strg, srvc))
	users_service.RegisterLessonServiceServer(grpcServer, service.NewLessonService(cfg, log, strg, srvc))
	users_service.RegisterTaskServiceServer(grpcServer, service.NewTaskService(cfg, log, strg, srvc))
	users_service.RegisterScheduleServiceServer(grpcServer, service.NewScheduleService(cfg, log, strg, srvc))
	users_service.RegisterJournalServiceServer(grpcServer, service.NewJournalService(cfg, log, strg, srvc))
	users_service.RegisterGroupServiceServer(grpcServer, service.NewGroupService(cfg, log, strg, srvc))
	users_service.RegisterBranchServiceServer(grpcServer, service.NewBranchService(cfg, log, strg, srvc))
	users_service.RegisterManagerServiceServer(grpcServer, service.NewManagerService(cfg, log, strg, srvc))
	users_service.RegisterEventServiceServer(grpcServer, service.NewEventService(cfg, log, strg, srvc))
	users_service.RegisterRoleServiceServer(grpcServer, service.NewRoleService(cfg, log, strg, srvc))
	users_service.RegisterAssignStudentServiceServer(grpcServer, service.NewAssignStudentService(cfg, log, strg, srvc))
	users_service.RegisterPaymentServiceServer(grpcServer, service.NewPaymentService(cfg, log, strg, srvc))
	users_service.RegisterScoreServiceServer(grpcServer, service.NewScoreService(cfg, log, strg, srvc))
	users_service.RegisterReportServiceServer(grpcServer, service.NewReportService(cfg, log, strg, srvc))
	reflection.Register(grpcServer)

	return
}
