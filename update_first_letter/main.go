package main

import (
	"log"
	"strings"
	"unicode"

	"eglass.com/entities"
	pinyin "github.com/mozillazg/go-pinyin"

	"eglass.com/utils"
)

var (
	mysql, sqlErr = utils.NewMysql(false, false)
)

func main() {
	if sqlErr != nil {
		log.Panicf("err : %v", sqlErr)
	}
	var result []entities.EOptometryUser
	err := mysql.SelectFrom("e_optometry_user").Where("first_letter is null").All(&result)
	if err != nil {
		log.Printf("err : %v", err)
	}

	// text := "#@%中国人"
	for _, member := range result {
		name := member.Name
		if name != "" {
			log.Printf("name is %s \n", name)
			names := []rune(name)
			first := string(names[0])
			var firstLetter string
			// patter := `^[A-Za-z]+$`
			// reg := regexp.MustCompile(patter)
			log.Printf("first name is %s \n", first)
			if unicode.Is(unicode.Scripts["Han"], names[0]) {
				// return true
				a := pinyin.NewArgs()
				pinyin := pinyin.Pinyin(first, a)
				log.Printf("letters is %v \n", pinyin)
				if (len(pinyin) != 0) && (len(pinyin[0][0]) != 0) {
					log.Printf("first name pinyin is %v\n", pinyin[0][0])
					firstLetter = string([]rune(pinyin[0][0])[0])
					log.Printf("first_letter is %v\n", firstLetter)
				}
			} else {
				log.Printf("first_letter is %v\n", first)
				firstLetter = first
			}
			firstLetter = strings.ToUpper(firstLetter)
			mysql.Update("e_optometry_user").Set("first_letter=?", firstLetter).Where("id=?", member.Id).Exec()
		}
	}

	log.Println(" over ")
}
