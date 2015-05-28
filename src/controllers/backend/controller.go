package backend

import (
	. "../../helpers/utilities"
)

type Contrller struct {
	Options
}

func (c Contrller) AdminIndexGet() {
	print("test")
}
