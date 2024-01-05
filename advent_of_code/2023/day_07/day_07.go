package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Hand struct {
	htype  uint8
	cards  [5]uint8
	bid    int
	cards2 string
}

func readHand(text string) Hand {

	values := strings.Fields(text)

	cardsMap := map[rune]uint8{
		'A': 12,
		'K': 11,
		'Q': 10,
		'J': 9,
		'T': 8,
		'9': 7,
		'8': 6,
		'7': 5,
		'6': 4,
		'5': 3,
		'4': 2,
		'3': 1,
		'2': 0,
	}

	cardsText := values[0]

	var cards [5]uint8
	for index, letter := range cardsText {
		cards[index] = cardsMap[letter]
	}

	bid, stringToIntConversionErr := strconv.Atoi(values[1])
	if stringToIntConversionErr != nil {
		log.Fatal(stringToIntConversionErr)
	}

	cumulativeCardsTypes := make([]uint8, 13)

	for _, card := range cards {
		cumulativeCardsTypes[card] = cumulativeCardsTypes[card] + 1
	}

	var higherCounter uint8 = 0
	var lowerCounter uint8 = 0

	for _, value := range cumulativeCardsTypes {
		if value >= higherCounter {
			lowerCounter = higherCounter
			higherCounter = value
		} else if value >= lowerCounter {
			lowerCounter = value
		}
	}

	typesMap := map[uint8]uint8{
		5: 6,
		4: 5,
		3: 3,
		2: 1,
		1: 0,
	}

	var htype uint8 = 0

	if higherCounter == 2 && lowerCounter == 2 {
		htype = 2
	} else if (higherCounter == 3 && lowerCounter == 2) || (higherCounter == 2 && lowerCounter == 3) {
		htype = 4
	} else {
		htype = typesMap[higherCounter]
	}

	hand := Hand{
		htype:  htype,
		cards:  cards,
		bid:    bid,
		cards2: cardsText,
	}

	return hand
}

func insertHand(newHand Hand, hands *[1000]Hand) {

AllLoops:
	for index, hand := range *hands {

		if hand.bid == 0 {
			hands[index] = newHand
			break
		}

		if hand.bid != 0 && newHand.htype < hand.htype {
			shiftCounter := 1
			for i := index + 1; i < 1000; i++ {
				if hands[i].bid != 0 {
					shiftCounter += 1
				} else {
					break
				}
			}
			for j := shiftCounter; j > 0; j-- {
				hands[index+j] = hands[index+j-1]
			}
			hands[index] = newHand
			break
		}

		if hand.bid != 0 && newHand.htype == hand.htype {

			for cardIndex, cardValue := range newHand.cards {
				tmp := hands[index].cards[cardIndex]
				if cardValue == tmp {
					continue
				}
				if cardValue < tmp {
					shiftCounter := 1
					for i := index + 1; i < 1000; i++ {
						if hands[i].bid != 0 {
							shiftCounter += 1
						} else {
							break
						}
					}
					for j := shiftCounter; j > 0; j-- {
						hands[index+j] = hands[index+j-1]
					}
					hands[index] = newHand
					break AllLoops
				}
				break
			}
		}
	}
}

func main() {

	start := time.Now()

	var hands [1000]Hand

	input, err := os.ReadFile("input/day_07")
	if err != nil {
		log.Fatal(err)
	}

	content := strings.Split(string(input), "\n")

	for _, line := range content {
		hand := readHand(line)
		insertHand(hand, &hands)
	}

	answer1 := 0

	for index, hand := range hands {
		answer1 += (index + 1) * hand.bid
	}

	elapsed := time.Since(start)
	fmt.Println(answer1)
	fmt.Println("time " + fmt.Sprint(elapsed))
}
