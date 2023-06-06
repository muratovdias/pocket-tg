package boltdb

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/muratovdias/pocket-tg/pkg/repository"
)

func InitDB() (*bolt.DB, error) {
	db, err := bolt.Open("pocket.db", 0600, nil)

	if err != nil {
		return nil, fmt.Errorf("repo: bolt: InitDB(): %w", err)
	}
	if err := CreateBuckets(db); err != nil {
		return nil, err
	}
	return db, nil
}

func CreateBuckets(db *bolt.DB) error {
	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(repository.AccessToken))
		if err != nil {
			return fmt.Errorf("repo: boltdb: CreateBuckets(): %w", err)
		}
		_, err = tx.CreateBucketIfNotExists([]byte(repository.RequestToken))
		if err != nil {
			return fmt.Errorf("repo: boltdb: CreateBuckets(): %w", err)
		}
		return nil
	}); err != nil {
		return fmt.Errorf("repo: boltdb: CreateBuckets(): %w", err)
	}
	return nil
}
