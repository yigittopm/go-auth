package utils

import (
	"log"
	"net/http"
	"os"
)

func Logger(h http.HandlerFunc) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
		log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	}

	return http.HandlerFunc(fn)
}
