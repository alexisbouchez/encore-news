package news

import (
	"context"
	"testing"
)

// TestSubmitAndRetrieve - test that the news is stored and retrieved from the database.
func TestSubmitAndRetrieve(t *testing.T) {
	// Test Submit()
	np := &SubmitNewsParams{
		Title: "Test Title",
		URL:   "https://example.com",
		Text:  "Test Text",
	}
	resp, err := Submit(context.Background(), np)
	if err != nil {
        t.Fatal(err)
	}

	if resp.Title != "Test Title" {
		t.Errorf("expected title to be 'Test Title', got '%s'", resp.Title)
	}
	if resp.URL != "https://example.com" {
		t.Errorf("expected url to be 'https://example.com', got '%s'", resp.URL)
	}
	if resp.Text != "Test Text" {
		t.Errorf("expected text to be 'Test Text', got '%s'", resp.Text)
	}

	// Test Get()
	news := resp

	resp, err = Get(context.Background(), news.ID)
	if err != nil {
		t.Fatal(err)
	}
	
	if resp.ID != news.ID {
		t.Errorf("expected id to be '%s', got '%s'", news.ID, resp.ID)
	}
	if resp.Title != news.Title {
		t.Errorf("expected title to be '%s', got '%s'", news.Title, resp.Title)
	}
	if resp.URL != news.URL {
		t.Errorf("expected url to be '%s', got '%s'", news.URL, resp.URL)
	}
	if resp.Text != news.Text {
		t.Errorf("expected text to be '%s', got '%s'", news.Text, resp.Text)
	}
}