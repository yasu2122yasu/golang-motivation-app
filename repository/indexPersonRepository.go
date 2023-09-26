package repository

import (
	"go_practice/database"
)

type IIndexPersonRepository interface {
	GetPersons() ([]Person, error)
}

type IndexPersonRepository struct {
	db database.IConnectDatabase
}

// リポジトリの作成
func NewIndexPersonRepository(db database.IConnectDatabase) *IndexPersonRepository {
	return &IndexPersonRepository{
		db,
	}
}

type Person struct {
	Id     int      `json:"id"`
	Name   string   `json:"name"`
	Gender string   `json:"gender"`
	Saving string   `json:"saving"`
}

type TrashScanner struct{}

func (TrashScanner) Scan(interface{}) error {
		return nil
}

func (isr IndexPersonRepository) GetPersons() ([]Person, error) {
	// データベースのハンドルを取得する
	db, err := isr.db.Connect()

	// SQLの実行
	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var persons []Person
	// SQLの実行
	for rows.Next() {
		var p Person

		err := rows.Scan(&p.Id, &p.Name, &p.Gender,&p.Saving, TrashScanner{}, TrashScanner{}, TrashScanner{})
		if err != nil {
			return nil, err
		}

		persons = append(persons, p)

	}

	// Check for errors from rows.Next()
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return persons, nil
}
