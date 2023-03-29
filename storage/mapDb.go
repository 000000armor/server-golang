package storage

type MapDB struct {
	data map[string]interface{}
}

func NewMapDb() *MapDB {
	return &MapDB{
		data: make(map[string]interface{}),
	}
}

func (db *MapDB) Set(key string, value interface{}) {
	db.data[key] = value
}

func (db *MapDB) Get(key string) (value interface{}) {
	value = db.data[key]
	return
}

func (db *MapDB) Delete(key string) {
	delete(db.data, key)
}
