package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	bugsnag "github.com/bugsnag/bugsnag-go"
)

func handler(ctx context.Context) error {
	bugsnag.Configure(bugsnag.Configuration{
		APIKey: "x",
		ProjectPackages: []string{
			"main",
		},
		Synchronous: true,
	})

	x := []string{"foo", "bar"}
	fmt.Println(x[2])

	return nil
}

func main() {
	lambda.Start(handler)
}
