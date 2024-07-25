package cli

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/spf13/cobra"

	"github.com/bandprotocol/chain/v2/x/restake/types"
)

// GetTxCmd returns a root CLI command handler for all x/restake transaction commands.
func GetTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Restake transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		GetTxCmdClaimRewards(),
	)

	return txCmd
}

// GetTxCmdClaimRewards creates a CLI command for claming rewards
func GetTxCmdClaimRewards() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claim-rewards [key1] [key2] ...",
		Short: "Claim rewards",
		Args:  cobra.MinimumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			var msgs []sdk.Msg

			if len(args) == 0 {
				queryClient := types.NewQueryClient(clientCtx)
				var rewards []*types.Reward

				offset := uint64(0)
				pageSize := uint64(1000)
				for {
					respRewards, err := queryClient.Rewards(context.Background(), &types.QueryRewardsRequest{
						LockerAddress: clientCtx.GetFromAddress().String(),
						Pagination: &query.PageRequest{
							Offset: offset,
							Limit:  pageSize,
						},
					})
					if err != nil {
						return err
					}

					rewards = append(rewards, respRewards.Rewards...)
					offset += pageSize

					if respRewards.Pagination.NextKey == nil {
						break
					}
				}

				// claim all possible reward pools (>= 1 unit or key is deactivated)
				for _, reward := range rewards {
					respKey, err := queryClient.Key(context.Background(), &types.QueryKeyRequest{
						Key: reward.Key,
					})
					if err != nil {
						continue
					}

					finalReward, _ := reward.Rewards.TruncateDecimal()
					if !finalReward.IsZero() || !respKey.Key.IsActive {
						args = append(args, reward.Key)
					}
				}
			}

			for _, arg := range args {
				msg := types.NewMsgClaimRewards(
					clientCtx.GetFromAddress(),
					arg,
				)

				msgs = append(msgs, msg)
			}

			if len(msgs) == 0 {
				return fmt.Errorf("no rewards to claim")
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgs...)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
