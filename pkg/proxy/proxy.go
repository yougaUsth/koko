package proxy

import (
	"github.com/ibuler/ssh"

	"cocogo/pkg/logger"
	"cocogo/pkg/model"
	"cocogo/pkg/service"
)

type ProxyServer struct {
	Session    ssh.Session
	User       *model.User
	Asset      *model.Asset
	SystemUser *model.SystemUser
}

func (p *ProxyServer) getSystemUserAuthOrManualSet() {
	info := service.GetSystemUserAssetAuthInfo(p.SystemUser.Id, p.Asset.Id)
	if p.SystemUser.LoginMode == model.LoginModeManual ||
		(p.SystemUser.Password == "" && p.SystemUser.PrivateKey == "") {
		logger.Info("Get password fom user input")
	}
	p.SystemUser.Password = info.Password
	p.SystemUser.PrivateKey = info.PrivateKey
}

func (p *ProxyServer) checkProtocol() bool {
	return true
}

func (p *ProxyServer) getSystemUserUsernameIfNeed() {

}

func (p *ProxyServer) validatePermission() bool {
	return true
}

func (p *ProxyServer) getServerConn() {

}

func (p *ProxyServer) sendConnectingMsg() {

}

func (p *ProxyServer) Proxy() {
	if !p.checkProtocol() {
		return
	}
	conn := SSHConnection{
		Host:     "192.168.244.185",
		Port:     "22",
		User:     "root",
		Password: "redhat",
	}
	ptyReq, _, ok := p.Session.Pty()
	if !ok {
		logger.Error("Pty not ok")
		return
	}
	err := conn.Connect(ptyReq.Window.Height, ptyReq.Window.Width, ptyReq.Term)
	if err != nil {
		return
	}

	sw := Switch{
		userSession: p.Session,
		serverConn:  &conn,
		parser:      parser,
	}
	_ = sw.Bridge()
	_ = conn.Close()
}
