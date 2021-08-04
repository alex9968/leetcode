package  main

import (
	"fmt"
	"io"
	"os"
	"strings"
)


func main() {
	IOWriter()
}

func IOReader() {
	reader := strings.NewReader("Clear is better than clever")
	p := make([]byte, 4)

	for {
		n, err := reader.Read(p)
		fmt.Println("n:", n)
		if err != nil{
			if err == io.EOF {
				fmt.Println("EOF:", n)
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(n, string(p[:n]))
	}
}

func IOWriter() {
	proverbs := []string{
		"Channels orchestrate mutexes serialize",
		"Cgo is not Go",
		"Errors are values",
		"Don't panic",
	}
	//var writer bytes.Buffer
	//for _, p := range proverbs {
	//	n, err := writer.Write([]byte(p))
	//	if err != nil {
	//		fmt.Println(err)
	//		os.Exit(1)
	//	}
	//	if n != len(p) {
	//		fmt.Println("failed to write data")
	//		os.Exit(1)
	//	}
	//}
	fmt.Print(proverbs)
	//fmt.Println(writer.String())
}