package backend

import (
	. "github.com/starboychina/martini-mvc/src/helpers/utilities"
)

type Contrller struct {
	PathOptions
}

func (c Contrller) AdminIndexGet() {
	print("test")
}
