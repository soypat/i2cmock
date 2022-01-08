package i2cmock

type Registered struct {
	addr uint16
	reg  []byte
}

func NewRegistered(addr uint16, regData []byte) Registered {
	return Registered{
		addr: addr,
		reg:  regData,
	}
}

func (rg Registered) Tx(addr uint16, w, r []byte) {
	if addr != rg.addr {
		// address does not correspond to peripheral receiver.
		return
	}
	reg, plen := getReg(w, r)
	switch {
	case plen == 0:
		// No work to do.

	case plen+reg > len(rg.reg):
		panic("register request exceeds memory map")

	case len(r) > 0:
		copy(r, rg.reg[reg:])

	case len(r) == 0:
		copy(rg.reg[reg:], w[1:])
	}
}
