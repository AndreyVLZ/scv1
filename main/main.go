package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	mRequest()
}

func mRequest() {
	//resp, err := http.Get("https://httpbin.org/get")
	// resp, err := http.Get("https://soundcloud.com/grum/under-your-skin-original-mix")
	resp, err := http.Get("https://a-v2.sndcdn.com/assets/0-b93ede05-3.js")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("HI " + string(body))

}
