package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDatabase() {
	dsn := "host=database port=5432 user=postgres password=postgres dbname=app sslmode=disable"
	dialector := postgres.Open(dsn)
	var err error
	if Db, err = gorm.Open(dialector); err != nil {
		fmt.Println("db error!!")
		panic(err)
	}
	fmt.Println("db connected!!")
	Migration(Db)
	UserSeeder(Db)

	createDateMaster(Db)
	insertDateMaster(Db)

}
func createDateMaster(Db *gorm.DB) error {
	if err := Db.Exec(`CREATE TABLE IF NOT EXISTS date_masters (
        date DATE PRIMARY KEY,
        day_of_week INT,
        is_weekend BOOLEAN,
        is_holiday BOOLEAN,
        quarter INT,
        year INT,
        month INT
    );`).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func insertDateMaster(Db *gorm.DB) error {
	if err := Db.Exec(`INSERT INTO date_masters
        SELECT
            generate_series AS date,
            EXTRACT(ISODOW FROM generate_series) AS day_of_week,
            EXTRACT(ISODOW FROM generate_series) IN (6, 7) AS is_weekend,
            FALSE AS is_holiday,
            EXTRACT(QUARTER FROM generate_series) AS quarter,
            EXTRACT(YEAR FROM generate_series) AS year,
            EXTRACT(MONTH FROM generate_series) AS month
        FROM
            generate_series('2023-01-01'::date, '2032-12-31'::date, '1 day');
    `).Error; err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
