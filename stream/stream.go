package stream

import (
	"gocv.io/x/gocv"
	"github.com/hybridgroup/mjpeg"
	"fmt"
)

var isStarted = false
var shouldSave = false
var webCam *gocv.VideoCapture
var vidWriter *gocv.VideoWriter
var stream = mjpeg.NewStream()
var err error

func catchIfFall() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}

func StartStream(shouldSave bool) error {
	if isStarted {
		return nil
	}
	webCam, err = gocv.VideoCaptureDevice(0)
	if err != nil {
		return err
	}
	if shouldSave {
		shouldSave = true
		err := StartVideoRecord()
		if err != nil {
			webCam.Close()
			return err
		}
	}
	go capture()
	isStarted = true
	return nil
}

func StartVideoRecord() error {
	vidWriter, err = gocv.VideoWriterFile("/Users/sakib/record.avi", "MJPG", 10, 1024, 720)
	return err
}

func StopStream() error {
	if isStarted {
		isStarted = false
		StopVideoRecord()
		return webCam.Close()
	}
	return nil
}

func StopVideoRecord() {
	if shouldSave {
		err := vidWriter.Close()
		if err != nil {
			fmt.Println(err)
		}
		shouldSave = false
	}
}

func capture() {
	defer catchIfFall()

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
		if shouldSave {
			err := vidWriter.Write(img)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func GetStream() *mjpeg.Stream {
	return stream
}
