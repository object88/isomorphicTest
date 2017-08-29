package cmd

import (
	"github.com/object88/isomorphicTest/client"
	"github.com/spf13/cobra"
)

func createStopCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stop",
		Short: "Stop the server",
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := client.NewClient()
			if err != nil {
				return err
			}

			defer c.DestroyClient()

			return c.RequestShutdown()
		},
	}

	return cmd
}
