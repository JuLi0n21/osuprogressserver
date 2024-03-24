package main

import (
	"flag"
	"fmt"
	"log"
	"osuprogressserver/api"
	"osuprogressserver/storage"
)

func main() {
	port := flag.String("port", ":3000", "Serverport")
	flag.Parse()
	storage, err := storage.NewSQLite("test.db")
	if err != nil {
		log.Fatal(err)
	}

	s := api.NewServer(*port, storage)

	fmt.Println("Webserver Running at", *port)
	log.Fatal(s.Start())

	defer storage.DB.Close()
}
