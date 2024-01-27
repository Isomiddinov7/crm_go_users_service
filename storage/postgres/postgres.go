package postgres

import (
	"context"
	"crm_go_users_service/config"
	"crm_go_users_service/storage"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db             *pgxpool.Pool
	super_admin    storage.SuperAdminRepoI
	teacher        storage.TeacherRepoI
	supportteacher storage.SupportTeacherRepoI
	administrator  storage.AdministrationRepoI
	student        storage.StudentRepoI
	lesson         storage.LessonRepoI
	task           storage.TaskRepoI
	schedule       storage.ScheduleRepoI
	journal        storage.JournalRepoI
	group          storage.GroupRepoI
	branch         storage.BranchRepoI
	manager        storage.ManagerRepoI
	event          storage.EventRepoI
	role           storage.RoleRepoI
	assign_student storage.AssignStudentRepoI
	payment        storage.PaymentRepoI
	score          storage.ScoreRepoI
	report         storage.ReportsRepoI
}

func NewPostgres(ctx context.Context, cfg config.Config) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))
	if err != nil {
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnections

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: pool,
	}, err
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (l *Store) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	args := make([]interface{}, 0, len(data)+2) // making space for arguments + level + msg
	args = append(args, level, msg)
	for k, v := range data {
		args = append(args, fmt.Sprintf("%s=%v", k, v))
	}
	log.Println(args...)
}

func (s *Store) SuperAdmin() storage.SuperAdminRepoI {
	if s.super_admin == nil {
		s.super_admin = NewSuperAdminRepo(s.db)
	}

	return s.super_admin
}

func (s *Store) Teacher() storage.TeacherRepoI {
	if s.teacher == nil {
		s.teacher = NewTeacherRepo(s.db)
	}

	return s.teacher
}

func (s *Store) SupportTeacher() storage.SupportTeacherRepoI {
	if s.supportteacher == nil {
		s.supportteacher = NewSupportTeacherRepo(s.db)
	}

	return s.supportteacher
}

func (s *Store) Administrator() storage.AdministrationRepoI {
	if s.administrator == nil {
		s.administrator = NewAdministrationRepo(s.db)
	}

	return s.administrator
}

func (s *Store) Student() storage.StudentRepoI {
	if s.student == nil {
		s.student = NewStudentRepo(s.db)
	}

	return s.student
}

func (s *Store) Lesson() storage.LessonRepoI {
	if s.lesson == nil {
		s.lesson = NewLessonRepo(s.db)
	}

	return s.lesson
}

func (s *Store) Task() storage.TaskRepoI {
	if s.task == nil {
		s.task = NewTaskRepo(s.db)
	}

	return s.task
}

func (s *Store) Schedule() storage.ScheduleRepoI {
	if s.schedule == nil {
		s.schedule = NewScheduleRepo(s.db)
	}

	return s.schedule
}

func (s *Store) Journal() storage.JournalRepoI {
	if s.journal == nil {
		s.journal = NewJournalRepo(s.db)
	}

	return s.journal
}

func (s *Store) Group() storage.GroupRepoI {
	if s.group == nil {
		s.group = NewGroupRepo(s.db)
	}

	return s.group
}

func (s *Store) Branch() storage.BranchRepoI {
	if s.branch == nil {
		s.branch = NewBranchRepo(s.db)
	}

	return s.branch
}

func (s *Store) Manager() storage.ManagerRepoI {
	if s.manager == nil {
		s.manager = NewManagerRepo(s.db)
	}

	return s.manager
}

func (s *Store) Event() storage.EventRepoI {
	if s.event == nil {
		s.event = NewEventRepo(s.db)
	}

	return s.event
}

func (s *Store) Role() storage.RoleRepoI {
	if s.role == nil {
		s.role = NewRoleRepo(s.db)
	}

	return s.role
}

func (s *Store) AssignStudent() storage.AssignStudentRepoI {
	if s.assign_student == nil {
		s.assign_student = NewAssignStudentRepo(s.db)
	}

	return s.assign_student
}

func (s *Store) Payment() storage.PaymentRepoI {
	if s.payment == nil {
		s.payment = NewPaymentRepo(s.db)
	}

	return s.payment
}

func (s *Store) Score() storage.ScoreRepoI {
	if s.score == nil {
		s.score = NewScoreRepo(s.db)
	}

	return s.score
}

func (s *Store) Reports() storage.ReportsRepoI {
	if s.report == nil {
		s.report = NewReportRepo(s.db)
	}

	return s.report
}
