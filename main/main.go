package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"
)

func main() {
	// mRequest()
	mb := GetBody("https://soundcloud.com/grum/under-your-skin-original-mix")
	ma := MyReg(`<script crossorigin src=\"(https:\/\/a-v2\.sndcdn\.com\/assets\/[a-zA-Z0-9\-]+\.js)"><\/script>`, mb)
	//myMap := GetBody("https://soundcloud.com/grum/under-your-skin-original-mix")
	ii := 0
	for _, v1 := range ma {
		// for _, v2 := range v1 {
		// 	fmt.Println(v2)
		// }
		fmt.Println(v1[1])
		mb2 := GetBody(v1[1])
		fmt.Println("--------")
		ma2 := MyReg(`client_id: \"([a-zA-Z0-9]{32})\"`, mb2)
		fmt.Println(ma2)
		ii++
		fmt.Println(ii)
		// ma2 := MyReg(`client_id: \"([a-zA-Z0-9]{32})\"`,mb2)
		// if len(ma2)>0{
		// 	fmt.Println("--------")
		// 	fmt.Println(ma2)
		// }
		// for _, v2_1 := range ma2 {

		// }
	}

}

func mRequest() {
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get("https://soundcloud.com/grum/under-your-skin-original-mix")
	// resp, err := http.Get("https://a-v2.sndcdn.com/assets/0-b93ede05-3.js")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	re := regexp.MustCompile(`<script crossorigin src=\"https:\/\/a-v2\.sndcdn\.com\/assets\/([a-zA-Z0-9\-]+\.js)"><\/script>`)

	fmt.Printf("%q\n", re.FindAllStringSubmatch(string(body), -1))

}

//GetBody ...
func GetBody(u string) string {
	timeout := time.Duration(100 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(u)
	MyErr(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	MyErr(err)
	// return MyReg(`<script crossorigin src=\"https:\/\/a-v2\.sndcdn\.com\/assets\/([a-zA-Z0-9\-]+\.js)"><\/script>`, string(body))
	// return MyReg(`<script crossorigin src=\"(https:\/\/a-v2\.sndcdn\.com\/assets\/[a-zA-Z0-9\-]+\.js)"><\/script>`, string(body))
	return string(body)
}

//MyErr ...
func MyErr(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

//MyReg ...
func MyReg(re string, b string) [][]string {
	r := regexp.MustCompile(re)
	//fmt.Printf("%q\n", r.FindAllStringSubmatch(b, -1))
	return r.FindAllStringSubmatch(b, -1)
}
