package main

import (
	"fmt"
	"log"

	"github.com/tokopedia/kol-marketplace/search"
)

type Posts struct {
	ID      int64  `json:"id"`
	Content string `json:"content"`
	Name 	string `json:"name"`
}

func main() {
	err := search.Init("http://172.21.44.18:9200")
	if err != nil {
		log.Fatal("Connect Error : ", err)
		return
	}
	field := make(map[string]string)
	field["content"] = "upin"
	filter := make(map[string]string)
	q := search.Query{
		Fields:  field,
		Filters: filter,
		Index:   "kol",
		Type:    "posts",
		Sort:    search.Sort{Field: "id", Ascending: true},
	}

	test := []Posts{}

	result, err := search.Search(q, 10, 0, test)
	if err != nil {
		log.Fatal("Query Error : ", err)
	}

	fmt.Println(result)

}
