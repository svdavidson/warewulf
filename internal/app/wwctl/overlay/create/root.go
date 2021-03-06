package create

import (
	"github.com/spf13/cobra"
)

var (
	baseCmd = &cobra.Command{
		Use:                "create [overlay name]",
		Short:              "Initialize a new Overlay",
		Long:               "Create a new Warewulf provisioning overlay",
		RunE:				CobraRunE,
		Args: 				cobra.ExactArgs(1),

	}
	SystemOverlay bool
	NoOverlayUpdate bool
)

func init() {
	baseCmd.PersistentFlags().BoolVarP(&SystemOverlay, "system", "s", false, "Show System Overlays as well")
	baseCmd.PersistentFlags().BoolVarP(&NoOverlayUpdate, "noupdate", "n", false, "Don't update overlays")
}

// GetRootCommand returns the root cobra.Command for the application.
func GetCommand() *cobra.Command {
	return baseCmd
}
