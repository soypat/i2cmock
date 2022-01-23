package i2cmock

import (
	"errors"
	"fmt"
	"sync"

	"periph.io/x/conn/v3/physic"
)

// compile time check of interface implementation
var _ i2c = (*Bus)(nil)

type Bus struct {
	// NoLock set to true disables
	// locking bus transaction.
	NoLock bool
	mu     sync.Mutex
	devs   []Peripheral
}

func (b *Bus) Add(p ...Peripheral) {
	b.devs = append(b.devs, p...)
}

func (b *Bus) Tx(addr uint16, w, r []byte) error {
	if len(b.devs) == 0 {
		return errors.New("no devices on bus. Did you close the bus before finishing?")
	}
	if !b.NoLock {
		b.mu.Lock()
		defer b.mu.Unlock()
	}
	for _, d := range b.devs {
		d.Tx(addr, w, r)
	}
	return nil
}

// Close should be called after finishing work.
func (b *Bus) Close() error {
	b.devs = b.devs[:0]
	return nil
}

func (b *Bus) String() (str string) {
	return fmt.Sprintf("i2cmock.bus{dev:%d}", len(b.devs))
}

func (b *Bus) SetSpeed(f physic.Frequency) error {
	if f == 0 {
		panic("zero frequency")
	}
	return nil
}

// ReadRegister transmits the register, restarts the connection as a read
// operation, and reads the response.
//
// Many I2C-compatible devices are organized in terms of registers. This method
// is a shortcut to easily read such registers. Also, it only works for devices
// with 7-bit addresses, which is the vast majority.
func (b *Bus) ReadRegister(address uint8, register uint8, data []byte) error {
	return b.Tx(uint16(address), []byte{register}, data)

}

// WriteRegister transmits first the register and then the data to the
// peripheral device.
//
// Many I2C-compatible devices are organized in terms of registers. This method
// is a shortcut to easily write to such registers. Also, it only works for
// devices with 7-bit addresses, which is the vast majority.
func (b *Bus) WriteRegister(address uint8, register uint8, data []byte) error {
	buf := make([]uint8, len(data)+1)
	buf[0] = register
	copy(buf[1:], data)
	return b.Tx(uint16(address), buf, nil)
}

func getReg(w, r []byte) (reg, length int) {
	lw := len(w)
	lr := len(r)
	switch {
	case lw == 0:
		// not a register operation
		reg = 0
		length = 0

	case lw != 1 && lr > 0:
		// Unhandled case.
		reg = 0
		length = 0

	default:
		// Normal case.
		reg = int(w[0])
		length = lr + lw - 1
	}
	return reg, length
}
