package starportcmd

import (
	"github.com/spf13/cobra"
)

// NewNetworkProfile returns a new command to show the address profile info on Starport Network.
func NewNetworkProfile() *cobra.Command {
	c := &cobra.Command{
		Use:   "profile",
		Short: "Show the address profile info",
		Args:  cobra.NoArgs,
		RunE:  networkProfileHandler,
	}
	c.Flags().AddFlagSet(flagNetworkFrom())
	c.Flags().AddFlagSet(flagSetKeyringBackend())
	c.Flags().AddFlagSet(flagSetHome())
	return c
}

func networkProfileHandler(cmd *cobra.Command, args []string) error {
	nb, err := newNetworkBuilder(cmd)
	if err != nil {
		return err
	}

	n, err := nb.Network()
	if err != nil {
		return err
	}

	campaigns, err := n.Campaigns(cmd.Context())
	if err != nil {
		return err
	}

	nb.Cleanup()
	return nil
}
