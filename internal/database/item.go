package database

import (
	"context"
	"time"
)

type CachedItem struct {
	ItemID    int       `json:"item_id"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetItem returns a cached item if it exists
func (db *Database) GetItem(ctx context.Context, itemID int) (*CachedItem, error) {
	row := db.Pool.QueryRow(ctx,
		`SELECT item_id, name, updated_at
         FROM item_metadata_cache
         WHERE item_id = $1`,
		itemID,
	)

	var item CachedItem
	err := row.Scan(&item.ItemID, &item.Name, &item.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

// InsertItem inserts a minimal cached item
func (db *Database) InsertItem(ctx context.Context, it CachedItem) error {
	_, err := db.Pool.Exec(ctx,
		`INSERT INTO item_metadata_cache (item_id, name, updated_at)
         VALUES ($1, $2, NOW())
         ON CONFLICT (item_id) DO NOTHING`,
		it.ItemID, it.Name,
	)
	return err
}
