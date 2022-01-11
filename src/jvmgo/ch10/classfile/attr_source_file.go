package classfile

/**
SourceFile_attribute {
	u2 	attribute_name_index;
	u4	attribute_length;
	u2	sourcefile_index;
}
SouceFile 属性用于指出源文件名，attribute_length 的值必须为 2
*/

type SourceFileAttribute struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

func (self *SourceFileAttribute) readInfo(reader *ClassReader) {
	self.sourceFileIndex = reader.readUint16()
}

func (self *SourceFileAttribute) FileName() string {
	return self.cp.getUtf8(self.sourceFileIndex)
}
