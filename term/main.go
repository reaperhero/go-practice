package main

import (
	"io"
	"log"
	"os"

	"github.com/pkg/term"
	"github.com/spf13/cobra"
)

func main() {
	var baud int
	var dev string

	cmd := &cobra.Command{
		Use:   "junkterm",
		Short: "junkterm is a quick and dirty serial terminal.",
		Run: func(cmd *cobra.Command, args []string) {
			term, err := term.Open(dev, term.Speed(baud), term.RawMode)
			if err != nil {
				log.Fatal(err)
			}
			go io.Copy(os.Stdout, term)
			io.Copy(term, os.Stdin)
		},
	}

	cmd.Flags().IntVarP(&baud, "baud", "b", 115200, "baud rate")
	cmd.Flags().StringVarP(&dev, "dev", "d", "/dev/tty", "device")

	cmd.Execute()
}