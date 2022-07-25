package repo

import (
	"context"

	"github.com/erwinhermanto31/quiz_master/entity"
)

type IQuestions interface {
	FindAllQuestion(ctx context.Context, req entity.Question) (res []entity.Question, err error)
	FindQuestion(ctx context.Context, req entity.Question) (res entity.Question, err error)
	CreateQuestion(ctx context.Context, req entity.Question) (err error)
	UpdateQuestion(ctx context.Context, req entity.Question) (err error)
	DeleteQuestion(ctx context.Context, req entity.Question) (err error)
}
