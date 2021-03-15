package handler

import (
	"encoding/json"
	"fmt"
	"govisualiser/api/sorting"
	"govisualiser/api/util"
	"io/ioutil"
	"log"
	"net/http"
)

const addr = "localhost:8000"

func StartVisualiser() {
	fmt.Printf("Listening on %s\n", addr)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	http.HandleFunc("/home", func(rw http.ResponseWriter, r *http.Request) {
		http.ServeFile(rw, r, "./public/assets/index.html")
	})
	http.HandleFunc("/api/v1/sort", sortArray)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func sortArray(rw http.ResponseWriter, r *http.Request) {
	r.Header.Set("Access-Control-Allow-Origin", "*")
	r.Header.Set("Content-Type", "multipart/form-data")
	if r.Method == "OPTIONS" {
		rw.WriteHeader(http.StatusOK)
		return
	}
	if r.Method != "POST" {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		r.Header.Set("Content-Type", "plain/text")
		rw.Write([]byte(err.Error()))
		return
	}
	arr, algorithm, err := util.ConvertJSONtoArray(data)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		r.Header.Set("Content-Type", "plain/text")
		rw.Write([]byte(err.Error()))
		return
	}
	changes, err := sorting.Sort(arr, algorithm)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		r.Header.Set("Content-Type", "plain/text")
		rw.Write([]byte(err.Error()))
		return
	}
	resp, err := json.Marshal(changes)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		r.Header.Set("Content-Type", "plain/text")
		rw.Write([]byte(err.Error()))
		return
	}
	r.Header.Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(resp)
}
