package cmd

import (
	"context"
	"sync"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func Command(ctx context.Context) *cobra.Command {
	c := &cobra.Command{
		Short:   "command example",
		Long:    "command example",
		Use:     "sub-cmd",
		Aliases: []string{"sc"},
	}

	fs := &pflag.FlagSet{}
	c.Flags().AddFlagSet(fs)
	c.RunE = func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithCancel(ctx)
		waitGroup := sync.WaitGroup{}

		<-ctx.Done()
		cancel()
		waitGroup.Wait()
		return nil
	}
	return c
}
