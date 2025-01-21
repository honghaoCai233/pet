package hashidv2

import (
	"testing"
)

func TestHashID_Encode(t *testing.T) {
	var (
		userHash = New(&Config{
			Prefix: "U0",
			Salt:   "user__",
			Type:   0,
		})
		shareHash = New(&Config{
			Type:   1,
			Prefix: "S0",
			Salt:   "shareGPT__",
		})
	)

	var (
		originUserID  = 3
		originShareID = 3
	)

	t.Logf("userID: %d encode: %s", originUserID, userHash.EncodeNotE(originUserID))
	t.Logf("shareGPTID: %d encode: %s", originShareID, shareHash.EncodeNotE(originShareID))
}
