# Nodejs compatibility library for Goja

This is a collection of Goja modules(from goja_nodejs + a http module that I added here because a was not able to add it there) that provide nodejs compatibility.

Example:

```go
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

```
