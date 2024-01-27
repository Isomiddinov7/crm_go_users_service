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

type TaskRepo struct {
	db *pgxpool.Pool
}

func NewTaskRepo(db *pgxpool.Pool) storage.TaskRepoI {
	return &TaskRepo{
		db: db,
	}
}

func (c *TaskRepo) Create(ctx context.Context, req *users_service.CreateTask) (resp *users_service.TaskPrimaryKey, err error) {
	var id = uuid.New()

	query := `INSERT INTO "task" (
				id,
				lesson_id,
				label,
				deadline,
				score,
				updated_at
			) VALUES ($1, $2, $3, $4, $5, now())
		`

	_, err = c.db.Exec(ctx,
		query,
		id.String(),
		req.LessonId,
		req.Label,
		req.Deadline,
		req.Score,
	)

	if err != nil {
		return nil, err
	}

	return &users_service.TaskPrimaryKey{Id: id.String()}, nil
}

func (c *TaskRepo) GetById(ctx context.Context, req *users_service.TaskPrimaryKey) (resp *users_service.Task, err error) {

	query := `
		SELECT
			id,
			lesson_id,
			label,
			deadline,
			score,
			created_at,
			updated_at
		FROM "task"
		WHERE id = $1
	`

	var (
		id        sql.NullString
		lesson_id sql.NullString
		label     sql.NullString
		deadline  sql.NullString
		score     sql.NullInt64
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&lesson_id,
		&label,
		&deadline,
		&score,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return resp, err
	}

	resp = &users_service.Task{
		Id:        id.String,
		LessonId:  lesson_id.String,
		Label:     label.String,
		Deadline:  deadline.String,
		Score:     score.Int64,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}
	return resp, nil
}

func (c *TaskRepo) GetList(ctx context.Context, req *users_service.GetListTaskRequest) (resp *users_service.GetListTaskResponse, err error) {

	resp = &users_service.GetListTaskResponse{}

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
			lesson_id,
			label,
			deadline,
			score,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "task"
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
			lesson_id sql.NullString
			label     sql.NullString
			deadline  sql.NullString
			score     sql.NullInt64
			createdAt sql.NullString
			updatedAt sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&lesson_id,
			&label,
			&deadline,
			&score,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.Tasks = append(resp.Tasks, &users_service.Task{
			Id:        id.String,
			LessonId:  lesson_id.String,
			Label:     label.String,
			Deadline:  deadline.String,
			Score:     score.Int64,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return
}

func (c *TaskRepo) Update(ctx context.Context, req *users_service.UpdateTask) (rowsAffected int64, err error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
			UPDATE
			    "task"
			SET
				lesson_id = :lesson_id,
				label = :label,
				deadline = :deadline,
				score = :score,
				updated_at    = now()
			WHERE
				id = :id`
	params = map[string]interface{}{
		"id":        req.GetId(),
		"lesson_id": req.GetLessonId(),
		"label":     req.GetLabel(),
		"deadline":  req.GetDeadline(),
		"score":     req.GetScore(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), nil
}

func (c *TaskRepo) Delete(ctx context.Context, req *users_service.TaskPrimaryKey) error {

	query := `DELETE FROM "task" WHERE id = $1`

	_, err := c.db.Exec(ctx, query, req.Id)

	if err != nil {
		return err
	}

	return nil
}
