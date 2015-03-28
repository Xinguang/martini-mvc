package utilities

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

func Test_Insert(t *testing.T) {
	db := getSession()
	person := Person{"Ale", "+55 53 8116 9639"}
	err := db.Write(person).Insert(person)
	t.Log(err)
}
func Test_Find(t *testing.T) {
	db := getSession()
	person := Person{}
	person.Name = "Ale"
	err := db.Read(person).Find(bson.M{"name": person.Name}).One(&person)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(person.Phone)
	}
}
func Test_Delete(t *testing.T) {
	db := getSession()
	person := Person{}
	person.Name = "Ale"
	err := db.Write(person).Remove(bson.M{"name": person.Name})
	t.Log(err)
}
func getTestDb() DbConfig {
	return DbConfig{
		Name:     "test",
		User:     "",
		Password: "",
		Host: dbHostConfig{
			Read:  "192.168.0.31",
			Write: "192.168.0.30",
		},
	}
}
func getSession() DbSession {
	dbc := getTestDb()
	sessionRead, errRead := mgo.Dial(dbc.Host.Read)
	if errRead != nil {
		panic(errRead)
	}
	sessionWrite, errWrite := mgo.Dial(dbc.Host.Write)
	if errWrite != nil {
		panic(errWrite)
	}
	sessionRead.SetMode(mgo.Monotonic, true)
	sessionWrite.SetMode(mgo.Monotonic, true)
	return DbSession{
		read:  sessionRead.DB(dbc.Name),
		write: sessionWrite.DB(dbc.Name),
	}
}

type Person struct {
	Name  string
	Phone string
}
