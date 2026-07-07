package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello Modules")
	// go get -u github.com/gorilla/mux

	// go env
	// find GOPATH='/Users/adi/go', open in terminal and run the command
	// ❯ cd /Users/adi/go
	// ❯ ls
	// bin pkg
	// ❯ cd pkg
	// ❯ ls
	// mod   sumdb
	// ❯ cd mod
	// ❯ ls
	// cache  <---- ⭐ all the packages downloaded by go get will be stored here ⭐
	// cel.dev
	// cloud.google.com
	// github.com
	// go.opentelemetry.io
	// go.opentelemetry.io@v0.1.0
	// go.starlark.net@v0.0.0-20231101134539-556fd59b42f6
	// go.uber.org
	// go.yaml.in
	// golang.org
	// gonum.org
	// google.golang.org
	// gopkg.in
	// gotest.tools
	// gotest.tools@v2.2.0+incompatible
	// honnef.co
	// k8s.io
	// mvdan.cc
	// sigs.k8s.io

	/*
		❯ cd cache/download
		❯ ls
		cel.dev             go.starlark.net     google.golang.org   mvdan.cc
		cloud.google.com    go.uber.org         gopkg.in            sigs.k8s.io
		github.com          go.yaml.in          gotest.tools        sumdb
		go.opencensus.io    golang.org          honnef.co
		go.opentelemetry.io gonum.org           k8s.io
		❯ cd github.com
		....      gorilla     ...
	*/
	greeter()

	/*
	    go mod graph    When you run go mod graph, Go prints out the module requirement graph 
		for your entire project. 
		It maps out exactly how every single package in your project is connected.

		go mod vendor    copies all the dependencies of the module to the vendor directory
						  it is just like node_modules in nodejs

		go run main.go compiles and runs the local source in main.go.
		If the required module is not already in your local module cache, Go may download that dependency once from the internet.
		After that, it usually reuses the cached copy on your machine.
		It does not fetch your own code from the internet.


		Copy all dependencies needed by this module into a local vendor directory.
		That creates a fully local copy of third-party code for builds.


		go run -mod=vendor main.go means:

		Run the program using the vendor directory only.
		Go will not try to resolve dependencies from the network.
		If something is missing from vendor, the build fails instead of downloading it.


		go run -mod=vendor main.go 
	*/

}
func greeter() {
	fmt.Println("Hello from Greeter function")

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Welcome to the HomePage!")
	// fmt.Println("Endpoint Hit: homePage")
	w.Write([]byte("<h1>Welcome to the HomePage!</h1>"))
}
