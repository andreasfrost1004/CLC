package database

import (
	"context"
	"time"
)

type CachedItem struct {
	ItemID    int
	Name      string
	Icon      string
	Level     int
	Quality   int
	Class     int
	Subclass  int
	Slot      int
	UpdatedAt time.Time
}

// for getting item data from cache
func (db *Database) GetItem(ctx context.Context, itemID int) (*CachedItem, error) {
	row := db.Pool.QueryRow(ctx,
		`SELECT item_id, name, icon, level, quality, class, subclass, slot, updated_at
         FROM item_metadata_cache
         WHERE item_id = $1`,
		itemID,
	)

	var item CachedItem
	err := row.Scan(
		&item.ItemID,
		&item.Name,
		&item.Icon,
		&item.Level,
		&item.Quality,
		&item.Class,
		&item.Subclass,
		&item.Slot,
		&item.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

// insert if not in cache
func (db *Database) InsertItem(ctx context.Context, it CachedItem) error {
	_, err := db.Pool.Exec(ctx,
		`INSERT INTO item_metadata_cache
         (item_id, name, icon, level, quality, class, subclass, slot, updated_at)
         VALUES ($1,$2,$3,$4,$5,$6,$7,$8,NOW())
         ON CONFLICT (item_id) DO NOTHING`,
		it.ItemID, it.Name, it.Icon, it.Level, it.Quality, it.Class, it.Subclass, it.Slot,
	)
	return err
}
