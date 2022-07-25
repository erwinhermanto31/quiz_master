package mysql

import (
	"context"

	"github.com/erwinhermanto31/quiz_master/entity"
	"github.com/erwinhermanto31/quiz_master/util"
	"github.com/jmoiron/sqlx"
)

type IMysql interface {
	Get(context.Context, *sqlx.DB, interface{}, *util.Query, string) error
	Select(context.Context, *sqlx.DB, *[]entity.Question, *util.Query, string) error
	CreateOrUpdate(ctx context.Context, db *sqlx.DB, data interface{}, query string) (lastId int64, err error)
}
