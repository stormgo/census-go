package main

import (
	"fmt"
	"net/url"
	"os"
	"github.com/stormasm/census-go/goreq"
)

func main() {

	key := os.Getenv("CENSUS_API_KEY")

	item3 := url.Values{}
	item3.Set("key", key)
	item3.Add("get", "P0010001,NAME")
	item3.Add("for", "state:02")

	item3.Encode()

	res3, err3 := goreq.Request{
		Method:      "GET",
		ContentType: "application/json",
		Uri:         "http://api.census.gov/data/2010/sf1",
		QueryString: item3,
	}.Do()

	str, _ := res3.Body.ToString()

	fmt.Println(err3)
	fmt.Println(str)
}
