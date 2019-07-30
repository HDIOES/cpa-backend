package test

import (
	"net/http"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

func TestSearchAnimesSuccess(t *testing.T) {
	const externalGenreID string = "1"
	const externalStudioID string = "1"
	const externalAnimeID string = "1"
	//fill database
	insertGenreErr := insertGenreToDatabase(t, externalGenreID, "trashcore", "трешкор", "anime")
	if insertGenreErr != nil {
		t.Fatal(insertGenreErr)
	}
	insertStudioErr := insertStudioToDatabase(t, externalStudioID, "trash studio", "треш студия", true, "/url.jpg")
	if insertStudioErr != nil {
		t.Fatal(insertStudioErr)
	}
	insertAnimeErr := insertAnimeToDatabase(t, externalAnimeID, "One Punch Man", "Один Удар Человек", "/url.jpg", "tv", "ongoing", 10, 5,
		time.Date(2009, 11, 17, 20, 20, 20, 0, time.UTC),
		time.Date(2009, 11, 17, 20, 20, 20, 0, time.UTC), "/url.jpg", false, externalStudioID, externalGenreID)
	if insertAnimeErr != nil {
		t.Fatal(insertAnimeErr)
	}
	//create request
	request, _ := http.NewRequest("GET", "/api/animes/search", nil)
	executeRequest(request)
	//asserts

}
