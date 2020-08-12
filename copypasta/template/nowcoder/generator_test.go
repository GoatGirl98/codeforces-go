package nowcoder

import (
	"fmt"
	"os"
	"testing"
)

func TestGenNowCoderTemplates(t *testing.T) {
	emailOrPhone := os.Getenv("NOWCODER_USERNAME")
	cipherPwd := os.Getenv("NOWCODER_CIPHER_PWD")
	const contestID = 6778
	contestDir := fmt.Sprintf("../../../misc/nowcoder/%d/", contestID)
	if err := GenNowCoderTemplates(emailOrPhone, cipherPwd, contestDir, contestID, `// github.com/EndlessCheng/codeforces-go`); err != nil {
		t.Fatal(err)
	}
}