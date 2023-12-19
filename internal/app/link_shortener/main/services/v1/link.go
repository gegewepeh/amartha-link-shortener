package v1

import (
	"context"
	"errors"
	"fmt"
	"link-shortener/internal/app/link_shortener/models"
	"link-shortener/internal/pkg/utils"
	"log"
	"os"
)

func (pool *Pool) GetLinks(ctx context.Context, id string, userId string) ([]models.Link, error) {
	var links = make([]models.Link, 0)
	log.Println("sini", userId)

	query := `SELECT * FROM links WHERE "userId" = $1`

	rows, err := pool.db.Query(ctx, query, userId)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

	for rows.Next() {
		values, err := rows.Values()
		log.Println("sini", values)
		if err != nil {
			log.Fatal("error while iterating dataset")
		}

		link := models.Link{
			FullLink: values[2].(string),
			Slug:     values[3].(string),
			Visit:    int(values[4].(int32)),
		}

		links = append(links, link)
	}

	return links, nil
}

func (pool *Pool) GetBySlugId(ctx context.Context, id string) (*models.Link, *utils.AppError) {
	var fullLink string
	var visit int

	queryUpdate := `UPDATE links
	SET visit = visit + 1
	WHERE slug = $1
	RETURNING "fullLink", "visit"`

	err := pool.db.QueryRow(ctx, queryUpdate, id).Scan(&fullLink, &visit)
	if err != nil {
		return nil, models.WrapError("Link", "NotFound", errors.New("Link Not Found"), nil)
	}

	link := models.Link{
		FullLink: fullLink,
		Visit:    visit,
	}

	return &link, nil
}

func (pool *Pool) CreateSlug(ctx context.Context, fullLink string, slug string, userId string) (*models.Link, *utils.AppError) {
	query := `INSERT INTO links("userId", "fullLink", slug, visit, "createdAt", "updatedAt")
	VALUES ($1, $2, $3, 0, NOW(), NOW())
	RETURNING slug`

	err := pool.db.QueryRow(ctx, query, userId, fullLink, slug).Scan(&slug)
	if err != nil {
		return nil, models.WrapError("Link", "InternalServerError", err, nil)
	}
	link := models.Link{
		Slug: "https://short.ly/" + slug,
	}

	return &link, nil
}

func (pool *Pool) UpdateSlugId(ctx context.Context, oldSlug string, newSlug string, userId string) (*models.Link, *utils.AppError) {
	var fullLink string
	var visit int

	queryUpdate := `UPDATE links
	SET slug = $1, visit = 0
	WHERE slug = $2 AND "userId" = $3
	RETURNING "fullLink", "visit", "slug"`

	err := pool.db.QueryRow(ctx, queryUpdate, newSlug, oldSlug, userId).Scan(&fullLink, &visit, &newSlug)
	if err != nil {
		return nil, models.WrapError("Link", "NotFound", errors.New("Link Not Found"), nil)
	}

	link := models.Link{
		FullLink: fullLink,
		Visit:    visit,
		Slug:     newSlug,
	}

	return &link, nil
}
