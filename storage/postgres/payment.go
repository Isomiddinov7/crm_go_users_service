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

type PaymentRepo struct {
	db *pgxpool.Pool
}

func NewPaymentRepo(db *pgxpool.Pool) storage.PaymentRepoI {
	return &PaymentRepo{
		db: db,
	}
}

func (c *PaymentRepo) Create(ctx context.Context, req *users_service.CreatePayment) (resp *users_service.PaymentPrimaryKey, err error) {
	var id = uuid.New()

	query := `INSERT INTO "payment" (
				id,
				student_id,
				branch_id,
				paid_sum,
				total_sum,
				course_count,
				updated_at
			) VALUES ($1, $2, $3, $4, $5, $6, now())
		`

	_, err = c.db.Exec(ctx,
		query,
		id.String(),
		req.StudentId,
		req.BranchId,
		req.PaidSum,
		req.TotalSum,
		req.CourseCount,
	)

	if err != nil {
		return nil, err
	}

	return &users_service.PaymentPrimaryKey{Id: id.String()}, nil
}

func (c *PaymentRepo) GetById(ctx context.Context, req *users_service.PaymentPrimaryKey) (resp *users_service.Payment, err error) {

	query := `
		SELECT
			id,
			student_id,
			branch_id,
			paid_sum,
			total_sum,
			course_count,
			created_at,
			updated_at
		FROM "payment"
		WHERE id = $1
	`

	var (
		id           sql.NullString
		student_id   sql.NullString
		branch_id    sql.NullString
		paid_sum     sql.NullString
		total_sum    sql.NullString
		course_count sql.NullString
		createdAt    sql.NullString
		updatedAt    sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&student_id,
		&branch_id,
		&paid_sum,
		&total_sum,
		&course_count,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return resp, err
	}

	resp = &users_service.Payment{
		Id:          id.String,
		StudentId:   student_id.String,
		BranchId:    branch_id.String,
		PaidSum:     paid_sum.String,
		TotalSum:    total_sum.String,
		CourseCount: course_count.String,
		CreatedAt:   createdAt.String,
		UpdatedAt:   updatedAt.String,
	}
	return resp, nil
}

func (c *PaymentRepo) GetList(ctx context.Context, req *users_service.GetListPaymentRequest) (resp *users_service.GetListPaymentResponse, err error) {

	resp = &users_service.GetListPaymentResponse{}

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
			student_id,
			branch_id,
			paid_sum,
			total_sum,
			course_count,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "payment"
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
			student_id   sql.NullString
			branch_id    sql.NullString
			paid_sum     sql.NullString
			total_sum    sql.NullString
			course_count sql.NullString
			createdAt    sql.NullString
			updatedAt    sql.NullString
		)
		err := rows.Scan(
			&resp.Count,
			&id,
			&student_id,
			&branch_id,
			&paid_sum,
			&total_sum,
			&course_count,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.Payments = append(resp.Payments, &users_service.Payment{
			Id:          id.String,
			StudentId:   student_id.String,
			BranchId:    branch_id.String,
			PaidSum:     paid_sum.String,
			TotalSum:    total_sum.String,
			CourseCount: course_count.String,
			CreatedAt:   createdAt.String,
			UpdatedAt:   updatedAt.String,
		})
	}

	return
}

func (c *PaymentRepo) Update(ctx context.Context, req *users_service.UpdatePayment) (rowsAffected int64, err error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
			UPDATE
			    "payment"
			SET
				student_id = :student_id,
				branch_id = :branch_id,
				paid_sum = :paid_sum,
				total_sum = :total_sum,
				course_count = :course_count,
				updated_at    = now()
			WHERE
				id = :id`
	params = map[string]interface{}{
		"id":           req.GetId(),
		"student_id":   req.GetStudentId(),
		"branch_id":    req.GetBranchId(),
		"paid_sum":     req.GetPaidSum(),
		"total_sum":    req.GetTotalSum(),
		"course_count": req.GetCourseCount(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), nil
}

func (c *PaymentRepo) Delete(ctx context.Context, req *users_service.PaymentPrimaryKey) error {

	query := `DELETE FROM "payment" WHERE id = $1`

	_, err := c.db.Exec(ctx, query, req.Id)

	if err != nil {
		return err
	}

	return nil
}
