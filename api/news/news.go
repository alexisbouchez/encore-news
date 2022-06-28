package news

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"net/url"

	"encore.dev/storage/sqldb"
)

type News struct {
	ID    string
	Title string
	URL   string
	Text  string
}

type SubmitNewsParams struct {
	Title string
	URL   string
	Text  string
}

type QueryResponse struct {
	News []*News
}

// Submit a news.
//encore:api public method=POST path=/news
func Submit(ctx context.Context, p *SubmitNewsParams) (*News, error) {
	if (p.Text != "" && p.URL != "") {
		return nil, errors.New("must provide either text or url")
	}

	if (p.URL != "") {
		_, err := url.ParseRequestURI(p.URL)
		if err != nil {
			return nil, errors.New("invalid url")
		}
	}

	id, err := generateID()
	if err != nil {
		return nil, err
	}

	if err := insert(ctx, id, p); err != nil {
		return nil, err
	}

	return &News{
		ID:    id,
		Title: p.Title,
		URL:   p.URL,
		Text:  p.Text,
	}, nil
}

// generateID generates a random short ID.
func generateID() (string, error) {
	var data [6]byte // 6 bytes of entropy
	if _, err := rand.Read(data[:]); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(data[:]), nil
}

// insert inserts a news into the database.
func insert(ctx context.Context, id string, p *SubmitNewsParams) error {
	_, err := sqldb.Exec(ctx, `
		INSERT INTO news (id, title, url, text)
		VALUES ($1, $2, $3, $4)
	`, id, p.Title, p.URL, p.Text)
	return err
}

// Get retrieves the news for a given id.
//encore:api public method=GET path=/news/:id
func Get(ctx context.Context, id string) (*News, error) {
	n := &News{ID: id}
	err := sqldb.QueryRow(ctx, `
		SELECT title, url, text FROM news
		WHERE id = $1
	`, id).Scan(&n.Title, &n.URL, &n.Text)
	return n, err
}

// Query retrieves the news for a given query.
//encore:api public method=GET path=/news
func Query(ctx context.Context) (*QueryResponse, error) {
	rows, err := sqldb.Query(ctx, `
		SELECT id, title, url, text FROM news
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to query news")
	}
	defer rows.Close()

	var news []*News
	for rows.Next() {
		n := &News{}
		if err := rows.Scan(&n.ID, &n.Title, &n.URL, &n.Text); err != nil {
			return nil, fmt.Errorf("failed to scan news")
		}
		news = append(news, n)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate over news")
	}

	return &QueryResponse{ 
		News: news, 
	}, nil
}
