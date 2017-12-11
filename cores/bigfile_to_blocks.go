package cores

import (
	"errors"
	"math"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func BigFileToBlocks(bigfilePath, outputDir string, blockSize int64, ext string) error {
	if _, err := os.Stat(outputDir); err != nil {
		os.Mkdir(outputDir, os.ModePerm)
	}
	fileInfo, err := os.Stat(bigfilePath)
	if err != nil {
		return err
	}
	fileSize := float64(fileInfo.Size())
	num := int(math.Ceil(fileSize / float64(blockSize)))

	fi, err := os.OpenFile(bigfilePath, os.O_RDONLY, os.ModePerm)
	defer fi.Close()

	if err != nil {
		return err
	}
	b := make([]byte, blockSize)
	var i int64 = 1
	for ; i <= int64(num); i++ {

		fi.Seek((i-1)*int64(blockSize), 0)

		if len(b) > int((fileInfo.Size() - (i-1)*int64(blockSize))) {
			b = make([]byte, fileInfo.Size()-(i-1)*int64(blockSize))
		}

		fi.Read(b)

		blockPath := path.Join(outputDir, strconv.Itoa(int(i))+ext)
		f, err := os.OpenFile(blockPath, os.O_CREATE|os.O_WRONLY, os.ModePerm)

		if err != nil {
			return err
		}
		f.Write(b)
		f.Close()
	}
	return nil
}

func Blocks2BigFile(blocksPath, bigfilePath, ext string, blockSize int64) {
	// TODO:
}

func Blocks2BigFileByDir(blocksDirPath, bigfilePath string) error {
	blocks := make(map[int]os.FileInfo)
	var block_keys []int

	visit := func(path string, f os.FileInfo, err error) error {
		if f.IsDir() {
			return nil
		}
		nameArr := strings.Split(f.Name(), ".")
		if iVal, err := strconv.Atoi(nameArr[0]); err == nil {
			blocks[iVal] = f
			block_keys = append(block_keys, iVal)
		} else {
			return err
		}
		return nil
	}
	if err := filepath.Walk(blocksDirPath, visit); err != nil {
		return err
	}
	sort.Ints(block_keys)

	bigFile, err := os.OpenFile(bigfilePath, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return errors.New("open bigfile error " + err.Error())
	}
	defer bigFile.Close()
	for _, item := range block_keys {
		f, err := os.OpenFile(path.Join(blocksDirPath, blocks[item].Name()), os.O_RDONLY, os.ModePerm)
		if err != nil {
			return err
		}
		blockBytes := make([]byte, blocks[item].Size())
		f.Read(blockBytes)
		f.Close()

		bigFile.Write(blockBytes)
	}
	return nil
}
