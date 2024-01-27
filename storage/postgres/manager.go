package postgres

import (
	"context"
	"crm_go_users_service/genproto/users_service"
	"crm_go_users_service/pkg/helper"
	"crm_go_users_service/storage"
	"database/sql"

	"github.com/jackc/pgx/v4/pgxpool"
)

type ManagerRepo struct {
	db *pgxpool.Pool
}

func NewManagerRepo(db *pgxpool.Pool) storage.ManagerRepoI {
	return &ManagerRepo{
		db: db,
	}
}

func (c *ManagerRepo) Create(ctx context.Context, req *users_service.CreateManager) (resp *users_service.ManagerPrimaryKey, err error) {
	id, err := helper.NewIncrementId(c.db, "id", "manager", "M", 8)

	query := `INSERT INTO "manager" (
				id,
				full_name,
				phone,
				salary,
				password,
				login,
				branch_id,
				role_id,
				updated_at
			) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, now())
		`

	_, err = c.db.Exec(ctx,
		query,
		id(),
		req.FullName,
		req.Phone,
		req.Saley,
		req.Password,
		req.Login,
		req.BranchId,
		"6eafc237-bc61-4d48-a884-94365dc58dcb",
	)

	if err != nil {
		return nil, err
	}

	return &users_service.ManagerPrimaryKey{Id: id()}, nil
}

func (c *ManagerRepo) GetById(ctx context.Context, req *users_service.ManagerPrimaryKey) (resp *users_service.Manager, err error) {

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
			created_at,
			updated_at
		FROM "manager"
		WHERE id = $1
	`

	var (
		id        sql.NullString
		full_name sql.NullString
		phone     sql.NullString
		salary    sql.NullString
		password  sql.NullString
		login     sql.NullString
		branch_id sql.NullString
		role_id   sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
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
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return resp, err
	}

	resp = &users_service.Manager{
		Id:        id.String,
		FullName:  full_name.String,
		Phone:     phone.String,
		Saley:     salary.String,
		Password:  password.String,
		Login:     login.String,
		BranchId:  branch_id.String,
		RoleId:    role_id.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}
	return resp, nil
}

func (c *ManagerRepo) GetList(ctx context.Context, req *users_service.GetListManagerRequest) (resp *users_service.GetListManagerResponse, err error) {

	resp = &users_service.GetListManagerResponse{}

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
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "manager"
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
			id        sql.NullString
			full_name sql.NullString
			phone     sql.NullString
			salary    sql.NullString
			password  sql.NullString
			login     sql.NullString
			branch_id sql.NullString
			role_id   sql.NullString
			createdAt sql.NullString
			updatedAt sql.NullString
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
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.Managers = append(resp.Managers, &users_service.Manager{
			Id:        id.String,
			FullName:  full_name.String,
			Phone:     phone.String,
			Saley:     salary.String,
			Password:  password.String,
			Login:     login.String,
			BranchId:  branch_id.String,
			RoleId:    role_id.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return
}

func (c *ManagerRepo) Update(ctx context.Context, req *users_service.UpdateManager) (rowsAffected int64, err error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
			UPDATE
			    "manager"
			SET
				full_name = :full_name,
				phone = :phone,
				salary = :salary,
				password = :password,
				login = :login,
				branch_id = :branch_id,
				role_id = :role_id,
				updated_at    = now()
			WHERE
				id = :id`
	params = map[string]interface{}{
		"id":        req.GetId(),
		"full_name": req.GetFullName(),
		"phone":     req.GetPhone(),
		"salary":    req.GetSaley(),
		"password":  req.GetPassword(),
		"login":     req.GetLogin(),
		"branch_id": req.GetBranchId(),
		"role_id":   req.GetRoleId(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), nil
}

func (c *ManagerRepo) Delete(ctx context.Context, req *users_service.ManagerPrimaryKey) error {

	query := `DELETE FROM "manager" WHERE id = $1`

	_, err := c.db.Exec(ctx, query, req.Id)

	if err != nil {
		return err
	}

	return nil
}
