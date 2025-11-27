package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"CLC/internal/database"
)

type ItemMetadata struct {
	ItemID int     `json:"item_id"`
	Name   string  `json:"name"`
	Icon   string  `json:"icon"`
	Slot   string  `json:"slot"`
	Tier   string  `json:"tier"`
	Weight float64 `json:"weight"`
}

func GetItem(db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			http.Error(w, "missing ?id=", http.StatusBadRequest)
			return
		}

		itemID, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		var item ItemMetadata

		err = db.Pool.QueryRow(ctx,
			`SELECT item_id, name, icon, slot, tier, weight 
             FROM item_metadata_cache 
             WHERE item_id = $1`,
			itemID,
		).Scan(
			&item.ItemID,
			&item.Name,
			&item.Icon,
			&item.Slot,
			&item.Tier,
			&item.Weight,
		)

		if err != nil {
			http.Error(w, "item not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(item)
	}
}
