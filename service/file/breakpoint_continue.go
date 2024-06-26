package file

import (
	"Mall/global"
	"Mall/model/file"
	"errors"

	"gorm.io/gorm"
)

type FileUploadAndDownloadService struct {
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: FindOrCreateFile
//@description: 上传文件时检测当前文件属性，如果没有文件则创建，有则返回文件的当前切片
//@param: fileMd5 string, fileName string, chunkTotal int
//@return: err error, file model.ExaFile

func (e *FileUploadAndDownloadService) FindOrCreateFile(fileMd5 string, fileName string, chunkTotal int) (err error, retFile file.ExaFile) {
	var cfile =file.ExaFile{}
	cfile.FileMd5 = fileMd5
	cfile.FileName = fileName
	cfile.ChunkTotal = chunkTotal

	if errors.Is(global.GVA_DB.Where("file_md5 = ? AND is_finish = ?", fileMd5, true).First(&retFile).Error, gorm.ErrRecordNotFound) {
		err = global.GVA_DB.Where("file_md5 = ? AND file_name = ?", fileMd5, fileName).Preload("ExaFileChunk").FirstOrCreate(&retFile, cfile).Error
		return err, retFile
	}
	cfile.IsFinish = true
	cfile.FilePath = retFile.FilePath
	err = global.GVA_DB.Create(&cfile).Error
	return err, cfile
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateFileChunk
//@description: 创建文件切片记录
//@param: id uint, fileChunkPath string, fileChunkNumber int
//@return: error

func (e *FileUploadAndDownloadService) CreateFileChunk(id uint, fileChunkPath string, fileChunkNumber int) error {
	var chunk file.ExaFileChunk
	chunk.FileChunkPath = fileChunkPath
	chunk.ExaFileID = id
	chunk.FileChunkNumber = fileChunkNumber
	err := global.GVA_DB.Create(&chunk).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFileChunk
//@description: 删除文件切片记录
//@param: fileMd5 string, fileName string, filePath string
//@return: error

func (e *FileUploadAndDownloadService) DeleteFileChunk(fileMd5 string, fileName string, filePath string) error {
	var chunks []file.ExaFileChunk
	var file file.ExaFile
	err := global.GVA_DB.Where("file_md5 = ? AND file_name = ?", fileMd5, fileName).First(&file).Update("IsFinish", true).Update("file_path", filePath).Error
	if err != nil {
		return err
	}
	err = global.GVA_DB.Where("exa_file_id = ?", file.ID).Delete(&chunks).Unscoped().Error
	return err
}
