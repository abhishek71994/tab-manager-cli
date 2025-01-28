package commands

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/abhishek71994/tab-manager-cli/internal/models"
	"github.com/abhishek71994/tab-manager-cli/internal/storage"
	"github.com/abhishek71994/tab-manager-cli/internal/utils"
)

type TabCommand struct {
	fileStore *storage.FileStore
}

func NewTabCommand(fileStore *storage.FileStore) *TabCommand {
	return &TabCommand{fileStore: fileStore}
}

func (tc *TabCommand) ListTabs(boardName string) error {
	boardDir, err := utils.GetBoardDir(boardName)

	if err != nil {
		return err
	}

	// Check if the board exists
	if _, err := os.Stat(boardDir); os.IsNotExist(err) {
		return fmt.Errorf("board '%s' does not exist", boardName)
	}

	jsonDir := boardDir + "/tablist.json"

	bytes, _ := os.ReadFile(jsonDir)

	var tabs []models.Tab

	jsonErr := json.Unmarshal(bytes, &tabs)

	if jsonErr != nil {
		return jsonErr
	}
	// for _, tab := range tabs {
	// 	fmt.Printf("Tab ID: %s\nTitle: %s\nURL: %s\n\n", tab.TabID, tab.Title, tab.URL)
	// }
	if err := utils.PrintListStyle(tabs); err != nil {
		log.Fatalf("Failed to print table: %v", err)
	}

	return nil
}

func (tc *TabCommand) CloseTabById(tabId string) error {
	// Execute chrome-cli with the correct syntax: "close" followed by tab ID
	cmd := exec.Command("chrome-cli", "close", "-t", tabId)

	// Capture both stdout and stderr
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to close tab %s: %v\nOutput: %s", tabId, err, string(output))
	}

	fmt.Printf("Tab closed successfully: %s\n", string(tabId))
	return nil
}
