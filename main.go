package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/go-tp1-annuary/store"
)

func main() {
	action := flag.String("action", "", "Action to perform: add, delete, update, search, list, reset")
	name := flag.String("name", "", "Contact name")
	phone := flag.String("phone", "", "Phone number")
	flag.Parse()

	kv := store.NewKVStore()

	switch *action {
	case "add":
		handleAdd(kv, *name, *phone)
	case "delete":
		handleDelete(kv, *name)
	case "update":
		handleUpdate(kv, *name, *phone)
	case "search":
		handleSearch(kv, *name)
	case "list":
		handleList(kv)
	case "reset":
		handleReset(kv)
	default:
		fmt.Println("Unsupported action. Use: add, delete, update, search, list, reset")
		os.Exit(1)
	}
}

func handleAdd(kv *store.ContactStore, name, phone string) {
	if name == "" || phone == "" {
		log.Fatal("Missing --name or --phone")
	}
	if err := kv.Add(name, phone); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Contact added.")
	}
}

func handleDelete(kv *store.ContactStore, name string) {
	if name == "" {
		log.Fatal("Missing --name")
	}
	if err := kv.Delete(name); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Contact deleted.")
	}
}

func handleUpdate(kv *store.ContactStore, name, phone string) {
	if name == "" || phone == "" {
		log.Fatal("Missing --name or --phone")
	}
	if err := kv.Update(name, phone); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Contact updated.")
	}
}

func handleSearch(kv *store.ContactStore, name string) {
	if name == "" {
		log.Fatal("Missing --name")
	}
	contact, err := kv.Search(name)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("%s -> %s\n", contact.Name, contact.Tel)
	}
}

func handleList(kv *store.ContactStore) {
	contacts := kv.List()
	if len(contacts) == 0 {
		fmt.Println("No contacts found.")
		return
	}
	for _, c := range contacts {
		fmt.Printf("Name: %s, Phone: %s\n", c.Name, c.Tel)
	}
}

func handleReset(kv *store.ContactStore) {
	if err := kv.Reset(); err != nil {
		fmt.Println("Error resetting contacts:", err)
	} else {
		fmt.Println("Contacts reset successfully.")
	}
}
