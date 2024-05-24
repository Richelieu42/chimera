package opencvKit

import (
	"fmt"
	"gocv.io/x/gocv"
	"image/color"
)

func Start() {
	// 步骤1：打开摄像头设备
	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		fmt.Println("打开摄像头设备失败:", err)
		return
	}
	defer webcam.Close()

	// 步骤2：加载人脸识别分类器
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	if !classifier.Load("haarcascade_frontalface_default.xml") {
		fmt.Println("加载分类器文件失败")
		return
	}

	// 步骤3：创建一个窗口用于显示图像
	window := gocv.NewWindow("Face Detection")
	defer window.Close()

	img := gocv.NewMat()
	defer img.Close()

	for {
		// 步骤4：从摄像头读取图像帧
		if ok := webcam.Read(&img); !ok || img.Empty() {
			fmt.Println("无法从摄像头读取图像帧")
			break
		}

		// 步骤5：将图像转换为灰度图像，因为人脸识别通常在灰度图像上进行
		gray := gocv.NewMat()
		defer gray.Close()

		gocv.CvtColor(img, &gray, gocv.ColorBGRToGray)

		// 步骤6：检测人脸
		rects := classifier.DetectMultiScale(gray)
		fmt.Printf("检测到 %d 个人脸\n", len(rects))

		// 步骤7：在图像上绘制人脸边界框
		for _, r := range rects {
			gocv.Rectangle(&img, r, color.RGBA{0, 255, 0, 0}, 2)
		}

		// 步骤8：显示图像
		window.IMShow(img)

		// 步骤9：等待用户按下ESC键退出
		if window.WaitKey(1) == 27 {
			break
		}
	}
}
