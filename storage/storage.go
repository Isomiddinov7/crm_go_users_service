package storage

import (
	"context"
	"crm_go_users_service/genproto/users_service"
)

type StorageI interface {
	CloseDB()
	SuperAdmin() SuperAdminRepoI
	Teacher() TeacherRepoI
	SupportTeacher() SupportTeacherRepoI
	Administrator() AdministrationRepoI
	Student() StudentRepoI
	Lesson() LessonRepoI
	Task() TaskRepoI
	Schedule() ScheduleRepoI
	Journal() JournalRepoI
	Group() GroupRepoI
	Branch() BranchRepoI
	Manager() ManagerRepoI
	Event() EventRepoI
	Role() RoleRepoI
	AssignStudent() AssignStudentRepoI
	Payment() PaymentRepoI
	Score() ScoreRepoI
	Reports() ReportsRepoI
}

type SuperAdminRepoI interface {
	Create(ctx context.Context, req *users_service.CreateSupperAdmin) (resp *users_service.SupperAdminPrimaryKey, err error)
	GetById(ctx context.Context, req *users_service.SupperAdminPrimaryKey) (resp *users_service.SupperAdmin, err error)
	GetList(ctx context.Context, req *users_service.GetListSupperAdminRequest) (resp *users_service.GetListSupperAdminResponse, err error)
	Update(ctx context.Context, req *users_service.UpdateSupperAdmin) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *users_service.SupperAdminPrimaryKey) error
}

type TeacherRepoI interface {
	Create(ctx context.Context, req *users_service.CreateTeacher) (resp *users_service.TeacherPrimaryKey, err error)
	GetById(ctx context.Context, req *users_service.TeacherPrimaryKey) (resp *users_service.Teacher, err error)
	GetList(ctx context.Context, req *users_service.GetListTeacherRequest) (resp *users_service.GetListTeacherResponse, err error)
	Update(ctx context.Context, req *users_service.UpdateTeacher) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *users_service.TeacherPrimaryKey) error
}

type SupportTeacherRepoI interface {
	Create(ctx context.Context, req *users_service.CreateSupportTeacher) (resp *users_service.SupportTeacherPrimaryKey, err error)
	GetById(ctx context.Context, req *users_service.SupportTeacherPrimaryKey) (resp *users_service.SupportTeacher, err error)
	GetList(ctx context.Context, req *users_service.GetListSupportTeacherRequest) (resp *users_service.GetListSupportTeacherResponse, err error)
	Update(ctx context.Context, req *users_service.UpdateSupportTeacher) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *users_service.SupportTeacherPrimaryKey) error
}

type AdministrationRepoI interface {
	Create(ctx context.Context, req *users_service.CreateAdministration) (resp *users_service.AdministrationPrimaryKey, err error)
	GetById(ctx context.Context, req *users_service.AdministrationPrimaryKey) (resp *users_service.Administration, err error)
	GetList(ctx context.Context, req *users_service.GetListAdministrationRequest) (resp *users_service.GetListAdministrationResponse, err error)
	Update(ctx context.Context, req *users_service.UpdateAdministration) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *users_service.AdministrationPrimaryKey) error
}

type StudentRepoI interface {
	Create(ctx context.Context, req *users_service.CreateStudent) (resp *users_service.StudentPrimaryKey, err error)
	GetById(ctx context.Context, req *users_service.StudentPrimaryKey) (resp *users_service.Student, err error)
	GetList(ctx context.Context, req *users_service.GetListStudentRequest) (resp *users_service.GetListStudentResponse, err error)
	Update(ctx context.Context, req *users_service.UpdateStudent) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *users_service.StudentPrimaryKey) error
}

type LessonRepoI interface {
	Create(ctx context.Context, req *users_service.CreateLesson) (resp *users_service.LessonPrimaryKey, err error)
	GetById(ctx context.Context, req *users_service.LessonPrimaryKey) (resp *users_service.Lesson, err error)
	GetList(ctx context.Context, req *users_service.GetListLessonRequest) (resp *users_service.GetListLessonResponse, err error)
	Update(ctx context.Context, req *users_service.UpdateLesson) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *users_service.LessonPrimaryKey) error
}

type TaskRepoI interface {
	Create(ctx context.Context, req *users_service.CreateTask) (resp *users_service.TaskPrimaryKey, err error)
	GetById(ctx context.Context, req *users_service.TaskPrimaryKey) (resp *users_service.Task, err error)
	GetList(ctx context.Context, req *users_service.GetListTaskRequest) (resp *users_service.GetListTaskResponse, err error)
	Update(ctx context.Context, req *users_service.UpdateTask) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *users_service.TaskPrimaryKey) error
}

type ScheduleRepoI interface {
	Create(ctx context.Context, req *users_service.CreateSchedule) (resp *users_service.SchedulePrimaryKey, err error)
	GetById(ctx context.Context, req *users_service.SchedulePrimaryKey) (resp *users_service.Schedule, err error)
	GetList(ctx context.Context, req *users_service.GetListScheduleRequest) (resp *users_service.GetListScheduleResponse, err error)
	Update(ctx context.Context, req *users_service.UpdateSchedule) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *users_service.SchedulePrimaryKey) error
}

type JournalRepoI interface {
	Create(ctx context.Context, req *users_service.CreateJournal) (resp *users_service.JournalPrimaryKey, err error)
	GetById(ctx context.Context, req *users_service.JournalPrimaryKey) (resp *users_service.Journal, err error)
	GetList(ctx context.Context, req *users_service.GetListJournalRequest) (resp *users_service.GetListJournalResponse, err error)
	Update(ctx context.Context, req *users_service.UpdateJournal) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *users_service.JournalPrimaryKey) error
}

type GroupRepoI interface {
	Create(ctx context.Context, req *users_service.CreateGroup) (resp *users_service.GroupPrimaryKey, err error)
	GetById(ctx context.Context, req *users_service.GroupPrimaryKey) (resp *users_service.Group, err error)
	GetList(ctx context.Context, req *users_service.GetListGroupRequest) (resp *users_service.GetListGroupResponse, err error)
	Update(ctx context.Context, req *users_service.UpdateGroup) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *users_service.GroupPrimaryKey) error
}

type BranchRepoI interface {
	Create(ctx context.Context, req *users_service.CreateBranch) (resp *users_service.BranchPrimaryKey, err error)
	GetById(ctx context.Context, req *users_service.BranchPrimaryKey) (resp *users_service.Branch, err error)
	GetList(ctx context.Context, req *users_service.GetListBranchRequest) (resp *users_service.GetListBranchResponse, err error)
	Update(ctx context.Context, req *users_service.UpdateBranch) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *users_service.BranchPrimaryKey) error
}

type ManagerRepoI interface {
	Create(ctx context.Context, req *users_service.CreateManager) (resp *users_service.ManagerPrimaryKey, err error)
	GetById(ctx context.Context, req *users_service.ManagerPrimaryKey) (resp *users_service.Manager, err error)
	GetList(ctx context.Context, req *users_service.GetListManagerRequest) (resp *users_service.GetListManagerResponse, err error)
	Update(ctx context.Context, req *users_service.UpdateManager) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *users_service.ManagerPrimaryKey) error
}

type EventRepoI interface {
	Create(ctx context.Context, req *users_service.CreateEvent) (resp *users_service.EventPrimaryKey, err error)
	GetById(ctx context.Context, req *users_service.EventPrimaryKey) (resp *users_service.Event, err error)
	GetList(ctx context.Context, req *users_service.GetListEventRequest) (resp *users_service.GetListEventResponse, err error)
	Update(ctx context.Context, req *users_service.UpdateEvent) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *users_service.EventPrimaryKey) error
}

type RoleRepoI interface {
	Create(ctx context.Context, req *users_service.CreateRole) (resp *users_service.RolePrimaryKey, err error)
	GetById(ctx context.Context, req *users_service.RolePrimaryKey) (resp *users_service.Role, err error)
	GetList(ctx context.Context, req *users_service.GetListRoleRequest) (resp *users_service.GetListRoleResponse, err error)
	Update(ctx context.Context, req *users_service.UpdateRole) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *users_service.RolePrimaryKey) error
}

type AssignStudentRepoI interface {
	Create(ctx context.Context, req *users_service.CreateAssignStudent) (resp *users_service.AssignStudentPrimaryKey, err error)
	GetById(ctx context.Context, req *users_service.AssignStudentPrimaryKey) (resp *users_service.AssignStudent, err error)
	GetList(ctx context.Context, req *users_service.GetListAssignStudentRequest) (resp *users_service.GetListAssignStudentResponse, err error)
	Update(ctx context.Context, req *users_service.UpdateAssignStudent) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *users_service.AssignStudentPrimaryKey) error
}

type PaymentRepoI interface {
	Create(ctx context.Context, req *users_service.CreatePayment) (resp *users_service.PaymentPrimaryKey, err error)
	GetById(ctx context.Context, req *users_service.PaymentPrimaryKey) (resp *users_service.Payment, err error)
	GetList(ctx context.Context, req *users_service.GetListPaymentRequest) (resp *users_service.GetListPaymentResponse, err error)
	Update(ctx context.Context, req *users_service.UpdatePayment) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *users_service.PaymentPrimaryKey) error
}

type ScoreRepoI interface {
	Create(ctx context.Context, req *users_service.CreateScore) (resp *users_service.ScorePrimaryKey, err error)
	GetById(ctx context.Context, req *users_service.ScorePrimaryKey) (resp *users_service.Score, err error)
	GetList(ctx context.Context, req *users_service.GetListScoreRequest) (resp *users_service.GetListScoreResponse, err error)
	Update(ctx context.Context, req *users_service.UpdateScore) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *users_service.ScorePrimaryKey) error
}

type ReportsRepoI interface {
	GetStudent(ctx context.Context, req *users_service.ReportRequest) (resp *users_service.StudentList, err error)
	GetTeacher(ctx context.Context, req *users_service.ReportRequest) (resp *users_service.TeacherList, err error)
	GetSupportTeacher(ctx context.Context, req *users_service.ReportRequest) (resp *users_service.SupportTeacherList, err error)
	GetAdministrator(ctx context.Context, req *users_service.ReportRequest) (resp *users_service.AdminList, err error)
	GetStudentGroup(ctx context.Context, req *users_service.GroupId) (resp *users_service.StudentList, err error)
	GetTeacherGroup(ctx context.Context, req *users_service.GroupId) (resp *users_service.TeacherId, err error)
}
