package ffmpeg

import (
	"strconv"

	"github.com/lijo-jose/gffmpeg/pkg/gffmpeg"
)

type Service interface {
	ExtractFrames(inFile, outDir string, fps int) error
}

type service struct {
	ff gffmpeg.GFFmpeg
}

func New(ff gffmpeg.GFFmpeg) (Service, error) {
	return &service{ff: ff}, nil
}

func (svc *service) ExtractFrames(inFile, outDir string, fps int) error {
	bd := gffmpeg.NewBuilder()
	bd = bd.SrcPath(inFile).VideoFilters("fps=" + strconv.Itoa(fps)).DestPath(outDir + "/frames%05d.jpg")
	svc.ff = svc.ff.Set(bd)

	ret  := svc.ff.Start(nil)
	if ret.Err != nil {
		return ret.Err
	}
	return nil
}
