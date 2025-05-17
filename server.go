package frpsp

import (
	"net/http"
	"encoding/json"
)

type ServeMux struct {
	http.ServeMux
}

var DefaultServeMux = &defaultServeMux
var defaultServeMux ServeMux

func NewServeMux() *ServeMux {
	return new(ServeMux)
}

func (mux *ServeMux) HandleLoginFunc(pattern string, handler func(*LoginObject)) {
	mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		var req LoginWrapper
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		o := &LoginObject{Req: &req.Content}
		handler(o)
		w.WriteHeader(http.StatusOK)

		var resp []byte
		switch o.Response() {
		case Deny:
			resp, _ = json.Marshal(map[string]any{
				"reject": true,
				"reject_reason": o.Reason(),
			})
		case AllowUnchanged:
			resp, _ = json.Marshal(map[string]any{
				"reject": false,
				"unchange": true,
			})
		case AllowChanged:
			resp, _ = json.Marshal(map[string]any{
				"unchange": false,
				"content": req.Content,
			})
		}
		w.Write(resp)
	})
}

func (mux *ServeMux) HandleNewProxyFunc(pattern string, handler func(*NewProxyObject)) {
	mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		var req NewProxyWrapper
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		o := &NewProxyObject{Req: &req.Content}
		handler(o)
		w.WriteHeader(http.StatusOK)

		var resp []byte
		switch o.Response() {
		case Deny:
			resp, _ = json.Marshal(map[string]any{
				"reject": true,
				"reject_reason": o.Reason(),
			})
		case AllowUnchanged:
			resp, _ = json.Marshal(map[string]any{
				"reject": false,
				"unchange": true,
			})
		case AllowChanged:
			resp, _ = json.Marshal(map[string]any{
				"unchange": false,
				"content": req.Content,
			})
		}
		w.Write(resp)
	})
}

func (mux *ServeMux) HandleCloseProxyFunc(pattern string, handler func(*CloseProxyObject)) {
	mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		var req CloseProxyWrapper
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		o := &CloseProxyObject{Req: &req.Content}
		handler(o)
		w.WriteHeader(http.StatusOK)

		var resp []byte
		switch o.Response() {
		case Deny:
			resp, _ = json.Marshal(map[string]any{
				"reject": true,
				"reject_reason": o.Reason(),
			})
		case AllowUnchanged:
			resp, _ = json.Marshal(map[string]any{
				"reject": false,
				"unchange": true,
			})
		case AllowChanged:
			resp, _ = json.Marshal(map[string]any{
				"unchange": false,
				"content": req.Content,
			})
		}
		w.Write(resp)
	})
}

func (mux *ServeMux) HandlePingFunc(pattern string, handler func(*PingObject)) {
	mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		var req PingWrapper
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		o := &PingObject{Req: &req.Content}
		handler(o)
		w.WriteHeader(http.StatusOK)

		var resp []byte
		switch o.Response() {
		case Deny:
			resp, _ = json.Marshal(map[string]any{
				"reject": true,
				"reject_reason": o.Reason(),
			})
		case AllowUnchanged:
			resp, _ = json.Marshal(map[string]any{
				"reject": false,
				"unchange": true,
			})
		case AllowChanged:
			resp, _ = json.Marshal(map[string]any{
				"unchange": false,
				"content": req.Content,
			})
		}
		w.Write(resp)
	})
}

func (mux *ServeMux) HandleNewWorkConnFunc(pattern string, handler func(*NewWorkConnObject)) {
	mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		var req NewWorkConnWrapper
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		o := &NewWorkConnObject{Req: &req.Content}
		handler(o)
		w.WriteHeader(http.StatusOK)

		var resp []byte
		switch o.Response() {
		case Deny:
			resp, _ = json.Marshal(map[string]any{
				"reject": true,
				"reject_reason": o.Reason(),
			})
		case AllowUnchanged:
			resp, _ = json.Marshal(map[string]any{
				"reject": false,
				"unchange": true,
			})
		case AllowChanged:
			resp, _ = json.Marshal(map[string]any{
				"unchange": false,
				"content": req.Content,
			})
		}
		w.Write(resp)
	})
}

func (mux *ServeMux) HandleNewUserConnFunc(pattern string, handler func(*NewUserConnObject)) {
	mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		var req NewUserConnWrapper
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		o := &NewUserConnObject{Req: &req.Content}
		handler(o)
		w.WriteHeader(http.StatusOK)

		var resp []byte
		switch o.Response() {
		case Deny:
			resp, _ = json.Marshal(map[string]any{
				"reject": true,
				"reject_reason": o.Reason(),
			})
		case AllowUnchanged:
			resp, _ = json.Marshal(map[string]any{
				"reject": false,
				"unchange": true,
			})
		case AllowChanged:
			resp, _ = json.Marshal(map[string]any{
				"unchange": false,
				"content": req.Content,
			})
		}
		w.Write(resp)
	})
}

func HandleLoginFunc(pattern string, handler func(*LoginObject)) {
	defaultServeMux.HandleLoginFunc(pattern, handler)
}

func HandleNewProxyFunc(pattern string, handler func(*NewProxyObject)) {
	defaultServeMux.HandleNewProxyFunc(pattern, handler)
}

func HandleCloseProxyFunc(pattern string, handler func(*CloseProxyObject)) {
	defaultServeMux.HandleCloseProxyFunc(pattern, handler)
}

func HandlePingFunc(pattern string, handler func(*PingObject)) {
	defaultServeMux.HandlePingFunc(pattern, handler)
}

func HandleNewWorkConnFunc(pattern string, handler func(*NewWorkConnObject)) {
	defaultServeMux.HandleNewWorkConnFunc(pattern, handler)
}

func HandleNewUserConnFunc(pattern string, handler func(*NewUserConnObject)) {
	defaultServeMux.HandleNewUserConnFunc(pattern, handler)
}

func ListenAndServe(addr string, handler http.Handler) error {
	if handler == nil {
		handler = DefaultServeMux
	}
	return http.ListenAndServe(addr, handler)
}
