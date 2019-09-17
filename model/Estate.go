package model

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"sync"
	"time"
)

const RequestTimeout = 5 * time.Second

type Estate struct {
	Title    string
	Link     string
	Phone    string
	LinkChan chan string
}

func (e *Estate) String() string {
	str := ""
	str += "Link: " + e.Link + "\n"
	str += "Title: " + e.Title + "\n"
	str += "Phone: " + e.Phone + "\n"

	return str
}

func (e *Estate) getPhone() {
	client := http.Client{
		Timeout: RequestTimeout,
	}
	req, _ := http.NewRequest("GET", e.Link, nil)
	req.Header.Set("Connection", "Keep-Alive")
	req.Header.Set("User-Agent", "Mozilla/5.0")
	res, err := client.Do(req)
	if err != nil {
		log.Printf("%s request timeout", e.Link)
		return
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return
	}
	re, _ := regexp.Compile(`(?ms)href="tel:(.*?)"`)
	m := re.FindSubmatch(body)
	if len(m) > 0 {
		e.Phone = string(m[1])
	}
	log.Println(e)
}

func collectEstate(body []byte) []*Estate {
	re, _ := regexp.Compile(`<a class="mbn-image" title="(.*?) href="(.*?)">`)
	listEstate := make([]*Estate, 0)
	for _, m := range re.FindAllSubmatch(body, -1) {
		listEstate = append(listEstate, &Estate{
			Title:    string(m[1]),
			Link:     string(m[2]),
			LinkChan: make(chan string),
		})
	}

	return listEstate
}

func getContentPage() {
	url := "https://muaban.net/ban-dat-da-nang-l15-c31?cp="
	num := 50
	client := http.Client{
		Timeout: RequestTimeout,
	}
	var wg sync.WaitGroup
	listEstate := make([]*Estate, 0)
	for i := 1; i <= num; i++ {
		wg.Add(1)
		link := url + strconv.Itoa(i)
		go func() {
			defer wg.Done()
			req, _ := http.NewRequest("GET", link, nil)
			req.Header.Set("Connection", "Keep-Alive")
			req.Header.Set("User-Agent", "Mozilla/5.0")
			res, err := client.Do(req)
			if err != nil {
				log.Printf("%s request timeout", link)
				return
			}
			body, err := ioutil.ReadAll(res.Body)
			res.Body.Close()
			if err != nil {
				return
			}
			es := collectEstate(body)
			for _, e := range es {
				if e != nil {
					//wg.Add(1)
					e.getPhone()
				}
			}
			listEstate = append(listEstate, es...)
		}()
	}
	wg.Wait()
	//log.Println(listEstate)
	log.Printf("total: %d", len(listEstate))
}
