package yaegi_interp

import (
	"reflect"
	"tsuki/extensions/yaegi_extract"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

var INTERP *interp.Interpreter

func SetupInterp() {
	i := interp.New(interp.Options{Unrestricted: false})

	symbols := stdlib.Symbols
	delete(symbols, "archive/tar/tar")
	delete(symbols, "archive/zip/zip")
	delete(symbols, "compress/gzip/gzip")
	delete(symbols, "compress/zlib/zlib")
	delete(symbols, "io/fs/fs")
	delete(symbols, "io/ioutil/ioutil")
	delete(symbols, "os/exec/exec")
	delete(symbols, "os/os")
	delete(symbols, "os/signal/signal")
	delete(symbols, "os/user/user")
	delete(symbols, "runtime/runtime")
	delete(symbols, "syscall/syscall")
	i.Use(symbols)
	i.Use(yaegi_extract.Symbols)

	INTERP = i
}

func EvaluateProvider(script string) (reflect.Value, error) {
	_, err := INTERP.Eval(script)
	if err != nil {
		return reflect.Value{}, err
	}

	v, err := INTERP.Eval("main.NewProvider")
	if err != nil {
		return reflect.Value{}, err
	}

	return v, err
}
