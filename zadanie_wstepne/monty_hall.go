package main

import (
	"fmt"
	"math/rand"
)

func zagrajRunde(changeChoice bool) bool {

	prizeBox := rand.Intn(3) + 1
	
    playerChoice := rand.Intn(3) + 1

	var revealedBox int
	for {
		revealedBox := rand.Intn(3) + 1
		if revealedBox != playerChoice && revealedBox != prizeBox {
			break
		}
	}

	if changeChoice {
		var newChoice int
		for {
			newChoice = rand.Intn(3) + 1
			if newChoice != 0 && newChoice != playerChoice && newChoice != revealedBox {
				break
			}
		}
		playerChoice = newChoice
	}

	return playerChoice == prizeBox
}

func main() {
    var liczbaRund int
	var strategia string

    fmt.Println("Podaj liczbę rund:")
    fmt.Scanln(&liczbaRund)

    for {
		
        fmt.Println("Zdecyduj czy chcesz zmianiać swój wybór (tak/nie)?")
        fmt.Scanln(&strategia)

        if strategia == "tak" || strategia == "nie" {
            break
        } else {
            fmt.Println("Proszę podać 'tak' lub 'nie'.")
        }
    }

	var zmianaWyboru bool
	if strategia == "tak" {
		zmianaWyboru = true
	} else {
		zmianaWyboru = false
	}

	wygrane := 0
    przegrane := 0

	for i := 0; i < liczbaRund; i++ {
        wygrana := zagrajRunde(zmianaWyboru)

        if wygrana {
            wygrane++
        } else {
            przegrane++
        }
    }

	if zmianaWyboru {
		fmt.Printf("Rozegrano %d rund, za każdym razem zmieniając wybór pudełka.\n", liczbaRund)
	} else {
		fmt.Printf("Rozegrano %d rund, nie zmianiając wyboru pudełka.\n", liczbaRund)
	}
	fmt.Printf("Liczba wygranych: %d\n", wygrane)
    fmt.Printf("Liczba porażek: %d\n", przegrane)
}