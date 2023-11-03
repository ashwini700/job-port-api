package service

import (
	"context"
	"job-port-api/internal/models"
	"reflect"
	"testing"
)

func TestService_AddJobDetails(t *testing.T) {
	type args struct {
		ctx     context.Context
		jobData models.Job
	}
	tests := []struct {
		name    string
		s       *Service
		args    args
		want    models.Job
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.AddJobDetails(tt.args.ctx, tt.args.jobData)
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
