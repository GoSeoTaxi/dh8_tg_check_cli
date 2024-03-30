// Binary gif-download implements example of gif backup.
package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/fatih/color"
	"github.com/go-faster/errors"
	"github.com/gotd/td/examples"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
)

const ver = "v0.0.1"

func fR() string {
	return strconv.Itoa(rand.Intn(10))
}

func run(ctx context.Context) error {

	ctx, cancel := context.WithTimeout(ctx, time.Duration(10*time.Second))
	defer cancel()

	var (
		appID   = int(gofakeit.Int8())
		appHash = gofakeit.ProductName()
		opt     telegram.Options
	)

	client := telegram.NewClient(appID, appHash, opt)
	flow := auth.NewFlow(examples.Terminal{PhoneNumber: "+70998" + fR() + fR() + fR() + fR() + fR() + fR()}, auth.SendCodeOptions{})
	return client.Run(ctx, func(ctx context.Context) error {
		if err := client.Auth().IfNecessary(ctx, flow); err != nil {
			return errors.Wrap(err, "auth")
		}

		return nil
	})
}

func main() {
	log.Printf(color.HiRedString("Starting CLI telegram checker to auth - version %v"), ver)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	err := run(ctx)
	if strings.Contains(err.Error(), "API_ID_INVALID") {
		os.Exit(0)
	}
	os.Exit(1)
}
