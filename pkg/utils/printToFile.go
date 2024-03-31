package utils

import "os"

func PrintToFile(pathWithName string, data []byte) error {
	err := func() error {
		f, err := os.OpenFile(pathWithName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.FileMode(0644))
		if err != nil {
			return err
		}
		_, err = f.Write(data)
		if err1 := f.Close(); err1 != nil && err == nil {
			err = err1
		}
		return err
	}()
	if err != nil {
		return err
	}

	return nil

}
