package fsystem

import (
	"errors"
	"github.com/spf13/afero"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var (
	Fs   = afero.NewOsFs()
	Mem  = afero.NewMemMapFs()
	FsRo = afero.NewReadOnlyFs(Fs)
)

func CreateFile(pathToFile string) (afero.File, error) {
	if err := Fs.MkdirAll(filepath.Dir(pathToFile), os.ModePerm); err != nil {
		return nil, err
	}
	return Fs.Create(pathToFile)
}

func FileExists(pathToFile string) (fs bool, mem bool, err error) {
	fs, err = afero.Exists(Fs, pathToFile)
	mem, err = afero.Exists(Mem, pathToFile)
	return fs, mem, err
}

func ReadAll(reader io.Reader) ([]byte, error) {
	return afero.ReadAll(reader)
}

func ReadRoFile(pathToFile string) ([]byte, error) {
	return afero.ReadFile(FsRo, pathToFile)
}

func ReadFsFile(pathToFile string) ([]byte, error) {
	return afero.ReadFile(Fs, pathToFile)
}

func ReadMemFile(pathToFile string) ([]byte, error) {
	return afero.ReadFile(Mem, pathToFile)
}

func ReadFile(Fs afero.Fs, pathToFile string) ([]byte, error) {
	return afero.ReadFile(Fs, pathToFile)
}

func ReadProperties(pathToFile string) (out map[string]string, err error) {
	fout, err := ReadFsFile(pathToFile)
	out = make(map[string]string)
	for _, line := range strings.Split(string(fout), "\n") {
		if !strings.Contains(line, "=") {
			continue
		}

		keyAndValue := strings.Split(line, "=")
		key := keyAndValue[0]

		if len(keyAndValue) > 2 {
			return nil, errors.New("malformed property: " + key)
		}

		var value string
		if len(keyAndValue) == 2 {
			value = keyAndValue[1]
		}
		out[key] = value
	}
	return out, nil
}

func WriteFsFile(pathToFile string, data []byte) error {
	return WriteFile(Fs, pathToFile, data)
}

func WriteMemFile(pathToFile string, data []byte) error {
	return WriteFile(Mem, pathToFile, data)
}

func WriteFile(fs afero.Fs, pathToFile string, data []byte) error {
	if err := fs.MkdirAll(filepath.Dir(pathToFile), os.ModePerm); err != nil {
		return nil
	}
	return afero.WriteFile(fs, pathToFile, data, os.ModePerm)
}
