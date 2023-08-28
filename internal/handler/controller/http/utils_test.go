package controllerhttp

import "testing"

func TestEmailValidation(t *testing.T) {
	emailList := []string{
		"111@gmail.com",
		"ab@yandex.ru",
		"armen@ro.ru",
	}

	for _, email := range emailList {
		err := validateEmail(email)
		if err != nil {
			t.Errorf("%s is valid email", email)
		}
	}
}

func TestEmailValidationWrong(t *testing.T) {
	emailList := []string{
		"@gmail.com",
		"abyandex.ru",
		"kiki@am.",
		"@ruchaam.ru",
		"limpopo@",
		"",
	}

	for _, email := range emailList {
		err := validateEmail(email)
		if err == nil {
			t.Errorf("%s is invalid email", email)
		}
	}
}
