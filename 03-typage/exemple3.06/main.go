package main

import "fmt"

func main() {
	commentaire1 := `Dans le système d'exploitation "Windows" le répertoire utilisateur est "C:\Users\"`
	commentaire2 := "Dans le système d'exploitation \"Windows\" le répertoire utilisateur est \"C:\\Users\\\""
	fmt.Println(commentaire1)
	fmt.Println(commentaire2)
}
