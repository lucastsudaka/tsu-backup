package pkg

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func ClearOutputFiles(bParams BackupParams) {
	fmt.Println("RUNNING \"ClearOutputFiles\"")
	files, err := ioutil.ReadDir(*bParams.targetDir)
	if err != nil {
		fmt.Println("err", err)
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].ModTime().After(files[j].ModTime())
	})

	backupFilesToRemove := []fs.FileInfo{}

	var countBackupFiles int = 0
	for _, value := range files {
		fmt.Println("countBackupFiles", countBackupFiles, "value", "limitOfBackupFiles", *bParams.limitOfBackupFiles, value.Name())

		if strings.Contains(value.Name(), *bParams.backupFilePrependName) && strings.Contains(value.Name(), "tar.zst") {
			countBackupFiles = countBackupFiles + 1
			if *bParams.limitOfBackupFiles < countBackupFiles {
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
		os.Remove(*bParams.targetDir + "/" + value.Name())
	}

}
