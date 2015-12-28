package zmqmgr

import (
	"time"
	"testing"
	"fmt"
	"os"
	pkgzmq "github.com/pebbe/zmq4"
	cmutils "gw.com.cn/dzhyun/utils.git/cmutils"
	dzhyun "gw.com.cn/dzhyun/dzhyun.git"
	pb "code.google.com/p/goprotobuf/proto"
)

func TestStart(t *testing.T) {
	if err:= DefaultZmqMgr.Init();err!=nil{
		t.Errorf("Init err[%v]", err)
	}
	if err := DefaultZmqMgr.Start("tcp://0.0.0.0:10199");err!=nil{
		t.Errorf("start err[%v]", err)
	}
	if err := DefaultZmqMgr.Stop(); err!=nil{
		t.Errorf("stop err[%v]", err)
	}
}

type MyProc struct{
	i int
}

func (this * MyProc)Proc(netData *RNetData) error{
	fmt.Println("id addr cmd ", netData.Id, string(netData.Addr), netData.Cmd)
	heartBeat, _ := netData.Msg.(*dzhyun.HeartBeat)
	sid := heartBeat.GetServiceId()
	fmt.Println("MyProc Ok",sid)
	if err := SendMsg(netData.Id, netData.Addr, netData.Cmd, heartBeat); err !=nil{
		fmt.Println("Proc send Err")
	}
	return nil
}



func TestNet(t *testing.T){
	if err:= DefaultZmqMgr.Init();err!=nil{
		t.Errorf("Init err[%v]", err)
	}
	if err := DefaultZmqMgr.Start("tcp://0.0.0.0:10199");err!=nil{
		t.Errorf("start err[%v]", err)
	}
	myProc := &MyProc{}
	DefaultZmqMgr.Reg(dzhyun.Command_HEART_BEAT, myProc)
	DSoc, err := pkgzmq.NewSocket(pkgzmq.DEALER)
	if err != nil {
		t.Errorf("NewSocket err[%v]", err)
	}
	state := dzhyun.StateType_RUNNING
	heartBeat := &dzhyun.HeartBeat{
		ServiceId: pb.String("/root/test1111"),
		Pid:       pb.Int32(int32(os.Getpid())),
		Result:    &state,
	}

	data, err0 := cmutils.ProtoPack(heartBeat, uint32(dzhyun.Command_HEART_BEAT), 12)
	if err0 != nil {
		t.Errorf("send heartbeat failed, proto pack error:%s", err0.Error())
	}
	
	DSoc.Connect("tcp://127.0.0.1:10199")
	
	if _, err := DSoc.SendBytes(data, pkgzmq.DONTWAIT); err!=nil{
		t.Errorf("send err[%v]", err)
	}
	
	if data, err := DSoc.RecvBytes(0); err!=nil{
		t.Errorf("recv err[%v]", err)
	}else{
		uCmd, _, reqid, err := cmutils.ProtoUnpack(data)
		if err != nil {
			t.Errorf("ProtoUnpack err[%s]", err.Error())
		}
		fmt.Println("recv data:", uCmd, reqid)
	}
	
	time.Sleep(time.Second*2)
	
	if err := DefaultZmqMgr.Stop(); err!=nil{
		t.Errorf("stop err[%v]", err)
	}
}

func TestNet2(t *testing.T){
	if err:= DefaultZmqMgr.Init();err!=nil{
		t.Errorf("Init err[%v]", err)
	}
	if err := DefaultZmqMgr.Start("tcp://0.0.0.0:10199");err!=nil{
		t.Errorf("start err[%v]", err)
	}
	myProc := &MyProc{}
	DefaultZmqMgr.Reg(dzhyun.Command_HEART_BEAT, myProc)
	DSoc, err := pkgzmq.NewSocket(pkgzmq.DEALER)
	if err != nil {
		t.Errorf("NewSocket err[%v]", err)
	}
	state := dzhyun.StateType_RUNNING
	heartBeat := &dzhyun.HeartBeat{
		ServiceId: pb.String("/root/test1111"),
		Pid:       pb.Int32(int32(os.Getpid())),
		Result:    &state,
	}

	data, err0 := cmutils.ProtoPack(heartBeat, uint32(dzhyun.Command_HEART_BEAT), 12)
	if err0 != nil {
		t.Errorf("send heartbeat failed, proto pack error:%s", err0.Error())
	}
	DSoc.SetIdentity("IdAddr")
	
	DSoc.Connect("tcp://127.0.0.1:10199")
	
	if _, err := DSoc.SendBytes(data, pkgzmq.DONTWAIT); err!=nil{
		t.Errorf("send err[%v]", err)
	}
	
	if data, err := DSoc.RecvBytes(0); err!=nil{
		t.Errorf("recv err[%v]", err)
	}else{
		uCmd, _, reqid, err := cmutils.ProtoUnpack(data)
		if err != nil {
			t.Errorf("ProtoUnpack err[%s]", err.Error())
		}
		fmt.Println("recv data:", uCmd, reqid)
	}
	
	time.Sleep(time.Second*2)
	
	if err := DefaultZmqMgr.Stop(); err!=nil{
		t.Errorf("stop err[%v]", err)
	}
}

type ExcepProc struct{
	
}

func (this * ExcepProc)Proc(netData *RNetData) error{
	return fmt.Errorf("ExcepProc faild(%d %v)", netData.Id, netData.Cmd)
}

func TestNet3(t *testing.T){
	if err:= DefaultZmqMgr.Init();err!=nil{
		t.Errorf("Init err[%v]", err)
	}
	if err := DefaultZmqMgr.Start("tcp://0.0.0.0:10199");err!=nil{
		t.Errorf("start err[%v]", err)
	}
	myProc := &ExcepProc{}
	DefaultZmqMgr.Reg(dzhyun.Command_HEART_BEAT, myProc)
	DSoc, err := pkgzmq.NewSocket(pkgzmq.DEALER)
	if err != nil {
		t.Errorf("NewSocket err[%v]", err)
	}
	state := dzhyun.StateType_RUNNING
	heartBeat := &dzhyun.HeartBeat{
		ServiceId: pb.String("/root/test1111"),
		Pid:       pb.Int32(int32(os.Getpid())),
		Result:    &state,
	}

	data, err0 := cmutils.ProtoPack(heartBeat, uint32(dzhyun.Command_HEART_BEAT), 12)
	if err0 != nil {
		t.Errorf("send heartbeat failed, proto pack error:%s", err0.Error())
	}
	
	DSoc.Connect("tcp://127.0.0.1:10199")
	
	if _, err := DSoc.SendBytes(data, pkgzmq.DONTWAIT); err!=nil{
		t.Errorf("send err[%v]", err)
	}
	
//	if data, err := DSoc.RecvBytes(0); err!=nil{
//		t.Errorf("recv err[%v]", err)
//	}else{
//		uCmd, _, reqid, err := cmutils.ProtoUnpack(data)
//		if err != nil {
//			t.Errorf("ProtoUnpack err[%s]", err.Error())
//		}
//		fmt.Println("recv data:", uCmd, reqid)
//	}
	
	time.Sleep(time.Second*2)
	
	if err := DefaultZmqMgr.Stop(); err!=nil{
		t.Errorf("stop err[%v]", err)
	}
}

func TestStart2(t *testing.T){
	if err:= DefaultZmqMgr.Init();err!=nil{
		t.Errorf("Init err[%v]", err)
	}
	if err := DefaultZmqMgr.Start("12873uyd");err==nil{
		t.Errorf("start need err but nil")
	}
	if err := DefaultZmqMgr.Stop(); err!=nil{
		t.Errorf("stop err[%v]", err)
	}
}
