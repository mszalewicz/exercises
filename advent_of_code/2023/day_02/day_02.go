package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type pixel struct {
	r, g, b int
}

type game struct {
	id     uint8
	pixels []pixel

	isPossible bool
}

func isValidGame(current_game *game) bool {

	max_r := 12
	max_g := 13
	max_b := 14

	for _, pixels := range current_game.pixels {
		if pixels.r > max_r || pixels.g > max_g || pixels.b > max_b {
			return false
		}
	}

	return true
}

func findMinimalPossibleSetsPower(batch *game) int {

	var r, g, b int

	for _, pixels := range batch.pixels {
		r = max(pixels.r)
		g = max(pixels.g)
		b = max(pixels.b)
	}

	return r * g * b
}

func main() {

	readFile, fileScanner := open_file("inputs/day_02.txt")
	var games []game

	start := time.Now()

	for fileScanner.Scan() {
		text := fileScanner.Text()

		game_info := strings.Split(text, ": ")[0]
		sets_info := strings.Split(text, ": ")[1]

		game_id, _ := strconv.Atoi(string(game_info[5:]))
		current_game := game{uint8(game_id), nil, true}

		for _, set := range strings.Split(sets_info, "; ") {

			current_pixel := pixel{0, 0, 0}

			for _, color := range strings.Split(set, ", ") {

				color_info := strings.Split(color, " ")
				color_value := color_info[0]
				color_type := color_info[1]

				switch color_type[0] {
				case 'r':
					current_pixel.r, _ = strconv.Atoi(color_value)
				case 'g':
					current_pixel.g, _ = strconv.Atoi(color_value)
				case 'b':
					current_pixel.b, _ = strconv.Atoi(color_value)
				}
			}

			current_game.pixels = append(current_game.pixels, current_pixel)
		}
		games = append(games, current_game)
	}

	var result1 int

	for _, game := range games {
		if !isValidGame(&game) {
			game.isPossible = false
			continue
		} else {
			result1 += int(game.id)
		}
	}

	var result2 int

	for _, game := range games {
		result2 += findMinimalPossibleSetsPower(&game)
	}

	elapsed := time.Since(start)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println("time " + fmt.Sprint(elapsed))

	readFile.Close()
}
