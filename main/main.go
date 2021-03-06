package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/quangngotan95/go-m3u8/m3u8"
)

// URLMain - ссылка для скачивания
var URLMain string = "https://soundcloud.com/grum/under-your-skin-original-mix"

var clientIdjs string = "https://a-v2.sndcdn.com/assets/49-d7adc028-3.js"
var myclientid string = "L1Tsmo5VZ0rup3p9fjY67862DyPiWGaG"

var clienID string

func main() {
	// main1()
	// main2()
	// GetClientID1()
	GetJsList(URLMain)
}

// GetJsList получение js-файлов
func GetJsList(u string) {
	b01 := GetBody(u)
	a01 := MyReg(`<script crossorigin src=\"(https:\/\/a-v2\.sndcdn\.com\/assets\/[a-zA-Z0-9\-]+\.js)"><\/script>`, b01)
	GetClientID(a01)
	a02 := MyReg(`https:\/\/api-v2\.soundcloud\.com\/media\/soundcloud:tracks:[0-9]+\/[0-9a-zA-Z-]+\/stream\/hls`, b01)

	b03 := GetBody(a02[0][0] + "?client_id=" + clienID)
	fmt.Println(b03)
}

// GetClientID получение client_id
func GetClientID(l [][]string) {
	if len(l) > 0 {
		for _, lv1 := range l {
			if len(lv1) >= 1 {
				fmt.Println(lv1[1])
				b001 := GetBody(lv1[1])
				WriteJSToFile(b001)
				ReadJSToFile()
				if clienID != "" {
					var err = os.Remove("WriteJSToFile.txt")
					if err != nil {
						panic(err)
					}
					fmt.Println(clienID)
					return
				}
			}

		}
	}
}

// WriteJSToFile ...
func WriteJSToFile(s string) {
	f1, err := os.OpenFile("WriteJSToFile.txt", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	l, err := f1.WriteString(s)
	if err != nil {
		fmt.Println(err)
		f1.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f1.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	f1.Close()
}

// ReadJSToFile ...
func ReadJSToFile() {
	configFile, err := ioutil.ReadFile("WriteJSToFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	configLines := strings.Split(string(configFile), "\n")
	for i := 0; i < len(configLines); i++ {
		if configLines[i] != "" {
			res1 := MyReg(`client_id:\"([a-zA-Z0-9]{32})\"`, configLines[i])
			// fmt.Println(res1)
			// fmt.Println(len(res1))
			if len(res1) >= 1 {
				clienID = res1[0][1]
				break
			}

		}
	}
}
func main1() {
	mb := GetBody("https://soundcloud.com/grum/under-your-skin-original-mix")
	ma := MyReg(`<script crossorigin src=\"(https:\/\/a-v2\.sndcdn\.com\/assets\/[a-zA-Z0-9\-]+\.js)"><\/script>`, mb)
	ii := 0
	for _, v1 := range ma {
		fmt.Println(v1[1])
		mb2 := GetBody(v1[1])
		fmt.Println("--------")
		ma2 := MyReg(`client_id: \"([a-zA-Z0-9]{32})\"`, mb2)
		fmt.Println(ma2)
		ii++
		fmt.Println(ii)
	}

}
func main2() {
	// fmt.Println(my_client_id)
	mb2 := GetBody("https://soundcloud.com/grum/under-your-skin-original-mix")
	ma2 := MyReg(`https:\/\/api-v2\.soundcloud\.com\/media\/soundcloud:tracks:[0-9]+\/[0-9a-zA-Z-]+\/stream\/hls`, mb2)
	mb3 := GetBody(ma2[0][0] + "?client_id=" + myclientid)
	var myjson map[string]string
	if err := json.Unmarshal([]byte(mb3), &myjson); err != nil {
		panic(err)
	}
	mb4 := GetBody(myjson["url"])
	// fmt.Println(mb4)
	playlist, _ := m3u8.ReadString(mb4)
	// fmt.Println(cmp.Diff("Hello World", "Hello Go"))
	// fmt.Println(playlist.Items[6])
	// ss := strings.Split(playlist.Items[2].String(), "\n")
	// fmt.Printf("%q\n", ss[1])
	// mb5 := GetBody(ss[1])
	for _, pv := range playlist.Items {
		ss := strings.Split(pv.String(), "\n")
		// WriteToFile(ss[1])
		// fmt.Printf("%q\n", ss[1])
		// WriteToFile("222\n")
		WriteToFile(GetBody(ss[1]))
		// fmt.Printf("%q\n", GetBody(ss[1]))
	}

	// mb5 := GetBody("https://cf-hls-media.sndcdn.com/media/3671770/3831429/YSPVFcL40sak.128.mp3?Policy=eyJTdGF0ZW1lbnQiOlt7IlJlc291cmNlIjoiKjovL2NmLWhscy1tZWRpYS5zbmRjZG4uY29tL21lZGlhLyovKi9ZU1BWRmNMNDBzYWsuMTI4Lm1wMyIsIkNvbmRpdGlvbiI6eyJEYXRlTGVzc1RoYW4iOnsiQVdTOkVwb2NoVGltZSI6MTU3OTE3NjU5Mn19fV19&Signature=OdW-~cv-mV1UxzpEWWuj-YbX9sAvQd58ovB-LC0-Wa-K7EwhvVvOzU~JchEMrL49Yd8XMtI10boRhcucJppGNgbeFNdDS9Mq1mt7Ok-3JfQjbtbyY-xrOkwdEjQ9bTtr0zac9WhC2ILnFb97Xcen6IqDX1YjRK3VBVtoln3exMxnyexpNHYjMAmUF5bIshVSPMqL6a~a5NjEt6jK1bRsxfYIPpRKjrKz3AWnwCwRhO8U2MdCHg96YaWzSecF5OhO5k7WcU4L0DEFjJfxS3DvMxHTdc6KwpS4IYTwuc1e~4tEybT6J59HIsjQYXxHm7R-HrpDiEjwcN39kS1stHR9jA__&Key-Pair-Id=APKAI6TU7MMXM5DG6EPQ")
	// fmt.Println(mb5)

}

//GetClientID1 ...gf
func GetClientID1() {
	jsString := GetBody(clientIdjs)
	fmt.Println()
	f1, err := os.OpenFile("testJS1.txt", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	l, err := f1.WriteString(jsString)
	if err != nil {
		fmt.Println(err)
		f1.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f1.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	configFile, err := ioutil.ReadFile("testJS1.txt")
	if err != nil {
		log.Fatal(err)
	}
	configLines := strings.Split(string(configFile), "\n")
	for i := 0; i < len(configLines); i++ {
		if configLines[i] != "" {
			// fmt.Println(configLines[i])
			res := MyReg(`client_id:\"([a-zA-Z0-9]{32})\"`, configLines[i])
			fmt.Println(res)
			fmt.Println(len(res))
			if len(res) >= 1 {
				return
			}

		}
	}

}

// WriteToFile ...
func WriteToFile(s string) {
	// f, err := os.Create("test.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	f, err := os.OpenFile("test1.mp3", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	l, err := f.WriteString(s)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
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
	return r.FindAllStringSubmatch(b, -1)
}
