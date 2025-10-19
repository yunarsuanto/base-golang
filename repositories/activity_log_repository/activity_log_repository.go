package activity_log_repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/yunarsuanto/base-go/constants"
	"github.com/yunarsuanto/base-go/models"
	"github.com/yunarsuanto/base-go/utils"
)

type activityLogRepository struct{}

func (activityLogRepository) Create(ctx context.Context, tx *sqlx.Tx, data models.CreateActivityLog) *constants.ErrorResponse {
	_, err := tx.NamedExecContext(
		ctx,
		createQuery,
		data,
	)
	if err != nil {
		return utils.ErrDatabase(err, models.ActivityLogDataName)
	}

	return nil
}
