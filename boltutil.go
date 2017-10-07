package boltutil

import (
	"errors"

	"github.com/boltdb/bolt"
)

var (
	ErrNotFound = errors.New("Not Found")
)

// Get gets a value from a bucket. The nested bucket is supported
func Get(tx *bolt.Tx, bucketNames []string, key interface{}, to interface{}) error {
	bucket := Bucket(tx, bucketNames)
	if bucket == nil {
		return ErrNotFound
	}

	keyB, err := ToKeyBytes(key)
	if err != nil {
		return err
	}

	v := bucket.Get(keyB)
	if v == nil {
		return ErrNotFound
	}

	err = Deserialize(v, to)
	if err != nil {
		return err
	}

	return nil
}

func Set(tx *bolt.Tx, bucketNames []string, key interface{}, value interface{}) error {
	bucket, err := CreateBucketIfNotExists(tx, bucketNames)
	if err != nil {
		return err
	}

	keyB, err := ToKeyBytes(key)
	if err != nil {
		return err
	}

	valueB, err := Serialize(value)
	if err != nil {
		return err
	}

	err = bucket.Put(keyB, valueB)
	if err != nil {
		return err
	}

	return nil
}

func Delete(tx *bolt.Tx, bucketNames []string, key interface{}) error {
	bucket, err := CreateBucketIfNotExists(tx, bucketNames)
	if err != nil {
		return err
	}

	keyB, err := ToKeyBytes(key)
	if err != nil {
		return err
	}

	return bucket.Delete(keyB)
}

func DeleteBucket(tx *bolt.Tx, bucketNames []string, key interface{}) error {
	bucket, err := CreateBucketIfNotExists(tx, bucketNames)
	if err != nil {
		return err
	}

	keyB, err := ToKeyBytes(key)
	if err != nil {
		return err
	}

	return bucket.DeleteBucket(keyB)
}

func Bucket(tx *bolt.Tx, bucketNames []string) *bolt.Bucket {
	var bucket *bolt.Bucket
	for i, bucketName := range bucketNames {
		if i == 0 {
			bucket = tx.Bucket([]byte(bucketName))
			if bucket == nil {
				return nil
			}
		} else {
			bucket = bucket.Bucket([]byte(bucketName))
			if bucket == nil {
				return nil
			}
		}
	}

	return bucket
}

func CreateBucketIfNotExists(tx *bolt.Tx, bucketNames []string) (*bolt.Bucket, error) {
	var bucket *bolt.Bucket
	for i, bucketName := range bucketNames {
		if i == 0 {
			bc, err := tx.CreateBucketIfNotExists([]byte(bucketName))
			if err != nil {
				return nil, err
			}
			bucket = bc
		} else {
			bc, err := bucket.CreateBucketIfNotExists([]byte(bucketName))
			if err != nil {
				return nil, err
			}
			bucket = bc
		}
	}

	return bucket, nil
}

func Cursor(tx *bolt.Tx, bucketNames []string) (*bolt.Cursor, error) {
	bucket := Bucket(tx, bucketNames)
	if bucket == nil {
		return nil, ErrNotFound
	}

	return bucket.Cursor(), nil
}
