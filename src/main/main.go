package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"../hack"
	"cloud.google.com/go/firestore"
)

var (
	firebaseID string
	exportPath string
	importPath string
)

func init() {
	flag.StringVar(&firebaseID, "p", "", "要操作的專案")
	flag.StringVar(&importPath, "i", "", "載入路徑")
	flag.StringVar(&exportPath, "e", "", "匯出路徑")
	flag.Parse()
}

func main() {
	// 認證
	ctx := context.Background()
	cli, err := firestore.NewClient(ctx, firebaseID)
	hack.PrintErrorExit(err)

	// 完成結束
	defer cli.Close()

	// 如果有匯出路徑則不進行匯入
	if len(exportPath) > 0 {
		ExportStored(ctx, cli, exportPath)
	} else if len(importPath) > 0 {
		// TODO: 匯入
	}
}

// ExportStored 匯出
func ExportStored(ctx context.Context, cli *firestore.Client, path string) {
	// 取得目前工作目錄
	wd, err := os.Getwd()
	hack.PrintErrorExit(err)
	// 取得所有集合
	snap, err := cli.Collections(ctx).GetAll()
	hack.PrintError(err)
	// 每個集合抓資料
	for _, snap := range snap {
		snap, err := snap.Documents(ctx).GetAll()
		hack.PrintError(err)
		for _, snap := range snap {
			b, _ := json.Marshal(snap.Data())
			if len(path) > 0 {
				newPath := fmt.Sprintf("%s/%s/%s", wd, path, snap.Ref.Parent.Path)
				_, err := os.Stat(newPath)
				if os.IsNotExist(err) {
					err = os.MkdirAll(newPath, os.ModePerm)
					hack.PrintError(err)
				}
				err = ioutil.WriteFile(fmt.Sprintf("%s/%s", newPath, fmt.Sprintf("%s.json", snap.Ref.ID)), b, os.ModePerm)
				hack.PrintError(err)
			} else {
				fmt.Printf("%s: %v\n", snap.Ref.Path, string(b))
			}
		}
	}
}
