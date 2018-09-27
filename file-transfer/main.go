package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"eglass.com/entities"

	"eglass.com/utils"

	"github.com/denverdino/aliyungo/oss"
)

const ACCESS_KEY_ID = "uslNjzUwTMSXnDF4"
const ACCESS_KEY_SECRET = "XTLgq8JSMIDPgZiGO6WZb7AHyqFPOJ"

var ossBucket *oss.Bucket

func GetOssBucket() *oss.Bucket {
	if ossBucket == nil {
		ossClient := oss.NewOSSClient("oss-cn-shanghai", isProd, ACCESS_KEY_ID, ACCESS_KEY_SECRET, false)
		ossBucket := ossClient.Bucket("epj-images")
		return ossBucket
	}
	return ossBucket
}

type OssFile struct {
	Source int32
	Url    string
}

var fileIdMap = make(map[string]OssFile)
var isProd = true //os.Getenv("prod") == "true"
var mysqlConn, err = utils.NewMysql(isProd, true)

var rediClient = utils.NewRedisClient(isProd)

func main() {
	files, err := rediClient.SMembers("delete-files").Result()
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		err := os.Remove(file)
		if err != nil {
			log.Printf("%v", err)
		}
	}
}
func sync() {
	ossBucket = GetOssBucket()
	if err != nil {
		panic(err)
	}
	var err error
	var colors []entities.NGoodsColorImages
	thirdPhone := "13405591880" // 15652813691

	err = mysqlConn.Select("a.id", "a.img_detail_0", "a.img_detail_1", "a.img_detail_2", "a.img_detail_3", "a.img_detail_4", "a.img_thumb", "a.img_leg", "a.img_face").From("n_goods_color_images a").Join("n_goods2 b").On("a.goods_id=b.id").Where("b.try_owner=? and a.try_on_use = 1", thirdPhone).
		All(&colors)
	if err != nil {
		panic(err)
	}
	log.Println(len(colors))
	log.Print(colors[0])
	for _, ci := range colors {
		checkFile(ci.Id, "img_detail_0", ci.ImgDetail0, true)
		checkFile(ci.Id, "img_detail_1", ci.ImgDetail1, true)
		checkFile(ci.Id, "img_detail_2", ci.ImgDetail2, true)
		checkFile(ci.Id, "img_detail_3", ci.ImgDetail3, true)
		checkFile(ci.Id, "img_detail_4", ci.ImgDetail4, true)
		checkFile(ci.Id, "img_thumb", ci.ImgThumb, true)
		checkFile(ci.Id, "img_leg", ci.ImgLeg, true)
		checkFile(ci.Id, "img_face", ci.ImgFace, true)
	}
	err = mysqlConn.Select("a.id", "a.img_detail_0", "a.img_detail_1", "a.img_detail_2", "a.img_detail_3", "a.img_detail_4", "a.img_thumb", "a.img_leg", "a.img_face").From("n_goods_color_images a").Join("n_goods2 b").On("a.goods_id=b.id").Where("b.try_owner!=? and a.try_on_use = 1", thirdPhone).
		All(&colors)
	if err != nil {
		panic(err)
	}
	log.Println(len(colors))
	log.Print(colors[0])
	for _, ci := range colors {
		checkFile(ci.Id, "img_detail_0", ci.ImgDetail0, false)
		checkFile(ci.Id, "img_detail_1", ci.ImgDetail1, false)
		checkFile(ci.Id, "img_detail_2", ci.ImgDetail2, false)
		checkFile(ci.Id, "img_detail_3", ci.ImgDetail3, false)
		checkFile(ci.Id, "img_detail_4", ci.ImgDetail4, false)
		checkFile(ci.Id, "img_thumb", ci.ImgThumb, false)
		checkFile(ci.Id, "img_leg", ci.ImgLeg, false)
		checkFile(ci.Id, "img_face", ci.ImgFace, false)
	}

}

func checkFile(id int32, columnName string, pathStr entities.NullString, ShouldSyncMap bool) {
	if !pathStr.Valid {
		return
	}
	if strings.Contains(pathStr.String, "https") {
		return
	}
	var err error
	ossFile, ok := fileIdMap[pathStr.String]
	if ok {
		updateUrl(id, columnName, ossFile.Url, ossFile.Source)
		return
	}
	// 文件所在路径
	path := fmt.Sprintf("/epj/www%s", pathStr.String)
	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		log.Printf("file %s may not exist, skip", path)
		mysqlConn.Exec("update n_goods_color_images set try_online = 0 where id = ? ", id)
		return
	}
	arr := strings.Split(path, "/")
	name := arr[len(arr)-1]
	filename := fmt.Sprintf("try-image/%s", name)
	err = ossBucket.PutFile(filename, file, oss.PublicRead, oss.Options{})
	if err != nil {
		log.Printf("file upload fail %s", path)
		return
	}
	url := fmt.Sprintf("https://img.epeijing.cn/%s", filename)
	updateUrl(id, columnName, url, 0)
	file.Close()
	// os.Remove(path)
	rediClient.SAdd("delete-files", path).Err()
	if ShouldSyncMap {
		fileIdMap[pathStr.String] = OssFile{id, url}
	}

}

func updateUrl(id int32, colunName, url string, sourceID int32) {
	mysqlConn.Exec(fmt.Sprintf("update n_goods_color_images set try_online = 1, source_id = ?, %s=? where id = ?", colunName), sourceID, url, id)
	if sourceID > 0 {
		key := fmt.Sprintf("color-goods-%d", sourceID)
		rediClient.SAdd(key, id).Err()
	}
}
