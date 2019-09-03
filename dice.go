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
		faces int
		dices int
	)

	// コマンドラインオプションの設定
	flag.IntVar(&faces, "f", 6, "The number of dice faces")
	flag.IntVar(&dices, "d", 1, "The number of dices to throw")
	flag.Parse()

	// サイコロを振り, 出目を出力
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < dices; i++ {
		fmt.Printf("%d ", (rand.Intn(faces) + 1))
	}
	fmt.Println()
}
