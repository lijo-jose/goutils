package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/lijo-jose/gffmpeg/pkg/gffmpeg"
	"github.com/lijo-jose/goutils/pkg/ffmpeg"

)

// Configuration defines the settings required during the app initiation
type Configuration struct {
	inputFilePath string
	destination   string
}

const (
	inputFilePathArg = "input"
	destinationArg   = "dest"
)

func init() {

	pflag.String(inputFilePathArg, "./dummy.ts", "input file for processing")
	pflag.String(destinationArg, "./", "destination directory for storing results and intermediate files")

	viper.AutomaticEnv()
	viper.BindPFlag(inputFilePathArg, pflag.Lookup(inputFilePathArg))
	viper.BindPFlag(destinationArg, pflag.Lookup(destinationArg))

	pflag.Parse()
}

func main(){
	config := Configuration{
		inputFilePath: viper.GetString(inputFilePathArg),
		destination:   viper.GetString(destinationArg),
	}

	ff, err := gffmpeg.NewGFFmpeg("/usr/bin/ffmpeg")
	if err != nil {
		fmt.Println(err)
		return
	}
	svc, err := ffmpeg.New(ff)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = svc.ExtractFrames(config.inputFilePath, config.destination, 3)
	if err != nil {
		fmt.Println(err)
	}
}
