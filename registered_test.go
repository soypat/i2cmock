package i2cmock

import "testing"

func TestMPU6050(t *testing.T) {
	const (
		MPU6050Addr        = 0x68
		ACCEL_XOUT_H uint8 = 0x3B
	)
	registers := [128]byte{
		ACCEL_XOUT_H: 0,
	}
	mpu6050 := NewRegistered(MPU6050Addr, registers[:])
	var b Bus
	b.Add(mpu6050)
	var buf [8]byte
	b.ReadRegister(MPU6050Addr, ACCEL_XOUT_H, buf[:])
	t.Logf("read registers, got %x", buf)
	for i := range buf {
		buf[i] = uint8(i)
	}
	t.Logf("writing to registers %x", buf)
	b.WriteRegister(MPU6050Addr, ACCEL_XOUT_H, buf[:])
	buf = [8]byte{} // zero out memory to make sure we don't just print previous value.
	b.ReadRegister(MPU6050Addr, ACCEL_XOUT_H, buf[:])
	t.Logf("read back registers1 %x", buf)
	buf = [8]byte{}
	b.ReadRegister(MPU6050Addr, ACCEL_XOUT_H, buf[:])
	t.Logf("read back registers2 %x", buf)

}
