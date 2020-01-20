package main

import "fmt"

// URLMain - ссылка для скачивания
var URLMain string = "https://soundcloud.com/grum/under-your-skin-original-mix"

// Mp3 ...
type Mp3 struct {
	URL string
}

// MyMp3 ...
var MyMp3 = Mp3{
	URL: "https://soundcloud.com/grum/under-your-skin-original-mix",
}

func main() {
	fmt.Println(MyMp3.URL)
}
