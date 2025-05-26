package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"sync"
)

type Contact struct {
	FirstName string
	Phone     string
}

type ContactStore struct {
	mu      sync.RWMutex
	records map[string]Contact
}

func NewContactStore() *ContactStore {
	return &ContactStore{
		records: make(map[string]Contact),
	}
}

func (cs *ContactStore) Add(name, firstName, phone string) error {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	if _, exists := cs.records[name]; exists {
		return fmt.Errorf("contact with name %s already exists", name)
	}
	cs.records[name] = Contact{FirstName: firstName, Phone: phone}
	return nil
}

func (cs *ContactStore) Update(name, firstName, phone string) error {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	if _, exists := cs.records[name]; !exists {
		return fmt.Errorf("contact with name %s does not exist", name)
	}
	cs.records[name] = Contact{FirstName: firstName, Phone: phone}
	return nil
}

func (cs *ContactStore) Delete(name string) error {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	if _, exists := cs.records[name]; !exists {
		return fmt.Errorf("contact with name %s does not exist", name)
	}
	delete(cs.records, name)
	return nil
}

func (cs *ContactStore) Search(name string) (Contact, error) {
	cs.mu.RLock()
	defer cs.mu.RUnlock()

	contact, exists := cs.records[name]
	if !exists {
		return Contact{}, fmt.Errorf("contact with name %s not found", name)
	}
	return contact, nil
}

func (cs *ContactStore) List() []string {
	cs.mu.RLock()
	defer cs.mu.RUnlock()

	names := make([]string, 0, len(cs.records))
	for name := range cs.records {
		names = append(names, name)
	}
	return names
}

func main() {
	action := flag.String("action", "", "Action to perform: add, search, list, delete, update")
	name := flag.String("name", "", "Contact's name")
	firstName := flag.String("firstname", "", "Contact's first name")
	phone := flag.String("phone", "", "Contact's phone number")
	flag.Parse()

	store := NewContactStore()

	switch *action {
	case "add":
		err := store.Add(*name, *firstName, *phone)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Contact added.")

	case "update":
		err := store.Update(*name, *firstName, *phone)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Contact updated.")

	case "delete":
		err := store.Delete(*name)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Contact deleted.")

	case "search":
		contact, err := store.Search(*name)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		out, _ := json.MarshalIndent(contact, "", "  ")
		fmt.Println(string(out))

	case "list":
		names := store.List()
		fmt.Println("Contacts:")
		for _, name := range names {
			fmt.Println("- ", name)
		}

	default:
		fmt.Println("Unknown action. Use: add, search, list, delete, update")
	}
}
