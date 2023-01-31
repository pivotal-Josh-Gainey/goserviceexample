package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {

//	fmt.Printf("A SERVICE/JOB/PROCESS CAN BE CODED TO DO ANYTHING")

	http.HandleFunc("/", handler)
	http.HandleFunc("/get-random-number", handlerRandomNumber)

	http.ListenAndServe(":"+os.Getenv("TCPPORT"), nil)

}


func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello World!")
}


func handlerRandomNumber(w http.ResponseWriter, r *http.Request){

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	number := r1.Intn(100)
	answer := strconv.Itoa(number)

	fmt.Fprintf(w, answer)
}



