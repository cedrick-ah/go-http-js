package main

import (
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/http"
	"github.com/dop251/goja_nodejs/require"
)

func main() {
	registry := new(require.Registry)
	
    runtime := goja.New()
	registry.Enable(runtime)
    http.Enable(runtime)
	console.Enable(runtime)

	runtime.RunString(`
		var data = http.get("https://www.google.com");
		console.log(data);
    `)
}
