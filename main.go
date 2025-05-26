package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/go-tp1-annuary/store"
)

func main() {
	action := flag.String("action", "", "Action to perform: add, delete, update, search, list")
	name := flag.String("name", "", "Contact name")
	phone := flag.String("phone", "", "Phone number")
	flag.Parse()

	kv := store.NewKVStore()

	switch *action {
	case "add":
		if *name == "" || *phone == "" {
			log.Fatal("Missing --name or --phone")
		}
		err := kv.Add(*name, *phone)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Contact added.")
		}
	case "delete":
		if *name == "" {
			log.Fatal("Missing --name")
		}
		err := kv.Delete(*name)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Contact deleted.")
		}
	case "update":
		if *name == "" || *phone == "" {
			log.Fatal("Missing --name or --phone")
		}
		err := kv.Update(*name, *phone)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Contact updated.")
		}
	case "search":
		if *name == "" {
			log.Fatal("Missing --name")
		}
		phone, err := kv.Search(*name)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("%s -> %s\n", *name, phone)
		}
	case "list":
		for _, c := range kv.List() {
			fmt.Println(c)
		}
	case "reset":
		err := kv.Reset()
		if err != nil {
			fmt.Println("Error resetting contacts:", err)
			return
		}
		fmt.Println("Contacts reset successfully.")
	default:
		fmt.Println("Unsupported action. Use: add, delete, update, search, list")
		os.Exit(1)
	}
}
