package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/ZupIT/ritchie-cli/pkg/security"
)

type logoutCmd struct {
	security.LogoutManager
}

func NewLogoutCmd(lm security.LogoutManager) *cobra.Command {
	l := logoutCmd{lm}
	return &cobra.Command{
		Use:     "logout",
		Short:   "User logout",
		Long:    "Destroy the user session of the organization",
		Example: "rit logout",
		RunE:    l.RunFunc(),
	}
}

func (l logoutCmd) RunFunc() CommandRunnerFunc {
	return func(cmd *cobra.Command, args []string) error {
		if err := l.Logout(); err != nil {
			return err
		}

		fmt.Println("Logout successful!")
		return nil
	}
}
