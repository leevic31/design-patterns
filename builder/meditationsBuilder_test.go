package main

import "testing"

func TestSetAuthor(t *testing.T) {
	book := &MeditationsBuilder{}
	book.setAuthor()
	if book.author != "Marcus Aurelius" {
		t.Fatalf("Book author is incorrect")
	}
}

func TestSetGenre(t *testing.T) {
	book := &MeditationsBuilder{}
	book.setGenre()
	if book.genre != "Philosophy" {
		t.Fatalf("Book genre is incorrect")
	}
}

func TestSetPages(t *testing.T) {
	book := &MeditationsBuilder{}
	book.setPages()
	if book.pages != 256 {
		t.Fatalf("Book pages is incorrect")
	}
}

func TestSetTitle(t *testing.T) {
	book := &MeditationsBuilder{}
	book.setTitle()
	if book.title != "Meditations" {
		t.Fatalf("Book title is incorrect")
	}
}