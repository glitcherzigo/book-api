package actions

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/glitcherzigo/book-api/grifts"
	"github.com/glitcherzigo/book-api/models"
	"github.com/savsgio/atreugo/v11"
)

func GetBooks(rc *atreugo.RequestCtx) error {
	books, err := grifts.SelectBooks()
	if err != nil {
		return rc.JSONResponse(atreugo.JSON{
			"error": err.Error(),
		}, http.StatusInternalServerError)
	}

	return rc.JSONResponse(atreugo.JSON{
		"books": books,
	})
}

func PostBook(rc *atreugo.RequestCtx) error {
	data := rc.PostBody()
	var book models.Book

	err := json.Unmarshal(data, &book)
	if err != nil {
		return rc.JSONResponse(atreugo.JSON{
			"error": err.Error(),
		}, http.StatusBadGateway)
	}

	err = grifts.InsertBook(book)
	if err != nil {
		return rc.JSONResponse(atreugo.JSON{
			"error": err.Error(),
		}, http.StatusInternalServerError)
	}

	return rc.JSONResponse(atreugo.JSON{
		"messgage": "Book Created successfully!",
	}, http.StatusCreated)
}

func GroupGenre(rc *atreugo.RequestCtx) error {
	games, err := grifts.GroupByGenre()
	if err != nil {
		return rc.JSONResponse(atreugo.JSON{
			"error": err.Error(),
		}, http.StatusInternalServerError)
	}

	return rc.JSONResponse(games, http.StatusOK)
}

func GetBookById(rc *atreugo.RequestCtx) error {
	idin := rc.UserValue("id")
	id := fmt.Sprintf("%v", idin)

	book, err := grifts.GetBookbyId(id)
	if err != nil {
		return rc.JSONResponse(atreugo.JSON{
			"error": err.Error(),
		})
	}

	return rc.JSONResponse(atreugo.JSON{
		"book": book,
	}, http.StatusGone)
}

func DeleteBook(rc *atreugo.RequestCtx) error {
	idin := rc.UserValue("id")
	id := fmt.Sprintf("%v", idin)

	err := grifts.DeleteBook(id)
	if err != nil {
		return rc.JSONResponse(atreugo.JSON{
			"error": err.Error(),
		})
	}

	return rc.JSONResponse(atreugo.JSON{
		"message": "Book deleted successfully!",
	}, http.StatusGone)
}

func GroupAuthor(rc *atreugo.RequestCtx) error {
	games, err := grifts.GroupByAuthor()
	if err != nil {
		return rc.JSONResponse(atreugo.JSON{
			"error": err.Error(),
		}, http.StatusInternalServerError)
	}

	return rc.JSONResponse(games, http.StatusOK)
}
