package main

import (
	"context"
	"notification-service/internal/api"
)

func main() {
	api.Handle(context.TODO())
}
