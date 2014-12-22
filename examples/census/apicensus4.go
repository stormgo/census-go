package main

import (
	"fmt"
	//"net/url"
	"os"
	"github.com/stormasm/census-go/goreq"
)

func main() {

	key := os.Getenv("CENSUS_API_KEY")

	type Item struct {
		Key string
		Get string
		For string
		In  string
	}

	item := Item {
		Key: key,
		Get: "P0010001,NAME",
		For: "county:*",
		In: "state:41",
	}

	res3, err3 := goreq.Request{
		Method:      "GET",
		ContentType: "application/json",
		Uri:         "http://api.census.gov/data/2010/sf1",
		QueryString: item,
	}.Do()

	str, _ := res3.Body.ToString()

	fmt.Println(err3)
	fmt.Println(str)

}
