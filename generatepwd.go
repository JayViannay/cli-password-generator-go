package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
	"os"
	"github.com/fatih/color"
)

var green = color.New(color.FgGreen)
var yellow = color.New(color.FgYellow)
var red = color.New(color.FgRed)
var blue = color.New(color.FgBlue)

func generatePassword(length int, useSpecial bool) string {
    characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
    if useSpecial {
        characters += "!@#$%^&*()_+{}[]:;<>,.?/~"
    }

    rand.Seed(time.Now().UnixNano())

    password := make([]byte, length)
    for i := range password {
        password[i] = characters[rand.Intn(len(characters))]
    }

    return string(password)
}

func main() {
	var (
		length       int
		useSpecial   bool
		savePwd   bool
	)
    flag.IntVar(&length, "length", 12, "Longueur du mot de passe")
    flag.BoolVar(&useSpecial, "special", false, "Utiliser des caractères spéciaux")
	flag.BoolVar(&savePwd, "save", false, "Sauvegarder le mot de passe dans un fichier")
    flag.Parse()

    password := generatePassword(length, useSpecial)
	if length < 8 {
		red.Printf("Le mot de passe est trop court !\n")
	} else if length > 32 { 
		red.Printf("Le mot de passe est trop long !\n")
	} else {
		green.Printf("Mot de passe généré : \n")
		fmt.Printf(password)
	}

	fmt.Print("\nÊtes-vous satisfait.e du mot de passe généré ? (oui/non): ")
	var anotherPassword string
	fmt.Scan(&anotherPassword)

	if anotherPassword == "non" {
		main()
	} else {
		var saveToFile string
		if savePwd == false {
			fmt.Print("Voulez-vous sauvegarder le mot de passe dans un fichier ? (oui/non): ")
			fmt.Scan(&saveToFile)
		}

		if saveToFile == "oui" || savePwd == true {
			fmt.Print("Entrez le nom du site : ")
			var siteName string
			fmt.Scan(&siteName)

			fmt.Printf("Voulez-vous vraiment sauvegarder le mot de passe pour %s ? (oui/non): ", siteName)
			var confirm string
			fmt.Scan(&confirm)

			if confirm == "oui" {
				err := savePasswordToFile(siteName, password)
				if err != nil {
					fmt.Printf("Erreur lors de la sauvegarde du mot de passe : %v\n", err)
				} else {
					fmt.Printf("Mot de passe sauvegardé pour %s.\n", siteName)
				}
			}

			fmt.Print("\nVoulez-vous générer un autre mot de passe ? (oui/non): ")
			var anotherPassword string
			fmt.Scan(&anotherPassword)

			if anotherPassword == "oui" {
				main()
			} else {
				fmt.Println("Merci d'avoir utilisé le générateur de mots de passe. Au revoir !")
			}
		} else {
			fmt.Println("Merci d'avoir utilisé le générateur de mots de passe. Au revoir !")
		}
	}
}

func savePasswordToFile(siteName, password string) error {
	fileName := fmt.Sprintf("%s.txt", siteName)

	// fichier en écriture
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Ecrire le mot de passe dans le fichier
	_, err = file.WriteString(password)
	if err != nil {
		return err
	}

	return nil
}