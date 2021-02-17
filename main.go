package main

import (
	"flag"
	"fmt"
	"github.com/sethgrid/multibar"
	"github.com/xfrr/goffmpeg/transcoder"
	"os"
	"strings"
	"sync"
)

func main() {

	flag.Parse()
	cExt := flag.Arg(0)
	aExt := flag.Arg(1)

	if cExt == "" || aExt == "" {
		fmt.Println("Format: ejConvert current_extension desired_extension")
		fmt.Println("ejConvert avi mp4")
	} else {
		path, _ := os.Getwd()
		d, _ := os.Open(path)
		defer d.Close()
		fi, _ := d.Readdir(-1)
		for _, fi := range fi {
			if fi.Mode().IsRegular() && strings.Contains(fi.Name(), "."+cExt) {
				name := fi.Name()[:strings.LastIndex(fi.Name(), ".")]
				ext := fi.Name()[strings.LastIndex(fi.Name(), ".")+1:]
				if ext == cExt {
					trans := new(transcoder.Transcoder)
					err := trans.Initialize(fi.Name(), name+"."+aExt)
					if err != nil {
						fmt.Println(err)
					}
					done := trans.Run(true)
					progress := trans.Output()

					progressBars, _ := multibar.New()
					smallTotal1 := 100
					barProgress1 := progressBars.MakeBar(smallTotal1, fi.Name())

					go progressBars.Listen()
					wg := &sync.WaitGroup{}
					wg.Add(1)
					//-------------------
					go func() {

						for msg := range progress {
							barProgress1(int(msg.Progress))
						}
						wg.Done()
					}()
					//-------------------
					wg.Wait()
					err = <-done
				}
			}
		}
	}
}
