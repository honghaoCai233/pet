package xphone

import (
	"regexp"

	"github.com/nyaruka/phonenumbers"
)

var (
	cnPhoneNumberReg, _ = regexp.Compile(`^(\+86)(1)[0-9]{10}$`)
)

func HideSensitive(phone string) string {
	length := len(phone)
	if length < 4 {
		return phone
	}
	if length <= 8 {
		return "****" + phone[length-4:]
	}
	return phone[:length-8] + "****" + phone[length-4:]
}

func GetLastFour(phone string) string {
	length := len(phone)
	if length < 4 {
		return phone
	}
	return phone[length-4:]
}

func IsCNPhoneNumber(phone string) bool {
	parse, err := phonenumbers.Parse(phone, "CN")
	if err != nil {
		return false
	}
	if *parse.CountryCode != 86 {
		return false
	}
	return cnPhoneNumberReg.MatchString(phone)
}
