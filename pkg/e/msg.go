package e

var MsgFlags = map[int]string{
	SUCCESS:       "ok",
	ERROR:         "fail",
	UNAUTHORIZED:  "unauthorized",
	NO_PERMISSION: "no permission",

	LOGIN_FAILED:         "Login failed",
	ACCOUNT_ERROR:        "Incorrect username or password",
	GENERATE_TOKEN_ERROR: "Failed to generate token",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
