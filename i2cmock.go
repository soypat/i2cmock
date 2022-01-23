package i2cmock

import "io"

type Peripheral interface {
	Tx(addr uint16, r, w []byte)
}

// I2C represents an I2C bus. It is notably implemented by the
// machine.I2C type. It was brought over from the tinygo/drivers I2C
// interface at https://github.com/tinygo-org/drivers
type i2c interface {
	io.Closer
	ReadRegister(addr uint8, r uint8, buf []byte) error
	WriteRegister(addr uint8, r uint8, buf []byte) error
	Tx(addr uint16, w, r []byte) error
}
