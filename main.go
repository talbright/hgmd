package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	heroku "github.com/talbright/heroku-go/v4"
)

var (
	username = flag.String("username", "", "api username")
	password = flag.String("password", "", "api password")
)

func main() {
	log.SetFlags(0)
	flag.Parse()

	heroku.DefaultTransport.Username = *username
	heroku.DefaultTransport.Password = *password

	h := heroku.NewService(heroku.DefaultClient)
	addons, err := h.AddOnList(context.TODO(), &heroku.ListRange{Field: "name"})
	if err != nil {
		log.Fatal(err)
	}
	for _, addon := range addons {
		fmt.Println(addon.Name)
	}
}
