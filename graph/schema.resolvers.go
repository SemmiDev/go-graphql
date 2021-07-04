package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/SemmiDev/go-graphql/repository"
	"strconv"

	"github.com/SemmiDev/go-graphql/graph/generated"
	"github.com/SemmiDev/go-graphql/graph/model"
	"github.com/SemmiDev/go-graphql/internal/auth"
	"github.com/SemmiDev/go-graphql/internal/links"
	"github.com/SemmiDev/go-graphql/internal/pkg/jwt"
	"github.com/SemmiDev/go-graphql/internal/users"
)

func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return &model.Link{}, fmt.Errorf("access denied")
	}

	var link links.Link
	link.Title = input.Title
	link.Address = input.Address
	link.User = user
	linkId := link.Save()
	grahpqlUser := &model.User{
		ID:   user.ID,
		Name: user.Username,
	}

	return &model.Link{
		ID:      strconv.FormatInt(linkId, 10),
		Title:   link.Title,
		Address: link.Address,
		User:    grahpqlUser}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	var user users.User
	user.Username = input.Username
	user.Password = input.Password
	user.Create()
	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	var resultLinks []*model.Link
	var dbLinks []links.Link
	dbLinks = links.GetAll()
	for _, link := range dbLinks {
		grahpqlUser := &model.User{
			ID:   link.User.ID,
			Name: link.User.Username,
		}
		resultLinks = append(
			resultLinks,
			&model.Link{
				ID:      link.ID,
				Title:   link.Title,
				Address: link.Address,
				User:    grahpqlUser,
			},
		)
	}
	return resultLinks, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	var user users.User
	user.Username = input.Username
	user.Password = input.Password
	correct := user.Authenticate()
	if !correct {
		// 1
		return "", &users.WrongUsernameOrPasswordError{}
	}
	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	username, err := jwt.ParseToken(input.Token)
	if err != nil {
		return "", fmt.Errorf("access denied")
	}
	token, err := jwt.GenerateToken(username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) CreateAuthor(ctx context.Context, firstName string, lastName string) (*model.Author, error) {
	var author model.Author
	author.FirstName = firstName
	author.LastName = lastName
	id, err := repository.CreateAuthor(author)
	if err != nil {
		return nil, err
	} else {
		return &model.Author{ID: strconv.FormatInt(id, 10), FirstName: author.FirstName, LastName: author.LastName}, nil
	}
}

func (r *mutationResolver) CreateBook(ctx context.Context, title string, author string) (*model.Book, error) {
	var book model.Book
	book.Title = title
	book.Author = &model.Author{
		ID: author,
	}

	id, err := repository.CreateBook(book)
	if err != nil {
		return nil, err
	}
	idStr := strconv.Itoa(int(id))
	createdBook, _ := repository.GetBooksByID(&idStr)
	return createdBook, nil
}

func (r *queryResolver) BookByID(ctx context.Context, id *string) (*model.Book, error) {
	book, err := repository.GetBooksByID(id)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (r *queryResolver) AllBooks(ctx context.Context) ([]*model.Book, error) {
	books, err := repository.GetAllBooks()
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (r *queryResolver) AuthorByID(ctx context.Context, id *string) (*model.Author, error) {
	author, err := repository.GetAuthorByID(id)
	if err != nil {
		return nil, err
	} else {
		return author, nil
	}
}

func (r *queryResolver) AllAuthors(ctx context.Context) ([]*model.Author, error) {
	authors, err := repository.GetAllAuthors()
	if err != nil {
		return nil, err
	} else {
		return authors, err
	}
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
