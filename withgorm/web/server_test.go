package web_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ar-sandbox3/level4/withgorm/db"
	"github.com/ar-sandbox3/level4/withgorm/db/models"
	"github.com/ar-sandbox3/level4/withgorm/db/seeder"
	"github.com/ar-sandbox3/level4/withgorm/web"
)

func TestHandler(t *testing.T) {
	conn, err := db.New(db.DefaultDSN)
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}

	if err := seeder.MigrateAndSeed(conn); err != nil {
		t.Fatalf("failed to migrate and seed: %v", err)
	}

	srv := httptest.NewServer(web.Handler(conn))
	resp, err := http.Get(srv.URL + "/departments")
	if err != nil {
		t.Fatalf("failed to get departments: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("failed to receive status OK, got: %v", resp.StatusCode)
	}

	var departments []models.Department
	if err := json.NewDecoder(resp.Body).Decode(&departments); err != nil {
		t.Fatalf("failed to decode departments: %v", err)
	}

	if departments[0].ID != 1 {
		t.Fatalf("got wrong data for first row: %d", departments[0].ID)
	}
}
