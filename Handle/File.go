package Handle

import (
	"os"
	"strconv"
)

func  Save(dstID int16,fileByte []byte) error {
	//创建用户的目录
	dstIDstr:=strconv.Itoa(int(dstID))
	err := os.Mkdir("./FileDir/"+dstIDstr, os.ModeDir)
	if err != nil {
		//如果是因为目录已存在而导致的创建错误则无视
		if !os.IsExist(err) {
			return err
		}
	}

	//在指定用户目录创建出该文件的空壳
	f, err := os.OpenFile("./FileDir/"+dstIDstr+"/"+"testFile", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	defer f.Close()
	//复制用户上传的文件从内存到本地
	f.Write(fileByte)
	return nil
}
