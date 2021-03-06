package pidfile

import (
	"errors"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/spf13/pflag"
)

func WritePIDFile(pidFilePath string) error {
	if pidFilePath == "" {
		return errors.New("empty pidFilePath")
	}
	err := ioutil.WriteFile(pidFilePath, []byte(strconv.Itoa(os.Getpid())), 0644)
	if err != nil {
		return err
	}
	return nil
}

func AddFlag(fs *pflag.FlagSet, pidFilePath *string, defaultValue string) {
	fs.StringVar(pidFilePath, "pid-file", defaultValue, "file to write process id")
}
