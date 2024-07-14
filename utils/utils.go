package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func IHandler(imgFolder string, mdFile string) {

	files, err := os.ReadDir(imgFolder)
	if err != nil {
		fmt.Println("[-] Error reading img folder:", err)
		return
	}

	var pnames []string
	for _, file := range files {
		pnames = append(pnames, file.Name())
	}

	mdContent, err := readMDFile(mdFile)
	if err != nil {
		fmt.Println("[-] Error reading markdown file:", err)
		return
	}

	for _, pname := range pnames {
		offset := strings.Index(mdContent, pname)
		if offset != -1 {
			fmt.Printf("[+] target img file %s's offset: %d\n", pname, offset)
		}
	}

	newFilenames := make(map[string]string)
	for i, oldFilename := range pnames {
		extension := filepath.Ext(oldFilename)
		newFilename := fmt.Sprintf("%d%s", i+1, extension)
		newFilenames[oldFilename] = newFilename
	}

	for oldFilename, newFilename := range newFilenames {

		oldFilepath := filepath.Join(imgFolder, oldFilename)
		newFilepath := filepath.Join(imgFolder, newFilename)
		if err := os.Rename(oldFilepath, newFilepath); err != nil {
			fmt.Println("[-] Error renaming file:", err)
			return
		}

		mdContent = strings.ReplaceAll(mdContent, oldFilename, newFilename)
	}

	if err := updateMDFile(mdFile, mdContent); err != nil {
		fmt.Println("[-] Error updating markdown file:", err)
		return
	}

	fmt.Printf("[!] Target Folder Files renamed and Target markdown file updated successfully.\r\n")
}

func readMDFile(filepath string) (string, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func updateMDFile(filepath, content string) error {
	return os.WriteFile(filepath, []byte(content), 0644)
}
