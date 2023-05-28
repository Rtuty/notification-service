package storage

import (
	"fmt"
	"notification-service/internal/storage"
	"testing"
)

func Test_GetStorageConfig(t *testing.T) {
	res, err := storage.GetStorageConfig()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}
