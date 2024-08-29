package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Student struct {
	ID          int
	Fname       string
	Lname       string
	DateOfBirth time.Time `db:"date_of_birth"`
	Email       string
	Address     string
	Gender      string
}

var db *sqlx.DB

func main() {
	var err error
	db, err = sqlxConnect()
	if err != nil {
		log.Fatal(err)
	}

	s := Student{
		Fname:       "Leon",
		Lname:       "Ashling",
		DateOfBirth: time.Date(1994, time.August, 14, 23, 51, 42, 0, time.UTC),
		Email:       "lashling5@senate.gov",
		Address:     "39 Kipling Pass",
		Gender:      "Male",
	}

	//adding student record to table
	sID, err := addStudent(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("addStudent id: %v \n", sID)

	//selecting student by ID
	st, err := studentByID(sID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("studentByID record: %v \n", st)

	students, err := fetchStudents()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("fetchStudents count: %v \n", len(students))
}

func sqlxConnect() (*sqlx.DB, error) {
	// Opening a database connection.
	db, err := sqlx.Open("mysql", "root:rootuser@tcp(localhost:3306)/go_db_demo?parseTime=true")
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected!")
	return db, nil
}

func addStudent(s Student) (int64, error) {
	query := "insert into students (fname, lname, date_of_birth, email, gender, address) values (?, ?, ?, ?, ?, ?);"
	result := db.MustExec(query, s.Fname, s.Lname, s.DateOfBirth, s.Email, s.Gender, s.Address)

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addSudent Error: %v", err)
	}

	return id, nil
}

func fetchStudents() ([]Student, error) {
	// A slice of Students to hold data from returned rows.
	var students []Student

	err := db.Select(&students, "SELECT * FROM students LIMIT 10")
	if err != nil {
		return nil, fmt.Errorf("fetchStudents %v", err)
	}

	return students, nil
}

func studentByID(id int64) (Student, error) {
	var st Student

	//if err := db.QueryRowx("SELECT * FROM students WHERE id = ?", id).StructScan(&st); err != nil {
	//	if err == sql.ErrNoRows {
	//		return st, fmt.Errorf("studentById %d: no such student", id)
	//	}
	//	return st, fmt.Errorf("studentById %d: %v", id, err)
	//}

	if err := db.Get(&st, "SELECT * FROM students WHERE id = ?", id); err != nil {
		if err == sql.ErrNoRows {
			return st, fmt.Errorf("studentById %d: no such student", id)
		}
		return st, fmt.Errorf("studentById %d: %v", id, err)
	}
	return st, nil
}
