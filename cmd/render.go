package cmd

import (
	"github.com/farpointlabs/k8s_templater/internal/pkg/command"
	"github.com/spf13/cobra"
)

var readPath string

var renderCommand = &cobra.Command{
	Use:   "render [folder to process]",
	Short: "Processes template files",
	Long:  `bigger description`,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		c := &command.RenderHandler{}
		c.List(args[0])
		return nil
	},
}

func init() {
	rootCmd.AddCommand(renderCommand)
}
