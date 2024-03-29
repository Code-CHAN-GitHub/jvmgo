package classfile

/**
CONSTANT_NameAndType_info {
	u1 tag;
	u2 name_index;
	u2 description_index;
}
*/

type ConstantNameAndTypeInfo struct {
	nameIndex        uint16
	descriptionIndex uint16
}

func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	self.descriptionIndex = reader.readUint16()
}
