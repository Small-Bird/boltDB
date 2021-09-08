package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

const dbFile = "info.db"
const bucketName = "person"

func CheckError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func write() {
	db, err := bolt.Open(dbFile, 0600, nil)
	CheckError(err)
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		CheckError(err)
		err = b.Put([]byte("001"), []byte("yangyang"))
		CheckError(err)
		err = b.Put([]byte("002"), []byte("yangyang1"))
		CheckError(err)
		err = b.Put([]byte("003"), []byte("yangyang2"))
		CheckError(err)
		return nil
	})
}
func read(key string) {
	db, err := bolt.Open(dbFile, 0600, nil)
	CheckError(err)
	defer db.Close()
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		name := bucket.Get([]byte(key))
		fmt.Println(string(name))
		return nil
	})
	CheckError(err)
}
func scan() {
	db, err := bolt.Open(dbFile, 0600, nil)
	CheckError(err)
	defer db.Close()
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		cursor := bucket.Cursor()
		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			fmt.Printf("%s,%s\n", k, v)
		}
		return nil
	})
	CheckError(err)
}
func main() {
	write()
	read("002")
	scan()
}
