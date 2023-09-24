package database

import (
	"database/sql"
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestNewConnectDatabase(t *testing.T) {
	tests := []struct {
		name string
		want ConnectDatabase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConnectDatabase(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConnectDatabase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConnectDatabase_Connect(t *testing.T) {
	tests := []struct {
		name    string
		c       ConnectDatabase
		want    *sql.DB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Connect()
			if (err != nil) != tt.wantErr {
				t.Errorf("ConnectDatabase.Connect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConnectDatabase.Connect() = %v, want %v", got, tt.want)
			}
		})
	}
}
