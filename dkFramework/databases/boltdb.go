package databases

import (
	"dkdns/dkFramework/logger"
	"fmt"
	"github.com/boltdb/bolt"
)

var (
	Boltdb *DB // 包级别的全局变量，用于存储数据库连接实例
)

// DB 是 BoltDB 的包装器，用于操作键值对存储
type DB struct {
	Path string   // 存储数据库文件的路径
	conn *bolt.DB // BoltDB 连接实例
}

func init() {
	Boltdb = NewBoltDB("./dkdns.db")
	err := Boltdb.Open()
	if err != nil {
		logger.Println("Error opening database:", err)
	}
}

// NewDB 创建一个新的 BoltDB 实例
func NewBoltDB(path string) *DB {
	return &DB{Path: path}
}

// Open 打开 BoltDB 数据库连接
func (db *DB) Open() error {
	conn, err := bolt.Open(db.Path, 0600, nil)
	if err != nil {
		return err
	}
	db.conn = conn
	return nil
}

// Close 关闭 BoltDB 数据库连接
func (db *DB) Close() error {
	if db.conn == nil {
		return nil
	}
	err := db.conn.Close()
	if err != nil {
		return err
	}
	db.conn = nil
	return nil
}

// Read 从指定的 bucket 中读取指定 key 对应的值
func (db *DB) Read(bucketName string, key string) ([]byte, error) {
	var value []byte
	err := db.conn.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return bolt.ErrBucketNotFound
		}
		value = b.Get([]byte(key))
		return nil
	})
	return value, err
}

// Write 将指定 key-value 对写入指定的 bucket 中
func (db *DB) Write(bucketName string, key string, value string) error {
	err := db.conn.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}
		err = b.Put([]byte(key), []byte(value))
		return err
	})
	return err
}

// Delete 根据指定的 key 删除指定 bucket 中的数据
func (db *DB) Delete(bucketName string, key string) error {
	err := db.conn.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return fmt.Errorf("Bucket %s not found", bucketName)
		}
		return b.Delete([]byte(key))
	})
	return err
}

func (db *DB) ReadBucket(bucketName string) (map[string]string, error) {
	data := make(map[string]string)

	err := db.conn.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return bolt.ErrBucketNotFound
		}

		return b.ForEach(func(k, v []byte) error {
			data[string(k)] = string(v)
			return nil
		})
	})

	if err != nil {
		return nil, err
	}

	return data, nil
}
