package activity_log_repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/models"
)

type ActivityLogRepositoryInterface interface {
	Create(ctx context.Context, tx *sqlx.Tx, data models.CreateActivityLog) *constants.ErrorResponse
}

func NewActivityLogRepository() ActivityLogRepositoryInterface {
	return &activityLogRepository{}
}
