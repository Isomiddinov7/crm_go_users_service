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

type BranchRepo struct {
	db *pgxpool.Pool
}

func NewBranchRepo(db *pgxpool.Pool) storage.BranchRepoI {
	return &BranchRepo{
		db: db,
	}
}

func (c *BranchRepo) Create(ctx context.Context, req *users_service.CreateBranch) (resp *users_service.BranchPrimaryKey, err error) {
	var id = uuid.New()

	query := `INSERT INTO "branch" (
				id,
				name,
				updated_at
			) VALUES ($1, $2, now())
		`

	_, err = c.db.Exec(ctx,
		query,
		id.String(),
		req.Name,
	)

	if err != nil {
		return nil, err
	}

	return &users_service.BranchPrimaryKey{Id: id.String()}, nil
}

func (c *BranchRepo) GetById(ctx context.Context, req *users_service.BranchPrimaryKey) (resp *users_service.Branch, err error) {

	query := `
		SELECT
			id,
			name,
			created_at,
			updated_at
		FROM "branch"
		WHERE id = $1
	`

	var (
		id                 sql.NullString
		name               sql.NullString
		createdAt          sql.NullString
		updatedAt          sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return resp, err
	}

	resp = &users_service.Branch{
		Id:               id.String,
		Name:             name.String,
		CreatedAt:        createdAt.String,
		UpdatedAt:        updatedAt.String,
	}
	return resp, nil
}

func (c *BranchRepo) GetList(ctx context.Context, req *users_service.GetListBranchRequest) (resp *users_service.GetListBranchResponse, err error) {

	resp = &users_service.GetListBranchResponse{}

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
			name,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "branch"
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
			id                 sql.NullString
			name               sql.NullString
			createdAt          sql.NullString
			updatedAt          sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&name,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.Branchs = append(resp.Branchs, &users_service.Branch{
			Id:               id.String,
			Name:             name.String,
			CreatedAt:        createdAt.String,
			UpdatedAt:        updatedAt.String,
		})
	}

	return
}

func (c *BranchRepo) Update(ctx context.Context, req *users_service.UpdateBranch) (rowsAffected int64, err error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
			UPDATE
			    "branch"
			SET
				name = :name,
				updated_at    = now()
			WHERE
				id = :id`
	params = map[string]interface{}{
		"id":                 req.GetId(),
		"name":               req.GetName(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), nil
}

func (c *BranchRepo) Delete(ctx context.Context, req *users_service.BranchPrimaryKey) error {

	query := `DELETE FROM "branch" WHERE id = $1`

	_, err := c.db.Exec(ctx, query, req.Id)

	if err != nil {
		return err
	}

	return nil
}
