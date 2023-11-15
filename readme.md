# ebpf demo

## 目录结构

```bash
├── Dockerfile  # 镜像文件
├── Makefile    # Makefile 
├── demo_bpfeb.go  # go generate 生成的文件 
├── demo_bpfeb.o
├── demo_bpfel.go
├── demo_bpfel.o
├── go.mod
├── go.sum
├── headers  # 内核头文件
├── kprobe.c  # ebpf hook 程序（内核态）
├── main.go  # 用户态，从 Map 中读取数据
└── readme.md
```

本程序演示了将一个 eBPF 程序附加到内核符号的过程。
这个 eBPF 程序将被附加到 sys_execve 内核函数的起始位置，并且每秒钟打印出它被调用的次数。

## 如何使用

1、生成 linux 内核头文件
```bash
cd headers
bash update.sh
```
headers目录下common.h文件是用于在使用 C 代码的示例中的紧凑版本的 `vmlinux.h`。

2、生成 `vmlinux.h` 文件
```bash
# 从linux内核5.2开始, 在编译内核时, 就会把内核的数据结构自动的镶嵌到 vmlinux 中, 还可以借助命令, 将这个定义导出到一个头文件中。
bpftool btf dump file /sys/kernel/btf/vmlinux format c > vmlinux.h
```

3、运行示例
```bash
# 方式1
make clean
root@instance-00qqerhq:~/ebpf/ebpf-learn# ./demo 
2023/11/15 20:02:02 Waiting for events..
2023/11/15 20:02:03 sys_execve called 0 times
2023/11/15 20:02:04 sys_execve called 1 times
2023/11/15 20:02:05 sys_execve called 1 times
2023/11/15 20:02:06 sys_execve called 1 times
2023/11/15 20:02:07 sys_execve called 1 times
2023/11/15 20:02:08 sys_execve called 1 times
2023/11/15 20:02:09 sys_execve called 1 times
2023/11/15 20:02:10 sys_execve called 3 times
2023/11/15 20:02:11 sys_execve called 3 times
2023/11/15 20:02:12 sys_execve called 3 times
2023/11/15 20:02:13 sys_execve called 3 times
2023/11/15 20:02:14 sys_execve called 3 times

# 方式2
root@instance-00qqerhq:~/ebpf/ebpf-learn# go run -exec sudo ./
2023/11/15 20:02:56 Waiting for events..
2023/11/15 20:02:57 sys_execve called 0 times
2023/11/15 20:02:58 sys_execve called 0 times
2023/11/15 20:02:59 sys_execve called 0 times
2023/11/15 20:03:00 sys_execve called 0 times
2023/11/15 20:03:01 sys_execve called 0 times
2023/11/15 20:03:02 sys_execve called 0 times
2023/11/15 20:03:03 sys_execve called 0 times
2023/11/15 20:03:04 sys_execve called 1 times
2023/11/15 20:03:05 sys_execve called 1 times
2023/11/15 20:03:06 sys_execve called 1 times
```
