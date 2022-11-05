package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type login_object struct {
	email        string
	password     string
	oneTimeToken string
}

func postlogin(rw http.ResponseWriter, req *http.Request) {
	(rw).Header().Set("Access-Control-Allow-Origin", "*")
	if req.URL.Path != "/api/login" {
		http.Error(rw, "404 not found.", http.StatusNotFound)
		return
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))

	var postBody login_object
	err = json.Unmarshal(body, &postBody)
	if err != nil {
		panic(err)
	}
	log.Println(postBody.email)
	log.Println(postBody.password)
	log.Println(postBody.oneTimeToken)
}

func main() {
	http.HandleFunc("/api/login", postlogin)

	fmt.Printf("Starting server post login on port 5000...\n")
	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Fatal(err)
	}
}
