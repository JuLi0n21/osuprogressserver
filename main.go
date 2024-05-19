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

	port := flag.String("port", ":4000", "Serverport")
	flag.Parse()
	storage, err := storage.NewSQLite("Sqlite.db")
	if err != nil {
		log.Fatal(err)
	}

	storage.MockScores(100)
	s := api.NewServer(*port, storage)

	fmt.Println("Webserver Running at", *port)
	log.Fatal(s.Start())

	defer storage.DB.Close()
}
