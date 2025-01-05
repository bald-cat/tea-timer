package main

import (
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"tgbot/router"
	"tgbot/telegram"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var tg = telegram.NewTelegram()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Unable to read request body", http.StatusInternalServerError)
			return
		}

		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Fatal("Error closing body")
			}
		}(r.Body)

		request, err := telegram.NewTgRequest(body)
		if err != nil {
			http.Error(w, "Unable to parse request body", http.StatusBadRequest)
			return
		}

		router.NewRouter(tg, request)
	})

	srv := &http.Server{
		Addr:    ":4000",
		Handler: mux,
	}

	log.Println("Listening on :4000")
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
