package utils

import "regexp"

func IsValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func IsValidNIK(nik string) bool {
	re := regexp.MustCompile(`^\d{16}$`)
	return re.MatchString(nik)
}

func IsNumeric(phone string) bool {
	re := regexp.MustCompile(`^\d+$`)
	return re.MatchString(phone)
}
