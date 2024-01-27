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

type AssignStudentRepo struct {
	db *pgxpool.Pool
}

func NewAssignStudentRepo(db *pgxpool.Pool) storage.AssignStudentRepoI {
	return &AssignStudentRepo{
		db: db,
	}
}

func (c *AssignStudentRepo) Create(ctx context.Context, req *users_service.CreateAssignStudent) (resp *users_service.AssignStudentPrimaryKey, err error) {
	var id = uuid.New()

	query := `INSERT INTO "assign_student" (
				id,
				event_id,
				student_id,
				updated_at
			) VALUES ($1, $2, $3, now())
		`

	_, err = c.db.Exec(ctx,
		query,
		id.String(),
		req.EventId,
		req.StudentId,
	)

	if err != nil {
		return nil, err
	}

	return &users_service.AssignStudentPrimaryKey{Id: id.String()}, nil
}

func (c *AssignStudentRepo) GetById(ctx context.Context, req *users_service.AssignStudentPrimaryKey) (resp *users_service.AssignStudent, err error) {

	query := `
		SELECT
			id,
			event_id,
			student_id,
			updated_at
		FROM "assign_student"
		WHERE id = $1
	`

	var (
		id         sql.NullString
		event_id   sql.NullString
		student_id sql.NullString
		createdAt  sql.NullString
		updatedAt  sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&event_id,
		&student_id,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return resp, err
	}

	resp = &users_service.AssignStudent{
		Id:        id.String,
		EventId:   event_id.String,
		StudentId: student_id.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}
	return resp, nil
}

func (c *AssignStudentRepo) GetList(ctx context.Context, req *users_service.GetListAssignStudentRequest) (resp *users_service.GetListAssignStudentResponse, err error) {

	resp = &users_service.GetListAssignStudentResponse{}

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
			event_id,
			student_id,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "assign_student"
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
			event_id   sql.NullString
			student_id sql.NullString
			createdAt  sql.NullString
			updatedAt  sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&event_id,
			&student_id,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.AssignStudents = append(resp.AssignStudents, &users_service.AssignStudent{
			Id:        id.String,
			EventId:   event_id.String,
			StudentId: student_id.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return
}

func (c *AssignStudentRepo) Update(ctx context.Context, req *users_service.UpdateAssignStudent) (rowsAffected int64, err error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
			UPDATE
			    "assign_student"
			SET
				event_id = :event_id,
				student_id = :student_id,
				updated_at    = now()
			WHERE
				id = :id`
	params = map[string]interface{}{
		"id":         req.GetId(),
		"event_id":   req.GetEventId(),
		"student_id": req.GetStudentId(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), nil
}

func (c *AssignStudentRepo) Delete(ctx context.Context, req *users_service.AssignStudentPrimaryKey) error {

	query := `DELETE FROM "assign_student" WHERE id = $1`

	_, err := c.db.Exec(ctx, query, req.Id)

	if err != nil {
		return err
	}

	return nil
}
