// internal/commands/board.go
package commands

import (
	"fmt"

	"github.com/abhishek71994/tab-manager-cli/internal/storage"
	"github.com/abhishek71994/tab-manager-cli/internal/utils"

	"os"
)

type BoardCommand struct {
	fileStore *storage.FileStore
}

func NewBoardCommand(fileStore *storage.FileStore) *BoardCommand {
	return &BoardCommand{fileStore: fileStore}
}

func (bc *BoardCommand) CreateBoard(boardName string) error {
	// Create the board directory
	boardDir, err := utils.GetBoardDir(boardName)
	if err != nil {
		return err
	}

	// Fetch Chrome tabs
	tabs, err := utils.FetchChromeTabs()
	if err != nil {
		return fmt.Errorf("failed to fetch Chrome tabs: %v\n\nTo fix this issue:\n1. Install chrome-cli using 'brew install chrome-cli'\n2. Ensure Google Chrome is running\n3. Grant chrome-cli accessibility permissions in System Preferences > Security & Privacy > Privacy > Accessibility", err)
	}

	// Save the tabs to tablist.json
	if err := bc.fileStore.SaveTabList(boardDir, tabs); err != nil {
		return err
	}

	fmt.Printf("Board created: %s\n", boardName)
	return nil
}

func (bc *BoardCommand) ListBoards() error {
	// Get the .tabman directory
	tabmanDir, err := utils.GetDataDir()
	if err != nil {
		return err
	}

	// List all directories inside .tabman
	entries, err := os.ReadDir(tabmanDir)
	if err != nil {
		return err
	}

	if len(entries) == 0 {
		fmt.Println("No boards found.")
		return nil
	}

	fmt.Println("Boards:")
	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Printf("- %s\n", entry.Name())
		}
	}

	return nil
}

func (bc *BoardCommand) DeleteBoard(boardName string) error {
	// Get the board directory
	boardDir, err := utils.GetBoardDir(boardName)
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

	fmt.Printf("Board deleted: %s\n", boardName)
	return nil
}
