# jvmgo
参看《自己动手写Java虚拟机》，利用 Go 实现 Java 虚拟机

**目录结构**

<img src="https://img-bed-1307643120.cos.ap-shanghai.myqcloud.com/img/202201112022338.png" style="zoom:50%;" />

- **classfile**：解析 class 字节码文件
- **classpath**：从磁盘加载 class 文件
- **instructions**：指令集
- **native**：实现必要的本地方法
- **rtda**：运行时数据区
- **cmd**：虚拟机启动参数
- **interpreter**：解释器
- **jvm**：为虚拟机启动进行准备
- **main**：虚拟机入口

