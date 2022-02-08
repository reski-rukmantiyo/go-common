package config

import (
	"os"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestLoadFullEnv(t *testing.T) {
	fileName := "/tmp/temporary.env"
	f, err := os.Create(fileName)
	if err != nil {
		t.Fatal(err)
	}
	content := "CacheDB=2\nCachePassword=\nCacheServer=localhost:6381"
	_, err = f.WriteString(content)
	if err != nil {
		t.Fatal(err)
	}
	err = LoadEnv(fileName, true)
	if err != nil {
		t.Fatal(err)
	}
	val, stat := os.LookupEnv("CacheDB")
	if !stat {
		t.Fatal("error")
	}
	assert.Equal(t, val, "2")
	val, stat = os.LookupEnv("CachePassword")
	if !stat {
		t.Fatal("error")
	}
	assert.Equal(t, val, "")
	val, stat = os.LookupEnv("CacheServer")
	if !stat {
		t.Fatal("error")
	}
	assert.Equal(t, val, "localhost:6381")
	os.Remove(fileName)
}

func TestLoadEnv(t *testing.T) {
	path, _ := GetPath()
	fileName := path + "/.lintasarta/temporary.env"
	f, err := os.Create(fileName)
	if err != nil {
		t.Fatal(err)
	}
	content := "CacheDB=2\nCachePassword=\nCacheServer=localhost:6381"
	_, err = f.WriteString(content)
	if err != nil {
		t.Fatal(err)
	}
	err = LoadEnv(fileName, true)
	if err != nil {
		t.Fatal(err)
	}
	val, stat := os.LookupEnv("CacheDB")
	if !stat {
		t.Fatal("error")
	}
	assert.Equal(t, val, "2")
	val, stat = os.LookupEnv("CachePassword")
	if !stat {
		t.Fatal("error")
	}
	assert.Equal(t, val, "")
	val, stat = os.LookupEnv("CacheServer")
	if !stat {
		t.Fatal("error")
	}
	assert.Equal(t, val, "localhost:6381")
	os.Remove(fileName)
}
