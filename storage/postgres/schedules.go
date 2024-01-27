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

type ScheduleRepo struct {
	db *pgxpool.Pool
}

func NewScheduleRepo(db *pgxpool.Pool) storage.ScheduleRepoI {
	return &ScheduleRepo{
		db: db,
	}
}

func (c *ScheduleRepo) Create(ctx context.Context, req *users_service.CreateSchedule) (resp *users_service.SchedulePrimaryKey, err error) {
	var id = uuid.New()

	query := `INSERT INTO "schedule" (
				id,
				start_time,
				end_time,
				journal_id,
				updated_at
			) VALUES ($1, $2, $3, $4, now())
		`

	_, err = c.db.Exec(ctx,
		query,
		id.String(),
		req.StartTime,
		req.EndTime,
		req.JournalId,
	)

	if err != nil {
		return nil, err
	}

	return &users_service.SchedulePrimaryKey{Id: id.String()}, nil
}

func (c *ScheduleRepo) GetById(ctx context.Context, req *users_service.SchedulePrimaryKey) (resp *users_service.Schedule, err error) {

	query := `
		SELECT
			id,
			start_time,
			end_time,
			journal_id,
			created_at,
			updated_at
		FROM "schedule"
		WHERE id = $1
	`

	var (
		id         sql.NullString
		start_time sql.NullString
		end_time   sql.NullString
		journal_id sql.NullString
		createdAt  sql.NullString
		updatedAt  sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&start_time,
		&end_time,
		&journal_id,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return resp, err
	}

	resp = &users_service.Schedule{
		Id:        id.String,
		StartTime: start_time.String,
		EndTime:   end_time.String,
		JournalId: journal_id.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}
	return resp, nil
}

func (c *ScheduleRepo) GetList(ctx context.Context, req *users_service.GetListScheduleRequest) (resp *users_service.GetListScheduleResponse, err error) {

	resp = &users_service.GetListScheduleResponse{}

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
			start_time,
			end_time,
			journal_id,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "schedule"
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
			start_time sql.NullString
			end_time   sql.NullString
			journal_id sql.NullString
			createdAt  sql.NullString
			updatedAt  sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&start_time,
			&end_time,
			&journal_id,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.Schedules = append(resp.Schedules, &users_service.Schedule{
			Id:        id.String,
			StartTime: start_time.String,
			EndTime:   end_time.String,
			JournalId: journal_id.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return
}

func (c *ScheduleRepo) Update(ctx context.Context, req *users_service.UpdateSchedule) (rowsAffected int64, err error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
			UPDATE
			    "schedule"
			SET
				start_time = :start_time,
				end_time = :end_time,
				journal_id = :journal_id,
				updated_at    = now()
			WHERE
				id = :id`
	params = map[string]interface{}{
		"id":         req.GetId(),
		"start_time": req.GetStartTime(),
		"end_time":   req.GetEndTime(),
		"journal_id": req.GetJournalId(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), nil
}

func (c *ScheduleRepo) Delete(ctx context.Context, req *users_service.SchedulePrimaryKey) error {

	query := `DELETE FROM "schedule" WHERE id = $1`

	_, err := c.db.Exec(ctx, query, req.Id)

	if err != nil {
		return err
	}

	return nil
}
