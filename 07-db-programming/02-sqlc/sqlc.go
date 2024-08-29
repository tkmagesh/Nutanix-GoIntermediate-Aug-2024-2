package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	conn, err := sql.Open("mysql", "theuser:thepass@tcp(localhost:3306)/thedb?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!")

	db := New(conn)

	//initializing record to be inserted
	newSt := addStudentParams{
		Fname:       "Leon",
		Lname:       "Ashling",
		DateOfBirth: time.Date(1994, time.August, 14, 23, 51, 42, 0, time.UTC),
		Email:       "lashling5@senate.gov",
		Gender:      "Male",
		Address:     "39 Kipling Pass",
	}
	// inserting the record
	sID, err := db.addStudent(context.Background(), newSt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("addSudent id: %v \n", sID)

	//retreive record by id
	st, err := db.studentByID(context.Background(), sID)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("studentByID record: %v \n", st)

	//fetching multiple records
	students, err := db.fetchStudents(context.Background())
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("fetchStudents count: %v \n", len(students))
}
