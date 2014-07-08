package storage

import (
	"fmt"
	"github.com/shinpei/comstock/model"
	"labix.org/v2/mgo"
	"log"
	"time"
)

type MongoStorager struct {
	mongopath string
}

const (
	MongoDatabaseName   string = "comstock"
	MongoCollectionName string = "command"
	MongoHost           string = "localhost"
)

func (ms *MongoStorager) Open() (err error) {
	return
}

func CreateMongoStorager() *MongoStorager {
	return &MongoStorager{}
}

func (ms *MongoStorager) StorageType() string {
	return "MongoStorager"

}

type Person struct {
	Name  string
	Phone string
}

//store command
func (ms *MongoStorager) Push(user *model.UserInfo, path string, cmd *model.Command) (err error) {
	hostname := MongoHost
	session, err := mgo.DialWithTimeout("mongodb://"+hostname, time.Duration(3)*time.Second)
	if err != nil {
		log.Fatal("Couldn't dial to ", hostname, ", ", err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(MongoDatabaseName).C(MongoCollectionName)
	err = c.Insert(cmd)
	if err != nil {
		log.Fatal("Couldn't insert")
	}
	if err != nil {
		log.Fatal(err)
	}

	return
}

func (ms *MongoStorager) Close() (err error) {
	return
}

func (ms *MongoStorager) FetchCommandFromNumber(user *model.UserInfo, num int) (cmd *model.Command, err error) {
	hostname := MongoHost
	var session *mgo.Session
	session, err = mgo.DialWithTimeout("mongodb://"+hostname, time.Duration(3)*time.Second)
	if err != nil {
		log.Fatal("Couldn't dial to ", hostname, ",", err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(MongoDatabaseName).C(MongoCollectionName)

	iter := c.Find(nil).Limit(100).Iter()
	var idx = 1
	var result model.Command
	for iter.Next(&result) {
		if idx == num {
			cmd = model.CreateCommand(result.Cmd)
			break
		}
		idx++
	}
	if err = iter.Close(); err != nil {
		log.Fatal(err)
	}
	return
}
func (ms *MongoStorager) List(user *model.UserInfo) (cmds []model.Command, err error) {
	hostname := MongoHost
	session, err := mgo.DialWithTimeout("mongodb://"+hostname, time.Duration(3)*time.Second)
	if err != nil {
		log.Fatal("Couldn't dial to ", hostname, ", ", err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(MongoDatabaseName).C(MongoCollectionName)
	//	ensureIndex(c)

	// TODO: make cmd array
	var result model.Command
	iter := c.Find(nil).Limit(100).Iter()
	var idx = 1
	for iter.Next(&result) {
		fmt.Printf("%d: %s\n", idx, result.Cmd)
		idx++
	}
	if err = iter.Close(); err != nil {
		log.Fatal(err)
	}
	return
}

func ensureIndex(col *mgo.Collection) {
	index := mgo.Index{
		Key:        []string{"cmd"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err := col.EnsureIndex(index)
	if err != nil {
		log.Fatal(err)
	}
}

func (m *MongoStorager) IsRequireLogin() bool {
	return false
}

func (m *MongoStorager) Status() (err error) {
	return
}
func (m *MongoStorager) CheckSession(user *model.UserInfo) bool {
	return true
}
