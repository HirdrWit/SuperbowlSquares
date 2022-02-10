package main

import (
	"encoding/csv"
	"log"
	"math/rand"
	"os"
	"strconv"
)

type Players struct {
	Name    string
	Squares int
}

func createPlayersList() []string {
	// open file
	f, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var players []string
	for i, line := range data {
		if i > 0 { // omit header line
			var rec Players
			for j, field := range line {
				if j == 0 {
					rec.Name = field
				} else if j == 2 {
					rec.Squares, _ = strconv.Atoi(field)
				}
			}
			for sqCount := 0; sqCount < rec.Squares; sqCount++ {
				players = append(players, rec.Name)
			}
		}
	}
	return players
}

func getRandomNumber(max int) int {
	return rand.Intn(max)
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func createGrid(playersList []string) [][]string {
	grid := make([][]string, 10)
	for i := 0; i < 10; i++ {
		grid[i] = make([]string, 10)
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			randomNumber := getRandomNumber(len(playersList))
			player := playersList[randomNumber]
			playersList = remove(playersList, randomNumber)
			grid[i][j] = player
		}

	}
	return grid
}

func writeToCSV(grid [][]string) error {
	file, err := os.Create("result.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range grid {
		if err := writer.Write(value); err != nil {
			return err
		}
	}
	return nil
}

func main() {

	// convert records to array of structs
	playersList := createPlayersList()
	grid := createGrid(playersList)
	err := writeToCSV(grid)
	if err != nil {
		log.Fatalln("error writing csv:", err)
	}

	// DEBUGGING CODE
	// // print the players list (debugging)
	// fmt.Printf("%+v\n", playersList)
	// // print the grid (debugging)
	// for i := 0; i < 10; i++ {
	// 	for j := 0; j < 10; j++ {
	// 		fmt.Printf("%+v", grid[i][j])
	// 	}
	// 	fmt.Printf("\n")
	// }
}
