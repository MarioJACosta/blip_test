package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

type Repository struct {
	Name          string
	ActivityScore int
}

func readCSV(filename string) (map[string]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %s\n", err.Error())
	}

	defer file.Close()

	reader := csv.NewReader(file)

	_, err = reader.Read()
	if err != nil {
		return nil, fmt.Errorf("Error reading header: %s\n", err.Error())
	}

	repoScores := make(map[string]int)

	for {
		row, err := reader.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			fmt.Printf("Error reading row: %s\n", err.Error())
			continue
		}

		if len(row) < 6 {
			continue
		}

		repoName := row[2]
		filesChanged, err := strconv.Atoi(row[3])
		if err != nil {
			fmt.Printf("Error converting filesChanged: %s\n", err.Error())
			continue
		}

		additions, err := strconv.Atoi(row[4])
		if err != nil {
			fmt.Printf("Error converting additions: %s\n", err.Error())
			continue
		}

		deletions, err := strconv.Atoi(row[5])
		if err != nil {
			fmt.Printf("Error converting deletions: %s\n", err.Error())
			continue
		}

		repoScores[repoName] += filesChanged + additions + deletions
	}

	return repoScores, nil
}

func rankRepositories(repoScores map[string]int) []Repository {
	repositories := make([]Repository, 0, len(repoScores))

	for name, score := range repoScores {
		repositories = append(repositories, Repository{Name: name, ActivityScore: score})
	}

	sort.Slice(repositories, func(i, j int) bool {
		return repositories[i].ActivityScore > repositories[j].ActivityScore
	})

	return repositories
}

func printTopRepositories(repositories []Repository, top int) {
	fmt.Printf("Top %d most active repositories \n", top)

	for i := 0; i < top && i < len(repositories); i++ {
		fmt.Printf("%d - Repository: %s has a score of: %d \n", i+1, repositories[i].Name, repositories[i].ActivityScore)
	}
}

func main() {
	const filename = "commits.csv"
	const top = 10

	repoScores, err := readCSV(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	repositories := rankRepositories(repoScores)
	printTopRepositories(repositories, top)
}
