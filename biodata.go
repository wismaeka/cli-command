package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type people struct {
	id         int
	name       string
	address    string
	work       string
	motivation string
}

var peopleData = []people{
	{id: 1, name: "Ani", address: "Jakarta", work: "Software Developer", motivation: "enchance skill for new programming language"},
	{id: 2, name: "Bimo", address: "Bandung", work: "Backend Developer", motivation: "want to know basic of Golang"},
	{id: 3, name: "Carly", address: "Surabaya", work: "IT Students", motivation: "enrich programming skill"},
	{id: 4, name: "Dede", address: "Bali", work: "Freelancer", motivation: "projects requirement for using Golang is increasing"},
}

func main() {
	str, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	for _, person := range peopleData {
		if str == person.id {
			// fmt.Println(person)
			fmt.Println(person.id, "name:", person.name, "address:", person.address, "work:", person.work, "motivation:", person.motivation)
		}
	}
	if str <= 0 || str > len(peopleData) {
		fmt.Println("There is no person matching your input.")
	}
}
