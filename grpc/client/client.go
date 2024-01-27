package client

import (
	"crm_go_users_service/config"
	"crm_go_users_service/genproto/users_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManagerI interface {
	TeacherService() users_service.TeacherServiceClient
	SupportTeacherService() users_service.SupportTeacherServiceClient
	AdministrationService() users_service.AdministrationServiceClient
	StudentService() users_service.StudentServiceClient
	LessonService() users_service.LessonServiceClient
	TaskService() users_service.TaskServiceClient
	ScheduleService() users_service.ScheduleServiceClient
	JournalService() users_service.JournalServiceClient
	GroupService() users_service.GroupServiceClient
	BranchService() users_service.BranchServiceClient
	ManagerService() users_service.ManagerServiceClient
	SchedulesService() users_service.SchedulesServiceClient
	EventService() users_service.EventServiceClient
	RoleService() users_service.RoleServiceClient
	AssignStudentService() users_service.AssignStudentServiceClient
	PaymentService() users_service.PaymentServiceClient
	ScoreService() users_service.ScoreServiceClient
	ReportService() users_service.ReportServiceClient
}

type grpcClients struct {
	teacherService        users_service.TeacherServiceClient
	supportteacherService users_service.SupportTeacherServiceClient
	administrationService users_service.AdministrationServiceClient
	studentService        users_service.StudentServiceClient
	lessonService         users_service.LessonServiceClient
	taskService           users_service.TaskServiceClient
	scheduleService       users_service.ScheduleServiceClient
	journalService        users_service.JournalServiceClient
	groupService          users_service.GroupServiceClient
	branchService         users_service.BranchServiceClient
	managerService        users_service.ManagerServiceClient
	schedulesService      users_service.SchedulesServiceClient
	eventService          users_service.EventServiceClient
	roleService           users_service.RoleServiceClient
	assign_studentService users_service.AssignStudentServiceClient
	paymentService        users_service.PaymentServiceClient
	scoreService          users_service.ScoreServiceClient
	reportService         users_service.ReportServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {
	connUsersService, err := grpc.Dial(
		cfg.UsersServiceHost+cfg.UsersGRPCPort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	return &grpcClients{
		teacherService:        users_service.NewTeacherServiceClient(connUsersService),
		supportteacherService: users_service.NewSupportTeacherServiceClient(connUsersService),
		administrationService: users_service.NewAdministrationServiceClient(connUsersService),
		studentService:        users_service.NewStudentServiceClient(connUsersService),
		lessonService:         users_service.NewLessonServiceClient(connUsersService),
		taskService:           users_service.NewTaskServiceClient(connUsersService),
		scheduleService:       users_service.NewScheduleServiceClient(connUsersService),
		journalService:        users_service.NewJournalServiceClient(connUsersService),
		groupService:          users_service.NewGroupServiceClient(connUsersService),
		branchService:         users_service.NewBranchServiceClient(connUsersService),
		managerService:        users_service.NewManagerServiceClient(connUsersService),
		schedulesService:      users_service.NewSchedulesServiceClient(connUsersService),
		eventService:          users_service.NewEventServiceClient(connUsersService),
		roleService:           users_service.NewRoleServiceClient(connUsersService),
		assign_studentService: users_service.NewAssignStudentServiceClient(connUsersService),
		paymentService:        users_service.NewPaymentServiceClient(connUsersService),
		scoreService:          users_service.NewScoreServiceClient(connUsersService),
		reportService:         users_service.NewReportServiceClient(connUsersService),
	}, nil
}

func (g *grpcClients) TeacherService() users_service.TeacherServiceClient {
	return g.teacherService
}

func (g *grpcClients) SupportTeacherService() users_service.SupportTeacherServiceClient {
	return g.supportteacherService
}

func (g *grpcClients) AdministrationService() users_service.AdministrationServiceClient {
	return g.administrationService
}

func (g *grpcClients) StudentService() users_service.StudentServiceClient {
	return g.studentService
}

func (g *grpcClients) LessonService() users_service.LessonServiceClient {
	return g.lessonService
}

func (g *grpcClients) TaskService() users_service.TaskServiceClient {
	return g.taskService
}

func (g *grpcClients) ScheduleService() users_service.ScheduleServiceClient {
	return g.scheduleService
}

func (g *grpcClients) JournalService() users_service.JournalServiceClient {
	return g.journalService
}

func (g *grpcClients) GroupService() users_service.GroupServiceClient {
	return g.groupService
}

func (g *grpcClients) BranchService() users_service.BranchServiceClient {
	return g.branchService
}

func (g *grpcClients) ManagerService() users_service.ManagerServiceClient {
	return g.managerService
}

func (g *grpcClients) SchedulesService() users_service.SchedulesServiceClient {
	return g.schedulesService
}

func (g *grpcClients) EventService() users_service.EventServiceClient {
	return g.eventService
}

func (g *grpcClients) RoleService() users_service.RoleServiceClient {
	return g.roleService
}

func (g *grpcClients) AssignStudentService() users_service.AssignStudentServiceClient {
	return g.assign_studentService
}

func (g *grpcClients) PaymentService() users_service.PaymentServiceClient {
	return g.paymentService
}

func (g *grpcClients) ScoreService() users_service.ScoreServiceClient {
	return g.scoreService
}

func (g *grpcClients) ReportService() users_service.ReportServiceClient {
	return g.reportService
}
