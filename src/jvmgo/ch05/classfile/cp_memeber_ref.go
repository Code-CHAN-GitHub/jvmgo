package classfile

/**
CONSTANT_Fieldref_info {
	u1 tag;
	u2 class_index;
	u2 name_and_type_index;
}

CONSTANT_Fieldref_info、CONSTANT_Methodref_info 和 CONSTANT_InterfaceMethodref_info 结构相同
*/

type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (self *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}

func (self *ConstantMemberrefInfo) ClassName() string {
	return self.cp.getClassName(self.classIndex)
}

func (self *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}

type ConstantFieldrefInfo struct {
	ConstantMemberrefInfo
}

//func (self ConstantFieldrefInfo) readInfo(reader *ClassReader) {
//	self.cpMemberInfo.readInfo(reader)
//}

type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo
}

//func (self ConstantMethodrefInfo) readInfo(reader *ClassReader) {
//	self.cpMemberInfo.readInfo(reader)
//}

type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberrefInfo
}

//func (self ConstantInterfaceMethodrefInfo) readInfo(reader *ClassReader) {
//	self.cpMemberInfo.readInfo(reader)
//}
