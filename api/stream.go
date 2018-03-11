package api

import (
	"net/http"
	"github.com/s4kibs4mi/belle-stream/stream"
	"github.com/hybridgroup/mjpeg"
	"fmt"
)

func recovery(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {
				resp := response{
					Status: http.StatusInternalServerError,
					Error:  fmt.Errorf("%v", rvr),
				}
				resp.ServeJSON(w)
				return
			}
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func startStream(w http.ResponseWriter, r *http.Request) {
	err := stream.StartStream(r.URL.Query().Get("save") == "true")
	if err != nil {
		resp := response{
			Status: http.StatusInternalServerError,
			Error:  err,
		}
		resp.ServeJSON(w)
		return
	}
	resp := response{
		Status:  http.StatusOK,
		Message: "Streaming server has been started",
	}
	resp.ServeJSON(w)
}

func stopStream(w http.ResponseWriter, r *http.Request) {
	err := stream.StopStream()
	if err != nil {
		resp := response{
			Status: http.StatusInternalServerError,
			Error:  err,
		}
		resp.ServeJSON(w)
		return
	}
	stream.StopVideoRecord()
	resp := response{
		Status:  http.StatusOK,
		Message: "Streaming server has been stopped",
	}
	resp.ServeJSON(w)
}

func serveStream() *mjpeg.Stream {
	return stream.GetStream()
}
