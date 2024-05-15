package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/vcaesar/bitmap"
)

// "go-vgo/robotgo"
var lock = 0

// var ifnotresponding bool = false
// var pid string

func activeIdle() {
	Idle := "IdleDragons"
	Idlepids, err := robotgo.FindIds(Idle)
	Idlepid := Idlepids[0]
	if err != nil {
		fmt.Println(err)
		return
	}
	err = robotgo.ActivePid(Idlepid)
	if err != nil {
		fmt.Println(err)
	}

	// mdata := robotgo.GetHandPid(Idlepid)

	// robotgo.SetActive(mdata)
}

func check_and_tap() {
	tempscreen := robotgo.CaptureScreen()
	// img := robotgo.ToImage(tempscreen)
	// imgo.Save("test.png", img)
	defer robotgo.FreeBitmap(tempscreen)

	brivbitmap1 := bitmap.Open("briv_unlock.png")
	brivbitmap2 := bitmap.Open("briv_unlock_dark.png")
	fx1, fy1 := bitmap.Find(brivbitmap1, tempscreen)
	fx2, fy2 := bitmap.Find(brivbitmap2, tempscreen)
	// println(fx, fy)
	if (fx1 > 0 && fy1 > 0) || (fx2 > 0 && fy2 > 0) {
		println("find!")
		activeIdle()
		robotgo.KeySleep = 100
		robotgo.KeyTap("f2")
		robotgo.KeyTap("f5")
		robotgo.KeyTap("f5")
		lock = 1
	} else {
		println("not find!")
	}
}

func main() {
	for {
		if lock == 1 {
			time.Sleep(60 * time.Second)
			activeIdle()
			check_and_tap()
			lock = 0
		}
		check_and_tap()
		time.Sleep(2 * time.Second)
	}
	// check_and_tap()
}
