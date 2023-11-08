package Hang

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"	
)


// Fonction pour choisir un mot au hasard depuis un fichier texte
func ChooseRandomWord(filename string) (string, error) {
	test := filename + ".txt"
	file, err := os.Open(test)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Lire chaque ligne du fichier
	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	// Générer un nombre aléatoire en fonction du temps actuel
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(words))

	return words[randomIndex], nil
}

var hangman = []string{
	`
  
=========`,
	`
  +---+
      |
      |
      |
      |
      |
=========`,
	`
  +---+
  |   |
  O   |
  |   |
      |
      |
=========`,
	`
  +---+
  |   |
  O   |
 /|\  |
      |
      |
=========`,
	`
  +---+
  |   |
  O   |
 /|\  |
 /    |
      |
=========`,
	`
  +---+
  |   |
  O   |
 /|\  |
 / \  |
      |
=========`,
}

// Fonction pour afficher la mort du pendu en fonction du nombre d'erreurs
func displayHangman(errors int) {
	if errors >= 0 && errors < len(hangman) {
		fmt.Println(hangman[errors])
	} else {
		fmt.Println("Le pendu est complet!")
	}
}

// Fonction pour jouer une partie du pendu
func PlayHangman() {
	score := 0
	// Choisissez un mot au hasard depuis le fichier "pli07.txt"
	randomWord, err := ChooseRandomWord("pli07")
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	// Convertir le mot en minuscules (pour la comparaison)
	randomWord = strings.ToLower(randomWord)

	// Initialisation du jeu
	wordLength := len(randomWord)
	guesses := make([]string, 0)
	errors := 0
	maxErrors := len(hangman) - 1

	// Masquer initialement 2 lettres pour les mots longs, 1 lettre pour les mots courts
	initialVisibleLetters := 2
	if wordLength <= 5 {
		initialVisibleLetters = 1
	}

	// Sélectionnez les lettres initialement visibles
	initialVisibleIndexes := rand.Perm(wordLength)
	for i := 0; i < initialVisibleLetters; i++ {
		guesses = append(guesses, string(randomWord[initialVisibleIndexes[i]]))
	}

	// Boucle principale du jeu
	for {
		// Afficher les lettres déjà testées
		fmt.Printf("Lettres déjà testées : %s\n", strings.Join(guesses, ", "))
	
		// Afficher l'état actuel du mot avec des lettres masquées
		display := make([]string, wordLength)
		for i, letter := range randomWord {
			if i < initialVisibleLetters || contains(guesses, string(letter)) {
				display[i] = string(letter)
			} else {
				display[i] = "_"
			}
		}
		fmt.Println("Mot à deviner:", strings.Join(display, " "))
	
		// Afficher l'état actuel du pendu
		displayHangman(errors)

		// Vérifier si le joueur a gagné
		if strings.Join(display, "") == randomWord {
			fmt.Println("Félicitations, vous avez gagné ! Le mot était :", randomWord)
			return
		}

		// Demander au joueur de deviner une lettre ou un mot complet
		var guess string
		fmt.Print("Devinez une lettre ou proposez un mot complet : ")
		_, _ = fmt.Scan(&guess)
		guess = strings.ToLower(guess)

		if len(guess) == 1 { // Le joueur a deviné une lettre
			// Vérifier si la lettre devinée est correcte
			if !strings.Contains(randomWord, guess) {
				fmt.Println("Désolé, la lettre", guess, "n'est pas dans le mot.")
				errors++
			} else {
				fmt.Println("Bonne devinette ! La lettre", guess, "est dans le mot.")
				guesses = append(guesses, guess)
			}
		} else if guess == randomWord { // Le joueur a proposé un mot complet correct
			fmt.Println("Félicitations, vous avez gagné !")
			fmt.Println("Félicitations, vous avez gagné ! Le mot était :", randomWord)
			return
		} else {
			// Le joueur a proposé un mot complet incorrect
			fmt.Println("Désolé, le mot proposé est incorrect. Vous perdez 2 points.")
			score -= 2 // Pénalité de 2 points
		}

		// Vérifier si le joueur a perdu
		if errors >= maxErrors {
			fmt.Println("Désolé, vous avez perdu ! Le mot était :", randomWord)
			return
		}
	}
}

func Hasard() {
	score := 0 // Score initial

	// Boucle pour continuer à jouer
	for {
		fmt.Println("Menu principal:")
		fmt.Println("1. Jouer au pendu")
		fmt.Println("2. Quitter")

		var choice int
		fmt.Print("Choisissez une option : ")
		_, _ = fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Début d'une nouvelle partie.")
			PlayHangman()
			score++
			fmt.Printf("Votre score actuel : %d\n", score)
		case 2:
			fmt.Println("Merci d'avoir joué au pendu !")
			return
		default:
			fmt.Println("Option invalide. Veuillez choisir une option valide.")
		}
	}
}

func contains(slice []string, item string) bool {
	for _, element := range slice {
		if element == item {
			return true
		}
	}
	return false
}