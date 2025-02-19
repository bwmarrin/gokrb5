package kadmin

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/bwmarrin/gokrb5/iana/nametype"
	"github.com/bwmarrin/gokrb5/test/testdata"
	"github.com/bwmarrin/gokrb5/types"
)

func TestChangePasswdData_Marshal(t *testing.T) {
	t.Parallel()
	chgpasswd := ChangePasswdData{
		NewPasswd: []byte("newpassword"),
		TargName:  types.NewPrincipalName(nametype.KRB_NT_PRINCIPAL, "testuser1"),
		TargRealm: "TEST.GOKRB5",
	}
	chpwdb, err := chgpasswd.Marshal()
	if err != nil {
		t.Fatalf("error marshaling change passwd data: %v\n", err)
	}
	b, err := hex.DecodeString(testdata.MarshaledChangePasswdData)
	if err != nil {
		t.Fatalf("Test vector read error: %v", err)
	}
	assert.Equal(t, b, chpwdb, "marshaled bytes of change passwd data not as expected")
}
