package stringx

import "strings"

// HidePhone masks phone number, keeps first 3 and last 4 digits
func HidePhone(phone string) string {
	return Hide(phone, 3, 4, '*')
}

// HideEmail masks email, keeps first 2 characters before @, leaves domain unchanged
func HideEmail(email string) string {
	at := strings.Index(email, "@")
	if at <= 2 {
		return email
	}
	// Use Hide function to mask characters before @, keep first 2 chars, suffix length 0
	return Hide(email[:at], 2, 0, '*') + email[at:]
}

// HideIDCard masks ID card number, keeps first 6 and last 4 digits
func HideIDCard(id string) string {
	return Hide(id, 6, 4, '*')
}

// HideBankCard masks bank card number, keeps first 4 and last 4 digits
func HideBankCard(card string) string {
	return Hide(card, 4, 4, '*')
}

func Hide(s string, prefix, suffix int, mask rune) string {
	runes := []rune(s)
	length := len(runes)

	// String too short, mask all characters
	if length <= prefix+suffix {
		return strings.Repeat(string(mask), length)
	}
	// Fixed 4 masks in the middle
	middle := strings.Repeat(string(mask), 4)
	return string(runes[:prefix]) + middle + string(runes[length-suffix:])
}
