# gocap
基于gopacket的抓包工具

## win10环境搭建步骤
#### 1、下载安装tdm-gcc [下载地址](http://tdm-gcc.tdragon.net/download)  
#### 2、tmd-gcc的bin目录加入环境变量(C:\TDM-GCC-64\bin)  
#### 3、下载安装winpcap [下载地址](https://www.winpcap.org/install/default.htm)  
#### 4、下载安装winpcap开发包 [下载地址](https://www.winpcap.org/devel.htm)  
```
开发包解压到C盘根目录,否则需要改gopacket代码
```
#### 5、由于winpcap开发包C:\WpdPack\Lib\x64目录下缺乏libpacket.a/libwpcap.a两个静态库，需要手动生成
```
(生成比较麻烦,已将静态库附带上传到github)
5.1 在c:\windows\system32目录下找到wpcap.dll和packet.dll,并copy到一个临时目录
5.2 linux执行:(gendef命令需安装MinGW,本人是装了WSL直接apt-get)
    gendef wpcap.dll
    gendef packet.dll
5.3 windows下执行:
    dlltool --as-flags=--64 -m i386:x86-64 -k --output-lib libwpcap.a --input-def wpcap.def
    dlltool --as-flags=--64 -m i386:x86-64 -k --output-lib libpacket.a --input-def packet.def
5.4 copy两个文件到C:\WpdPack\Lib\x64
```

#### 6、go get github.com/google/gopacket
