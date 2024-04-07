package main

import (
	"flag"
	"fmt"
	"log"
	"osuprogressserver/api"
	"osuprogressserver/storage"
)

func main() {

	for i := 0; i < 100; i++ {
		//fmt.Println(i++)
	}

	port := flag.String("port", ":3000", "Serverport")
	flag.Parse()
	storage, err := storage.NewSQLite("Data/test.db")
	if err != nil {
		log.Fatal(err)
	}

	s := api.NewServer(*port, storage)

	fmt.Println("Webserver Running at", *port)
	log.Fatal(s.Start())

	defer storage.DB.Close()
}
