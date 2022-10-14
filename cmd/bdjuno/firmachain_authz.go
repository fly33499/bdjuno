package main

import (
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authztype "github.com/cosmos/cosmos-sdk/x/authz"
)

func between(value string, a string, b string) string {
	// Get substring between two strings.
	posFirst := strings.Index(value, a)
	if posFirst == -1 {
		return ""
	}

	length := len(value)

	posLast := strings.Index(value[posFirst:length], b)
	if posLast == -1 {
		return ""
	}
	return (value[posFirst:length])[:posLast]
}

func FirmaChainAuthzMessagesParser(_ codec.Codec, cosmosMsg sdk.Msg) ([]string, error) {

	switch msg := cosmosMsg.(type) {

	case *authztype.MsgGrant:
		var stringArray = []string{msg.Grantee, msg.Granter}
		return stringArray, nil

	case *authztype.MsgRevoke:
		var stringArray = []string{msg.Grantee, msg.Granter}
		return stringArray, nil

	case *authztype.MsgExec:
		msgs, _ := msg.GetMessages()

		var stringArray = []string{}
		stringArray = append(stringArray, msg.Grantee)

		total := len(msgs)

		for i := 0; i < total; i++ {
			msgText := msgs[i].String()
			slice := strings.Split(msgText, " ")

			for _, str := range slice {

				userAddress := between(str, "firma1", "\"")

				if len(userAddress) > 0 {
					stringArray = append(stringArray, userAddress)
				}

				valoperAddress := between(str, "firmavaloper1", "\"")

				if len(valoperAddress) > 0 {
					stringArray = append(stringArray, valoperAddress)
				}
			}

		}

		return stringArray, nil
	}

	return nil, nil
}
