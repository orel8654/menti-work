package internal_middleware

import b64 "encoding/base64"

func PasswordEncoding(data string) string {
	return b64.StdEncoding.EncodeToString([]byte(data))
}
