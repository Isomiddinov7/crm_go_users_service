package postgres

import (
	"context"
	"crm_go_users_service/genproto/users_service"
	"crm_go_users_service/pkg/helper"
	"crm_go_users_service/storage"
	"database/sql"

	"github.com/jackc/pgx/v4/pgxpool"
)

type StudentRepo struct {
	db *pgxpool.Pool
}

func NewStudentRepo(db *pgxpool.Pool) storage.StudentRepoI {
	return &StudentRepo{
		db: db,
	}
}

func (c *StudentRepo) Create(ctx context.Context, req *users_service.CreateStudent) (resp *users_service.StudentPrimaryKey, err error) {
	id, err := helper.NewIncrementId(c.db, "id", "student", "S", 8)

	query := `INSERT INTO "student" (
				id,
				full_name,
				phone,
				paid_sum,
				cource_count,
				total_sum,
				group_id,
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
		req.PaidSum,
		req.CourceCount,
		req.TotalSum,
		req.GroupId,
		req.BranchId,
		"a3eb6ca6-c1af-4b01-947b-83ab8bc943bb",
	)

	if err != nil {
		return nil, err
	}

	return &users_service.StudentPrimaryKey{Id: id()}, nil
}

func (c *StudentRepo) GetById(ctx context.Context, req *users_service.StudentPrimaryKey) (resp *users_service.Student, err error) {

	query := `
		SELECT
			id,
			full_name,
			phone,
			paid_sum,
			cource_count,
			total_sum,
			group_id,
			branch_id,
			role_id,
			created_at,
			updated_at
		FROM "student"
		WHERE id = $1
	`

	var (
		id           sql.NullString
		full_name    sql.NullString
		phone        sql.NullString
		paid_sum     sql.NullString
		cource_count sql.NullInt64
		total_sum    sql.NullString
		group_id     sql.NullString
		branch_id    sql.NullString
		role_id      sql.NullString
		createdAt    sql.NullString
		updatedAt    sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&full_name,
		&phone,
		&paid_sum,
		&cource_count,
		&total_sum,
		&group_id,
		&branch_id,
		&role_id,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return resp, err
	}

	resp = &users_service.Student{
		Id:          id.String,
		FullName:    full_name.String,
		Phone:       phone.String,
		PaidSum:     paid_sum.String,
		CourceCount: cource_count.Int64,
		TotalSum:    total_sum.String,
		GroupId:     group_id.String,
		BranchId:    branch_id.String,
		RoleId:      role_id.String,
		CreatedAt:   createdAt.String,
		UpdatedAt:   updatedAt.String,
	}
	return resp, nil
}

func (c *StudentRepo) GetList(ctx context.Context, req *users_service.GetListStudentRequest) (resp *users_service.GetListStudentResponse, err error) {

	resp = &users_service.GetListStudentResponse{}

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
			paid_sum,
			cource_count,
			total_sum,
			group_id,
			branch_id,
			role_id,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "student"
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
			id           sql.NullString
			full_name    sql.NullString
			phone        sql.NullString
			paid_sum     sql.NullString
			cource_count sql.NullInt64
			total_sum    sql.NullString
			group_id     sql.NullString
			branch_id    sql.NullString
			role_id      sql.NullString
			createdAt    sql.NullString
			updatedAt    sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&full_name,
			&phone,
			&paid_sum,
			&cource_count,
			&total_sum,
			&group_id,
			&branch_id,
			&role_id,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.Students = append(resp.Students, &users_service.Student{
			Id:          id.String,
			FullName:    full_name.String,
			Phone:       phone.String,
			PaidSum:     paid_sum.String,
			CourceCount: cource_count.Int64,
			TotalSum:    total_sum.String,
			GroupId:     group_id.String,
			BranchId:    branch_id.String,
			RoleId:      role_id.String,
			CreatedAt:   createdAt.String,
			UpdatedAt:   updatedAt.String,
		})
	}

	return
}

func (c *StudentRepo) Update(ctx context.Context, req *users_service.UpdateStudent) (rowsAffected int64, err error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
			UPDATE
			    "student"
			SET
				full_name	  = :full_name,
				phone		  = :phone,
				paid_sum 	  = :paid_sum,
				cource_count  = :cource_count,
				total_sum	  = :total_sum,
				group_id = :group_id,
				branch_id = :branch_id,
				role_id = :role_id,
				updated_at    = now()
			WHERE
				id = :id`
	params = map[string]interface{}{
		"id":           req.GetId(),
		"full_name":    req.GetFullName(),
		"phone":        req.GetPhone(),
		"paid_sum":     req.GetPaidSum(),
		"cource_count": req.GetCourceCount(),
		"total_sum":    req.GetTotalSum(),
		"group_id":     req.GetGroupId(),
		"branch_id":    req.GetBranchId(),
		"role_id":      req.GetRoleId(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), nil
}

func (c *StudentRepo) Delete(ctx context.Context, req *users_service.StudentPrimaryKey) error {

	query := `DELETE FROM "student" WHERE id = $1`

	_, err := c.db.Exec(ctx, query, req.Id)

	if err != nil {
		return err
	}

	return nil
}
