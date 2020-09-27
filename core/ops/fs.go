package ops

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/elsaland/elsa/cmd"
	"github.com/elsaland/quickjs"
	"github.com/spf13/afero"
)

type FsDriver struct {
	Perms cmd.Perms
	Fs    afero.Fs
}

var _ io.Reader = (*os.File)(nil)

func (fs *FsDriver) ReadFile(ctx *quickjs.Context, path quickjs.Value) quickjs.Value {
	data, err := afero.ReadFile(fs.Fs, path.String())
	if err != nil {
		log.Fatal(err)
	}
	return ctx.String(string(data))
}

func (fs *FsDriver) WriteFile(ctx *quickjs.Context, path quickjs.Value, content quickjs.Value) quickjs.Value {
	err := afero.WriteFile(fs.Fs, path.String(), []byte(content.String()), 0777)
	if err != nil {
		log.Fatal(err)
	}
	return ctx.Bool(true)
}

func (fs *FsDriver) Exists(ctx *quickjs.Context, path quickjs.Value) quickjs.Value {
	data, err := afero.Exists(fs.Fs, path.String())
	if err != nil {
		log.Fatal(err)
	}
	return ctx.Bool(data)
}

func (fs *FsDriver) DirExists(ctx *quickjs.Context, path quickjs.Value) quickjs.Value {
	data, err := afero.DirExists(fs.Fs, path.String())
	if err != nil {
		log.Fatal(err)
	}
	return ctx.Bool(data)
}

func (fs *FsDriver) Cwd(ctx *quickjs.Context) quickjs.Value {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("%v", err)
		os.Exit(1)
	}
	return ctx.String(dir)
}

type FileInfo struct {
	Name    string
	Size    int64
	Mode    os.FileMode
	ModTime time.Time
	IsDir   bool
}

func (fs *FsDriver) Stats(ctx *quickjs.Context, path quickjs.Value) quickjs.Value {
	entry, err := fs.Fs.Stat(path.String())
	if err != nil {
		fmt.Println("%v", err)
		os.Exit(1)
	}
	f := FileInfo{
		Name:    entry.Name(),
		Size:    entry.Size(),
		Mode:    entry.Mode(),
		ModTime: entry.ModTime(),
		IsDir:   entry.IsDir(),
	}
	output, err := json.Marshal(f)
	if err != nil {
		log.Fatal(err)
	}
	return ctx.String(string(output))
}

func (fs *FsDriver) Remove(ctx *quickjs.Context, path quickjs.Value) quickjs.Value {
	err := fs.Fs.Remove(path.String())
	if err != nil {
		fmt.Println("%v", err)
		os.Exit(1)
	}
	return ctx.Bool(true)
}
