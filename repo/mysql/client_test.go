package mysql

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/erwinhermanto31/quiz_master/entity"
	"github.com/erwinhermanto31/quiz_master/util"
	"github.com/jmoiron/sqlx"
)

// test function that get list of lof file in n minutes
func TestGet(t *testing.T) {
	// assert := assert.New(t)
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	client := NewClient()
	query := &util.Query{
		Filter: map[string]interface{}{
			"no$eq!": 1,
		},
	}
	var data entity.Question

	rows := sqlmock.NewRows([]string{"question"}).
		AddRow(1)
	mock.ExpectQuery(regexp.QuoteMeta(QueryFindQuestions)).
		WillReturnRows(rows)

	err = client.Get(context.Background(), sqlxDB, &data, query, QueryFindQuestions)
	if err != nil {
		t.Log(err)
		t.Error("Failed to Process Files")
	}
}

func TestSelect(t *testing.T) {
	// assert := assert.New(t)
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	client := NewClient()
	query := &util.Query{
		Sort: "no",
	}
	var data []entity.Question

	rows := sqlmock.NewRows([]string{"question"}).
		AddRow(1)
	mock.ExpectQuery(regexp.QuoteMeta(QueryFindQuestions)).
		WillReturnRows(rows)

	err = client.Select(context.Background(), sqlxDB, &data, query, QueryFindQuestions)
	if err != nil {
		t.Log(err)
		t.Error("Failed to Process Files")
	}
}

func TestCreateOrUpdate(t *testing.T) {
	// assert := assert.New(t)
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	// mock.ExpectBegin()
	client := NewClient()
	data := &entity.Question{
		No:       1,
		Question: "test",
		Answer:   "25",
	}

	mock.ExpectQuery(QueryInsertQuestion)

	_, err = client.CreateOrUpdate(context.Background(), sqlxDB, &data, QueryInsertQuestion)
	if err == nil {
		t.Log(err)
		t.Error("Failed to Process Files")
	}
}
