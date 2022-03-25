package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/DanielMartin96/goddit/internal/post"
)

type PostRow struct {
	ID string
	Slug sql.NullString
	Author sql.NullString
	Body sql.NullString
}

func convertPostRowToPost(p PostRow) post.Post {
	return post.Post{
		ID: p.ID,
		Slug: p.Slug.String,
		Author: p.Author.String,
		Body: p.Body.String,
	}
}

func (d * Database) GetPost(ctx context.Context, uuid string) (post.Post, error) {
	var pstRow PostRow
	row := d.Client.QueryRowContext(
		ctx,
		`SELECT id, slug, author, body
		FROM posts
		WHERE id = $1`,
		uuid,
	)
	err := row.Scan(&pstRow.ID, &pstRow.Slug, &pstRow.Author, &pstRow.Body)
	if err != nil {
		return post.Post{}, fmt.Errorf("error fetching post by uuid: %w", err)
	}


	return convertPostRowToPost(pstRow), nil
}