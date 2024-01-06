package qi

import (
	"bufio"
	"fmt"
	"github.com/gookit/color"
	"os"
	"strings"
)

var (
	green  = []*color.Style256{color.S256(46), color.S256(47), color.S256(48), color.S256(49), color.S256(50), color.S256(51)}
	pink   = []*color.Style256{color.S256(214), color.S256(215), color.S256(216), color.S256(217), color.S256(218), color.S256(219)}
	yellow = []*color.Style256{color.S256(226), color.S256(227), color.S256(228), color.S256(229), color.S256(230), color.S256(231)}
)

func ReadLinesFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("打开文件错误: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("文件读取错误: %v", err)
	}
	return lines, nil
}

func gradient(text string, coloRR []*color.Style256) string {
	lines := strings.Split(text, "\n")

	var output string

	t := len(text) / 6
	i := 0
	j := 0
	for l := 0; l < len(lines); l++ {
		str := strings.Split(lines[l], "")
		for _, x := range str {
			j++
			output += coloRR[i].Sprint(x)
			if j > t {
				i++
				j = 0
			}
		}
		if len(lines) != 0 {
			output += "\n"
		}
	}

	return strings.TrimRight(output, "\n")
}

func Logo() {
	logo1 := "                                        \n                                        \n██\\   ██\\  ██████\\  ███████\\   ██████\\  \n██ |  ██ |██  __██\\ ██  __██\\ ██  __██\\ \n██ |  ██ |██ /  ██ |██ |  ██ |██ /  ██ |\n██ |  ██ |██ |  ██ |██ |  ██ |██ |  ██ |\n\\███████ |\\██████  |██ |  ██ |\\███████ |\n \\____██ | \\______/ \\__|  \\__| \\____██ |\n██\\   ██ |                    ██\\   ██ |\n\\██████  |                    \\██████  |\n \\______/                      \\______/ "
	fmt.Println(gradient(logo1, yellow))
	fmt.Println(gradient("by qi4l", yellow))
}
