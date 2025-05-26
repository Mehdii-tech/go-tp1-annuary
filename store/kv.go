package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"
)

type Contact struct {
	Name string `json:"name"`
	Tel  string `json:"tel"`
}

type ContactStore struct {
	mu       sync.RWMutex
	contacts map[string]Contact
}

const dataFile = "contacts.json"

func NewKVStore() *ContactStore {
	cs := &ContactStore{
		contacts: make(map[string]Contact),
	}
	cs.load()
	return cs
}

func (cs *ContactStore) load() {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	file, err := os.ReadFile(dataFile)
	if err != nil {
		return
	}

	var loaded []Contact
	if err := json.Unmarshal(file, &loaded); err != nil {
		fmt.Println("Error decoding data file:", err)
		return
	}

	for _, c := range loaded {
		cs.contacts[c.Name] = c
	}
}

func (cs *ContactStore) save(contacts []Contact) error {
	data, err := json.MarshalIndent(contacts, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling contacts:", err)
		return err
	}

	if err := os.WriteFile(dataFile, data, 0644); err != nil {
		fmt.Println("Error writing file:", err)
		return err
	}
	return nil
}

func (cs *ContactStore) Add(name, tel string) error {
	cs.mu.Lock()
	if _, exists := cs.contacts[name]; exists {
		cs.mu.Unlock()
		return fmt.Errorf("contact with name '%s' already exists", name)
	}
	cs.contacts[name] = Contact{Name: name, Tel: tel}

	all := make([]Contact, 0, len(cs.contacts))
	for _, c := range cs.contacts {
		all = append(all, c)
	}
	cs.mu.Unlock()

	return cs.save(all)
}

func (cs *ContactStore) Search(name string) (Contact, error) {
	cs.mu.RLock()
	defer cs.mu.RUnlock()

	contact, ok := cs.contacts[name]
	if !ok {
		return Contact{}, errors.New("contact not found")
	}
	return contact, nil
}

func (cs *ContactStore) Update(name, tel string) error {
	cs.mu.Lock()
	if _, ok := cs.contacts[name]; !ok {
		cs.mu.Unlock()
		return errors.New("contact not found")
	}
	cs.contacts[name] = Contact{Name: name, Tel: tel}

	all := make([]Contact, 0, len(cs.contacts))
	for _, c := range cs.contacts {
		all = append(all, c)
	}
	cs.mu.Unlock()

	return cs.save(all)
}

func (cs *ContactStore) Delete(name string) error {
	cs.mu.Lock()
	if _, ok := cs.contacts[name]; !ok {
		cs.mu.Unlock()
		return errors.New("contact not found")
	}
	delete(cs.contacts, name)

	all := make([]Contact, 0, len(cs.contacts))
	for _, c := range cs.contacts {
		all = append(all, c)
	}
	cs.mu.Unlock()

	return cs.save(all)
}

func (cs *ContactStore) List() []Contact {
	cs.mu.RLock()
	defer cs.mu.RUnlock()

	all := []Contact{}
	for _, c := range cs.contacts {
		all = append(all, c)
	}
	return all
}

func (cs *ContactStore) Reset() error {
	cs.contacts = make(map[string]Contact)
	err := os.Remove("contacts.json")
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}
