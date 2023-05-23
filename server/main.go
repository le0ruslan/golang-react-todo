package main

import (
	"fmt"
	"githib.com/le0ruslan/golang-react-todo/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("starting server on port 9000....")

	log.Fatal(http.ListenAndServe(":9000", r))
}
