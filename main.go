/*
Title: main.go

Author: Troy Caro <twc17@pitt.edu>
Date Modified: 08/27/2018
Version: 0.0.2

Purpose:
	Hello World! web application
*/

package main

import (
	"context"
	"html/template"
	"io/ioutil"
	"log"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	version = "v0.0.2"
	title = "YAY!"
)

type Page struct {
	Title string
	Body []byte
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, version)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	p, _ := loadPage(title)
	renderTemplate(w, "index", p)
}

func main() {
	log.Println("Starting hello-world application")

	http.HandleFunc("/", rootHandler)

	http.HandleFunc("/health", healthHandler)

	http.HandleFunc("/version", versionHandler)

	s := http.Server{Addr: ":80"}
	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	log.Println("Shutdown received, closing...")

	s.Shutdown(context.Background())
}
