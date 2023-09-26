package service

import (
	"go_practice/repository"
	"reflect"
	"testing"
)

func TestIndexPersonService(t *testing.T) {
	tests := []struct {
		name    string
		want    []repository.Person
		wantErr bool
	}{
		{
			"success",
			[]repository.Person{
				{
					Id:     1,
					Name:   "Tara",
					Gender: "Male",
					Saving: "1000",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IndexPersonService()
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
