package file_srvc

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) Upload(file *multipart.FileHeader, folder string) (string, error) {
	if file == nil {
		return "", nil
	}

	filename := filepath.Base(file.Filename)

	if _, err := os.Stat("./media/" + folder); errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll("./media/"+folder, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	{
		splitString := strings.Split(filename, ".")
		extra := strconv.Itoa(int(time.Now().Unix()))
		if len(splitString) >= 2 {
			splitString[len(splitString)-2] = splitString[len(splitString)-2] + "-" + extra
		} else {
			splitString[0] += "-" + extra
		}
		filename = strings.Join(splitString, ".")
	}
	filename = strings.ReplaceAll(filename, " ", "-")

	dst := "./media/" + folder + "/" + filename

	src, err := file.Open()
	if err != nil {
		return "", err
	}

	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return "", err
	}

	defer out.Close()

	_, err = io.Copy(out, src)

	if err != nil {
		return "", err
	}

	return "/media/" + folder + "/" + filename, nil
}

func (s *Service) Service_() {

}
