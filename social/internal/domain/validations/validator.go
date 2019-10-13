package validations

import (
	"regexp"
	"strings"
)

func CheckSlug(str string) (matched bool, err error) {
	matched, err = regexp.MatchString(`^[a-z0-9-_]+$`, str)
	return
}

func CheckEmail(str string) (matched bool, err error) {
	matched, err = regexp.MatchString(`^([a-z0-9_\\.-]+)@([a-z0-9_\\.-]+)\.([a-z.]{2,7})$`, str)
	return
}

func CheckPassword(str string) (matched bool, err error) {
	matched, err = regexp.MatchString(`^[\wd_@-]{6,25}$`, str)
	return
}

func CheckUrl(str string) (matched bool, err error) {
	//matched, err = regexp.MatchString(`^http(s)?:\/\/[\w.-]+(?:\.[\w\.-]+)+[\w\-\._~:/?#[\]@!\$&'\(\)\*\+,;=.]+$`, str)
	matched, err = regexp.MatchString(`^http(s)?:\/\/`, str)
	return
}

// check phone number
func PhoneClean(phone string) string {
	re := regexp.MustCompile(`[\s()-]+`)
	phone = re.ReplaceAllString(phone, "")
	if !strings.HasPrefix(phone, "+") {
		phone = "+" + phone
	}
	return phone
}

// clean space
func SpaceClean(str string) string {
	re := regexp.MustCompile(`[\s]+`)
	str = re.ReplaceAllString(str, "")
	return str
}

// clean number
func CheckPhone(str string) (matched bool, err error) {
	matched, err = regexp.MatchString(`^\+[\d]+$`, str)
	return
}
