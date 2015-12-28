/*
*  作者：pingzilao 531169454@qq.com
*/

package app

import(
	"fmt"
	"time"
	"io/ioutil"
	"net/http"
	"strings"
	l4g "github.com/alecthomas/log4go"
	pb "code.google.com/p/goprotobuf/proto"
	simplejson "github.com/bitly/go-simplejson"
	pkgzmqmgr "gw.com.cn/dzhyun/cm.git/zmqmgr"
)

type MsgMgr struct{
	mAppMgr *AppManager
}

func (this* MsgMgr) Start(cmAddr string)error {
	if err := pkgzmqmgr.Init();err!=nil{
		return err
	}
	pkgzmqmgr.Reg(dzhyun.Command_HEART_BEAT,          &HbProc{this.mAppMgr})
	if err := pkgzmqmgr.Start(cmAddr);err!=nil{
		return err
	}
	return nil
}

func (this* MsgMgr) Stop() (err error) {
	err = pkgzmqmgr.Stop()
	return
}

/*****************************HbProc********************************/
type HbProc struct{
	mAppMgr *AppManager
}

func (this * HbProc)Proc(netData *pkgzmqmgr.RNetData) error{
	msg, ok := netData.Msg.(*dzhyun.HeartBeat)
	if !ok {
		return fmt.Errorf("RNetData(%d %v) Not HeartBeat", netData.Id, netData.Cmd)
	}
	this.heartBeat(msg, netData)
	return nil
}

func (this *HbProc) heartBeat(heartbt *dzhyun.HeartBeat, netData *pkgzmqmgr.RNetData) error {
	errCode := 0
	heartRet := &dzhyun.HeartBeatRet{
		ServiceId: pb.String(heartbt.GetServiceId()),
		ErrCode:   &errCode,
	}

	if err:= pkgzmqmgr.SendMsg(netData.Id, netData.Addr, dzhyun.Command_HEART_BEAT_RET, heartRet);err!=nil{
		l4g.Error("ZmqMgr SendMsg Err[%s]", err.Error())
	}

	return nil
}