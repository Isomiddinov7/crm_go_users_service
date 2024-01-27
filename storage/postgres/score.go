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

type ScoreRepo struct {
	db *pgxpool.Pool
}

func NewScoreRepo(db *pgxpool.Pool) storage.ScoreRepoI {
	return &ScoreRepo{
		db: db,
	}
}

func (c *ScoreRepo) Create(ctx context.Context, req *users_service.CreateScore) (resp *users_service.ScorePrimaryKey, err error) {
	var id = uuid.New()

	query := `INSERT INTO "score" (
				id,
				task_id,
				student_id,
				updated_at
			) VALUES ($1, $2, $3, now())
		`

	_, err = c.db.Exec(ctx,
		query,
		id.String(),
		req.TaskId,
		req.StudentId,
	)

	if err != nil {
		return nil, err
	}

	return &users_service.ScorePrimaryKey{Id: id.String()}, nil
}

func (c *ScoreRepo) GetById(ctx context.Context, req *users_service.ScorePrimaryKey) (resp *users_service.Score, err error) {

	query := `
		SELECT
			id,
			task_id,
			student_id,
			created_at,
			updated_at
		FROM "score"
		WHERE id = $1
	`

	var (
		id         sql.NullString
		task_id    sql.NullString
		student_id sql.NullString
		createdAt  sql.NullString
		updatedAt  sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&task_id,
		&student_id,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return resp, err
	}

	resp = &users_service.Score{
		Id:        id.String,
		TaskId:    task_id.String,
		StudentId: student_id.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}
	return resp, nil
}

func (c *ScoreRepo) GetList(ctx context.Context, req *users_service.GetListScoreRequest) (resp *users_service.GetListScoreResponse, err error) {

	resp = &users_service.GetListScoreResponse{}

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
			task_id,
			student_id,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "score"
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
			id         sql.NullString
			task_id    sql.NullString
			student_id sql.NullString
			createdAt  sql.NullString
			updatedAt  sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&task_id,
			&student_id,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.Scores = append(resp.Scores, &users_service.Score{
			Id:        id.String,
			TaskId:    task_id.String,
			StudentId: student_id.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return
}

func (c *ScoreRepo) Update(ctx context.Context, req *users_service.UpdateScore) (rowsAffected int64, err error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
			UPDATE
			    "score"
			SET
				task_id = :task_id,
				student_id = :student_id,
				updated_at    = now()
			WHERE
				id = :id`
	params = map[string]interface{}{
		"id":         req.GetId(),
		"task_id":    req.GetTaskId(),
		"student_id": req.GetStudentId(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), nil
}

func (c *ScoreRepo) Delete(ctx context.Context, req *users_service.ScorePrimaryKey) error {

	query := `DELETE FROM "score" WHERE id = $1`

	_, err := c.db.Exec(ctx, query, req.Id)

	if err != nil {
		return err
	}

	return nil
}
