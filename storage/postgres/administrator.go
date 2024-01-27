package postgres

import (
	"context"
	"crm_go_users_service/genproto/users_service"
	"crm_go_users_service/pkg/helper"
	"crm_go_users_service/storage"
	"database/sql"

	"github.com/jackc/pgx/v4/pgxpool"
)

type AdministrationRepo struct {
	db *pgxpool.Pool
}

func NewAdministrationRepo(db *pgxpool.Pool) storage.AdministrationRepoI {
	return &AdministrationRepo{
		db: db,
	}
}

func (c *AdministrationRepo) Create(ctx context.Context, req *users_service.CreateAdministration) (resp *users_service.AdministrationPrimaryKey, err error) {
	id, err := helper.NewIncrementId(c.db, "id", "administrator", "A", 8)

	query := `INSERT INTO "administrator" (
				id,
				full_name,
				phone,
				password,
				login,
				salary,
				ielts_score,
				branch_id,
				role_id,
				updated_at
			) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, now())
		`

	_, err = c.db.Exec(ctx,
		query,
		id(),
		req.FullName,
		req.Phone,
		req.Password,
		req.Login,
		req.Salary,
		req.IeltsScore,
		req.BranchId,
		"4fc92739-e7ed-47be-9042-aedab5c660e6",
	)

	if err != nil {
		return nil, err
	}

	return &users_service.AdministrationPrimaryKey{Id: id()}, nil
}

func (c *AdministrationRepo) GetById(ctx context.Context, req *users_service.AdministrationPrimaryKey) (resp *users_service.Administration, err error) {

	query := `
		SELECT
			id,
			full_name,
			phone,
			password,
			login,
			salary,
			ielts_score,
			branch_id,
			role_id,
			created_at,
			updated_at
		FROM "administrator"
		WHERE id = $1
	`

	var (
		id          sql.NullString
		full_name   sql.NullString
		phone       sql.NullString
		password    sql.NullString
		login       sql.NullString
		salary      sql.NullString
		ielts_score sql.NullString
		branch_id   sql.NullString
		role_id     sql.NullString
		createdAt   sql.NullString
		updatedAt   sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&full_name,
		&phone,
		&password,
		&login,
		&salary,
		&ielts_score,
		&branch_id,
		&role_id,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return resp, err
	}

	resp = &users_service.Administration{
		Id:         id.String,
		FullName:   full_name.String,
		Phone:      phone.String,
		Password:   password.String,
		Login:      login.String,
		Salary:     salary.String,
		IeltsScore: ielts_score.String,
		BranchId:   branch_id.String,
		RoleId:     role_id.String,
		CreatedAt:  createdAt.String,
		UpdatedAt:  updatedAt.String,
	}
	return resp, nil
}

func (c *AdministrationRepo) GetList(ctx context.Context, req *users_service.GetListAdministrationRequest) (resp *users_service.GetListAdministrationResponse, err error) {

	resp = &users_service.GetListAdministrationResponse{}

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
			password,
			login,
			salary,
			ielts_score,
			branch_id,
			role_id,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "administrator"
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
			id          sql.NullString
			full_name   sql.NullString
			phone       sql.NullString
			password    sql.NullString
			login       sql.NullString
			salary      sql.NullString
			ielts_score sql.NullString
			branch_id   sql.NullString
			role_id     sql.NullString
			createdAt   sql.NullString
			updatedAt   sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&full_name,
			&phone,
			&password,
			&login,
			&salary,
			&ielts_score,
			&branch_id,
			&role_id,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.Administrations = append(resp.Administrations, &users_service.Administration{
			Id:         id.String,
			FullName:   full_name.String,
			Phone:      phone.String,
			Password:   password.String,
			Login:      login.String,
			Salary:     salary.String,
			IeltsScore: ielts_score.String,
			BranchId:   branch_id.String,
			RoleId:     role_id.String,
			CreatedAt:  createdAt.String,
			UpdatedAt:  updatedAt.String,
		})
	}

	return
}

func (c *AdministrationRepo) Update(ctx context.Context, req *users_service.UpdateAdministration) (rowsAffected int64, err error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
			UPDATE
			    "administrator"
			SET
				full_name = :full_name,
				phone = :phone,
				password = :password,
				login = :login,
				salary = :salary,
				ielts_score = :ielts_score,
				branch_id = :branch_id,
				role_id = :role_id,
				updated_at    = now()
			WHERE
				id = :id`
	params = map[string]interface{}{
		"id":          req.GetId(),
		"full_name":   req.GetFullName(),
		"phone":       req.GetPhone(),
		"password":    req.GetPassword(),
		"login":       req.GetLogin(),
		"salary":      req.GetSalary(),
		"ielts_score": req.GetIeltsScore(),
		"branch_id":   req.GetBranchId(),
		"role_id":     req.GetRoleId(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), nil
}

func (c *AdministrationRepo) Delete(ctx context.Context, req *users_service.AdministrationPrimaryKey) error {

	query := `DELETE FROM "administrator" WHERE id = $1`

	_, err := c.db.Exec(ctx, query, req.Id)

	if err != nil {
		return err
	}

	return nil
}
