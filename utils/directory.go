package utils

import (
	"Mall/global"
	"os"

	"go.uber.org/zap"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			global.GVA_LOG.Debug("create directory" + v)
			if err := os.MkdirAll(v, os.ModePerm); err != nil {
				global.GVA_LOG.Error("create directory"+v, zap.Any(" error:", err))
				return err
			}
		}
	}
	return err
}
