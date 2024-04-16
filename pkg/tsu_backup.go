package pkg

import (
	"fmt"
	"github.com/spf13/cast"
	"math/rand/v2"
)

func TsuBackupMain(bParams BackupParams) {
	appConfig, _ := NewAppConfig()
	bParams.session = cast.ToString(rand.Uint64())

	fmt.Println(" INIT BACKUP " + bParams.session)
	//REFACTOR TO PARAMS, TOO LAZY NOW

	if *bParams.limitOfBackupFiles != 0 {
		ClearOutputFiles(bParams)
	}
	fmt.Println("backupFilePrependName", *bParams.backupFilePrependName)

	fmt.Println("sourceDir", *bParams.sourceDir)
	fmt.Println("targetDir", *bParams.targetDir)
	fmt.Println("isCompress", *bParams.isBackup)
	if *bParams.sourceDir == "" {
		fmt.Println("missing -sourceDir")
	}
	if *bParams.targetDir == "" {
		fmt.Println("missing -targetDir")
	}
	if *bParams.backupMariadb {
		MariaDbBackup(appConfig, bParams)
	}

	if *bParams.isBackup {
		Compress(bParams)
	}

	if *bParams.limitOfBackupFiles != 0 {
		ClearOutputFiles(bParams)
	}

	fmt.Println(" DONE BACKUP " + bParams.session)

	//appConfig
}
