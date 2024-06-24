package regex

import "regexp"

func ValidCep(cep string) bool {
	regex := regexp.MustCompile(`^[0-9]{8}$`)
	return regex.MatchString(cep)
}
