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
	item3.Add("for", "place:02000")
	item3.Add("in", "state:35")

	res3, err3 := goreq.Request{
		Method:      "GET",
		ContentType: "application/json",
		Uri:         "http://api.census.gov/data/2010/sf1",
		QueryString: item3,
	}.Do()

	//fmt.Println(err3, res3)

	str, _ := res3.Body.ToString()

	fmt.Println(err3)
	fmt.Println(str)

}
