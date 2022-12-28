package main

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"github.com/xuri/excelize/v2"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	codes := generateRedeemCode(10)
	st := time.Now()

	f := excelize.NewFile()
	// 创建一个工作表
	index := f.NewSheet("Sheet1")
	// 单元格格式为居中
	style, _ := f.NewStyle(`{"alignment":{"horizontal": "center", "vertical": "center"}}`)
	// 设置列宽度
	if err := f.SetColWidth("Sheet1", "A", "F", 16); err != nil {
		fmt.Println(err)
	}
	// 设置表头
	_ = f.SetCellValue("Sheet1", "A1", "ID")
	_ = f.SetCellValue("Sheet1", "B1", "code")
	_ = f.SetCellValue("Sheet1", "C1", "qrcode")

	for i, code := range codes {
		i += 2
		// 设置行高度
		_ = f.SetRowHeight("Sheet1", i, 90)
		// 设置单元格格式
		_ = f.SetCellStyle("Sheet1", "A"+strconv.Itoa(i), "B"+strconv.Itoa(i), style)
		// 二维码内容
		content := "https://www.baidu.com?code=" + code
		// 生产二维码
		qrCode, _ := qrcode.Encode(content, qrcode.Medium, 112)
		// 设置单元格的值
		_ = f.SetCellValue("Sheet1", "A"+strconv.Itoa(i), i)
		_ = f.SetCellValue("Sheet1", "B"+strconv.Itoa(i), code)
		_ = f.AddPictureFromBytes("Sheet1", "C"+strconv.Itoa(i), "", "qrcode", ".png", qrCode)
	}
	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)
	// 根据指定路径保存文件
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}

	sp := time.Since(st)
	fmt.Printf("总耗时：%f s", sp.Seconds())
}

func generateRedeemCode(cnt int) []string {
	initSource := "ABCDEFGHIGKMLNOPQRSTUVWXYZ"
	size := len(initSource)
	var res []string
	rand.Seed(time.Now().Unix())
	for i := 0; i < cnt; i++ {
		index := rand.Intn(size)
		res = append(res, string(initSource[index]))
	}
	return res
}
