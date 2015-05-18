package backend

import (
	. "../../helpers/utilities"
)

type Contrller struct {
	PathOptions
}

func (c Contrller) AdminIndexGet() {
	print("test")
}
