package classfile

/**
CONSTANT_Class_info {
	u1 tag;
	u2 name_index;
}

与 CONSTANT_String_info 类似，并不存储实际值，而是指向字符串常量池中 utf8 的索引
类、超类以及接口的索引指向的都是 CONSTANT_Class_info
*/

type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
}

// 按索引从常量池中查找字符串
func (self *ConstantClassInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}

//func (self *ConstantClassInfo) String() string {
//	return self.cp.getUtf8(self.nameIndex)
//}
