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

type JournalRepo struct {
	db *pgxpool.Pool
}

func NewJournalRepo(db *pgxpool.Pool) storage.JournalRepoI {
	return &JournalRepo{
		db: db,
	}
}

func (c *JournalRepo) Create(ctx context.Context, req *users_service.CreateJournal) (resp *users_service.JournalPrimaryKey, err error) {
	var id = uuid.New()

	query := `INSERT INTO "journal" (
				id,
				group_id,
				"from",
				"to",
				student_id,
				updated_at
			) VALUES ($1, $2, $3, $4, $5, now())
		`

	_, err = c.db.Exec(ctx,
		query,
		id.String(),
		req.GroupId,
		req.From,
		req.To,
		req.StudentId,
	)

	if err != nil {
		return nil, err
	}

	return &users_service.JournalPrimaryKey{Id: id.String()}, nil
}

func (c *JournalRepo) GetById(ctx context.Context, req *users_service.JournalPrimaryKey) (resp *users_service.Journal, err error) {

	query := `
		SELECT
			id,
			group_id,
			"from",
			"to",
			student_id,
			created_at,
			updated_at
		FROM "journal"
		WHERE id = $1
	`

	var (
		id         sql.NullString
		group_id   sql.NullString
		from       sql.NullString
		to         sql.NullString
		student_id sql.NullString
		createdAt  sql.NullString
		updatedAt  sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&group_id,
		&from,
		&to,
		&student_id,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return resp, err
	}

	resp = &users_service.Journal{
		Id:        id.String,
		GroupId:   group_id.String,
		From:      from.String,
		To:        to.String,
		StudentId: student_id.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}
	return resp, nil
}

func (c *JournalRepo) GetList(ctx context.Context, req *users_service.GetListJournalRequest) (resp *users_service.GetListJournalResponse, err error) {

	resp = &users_service.GetListJournalResponse{}

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
			group_id,
			"from",
			"to",
			student_id,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "journal"
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
			group_id   sql.NullString
			from       sql.NullString
			to         sql.NullString
			student_id sql.NullString
			createdAt  sql.NullString
			updatedAt  sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&group_id,
			&from,
			&to,
			&student_id,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.Journals = append(resp.Journals, &users_service.Journal{
			Id:        id.String,
			GroupId:   group_id.String,
			From:      from.String,
			To:        to.String,
			StudentId: student_id.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return
}

func (c *JournalRepo) Update(ctx context.Context, req *users_service.UpdateJournal) (rowsAffected int64, err error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
			UPDATE
			    "journal"
			SET
				group_id = :group_id,
				"from" = :"from",
				"to" = :"to",
				student_id = :student_id,
				updated_at    = now()
			WHERE
				id = :id`
	params = map[string]interface{}{
		"id":         req.GetId(),
		"group_id":   req.GetGroupId(),
		"from":       req.GetFrom(),
		"to":         req.GetTo(),
		"student_id": req.GetStudentId(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), nil
}

func (c *JournalRepo) Delete(ctx context.Context, req *users_service.JournalPrimaryKey) error {

	query := `DELETE FROM "journal" WHERE id = $1`

	_, err := c.db.Exec(ctx, query, req.Id)

	if err != nil {
		return err
	}

	return nil
}
