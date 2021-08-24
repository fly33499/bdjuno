package main

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/juno/modules/messages"
	evmtypes "github.com/tharsis/ethermint/x/evm/types"

	"github.com/ethereum/go-ethereum/crypto"
)

// EVMMessagesParser returns the list of all the accounts involved in the given
// message if it's related to the x/evm module. The account addresses are returned
// in both bech32 and hex formats.
func EVMMessagesParser(cdc codec.Codec, cosmosMsg sdk.Msg) ([]string, error) {
	// nolint: gocritic
	switch msg := cosmosMsg.(type) {

	case *evmtypes.MsgEthereumTx:
		var data evmtypes.TxData
		err := cdc.UnpackAny(msg.Data, &data)
		if err != nil {
			return nil, err
		}

		// Sender address is obtained from the signature and cached to msg.From
		sender, err := msg.GetSender(data.GetChainID())
		if err != nil {
			return nil, err
		}

		senderBech32 := sdk.AccAddress(sender.Bytes()).String()
		senderHex := sender.Hex()

		to := data.GetTo()

		if to != nil {
			toBech32 := sdk.AccAddress(to.Bytes()).String()
			toHex := to.Hex()
			return []string{senderBech32, senderHex, toBech32, toHex}, nil
		}

		// when the recipient address **IS** defined, we are performing a EVM Contract creation
		contract := crypto.CreateAddress(sender, data.GetNonce())
		contractBech32 := sdk.AccAddress(contract.Bytes()).String()
		contractHex := contract.Hex()

		return []string{senderBech32, senderHex, contractBech32, contractHex}, nil
	}

	return nil, messages.MessageNotSupported(cosmosMsg)
}
