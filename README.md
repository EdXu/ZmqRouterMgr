# ZmqRouterMgr
封装的zeromq的Router管理模块，支持回调和正常处理

文件结构：
        |-----zmqmgr                    包名
        |       |----zmqmgr.go          zmq管理模块
		|		|----zmqmgr_test.go     单元测试
		|
		|------app                      包名
		|       |----msgmgr.go          具体zmq消息处理
	    |       |----msgmgr_test.go
		

网络数据协议采用protobuf，使用者需自行修改。

zmqmgr使用：
      初始化：Init()；Start(cmAddr string)；
	  停止：Stop()；
	  协议处理：1.正常处理：UseRevChan()接口返回chan，推送网络消息。
	            2.回调模式：Reg(cmd dzhyun.Command, netDataProc INetDataProc)注册命令字，及其处理接口。
				
