package interp

import (
	"reflect"
	"tsuki/backend/yaegi/extract"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

func NewInterp() *interp.Interpreter {
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
	i.Use(extract.Symbols)

	return i
}

func EvaluateProvider(script string) (reflect.Value, error) {
	interpreter := NewInterp()
	_, err := interpreter.Eval(script)
	if err != nil {
		return reflect.Value{}, err
	}

	v, err := interpreter.Eval("main.NewProvider")
	if err != nil {
		return reflect.Value{}, err
	}

	return v, err
}
