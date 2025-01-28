package specs

import (
	"testing"

	"github.com/abhishek71994/tab-manager-cli/cmd/cli"
	"github.com/abhishek71994/tab-manager-cli/specs/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNewApp(t *testing.T) {
	// Setup
	fileStore := mocks.NewMockFileStore()

	// Execute
	app := cli.NewApp(fileStore)

	// Assert basic app properties
	assert.Equal(t, "tabman", app.Name)
	assert.Equal(t, "A CLI tool to manage tabs.", app.Usage)

	// Test command configuration
	tests := []struct {
		name         string
		commandName  string
		aliasName    string
		expectedFlag string
	}{
		{
			name:         "create-board command",
			commandName:  "create-board",
			aliasName:    "cb",
			expectedFlag: "name",
		},
		{
			name:        "list-boards command",
			commandName: "list-boards",
			aliasName:   "lb",
		},
		{
			name:         "delete-board command",
			commandName:  "delete-board",
			aliasName:    "deb",
			expectedFlag: "name",
		},
		{
			name:         "list-tab command",
			commandName:  "list-tab",
			aliasName:    "lt",
			expectedFlag: "name",
		},
		{
			name:         "close-tab command",
			commandName:  "close-tab",
			aliasName:    "ct",
			expectedFlag: "id",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Find command
			var cmd *cli.Command
			for _, c := range app.Commands {
				if c.Name == tt.commandName {
					cmd = c
					break
				}
			}

			// Assert command exists
			assert.NotNil(t, cmd, "Command %s should exist", tt.commandName)

			// Assert alias
			assert.Contains(t, cmd.Aliases, tt.aliasName, "Command should have correct alias")

			// Assert action is configured
			assert.NotNil(t, cmd.Action, "Command should have an action configured")

			// Assert flags if expected
			if tt.expectedFlag != "" {
				var flagFound bool
				for _, flag := range cmd.Flags {
					if flag.Names()[0] == tt.expectedFlag {
						flagFound = true
						break
					}
				}
				assert.True(t, flagFound, "Command should have %s flag", tt.expectedFlag)
			}
		})
	}
}
