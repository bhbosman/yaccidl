package yaccidl

import (
	"fmt"
	"go.uber.org/fx"
	"os/user"
	"path/filepath"
	"strings"
	"sync"
)

type FileInformationEntry struct {
	Path     string
	FolderId string
}

type FileInformation struct {
	entries map[string]FileInformationEntry
	counter map[string]int
	mutex sync.Mutex
}

func NewFileInformation() *FileInformation {
	fi := &FileInformation{
		entries: make(map[string]FileInformationEntry),
		counter: make(map[string]int),
	}
	return fi
}

func ProvideFileInformation() fx.Option {
	return fx.Provide(func() *FileInformation {
		return NewFileInformation()
	})
}

func (self *FileInformation) AddPath(path string) FileInformationEntry {
	self.mutex.Lock()
	defer  self.mutex.Unlock()
	if entry, ok := self.entries[path]; ok {
		return entry
	}

	if !filepath.IsAbs(path) {
		usr, _ := user.Current()
		homeDir := usr.HomeDir
		if path == "~" {
			path = homeDir
		} else if strings.HasPrefix(path, "~/") {
			path = filepath.Join(homeDir, path[2:])
		}
	}
	if entry, ok := self.entries[path]; ok {
		return entry
	}

	aliasName := filepath.Base(path)
	count := 1
	var ok bool
	if count, ok = self.counter[aliasName]; ok {
		count++
	}
	self.counter[aliasName] = count
	folderId := fmt.Sprintf("__%v__", aliasName)
	if count > 1 {
		folderId = fmt.Sprintf("__%v_%d__", aliasName, count)
	}

	entry := FileInformationEntry{
		Path:     path,
		FolderId: folderId,
	}
	self.entries[path] = entry

	return entry
}
