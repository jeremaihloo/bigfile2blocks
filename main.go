package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kingpin"
	"github.com/jeremaihloo/bigfile2blocks/cores"
	// "fmt"
	"strconv"
)

const defaultBlockSize = 1024 * 1024

var (
	app = kingpin.New("bigfile2blocks", "A command-line bigfile2blocks application.")

	blocksCommand  = app.Command("blocks", "Blocks a big file")
	combineCommand = app.Command("combine", "Combine blocks into a big file")

	bigFilePath = app.Flag("bigfile", "Big file path.").Required().String()
	outPutDir   = app.Flag("outputs", "Number of packets to send").Required().String()
	hash        = app.Flag("hash", "Big file hash to check").Bool()

	debug   = app.Flag("debug", "Enable debug mode.").Bool()
	timeout = app.Flag("timeout", "Timeout waiting for ping.").Default("5s").OverrideDefaultFromEnvar("PING_TIMEOUT").Short('t').Duration()

	blockSize = app.Flag("block-size", "Block size").Default(strconv.Itoa(defaultBlockSize)).Int64()
	ext       = app.Flag("block-ext", "Block extension name").Default(".block").String()
)

func main() {
	app.Version("0.0.1")

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {

	case blocksCommand.FullCommand():
		if *hash {
			md5, _ := cores.Md5(*bigFilePath, *blockSize)
			fmt.Printf("hash: %s\n", md5)
		}
		err := cores.BigFileToBlocks(*bigFilePath, *outPutDir, *blockSize, *ext)
		if err != nil {
			if *debug {
				panic(err)
			}
			fmt.Println(err)
			return
		}
		fmt.Printf("blocks saved into %s\n", *outPutDir)

	case combineCommand.FullCommand():
		err := cores.Blocks2BigFileByDir(*outPutDir, *bigFilePath)
		if err != nil {
			if *debug {
				panic(err)
			}
			fmt.Println(err)
			return
		}
		if *hash {
			md5, _ := cores.Md5(*bigFilePath, *blockSize)
			fmt.Printf("hash: %s\n", md5)
		}
	}

}
