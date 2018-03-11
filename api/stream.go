package api

import (
	"net/http"
	"github.com/s4kibs4mi/belle-stream/stream"
	"github.com/hybridgroup/mjpeg"
)

func startStream(w http.ResponseWriter, r *http.Request) {
	err := stream.StartStream()
	if err != nil {
		resp := response{
			Status: http.StatusInternalServerError,
			Error:  err,
		}
		resp.ServeJSON(w)
		return
	}
	if r.URL.Query().Get("save") == "true" {
		stream.StartVideoRecord()
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
