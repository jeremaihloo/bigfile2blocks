package cores

import (
	"crypto/md5"
	"fmt"
	"io"
	"math"
	"os"
)

func Md5SmallFile(path string) (md5Value string, err error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	md5h := md5.New()
	io.Copy(md5h, file)
	return fmt.Sprintf("%x", md5h.Sum([]byte(""))), nil //md5
}

func Md5BigFile(path string, filechunk int64) (md5Value string, err error) {
	file, err := os.Open(path)

	if err != nil {
		panic(err.Error())
	}

	defer file.Close()

	// calculate the file size
	info, _ := file.Stat()

	filesize := info.Size()

	blocks := int(math.Ceil(float64(filesize) / float64(filechunk)))

	hash := md5.New()

	for i := 0; i < blocks; i++ {
		r := int64(i) * filechunk
		blocksize := int(math.Min(float64(filechunk), float64(filesize-int64(r))))
		buf := make([]byte, blocksize)

		file.Read(buf)
		io.WriteString(hash, string(buf)) // append into the hash
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func Md5(path string, filechunk int64) (md5Val string, err error) {
	f, err := os.Stat(path)
	if f.Size() > filechunk {
		return Md5BigFile(path, filechunk)
	}
	return Md5SmallFile(path)
}
