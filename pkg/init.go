package pkg

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type BackupParams struct {
	sourceDir             *string
	targetDir             *string
	isBackup              *bool
	backupFilePrependName *string
	limitOfBackupFiles    *int
	encodeLevel           *int
	backupMariadb         *bool
	crontab               *string
	runNow                *bool
	session               string
}

func Init() {
	defer fmt.Println("BYE FROM TSU-BACKUP")
	bParams := BackupParams{}
	bParams.sourceDir = flag.String("sourceDir", "/backup-source", " TODO  ")
	bParams.targetDir = flag.String("targetDir", "/backups", " TODO  ")
	bParams.isBackup = flag.Bool("backup", true, " TODO ")
	bParams.limitOfBackupFiles = flag.Int("limitOfBackupFiles", 5, " TODO ")
	bParams.encodeLevel = flag.Int("encodeLevel", 2, " TODO  ")
	bParams.backupMariadb = flag.Bool("backupMariadb", true, " TODO  ")
	bParams.backupFilePrependName = flag.String("prependName", "tsu-backup", " TODO  ")
	bParams.crontab = flag.String("crontab", "0 * * * *", " TODO  ") //DEFAULT 0 * * * * = every hour https://crontab.cronhub.io/
	bParams.runNow = flag.Bool("runNow", true, " TODO  ")
	flag.Parse()

	tsuBackup := func() {
		TsuBackupMain(bParams)
	}
	scheduler, job, err := Job(tsuBackup, bParams)
	if err != nil {
		panic(err)
	}

	if *bParams.runNow {
		time.Sleep(time.Second)
		job.RunNow()
	}

	if *bParams.crontab == "" {
		scheduler.Shutdown()
		return
	} else {
		keepRunning()
		scheduler.Shutdown()
	}

}

func keepRunning() {
	go forever()
	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel
}
func forever() {
	for {
		time.Sleep(time.Second)
	}
}
