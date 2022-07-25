package action

import (
	"context"
	"log"

	"github.com/erwinhermanto31/quiz_master/entity"
	"github.com/erwinhermanto31/quiz_master/repo"
)

type GetQuestion struct {
	repoUsers repo.IQuestions
}

func NewGetQuestion() *GetQuestion {
	return &GetQuestion{
		repoUsers: repo.NewQuestion(),
	}
}

func (a *GetQuestion) Handler(ctx context.Context, req entity.Question) (res entity.Question, err error) {
	data, err := a.repoUsers.FindQuestion(ctx, req)
	if err != nil {
		log.Printf("[Handler] FindQuestion : %v", err)
		return res, err
	}

	return data, nil
}
