package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Check if a directory path was provided as a command-line argument.
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go /path/to/directory")
		return
	}

	// Extract the directory path from the command-line argument.
	directory := os.Args[1]

	// Use the filepath package to clean and normalize the directory path.
	directory = filepath.Clean(directory)

	// List all files in the specified directory.
	files, err := os.ReadDir(directory)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	// Extract the series name from the parent directory of the episodes.
	seriesName := filepath.Base(filepath.Dir(directory))

	// Check if there's a "season" directory.
	season, err := extractSeasonNumber(directory)
	if err != nil {
		// If the "season" directory doesn't exist, assume it's a single-season series.
		fmt.Println("Assuming it's a single-season series.")
		season := "1" // Assume it's the first season
		seriesName := filepath.Base(directory)

		// Rename files in the directory.
		for i, file := range files {
			if shouldRename(file.Name()) {
				// Construct the new episode number with one leading "0" after "E".
				episode := fmt.Sprintf("E%02d", i+1)

				// Construct the new file name.
				newName := fmt.Sprintf("%s - S%s%s%s", seriesName, season, episode, filepath.Ext(file.Name()))

				oldPath := filepath.Join(directory, file.Name())
				newPath := filepath.Join(directory, newName)
				err := os.Rename(oldPath, newPath)
				if err != nil {
					fmt.Printf("Error renaming %s to %s: %v\n", file.Name(), newName, err)
				} else {
					fmt.Printf("Renamed %s to %s\n", file.Name(), newName)
				}
			}
		}
		return
	}

	// Rename files in the directory.
	for i, file := range files {
		if shouldRename(file.Name()) {
			// Construct the new episode number with one leading "0" after "E".
			episode := fmt.Sprintf("E%02d", i+1)

			// Construct the new file name.
			newName := fmt.Sprintf("%s - S%s%s%s", seriesName, season, episode, filepath.Ext(file.Name()))

			oldPath := filepath.Join(directory, file.Name())
			newPath := filepath.Join(directory, newName)
			err := os.Rename(oldPath, newPath)
			if err != nil {
				fmt.Printf("Error renaming %s to %s: %v\n", file.Name(), newName, err)
			} else {
				fmt.Printf("Renamed %s to %s\n", file.Name(), newName)
			}
		}
	}
}

func shouldRename(fileName string) bool {
	return strings.HasSuffix(fileName, ".mp4") || strings.HasSuffix(fileName, ".mkv")
}

func extractSeasonNumber(folderPath string) (string, error) {
	// Convert folderPath to lower case for case-insensitive matching
	lowerFolderPath := strings.ToLower(folderPath)

	// Check if folderPath contains "season" or "series"
	if !strings.Contains(lowerFolderPath, "season") && !strings.Contains(lowerFolderPath, "series") {
		return "", errors.New("folder path does not contain 'season' or 'series'")
	}

	parts := strings.Split(filepath.Base(folderPath), " ")
	season := parts[len(parts)-1]
	return season, nil
}
