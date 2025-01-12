package database

import (
	"fmt"
	"log"
	"testing"
)

func TestGetPoetry(t *testing.T) {

	InitDB()
	ID := "5b8b9572e116fb3714e6faba"
	fetchedResponse, err := FetchFromDB(ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Fetched data: %+v\n", fetchedResponse)

	respondList, err := FetchRandomFromDB(DB, 5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("get List:")
	for _, v := range respondList {
		fmt.Printf("Fetched data: %+v\n", v)
	}
}
