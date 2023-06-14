package keeper_test

import (
	"github.com/bandprotocol/chain/v2/x/tss/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (s *KeeperTestSuite) TestGetSetDEQueue() {
	ctx, k := s.ctx, s.app.TSSKeeper
	address, _ := sdk.AccAddressFromBech32("band1m5lq9u533qaya4q3nfyl6ulzqkpkhge9q8tpzs")
	deQueue := types.DEQueue{
		Head: 1,
		Tail: 2,
	}

	// Set DEQueue
	k.SetDEQueue(ctx, address, deQueue)

	// Get DEQueue
	got := k.GetDEQueue(ctx, address)

	s.Require().Equal(deQueue, got)
}

func (s *KeeperTestSuite) TestGetSetDE() {
	ctx, k := s.ctx, s.app.TSSKeeper
	address, _ := sdk.AccAddressFromBech32("band1m5lq9u533qaya4q3nfyl6ulzqkpkhge9q8tpzs")
	index := uint64(1)
	de := types.DE{
		PubD: []byte("D"),
		PubE: []byte("E"),
	}

	// Set DE
	k.SetDE(ctx, address, index, de)

	// Get DE
	got, err := k.GetDE(ctx, address, index)

	s.Require().NoError(err)
	s.Require().Equal(de, got)
}

func (s *KeeperTestSuite) TestDeleteDE() {
	ctx, k := s.ctx, s.app.TSSKeeper
	address, _ := sdk.AccAddressFromBech32("band1m5lq9u533qaya4q3nfyl6ulzqkpkhge9q8tpzs")
	index := uint64(1)
	de := types.DE{
		PubD: []byte("D"),
		PubE: []byte("E"),
	}

	// Set DE
	k.SetDE(ctx, address, index, de)

	// Get DE
	k.DeleteDE(ctx, address, index)

	// Try to get the deleted DE
	got, err := k.GetDE(ctx, address, index)

	s.Require().ErrorIs(types.ErrDENotFound, err)
	s.Require().Equal(types.DE{}, got)
}

func (s *KeeperTestSuite) TestHandleSetDEs() {
	ctx, k := s.ctx, s.app.TSSKeeper
	address, _ := sdk.AccAddressFromBech32("band1m5lq9u533qaya4q3nfyl6ulzqkpkhge9q8tpzs")
	des := []types.DE{
		{
			PubD: []byte("D"),
			PubE: []byte("E"),
		},
	}

	// Handle setting DEs
	k.HandleSetDEs(ctx, address, des)

	// Get DEQueue
	deQueue := k.GetDEQueue(ctx, address)

	// Check that all DEs have been stored correctly
	s.Require().Equal(uint64(len(des)), deQueue.Tail)
	for i := uint64(0); i < deQueue.Tail; i++ {
		gotDE, err := k.GetDE(ctx, address, i)
		s.Require().NoError(err)
		s.Require().Equal(des[i], gotDE)
	}
}

func (s *KeeperTestSuite) TestPollDE() {
	ctx, k := s.ctx, s.app.TSSKeeper
	address, _ := sdk.AccAddressFromBech32("band1m5lq9u533qaya4q3nfyl6ulzqkpkhge9q8tpzs")
	des := []types.DE{
		{
			PubD: []byte("D"),
			PubE: []byte("E"),
		},
	}
	index := uint64(1)

	// Set DE and DEQueue
	k.HandleSetDEs(ctx, address, des)

	// Poll DE
	polledDE, err := k.PollDE(ctx, address)
	s.Require().NoError(err)

	// Ensure polled DE is equal to original DE
	s.Require().Equal(des[0], polledDE)

	// Attempt to get deleted DE
	deletedDE, err := k.GetDE(ctx, address, index)

	// Should return error
	s.Require().ErrorIs(types.ErrDENotFound, err)
	s.Require().Equal(types.DE{}, deletedDE)
}
