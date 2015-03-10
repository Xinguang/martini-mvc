package backend

import (
	. "github.com/starboychina/martini-mvc/src/options"
)

type Contrller struct {
	PathOptions
}

func (c Contrller) AdminIndexGet() {
	print("test")
}
