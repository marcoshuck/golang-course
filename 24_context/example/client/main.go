package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	//ctx := context.Background()
	//
	//ctx = context.WithValue(ctx, "request_id", "1234")
	//
	//req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:1234", nil)
	//if err != nil {
	//	log.Fatalln("Error:", err)
	//}
	//
	//res, err := http.DefaultClient.Do(req)
	//defer res.Body.Close()

	res, err := http.Get("http://localhost:1234")

	if err != nil {
		log.Fatalln("Error:", err)
	}

	if res.StatusCode != http.StatusOK {
		log.Fatalln("Status error:", res.StatusCode, res.Status)
	}

	if _, err = io.Copy(os.Stdout, res.Body); err != nil {
		log.Fatalln("Error while copying to stdout:", err)
	}

	log.Println("Success")
}
