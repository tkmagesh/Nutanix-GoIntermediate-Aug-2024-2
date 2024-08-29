-- name: addStudent :execlastid
insert into students (fname, lname, date_of_birth, email, gender, address) values (?, ?, ?, ?, ?, ?);

-- name: studentByID :one
SELECT * FROM students WHERE id = ?;

-- name: fetchStudents :many
SELECT * FROM students LIMIT 10;

-- name: updateStudent :exec
update students set fname=?, lname=?, email=? where ID=?