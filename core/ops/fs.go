package ops

import (
	"encoding/json"
	"io"
	"os"
	"time"

	"github.com/elsaland/elsa/core/options"
	"github.com/elsaland/elsa/util"

	"github.com/elsaland/quickjs"
	"github.com/spf13/afero"
)

type FsDriver struct {
	Perms *options.Perms
	Fs    afero.Fs
}

var _ io.Reader = (*os.File)(nil)

func (fs *FsDriver) ReadFile(ctx *quickjs.Context, path quickjs.Value) quickjs.Value {
	data, err := afero.ReadFile(fs.Fs, path.String())
	util.Check(err)
	return ctx.String(string(data))
}

func (fs *FsDriver) WriteFile(ctx *quickjs.Context, path quickjs.Value, content quickjs.Value) quickjs.Value {
	err := afero.WriteFile(fs.Fs, path.String(), []byte(content.String()), 0777)
	util.Check(err)
	return ctx.Bool(true)
}

func (fs *FsDriver) Exists(ctx *quickjs.Context, path quickjs.Value) quickjs.Value {
	data, err := afero.Exists(fs.Fs, path.String())
	util.Check(err)
	return ctx.Bool(data)
}

func (fs *FsDriver) DirExists(ctx *quickjs.Context, path quickjs.Value) quickjs.Value {
	data, err := afero.DirExists(fs.Fs, path.String())
	util.Check(err)
	return ctx.Bool(data)
}

func (fs *FsDriver) Cwd(ctx *quickjs.Context) quickjs.Value {
	dir, err := os.Getwd()
	util.Check(err)
	return ctx.String(dir)
}

type FileInfo struct {
	Name    string
	Size    int64
	Mode    os.FileMode
	ModTime time.Time
	IsDir   bool
}

func (fs *FsDriver) Stat(ctx *quickjs.Context, path quickjs.Value) quickjs.Value {
	entry, err := fs.Fs.Stat(path.String())
	util.Check(err)
	f := FileInfo{
		Name:    entry.Name(),
		Size:    entry.Size(),
		Mode:    entry.Mode(),
		ModTime: entry.ModTime(),
		IsDir:   entry.IsDir(),
	}
	output, err := json.Marshal(f)
	util.Check(err)
	return ctx.String(string(output))
}

func (fs *FsDriver) Remove(ctx *quickjs.Context, path quickjs.Value) quickjs.Value {
	err := fs.Fs.Remove(path.String())
	util.Check(err)
	return ctx.Bool(true)
}

func (fs *FsDriver) Mkdir(ctx *quickjs.Context, path quickjs.Value) quickjs.Value {
	err := fs.Fs.Mkdir(path.String(), os.FileMode(0777))
	util.Check(err)
	return ctx.Bool(true)
}
