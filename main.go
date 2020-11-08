package main

import (
	"fmt"
)

var topPos []string
var middlePos []string
var bottomPos []string
var inputValue int64

func main() {


	var arrayOfInput []int64


	if inputValue != 0 {
		// Permettant de déconstruire inputValue en un tableau d'entier
		arrayOfInput = extractNumbers(inputValue)

		// Appel la méthode qui va utiliser les valeurs du tableau d'entier pour créer les formes "digital"
		for _, value := range arrayOfInput{
			numberToDigit(value)
		}
		//Execute display() dans l'ordre pour afficher le resultat de la suite d'entier
		executeDisplay(topPos, middlePos, bottomPos)
	}else{
		enterValidInput()
	}


}

func enterValidInput(){
	fmt.Print("Entrez vos numéros: ")
	removeAllArrayElement()

	_, err := fmt.Scan(&inputValue)
	if err != nil || inputValue == 0 {
		fmt.Println("Veuillez entrer une suite d'entier de 1 à 9 sans espace")
		fmt.Println("Exemple: '15689'")
	}
	main()

}

/**
*Récupère une suite de numéro et les split dans un tableau
*"123456" => [1, 2, 3, 4, 5, 6]int
 */
func extractNumbers(inputValue int64) []int64 {
	var arrayOfInput []int64
	for inputValue != 0 {
		remainModulo := inputValue % 10
		arrayOfInput = append([]int64{remainModulo}, arrayOfInput...)
		inputValue= inputValue/10
	}
	return arrayOfInput
}

/**
*Permet de créer les formes "digital" à partir de tableau d'entier, remplissant 3 tableau de string topPos[], middlePos[], bottom[]
*"2" => topPos[" _ "]
*	 => middlePos[" _|"]
*	 => bottomPos["|_ "]
*En remplissant de manière séparé les tableaux en fonction de la position des "digits" on peut alors afficher en ligne "digits" à la manière d'un écran LCD
* /!\ Important: Ne pas oublier les espaces lors de la déclaration de digits dans ces tableaux
 */
func numberToDigit (value int64) {
	switch value {
		case 1:
			topPos = append(topPos, "   ")
			middlePos = append(middlePos, "  |")
			bottomPos = append(bottomPos, "  |")
		case 2:
			topPos = append(topPos, " _ ")
			middlePos = append(middlePos, " _|")
			bottomPos = append(bottomPos, "|_ ")
		case 3:
			topPos = append(topPos, " _ ")
			middlePos = append(middlePos, " _|")
			bottomPos = append(bottomPos, " _|")
		case 4:
			topPos = append(topPos, "   ")
			middlePos = append(middlePos, "|_|")
			bottomPos = append(bottomPos, "  |")
		case 5:
			topPos = append(topPos, " _ ")
			middlePos = append(middlePos, "|_ ")
			bottomPos = append(bottomPos, " _|")
		case 6:
			topPos = append(topPos, " _ ")
			middlePos = append(middlePos, "|_ ")
			bottomPos = append(bottomPos, "|_|")
		case 7:
			topPos = append(topPos, " _ ")
			middlePos = append(middlePos, "  |")
			bottomPos = append(bottomPos, "  |")
		case 8:
			topPos = append(topPos, " _ ")
			middlePos = append(middlePos, "|_|")
			bottomPos = append(bottomPos, "|_|")
		case 9:
			topPos = append(topPos, " _ ")
			middlePos = append(middlePos, "|_|")
			bottomPos = append(bottomPos, " _|")
	}
}

/**
*Permet de recuperer un tableau selon le niveau des digits (top, middle ou bottom)
*top[]=>"123" =>    _   _
* Le numéro 1 n'a pas de top donc string vide
* Le numéro 2 et 3 ont un top "_"
* Apres la boucle on Println pour revenir à la ligne et ainsi entamer le prochain niveau
 */
func display(digitPosition[]string){
	for _, digit := range digitPosition{
		fmt.Print(digit)
	}
	fmt.Println()
}

/**
* Execute dans l'ordre la fonction display
* 13 => 	_
*		 |  _|
*		 |	_|
* On remet à nil les tableaux de niveau pour ensuite reCall la méthode main pour donner la possibilité à l'utilisateur d'entré d'autre numéros
*/
func executeDisplay(top []string, middle []string, bottom[]string){
	display(top)
	display(middle)
	display(bottom)
	removeAllArrayElement()
	main()
}

/**
*Permet de remove les élements des tableaux de niveau
 */
func removeAllArrayElement(){
	topPos = nil
	middlePos = nil
	bottomPos = nil
	inputValue = 0
}



