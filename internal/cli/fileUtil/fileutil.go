package fileUtil

import (
	"io"
	"io/ioutil"
	"os"
)

type FileUtil interface {
	GetIoReader(string) (io.Reader,error)
	IsFileExist(string) bool
	ReadFile(string) (string,error)
	CreateDirIfNotExist(string) error
	CreateFile(string) error
	WriteFile(string,string) error
}

type fileUtil struct {

}

func NewFileUtil() FileUtil {
	return &fileUtil{}
}

func (f *fileUtil) GetIoReader(filePath string) (io.Reader,error) {
	ioReader, err := os.Open(filePath)
	if err != nil {
		return nil,err
	}
	return ioReader,nil
}

func (f *fileUtil) IsFileExist(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}

func (f *fileUtil) ReadFile(filePath string) (string,error) {
	content,err:=ioutil.ReadFile(filePath)
	if err!=nil {
		return "",err
	}
	return string(content),nil
}

func (f *fileUtil) CreateDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *fileUtil) CreateFile(filePath string) error {
	_,err:=os.Create(filePath)
	if err!=nil {
		return err
	}
	return nil
}

func (f *fileUtil) WriteFile(filePath string,content string) error {
	contentInByte:=[]byte(content)
	err:=ioutil.WriteFile(filePath,contentInByte,0777)
	if err!=nil {
		return err
	}
	return nil
}