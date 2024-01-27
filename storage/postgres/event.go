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

type EventRepo struct {
	db *pgxpool.Pool
}

func NewEventRepo(db *pgxpool.Pool) storage.EventRepoI {
	return &EventRepo{
		db: db,
	}
}

func (c *EventRepo) Create(ctx context.Context, req *users_service.CreateEvent) (resp *users_service.EventPrimaryKey, err error) {
	var id = uuid.New()

	query := `INSERT INTO "event" (
				id,
				topic,
				start_time,
				day,
				branch_id,
				updated_at
			) VALUES ($1, $2, $3, $4, $5, now())
		`

	_, err = c.db.Exec(ctx,
		query,
		id.String(),
		req.Topic,
		req.StartTime,
		req.Day,
		req.BranchId,
	)

	if err != nil {
		return nil, err
	}

	return &users_service.EventPrimaryKey{Id: id.String()}, nil
}

func (c *EventRepo) GetById(ctx context.Context, req *users_service.EventPrimaryKey) (resp *users_service.Event, err error) {

	query := `
		SELECT
			id,
			topic,
			start_time,
			day,
			branch_id,
			created_at,
			updated_at
		FROM "event"
		WHERE id = $1
	`

	var (
		id         sql.NullString
		topic      sql.NullString
		start_time sql.NullString
		day        sql.NullString
		branch_id  sql.NullString
		createdAt  sql.NullString
		updatedAt  sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&topic,
		&start_time,
		&day,
		&branch_id,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return resp, err
	}

	resp = &users_service.Event{
		Id:        id.String,
		Topic:     topic.String,
		StartTime: start_time.String,
		Day:       day.String,
		BranchId:  branch_id.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}
	return resp, nil
}

func (c *EventRepo) GetList(ctx context.Context, req *users_service.GetListEventRequest) (resp *users_service.GetListEventResponse, err error) {

	resp = &users_service.GetListEventResponse{}

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
			topic,
			start_time,
			day,
			branch_id,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "event"
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
			topic      sql.NullString
			start_time sql.NullString
			day        sql.NullString
			branch_id  sql.NullString
			createdAt  sql.NullString
			updatedAt  sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&topic,
			&start_time,
			&day,
			&branch_id,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.Events = append(resp.Events, &users_service.Event{
			Id:        id.String,
			Topic:     topic.String,
			StartTime: start_time.String,
			Day:       day.String,
			BranchId:  branch_id.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return
}

func (c *EventRepo) Update(ctx context.Context, req *users_service.UpdateEvent) (rowsAffected int64, err error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
			UPDATE
			    "event"
			SET
				topic = :topic,
				start_time = :start_time,
				day = :day,
				branch_id = :branch_id,
				updated_at    = now()
			WHERE
				id = :id`
	params = map[string]interface{}{
		"id":         req.GetId(),
		"topic":      req.GetTopic(),
		"start_time": req.GetStartTime(),
		"day":        req.GetDay(),
		"branch_id":  req.GetBranchId(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), nil
}

func (c *EventRepo) Delete(ctx context.Context, req *users_service.EventPrimaryKey) error {

	query := `DELETE FROM "event" WHERE id = $1`

	_, err := c.db.Exec(ctx, query, req.Id)

	if err != nil {
		return err
	}

	return nil
}
