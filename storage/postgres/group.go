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

type GroupRepo struct {
	db *pgxpool.Pool
}

func NewGroupRepo(db *pgxpool.Pool) storage.GroupRepoI {
	return &GroupRepo{
		db: db,
	}
}

func (c *GroupRepo) Create(ctx context.Context, req *users_service.CreateGroup) (resp *users_service.GroupPrimaryKey, err error) {
	var id = uuid.New()

	query := `INSERT INTO "group" (
				id,
				"uniqueID",
				branch_id,
				type,
				teacher_id,
				support_teacher_id,
				updated_at
			) VALUES ($1, $2, $3, $4, $5, $6, now())
		`

	_, err = c.db.Exec(ctx,
		query,
		id.String(),
		req.UniqueID,
		req.BranchId,
		req.Type,
		req.TeacherId,
		req.SupportTeacherId,
	)

	if err != nil {
		return nil, err
	}

	return &users_service.GroupPrimaryKey{Id: id.String()}, nil
}

func (c *GroupRepo) GetById(ctx context.Context, req *users_service.GroupPrimaryKey) (resp *users_service.Group, err error) {

	query := `
		SELECT
			id,
			"uniqueID",
			branch_id,
			type,
			teacher_id,
			support_teacher_id,
			created_at,
			updated_at
		FROM "group"
		WHERE id = $1
	`

	var (
		id                 sql.NullString
		uniqueID           sql.NullString
		branch_id          sql.NullString
		types              sql.NullString
		teacher_id         sql.NullString
		support_teacher_id sql.NullString
		createdAt          sql.NullString
		updatedAt          sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&uniqueID,
		&branch_id,
		&types,
		&teacher_id,
		&support_teacher_id,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return resp, err
	}

	resp = &users_service.Group{
		Id:               id.String,
		UniqueID:         uniqueID.String,
		BranchId:         branch_id.String,
		Type:             types.String,
		TeacherId:        teacher_id.String,
		SupportTeacherId: support_teacher_id.String,
		CreatedAt:        createdAt.String,
		UpdatedAt:        updatedAt.String,
	}
	return resp, nil
}

func (c *GroupRepo) GetList(ctx context.Context, req *users_service.GetListGroupRequest) (resp *users_service.GetListGroupResponse, err error) {

	resp = &users_service.GetListGroupResponse{}

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
			"uniqueID",
			branch_id,
			type,
			teacher_id,
			support_teacher_id,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "group"
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
			uniqueID           sql.NullString
			teacher_id         sql.NullString
			types              sql.NullString
			branch_id          sql.NullString
			support_teacher_id sql.NullString
			createdAt          sql.NullString
			updatedAt          sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&uniqueID,
			&branch_id,
			&types,
			&teacher_id,
			&support_teacher_id,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.Groups = append(resp.Groups, &users_service.Group{
			Id:               id.String,
			UniqueID:         uniqueID.String,
			BranchId:         branch_id.String,
			Type:             types.String,
			TeacherId:        teacher_id.String,
			SupportTeacherId: support_teacher_id.String,
			CreatedAt:        createdAt.String,
			UpdatedAt:        updatedAt.String,
		})
	}

	return
}

func (c *GroupRepo) Update(ctx context.Context, req *users_service.UpdateGroup) (rowsAffected int64, err error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
			UPDATE
			    "group"
			SET
				"uniqueID" = :"uniqueID",
				branch_id = :branch_id,
				type = :type,
				teacher_id = :teacher_id,
				support_teacher_id = :support_teacher_id,
				updated_at    = now()
			WHERE
				id = :id`
	params = map[string]interface{}{
		"id":                 req.GetId(),
		"uniqueID":           req.GetUniqueID(),
		"branch_id":          req.GetBranchId(),
		"type":               req.GetType(),
		"teacher_id":         req.GetTeacherId(),
		"support_teacher_id": req.GetSupportTeacherId(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), nil
}

func (c *GroupRepo) Delete(ctx context.Context, req *users_service.GroupPrimaryKey) error {

	query := `DELETE FROM "group" WHERE id = $1`

	_, err := c.db.Exec(ctx, query, req.Id)

	if err != nil {
		return err
	}

	return nil
}
