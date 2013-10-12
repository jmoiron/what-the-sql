package main

import (
	"fmt"
	"github.com/jmoiron/modl"
	"github.com/jmoiron/sqlx"
	"github.com/mailgun/gocql/uuid"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"os"
	"path"
)

const DBROOT = "/dev/shm/"

var PORT = "3312"

type Employee struct {
	EmployeeID   int
	DepartmentID int // -> Department.DepartmentID
	BossID       int // -> Employee.EmployeeID
	Name         string
	Salary       int
}

type Department struct {
	DepartmentID int
	Name         string
}

type Status struct {
	CurrentQuestion int
}

func exists(id string) bool {
	dbpath := path.Join(DBROOT, id+".db")
	_, err := os.Stat(dbpath)
	return err == nil
}

func OpenDatabase(id string) (*modl.DbMap, error) {
	db, err := sqlx.Connect("sqlite3", path.Join(DBROOT, id+".db"))
	if err != nil {
		return nil, err
	}
	dbm := modl.NewDbMap(&db.DB, modl.SqliteDialect{})
	dbm.AddTableWithName(Employee{}, "Employees").SetKeys(true, "EmployeeID")
	dbm.AddTableWithName(Department{}, "Departments").SetKeys(true, "DepartmentID")
	dbm.AddTableWithName(Status{}, "_status").SetKeys(true, "CurrentQuestion")
	return dbm, err
}

func CreateDatabase(id string) (*modl.DbMap, error) {
	// create a database in DBROOT/id
	alreadyExists := exists(id)
	dbm, err := OpenDatabase(id)
	if err != nil {
		return nil, err
	}
	if !alreadyExists {
		err = dbm.CreateTables()
		if err != nil {
			return nil, err
		}
		LoadData(dbm)
	}
	return dbm, nil
}

func CreateNewDatabase() (*modl.DbMap, string, error) {
	u := uuid.RandomUUID().String()
	dbm, err := CreateDatabase(u)
	if err != nil {
		LoadData(dbm)
	}
	return dbm, u, err
}

func main() {
	fmt.Println("Listening on :" + PORT)
	fmt.Println(http.ListenAndServe(":"+PORT, nil))
}

func init() {
	//sqlx.NameMapper = func(s string) string { return s }
}
