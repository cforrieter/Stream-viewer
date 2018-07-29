package main

import (
	"log"

	"encoding/json"
	"net/http"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/julienschmidt/httprouter"
)

type message struct {
	Author  string `bson:"author"`
	Message string `bson:"message"`
}

var initialSession *mgo.Session

func init() {
	var err error
	initialSession, err = mgo.Dial("mongodb://admin:admin123@ds139970.mlab.com:39970/heroku_4ttnstzg")

	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {

	router := httprouter.New()

	router.POST("/api/messages", handleNewMessages)
	router.GET("/api/stats", handleGetStats)

	log.Fatal(http.ListenAndServe(":8080", router))

}

func handleNewMessages(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	log.Println(r)
}

func handleGetStats(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Write(calculateStats())
}

func addMessages(messages []message) {

	session := initialSession.Copy()
	defer session.Close()

	collection := session.DB("heroku_4ttnstzg").C("messages")

	for _, message := range messages {
		collection.Insert(message)
	}
}

func calculateStats() []byte {

	session := initialSession.Copy()
	defer session.Close()

	collection := session.DB("heroku_4ttnstzg").C("messages")

	pipeline := []bson.M{
		{"$match": bson.M{}},
		{"$group": bson.M{"_id": "$author", "count": bson.M{"$sum": 1}, "longest_message": bson.M{"$max": "$message"}}},
	}

	var result []bson.M
	err := collection.Pipe(pipeline).All(&result)

	if err != nil {
		log.Println(err.Error())
	}

	resultJson, err := json.Marshal(result)
	return resultJson

}
