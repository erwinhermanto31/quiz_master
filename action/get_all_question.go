package action

import (
	"context"
	"log"

	"github.com/erwinhermanto31/quiz_master/entity"
	"github.com/erwinhermanto31/quiz_master/repo"
)

type GetAllQuestion struct {
	repoQuestion repo.IQuestions
}

func NewGetAllQuestion() *GetAllQuestion {
	return &GetAllQuestion{
		repoQuestion: repo.NewQuestion(),
	}
}

func (a *GetAllQuestion) Handler(ctx context.Context, req entity.Question) (res []entity.Question, err error) {
	data, err := a.repoQuestion.FindAllQuestion(ctx, req)
	if err != nil {
		log.Printf("[Handler] FindAllQuestion : %v", err)
		return res, err
	}

	return data, nil
}
