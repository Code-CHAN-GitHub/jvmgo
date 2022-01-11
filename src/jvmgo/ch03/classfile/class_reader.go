package classfile

import "encoding/binary"

/**
Java 虚拟机规范定义了 u1、u2 和 u4 三种数据类型来表示 1、2 和 4 字节无符号整数，
分别对应 Go 语言的 uint8、uint16 和 uint32 类型
*/

type ClassReader struct {
	data []byte
}

func (self *ClassReader) readUint8() uint8 {
	val := self.data[0]
	self.data = self.data[1:]
	return val
}

func (self *ClassReader) readUint16() uint16 {
	// uint16(b[1]) | uint16(b[0])<<8
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}

func (self *ClassReader) readUint32() uint32 {
	// uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}

func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

// 读取 uint16 表
func (self *ClassReader) readUint16s() []uint16 {
	/**
	相同的多条数据一般按表（table）的形式存储在 class 文件中。
	表由表头和表项组成，表头是 u2 或 u4 整数, 假设表头是 n，
	后面就紧跟这 n 个表项数据
	*/
	n := self.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = self.readUint16()
	}
	return s
}

// 读取指定数量的字节
func (self *ClassReader) readBytes(n uint32) []byte {
	bytes := self.data[:n]
	self.data = self.data[n:]
	return bytes
}
