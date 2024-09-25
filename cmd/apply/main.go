package apply

import (
	config "github.com/hamza-007/ret3iac/pkg/utils/config"

	resourceService "github.com/hamza-007/ret3iac/pkg/service/resource"

	cobra "github.com/spf13/cobra"
)

func Command() *cobra.Command {
	// Command
	cmd := &cobra.Command{
		Use:   "apply",
		Short: "Apply config file (.yaml/.yml)",
	}

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadConfig(args[0])
		if err != nil {
			return err
		}

		if err := resourceService.ResourceWriter().Interpret(cfg.Resources); err != nil {
			return err
		}

		return nil
	}
	return cmd
}
