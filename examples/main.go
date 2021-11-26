package main

import (
	"goinvest"
	"log"
)

func main() {
	client := goinvest.NewClient()
	_, err := client.GetScreener(goinvest.ScreenerTask{
		Country:  "USA",
		Sector:   "all",
		Industry: "all",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
