package main

import (
    "fmt"
	"strconv"
	"math"
	"math/big"
	"strings"
)

var iloscWywolan int

func fibonacci(n int, value int) int {
	if (n == value) {
		iloscWywolan++
	}
	if n <= 1 {
		return n
	}
	return fibonacci(n-1, value) + fibonacci(n-2, value)
}

func silnia(n int) *big.Int {
	if n < 0 {
		return big.NewInt(0)
	}
	if n == 0 {
		return big.NewInt(1)
	}
	result := big.NewInt(int64(n))
	for i := n - 1; i > 0; i-- {
		result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}

func normalizeString(s string) string {
	polishToEnglish := map[string]string{
		"ą": "a",
		"ć": "c",
		"ę": "e",
		"ł": "l",
		"ń": "n",
		"ó": "o",
		"ś": "s",
		"ź": "z",
		"ż": "z",
	}

	converted := ""
	for _, char := range s {
		if replacement, ok := polishToEnglish[strings.ToLower(string(char))]; ok {
			converted += replacement
		} else {
			converted += string(char)
		}
	}
	return converted
}

func zawieraWszystkieLiczby(ciąg string, tablica []string) bool {
	licznik := make(map[string]int)

	for _, podciąg := range tablica {
		licznik[podciąg]++
	}

	for _, podciąg := range tablica {
		if strings.Count(ciąg, podciąg) < licznik[podciąg] {
			return false
		}
	}

	return true

}

func main() {
	var imie string
	var nazwisko string

	fmt.Println("Podaj imię: ")
    fmt.Scanln(&imie)
	fmt.Println("Podaj nazwisko: ")
	fmt.Scanln(&nazwisko)

	imie = normalizeString(imie)
	nazwisko = normalizeString(nazwisko)

	imie3 := imie[:3]
	nazwisko3 := nazwisko[:3]
	nick:=strings.ToLower(imie3+nazwisko3)
	fmt.Println("Wygenerowano nick:", nick)
	var nickAscii string
	var tablicaWartosciAscii []string
	for _, znak := range nick {
		ascii:=int(znak)
		stringAscii := strconv.Itoa(ascii)
		tablicaWartosciAscii = append(tablicaWartosciAscii, stringAscii)
		nickAscii+=stringAscii
	} 
	
	var silnaLiczba int

	for i := 0; i <= 1000; i++ {
        wynik := silnia(i).String()
		if zawieraWszystkieLiczby(wynik, tablicaWartosciAscii) {
			fmt.Printf("Silna liczba wynosi: %d. Silnia tej liczby wynosi: %s\n", i, wynik)
			silnaLiczba = i
			break
		}
    }

	fmt.Println("Silna liczba:", silnaLiczba)

	var poprzedniaIloscWywolan int

	for i := 1; i < 30; i++ {
		iloscWywolan = 0
		fibonacci(30, i)
		if (poprzedniaIloscWywolan > silnaLiczba && iloscWywolan <= silnaLiczba) {
			roznicaPoprzednia := math.Abs(float64(silnaLiczba - poprzedniaIloscWywolan))
			roznicaObecna := math.Abs(float64(silnaLiczba - iloscWywolan))
			if roznicaObecna < roznicaPoprzednia {
				fmt.Println("Słaba liczba:", i)
			} else {
				fmt.Println("Słaba liczba:", i-1)
			}
			break
		}
		poprzedniaIloscWywolan = iloscWywolan
	}

}

