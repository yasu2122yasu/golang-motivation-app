package service

import (
	"go_practice/repository"
	"reflect"
	"testing"
)

func TestIndexSavingService(t *testing.T) {
	tests := []struct {
		name    string
		want    []repository.Person
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IndexSavingService()
			if (err != nil) != tt.wantErr {
				t.Errorf("IndexSavingService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IndexSavingService() = %v, want %v", got, tt.want)
			}
		})
	}
}
