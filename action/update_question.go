package action

import (
	"context"
	"log"

	"github.com/erwinhermanto31/quiz_master/entity"
	"github.com/erwinhermanto31/quiz_master/repo"
)

type UpdateQuestion struct {
	repoUsers repo.IQuestions
}

func NewUpdateQuestion() *UpdateQuestion {
	return &UpdateQuestion{
		repoUsers: repo.NewQuestion(),
	}
}

func (a *UpdateQuestion) Handler(ctx context.Context, req entity.Question) (err error) {
	err = a.repoUsers.UpdateQuestion(ctx, req)
	if err != nil {
		log.Printf("[Handler] UpdateQuestion : %v", err)
		return err
	}
	return nil
}
