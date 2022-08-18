package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func get(url string, i int) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	_ = body
	log.Printf("finished get %v ", i)
}

func main() {

	var wg sync.WaitGroup
	for true {
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func(i int) {
				get("https://www.test.com/", i)
				wg.Done()
			}(i)
		}
		wg.Wait()
	}
}
