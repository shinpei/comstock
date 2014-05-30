package engine

import (
	"fmt"
)

func (e *Engine) Status() {
	fmt.Println("[Comstock envieonment]")
	m := e.env.AsMap()
	for k, v := range m {
		fmt.Println(k, ":", v)
	}
	fmt.Println("")
	fmt.Println("[Storager environment]")
	_ = e.storager.Status()
}
