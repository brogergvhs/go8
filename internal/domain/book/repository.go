package book

import (
	"context"

	"github.com/gmhafiz/go8/internal/model"
	"github.com/gmhafiz/go8/internal/resource"
)

//type Repository interface {
//	Create(ctx context.Context, book *model.Book) (*model.Book, error)
//	All(ctx context.Context) ([]sqlx.BookDB, error)
//}

type Repository interface {
	Create(ctx context.Context, book *model.Book) (*model.Book, error)
	All(ctx context.Context) ([]resource.BookDB, error)
	Close()
	Drop() error
	Up() error
}
