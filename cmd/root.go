package cmd

import (
	cobra "github.com/spf13/cobra"

	apply "github.com/hamza-007/ret3iac/cmd/apply"
)

func Execute() error {
	// Command
	root := &cobra.Command{
		Use:   "Ret3iac",
		Short: "Ret3iac CLI",
		Long:  "Ret3iac CLI",
	}

	root.AddCommand(apply.Command())

	return root.Execute()
}
