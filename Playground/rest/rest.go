package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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

type User struct {
	Username string `Json: Username string`
	Password string `Json: Password string`
}

func createUser(rw http.ResponseWriter, request *http.Request) {
	fmt.Println(request.RemoteAddr)
	fmt.Println(request.Method)
	if request.Method != "POST" {
		io.WriteString(rw, "HTTP Method is not supported.")
		return
	}

	body, err := ioutil.ReadAll(io.LimitReader(request.Body, 512))
	if err != nil {
		panic(err)
	}
	if err := request.Body.Close(); err != nil {
		panic(err)
	}
	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
		rw.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(rw).Encode(err); err != nil {
			panic(err)
		}
	}
	fmt.Print(user)
	rw.WriteHeader(http.StatusOK)
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
