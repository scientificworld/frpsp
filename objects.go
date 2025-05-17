package frpsp

type ResponseType int

const (
	Deny ResponseType = iota
	AllowChanged
	AllowUnchanged
)

type ObjectBase struct {
	response ResponseType
	reason string
}

func (o *ObjectBase) Resolve(change bool) {
	if change {
		o.response = AllowChanged
	} else {
		o.response = AllowUnchanged
	}
}

func (o *ObjectBase) Reject(reason string) {
	o.response = Deny
	o.reason = reason
}

func (o *ObjectBase) Response() ResponseType {
	return o.response
}

func (o *ObjectBase) Reason() string {
	return o.reason
}

type LoginObject struct {
	ObjectBase
	Req *Login
}

type NewProxyObject struct {
	ObjectBase
	Req *NewProxy
}

type CloseProxyObject struct {
	ObjectBase
	Req *CloseProxy
}

type PingObject struct {
	ObjectBase
	Req *Ping
}

type NewWorkConnObject struct {
	ObjectBase
	Req *NewWorkConn
}

type NewUserConnObject struct {
	ObjectBase
	Req *NewUserConn
}
