package imgin

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/Kagami/go-face"
	"github.com/kettek/apng"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"github.com/vincent-petithory/dataurl"
	"golang.org/x/image/webp"
)

// ================================================================
//
// ================================================================
var (
	ErrInvalidInput        = errors.New("Invalid input")               // 400
	ErrServiceUnavailable  = errors.New("Service unavailable")         // 503
	ErrDecodeImageFailed   = errors.New("Decode image failed")         // 500
	ErrEncodeToJpegFailed  = errors.New("Encode to jpeg failed")       // 500
	ErrFileNotExists       = errors.New("File is not exists")          // 500
	ErrReadFileFailed      = errors.New("Read file failed")            // 500
	ErrParseMimeTypeFailed = errors.New("Mime type parse failed")      // 500
	ErrMimeType            = errors.New("Request URL is not an image") // 400
	ErrPayloadReading      = errors.New("Cannot read payload body")    // 500
	ErrContent             = errors.New("Invalid Content")             // 400
)

func ErrStatusCode(err error) int {
	switch err {
	case ErrInvalidInput, ErrMimeType, ErrContent:
		return http.StatusBadRequest
	case ErrServiceUnavailable:
		return http.StatusServiceUnavailable
	default:
		return http.StatusInternalServerError
	}
}

// ================================================================
//
// ================================================================
type ImgInput struct {
	Src       string      `json:"src" form:"src" binding:"required"`
	Image     image.Image `json:"-" form:"-"`
	JpegBytes []byte      `json:"-" form:"-"`
}

func (i *ImgInput) Validate(dirUploads string) error {
	if dirUploads != "" {
		if _, err := os.Stat(filepath.Join(dirUploads, i.Src)); err != nil {
			return ErrFileNotExists
		} else if payload, err := os.ReadFile(filepath.Join(dirUploads, i.Src)); err != nil {
			return ErrReadFileFailed
		} else if mediaType := http.DetectContentType(payload); !slices.Contains(ImageMIMETypes, mediaType) {
			return ErrInvalidInput
		} else if img, err := DecodeToImage(payload, mediaType); err != nil {
			return ErrDecodeImageFailed
		} else if jpegbytes, err := EncodeToJpeg(img); err != nil {
			return ErrEncodeToJpegFailed
		} else {
			i.Image = img
			i.JpegBytes = jpegbytes
		}
	} else if u, err := url.Parse(i.Src); err != nil {
		return ErrInvalidInput
	} else {
		switch {
		case strings.HasPrefix(u.Scheme, "http"):
			i.Src = u.String()
			if resp, err := http.Get(i.Src); err != nil {
				return ErrServiceUnavailable
			} else {
				defer resp.Body.Close()
				if resp.StatusCode >= 400 {
					return ErrServiceUnavailable
				} else if payload, mediaType, err := GetImagePayload(resp); err != nil {
					return err
				} else if img, err := DecodeToImage(payload, mediaType); err != nil {
					return ErrDecodeImageFailed
				} else if jpegbytes, err := EncodeToJpeg(img); err != nil {
					return ErrEncodeToJpegFailed
				} else {
					i.Image = img
					i.JpegBytes = jpegbytes
				}
			}
		case strings.HasPrefix(u.Scheme, "data"):
			if du, err := dataurl.DecodeString(i.Src); err != nil {
				return ErrInvalidInput
			} else if mediaType := du.MediaType.ContentType(); !slices.Contains(ImageMIMETypes, mediaType) {
				return ErrInvalidInput
			} else if img, err := DecodeToImage(du.Data, mediaType); err != nil {
				return ErrDecodeImageFailed
			} else if jpegbytes, err := EncodeToJpeg(img); err != nil {
				return ErrEncodeToJpegFailed
			} else {
				i.Image = img
				i.JpegBytes = jpegbytes
			}
		default:
			return ErrInvalidInput
		}
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
func GetImagePayload(resp *http.Response) ([]byte, string, error) {
	if mediaType, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type")); err != nil {
		return nil, "", ErrParseMimeTypeFailed
	} else if !slices.Contains(ImageMIMETypes, mediaType) {
		return nil, "", ErrMimeType
	} else if payload, err := io.ReadAll(resp.Body); err != nil {
		return nil, "", ErrPayloadReading
	} else if mediaType := http.DetectContentType(payload); !slices.Contains(ImageMIMETypes, mediaType) {
		return nil, "", ErrContent
	} else {
		return payload, mediaType, nil
	}
}

// ================================================================
//
// ================================================================
func DecodeToImage(payload []byte, mediaType string) (image.Image, error) {
	switch mediaType {
	case IMAGE_APNG:
		if img, err := apng.Decode(bytes.NewReader(payload)); err != nil {
			return nil, fmt.Errorf("Unable to decode: apng.")
		} else {
			return img, nil
		}

	case IMAGE_AVIF:
		return nil, fmt.Errorf("Unable to decode: avif.")

	case IMAGE_GIF:
		if img, err := gif.Decode(bytes.NewReader(payload)); err != nil {
			return nil, fmt.Errorf("Unable to decode: gif.")
		} else {
			return img, nil
		}

	case IMAGE_JPEG:
		if img, err := jpeg.Decode(bytes.NewReader(payload)); err != nil {
			return nil, fmt.Errorf("Unable to decode: jepg.")
		} else {
			return img, nil
		}

	case IMAGE_PNG:
		if img, err := png.Decode(bytes.NewReader(payload)); err != nil {
			return nil, fmt.Errorf("Unable to decode: png.")
		} else {
			return img, nil
		}

	case IMAGE_SVGXML:
		return nil, fmt.Errorf("Unable to decode: svg+xml.")

	case IMAGE_WEBP:
		if img, err := webp.Decode(bytes.NewReader(payload)); err != nil {
			return nil, fmt.Errorf("Unable to decode: webp.")
		} else {
			return img, nil
		}
	}

	return nil, fmt.Errorf("Invalid image.")
}

// ================================================================
//
// ================================================================
func EncodeToJpeg(img image.Image) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, img, nil); err != nil {
		return nil, fmt.Errorf("Unable to encode: jpeg.")
	}

	return buf.Bytes(), nil
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
func CropImage(img image.Image, crop image.Rectangle) (image.Image, error) {
	type subImager interface {
		SubImage(r image.Rectangle) image.Image
	}

	// img is an Image interface. This checks if the underlying value has a
	// method called SubImage. If it does, then we can use SubImage to crop the
	// image.
	simg, ok := img.(subImager)
	if !ok {
		return nil, fmt.Errorf("image does not support cropping")
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
