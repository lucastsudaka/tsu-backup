package pkg

import (
	"fmt"
	"github.com/spf13/cast"
	"os"
	"os/exec"
)

// TODO: Refactor to interface  | Too lazy now
func MariaDbBackup(config IAppConfig, targetDir string, session string) error {

	dbPort := config.GetConfigValueByKey("APP_BACKEND_MASTER_DB_PORT")
	dbPassword := config.GetConfigValueByKey("APP_BACKEND_MASTER_DB_PASSWORD")
	dbHost := config.GetConfigValueByKey("APP_BACKEND_MASTER_DB_HOST")
	dbUser := config.GetConfigValueByKey("APP_BACKEND_MASTER_DB_USER")

	backupTo := targetDir + "/tsu-backup_" + cast.ToString(session) + ".sql"

	fmt.Println("backupTo", backupTo)

	cmd := exec.Command("mariadb-dump", "-h"+dbHost, "-u"+dbUser, "-p"+dbPassword, "-P"+dbPort, "--lock-tables", "--all-databases", "--result-file="+backupTo)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("err", err)
	}

	return nil
}
