package store

import (
	"testing"
)

func setupTest(t *testing.T) *ContactStore {
	t.Helper()
	store := NewKVStore()
	if err := store.Reset(); err != nil {
		t.Fatalf("failed to reset store: %v", err)
	}
	return store
}

func TestAddContact(t *testing.T) {
	store := setupTest(t)
	err := store.Add("Alice", "0123456789")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	contact, err := store.Search("Alice")
	if err != nil {
		t.Errorf("expected contact Alice to exist")
	}
	if contact.Tel != "0123456789" {
		t.Errorf("expected tel 0123456789, got %s", contact.Tel)
	}
}

func TestSearchContact(t *testing.T) {
	store := setupTest(t)
	if err := store.Add("Bob", "0987654321"); err != nil {
		t.Fatalf("failed to add contact: %v", err)
	}

	contact, err := store.Search("Bob")
	if err != nil {
		t.Fatalf("expected to find Bob, got error: %v", err)
	}

	if contact.Tel != "0987654321" {
		t.Errorf("expected tel 0987654321, got %s", contact.Tel)
	}
}

func TestUpdateContact(t *testing.T) {
	store := setupTest(t)
	if err := store.Add("Charlie", "111"); err != nil {
		t.Fatalf("failed to add contact: %v", err)
	}
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
	store := setupTest(t)
	if err := store.Add("Diana", "333"); err != nil {
		t.Fatalf("failed to add contact: %v", err)
	}
	err := store.Delete("Diana")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_, err = store.Search("Diana")
	if err == nil {
		t.Errorf("expected error when searching deleted contact")
	}
}
