package pkg

import (
	"flag"
	"fmt"
	"github.com/spf13/cast"
	"math/rand/v2"
)

func Init() {
	appConfig, _ := NewAppConfig()
	backupFilePrependName := flag.String("prependName", "tsu-backup", " TODO  ")

	sourceDir := flag.String("sourceDir", "/backup-source", " TODO  ")
	targetDir := flag.String("targetDir", "/backups", " TODO  ")
	isBackup := flag.Bool("backup", true, " TODO ")
	limitOfBackupFiles := flag.Int("limitOfBackupFiles", 3, " TODO ")
	encodeLevel := flag.Int("encodeLevel", 2, " TODO  ")
	backupMariadb := flag.Bool("backupMariadb", true, " TODO  ")

	flag.Parse()
	session := cast.ToString(rand.Int())

	if *limitOfBackupFiles != 0 {
		CleanOutput(*targetDir, *limitOfBackupFiles, *backupFilePrependName, session)
	}
	fmt.Println("backupFilePrependName", *backupFilePrependName)

	fmt.Println("sourceDir", *sourceDir)
	fmt.Println("targetDir", *targetDir)
	fmt.Println("isCompress", *isBackup)
	if *sourceDir == "" {
		fmt.Println("missing -sourceDir")
	}
	if *targetDir == "" {
		fmt.Println("missing -targetDir")
	}
	if *backupMariadb {
		MariaDbBackup(appConfig, *targetDir, session)
	}

	if *isBackup {
		Compress(*sourceDir, *targetDir, *backupFilePrependName, *encodeLevel, session)
	}

	if *limitOfBackupFiles != 0 {
		CleanOutput(*targetDir, *limitOfBackupFiles, *backupFilePrependName, session)
	}

	//appConfig

}
