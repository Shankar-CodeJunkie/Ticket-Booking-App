package helper

func IsValidUserName(firstName string, lastName string) bool {
	if len(firstName) >= 2 && len(lastName) >= 2 {
		return true
	} else {
		return false
	}
}
