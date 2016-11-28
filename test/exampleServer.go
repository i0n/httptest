package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/namsral/flag"
)

var (
	port = flag.String("port", "8080", "The port that the health check blocks on.")
)

// Message Simple message struct
type Message struct {
	Body string
}

func main() {
	flag.Parse()
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		var message = "\"I'm pretty sure there's a lot more to life than being really, really, ridiculously good looking.\" - Derek Zoolander\n"
		if r.Header.Get("Content-Type") == "application/json" {
			w.Header().Set("Content-Type", "application/json")
			json, err := json.Marshal(&Message{Body: message})
			if err != nil {
				panic(err)
			}
			w.Write([]byte(json))
		} else {
			w.Write([]byte(message))
		}
	})
	http.HandleFunc("/yummydata", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") == "application/json" {
			w.Header().Set("Content-Type", "application/json")
			if r.Method == "POST" {
				defer r.Body.Close()
				body, err := ioutil.ReadAll(r.Body)
				if err != nil {
					panic(err)
				}
				message := &Message{}
				err = json.Unmarshal(body, message)
				if err != nil {
					panic(err)
				}
				json, err := json.Marshal(message)
				if err != nil {
					panic(err)
				}
				w.WriteHeader(http.StatusCreated)
				w.Write([]byte(json))
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})

	log.Println("Listening on 0.0.0.0:" + *port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+*port, nil))
}
