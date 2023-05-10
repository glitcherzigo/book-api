package grifts

import (
	"context"

	"github.com/glitcherzigo/book-api/models"

	"github.com/edgedb/edgedb-go"
)

func InsertBook(book models.Book) error {
	client := EdgedbClient()
	defer client.Close()

	var inserted struct {
		Id edgedb.UUID `edgedb:"id" json:"Id"`
	}

	ctx := context.Background()

	query := `
		INSERT Book {
			title := <str>$title,
			author := <str>$author,
			genre := <str>$genre,
			release_date := <datetime>$release_date,
			quantity := <int16>$quantity,
			sales := <bigint>$sales
		}
	`

	args := map[string]interface{}{
		"title":        book.Title,
		"author":       book.Author,
		"genre":        book.Genre,
		"release_date": book.Release_date,
		"quantity":     book.Quantity,
		"sales":        book.Sales,
	}

	err := client.QuerySingle(ctx, query, &inserted, args)
	if err != nil {
		return err
	}

	return nil
}

func SelectBooks() ([]models.Book, error) {
	client := EdgedbClient()
	defer client.Close()

	var books []models.Book

	ctx := context.Background()

	query := `
		SELECT Book {
			id,
			title,
			author,
			genre,
			release_date,
			quantity,
			sales
		}
	`

	err := client.Query(ctx, query, &books)
	if err != nil {
		return nil, err
	}

	return books, err
}

func GroupByGenre() ([]models.Groups, error) {
	client := EdgedbClient()
	defer client.Close()

	var books []models.Groups

	ctx := context.Background()

	query := `
		GROUP Book {
			id,
			title,
			author,
			genre,
			release_date,
			quantity,
			sales
		} BY .genre
	`

	err := client.Query(ctx, query, &books)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func GroupByAuthor() ([]models.Groups, error) {
	client := EdgedbClient()
	defer client.Close()

	var books []models.Groups

	ctx := context.Background()

	query := `
		GROUP Book {
			id,
			title,
			author,
			genre,
			release_date,
			quantity,
			sales
		} BY .author
	`

	err := client.Query(ctx, query, &books)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func DeleteBook(idstr string) error {
	client := EdgedbClient()
	defer client.Close()
	id, err := edgedb.ParseUUID(idstr)
	if err != nil {
		return err
	}

	ctx := context.Background()

	query := `
		DELETE Book
		FILTER .id = <uuid>$id;
	`

	args := map[string]interface{}{
		"id": id,
	}

	err = client.Execute(ctx, query, args)
	if err != nil {
		return err
	}

	return nil
}

func GetBookbyId(idstr string) (*models.Book, error) {
	client := EdgedbClient()
	defer client.Close()

	var book = new(models.Book)

	ctx := context.Background()

	query := `
		SELECT Book {
			id,
			title,
			author,
			genre,
			release_date,
			quantity,
			sales,
		}
		FILTER .id = <uuid>$id;
	`

	id, err := edgedb.ParseUUID(idstr)
	if err != nil {
		return nil, err
	}

	args := map[string]interface{}{
		"id": id,
	}

	err = client.QuerySingle(ctx, query, book, args)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func UpdateQuality(idstr string, checkout int16) error {
	client := EdgedbClient()
	defer client.Close()
	id, err := edgedb.ParseUUID(idstr)
	if err != nil {
		return err
	}

	book, err := GetBookbyId(idstr)
	if err != nil {
		return err
	}

	q := book.Quantity - checkout
	var updated struct {
		Id edgedb.UUID `edgedb:"id" json:"Id"`
	}

	ctx := context.Background()

	query := `
		UPDATE Book 
		FILTER .id = <uuid>$id

		SET {
			quantity := <int16>$quantity,
		};
	`

	args := map[string]interface{}{
		"id":       id,
		"quantity": q,
	}

	err = client.Query(ctx, query, &updated, args)
	if err != nil {
		return err
	}

	return nil
}
