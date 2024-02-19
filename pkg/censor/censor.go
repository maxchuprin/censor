//инструментарий проверки строк на вхождение заданных бан-слов
package censor

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

var banlist = "banlist.json"
var list []string
var banRegex *regexp.Regexp

func init() {
	list = readBanList()

	//делаем регулярку на все бан-слова
	join := strings.Join(list, "|")
	banRegex = regexp.MustCompile(join)
}

//Проверяет строку на вхождение бан-слов их бан-листа
func Censored(s string) bool {
	if banRegex.MatchString(s) {
		return true
	}
	return false
}

//Парсинг бан-листа
func readBanList() []string {
	f, err := os.OpenFile(banlist, os.O_RDONLY, 0777)
	if err != nil {
		log.Fatal("Cannot read ban list file: ", err)
	}
	defer f.Close()

	buf, err := io.ReadAll(f)
	if err != nil {
		log.Fatal("Cannot read ban list file: ", err)
	}

	conf := struct {
		List []string
	}{}

	err = json.Unmarshal(buf, &conf)
	if err != nil {
		log.Fatal("Cannot parse ban list : ", err)
	}

	return conf.List
}
