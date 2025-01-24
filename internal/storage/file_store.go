// internal/storage/file_store.go
package storage

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/abhishek71994/tab-manager-cli/internal/models"
)

type FileStore struct{}

func NewFileStore() *FileStore {
	return &FileStore{}
}

// SaveTabList saves the given tabs to tablist.json in the specified board directory.
func (fs *FileStore) SaveTabList(boardDir string, tabs []models.Tab) error {
	// Define the tablist.json file path
	filePath := filepath.Join(boardDir, "tablist.json")
	log.Printf("SaveTabList: filePath = %s", filePath) // Debug log

	// // Log all tabs being saved
	// log.Println("Saving the following tabs:")
	// for _, tab := range tabs {
	// 	log.Printf("- TabID: %s, Title: %s, URL: %s", tab.TabID, tab.Title, tab.URL)
	// }

	// Create the board directory if it doesn't exist
	if err := os.MkdirAll(boardDir, 0755); err != nil {
		return fmt.Errorf("failed to create board directory: %v", err)
	}

	// Save the tabs to tablist.json
	file, err := json.MarshalIndent(tabs, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal tabs: %v", err)
	}

	log.Printf("SaveTabList: Writing to file: %s", filePath) // Debug log
	return os.WriteFile(filePath, file, 0644)
}
