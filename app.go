package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/tokopedia/kol-marketplace/search"
)

type Posts struct {
	ID      int64  `json:"id,omitempty"`
	Content string `json:"content,omitempty"`
	Name    string `json:"name,omitempty"`
}

func main() {
	err := search.Init("http://10.255.13.51:9200", "kolmarketplace")
	if err != nil {
		log.Fatal("Connect Error : ", err)
		return
	}

	err = search.InsertIndex(context.Background(), "test2")
	if err != nil {
		fmt.Println(err)
	}

	post := Posts{ID: 8939359, Name: "Content Buyer 2", Content: "{}"}
	doc := search.BuildDoc(strconv.FormatInt(post.ID, 10), "test2", "users", post)

	// //example insert
	err = search.Create(context.Background(), doc)
	if err != nil {
		log.Println(err)
	}

	field := make(map[string]string)
	field["name"] = "content"
	filter := make(map[string]string)
	q := search.Query{
		Fields:  field,
		Filters: filter,
		Index:   "test2",
		Type:    "users",
		Sort:    search.Sort{Field: "id", Ascending: true},
	}

	test := []Posts{}

	result, err := search.Search(context.Background(), q, 10, 0, test)

	if err != nil {
		log.Fatal("Query Error : ", err)
	}

	//Casting the result to desired type
	fmt.Println(result.([]Posts))

}
