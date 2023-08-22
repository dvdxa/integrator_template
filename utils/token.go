package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"regexp"
)

func CheckToken(expectedToken, Token string) bool {

	var re = regexp.MustCompile(`(?m)^\s*(Bearer)\s+(` + regexp.QuoteMeta(Token) + `)\s*$`)

	return re.MatchString(expectedToken)
}

func CreateSign(body interface{}, secretKey string) (sign string, err error) {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("failed to marshal: %v", err)
	}
	hasher := hmac.New(sha1.New, []byte(secretKey))
	hasher.Write(bodyBytes)
	signature := hex.EncodeToString(hasher.Sum(nil))

	return signature, nil
}
