package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type login_object struct {
	Email        string
	Password     string
	OneTimeToken string
}

func postlogin(rw http.ResponseWriter, req *http.Request) {
	(rw).Header().Set("Access-Control-Allow-Origin", "*")
	if req.URL.Path != "/api/login" {
		http.Error(rw, "404 not found.", http.StatusNotFound)
		return
	}
	decoder := json.NewDecoder(req.Body)
	var postBody login_object
	err := decoder.Decode(&postBody)
	if err != nil {
		panic(err)
	}

	log.Println(postBody.Email)
	if (postBody.Email == "c137@onecause.com") && (postBody.Password == "#th@nH@rm#y#r!$100%D0p#") {
		log.Println(postBody.Email)
		log.Println(postBody.Password)
		log.Println(postBody.OneTimeToken)
		rw.WriteHeader(http.StatusOK)
		resp := make(map[string]string)
		resp["message"] = "Status OK"
		jsonResp, err := json.Marshal(resp)
		rw.Write(jsonResp)
		if err != nil {
			panic(err)
		}
		return
	} else {
		log.Println("login validation failed")
		rw.WriteHeader(http.StatusUnauthorized)
		resp := make(map[string]string)
		resp["message"] = "Status not okay"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			panic(err)
		}
		rw.Write(jsonResp)
		return
	}
}

func main() {
	http.HandleFunc("/api/login", postlogin)

	fmt.Printf("Starting server post login on port 5000...\n")
	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Fatal(err)
	}
}
