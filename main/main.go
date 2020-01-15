package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

}

func mRequest() {
	//resp, err := http.Get("https://httpbin.org/get")
	resp, err := http.Get("")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))

}
