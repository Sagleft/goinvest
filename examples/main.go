package main

import (
	"encoding/json"
	"fmt"
	"goinvest"
	"log"
)

func main() {
	client := goinvest.NewClient()
	response, err := client.GetScreener(goinvest.ScreenerTask{
		Country:  "USA",
		Sector:   "all",
		Industry: "all",
	})
	if err != nil {
		log.Fatalln(err)
	}

	if len(response.Data) == 0 {
		log.Fatalln("no shares found")
	}

	fmt.Println(toJSON(response.Data[0]))
}

func toJSON(v interface{}) string {
	data, _ := json.MarshalIndent(v, "", "\t")
	return string(data)
}
