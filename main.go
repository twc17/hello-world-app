package main

import (
	"context"
	"html/tmplate"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	version = "v0.0.1"
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


func main() {
	log.Println("Starting hello-world application")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World! - CI/CD Test\n -Troy\n")
	})

	http.HandleFunc("/health", healthHandler)

	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, version)
	})

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
