package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nathanielhall/cloud-native-go/config"
)
 

func main() {
	appConfig := config.AppConfig()
	mux := http.NewServeMux()
	mux.HandleFunc("/", Hello)	
	
	address := fmt.Sprintf(":%d", appConfig.Server.Port)
	log.Printf("Starting server 2 %s\n", address)

	s := &http.Server{
		Addr:         address,
		Handler:      mux,
		ReadTimeout:  appConfig.Server.TimeoutRead,
		WriteTimeout: appConfig.Server.TimeoutWrite,
		IdleTimeout:  appConfig.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}
 
func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}






// log.Println("Starting server :8080")
// s := &http.Server{
// 	Addr:         ":8080",
// 	Handler:      mux,
// 	ReadTimeout:  30 * time.Second,
// 	WriteTimeout: 30 * time.Second,
// 	IdleTimeout:  120 * time.Second,
// }
// if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
// 	log.Fatal("Server startup failed")
// }
// }
// func Greet(w http.ResponseWriter, r *http.Request) {
// fmt.Fprintf(w, "Hello World!")
// }