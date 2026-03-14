package cmd

import (
	"context"
	"fmt"
	"net/http"

	"github.com/divtxt/devhttps/internal/validate"
	"github.com/urfave/cli/v3"
)

func newHttpCommand() *cli.Command {
	return &cli.Command{
		Name:      "http",
		Usage:     "Run a simple HTTP server for testing",
		ArgsUsage: "[port]",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			port := "8000"
			if cmd.Args().Len() > 0 {
				port = cmd.Args().Get(0)
				if _, err := validate.Port(port); err != nil {
					fmt.Printf("Invalid port: %s\n", err)
					return cli.Exit("", 1)
				}
			}
			mux := http.NewServeMux()
			mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "hello (devhttps http %s)\n", port)
			})
			fmt.Printf("HTTP server running on http://127.0.0.1:%s\n", port)
			fmt.Printf("(use Ctrl-C to stop)\n")
			if err := http.ListenAndServe("127.0.0.1:"+port, mux); err != nil {
				fmt.Printf("Error: %s\n", err)
				return cli.Exit("", 1)
			}
			return nil
		},
	}
}
