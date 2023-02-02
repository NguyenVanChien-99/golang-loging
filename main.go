package main

import (
	"context"
	"github.com/NguyenVanChien-99/golang-loging/log"
)

func main() {
	ctx := context.WithValue(context.Background(), log.MIDDLEWARE_REQUEST_UID, "123456789")
	log.Info(ctx, "hello1")
	log.Infof(ctx, "Hello %d", 2)

	log.Error(ctx, "Hello 3")
	log.Errorf(ctx, "Hello %d", 4)

	ctx2 := context.Background()
	log.Info(ctx2, "Hello 5")

}
