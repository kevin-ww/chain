package leveldb

import (
	"github.com/syndtr/goleveldb/leveldb"
)



type Storage struct{
	instance *leveldb.DB
}

func NewStorage(dbpath string) (INSTANCE *Storage){

	var storage = 	&Storage{
		instance:database("db"),
	}
	return storage
}


func  database(DB_PATH string) (INSTANCE *leveldb.DB) {

	debug("Create new LevelDB instance.")

	INSTANCE, err := leveldb.OpenFile(DB_PATH, nil)

	if err != nil {
		debug("Create new LevelDB instance has exception "+err.Error())
	}

	defer INSTANCE.Close()

	return INSTANCE
}


func (s *Storage) setDatabasePath(dbpath string){

}

func (s *Storage) destory(){
	//on close
	//on destory
}