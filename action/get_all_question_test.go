package action

import (
	"context"
	"testing"

	"github.com/erwinhermanto31/quiz_master/entity"
	"github.com/erwinhermanto31/quiz_master/repo/mocks_repo"
	"github.com/stretchr/testify/mock"
)

func initGetAllQuestionTest() *GetAllQuestion {
	mockRepo = &mocks_repo.IQuestions{}

	return &GetAllQuestion{
		repoQuestion: mockRepo,
	}
}

func TestNewGetAllQuestion(t *testing.T) {
	t.Run("TestNewGetAllQuestion", func(t *testing.T) {

		tx := NewGetAllQuestion()
		if tx == nil {
			t.Errorf("NewGetAllQuestion() = %+v, want %+v", tx, tx)
		}
	})
}

func TestGetAllQuestion(t *testing.T) {
	m := initGetAllQuestionTest()

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
				mockRepo.Mock.On("FindAllQuestion", mock.Anything, mock.Anything).Return([]entity.Question{}, nil).Once()
			},
		},
	}
	for _, tt := range tests {
		tt.fn()
		t.Run(tt.name, func(t *testing.T) {
			_, err := m.Handler(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handler() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}
