package repo

import (
	"context"
	"log"

	"github.com/erwinhermanto31/quiz_master/entity"
	"github.com/erwinhermanto31/quiz_master/repo/mysql"
	"github.com/erwinhermanto31/quiz_master/util"
	"github.com/jmoiron/sqlx"
)

type Question struct {
	iMysql       mysql.IMysql
	quizMasterDB *sqlx.DB
}

func NewQuestion() *Question {
	return &Question{
		iMysql:       mysql.NewClient(),
		quizMasterDB: mysql.QuizMasterDB,
	}
}

func (r *Question) FindAllQuestion(ctx context.Context, req entity.Question) (res []entity.Question, err error) {
	query := &util.Query{
		Sort: "no",
	}

	err = r.iMysql.Select(ctx, r.quizMasterDB, &res, query, mysql.QueryFindQuestions)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r *Question) FindQuestion(ctx context.Context, req entity.Question) (res entity.Question, err error) {
	query := &util.Query{
		Filter: map[string]interface{}{
			"no$eq!": req.No,
		},
	}

	err = r.iMysql.Get(ctx, r.quizMasterDB, &res, query, mysql.QueryFindQuestions)
	if err != nil {
		return res, err
	}
	return res, err
}

func (r *Question) CreateQuestion(ctx context.Context, req entity.Question) (err error) {
	_, err = r.iMysql.CreateOrUpdate(ctx, r.quizMasterDB, &req, mysql.QueryInsertQuestion)
	if err != nil {
		log.Printf("[CreateQuestion] get error: %v", err)
		return err
	}

	return nil
}

func (r *Question) UpdateQuestion(ctx context.Context, req entity.Question) (err error) {
	_, err = r.iMysql.CreateOrUpdate(ctx, r.quizMasterDB, &req, mysql.QueryUpdateQuestion)
	if err != nil {
		log.Printf("[UpdateQuestion] get error: %v", err)
		return err
	}

	return nil
}

func (r *Question) DeleteQuestion(ctx context.Context, req entity.Question) (err error) {
	_, err = r.iMysql.CreateOrUpdate(ctx, r.quizMasterDB, &req, mysql.QueryDeleteQuestion)
	if err != nil {
		log.Printf("[DeleteQuestion] get error: %v", err)
		return err
	}

	return nil
}
