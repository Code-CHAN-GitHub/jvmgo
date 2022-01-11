package lang

import (
	"jvmgo/ch09/native"
	"jvmgo/ch09/rtda"
	"math"
)

func init() {
	native.Register("java/lang/Double",
		"doubleToRawLongBits",
		"(D)J",
		doubleToRawLongBits)
}

// public static native long doubleToRawLongBits(double value);
// (D)J
func doubleToRawLongBits(frame *rtda.Frame) {
	value := frame.LocalVars().GetDouble(0)
	bits := math.Float64bits(value) // todo
	frame.OperandStack().PushLong(int64(bits))
}
