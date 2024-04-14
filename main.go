package main

import (
	"flag"
	"fmt"
	"log"
	"osuprogressserver/api"
	"osuprogressserver/storage"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic("uhh no .env")
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
