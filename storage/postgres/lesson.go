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

type LessonRepo struct {
	db *pgxpool.Pool
}

func NewLessonRepo(db *pgxpool.Pool) storage.LessonRepoI {
	return &LessonRepo{
		db: db,
	}
}

func (c *LessonRepo) Create(ctx context.Context, req *users_service.CreateLesson) (resp *users_service.LessonPrimaryKey, err error) {
	var id = uuid.New()

	query := `INSERT INTO "lesson" (
				id,
				schedule_id,
				lesson,
				updated_at
			) VALUES ($1, $2, $3, now())
		`

	_, err = c.db.Exec(ctx,
		query,
		id.String(),
		req.ScheduleId,
		req.Lesson,
	)

	if err != nil {
		return nil, err
	}

	return &users_service.LessonPrimaryKey{Id: id.String()}, nil
}

func (c *LessonRepo) GetById(ctx context.Context, req *users_service.LessonPrimaryKey) (resp *users_service.Lesson, err error) {

	query := `
		SELECT
			id,
			schedule_id,
			lesson,
			created_at,
			updated_at
		FROM "lesson"
		WHERE id = $1
	`

	var (
		id          sql.NullString
		schedule_id sql.NullString
		lesson      sql.NullString
		createdAt   sql.NullString
		updatedAt   sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&schedule_id,
		&lesson,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return resp, err
	}

	resp = &users_service.Lesson{
		Id:         id.String,
		ScheduleId: schedule_id.String,
		Lesson:     lesson.String,
		CreatedAt:  createdAt.String,
		UpdatedAt:  updatedAt.String,
	}
	return resp, nil
}

func (c *LessonRepo) GetList(ctx context.Context, req *users_service.GetListLessonRequest) (resp *users_service.GetListLessonResponse, err error) {

	resp = &users_service.GetListLessonResponse{}

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
			schedule_id,
			lesson,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "lesson"
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
			schedule_id sql.NullString
			lesson      sql.NullString
			createdAt   sql.NullString
			updatedAt   sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&schedule_id,
			&lesson,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.Lessons = append(resp.Lessons, &users_service.Lesson{
			Id:         id.String,
			ScheduleId: schedule_id.String,
			Lesson:     lesson.String,
			CreatedAt:  createdAt.String,
			UpdatedAt:  updatedAt.String,
		})
	}

	return
}

func (c *LessonRepo) Update(ctx context.Context, req *users_service.UpdateLesson) (rowsAffected int64, err error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
			UPDATE
			    "lesson"
			SET
				schedule_id = :schedule_id,
				lesson = :lesson,
				updated_at    = now()
			WHERE
				id = :id`
	params = map[string]interface{}{
		"id":          req.GetId(),
		"schedule_id": req.GetScheduleId(),
		"lesson":      req.GetLesson(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), nil
}

func (c *LessonRepo) Delete(ctx context.Context, req *users_service.LessonPrimaryKey) error {

	query := `DELETE FROM "lesson" WHERE id = $1`

	_, err := c.db.Exec(ctx, query, req.Id)

	if err != nil {
		return err
	}

	return nil
}
