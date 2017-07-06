package main

import (
	"bytes"
	"bufio"
    "fmt"
    "io"
    "os"
)
	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func Use(vals ...interface{}) {
    for _, val := range vals {
        _ = val
    }
}

func main() {
	// open iplog.txt. Send error if empty or not in folder.
	file, err := os.Open("iplog.txt")
	check(err)
	
	reader := bufio.NewReader(file)
	
	first4  := []byte("0000")
	filename := "0000"
	
	// temp is used as a placeholder for the following loop. It will delete itself at the end of the program
	currentfile, err := os.OpenFile("temp", os.O_RDONLY|os.O_CREATE, 0666)
	
	
	for {
		line, _, err := reader.ReadLine()
		
		// once it hits the end of file, exit loop
		if err == io.EOF {
			break
		}
		
		// if new day, create new file, append line to file
		if !bytes.Equal(first4, line[0:4]) {
			first4 = line[0:4]
			filename = fmt.Sprintf("iplog_archive_%s2017.txt", first4)
			currentfile, err = os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0666)
			check(err)
		}
		
		// else still on current day, add line to current file
		n, err := currentfile.WriteString(string(line)+"\n")
		Use(n)
	}
	
	// delete temp and exit
	rem := os.Remove("temp")
	Use(rem)
}