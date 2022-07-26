package repo

import (
	"context"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/erwinhermanto31/quiz_master/entity"
	"github.com/erwinhermanto31/quiz_master/repo/mysql"
	"github.com/erwinhermanto31/quiz_master/repo/mysql/mocks_mysql"
	"github.com/erwinhermanto31/quiz_master/util"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/mock"
)

var (
	mockMysql *mocks_mysql.IMysql
)

func initQuestionTest() *Question {
	mockMysql = &mocks_mysql.IMysql{}
	db, _, _ := sqlmock.New()

	// defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	return &Question{
		iMysql:       mockMysql,
		quizMasterDB: sqlxDB,
	}
}

func TestNewQuestion(t *testing.T) {
	t.Run("TestNewQuestion", func(t *testing.T) {

		tx := NewQuestion()
		if tx == nil {
			t.Errorf("NewQuestion() = %+v, want %+v", tx, tx)
		}
	})
}

func TestFindQuestion(t *testing.T) {
	m := initQuestionTest()

	no := int(1)
	query := &util.Query{
		Filter: map[string]interface{}{
			"no$eq!": no,
		},
	}

	type args struct {
		ctx     context.Context
		request entity.Question
	}
	tests := []struct {
		name    string
		args    args
		want    entity.Question
		wantErr bool
		fn      func()
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				request: entity.Question{
					No: 1,
				},
			},
			want:    entity.Question{},
			wantErr: false,
			fn: func() {
				mockMysql.Mock.On("Get", mock.Anything, m.quizMasterDB, &entity.Question{Question: ""}, query, mysql.QueryFindQuestions).Return(nil, nil).Once()
			},
		},
	}
	for _, tt := range tests {
		tt.fn()
		t.Run(tt.name, func(t *testing.T) {
			got, err := m.FindQuestion(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindQuestion() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err == nil {
				if !reflect.DeepEqual(got.Question, tt.want.Question) {
					t.Errorf("FindQuestion() got = %v, want %v", got.Question, tt.want.Question)
				}
			}
		})
	}
}

func TestFindAllQuestion(t *testing.T) {
	m := initQuestionTest()

	// no := int(1)
	query := &util.Query{
		Sort: "no",
	}

	type args struct {
		ctx     context.Context
		request entity.Question
	}
	tests := []struct {
		name    string
		args    args
		want    *entity.Question
		wantErr bool
		fn      func()
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				request: entity.Question{
					No: 1,
				},
			},
			want:    nil,
			wantErr: false,
			fn: func() {
				mockMysql.Mock.On("Select", mock.Anything, m.quizMasterDB, mock.Anything, query, mysql.QueryFindQuestions).Return(nil).Once()
			},
		},
	}
	for _, tt := range tests {
		tt.fn()
		t.Run(tt.name, func(t *testing.T) {
			_, err := m.FindAllQuestion(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindAllQuestion() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateQuestion(t *testing.T) {
	m := initQuestionTest()

	type args struct {
		ctx     context.Context
		request entity.Question
	}
	tests := []struct {
		name    string
		args    args
		want    *entity.Question
		wantErr bool
		fn      func()
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				request: entity.Question{
					No: 1,
				},
			},
			want:    nil,
			wantErr: false,
			fn: func() {
				mockMysql.Mock.On("CreateOrUpdate", mock.Anything, m.quizMasterDB, mock.Anything, mysql.QueryInsertQuestion).Return(int64(0), nil).Once()
			},
		},
	}
	for _, tt := range tests {
		tt.fn()
		t.Run(tt.name, func(t *testing.T) {
			err := m.CreateQuestion(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateQuestion() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
