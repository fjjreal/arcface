package routers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "hr-face-free/libs/hr/engine"
    util "hr-face-free/utils"
    "hr-face-free/conf"
    "encoding/base64"
    "fmt"
    "crypto/md5"
)

func helloHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "Hello World!",
    })
}

func activeArcSoft(c *gin.Context){

	arcConf := conf.LoadArcConfig()

	appID   := arcConf.AppId
	appKEY  := arcConf.AppKey

	message := "激活成功"
	status  := 1
	if err := engine.Activation(appID, appKEY); err != nil {
		status  = 0
		message = err.Error()
	}
	c.JSON(http.StatusOK, gin.H{
		"status" : status,
        "message": message,
        "appID": appID,
    })
}

func getFeature(c *gin.Context){
	path := "/data/src/public/test/a.jpeg"

	status, message := feature(path)
	c.JSON(http.StatusOK, gin.H{
		"status" : status,
        "message": message,
    })
}
func feature(imagePath string) (int, string) {
	// 初始化引擎
	engineObj, err := engine.NewFaceEngine(engine.DetectModeImage,
		engine.OrientPriority0,
		12,
		50,
		engine.EnableFaceDetect|engine.EnableFaceRecognition)
	if err != nil {
		return 0, err.Error()
	}
	// 加载图片
	width, height := util.GetImageWidthAndHeight(imagePath)
	imageData := util.GetResizedBGR(imagePath)
	// 检测人脸
	info, err := engineObj.DetectFaces(width-width%4, height, engine.ColorFormatBGR24, imageData)
	if err != nil {
		return 0, err.Error()
	}
	// 获取单人脸信息
	singleFaceInfoArr := engine.GetSingleFaceInfo(info)
	if(len(singleFaceInfoArr) == 0){
		return 0, "get singleFaceInfo defeat, sdk error"
	}
	// 获取人脸特征码
	f1, err := engineObj.FaceFeatureExtract(width-width%4, height, engine.ColorFormatBGR24, imageData, singleFaceInfoArr[0])
	if err != nil {
		return 0, err.Error()
	}
	feature := f1.Feature
	f1.Release()
	// 销毁引擎
	if err = engineObj.Destroy(); err != nil {
		return 0, err.Error()
	}
	return 1, base64.StdEncoding.EncodeToString(feature)
}

func doCompare(c *gin.Context){

	//需要比对的图片
	name := c.DefaultQuery("name", "b")
	path := "/data/src/public/test/" + name + ".jpeg"

	// 已有的特征值
	// /public/test/a.jpeg
	exitedFearureBase64 := "AID6RAAAoEFpV9c9ojZcve7nkTvPxeO7Yc1RPdi+Q73Ao5C9bRy/PVEm2jyP6py9PQSHva0vMbw9N6S9UmffuHH5tz1Bn5y9SAYxvGtkEb0QuFE9OqdkPZdmoL0Njqk9xV8EPg7B0rx1LYQ70vTdvNLM+D0HT5q9sI6JPU5eozy2oCM+jla1vZUGp7zJPnQ8XqWVvJcYSrsuXdG9QR7bPPn/Gj0JHhs9s2lPO4Kr2j3ucbm9kndIvQgGVz0Yjp09ZcAkvRaV5T377Z87UIPuvGDp0rws/4O8YY8SPfiOMrslmIY9ykOwPOmIJj22BIm9OvsXPplj9zrCmdE8zffCvVcayryj2za9xy+NPXgDRTy2syi86ftlPZFWPr2hphC7UVDGvJbDrj1qQ/k8cSANPduDjT338qs83VoIvdsSqLyn+Ny9YoIYPNKj+j3P3Co8h5jnvGfMB7zqNAe96nvmPUZGJzxnWJo94HhbvRZkurzFDWu6ffbgvR7lxL0+wq25JggMvbmBwTth5UK76Xf6u+EpDr5RPEG9L+8xPUGLabxAHBW9s5L3Pa2r2j1Lo1A9VoAcvFWTk7yI9608cB/fPAOqt7wn/c887s9vPc6xNr15zK89HwtvPTH65zxqmnu7edxovCWpEj3VsLC7hTQvvXC4mr2Pwse97dAwPQNeOb2p3wu+eE/Zver0vD2LCuk77bCCPdIYvjv3DxE+IgFFvJbbFz3IKfO83q85PWkMMb4I2Im9IqoDPSbv8z2bB9c9WxapPHlt5rxdjxw9WPmdPBm/IL2tvWU8T9srvG1JzL0x+k48qkYxvYEgd72sGIs9qXmHvaPbprsw6RO9rhIrvZsWSb2dVsM767VSPTq3Fj3DnJQ9nwzsvdjVGD7tcdY8XdGlPHM8eT2AMEs9f+i7vAHt+j0RvH28+BPFPUlmqj2y/pQ6b6bvvd2567zq+g+9lxY4PeNEzrzQvYw9xVuDPWltIb23kvI7pj4wPW7oyz0pq608NvPRvSDT27tXCw4+34SePajZUT0j5T29Lk4vPX4gC70KuI89WJFYvJLakb1zVi++FhmhPLiKmrvmzfA95SbNPGm0Dj0mDFA9Noyou4OOx70g9b47dwhPvU5LKT1yhAO9dqoKPhLuWT2+8R696Ed3vc1khb3Tw447amZwvXlPcDw4nXC8D+C6O9Or0rzvmv+6vxsbPh6Y5T1oKJk7tZ7yPVcNEL13XLe8lxSSvU3vY73/ZCM7NdB/PR8jD73LkJ694LNBPWqNvrzQxem8L2OKvd22cLyWGou8soVnPYNzDb3TmMO6GGgxvngvoTuBi2a9isSZPe/LE72ZoeO8qOSqPQ6w6D3pKzc9N2pWPMaeyj0gECE9"

	fts := map[string]string{"/public/test/a.jpeg":exitedFearureBase64}

	status, message, confidenceLevel := compare(path, fts, 0.8)
	target := ""
	if status == 1 {
		target = message
	}

	c.JSON(http.StatusOK, gin.H{
		"status" : status,
        "message": message,
        "target": target,
        "confidenceLevel": confidenceLevel, //相似值
    })
}

//图片路径、已有的特征值、比对阈值
func compare(imagePath string, features map[string]string, compareTargetNum float32) (int, string, float32) {
	
	var confidenceLevel float32 = 0
	var status int = 0

	// 初始化引擎
	engineObj, err := engine.NewFaceEngine(engine.DetectModeImage,
		engine.OrientPriority0,
		12,
		50,
		engine.EnableFaceDetect|engine.EnableFaceRecognition)
	if err != nil {
		return status, err.Error(), confidenceLevel
	}
	defer engineObj.Destroy()

	status, f1Base64 := feature(imagePath)
	if status != 1 {
		return status, f1Base64, confidenceLevel
	}
	str, _ := base64.StdEncoding.DecodeString(f1Base64)
	f1 := engine.ConvertToFaceFeaturePro([]byte(str))

	var f2 engine.FaceFeature
	var target string
	var end int = 0
	var max float32 = 0
	for k, v := range features {
		fmt.Println("person_id:", k,", face_feature_md5:", toMd5(v))
		str, _ = base64.StdEncoding.DecodeString(v)
		f2 = engine.ConvertToFaceFeaturePro([]byte(str))
		confidenceLevel, _ = engineObj.FaceFeatureComparePro(f1, f2, 1)
		if max < confidenceLevel {
			max = confidenceLevel
		}
		if (confidenceLevel <= 1 && confidenceLevel > compareTargetNum) {
			target = k
			status = 1
			break;
		}
		status = 2
		end++
	}
	fmt.Println("匹配图片次数：", end)
	fmt.Println("最终匹配值：", max)

	f1.Release()
	f2.Release()

	return status, target, confidenceLevel
}


// SetupRouter 配置路由信息
func SetupRouter() *gin.Engine {
    r := gin.Default()
    r.GET("/", helloHandler)
    r.GET("/active", activeArcSoft)
    r.GET("/test", getFeature)
    r.GET("/compare", doCompare)
    return r
}

func toMd5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}