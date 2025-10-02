package main

import (
	 "context"
	 "fmt"
	 "log"
	 "net/http"

	 "golang.ngrok.com/ngrok/v2"
)

func main() {
	 l, err := ngrok.Listen(context.Background())
	 if err != nil {
			 log.Fatal(err)
	 }
	 fmt.Println("endpoint url: ", l.URL())
	 http.Serve(l, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			 fmt.Fprintln(w, "Hello from your ngrok-delivered Go app!")
	 }))
}
