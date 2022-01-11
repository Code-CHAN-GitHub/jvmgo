package classpath

import (
	"os"
	"strings"
)

/**
路径列表分隔符 ":"
java -cp path/to/classes:lib/a.jar:lib/b.jar ...
*/
const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	// 加载 class 文件
	//  @param className class 文件的相对路径，例如 java/lang/Object.class
	// 	@return []byte 读取到的字节数据
	readClass(className string) ([]byte, Entry, error)
	// 类似于 Java 的 toString 方法
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildCardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") ||
		strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") ||
		strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}
