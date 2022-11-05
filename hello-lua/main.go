package main

import (
	"log"

	lua "github.com/yuin/gopher-lua"
)

func main() {
	L := lua.NewState()
	defer L.Close()

	if err := L.DoFile(`hello.lua`); err != nil {
		panic(err)
	}

	lv := L.GetGlobal("hello")
	if f, ok := lv.(*lua.LFunction); ok {
		log.Println(f.String())

		// nolint: exhaustivestruct
		if err := lua.NewState().CallByParam(lua.P{
			Fn:      f,
			NRet:    0,
			Protect: true,
		}); err != nil {
			log.Fatal(err)
		}
	}

	if lv.Type() != lua.LTFunction {
		log.Fatalf("function required but %s is given", lv.Type())
	}
}
