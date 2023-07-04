package cli

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	"github.com/stretchr/testify/require"

	"github.com/bandprotocol/chain/v2/x/tss/types"
)

func TestCombineGrantMsgs(t *testing.T) {
	granter := sdk.AccAddress([]byte("granter"))
	grantee := sdk.AccAddress([]byte("grantee"))
	msgGrants := types.GetMsgGrants()
	expiration := time.Now()

	msgs, err := combineGrantMsgs(granter, grantee, msgGrants, &expiration)
	require.NoError(t, err)
	require.Len(t, msgs, len(msgGrants))

	for i, msg := range msgs {
		// cast type
		msgGrant, ok := msg.(*authz.MsgGrant)
		require.True(t, ok)

		authorization, err := msgGrant.GetAuthorization()
		require.NoError(t, err)

		require.Equal(t, msgGrants[i], authorization.MsgTypeURL())
		require.Equal(t, granter.String(), msgGrant.Granter)
		require.Equal(t, grantee.String(), msgGrant.Grantee)
		require.Equal(t, expiration, *msgGrant.Grant.Expiration)
	}
}

func TestCombineRevokeMsgs(t *testing.T) {
	granter := sdk.AccAddress([]byte("granter"))
	grantee := sdk.AccAddress([]byte("grantee"))
	msgRevokes := []string{"revoke1", "revoke2"}

	msgs, err := combineRevokeMsgs(granter, grantee, msgRevokes)
	require.NoError(t, err)
	require.Len(t, msgs, len(msgRevokes))

	for i, msg := range msgs {
		// cast type
		msgRevoke, ok := msg.(*authz.MsgRevoke)
		require.True(t, ok)
		require.Equal(t, granter.String(), msgRevoke.Granter)
		require.Equal(t, grantee.String(), msgRevoke.Grantee)
		require.Equal(t, msgRevokes[i], msgRevoke.MsgTypeUrl)
	}
}

func TestParseComplaints(t *testing.T) {
	// 1. Test with empty string
	complaints, err := parseComplaints("")
	require.NoError(t, err)
	require.Empty(t, complaints)

	// 2. Test with a valid file
	// Write a valid JSON to a temp file
	tempFile, err := ioutil.TempFile("", "complaints")

	require.NoError(t, err)
	defer os.Remove(tempFile.Name())

	validJSON := `{
		"complaints": [
			{
				"i": 1,
				"j": 2,
				"key_sym": "A12yoSWiMwC+8k5XiD9UdQOrJZipntB9ZdSCtOof+O0m",
				"sig": "Aj1c3dvb5QNZAjHpqAljSM8n2TcUAh/q75GzwJVTcjujxdE324C0ZCgl5IxCVFDxRzHnzTwjl6u0sscOZacLBi4="
			}
		]
	}`
	_, err = tempFile.WriteString(validJSON)
	require.NoError(t, err)

	complaints, err = parseComplaints(tempFile.Name())
	require.NoError(t, err)
	require.Equal(t, 1, len(complaints))

	// 3. Test with a non-existent file
	complaints, err = parseComplaints("non-existent-file")
	require.Error(t, err)
	require.Nil(t, complaints)

	// 4. Test with an invalid JSON file
	invalidFile, err := ioutil.TempFile("", "invalid-complaints")
	require.NoError(t, err)
	defer os.Remove(invalidFile.Name())

	invalidJSON := `[{invalidJSON}]`
	_, err = invalidFile.WriteString(invalidJSON)
	require.NoError(t, err)

	complaints, err = parseComplaints(invalidFile.Name())
	require.Error(t, err)
	require.Nil(t, complaints)
}
