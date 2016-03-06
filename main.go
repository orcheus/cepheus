package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"./core"

	"github.com/gorilla/mux"
)

var (
	cm = core.NewCtxManager()
)

// ****** POGO ***********

// ****** Data ***********

// ****** RET API ***********

// register a new context
func handleRegisterContext(res http.ResponseWriter, req *http.Request) {
	r := new(core.RegisterContextRequest)
	err := json.NewDecoder(req.Body).Decode(r)
	if err != nil {
		res.WriteHeader(401)
		fmt.Fprintf(res, "bad input")
		log.Printf("error : %v", err)
	} else {
		fmt.Fprintf(res, "register context %+v", r)

	}
}

func handleGetContext(res http.ResponseWriter, req *http.Request) {
	v := mux.Vars(req)
	fmt.Fprintf(res, "get context %v", v["EntityID"])
}

func initRouter(r *mux.Router) {
	r.HandleFunc("/registerContext", handleRegisterContext)
	r.HandleFunc("/discoverContextAvailability", handleRegisterContext)
	r.HandleFunc("/subscribeContextAvailability", handleRegisterContext)
	r.HandleFunc("/updateContextAvailabilitySubscription", handleRegisterContext)
	r.HandleFunc("/unsubscribeContextAvailability", handleRegisterContext)

	r.HandleFunc("/contextEntities/{EntityID}", handleGetContext)
	r.HandleFunc("/contextEntities/{EntityID}/attributes", handleRegisterContext)
	r.HandleFunc("/contextEntities/{EntityID}/attributes/{attributeName}", handleRegisterContext)
	r.HandleFunc("/contextEntities/{EntityID}/attributeDomains/{attributeDomainName}", handleRegisterContext)
	r.HandleFunc("/contextEntityTypes/{tyeName}", handleRegisterContext)
	r.HandleFunc("/contextEntityTypes/{typeName}/attributes", handleRegisterContext)
	r.HandleFunc("/contextEntityTypes/{typeName}/attributes/{attributeName}", handleRegisterContext)
	r.HandleFunc("/contextEntityTypes/{typeName}/attributeDomains/{attributeDomainName}", handleRegisterContext)
	r.HandleFunc("/contextAvailabilitySubscriptions", handleRegisterContext)
	r.HandleFunc("/contextAvailabilitySubscriptions/{subscriptionID}", handleRegisterContext)

}

var (
	port = 8080
)

func main() {

	log.Println("                _")
	log.Println("               | |")
	log.Println("  ___ ___ _ __ | |__   ___ _   _ ___")
	log.Println(" / __/ _ \\ '_ \\| '_ \\ / _ \\ | | / __|")
	log.Println("| (_|  __/ |_) | | | |  __/ |_| \\__ \\")
	log.Println(" \\___\\___| .__/|_| |_|\\___|\\__,_|___/")
	log.Println("         | |")
	log.Println("         |_|")

	log.Println("starting...")
	defer log.Println("...stopping.")

	r := mux.NewRouter()

	initRouter(r)

	log.Printf("listening on port %d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), r)

	if err != nil {
		log.Fatal(err)
	}

}
