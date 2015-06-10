package utilities

import (
	"github.com/go-martini/martini"
	"gopkg.in/mgo.v2"
	"reflect"
	"strings"
)

type DbConfig struct {
	Name     string       `json:"name"`
	User     string       `json:"user"`
	Password string       `json:"password"`
	Host     dbHostConfig `json:"host"`
}
type dbHostConfig struct {
	Read  string `json:"read"`
	Write string `json:"write"`
}
type DbSession struct {
	read  *mgo.Database
	write *mgo.Database
}

func DataHelper(dbc DbConfig) martini.Handler {
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

	return func(c martini.Context) {
		sR := sessionRead.Clone()
		sW := sessionRead.Clone()

		var db = DbSession{
			read:  sR.DB(dbc.Name),
			write: sW.DB(dbc.Name),
		}
		c.Map(db)
		defer sR.Close()
		defer sW.Close()
		c.Next()
	}

}

func (db DbSession) Read(i interface{}) *mgo.Collection {
	return db.read.C(db.getTableName(i))
}
func (db DbSession) Write(i interface{}) *mgo.Collection {
	return db.write.C(db.getTableName(i))
}

func (db DbSession) getTableName(i interface{}) string {
	v := reflect.ValueOf(i)
	if v.Type().Kind() == reflect.Slice {
		return strings.ToLower(v.Type().Elem().Name())
		
	}
	return strings.ToLower(v.Type().Name())
}

func (db DbSession) Populate(result interface{},filedNmae string ,filedinterface interface{}){

//	for key, value := range result {
//		print(key)
//		print(value)
//	}
//	db.Read(filedinterface).Find().All(&filedinterface)
	
}
//item := models.Item{}
/*
	users := []models.User{}
	db.Read(models.User{}).Find(bson.M{}).All(&users)
	fmt.Println(users)
*/
/*

type Person struct {
	Name  string
	Phone string
}
func test() {
	session, err := mgo.Dial("192.168.0.30")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Phone)
}
*/
