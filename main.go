package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/ashkrai/orders-api/application"
)

func main() {
	app := application.New()

	ctx , cancel := signal.NotifyContext(context.Background(), os.Interrupt,os.Kill)
	defer cancel()
	err := app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app ", err)
	}
}
