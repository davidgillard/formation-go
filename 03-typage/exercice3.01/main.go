package main

import (
	"fmt"
	"unicode"
)

func verifMotDePasse(pw string) bool {
	pwR := []rune(pw)
	if len(pwR) < 8 {
		return false
	}
	//definition des critéres acceptable
	contientMajuscule := false
	contientMiniscule := false
	contientNombre := false
	contientSymbole := false

	for _, v := range pwR {
		if unicode.IsUpper(v) {
			contientMajuscule = true
		}
		if unicode.IsLower(v) {
			contientMiniscule = true
		}

		if unicode.IsNumber(v) {
			contientNombre = true
		}

		if unicode.IsSymbol(v) || unicode.IsPunct(v) {
			contientSymbole = true
		}
	}
	return contientMajuscule && contientMiniscule && contientNombre && contientSymbole
}

func main() {
	// if verifMotDePasse("") {
	// 	fmt.Println("mot de passe respecte les critères")
	// } else {
	// 	fmt.Println("les critéres pour le mot de passe ne sont pas respecté")
	// }

	if verifMotDePasse("y_v0DJwng8:?1ocfL(m)%<wcL)dt0M") {
		fmt.Println("Le mot de passe respecte les critères")
	} else {
		fmt.Println("les critéres pour le mot de passe ne sont pas respecté")
	}
}
