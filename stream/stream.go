package stream

import (
	"gocv.io/x/gocv"
	"github.com/hybridgroup/mjpeg"
	"fmt"
)

var isStarted = false
var webCam *gocv.VideoCapture
var stream = mjpeg.NewStream()
var err error

func StartStream() error {
	if isStarted {
		return nil
	}
	webCam, err = gocv.VideoCaptureDevice(0)
	if err != nil {
		return err
	}
	go capture()
	isStarted = true
	return nil
}

func StartVideoRecord() {

}

func StopStream() error {
	if isStarted {
		isStarted = false
		return webCam.Close()
	}
	return nil
}

func StopVideoRecord() {

}

func capture() {
	img := gocv.NewMat()
	for {
		if !isStarted {
			break
		}
		if ok := webCam.Read(img); !ok {
			fmt.Printf("cannot read device\n")
			break
		}
		if img.Empty() {
			continue
		}

		buf, _ := gocv.IMEncode(".jpg", img)
		stream.UpdateJPEG(buf)
	}
}

func GetStream() *mjpeg.Stream {
	return stream
}
