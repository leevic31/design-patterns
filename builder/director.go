package main

type Director struct {
	builder BookBuilder
}

func newDirector(b BookBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b BookBuilder) {
	d.builder = b
}

func (d *Director) buildBook() Book {
	d.builder.setAuthor()
	d.builder.setGenre()
	d.builder.setPages()
	d.builder.setTitle()
	return d.builder.getBook()
}