package actions

import (
	"github.com/glitcherzigo/book-api/grifts"
	"github.com/savsgio/atreugo/v11"
)

func App() *atreugo.Atreugo {
	cfg := grifts.Settings()

	srv := atreugo.New(*cfg)

	srv.GET("/", GetBooks)
	srv.POST("/", PostBook)
	srv.GET("/genre", GroupGenre)
	srv.GET("/author", GroupAuthor)
	srv.GET("/{id}", GetBookById)
	srv.DELETE("/{id}", DeleteBook)

	return srv
}
