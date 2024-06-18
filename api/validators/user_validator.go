package validators

func ComparePasswords(password, confirmPassword string) bool {
	return password == confirmPassword
}
