package pkg

import (
	"context"
	"fmt"
	"github.com/klauspost/compress/zstd"
	"github.com/mholt/archiver/v4"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func Compress(bParams BackupParams) error {
	fmt.Println("Compress init")
	defer fmt.Println("Compress done ")

	dateX := time.Now().Format("2006-02-01_04-05")

	backupFileName := *bParams.backupFilePrependName + "__" + dateX
	backupFileName += ".tar.zst"

	filesToBackup := map[string]string{
		*bParams.sourceDir: "",
	}

	///INCLUDE FILES
	filesDir, err := ioutil.ReadDir(*bParams.targetDir)
	for _, value := range filesDir {
		if strings.Contains(value.Name(), bParams.session) {
			filesToBackup[*bParams.targetDir+"/"+value.Name()] = ""
		}
	}

	files, err := archiver.FilesFromDisk(nil, filesToBackup)
	if err != nil {
		return err
	}
	out, err := os.Create(*bParams.targetDir + "/" + backupFileName)
	if err != nil {
		return err
	}
	defer out.Close()

	// we can use the CompressedArchive type to gzip a tarball
	// (compression is not required; you could use Tar directly)

	format := archiver.CompressedArchive{
		Compression: archiver.Zstd{
			EncoderOptions: append([]zstd.EOption(nil), zstd.WithEncoderLevel(2), zstd.WithEncoderConcurrency(1)),
		},
		Archival: archiver.Tar{
			ContinueOnError: false,
		},
	}

	// create the archive
	err = format.Archive(context.Background(), out, files)
	if err != nil {
		return err
	}

	return nil

}
