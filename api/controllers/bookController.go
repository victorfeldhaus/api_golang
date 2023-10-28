package controller

import (
	entities "app/api/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type bookController struct {
	books []entities.Book
}

func NewBookController() *bookController {
	return &bookController{}
}

func (b *bookController) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, b.books)
}

func (b *bookController) Create(ctx *gin.Context) {
	book := entities.NewBook()

	if err := ctx.BindJSON(&book); err != nil {
		return
	}

	b.books = append(b.books, *book)

	ctx.JSON(http.StatusOK, b.books)
}

func (b *bookController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	for idx, book := range b.books {
		if book.ID == id {
			b.books = append(b.books[0:idx], b.books[idx+1:]...)
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{
		"error": "Book not found",
	})
}
