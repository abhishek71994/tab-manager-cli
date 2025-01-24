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

	// Define the CLI app
	app := &cli.App{
		Name:  "tabman", // Updated to match the binary name
		Usage: "A CLI tool to manage boards, categories, and tabs.",
		Commands: []*cli.Command{
			{
				Name:  "create-board",
				Usage: "Create a new board",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "name",
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
				Name:  "list-boards",
				Usage: "List all boards",
				Action: func(c *cli.Context) error {
					return boardCmd.ListBoards()
				},
			},
			{
				Name:  "delete-board",
				Usage: "Delete a board by name",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "name",
						Usage:    "Name of the board to delete",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					boardName := c.String("name")
					return boardCmd.DeleteBoard(boardName)
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
