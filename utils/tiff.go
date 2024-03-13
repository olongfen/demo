package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

func ProcessTiff2MbtilesFromDir(dir string) error {

	dirs, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, v := range dirs {
		if !strings.HasSuffix(v.Name(), ".tif") {
			continue
		}
		// tiff转成拇指图
		name := strings.Split(v.Name(), ".tif")[0]
		fileDir := path.Join(dir, v.Name())
		fmt.Println(fileDir)
		mbtileFile := path.Join(dir, name+".mbtiles")
		cmdGdalTranslateJPEG := exec.Command("gdal_translate", "-of", "JPEG", "-co", "COMPRESS=JPEG", "-outsize",
			"512", "512", fileDir, path.Join(dir, name+".jpg"))
		fmt.Println(cmdGdalTranslateJPEG.String())
		if err = cmdGdalTranslateJPEG.Run(); err != nil {
			fmt.Println("转换jpg失败", err.Error())
			os.Remove(path.Join(dir, name+".jpg"))
		}

		// tiff转成mbtiles
		cmdGdalTranslateMbtile := exec.Command("gdal_translate", "-of", "MBTiles", fileDir, mbtileFile)
		fmt.Println(cmdGdalTranslateMbtile.String())
		if err = cmdGdalTranslateMbtile.Run(); err != nil {
			fmt.Println("转换mbtiles失败", err.Error())
			os.Remove(mbtileFile)
			continue
		}
		cmdGdalAddo := exec.Command("gdaladdo", "-r", "average", mbtileFile, "2", "4", "8", "16", "32", "64", "128", "256")
		if err = cmdGdalAddo.Run(); err != nil {
			fmt.Println("转换mbtiles1失败", err.Error())
		}
	}
	return nil
}
