

问：在协程多级调用，报panic 会按层级调用 defer吗？
答：panic 如果不在当前协程处理错误，则会直接调用`runtime.fatalpanic`结束进程，并用`runtime.printpanics`打印错误栈。
    所以，panic必须在当前协程捕获并恢复，不然则会造成进程退出。

参考1：https://zhuanlan.zhihu.com/p/346514343
参考2：https://mp.weixin.qq.com/s?__biz=MzkyNzI1NzM5NQ==&mid=2247484780&idx=1&sn=0468a1a4dc27c09732804798e5609def