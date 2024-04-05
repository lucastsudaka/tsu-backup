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

func Compress(sourceDir string, destinationDir string, backupFilePrependName string, encodeLevel int, session string) error {
	fmt.Println("Compress init")
	defer fmt.Println("Compress done ")

	dateX := time.Now().Format("2006-02-01_04-05")

	backupFileName := backupFilePrependName + "__" + dateX
	backupFileName += ".tar.zst"

	filesToBackup := map[string]string{
		sourceDir: "",
	}

	///INCLUDE FILES
	filesDir, err := ioutil.ReadDir(destinationDir)
	for _, value := range filesDir {
		if strings.Contains(value.Name(), session) {
			filesToBackup[destinationDir+"/"+value.Name()] = ""
		}
	}

	files, err := archiver.FilesFromDisk(nil, filesToBackup)
	if err != nil {
		return err
	}
	out, err := os.Create(destinationDir + "/" + backupFileName)
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
	//
	//z := archiver.Zstd{
	//	EncoderOptions: nil,
	//	DecoderOptions: nil,
	//}.Op
	//
	////Tar: &archiver.Tar{
	////	OverwriteExisting:      false,
	////	MkdirAll:               false,
	////	ImplicitTopLevelFolder: false,
	////	StripComponents:        0,
	////	ContinueOnError:        false,
	////},
	//
	//z = z.O
	//err := z.Archive([]string{source}, destination+"/"+backupFileName)
	//if err != nil {
	//	fmt.Println("_________err", err)
	//}
	//z.Close()
	//if err != nil {
	//	fmt.Println("_________err", err)
	//}
	return nil

}
