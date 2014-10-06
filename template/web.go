package main

import (
	"fmt"
	"net/http"
	"os"
	"log"
	"runtime"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name string
	Location string
}

func main() {
	http.HandleFunc("/", hello)
	bind := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	fmt.Printf("listening on %s...", bind)
	err := http.ListenAndServe(bind, nil)
	if err != nil {
		panic(err)
	}
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
	err = c.Insert(&Person{"El Chamakito", "Del Varrio"},
		&Person{"El Chamakito", "De las Chelas"})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"location": "Chelas"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Json:", result)
	//	fmt.Fprintf(res, "Welcome to g4capi, powered by GO %s", runtime.Version())
	fmt.Fprintf(res, "Welcome to %s universe, powered by %s", result, runtime.Version())

}
