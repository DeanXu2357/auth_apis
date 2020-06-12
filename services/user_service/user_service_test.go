package user_service

import (
	"github.com/jinzhu/gorm"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name        string
		wantService *UserService
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotService, err := New()
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotService, tt.wantService) {
				t.Errorf("New() gotService = %v, want %v", gotService, tt.wantService)
			}
		})
	}
}

func TestUserService_Create(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		data map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &UserService{
				DB: tt.fields.DB,
			}
			if err := service.Create(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserService_GetUserByUUID(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		uuid string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &UserService{
				DB: tt.fields.DB,
			}
			if err := service.GetUserByUUID(tt.args.uuid); (err != nil) != tt.wantErr {
				t.Errorf("GetUserByUUID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}