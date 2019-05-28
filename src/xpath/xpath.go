package xpath

import (
	// xmlpath "gopkg.in/xmlpath.v2"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type HrefList struct {
	name string
	url  string
}

func Xpath() {
	htmlContext := Request("https://mirrors.xtom.com/debian/dists/")
	fmt.Println(htmlContext)
	// xmlpath.Parse(htmlContext)

}

func Request(url string) string {
	client := http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.11; rv:43.0) Gecko/20100101 Firefox/43.0")
	response, err := client.Do(request)
	if err != nil {
		log.Println("http.NewRequest(\"GET\", url, nil)")
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return string(body)
}
