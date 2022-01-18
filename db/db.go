package db

import (
	"database/sql"
	"fmt"
	"log"
	model "mingo/audit/model"
	"os"
	"sync"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

type DAO interface {
	Insert(ev model.Event)
	Free()
	GetEvents(query model.EventQuery) ([]model.Event, error)
}

type daoImpl struct {
	db *sql.DB
}

func (d *daoImpl) Insert(ev model.Event) {
	fmt.Println("Insertando en DB")
	insertEvent(d.db, ev)
}

func (d *daoImpl) GetEvents(query model.EventQuery) ([]model.Event, error) {
	fmt.Print("DAO GetEvents ")
	fmt.Print("eventType:")
	fmt.Println(query.EventType)
	fmt.Print(" startDate:")
	fmt.Println(query.StartDate)
	fmt.Print(" endDate:")
	fmt.Println(query.EndDate)

	statement, err := d.db.Prepare("SELECT * FROM event where eventType=? and (datetime(createdTime) BETWEEN datetime(?) AND datetime(?))")
	if err != nil {
		fmt.Println("fatal")
		log.Fatal(err)
	}
	row, err := statement.Query(query.EventType, query.StartDate, query.EndDate)
	if err != nil {
		fmt.Println("fatal")
		log.Fatal(err)
	}
	defer row.Close()
	var list []model.Event
	for row.Next() {
		var ev model.Event
		var id int
		if err := row.Scan(&id, &ev.EventType, &ev.CreatedTime, &ev.DynamicData, &ev.Gravity); err != nil {
			fmt.Print("ERROR EN LA CONSULTA")
			return list, err
		}
		log.Println("Event: ", ev.EventType, " ", ev.CreatedTime)
		list = append(list, ev)
	}
	if err = row.Err(); err != nil {
		return list, err
	}
	fmt.Println("retornando lista")
	return list, nil
}

func (d *daoImpl) Free() {
	d.db.Close()
}

func createTable(db *sql.DB) {
	createEventsTableSQL := `CREATE TABLE event (
		"idEvent" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"eventType" TEXT NOT NULL,
		"createdTime" TEXT NOT NULL,
		"dynamicData" TEXT,
		"gravity" TEXT
	  );` // SQL Statement for Create Table

	log.Println("Create events table...")
	statement, err := db.Prepare(createEventsTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("events table created")
}

func insertEvent(db *sql.DB, ev model.Event) {
	log.Print("Inserting event record ...")
	fmt.Print("createdTime: ")
	fmt.Println(ev.CreatedTime)
	insertStudentSQL := `INSERT INTO event(eventType,createdTime,dynamicData,gravity) VALUES (?,?,?,?)`
	statement, err := db.Prepare(insertStudentSQL) // Prepare statement.
	if err != nil {
		fmt.Println("fatal")
		log.Fatalln(err.Error())
	}
	fmt.Println("executing insert")
	res, err := statement.Exec(ev.EventType, ev.CreatedTime, ev.DynamicData, ev.Gravity)
	if err != nil {
		fmt.Println("fatal")
		log.Fatalln(err.Error())
	} else {
		fmt.Println(res.RowsAffected())
	}
	fmt.Println("inserción en BD correcta")
}

var dao *daoImpl

var mu sync.Mutex

func GetDAO() DAO {
	mu.Lock()
	if dao == nil {
		dao = new(daoImpl)
		os.Remove("sqlite-database.db")
		// SQLite is a file based database.
		fmt.Println("Creating sqlite-database.db...")
		file, err := os.Create("sqlite-database.db") // Create SQLite file
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
		log.Println("sqlite-database.db file created")

		sqliteDatabase, err := sql.Open("sqlite3", "./sqlite-database.db?_journal_mode=WAL") // Open the created SQLite File
		if err != nil {
			fmt.Println("ERRORRRRRRRRRR")
			log.Fatal(err.Error())
		}
		fmt.Print("sqlitedb: ")
		fmt.Println(sqliteDatabase)
		dao.db = sqliteDatabase
		createTable(sqliteDatabase) // Create Database Tables
	}
	mu.Unlock()
	return dao
}
