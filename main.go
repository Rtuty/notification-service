package main

import (
	"context"
	"fmt"
	"notification-service/internal/storage"
)

func main() {
	sc, err := storage.GetStorageConfig()
	if err != nil {
		panic(err)
	}

	db, err := storage.NewClient(context.Background(), 5, sc)
	if err != nil {
		panic(err)
	}
	fmt.Println(db)
	// if err := db.AddClient(context.TODO(), entities.Client{}); err != nil {
	// 	panic(err)
	// }
}
