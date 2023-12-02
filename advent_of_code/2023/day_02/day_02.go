package main

import (
	"fmt"
	"strconv"
	"strings"
)

type pixel struct {
	r, g, b int
}

type game struct {
	id     uint8
	pixels []pixel

	isPossible bool
}

func isValidGame(current_game game) bool {

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

func findMinimalPossibleSetsPower(batch game) int {

	var (
		r int
		g int
		b int
	)

	for _, pixels := range batch.pixels {
		if pixels.r > r {
			r = pixels.r
		}
		if pixels.g > g {
			g = pixels.g
		}
		if pixels.b > b {
			b = pixels.b
		}
	}

	return r * g * b
}

func main() {

	readFile, fileScanner := open_file("inputs/day_02_a.txt")
	var games []game

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
					current_pixel.r, _ = strconv.Atoi(string(color_value))
				case 'g':
					current_pixel.g, _ = strconv.Atoi(string(color_value))
				case 'b':
					current_pixel.b, _ = strconv.Atoi(string(color_value))
				}
			}

			current_game.pixels = append(current_game.pixels, current_pixel)
		}
		games = append(games, current_game)
	}

	var counter_solution_a int

	for _, game := range games {
		if !isValidGame(game) {
			game.isPossible = false
			continue
		} else {
			counter_solution_a += int(game.id)
		}
	}

	var counter_solution_b int

	for _, game := range games {
		counter_solution_b += findMinimalPossibleSetsPower(game)
	}

	fmt.Println(counter_solution_a)
	fmt.Println(counter_solution_b)

	readFile.Close()
}
