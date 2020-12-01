package service

import (
	"github.com/hpcng/warewulf/internal/app/wwctl/services/dhcp"
	"github.com/spf13/cobra"
)

var (
	baseCmd = &cobra.Command{
		Use:   "service",
		Short: "Initialize Warewulf services",
		Long:  "Warewulf Initialization",
	}
)

func init() {
	baseCmd.AddCommand(dhcp.GetCommand())
}

// GetRootCommand returns the root cobra.Command for the application.
func GetCommand() *cobra.Command {
	return baseCmd
}