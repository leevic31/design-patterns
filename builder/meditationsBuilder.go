package main

type MeditationsBuilder struct {
	author string
	pages int
	genre string
	title string
}

func newMeditationsBuilder() *MeditationsBuilder {
	return &MeditationsBuilder{}
}

func (b *MeditationsBuilder) setAuthor() {
	b.author = "Marcus Aurelius"
}

func (b *MeditationsBuilder) setGenre() {
	b.genre = "Philosophy"
}

func (b *MeditationsBuilder) setPages() {
	b.pages = 256
}

func (b *MeditationsBuilder) setTitle() {
	b.title = "Meditations"
}

func (b *MeditationsBuilder) getBook() Book {
	return Book{
		author: b.author,
		pages: b.pages,
		genre: b.genre,
		title: b.title,
	}
}