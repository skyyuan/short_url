package conf

import (
	"os"
	"bufio"
	"io"
	"strings"
	"errors"
)

type Config struct {
	content map[string]string
}

var config Config

func init() {
	config.content = make(map[string]string)

	file, err := os.Open("/home/zereker/Documents/Go/src/short_url/conf/mongo.conf")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	in := bufio.NewReader(file)
	for {
		line, _, err := in.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		s := strings.TrimSpace(string(line))
		split := strings.Split(s, " = ")
		config.content[split[0]] = split[1]
	}
}

func Read(key string) (string, error) {
	value := config.content[key]
	if value != "" {
		return value, nil
	}
	return "", errors.New("不存在" + key)
}
