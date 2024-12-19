package http

import (
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

const ModuleName = "http"

type HttpModule struct {
	runtime     *goja.Runtime
}

func (h *HttpModule) get(call goja.FunctionCall) goja.Value {
	res, err := Get(call.Argument(0).String())
	if err != nil {
		if _, ok := err.(*goja.Exception); !ok {
			panic(h.runtime.NewGoError(err))
		}
		panic(err)
	}
	ret := h.runtime.ToValue(res)
	return ret
}

func Require(runtime *goja.Runtime, module *goja.Object) {
	h := &HttpModule{
		runtime: runtime,
	}
	exports := module.Get("exports").(*goja.Object)
	exports.Set("get", h.get)
}

func Enable(runtime *goja.Runtime) {
	runtime.Set("http", require.Require(runtime, ModuleName))
}

func init() {
	require.RegisterCoreModule(ModuleName, Require)
}
