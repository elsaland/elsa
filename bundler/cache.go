package bundler

import (
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type ElsaCache struct {
	dir string
}

func (cache *ElsaCache) BuildFileName(uri string) string {
	fileUrl, _ := url.Parse(uri)
	path := path.Join(cache.dir, fileUrl.Host, fileUrl.Path)
	return path
}

func (cache *ElsaCache) PathToUrl(path string) string {
	parts := strings.Split(path, "/")[2:]
	url, _ := url.Parse("https://" + strings.Join(parts, "/"))
	return url.String()
}

func (cache *ElsaCache) UrlToPath(url string) string {
	parts := strings.Split(url, "//")[1]
	path := path.Join(cache.dir, parts)
	return path
}

func (cache *ElsaCache) InCache(path string) bool {
	return strings.HasPrefix(path, cache.dir)
}

func (cache *ElsaCache) Exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func (cache *ElsaCache) Create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0750); err != nil {
		return nil, err
	}
	return os.Create(p)
}
