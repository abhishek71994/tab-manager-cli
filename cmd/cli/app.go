// internals/commands/tab.go
package cli

import (
	"log"
	"os"

	"github.com/abhishek71994/tab-manager-cli/internal/commands"
	"github.com/abhishek71994/tab-manager-cli/internal/storage"
	"github.com/urfave/cli/v2"
)

func NewApp(fileStore *storage.FileStore) *cli.App {
	// Initialize the BoardCommand
	boardCmd := commands.NewBoardCommand(fileStore)
	tabCmd := commands.NewTabCommand(fileStore)

	// Define the CLI app
	app := &cli.App{
		Name:  "tabman", // Updated to match the binary name
		Usage: "A CLI tool to manage tabs.",
		Commands: []*cli.Command{
			{
				Name:    "create-board",
				Usage:   "Create a new board",
				Aliases: []string{"cb"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "name",
						Aliases:  []string{"n"},
						Usage:    "Name of the board",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					boardName := c.String("name")
					return boardCmd.CreateBoard(boardName)
				},
			},
			{
				Name:    "list-boards",
				Usage:   "List all boards",
				Aliases: []string{"lb"},
				Action: func(c *cli.Context) error {
					return boardCmd.ListBoards()
				},
			},
			{
				Name:    "delete-board",
				Usage:   "Delete a board by name",
				Aliases: []string{"deb"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "name",
						Aliases:  []string{"n"},
						Usage:    "Name of the board to delete",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					boardName := c.String("name")
					return boardCmd.DeleteBoard(boardName)
				},
			},
			{
				Name:    "list-tab",
				Usage:   "List tabs inside a board",
				Aliases: []string{"lt"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "name",
						Aliases:  []string{"n"},
						Usage:    "Name of the board",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					boardName := c.String("name")
					return tabCmd.ListTabs(boardName)
				},
			},
			{
				Name:    "close-tab",
				Usage:   "close tab by id",
				Aliases: []string{"ct"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "id",
						Usage:    "Tab ID",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					tabId := c.String("id")
					return tabCmd.CloseTabById(tabId)
					//maybe delete the entry from the file
				},
			},
			{
				Name:    "list-current-tabs",
				Usage:   "list all the current tabs on chrome",
				Aliases: []string{"curt"},
				Action: func(c *cli.Context) error {
					return tabCmd.ListCurrentTabs()
				},
			},
		},
	}

	return app
}

func Run() {
	// Initialize file storage
	fileStore := storage.NewFileStore()

	// Create and run the CLI app
	app := NewApp(fileStore)
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
