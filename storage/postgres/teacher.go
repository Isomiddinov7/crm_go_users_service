package postgres

import (
	"context"
	"crm_go_users_service/genproto/users_service"
	"crm_go_users_service/pkg/helper"
	"crm_go_users_service/storage"
	"database/sql"

	"github.com/jackc/pgx/v4/pgxpool"
)

type TeacherRepo struct {
	db *pgxpool.Pool
}

func NewTeacherRepo(db *pgxpool.Pool) storage.TeacherRepoI {
	return &TeacherRepo{
		db: db,
	}
}

func (c *TeacherRepo) Create(ctx context.Context, req *users_service.CreateTeacher) (resp *users_service.TeacherPrimaryKey, err error) {
	id, err := helper.NewIncrementId(c.db, "id", "teacher", "T", 8)
	query := `INSERT INTO "teacher" (
				id,
				full_name,
				phone,
				salary,
				password,
				login,
				branch_id
				role_id,
				months_worked,
				total_sum,
				ielts_score,
				updated_at
			) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, now())
		`
	_, err = c.db.Exec(ctx,
		query,
		id(),
		req.FullName,
		req.Phone,
		req.Salary,
		req.Password,
		req.Login,
		req.BranchId,
		"beeae8a2-87b0-4f67-87c1-f6d6754b6645",
		req.MonthsWorked,
		req.TotalSum,
		req.IeltsScore,
	)

	if err != nil {
		return nil, err
	}

	return &users_service.TeacherPrimaryKey{Id: id()}, nil
}

func (c *TeacherRepo) GetById(ctx context.Context, req *users_service.TeacherPrimaryKey) (resp *users_service.Teacher, err error) {

	query := `
		SELECT
			id,
			full_name,
			phone,
			salary,
			password,
			login,
			branch_id,
			role_id,
			months_worked,
			total_sum,
			ielts_score,
			created_at,
			updated_at
		FROM "teacher"
		WHERE id = $1
	`

	var (
		id            sql.NullString
		full_name     sql.NullString
		phone         sql.NullString
		salary        sql.NullString
		password      sql.NullString
		login         sql.NullString
		branch_id     sql.NullString
		role_id       sql.NullString
		months_worked sql.NullString
		total_sum     sql.NullString
		ielts_score   sql.NullString
		createdAt     sql.NullString
		updatedAt     sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&full_name,
		&phone,
		&salary,
		&password,
		&login,
		&branch_id,
		&role_id,
		&months_worked,
		&total_sum,
		&ielts_score,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return resp, err
	}

	resp = &users_service.Teacher{
		Id:           id.String,
		FullName:     full_name.String,
		Phone:        phone.String,
		Salary:       salary.String,
		Password:     password.String,
		Login:        login.String,
		RoleId:       role_id.String,
		BranchId:     branch_id.String,
		MonthsWorked: months_worked.String,
		TotalSum:     total_sum.String,
		IeltsScore:   ielts_score.String,
		CreatedAt:    createdAt.String,
		UpdatedAt:    updatedAt.String,
	}
	return resp, nil
}

func (c *TeacherRepo) GetList(ctx context.Context, req *users_service.GetListTeacherRequest) (resp *users_service.GetListTeacherResponse, err error) {

	resp = &users_service.GetListTeacherResponse{}

	var (
		query  string
		limit  = ""
		offset = " OFFSET 0 "
		params = make(map[string]interface{})
		filter = " WHERE TRUE"
		sort   = " ORDER BY created_at DESC"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			full_name,
			phone,
			salary,
			password,
			login,
			branch_id,
			role_id,
			months_worked,
			total_sum,
			ielts_score,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "teacher"
	`

	if req.GetLimit() > 0 {
		limit = " LIMIT :limit"
		params["limit"] = req.Limit
	}

	if req.GetOffset() > 0 {
		offset = " OFFSET :offset"
		params["offset"] = req.Offset
	}

	query += filter + sort + offset + limit

	query, args := helper.ReplaceQueryParams(query, params)
	rows, err := c.db.Query(ctx, query, args...)

	if err != nil {
		return resp, err
	}

	for rows.Next() {
		var (
			id            sql.NullString
			full_name     sql.NullString
			phone         sql.NullString
			salary        sql.NullString
			password      sql.NullString
			login         sql.NullString
			branch_id     sql.NullString
			role_id       sql.NullString
			months_worked sql.NullString
			total_sum     sql.NullString
			ielts_score   sql.NullString
			createdAt     sql.NullString
			updatedAt     sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&full_name,
			&phone,
			&salary,
			&password,
			&login,
			&branch_id,
			&role_id,
			&months_worked,
			&total_sum,
			&ielts_score,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.Teachers = append(resp.Teachers, &users_service.Teacher{
			Id:           id.String,
			FullName:     full_name.String,
			Phone:        phone.String,
			Salary:       salary.String,
			Password:     password.String,
			Login:        login.String,
			BranchId:     branch_id.String,
			RoleId:       role_id.String,
			MonthsWorked: months_worked.String,
			TotalSum:     total_sum.String,
			IeltsScore:   ielts_score.String,
			CreatedAt:    createdAt.String,
			UpdatedAt:    updatedAt.String,
		})
	}

	return
}

func (c *TeacherRepo) Update(ctx context.Context, req *users_service.UpdateTeacher) (rowsAffected int64, err error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
			UPDATE
			    "teacher"
			SET
				full_name = :full_name,
				phone = :phone,
				salary = :salary,
				password = :password,
				login = :login,
				branch_id = :branch_id,
				role_id = :role_id,
				months_worked = :months_worked,
				total_sum = :total_sum,
				ielts_score = :ielts_score,
				updated_at    = now()
			WHERE
				id = :id`
	params = map[string]interface{}{
		"id":            req.GetId(),
		"full_name":     req.GetFullName(),
		"phone":         req.GetPhone(),
		"salary":        req.GetSalary(),
		"password":      req.GetPassword(),
		"login":         req.GetLogin(),
		"branch_id":     req.GetBranchId(),
		"role_id":       req.GetRoleId(),
		"months_worked": req.GetMonthsWorked(),
		"total_sum":     req.GetTotalSum(),
		"ielts_score":   req.GetIeltsScore(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), nil
}

func (c *TeacherRepo) Delete(ctx context.Context, req *users_service.TeacherPrimaryKey) error {

	query := `DELETE FROM "teacher" WHERE id = $1`

	_, err := c.db.Exec(ctx, query, req.Id)

	if err != nil {
		return err
	}

	return nil
}
