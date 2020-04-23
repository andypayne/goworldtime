package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/andypayne/goworldtime/controllers"
)

func main() {
	fmt.Println("Go World Time Web Service")
	port := flag.Int("port", 4040, "The port to listen on")
	flag.Parse()
	err := startWebService(*port)
	fmt.Println(err)
}

func startWebService(port int) error {
	fmt.Println("Starting service on port", port)
	controllers.RegisterControllers()
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	return nil
}
