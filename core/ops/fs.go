package ops

import (
	"encoding/binary"
	"io"
	"os"

	"github.com/elsaland/elsa/cmd"
	"github.com/lithdew/quickjs"
	"github.com/spf13/afero"
)

type FsDriver struct {
	Perms cmd.Perms
	Fs    afero.Fs
}

var _ io.Reader = (*os.File)(nil)

func (fs *FsDriver) CreateFS(path string) afero.Fs {
	appfs := afero.NewOsFs()
	return appfs
}

func (fs *FsDriver) ReadFile(ctx quickjs.Context, path quickjs.Value) quickjs.Value {
	data, err := afero.ReadFile(fs.Fs, path.String())
	if err != nil {
		return ctx.ThrowError(err)
	}
	return ctx.Uint32(binary.LittleEndian.Uint32(data))
}

func (fs *FsDriver) Exists(ctx quickjs.Context, path quickjs.Value) quickjs.Value {
	data, err := afero.Exists(fs.Fs, path.String())
	if err != nil {
		return ctx.ThrowError(err)
	}
	return ctx.Bool(data)
}

func (fs *FsDriver) DirExists(ctx quickjs.Context, path quickjs.Value) quickjs.Value {
	data, err := afero.DirExists(fs.Fs, path.String())
	if err != nil {
		return ctx.ThrowError(err)
	}
	return ctx.Bool(data)
}
