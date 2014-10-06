package main

import (
	"fmt"
	"net/http"
	"log"
	"runtime"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/go-martini/martini"
)

type Person struct {
	Name string
	Location string
}

func main() {
	m := martini.Classic()
	m.Get("/", hello)
	m.Run()
}

func hello(res http.ResponseWriter, req *http.Request) {
	session, err := mgo.Dial("mongodb://localhost/test")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("g3").C("people")
	err = c.Insert(&Person{"Golang Giraltar Community", "Gibraltar"},
		&Person{"Golang Community", "UK"})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"location": "Gibraltar"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Json:", result)
	fmt.Fprintf(res, "Welcome, %s, powered by %s", result, runtime.Version())

}
