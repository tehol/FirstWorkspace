package rest

import (
	"io"
	"log"
	"net/http"
)

func StartRestServer(listenAddr string) {
	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/createuser"] = createUser
	mux["/authenticate"] = authenticate
	err := http.ListenAndServeTLS(listenAddr, "cacert.pem", "ca.key", &restHandler{})
	if err != nil {
		log.Fatal(err)
	}

}

func authenticate(rw http.ResponseWriter, request *http.Request) {
	io.WriteString(rw, "Authenticate is called.")
}

func createUser(rw http.ResponseWriter, request *http.Request) {

	io.WriteString(rw, "Create user is called.")

}

var mux map[string]func(http.ResponseWriter, *http.Request)

type restHandler struct{}

func (*restHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}

	io.WriteString(w, "My server: "+r.URL.String())
}
