package classfile

/**
CONSTANT_String_info {
	u1 tag;
	u2 string_index;
}

CONSTANT_String_info 本身不存放字符串数据，只存了常量池索引，
这个索引指向一个 CONSTANT_Utf8_info 常量
*/

type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.stringIndex = reader.readUint16()
}

// 按索引从常量池中查找字符串
func (self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.stringIndex)
}
