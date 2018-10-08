package leveldb

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
)



//func instantiate(dbpath string){
//	if dbpath ==""{
//		log.Fatal("db path must not be empty")
//	}
//}

type Bucket struct{

}


type Record struct{
	Id string

}

type BucketDB struct {
	bucketName string
	bucket *leveldb.DB
}

func GetBucketDB() (bucketDb *BucketDB) {

	if bucketDb == nil {
		INSTANCE := NewStorage("/tmp/db")
		bucketDb := NewBucketDB(INSTANCE, "level-a")
		fmt.Printf("%v", bucketDb)
	}

	return bucketDb
}

func NewBucketDB(storage *Storage, bucket string) (bucketDb *BucketDB) {

	return nil
}

func (b *Bucket) insert(record interface{}) {
	GetBucketDB().bucket.Put(nil,nil, nil)

}

func (b *Bucket) update(record interface{}) {

}

func (b *Bucket) query(q interface{}) {

}
