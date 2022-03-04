package i2cmock

import (
	"errors"
	"fmt"
	"sync"

	"periph.io/x/conn/v3"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/conn/v3/spi"
)

var _ spibus = (*SPIBus)(nil)

type SPIBus struct {
	// NoLock set to true disables
	// locking bus transaction.
	NoLock bool
	mu     sync.Mutex
	devs   []Peripheral
}

func (b *SPIBus) Add(p ...Peripheral) {
	b.devs = append(b.devs, p...)
}

func (s *SPIBus) Tx(w, r []byte) error {
	if len(s.devs) == 0 {
		return errors.New("no devices on bus. Did you close the bus before finishing?")
	}
	if !s.NoLock {
		s.mu.Lock()
		defer s.mu.Unlock()
	}
	for _, d := range s.devs {
		d.Tx(0, w, r)
	}
	return nil
}

// Close should be called after finishing work.
func (b *SPIBus) Close() error {
	b.devs = b.devs[:0]
	return nil
}

func (b *SPIBus) String() (str string) {
	return fmt.Sprintf("i2cmock.SPIBus{Ndevs:%d}", len(b.devs))
}

func (b *SPIBus) SetSpeed(f physic.Frequency) error {
	if f == 0 {
		panic("zero frequency")
	}
	return nil
}

func (b *SPIBus) TxPackets(packets []spi.Packet) error {
	for _, p := range packets {
		err := b.Tx(p.W, p.R)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *SPIBus) Duplex() conn.Duplex { return conn.Full }
func (b *SPIBus) Connect(f physic.Frequency, mode spi.Mode, bits int) (spi.Conn, error) {
	return b, nil
}
func (b *SPIBus) LimitSpeed(f physic.Frequency) error { return nil }
