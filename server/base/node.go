package main

import godiscovery "github.com/fananchong/go-discovery"

type Node struct {
	godiscovery.Node
}

func NewNode() *Node {
	this := &Node{}
	this.Node.Init(this)
	return this
}

func (this *Node) OnNodeUpdate(nodeType int, id string, data []byte) {
	xlog.Infoln("OnNodeUpdate: nodeType =", nodeType, "id =", id, "data =", data)
}

func (this *Node) OnNodeJoin(nodeType int, id string, data []byte) {
	xlog.Infoln("OnNodeJoin: nodeType =", nodeType, "id =", id, "data =", data)
}

func (this *Node) OnNodeLeave(nodeType int, id string) {
	xlog.Infoln("OnNodeLeave: nodeType =", nodeType, "id =", id)
}

func (this *Node) GetPutData() string {
	return ""
}