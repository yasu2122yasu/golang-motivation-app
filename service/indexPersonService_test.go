package service

import (
	mock_repository "go_practice/mock/repository"
	"go_practice/repository"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestNewIndexSavingService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	isrMock := mock_repository.NewMockIIndexPersonRepository(ctrl)

	tests := []struct {
		name string
		args repository.IIndexPersonRepository
		want *IndexSavingService
	}{
		{
			"success",
			isrMock,
			&IndexSavingService{isrMock},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIndexSavingService(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIndexSavingService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndexSavingService_IndexSaving(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	isrMock := mock_repository.NewMockIIndexPersonRepository(ctrl)

	// モックの期待値を設定
	expectedPersons := []repository.Person{
		{Id: 1, Name: "John", Gender: "Female", Saving: "1000"},
		{Id: 2, Name: "Tako", Gender: "Male", Saving: "2000"},
	}
	isrMock.EXPECT().GetPersons().Return(expectedPersons, nil).AnyTimes()

	tests := []struct {
		name    string
		iss     *IndexSavingService
		want    []repository.Person
		wantErr bool
	}{
		{
			"success",
			&IndexSavingService{isrMock},
			expectedPersons,
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.iss.IndexSaving()
			if (err != nil) != tt.wantErr {
				t.Errorf("IndexSavingService.IndexSaving() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IndexSavingService.IndexSaving() = %v, want %v", got, tt.want)
			}
		})
	}
}
