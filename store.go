package main

import (
	"errors"
	"fmt"
)

type Contact struct {
	Name string
	Tel  string
}

type ContactStore struct {
	contacts map[string]Contact
}

func NewKVStore() *ContactStore {
	return &ContactStore{
		contacts: make(map[string]Contact),
	}
}

func (cs *ContactStore) Add(name, tel string) error {
	if _, exists := cs.contacts[name]; exists {
		return fmt.Errorf("contact with name '%s' already exists", name)
	}
	cs.contacts[name] = Contact{Name: name, Tel: tel}
	return nil
}

func (cs *ContactStore) Search(name string) (Contact, error) {
	contact, ok := cs.contacts[name]
	if !ok {
		return Contact{}, errors.New("contact not found")
	}
	return contact, nil
}

func (cs *ContactStore) Update(name, tel string) error {
	if _, ok := cs.contacts[name]; !ok {
		return errors.New("contact not found")
	}
	cs.contacts[name] = Contact{Name: name, Tel: tel}
	return nil
}

func (cs *ContactStore) Delete(name string) error {
	if _, ok := cs.contacts[name]; !ok {
		return errors.New("contact not found")
	}
	delete(cs.contacts, name)
	return nil
}

func (cs *ContactStore) List() []Contact {
	all := []Contact{}
	for _, c := range cs.contacts {
		all = append(all, c)
	}
	return all
}
