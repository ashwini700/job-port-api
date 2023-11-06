package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"

	"job-port-api/internal/auth"
	"job-port-api/internal/models"
	"job-port-api/internal/repository"
)

func TestService_AddJobDetails(t *testing.T) {
	type args struct {
		ctx     context.Context
		jobData models.Job
	}
	tests := []struct {
		name             string
		args             args
		want             models.Job
		wantErr          bool
		mockRepoResponse func() (models.Job, error)
	}{
		{
			name: "error in db",
			args: args{
				ctx: context.Background(),
			},
			want:    models.Job{},
			wantErr: true,
			mockRepoResponse: func() (models.Job, error) {
				return models.Job{}, errors.New("test error")
			},
		},
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				jobData: models.Job{
					Cid:     1,
					JobRole: "developer",
					Salary:  "200000",
				},
			},
			want: models.Job{
				Cid:     1,
				JobRole: "developer",
				Salary:  "200000",
			},
			wantErr: false,
			mockRepoResponse: func() (models.Job, error) {
				return models.Job{
					Cid:     1,
					JobRole: "developer",
					Salary:  "200000",
				}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockRepo := repository.NewMockUserRepo(mc)
			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().CreateJob(tt.args.ctx, tt.args.jobData).Return(tt.mockRepoResponse()).AnyTimes()
			}
			s, _ := NewService(mockRepo, &auth.Auth{})
			got, err := s.AddJobDetails(tt.args.ctx, tt.args.jobData)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.AddJobDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.AddJobDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}
