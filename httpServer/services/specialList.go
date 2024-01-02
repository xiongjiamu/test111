package special_list

import (
	"dkdns/dkFramework/databases"
)

const specialListBucket = "special_list"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

// Read 从特殊列表中读取指定 key 对应的值
func (s *Service) Read(key string) (string, error) {
	value, err := databases.Boltdb.Read(specialListBucket, key)
	if err != nil {
		return "", err
	}
	return string(value), nil
}

// Write 向特殊列表中写入指定 key-value 对
func (s *Service) Write(key string, value string) error {
	err := databases.Boltdb.Write(specialListBucket, key, value)
	return err
}

// Delete 根据指定的 key 从特殊列表中删除数据
func (s *Service) Delete(key string) error {
	err := databases.Boltdb.Delete(specialListBucket, key)
	return err
}

// ReadAll 读取特殊列表中所有的数据
func (s *Service) ReadAll() (map[string]string, error) {
	return databases.Boltdb.ReadBucket(specialListBucket)
}
