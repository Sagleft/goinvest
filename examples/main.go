package main

import (
	"goinvest"
	"log"
)

func main() {
	client := goinvest.NewClient()
	_, err := client.GetScreener(goinvest.ScreenerTask{})
	if err != nil {
		log.Fatalln(err)
	}
}
