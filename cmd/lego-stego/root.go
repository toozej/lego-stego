package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "lego-stego",
	Short: "CLI wrapped steganography tool",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
