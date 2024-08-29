package main

/*
#include <stdio.h>
#include <stdlib.h>

char* hello() {
    char* message = "Hello World!\n";
    return message;
}
*/
import "C"

import (
    "log"
    "time"
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"
)

type Response struct {
    Message string `json:"message"`
}

func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        next.ServeHTTP(w, r)

        log.Printf(
            "%s %s %s %s",
            r.Method,
            r.RequestURI,
            r.RemoteAddr,
            time.Since(start),
        )
    })
}

func main() {
    r := mux.NewRouter()

    corsOptions := handlers.CORS(
        handlers.AllowedOrigins([]string{"http://localhost"}),
        handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
        handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
        handlers.ExposedHeaders([]string{"X-My-Custom-Header"}),
        handlers.AllowCredentials(),
        handlers.MaxAge(3600),
    )
    r.Use(corsOptions)
    r.Use(LoggingMiddleware)

    r.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
        message := C.GoString(C.hello())
        response := Response{Message: message}
        json.NewEncoder(w).Encode(response)
    }).Methods("GET")

    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}