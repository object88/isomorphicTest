package cmd

import (
	"github.com/object88/isomorphicTest/client"
	"github.com/spf13/cobra"
)

func createStartCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start the server",
		RunE: func(_ *cobra.Command, _ []string) error {
			c, err := client.NewClient()
			if err != nil {
				return err
			}

			err = c.RequestStartup()
			if err == nil {
				c.DestroyClient()
				return nil
			}

			c, err = c.RequestNewService()
			if err != nil {
				return err
			}

			c.DestroyClient()

			return nil
		},
	}

	return cmd
}
