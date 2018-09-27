package sta

import (
	"os"
	"time"
	"strings"
	"path/filepath"
)

const Prefix = "/epj/www"
// const Prefix = "/Users/amyli/go/src/eglass.com"

type FileStat struct {
	Name string
	Size int64 
	ModTime time.Time
	isDir bool
}
type FilesStats map[string]FileStat


func GetWalkFn(prefix string, dirStats FilesStats) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error)error {
		if err != nil {
			return err
		}
		if !strings.Contains(path, ".git") {
			relativePath := path[len(prefix):]
			dirStats[relativePath] = FileStat{
				Name: info.Name(),
				Size: info.Size(),
			}
		}
		return nil
	}
}



func Scan(dir string)(FilesStats,error) {
	var dirStats = make(FilesStats)
	error := filepath.Walk(Prefix + dir, GetWalkFn(Prefix + dir, dirStats))
	if (error != nil) { return nil, error }
	return dirStats, nil
}