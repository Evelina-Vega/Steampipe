package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

//export STEAMPIPE_LOG_LEVEL=trace steampipe query
var instanceQuery = `
		select instance_id from aws_ec2_instance
`

var s3FailedQuery = `
		select name from aws_cloudtrail_trail
`

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// use db connection
func queryPostgresForRepos(query string) {
	host := "localhost"
	port := 9193
	user := "steampipe"
	password := "c08a_4d27_97ad"
	dbname := "steampipe"
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	rows, _ := db.Query(query)
	var names string
	for rows.Next() {
		err = rows.Scan(&names)
		fmt.Printf("%s \n", names)
		CheckError(err)
	}
}

func main() {
	queryPostgresForRepos(instanceQuery)

	queryPostgresForRepos(s3FailedQuery)
}
