package fileutil

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gabriel-vasile/mimetype"
)

func LineCounter(filelocation string) (int, error) {
	file, err := os.Open(filelocation)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	lineCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return lineCount, nil
}

func IsFileExist(filelocation string) (bool, error) {
	fi, err := os.Stat(filelocation)
	if err != nil {
		return false, err
	}
	mode := fi.Mode()
	return mode.IsRegular(), nil
}

func IsDirExist(dirlocation string) (bool, error) {
	fi, err := os.Stat(dirlocation)
	if err != nil {
		return false, err
	}
	mode := fi.Mode()
	return mode.IsDir(), nil
}

func GetFileContentType(out *os.File) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

func CreateDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func IsTextFile(filelocation string) (bool, error) {
	mime, err := mimetype.DetectFile(filelocation)
	if err != nil {
		return false, err
	}
	return strings.Contains(mime.String(), "text/plain"), nil
}

func IsImageFile(filelocation string) (bool, error) {
	mime, err := mimetype.DetectFile(filelocation)
	if err != nil {
		return false, err
	}
	return strings.Contains(mime.String(), "image/"), nil
}

func ReadLine(filelocation string, lineNum int) (line string, lastLine int, err error) {
	if lineNum < 1 {
		return "", 0, fmt.Errorf("read line invalid request: line %d", lineNum)
	}

	file, err := os.Open(filelocation)
	if err != nil {
		return "", 0, err
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		lastLine++
		if lastLine == lineNum {
			// you can return sc.Bytes() if you need output in []bytes
			return sc.Text(), lastLine, sc.Err()
		}
	}
	return line, lastLine, io.EOF
}

func GetFileSize(filelocation string) (int64, error) {
	fi, err := os.Stat(filelocation)
	if err != nil {
		return 0, err
	}
	// get the size
	return fi.Size(), nil
}

func AppendCSV(filelocation string, csvdata []string, addnewline bool) error {
	f, err := os.OpenFile(filelocation, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if addnewline {
		f.WriteString("\n")
	}

	w := csv.NewWriter(f)
	err = w.Write(csvdata)
	if err != nil {
		return err
	}
	w.Flush()
	return nil
}

func ReadLines(filelocation string) ([]string, error) {
	file, err := os.Open(filelocation)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func WriteLines(lines []string, filelocation string) error {
	file, err := os.Create(filelocation)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func MoveFile(currentLocation string, newLocation string) error {
	isExist, _ := IsFileExist(newLocation)
	if isExist {
		return fmt.Errorf("%s is exist", newLocation)
	}

	err := os.Rename(currentLocation, newLocation)
	if err != nil {
		return err
	}

	return nil

}

func CopyFile(source string, destination string) error {
	isExist, _ := IsFileExist(destination)
	if isExist {
		return fmt.Errorf("%s is exist", destination)
	}

	from, err := os.Open(source)
	if err != nil {
		return err
	}
	defer from.Close()

	to, err := os.OpenFile(destination, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		return err
	}
	return nil
}

func IsNewLineAtEOF(filelocation string) (bool, error) {
	file, err := os.Open(filelocation)
	if err != nil {
		return false, err
	}
	defer file.Close()
	buf := make([]byte, 10)
	lastBuf := make([]byte, 10)
	lastByteRead := 0
	for {
		byteread, err := file.Read(buf)

		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}

			break
		}
		lastBuf = buf
		lastByteRead = byteread
	}
	lastReadString := string(lastBuf[:lastByteRead])
	lastString := string(lastReadString[len(lastReadString)-1])

	return lastString == "\n", nil
}

func DeleteFile(filelocation string) error {
	ok, err := IsFileExist(filelocation)
	if err == nil {
		if !ok {
			return fmt.Errorf("%s not file", filelocation)
		}
		err = os.Remove(filelocation)
		if err != nil {
			return err
		}
		return nil
	}
	return err

}

func DeleteFolder(dirLocation string) error {
	ok, err := IsDirExist(dirLocation)
	if err == nil {
		if !ok {
			return fmt.Errorf("%s not directory", dirLocation)
		}
		err = os.Remove(dirLocation)
		if err != nil {
			return err
		}
		return nil

	}
	return err
}

func ExtractFileName(filelocation string) string {
	filenamewithext := filepath.Base(filelocation)
	extension := filepath.Ext(filenamewithext)
	return filenamewithext[0 : len(filenamewithext)-len(extension)]
}

func ExtractExtension(filelocation string) string {
	filenamewithext := filepath.Base(filelocation)
	extension := filepath.Ext(filenamewithext)
	return extension
}

func ListFileInDirectory(rootdirectory string, isfindsubdir bool) ([]string, error) {
	lastDirString := rootdirectory[len(rootdirectory)-1:]
	if lastDirString != string(os.PathSeparator) {
		rootdirectory = rootdirectory + string(os.PathSeparator)
	}

	var files []string
	err := filepath.Walk(rootdirectory, visit(&files, rootdirectory, isfindsubdir))
	if err != nil {
		return []string{}, err
	}
	return files, nil
}

func visit(files *[]string, rootdirectory string, isfindsubdir bool) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			if filepath.Dir(rootdirectory) == filepath.Dir(path) && !isfindsubdir {
				if string(filepath.Base(path)[0]) != "." {
					*files = append(*files, path)
				}
			}

			if isfindsubdir {
				if string(filepath.Base(path)[0]) != "." {
					*files = append(*files, path)
				}
			}

		}
		return nil
	}
}

func CheckValidTypeFile(bytefile []byte, fileTypeCheck []string) bool {
	correctTypeFile := false

	if len(fileTypeCheck) > 0 {
		fileType := mimetype.Detect(bytefile)
		fileTypeString := fileType.String()
		fileTypeString = strings.ToLower(fileTypeString)
		for _, item := range fileTypeCheck {
			fileTypeCheck := strings.ToLower(item)
			if strings.Contains(fileTypeString, fileTypeCheck) {
				correctTypeFile = true
				break
			}
		}
	}

	return correctTypeFile
}
