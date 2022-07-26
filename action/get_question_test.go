package action

import (
	"context"
	"reflect"
	"testing"

	"github.com/erwinhermanto31/quiz_master/entity"
	"github.com/erwinhermanto31/quiz_master/repo/mocks_repo"
	"github.com/stretchr/testify/mock"
)

var (
	mockRepo *mocks_repo.IQuestions
)

func initQuestionTest() *GetQuestion {
	mockRepo = &mocks_repo.IQuestions{}

	return &GetQuestion{
		repoQuestion: mockRepo,
	}
}

func TestNewGetQuestion(t *testing.T) {
	t.Run("TestNewGetQuestion", func(t *testing.T) {

		tx := NewGetQuestion()
		if tx == nil {
			t.Errorf("NewGetQuestion() = %+v, want %+v", tx, tx)
		}
	})
}

func TestGetQuestion(t *testing.T) {
	m := initQuestionTest()

	// no := int(1)

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
					No:       1,
					Question: "",
					Answer:   "",
				},
			},
			want:    entity.Question{},
			wantErr: false,
			fn: func() {
				mockRepo.Mock.On("FindQuestion", mock.Anything, mock.Anything).Return(entity.Question{}, nil).Once()
			},
		},
	}
	for _, tt := range tests {
		tt.fn()
		t.Run(tt.name, func(t *testing.T) {
			got, err := m.Handler(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handler() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err == nil {
				if !reflect.DeepEqual(got.Question, tt.want.Question) {
					t.Errorf("Handler() got = %v, want %v", got.Question, tt.want.Question)
				}
			}
		})
	}
}
