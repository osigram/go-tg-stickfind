package ocr

type OCR interface {
	ParseImage(image []byte, width, height int) (string, error)
}

type Getter = func(apiKey string) (OCR, error)
