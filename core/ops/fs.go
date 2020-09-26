package ops

import (
  "io"
  "log"
  "os"

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

func (fs *FsDriver) DirExists(ctx quickjs.Context, path quickjs.Value) quickjs.Value {
	data, err := afero.DirExists(fs.Fs, path.String())
	if err != nil {
    log.Fatal(err)
	}
	return ctx.Bool(data)
}
