package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/boltdb/bolt"
)

func main() {
	// simpleTest()
	batchTest()
}

func batchTest() {
	execBatchPut()
	execLoopPut()
}

func execBatchPut() {
	var bucket = []byte("ClumsyTenz")

	// db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	db, err := bolt.Open("my.db", 0644, nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	startTime := time.Now()
	count := 1000 // 开启testCount个写协程

	wg := &sync.WaitGroup{}
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(bucket []byte, i int, group *sync.WaitGroup) {
			defer group.Done()
			// 改成db.Batch用于测试Batch API
			db.Batch(func(tx *bolt.Tx) error {
				kv := strconv.Itoa(i)
				b, err := tx.CreateBucketIfNotExists(bucket)
				if err != nil {
					return err
				}
				// b := tx.Bucket(bucket)
				err = b.Put([]byte(kv), []byte(kv))
				return err
			})
		}(bucket, i, wg)
	}

	// 主协程阻塞等待写协程执行完成
	wg.Wait()

	fmt.Printf("batch time cost = %v\n", time.Since(startTime))
}

func execLoopPut() {
	var bucket = []byte("ClumsyTenz")

	// db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	db, err := bolt.Open("my.db", 0644, nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	startTime := time.Now()
	count := 1000 // 开启testCount个写协程

	wg := &sync.WaitGroup{}
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(bucket []byte, i int, group *sync.WaitGroup) {
			defer group.Done()
			// 改成db.Batch用于测试Batch API
			db.Update(func(tx *bolt.Tx) error {
				kv := strconv.Itoa(i)
				b, err := tx.CreateBucketIfNotExists(bucket)
				if err != nil {
					return err
				}
				// b := tx.Bucket(bucket)
				err = b.Put([]byte(kv), []byte(kv))
				return err
			})
		}(bucket, i, wg)
	}

	// 主协程阻塞等待写协程执行完成
	wg.Wait()

	fmt.Printf("loop time cost = %v\n", time.Since(startTime))
}

func simpleTest() {
	var world = []byte("greeting")

	// db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	db, err := bolt.Open("my.db", 0644, nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	key := []byte("hello")
	value := []byte("Hello World!")

	// store some data
	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(world)
		if err != nil {
			return err
		}

		err = bucket.Put(key, value)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	// retrieve the data
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(world)
		if bucket == nil {
			return fmt.Errorf("bucket %s not found", world)
		}

		val := bucket.Get(key)
		fmt.Println(string(val))

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
