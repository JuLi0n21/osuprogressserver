package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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
	mockdata := flag.Bool("mock", false, "Generate Fake Data on startup")
	flag.Parse()
	storage, err := storage.NewSQLite("Sqlite.db")
	if err != nil {
		log.Fatal(err)
	}

	if *mockdata {
		fmt.Println("Creating Fake Data")
		storage.MockScores(100000)
	}

	env := os.Getenv("ENV")

	s := api.NewServer(*port, storage, env)

	fmt.Println("Webserver Running at", *port)
	log.Fatal(s.Start())

	defer storage.DB.Close()
}
