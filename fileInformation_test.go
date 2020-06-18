package yaccidl

import (
	"github.com/stretchr/testify/assert"
	"os/user"
	"path/filepath"
	"testing"
)

func TestFileInformation_Add(t *testing.T) {
	t.Run("", func(t *testing.T) {
		fi := NewFileInformation("/")
		assert.Equal(t, "/", fi.basePath)
		entry := fi.AddPath("/")
		assert.Equal(t, "/", entry.Path)

		entry = fi.AddPath("/aa")
		assert.Equal(t, "/aa", entry.Path)
	})

	t.Run("", func(t *testing.T) {
		fi := NewFileInformation("./dddd/fn.txt")
		assert.Equal(t, "dddd", fi.basePath)

		entry := fi.AddPath("/")
		assert.Equal(t, "/", entry.Path)

		entry = fi.AddPath("./aa")
		assert.Equal(t, "dddd/aa", entry.Path)

		entry = fi.AddPath("./..")
		assert.Equal(t, ".", entry.Path)

		entry = fi.AddPath("./../..")
		assert.Equal(t, "..", entry.Path)

		entry = fi.AddPath("./../../aa")
		assert.Equal(t, "../aa", entry.Path)
	})
}

func TestFileInformation_New(t *testing.T) {
	t.Run("", func(t *testing.T) {
		fi := NewFileInformation("/")
		assert.Equal(t, "/", fi.basePath)
	})

	t.Run("", func(t *testing.T) {
		fi := NewFileInformation("/abc/sss/fn.txt")
		assert.Equal(t, "/abc/sss", fi.basePath)
	})
	t.Run("", func(t *testing.T) {
		fi := NewFileInformation("")
		assert.Equal(t, ".", fi.basePath)
	})

	t.Run("", func(t *testing.T) {
		fi := NewFileInformation("./ss/fn.txt")
		assert.Equal(t, "ss", fi.basePath)
	})

	t.Run("", func(t *testing.T) {
		fi := NewFileInformation("../../")
		assert.Equal(t, "../..", fi.basePath)
	})
	t.Run("", func(t *testing.T) {
		fi := NewFileInformation("~/../fn.txt")
		usr, _ := user.Current()
		homeDir := usr.HomeDir
		assert.Equal(t, filepath.Dir(homeDir), fi.basePath)
	})
}
