package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// define deck as a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"diamonds", "spades", "hearts", "clubs"}
	cardValues := []string{"one", "two", "three"}

	for _, suit := range cardSuits {
		for _, val := range cardValues {
			cards = append(cards, val+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, numCards int) (deck, deck) {
	return d[:numCards], d[numCards:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func loadFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	str := strings.Split(string(bs), ",")
	return deck(str)
}

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		i2 := r.Intn(len(d) - 1)
		d[i], d[i2] = d[i2], d[i]
	}
}
