package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// コマンドラインオプションで与える値の変数定義
	var (
		faces uint
		dices uint
	)

	// コマンドラインオプションの設定
	flag.UintVar(&faces, "f", 6, "The number of dice faces")
	flag.UintVar(&dices, "d", 1, "The number of dices to throw")
	flag.Parse()

	// サイコロを振り, 出目を出力
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < int(dices); i++ {
		fmt.Printf("%d ", (rand.Intn(int(faces)) + 1))
	}
	fmt.Println()
}
