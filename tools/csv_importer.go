//go:build ignore
// +build ignore

package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	records, err := readCsvFile("authors.csv")

	if err != nil {
		log.Fatal(err)
	}

	var (
		host     = "127.0.0.1"
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USERNAME")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
	)

	db, err := sql.Open("mysql", user+":"+password+"@tcp("+host+":"+port+")/"+dbname+"?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	for index, record := range records {
		if index == 0 {
			continue
		}

		sqlStatement := fmt.Sprintf("INSERT INTO authors(name) VALUES ('%s')", record[0])

		res, err := db.Exec(sqlStatement)

		if err != nil {
			fmt.Println("err", err)
			err = nil
			continue
		}

		fmt.Println(res)
		fmt.Println(record[0], "inserted!")
	}
}

func readCsvFile(fileName string) (records [][]string, err error) {
	file, err := os.Open(fileName)

	if err != nil {
		return
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err = csvReader.ReadAll()

	if err != nil {
		return
	}

	return
}
