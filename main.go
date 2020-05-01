package main

import (
	"net"
	"net/http"

	"log"
)

var port string = "8080"
var frontHost = net.JoinHostPort("front.lab.test.com", port)
var labHost = net.JoinHostPort("lab.test.com", port)
var graphsHost = net.JoinHostPort("graphs.lab.test.com", port)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request for %q, cookies: %v\n", r.Host, r.Cookies())
	switch r.Host {
	case frontHost:
		frontendHandler(w, r)
		return
	case labHost:
		labHandler(w, r)
		return
	case graphsHost:
		graphsHandler(w, r)
		return
	}
}

func frontendHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Frontend")
	http.SetCookie(w, &http.Cookie{
		Name:   "FrontendToLab",
		Value:  "A cookie from the frontend to lab",
		Domain: "lab.test.com",
	})
	http.SetCookie(w, &http.Cookie{
		Name:  "Frontend",
		Value: "A cookie from the frontend",
	})
	http.SetCookie(w, &http.Cookie{
		Name:   "FrontendToGraphs",
		Value:  "A cookie from the frontend to graphs",
		Domain: "graphs.lab.test.com",
	})
	w.WriteHeader(http.StatusOK)
}

func labHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Lab")
	http.SetCookie(w, &http.Cookie{
		Name:  "Lab",
		Value: "A cookie from the lab",
	})
	http.SetCookie(w, &http.Cookie{
		Name:   "LabExplicit",
		Value:  "A cookie from the lab (Explicit)",
		Domain: "lab.test.com",
	})
	http.SetCookie(w, &http.Cookie{
		Name:   "LabToFrontend",
		Value:  "A cookie from the lab to the frontend",
		Domain: "front.lab.test.com",
	})
	w.WriteHeader(http.StatusOK)
}

func graphsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Graphs")
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.ListenAndServe(net.JoinHostPort("", port), http.HandlerFunc(handler))
}
