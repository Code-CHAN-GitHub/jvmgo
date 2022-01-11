package control

import (
	"jvmgo/ch10/instructions/base"
	"jvmgo/ch10/rtda"
)

/**

	switch(i) {
		case 0: return 0;
		case 1: return 1;
		case 2: return 2;
		default: return -1;
	}

   0: iload_0
   1: tableswitch   { // 0 to 2
                 0: 28
                 1: 30
                 2: 32
           default: 34
      }
  28: iconst_0
  29: ireturn
  30: iconst_1
  31: ireturn
  32: iconst_2
  33: ireturn
  34: iconst_m1
  35: ireturn
*/

type TABLE_SWITCH struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffsets   []int32
}

func (self *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding() //  tableswitch 指令操作码的后面有 0 ～ 3 字节的 padding，以保证 defaultOffset 在字节码中的地址是 4 的倍数
	self.defaultOffset = reader.ReadInt32()
	self.low = reader.ReadInt32()
	self.high = reader.ReadInt32()
	jumpOffsetsCount := self.high - self.low + 1
	self.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (self *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()

	var offset int
	if index >= self.low && index <= self.high {
		offset = int(self.jumpOffsets[index-self.low])
	} else {
		offset = int(self.defaultOffset)
	}

	base.Branch(frame, offset)
}
