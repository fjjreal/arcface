package conf

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type ArcConf struct {
    AppId   string `json:"ARC_APPID"`
    AppKey  string `json:"ARC_APPKEY"`
}

func LoadArcConfig () *ArcConf {

	file, err := os.Open("conf/arc.json")
	if err != nil {
		path, err := os.Executable()
		if err != nil {
		    panic(err)
		}
		dir := filepath.Dir(path)
		file, err = os.Open(dir + "/conf/arc.json")
		if err != nil {
		    panic(err)
		}
	}

	dcobj := ArcConf{}
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fd, err := ioutil.ReadAll(file)
	conf := string(fd)
	//json解析到结构体里面
	err = json.Unmarshal([]byte(conf), &dcobj)
	if err != nil {
		panic(err)
	}
	return &dcobj
}

