package main

import "os"

func append_file(lines string) error {
	pwd :="/mnt/LONTAS/ExpControl/dire15"
	file, err := os.OpenFile(pwd + "/results.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
    		return err
	}
	defer file.Close()

	if _, err = file.WriteString(lines); err != nil {
    		return err
	}

	return err
}
