package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	mirrors "restfulServices/mirrorFinder/data"
	"time"
)

type response struct {
	FasterURL string        `json:"fastest_url"`
	Latency   time.Duration `json:"latency"`
}

func main() {
	http.HandleFunc("/fastest-mirror", handleFunHandler)

	port := ":8000"
	server := &http.Server{
		Addr:           port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("Starting srrver on port %sn", port)
	log.Fatal(server.ListenAndServe())

}

func handleFunHandler(w http.ResponseWriter, r *http.Request) {
	res := findFastest(mirrors.MirrorList)
	resJson, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resJson)
}

func findFastest(urls []string) response {
	urlChan := make(chan string)
	latencyChan := make(chan time.Duration)

	for _, url := range urls {
		mirrorUrl := url

		go func() {
			start := time.Now()
			_, err := http.Get(mirrorUrl + "/README")
			latency := time.Now().Sub(start) / time.Millisecond

			if err == nil {
				urlChan <- mirrorUrl
				latencyChan <- latency
			}
			log.Printf("Got the best mirror: %s with latency : %s", mirrorUrl, latency)
		}()
	}

	return response{<-urlChan, <-latencyChan}
}
