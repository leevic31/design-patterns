package main

type BookBuilder interface {
	setAuthor() string
	setPages() int
	setGenre() string
	setTitle() string
	getBook() Book
}