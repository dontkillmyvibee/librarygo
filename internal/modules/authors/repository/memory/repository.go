package authorsmemory

import (
	"context"
	"fmt"
	authorsservice "library/internal/modules/authors/service"
	"sync"
	"time"
	"uuid"

	coreerrors "library/internal/core/errors"
	authorsdomain "library/internal/modules/authors/domain"
)

type Repository struct {
	mu      sync.Mutex
	authors map[uuid.UUID]authorsdomain.Author
}

func New() *Repository {
	return &Repository{
		authors: make(map[uuid.UUID]authorsdomain.Author),
	}
}

func (r *Repository) Create(ctx context.Context, author authorsdomain.Author) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if err := ctx.Err(); err != nil {
		return err
	}

	if _, exists := r.authors[author.ID]; exists {
		return fmt.Errorf("author ID already exists: %w", coreerrors.ErrConflict)
	}

	if author.MiddleName != nil {
		middleName := *author.MiddleName
		author.MiddleName = &middleName
	}

	r.authors[author.ID] = author
	return nil
}

func (r *Repository) GetAuthor(ctx context.Context, id uuid.UUID) (authorsdomain.Author, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if err := ctx.Err(); err != nil {
		return authorsdomain.Author{}, err
	}

	author, ok := r.authors[id]
	if !ok {
		return authorsdomain.Author{}, coreerrors.ErrNotFound
	}

	if author.MiddleName != nil {
		middleName := *author.MiddleName
		author.MiddleName = &middleName
	}

	return author, nil
}

func (r *Repository) GetAuthors(ctx context.Context) ([]authorsdomain.Author, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	authors := make([]authorsdomain.Author, 0, len(r.authors))

	for _, author := range r.authors {
		if author.MiddleName != nil {
			middleName := *author.MiddleName
			author.MiddleName = &middleName
		}

		authors = append(authors, author)
	}

	return authors, nil
}

func (r *Repository) Delete(ctx context.Context, id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if err := ctx.Err(); err != nil {
		return err
	}
	if _, ok := r.authors[id]; !ok {
		return coreerrors.ErrNotFound
	}

	delete(r.authors, id)

	return nil
}

func (r *Repository) Patch(
	ctx context.Context,
	id uuid.UUID,
	input authorsservice.PatchAuthorInput,
	updatedAt time.Time,
) (authorsdomain.Author, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if err := ctx.Err(); err != nil {
		return authorsdomain.Author{}, err
	}

	author, exists := r.authors[id]
	if !exists {
		return authorsdomain.Author{}, coreerrors.ErrNotFound
	}

	if input.FirstName != nil {
		author.FirstName = *input.FirstName
	}

	if input.LastName != nil {
		author.LastName = *input.LastName
	}

	if input.MiddleNameSet {
		author.MiddleName = input.MiddleName
	}

	author.UpdatedAt = updatedAt

	if author.MiddleName != nil {
		middleName := *author.MiddleName
		author.MiddleName = &middleName
	}

	r.authors[id] = author

	if author.MiddleName != nil {
		middleName := *author.MiddleName
		author.MiddleName = &middleName
	}

	return author, nil
}
