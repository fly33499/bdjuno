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

	//fmt.Println("hello world")
	//fmt.Println("hello world")

	switch msg := cosmosMsg.(type) {
	case *authztype.MsgExec:

		//fmt.Println("msg.Grantee : " + msg.Grantee)
		//fmt.Println(msg.GetMessages())
		//fmt.Println("hello world")

		msgs, _ := msg.GetMessages()

		str1 := msgs[0].String()
		slice := strings.Split(str1, " ")

		var stringArray = []string{}

		stringArray = append(stringArray, msg.Grantee)

		for _, str := range slice {

			//fmt.Println(str)
			//fmt.Println("----------")

			result := between(str, "firma1", "\"")
			//fmt.Println(result)
			//fmt.Println(len(result))

			if len(result) > 0 {
				stringArray = append(stringArray, result)
			}

			result2 := between(str, "firmavaloper1", "\"")
			//fmt.Println(result2)
			//fmt.Println(len(result2))

			if len(result2) > 0 {
				stringArray = append(stringArray, result2)
			}

			//fmt.Println("----------")
		}

		return stringArray, nil
	}

	return nil, nil
}
