package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"rename2/utils"
)

var imgFolder string
var mdFile string

var rootCmd = &cobra.Command{
	Use:   "rename2",
	Short: "A tool to rename image files and update markdown file ",
	Run: func(cmd *cobra.Command, args []string) {
		if imgFolder == "" || mdFile == "" {
			fmt.Printf("Error: File Folder and TargetFile must be specified.\r\n\r\n")
			fmt.Println("rename2 tool's help: ")
			cmd.Help()
			return
		}

		utils.IHandler(imgFolder, mdFile)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&imgFolder, "imgFolder", "d", "", "Path to the image folder")
	rootCmd.PersistentFlags().StringVarP(&mdFile, "mdfile", "f", "", "Path to the markdown file")
}
