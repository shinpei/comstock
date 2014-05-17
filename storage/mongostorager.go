package storage

import (
	"fmt"
	"github.com/shinpei/comstock/model"
	"labix.org/v2/mgo"
	"log"
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

//store command
func (ms *MongoStorager) Push(path string, cmd *model.Command) (err error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal("Couldn't dial")
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(MongoDatabaseName).C(MongoCollectionName)
	err = c.Insert(cmd)
	println("Insert done")

	println(cmd)
	if err != nil {
		log.Fatal(err)
	}

	return
}

func (ms *MongoStorager) Close() (err error) {
	return
}

func (ms *MongoStorager) FetchCommandFromNumber(num int) (cmd *model.Command) {
	return
}
func (ms *MongoStorager) List() (err error) {
	hostname := MongoHost
	session, err := mgo.Dial(hostname)
	if err != nil {
		log.Fatal("Couldn't dial " + hostname)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(MongoDatabaseName).C(MongoCollectionName)
	ensureIndex(c)

	var results []model.Command
	scmd := &model.Command{}
	err = c.Find(scmd).All(&results)
	if err != nil {
		log.Fatal("Couldn't fetch ")
	}
	fmt.Println("result: ", results[0].Cmd())

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
