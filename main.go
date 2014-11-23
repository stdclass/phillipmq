package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

var messages = []string{};

const (
    CONN_HOST = "localhost"
    CONN_PORT = "4390"
)

func peek(w http.ResponseWriter, r *http.Request) {
    
    pool := r.URL.Path[6:]
    
 	fmt.Fprintf(w, "peek %s!", channel)
}

func push(w http.ResponseWriter, r *http.Request) {
    
    data, err := ioutil.ReadAll(r.Body)
    
    message := string(data)
    
    pool := r.URL.Path[6:]
    
    if err != nil {
        fmt.Println("Error reading: ", err.Error())
    }
    
    messages = append(messages, message)
    
    fmt.Fprintf(w, "push to pool %s!", channel)
    fmt.Fprintf(w, "messagse %s!", messages)
    
    
}

func main() {
    http.HandleFunc("/peek/", peek)
    http.HandleFunc("/push/", push)
    http.ListenAndServe(CONN_HOST + ":" + CONN_PORT, nil)
}