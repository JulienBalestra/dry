package main

import (
	"context"
	"os"

	"github.com/JulienBalestra/dry/pkg/exit"
	"github.com/JulienBalestra/dry/tree/cmd"
)

func main() {
	c := cmd.Command(context.TODO())
	err := c.Execute()
	code := exit.Exit(err)
	os.Exit(code)
}
