package repository

import (
	"project/internal/model"
	"reflect"
	"testing"
)

func TestRepo_CreateCompany(t *testing.T) {
	type args struct {
		u model.Company
	}
	tests := []struct {
		name    string
		r       *Repo
		args    args
		want    model.Company
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.CreateCompany(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repo.CreateCompany() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repo.CreateCompany() = %v, want %v", got, tt.want)
			}
		})
	}
}
