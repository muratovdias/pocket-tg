package boltdb

import (
	"errors"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/muratovdias/pocket-tg/pkg/repository"
	"strconv"
)

type Repo struct {
	db *bolt.DB
}

func NewRepo(db *bolt.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) Save(chatID int64, token string, bucket repository.Bucket) error {
	// Start transaction
	err := r.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return fmt.Errorf("repo: boltdb: Save(): %w", err)
		}
		err = b.Put(intToBytes(chatID), []byte(token))
		if err != nil {
			return fmt.Errorf("repo: boltdb: Save(): %w", err)
		}
		// Commit
		return nil
	})
	if err != nil {
		return fmt.Errorf("repo: boltdb: Save(): %w", err)
	}
	return nil
}
func (r *Repo) Get(chatID int64, bucket repository.Bucket) (string, error) {
	var token string

	err := r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		data := b.Get([]byte(intToBytes(chatID)))
		token = string(data)

		return nil
	})

	if err != nil {
		return "", fmt.Errorf("repo: boltdb: Get(): %w", err)
	}
	if token == "" {
		return "", errors.New("token not found")
	}
	return token, nil
}
func intToBytes(n int64) []byte {
	return []byte(strconv.FormatInt(n, 10))
}
