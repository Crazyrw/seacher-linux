package utils

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"testBtree/searcher/btree"
)

var AllStopsWords = make(map[string]int)

func init() {
	getAllStopWords()
}

// RemovePunctuation 移除所有的标点符号
func RemovePunctuation(str string) string {
	reg := regexp.MustCompile(`\p{P}+`)
	return reg.ReplaceAllString(str, "")
}

// RemoveSpace 移除所有的空格
func RemoveSpace(str string) string {
	reg := regexp.MustCompile(`\s+`)
	return reg.ReplaceAllString(str, "")
}

//百度停用词
func getAllStopWords() {
	file, err := os.Open("searcher/words/stopwords.txt")
	if err != nil {
		log.Fatalln("stopwords.txt 打开失败")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		AllStopsWords[scanner.Text()] = 1
	}
}

func OpenFile(path string) *os.File {
	//打开文件 覆盖写入文件
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(path, " 打开失败 ", err.Error())
		return nil
	}
	return file
}

func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// BtreeToDisk protobuf > json > gob 序列化的效率
// BtreeToDisk 对btree进行序列化存储
func BtreeToDisk(btree *btree.BPlusTree, path string) {
	marshal, err := json.Marshal(btree)
	if err != nil {
		log.Fatalln("json序列化失败 ", err.Error())
	}
	log.Println("json序列化结果 ", marshal)
	err = ioutil.WriteFile(path, marshal, 0644)
	if err != nil {
		log.Fatalln("写入 ", path, "失败 ", err.Error())
	}
}
