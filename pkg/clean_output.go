package pkg

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func CleanOutput(path string, limitOfBackupFiles int, backupFilePrependName string, session string) {
	fmt.Println("RUNNING \"CleanOutput\"")
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("err", err)
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].ModTime().After(files[j].ModTime())
	})

	backupFilesToRemove := []fs.FileInfo{}

	var countBackupFiles int = 0
	for _, value := range files {
		fmt.Println("countBackupFiles", countBackupFiles, "value", "limitOfBackupFiles", limitOfBackupFiles, value.Name())

		if strings.Contains(value.Name(), backupFilePrependName) && strings.Contains(value.Name(), "tar.zst") {
			countBackupFiles = countBackupFiles + 1
			if limitOfBackupFiles <= countBackupFiles {
				fmt.Println("Adding file to remove: ", value.Name())
				backupFilesToRemove = append(backupFilesToRemove, value)
			}
			//fmt.Println(index, value)
		}

		if strings.Contains(value.Name(), "tsu-backup") && strings.Contains(value.Name(), ".sql") {
			fmt.Println("Adding file to remove: ", value.Name())
			backupFilesToRemove = append(backupFilesToRemove, value)
		}
	}

	for _, value := range backupFilesToRemove {
		os.Remove(path + "/" + value.Name())
	}

}
