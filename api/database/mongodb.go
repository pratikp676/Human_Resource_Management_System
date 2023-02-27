package database

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
)

var db *mgo.Database

func InitDB() {
	session, err := mgo.Dial("localhost")
	fmt.Println(session)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	db = session.DB("EmployeeDetails")

}

func Collection() *mgo.Collection {
	return db.C("EmployeeData")
}
