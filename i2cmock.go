package i2cmock

import (
	"io"

	"periph.io/x/conn/v3/physic"
)

type Peripheral interface {
	Tx(addr uint16, r, w []byte)
}

// I2C represents an I2C bus. It is notably implemented by the
// machine.I2C type. It was brought over from the tinygo/drivers I2C
// interface at https://github.com/tinygo-org/drivers
type i2c interface {
	io.Closer
	periphBus
	ReadRegister(addr uint8, r uint8, buf []byte) error
	WriteRegister(addr uint8, r uint8, buf []byte) error
	Tx(addr uint16, w, r []byte) error
}

type spibus interface {
	io.Closer
	periphBus
	Tx(w, r []byte) error
}
type periphBus interface {
	String() string
	SetSpeed(f physic.Frequency) error
}
