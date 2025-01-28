// internal/utils/chrome_tabs.go
package utils

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/abhishek71994/tab-manager-cli/internal/models"
)

// FetchChromeTabs fetches all open tabs in Chrome using chrome-cli.
func FetchChromeTabs() ([]models.Tab, error) {
	// Run chrome-cli list tabs to get all open tabs
	cmd := exec.Command("chrome-cli", "list", "tabs")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to run chrome-cli: %v", err)
	}

	// Parse the output
	lines := strings.Split(string(output), "\n")
	var tabs []models.Tab
	for _, line := range lines {
		if line == "" {
			continue
		}

		// Extract tabId from the line
		if !strings.HasPrefix(line, "[") {
			continue
		}
		parts := strings.Split(line, "]")
		if len(parts) < 2 {
			continue
		}
		tabId := strings.Split(parts[0], ":")[1]

		// Fetch tab details using chrome-cli info -t <tabId>
		infoCmd := exec.Command("chrome-cli", "info", "-t", tabId)
		infoOutput, err := infoCmd.Output()
		if err != nil {
			log.Printf("Failed to fetch info for tab %s: %v", tabId, err)
			continue
		}

		// log.Printf("chrome-cli info -t %s output: %s", tabId, string(infoOutput)) // Debug log

		// Parse the tab details
		tab := parseTabInfo(tabId, string(infoOutput))
		if tab != nil {
			tabs = append(tabs, *tab)
		}
	}

	return tabs, nil
}

// parseTabInfo parses the output of chrome-cli info -t <tabId>.
func parseTabInfo(tabId string, info string) *models.Tab {
	lines := strings.Split(info, "\n")
	var tab models.Tab
	tab.TabID = tabId // Set the tabId
	for _, line := range lines {
		if strings.HasPrefix(line, "Title:") {
			tab.Title = strings.TrimSpace(strings.TrimPrefix(line, "Title:"))
		} else if strings.HasPrefix(line, "Url:") {
			tab.URL = strings.TrimSpace(strings.TrimPrefix(line, "Url:"))
		}
	}

	if tab.Title == "" || tab.URL == "" {
		return nil
	}

	return &tab
}

// CheckChromeCLI verifies if chrome-cli is installed and working.
func CheckChromeCLI() error {
	// Check if chrome-cli is installed
	if _, err := exec.LookPath("chrome-cli"); err != nil {
		return fmt.Errorf("chrome-cli is not installed. Please install it using 'brew install chrome-cli'")
	}

	// Check if chrome-cli is working by running a simple command
	cmd := exec.Command("chrome-cli", "list", "tabs")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("chrome-cli is not working. Ensure Google Chrome is running and chrome-cli has the necessary permissions")
	}

	return nil
}
