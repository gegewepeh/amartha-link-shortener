package network

import (
	"fmt"
	"net/http"
	"time"

	utils "link-shortener/internal/pkg/utils"
)

// HTTPHandler is generic handler to all routes
type HTTPHandler func(w http.ResponseWriter, r *http.Request) (*Response, *utils.AppError)

// ServeHTTP is http.Handler interface implementation
func (fn HTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	reqID := r.Header.Get("X-Request-Id")
	if reqID == "" {
		reqID = utils.RandString(8, true)
	}
	ctx = utils.SetRequestID(ctx, reqID)

	utils.Log(ctx, fmt.Sprintf("START %s %s %s", r.Method, r.URL.Path, getIP(r)))

	start := time.Now()
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")

	r = r.WithContext(ctx)
	resp, err := fn(w, r)

	elapsed := time.Since(start)
	if err != nil {
		utils.Log(ctx, "%v", err)

		w.WriteHeader(err.Code)
		w.Write(err.ToJSON())

		utils.Log(ctx, "END %s %s %s %d %s", r.Method, r.URL.Path, getIP(r), err.Code, elapsed.String())
		return
	}

	if resp.Code == 0 {
		resp.Code = http.StatusOK
	}

	data, err := resp.ToJSON()

	if err != nil {
		utils.Log(ctx, "%v", err)

		w.WriteHeader(err.Code)
		w.Write(err.ToJSON())

		utils.Log(ctx, "END %s %s %s %d %s", r.Method, r.URL.Path, getIP(r), err.Code, elapsed.String())
		return
	}

	w.WriteHeader(resp.Code)
	w.Write(data)

	utils.Log(ctx, "END %s %s %s %d %s", r.Method, r.URL.Path, getIP(r), resp.Code, elapsed.String())
	return
}

func getIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}
