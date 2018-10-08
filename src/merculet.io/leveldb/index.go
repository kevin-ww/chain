package leveldb

import "github.com/gpmgo/gopm/modules/log"

func instantiate(dbpath string){
	if dbpath ==""{
		log.Fatal("db path must not be empty")
	}
}


type BucketDB struct{

}

func NewBucketDB(bucket string){

	//var bucket = NewStorage()
}


func (b *BucketDB) insert(record interface{}){

}

func (b *BucketDB) update(record interface{}){

}

func (b *BucketDB) query(q interface{}){

}