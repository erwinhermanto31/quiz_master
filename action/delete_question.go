package action

import (
	"context"
	"log"

	"github.com/erwinhermanto31/quiz_master/entity"
	"github.com/erwinhermanto31/quiz_master/repo"
)

type DeleteQuestion struct {
	repoQuestion repo.IQuestions
}

func NewDeleteQuestion() *DeleteQuestion {
	return &DeleteQuestion{
		repoQuestion: repo.NewQuestion(),
	}
}

func (a *DeleteQuestion) Handler(ctx context.Context, req entity.Question) (err error) {
	err = a.repoQuestion.DeleteQuestion(ctx, req)
	if err != nil {
		log.Printf("[Handler] DeleteQuestion : %v", err)
		return err
	}
	return nil
}
