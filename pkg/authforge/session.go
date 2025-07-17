package authforge

func CreateOrExpireToken(userId string) string {
	return *new(string)
}

func SessionValidation(sessionId string) bool { return true }

func RevokeSession() (bool, error) { return true, nil }
