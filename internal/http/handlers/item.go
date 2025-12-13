package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"CLC/internal/database"
	"CLC/internal/wowhead"
)

func GetItem(db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			http.Error(w, "missing ?id=", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		// 1. Try DB first
		cached, err := db.GetItem(ctx, id)
		if err == nil { // found
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(cached)
			return
		}

		// 2. Fetch from Wowhead
		xmlItem, err := wowhead.FetchItemXML(id)
		if err != nil {
			http.Error(w, "wowhead fetch failed: "+err.Error(), 500)
			return
		}

		item := database.CachedItem{
			ItemID: xmlItem.Item.ID,
			Name:   xmlItem.Item.Name,
		}

		// 3. Insert minimal item data into DB
		if err := db.InsertItem(ctx, item); err != nil {
			http.Error(w, "db insert failed: "+err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(item)
	}
}
