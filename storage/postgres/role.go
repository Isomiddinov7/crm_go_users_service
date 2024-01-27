package postgres

import (
	"context"
	"crm_go_users_service/genproto/users_service"
	"crm_go_users_service/pkg/helper"
	"crm_go_users_service/storage"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type RoleRepo struct {
	db *pgxpool.Pool
}

func NewRoleRepo(db *pgxpool.Pool) storage.RoleRepoI {
	return &RoleRepo{
		db: db,
	}
}

func (c *RoleRepo) Create(ctx context.Context, req *users_service.CreateRole) (resp *users_service.RolePrimaryKey, err error) {
	var id = uuid.New()

	query := `INSERT INTO "role" (
				id,
				role
			) VALUES ($1, $2)
		`

	_, err = c.db.Exec(ctx,
		query,
		id.String(),
		req.Role,
	)

	if err != nil {
		return nil, err
	}

	return &users_service.RolePrimaryKey{Id: id.String()}, nil
}

func (c *RoleRepo) GetById(ctx context.Context, req *users_service.RolePrimaryKey) (resp *users_service.Role, err error) {

	query := `
		SELECT
			id,
			role
		FROM "role"
		WHERE id = $1
	`

	var (
		id        sql.NullString
		role      sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&role,
	)
	if err != nil {
		return resp, err
	}

	resp = &users_service.Role{
		Id:        id.String,
		Role:      role.String,
	}
	return resp, nil
}

func (c *RoleRepo) GetList(ctx context.Context, req *users_service.GetListRoleRequest) (resp *users_service.GetListRoleResponse, err error) {

	resp = &users_service.GetListRoleResponse{}

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
			role
		FROM "role"
	`

	if req.Limit > 0 {
		limit = " LIMIT :limit"
		params["limit"] = req.Limit
	}

	if req.Offset > 0 {
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
			role      sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&role,
		)

		if err != nil {
			return resp, err
		}

		resp.Roles = append(resp.Roles, &users_service.Role{
			Id:        id.String,
			Role:      role.String,
		})
	}

	return
}

func (c *RoleRepo) Update(ctx context.Context, req *users_service.UpdateRole) (rowsAffected int64, err error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
			UPDATE
			    "role"
			SET
				role = :role
			WHERE
				id = :id`
	params = map[string]interface{}{
		"id":   req.Id,
		"role": req.Role,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), nil
}

func (c *RoleRepo) Delete(ctx context.Context, req *users_service.RolePrimaryKey) error {

	query := `DELETE FROM "role" WHERE id = $1`

	_, err := c.db.Exec(ctx, query, req.Id)

	if err != nil {
		return err
	}

	return nil
}
