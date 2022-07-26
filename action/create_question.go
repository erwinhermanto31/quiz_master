package action

import (
	"context"
	"database/sql"
	stringErr "errors"
	"log"
	"strconv"

	"github.com/erwinhermanto31/quiz_master/entity"
	"github.com/erwinhermanto31/quiz_master/repo"
)

type CreateQuestion struct {
	repoQuestion repo.IQuestions
}

func NewCreateQuestion() *CreateQuestion {
	return &CreateQuestion{
		repoQuestion: repo.NewQuestion(),
	}
}

func (a *CreateQuestion) Handler(ctx context.Context, req entity.Question) (err error) {
	question, err := a.repoQuestion.FindQuestion(ctx, req)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Printf("[Handler] Find Question : %v", err)
			return err
		}
	}

	no := strconv.Itoa(question.No)
	if question.No == req.No {
		return stringErr.New("Question no " + no + " already existed!")
	}

	err = a.repoQuestion.CreateQuestion(ctx, req)
	if err != nil {
		log.Printf("[Handler] Create Question : %v", err)
		return err
	}
	return nil
}
