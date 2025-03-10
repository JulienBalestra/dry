package main

import (
	"context"
	"os"
	"sync"

	"github.com/JulienBalestra/dry/pkg/exit"
	"github.com/JulienBalestra/dry/pkg/signals"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func main() {
	c := &cobra.Command{
		Short:   "command example",
		Long:    "command example",
		Use:     "sub-cmd",
		Aliases: []string{"sc"},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := context.WithCancel(context.Background())
			waitGroup := sync.WaitGroup{}
			waitGroup.Add(1)
			go func() {
				signals.NotifySignals(ctx, func() {})
				waitGroup.Done()
			}()
			waitGroup.Wait()
			cancel()
			<-ctx.Done()
			return nil
		},
	}
	fs := &pflag.FlagSet{}
	c.Flags().AddFlagSet(fs)
	err := c.Execute()
	code := exit.Exit(err)
	os.Exit(code)
}
