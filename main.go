package main

import(
    "log"
    "net/http"
)

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        w.Write([]byte("Halo Test"))
    })

    log.Println("server running in port 8080")

    err := http.ListenAndServe(":8080", mux)
    log.Fatal(err)
}
