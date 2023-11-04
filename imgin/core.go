package imgin

import (
	"bytes"
	"database/sql/driver"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/Kagami/go-face"
	resph "github.com/hexcraft-biz/misc/resp"
	_ "github.com/neofelisho/apng"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"github.com/vincent-petithory/dataurl"
	_ "golang.org/x/image/webp"
)

// ================================================================
//
// ================================================================
type ImgInput struct {
	Src        string      `json:"src" form:"src" binding:"required"`
	Image      image.Image `json:"-" form:"-"`
	JpegBytes  []byte      `json:"-" form:"-"`
	DirUploads string      `json:"-" form:"-"`
}

func (i *ImgInput) Validate() *resph.Resp {
	u, err := url.Parse(i.Src)
	if err != nil {
		return resph.ErrBadRequest
	}

	switch {
	case strings.HasPrefix(u.Scheme, "http"):
		i.Src = u.String()
		if resp, err := http.Get(i.Src); err != nil {
			return resph.NewError(http.StatusServiceUnavailable, err, nil)
		} else {
			if resp.StatusCode >= 400 {
				return resph.ErrBadRequest
			} else if img, err := DecodeImageFromResponse(resp); err != nil {
				return err
			} else if jpegbytes, err := EncodeToJpeg(img); err != nil {
				return err
			} else {
				i.Image = img
				i.JpegBytes = jpegbytes
			}
		}

	case strings.HasPrefix(u.Scheme, "data"):
		if du, err := dataurl.DecodeString(i.Src); err != nil {
			return resph.ErrBadRequest
		} else if img, _, err := image.Decode(bytes.NewReader(du.Data)); err != nil {
			return resph.NewError(http.StatusInternalServerError, err, nil)
		} else if jpegbytes, err := EncodeToJpeg(img); err != nil {
			return err
		} else {
			i.Image = img
			i.JpegBytes = jpegbytes
		}

	case i.DirUploads != "":
		if file, err := os.Open(filepath.Join(i.DirUploads, i.Src)); err != nil {
			return resph.NewError(http.StatusInternalServerError, err, nil)
		} else {
			defer file.Close()
			if img, _, err := image.Decode(file); err != nil {
				return resph.NewError(http.StatusInternalServerError, err, nil)
			} else if jpegbytes, err := EncodeToJpeg(img); err != nil {
				return err
			} else {
				i.Image = img
				i.JpegBytes = jpegbytes
			}
		}

	default:
		return resph.ErrBadRequest
	}

	return nil
}

// ================================================================
//
// ================================================================
const FaceDistThreshold = 0.15

type Threshold float64

func (t *Threshold) Validate() {
	if *t == 0.0 {
		*t = FaceDistThreshold
	} else if *t < 0.0 {
		*t = 0.01
	} else if *t > 0.99 {
		*t = 0.99
	}
}

// ================================================================
//
// ================================================================
const (
	IMAGE_APNG   = "image/apng"
	IMAGE_AVIF   = "image/avif"
	IMAGE_GIF    = "image/gif"
	IMAGE_JPEG   = "image/jpeg"
	IMAGE_PNG    = "image/png"
	IMAGE_SVGXML = "image/svg+xml"
	IMAGE_WEBP   = "image/webp"
)

var ImageMIMETypes = []string{
	IMAGE_APNG,
	//IMAGE_AVIF,
	IMAGE_GIF,
	IMAGE_JPEG,
	IMAGE_PNG,
	//IMAGE_SVGXML,
	IMAGE_WEBP,
}

// ================================================================
//
// ================================================================
func DecodeImageFromResponse(resp *http.Response) (image.Image, *resph.Resp) {
	defer resp.Body.Close()
	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, resph.NewError(http.StatusInternalServerError, err, nil)
	}
	return img, nil
}

// ================================================================
//
// ================================================================
func EncodeToJpeg(img image.Image) ([]byte, *resph.Resp) {
	buf := new(bytes.Buffer)
	err := jpeg.Encode(buf, img, nil)
	return buf.Bytes(), resph.NewError(http.StatusInternalServerError, err, nil)
}

// ================================================================
//
// ================================================================
func JpegToDataUrl(payload []byte) string {
	return "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(payload)
}

// ================================================================
//
// ================================================================
func CropImage(img image.Image, crop image.Rectangle) (image.Image, *resph.Resp) {
	type subImager interface {
		SubImage(r image.Rectangle) image.Image
	}

	// img is an Image interface. This checks if the underlying value has a
	// method called SubImage. If it does, then we can use SubImage to crop the
	// image.
	simg, ok := img.(subImager)
	if !ok {
		return nil, resph.NewErrorWithMessage(http.StatusInternalServerError, "Image cropping failed", nil)
	}

	return simg.SubImage(crop), nil
}

// ================================================================
//
// ================================================================
func Mp4ToImages(sour, destDir string, fps float32) error {
	return ffmpeg.Input(sour).Output(
		filepath.Join(destDir, "%04d.jpeg"),
		ffmpeg.KwArgs{
			"v":      "error",
			"vf":     fmt.Sprintf("fps=%.2f", fps),
			"format": "image2",
			"vcodec": "mjpeg",
		},
	).Run()
}

// ================================================================
//
// ================================================================
const DimensionCount = 128

func SquaredDist(f1, f2 face.Descriptor) float64 {
	sum, diff := float64(0), float64(0)
	for i := 0; i < DimensionCount; i += 1 {
		diff = float64(f1[i] - f2[i])
		sum += diff * diff
	}

	return sum
}

// ================================================================
//
// ================================================================
type Descriptor face.Descriptor

func (d Descriptor) Value() (driver.Value, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, d); err != nil {
		return nil, err
	} else {
		return buf.Bytes(), nil
	}
}

func (d *Descriptor) Scan(src any) error {
	if src != nil {
		buf := bytes.NewReader(src.([]byte))
		return binary.Read(buf, binary.LittleEndian, d)
	}

	return nil
}
