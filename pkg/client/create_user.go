package client

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-manage-service/management"
	"github.com/spf13/cobra"
)

const createUsersHelp = `Creats a user account

This command creates a user account on ther server.
`

func cmdCreateUser(cfg *config) *cobra.Command {
	var username string
	var password string
	var orgaLvl string

	cmd := &cobra.Command{
		Use:   "create_user",
		Short: "Creates a user account.",
		Long:  createUsersHelp,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := context.WithTimeout(context.Background(), cfg.timeout)
			defer cancel()

			service := connect(ctx, cfg.address)

			req := &management.CreateUserRequest{
				Username:                    username,
				Password:                    password,
				OrganisationManagementLevel: orgaLvl,
			}

			if _, err := service.CreateUser(ctx, req); err != nil {
				return fmt.Errorf("create user: %w", err)
			}
			return nil
		},
	}

	cmd.Flags().StringVarP(&username, "username", "u", "admin", "Name of the user account")
	cmd.Flags().StringVarP(&password, "password", "p", "admin", "Password for the user")
	cmd.Flags().StringVarP(&orgaLvl, "management_level", "m", "superadmin", "Set organisation management level")

	return cmd
}