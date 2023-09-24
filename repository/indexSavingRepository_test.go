package repository

import (
	"go_practice/database"
	"reflect"
	"testing"
)

func TestNewIndexSavingRepository(t *testing.T) {
	type args struct {
		db database.IConnectDatabase
	}
	tests := []struct {
		name string
		args args
		want *IndexSavingRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIndexSavingRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIndexSavingRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrashScanner_Scan(t *testing.T) {
	type args struct {
		in0 interface{}
	}
	tests := []struct {
		name    string
		tr      TrashScanner
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.tr.Scan(tt.args.in0); (err != nil) != tt.wantErr {
				t.Errorf("TrashScanner.Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIndexSavingRepository_GetPersons(t *testing.T) {
	tests := []struct {
		name    string
		isr     IndexSavingRepository
		want    []Person
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.isr.GetPersons()
			if (err != nil) != tt.wantErr {
				t.Errorf("IndexSavingRepository.GetPersons() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IndexSavingRepository.GetPersons() = %v, want %v", got, tt.want)
			}
		})
	}
}
