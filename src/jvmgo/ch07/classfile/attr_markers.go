package classfile

/**
Deprecated_attribute {
	u2	attribute_name_index;
	u4	attribute_length;
}

Synthetic_attribute {
	u2	attribute_name_index;
	u4	attribute_length;
}

Deprecate 属性用于指出类、接口、字段或方法以及不建议使用
Synthetic 属性用来标记源文件中不存在、由编译器删除的类成员
由于不包含数据，所以两个属性的 attribute_length 值必须为 0
*/

type DeprecatedAttribute struct {
	MarkerAttribute
}

type SyntheticAttribute struct {
	MarkerAttribute
}

type MarkerAttribute struct {
}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}
