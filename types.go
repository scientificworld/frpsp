package frpsp

const (
	TCP = "tcp"
	UDP = "udp"
	HTTP = "http"
	HTTPS = "https"
	TCPMUX = "tcpmux"
	STCP = "stcp"
	SUDP = "sudp"
	XTCP = "xtcp"
)

type UserInfo struct {
	User string `json:"user"`
	Metas map[string]string `json:"metas"`
	RunID string `json:"run_id"`
}

type Login struct {
	Version string `json:"version"`
	Hostname string `json:"hostname"`
	OS string `json:"os"`
	Arch string `json:"arch"`
	User string `json:"user"`
	Timestamp int64 `json:"timestamp"`
	PrivilegeKey string `json:"privilege_key"`
	RunID string `json:"run_id"`
	PoolCount int `json:"pool_count"`
	Metas map[string]string `json:"metas"`
	ClientAddress string `json:"client_address"`
}

type NewProxy struct {
	User UserInfo `json:"user"`
	ProxyName string `json:"proxy_name"`
	ProxyType string `json:"proxy_type"`
	UseEncryption bool `json:"use_encryption"`
	UseCompression bool `json:"use_compression"`
	BandwidthLimit string `json:"bandwidth_limit"`
	BandwidthLimitMode string `json:"bandwidth_limit_mode"`
	Group string `json:"group"`
	GroupKey string `json:"group_key"`
	RemotePort int `json:"remote_port"`
	CustomDomains []string `json:"custom_domains"`
	Subdomain string `json:"subdomain"`
	Locations []string `json:"locations"`
	HTTPUser string `json:"http_user"`
	HTTPPwd string `json:"http_pwd"`
	HTTPHeaderRewrite string `json:"http_header_rewrite"`
	Headers map[string]string `json:"headers"`
	SK string `json:"sk"`
	Multiplexer string `json:"multiplexer"`
	Metas map[string]string `json:"metas"`
}

type CloseProxy struct {
	User UserInfo `json:"user"`
	ProxyName string `json:"proxy_name"`
}

type Ping struct {
	User UserInfo `json:"user"`
	Timestamp int64 `json:"timestamp"`
	PrivilegeKey string `json:"privilege_key"`
}

type NewWorkConn struct {
	User UserInfo `json:"user"`
	RunID string `json:"run_id"`
	Timestamp int64 `json:"timestamp"`
	PrivilegeKey string `json:"privilege_key"`
}

type NewUserConn struct {
	User UserInfo `json:"user"`
	ProxyName string `json:"proxy_name"`
	ProxyType string `json:"proxy_type"`
	RemoteAddr string `json:"remote_addr"`
}

type LoginWrapper struct {
	Content Login `json:"content"`
}

type NewProxyWrapper struct {
	Content NewProxy `json:"content"`
}

type CloseProxyWrapper struct {
	Content CloseProxy `json:"content"`
}

type PingWrapper struct {
	Content Ping `json:"content"`
}

type NewWorkConnWrapper struct {
	Content NewWorkConn `json:"content"`
}

type NewUserConnWrapper struct {
	Content NewUserConn `json:"content"`
}
