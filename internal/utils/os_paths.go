// internal/utils/os_paths.go
package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// GetDataDir returns the path to the .tabman directory.
func GetDataDir() (string, error) {
	// Get the user's home directory
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	// Define the .tabman directory
	tabmanDir := filepath.Join(home, ".tabman")

	// Create the directory if it doesn't exist
	if err := os.MkdirAll(tabmanDir, 0755); err != nil {
		return "", err
	}

	return tabmanDir, nil
}

// GetBoardDir returns the full path to a board's directory.
func GetBoardDir(boardName string) (string, error) {
	tabmanDir, err := GetDataDir()
	if err != nil {
		return "", err
	}

	// Define the board directory
	boardDir := filepath.Join(tabmanDir, boardName)

	log.Printf("GetBoardDir: %s", boardDir) // Debug log
	return boardDir, nil
}

// CreateBoardDir creates a new board directory inside .tabman.
func CreateBoardDir(boardName string) (string, error) {
	boardDir, err := GetBoardDir(boardName)
	if err != nil {
		return "", err
	}

	// Check if the board already exists
	if _, err := os.Stat(boardDir); !os.IsNotExist(err) {
		return "", fmt.Errorf("board '%s' already exists", boardName)
	}

	// Create the directory if it doesn't exist
	if err := os.MkdirAll(boardDir, 0755); err != nil {
		return "", err
	}

	return boardDir, nil
}

// ListBoardDirs returns a list of all board directories inside .tabman.
func ListBoardDirs() ([]string, error) {
	tabmanDir, err := GetDataDir()
	if err != nil {
		return nil, err
	}

	// List all directories inside .tabman
	entries, err := os.ReadDir(tabmanDir)
	if err != nil {
		return nil, err
	}

	var boards []string
	for _, entry := range entries {
		if entry.IsDir() {
			boards = append(boards, entry.Name())
		}
	}

	return boards, nil
}

// DeleteBoardDir deletes a board directory from .tabman.
func DeleteBoardDir(boardName string) error {
	boardDir, err := GetBoardDir(boardName)
	if err != nil {
		return err
	}

	// Check if the board exists
	if _, err := os.Stat(boardDir); os.IsNotExist(err) {
		return fmt.Errorf("board '%s' does not exist", boardName)
	}

	// Delete the board directory
	if err := os.RemoveAll(boardDir); err != nil {
		return err
	}

	return nil
}
