package main

import (
	"testing"
)

func TestAddContact(t *testing.T) {
	store := NewKVStore()
	err := store.Add("Alice", "0123456789")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if _, exists := store.contacts["Alice"]; !exists {
		t.Errorf("expected contact Alice to exist")
	}
}

func TestSearchContact(t *testing.T) {
	store := NewKVStore()
	store.Add("Bob", "0987654321")

	contact, err := store.Search("Bob")
	if err != nil {
		t.Fatalf("expected to find Bob, got error: %v", err)
	}

	if contact.Tel != "0987654321" {
		t.Errorf("expected tel 0987654321, got %s", contact.Tel)
	}
}

func TestUpdateContact(t *testing.T) {
	store := NewKVStore()
	store.Add("Charlie", "111")
	err := store.Update("Charlie", "222")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	contact, _ := store.Search("Charlie")
	if contact.Tel != "222" {
		t.Errorf("expected updated tel 222, got %s", contact.Tel)
	}
}

func TestDeleteContact(t *testing.T) {
	store := NewKVStore()
	store.Add("Diana", "333")
	err := store.Delete("Diana")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_, err = store.Search("Diana")
	if err == nil {
		t.Errorf("expected error when searching deleted contact")
	}
}
