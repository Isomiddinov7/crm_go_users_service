package postgres

import (
	"context"
	"crm_go_users_service/genproto/users_service"
	"crm_go_users_service/storage"
	"database/sql"
	"errors"

	"github.com/jackc/pgx/v4/pgxpool"
)

type ReportRepo struct {
	db *pgxpool.Pool
}

func NewReportRepo(db *pgxpool.Pool) storage.ReportsRepoI {
	return &ReportRepo{
		db: db,
	}
}

func (r *ReportRepo) GetStudent(ctx context.Context, req *users_service.ReportRequest) (resp *users_service.StudentList, err error) {
	resp = &users_service.StudentList{}
	query := `
		SELECT 
			full_name,
			phone,
			paid_sum,
			cource_count,
			total_sum
		FROM "student"
		GROUP BY "id"
		HAVING branch_id = $1
	`

	rows, err := r.db.Query(ctx, query, req.BranchId)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			full_name    sql.NullString
			phone        sql.NullString
			paid_sum     sql.NullString
			cource_count sql.NullString
			total_sum    sql.NullString
		)

		err = rows.Scan(
			&full_name,
			&phone,
			&paid_sum,
			&cource_count,
			&total_sum,
		)
		if err != nil {
			return resp, err
		}
		resp.StudentList = append(resp.StudentList, &users_service.GetStudent{
			FullName:    full_name.String,
			Phone:       phone.String,
			PaidSum:     paid_sum.String,
			CourceCount: cource_count.String,
			TotalSum:    total_sum.String,
		})

	}

	if len(resp.StudentList) == 0 {
		return nil, errors.New("no data found")
	}

	return
}

func (r *ReportRepo) GetTeacher(ctx context.Context, req *users_service.ReportRequest) (resp *users_service.TeacherList, err error) {
	resp = &users_service.TeacherList{}
	query := `
		SELECT 
			full_name,
			phone,
			salary,
			password,
			login,
			months_worked,
			total_sum,
			ielts_score
		FROM "teacher"
		GROUP BY "id"
		HAVING branch_id = $1
	`

	rows, err := r.db.Query(ctx, query, req.BranchId)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			full_name     sql.NullString
			phone         sql.NullString
			salary        sql.NullString
			password      sql.NullString
			login         sql.NullString
			months_worked sql.NullString
			total_sum     sql.NullString
			ielts_score   sql.NullString
		)

		err = rows.Scan(
			&full_name,
			&phone,
			&salary,
			&password,
			&login,
			&months_worked,
			&total_sum,
			&ielts_score,
		)
		if err != nil {
			return resp, err
		}
		resp.TeacherList = append(resp.TeacherList, &users_service.GetTeacher{
			FullName:     full_name.String,
			Phone:        phone.String,
			Salary:       salary.String,
			Password:     password.String,
			Login:        login.String,
			MonthsWorked: months_worked.String,
			TotalSum:     total_sum.String,
			IeltsScore:   ielts_score.String,
		})

	}

	if len(resp.TeacherList) == 0 {
		return nil, errors.New("no data found")
	}

	return
}

func (r *ReportRepo) GetSupportTeacher(ctx context.Context, req *users_service.ReportRequest) (resp *users_service.SupportTeacherList, err error) {
	resp = &users_service.SupportTeacherList{}
	query := `
		SELECT 
			full_name,
			phone,
			salary,
			password,
			login,
			months_worked,
			total_sum,
			ielts_score
		FROM "support_teacher"
		GROUP BY "id"
		HAVING branch_id = $1
	`

	rows, err := r.db.Query(ctx, query, req.BranchId)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			full_name     sql.NullString
			phone         sql.NullString
			salary        sql.NullString
			password      sql.NullString
			login         sql.NullString
			months_worked sql.NullString
			total_sum     sql.NullString
			ielts_score   sql.NullString
		)

		err = rows.Scan(
			&full_name,
			&phone,
			&salary,
			&password,
			&login,
			&months_worked,
			&total_sum,
			&ielts_score,
		)
		if err != nil {
			return resp, err
		}
		resp.SupportTeacherList = append(resp.SupportTeacherList, &users_service.GetSupportTeacher{
			FullName:     full_name.String,
			Phone:        phone.String,
			Salary:       salary.String,
			Password:     password.String,
			Login:        login.String,
			MonthsWorked: months_worked.String,
			TotalSum:     total_sum.String,
			IeltsScore:   ielts_score.String,
		})

	}

	if len(resp.SupportTeacherList) == 0 {
		return nil, errors.New("no data found")
	}

	return
}

func (r *ReportRepo) GetAdministrator(ctx context.Context, req *users_service.ReportRequest) (resp *users_service.AdminList, err error) {
	resp = &users_service.AdminList{}
	query := `
		SELECT 
			full_name,
			phone,
			salary,
			password,
			login,
			months_worked,
			total_sum,
			ielts_score
		FROM "administrator"
		GROUP BY "id"
		HAVING branch_id = $1
	`

	rows, err := r.db.Query(ctx, query, req.BranchId)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			full_name     sql.NullString
			phone         sql.NullString
			salary        sql.NullString
			password      sql.NullString
			login         sql.NullString
			months_worked sql.NullString
			total_sum     sql.NullString
			ielts_score   sql.NullString
		)

		err = rows.Scan(
			&full_name,
			&phone,
			&salary,
			&password,
			&login,
			&months_worked,
			&total_sum,
			&ielts_score,
		)
		if err != nil {
			return resp, err
		}
		resp.AdminList = append(resp.AdminList, &users_service.GetAdministrator{
			FullName:   full_name.String,
			Phone:      phone.String,
			Salary:     salary.String,
			Password:   password.String,
			Login:      login.String,
			IeltsScore: ielts_score.String,
		})

	}

	if len(resp.AdminList) == 0 {
		return nil, errors.New("no data found")
	}

	return
}

func (r *ReportRepo) GetTeacherGroup(ctx context.Context, req *users_service.GroupId) (resp *users_service.TeacherId, err error) {
	resp = &users_service.TeacherId{}
	query := `
		SELECT
			"teacher_id"
		FROM
			"group"
		WHERE id = $1;
	`
	rows, err := r.db.Query(ctx, query, req.GroupId)

	if err != nil {
		return resp, err
	}

	for rows.Next() {
		var (
			teacher_id sql.NullString
		)

		err := rows.Scan(
			&teacher_id,
		)

		if err != nil {
			return resp, err
		}
	}

	return
}

func (r *ReportRepo) GetStudentGroup(ctx context.Context, req *users_service.GroupId) (resp *users_service.StudentList, err error) {
	resp = &users_service.StudentList{}
	query := `
		SELECT 
			full_name,
			phone,
			paid_sum,
			cource_count,
			total_sum
		FROM "student"
		GROUP BY "id"
		HAVING group_id = $1
	`

	rows, err := r.db.Query(ctx, query, req.GroupId)
	if err != nil {

		return resp, err
	}

	for rows.Next() {
		var (
			full_name    sql.NullString
			phone        sql.NullString
			paid_sum     sql.NullString
			cource_count sql.NullString
			total_sum    sql.NullString
		)

		err = rows.Scan(
			&full_name,
			&phone,
			&paid_sum,
			&cource_count,
			&total_sum,
		)
		if err != nil {
			return resp, err
		}
		resp.StudentList = append(resp.StudentList, &users_service.GetStudent{
			FullName:    full_name.String,
			Phone:       phone.String,
			PaidSum:     paid_sum.String,
			CourceCount: cource_count.String,
			TotalSum:    total_sum.String,
		})

	}

	if len(resp.StudentList) == 0 {
		return nil, errors.New("no data found")
	}

	return
}
