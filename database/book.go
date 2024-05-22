package database

import (
	"github.com/go-sql-driver/mysql"
)

func GetAllBookList() ([]Book, error) {
	books := []Book{}

	rows, err := DBConn.Query(`SELECT id, name, author, pages, publication_date FROM books order by id`)
	if err != nil {
		return books, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var author string
		var pages int
		var publicationDate mysql.NullTime

		err = rows.Scan(&id, &name, &author, &pages, &publicationDate)
		if err != nil {
			return books, err
		}

		currentBook := Book{ID: id, Name: name, Author: author, Pages: pages}
		if publicationDate.Valid {
			currentBook.PublicationDate = publicationDate.Time
		}

		books = append(books, currentBook)
	}

	return books, err
}
