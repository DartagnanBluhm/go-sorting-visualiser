package handler

import (
	"encoding/json"
	"govisualiser/api/sorting"
	"govisualiser/api/util"
	"io/ioutil"
	"log"
	"net/http"
)

const addr = "localhost:8000"

func StartVisualiser() {
	log.Printf("Listening on %s\n", addr)
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
	r.Header.Set("Content-Type", "application/json")
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(`{"status":"error","message":"error reading request body -> ` + err.Error() + `"}`))
		return
	}
	arr, algorithm, err := util.ConvertJSONtoArray(data)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`{"status":"error","message":"error converting json to array -> ` + err.Error() + `"}`))
		return
	}
	changes, err := sorting.Sort(arr, algorithm)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`{"status":"error","message":"error sorting request -> ` + err.Error() + `"}`))
		return
	}
	resp, err := json.Marshal(changes)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`{"status":"error","message":"error marshalling sorted array -> ` + err.Error() + `"}`))
		return
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write(resp)
}
