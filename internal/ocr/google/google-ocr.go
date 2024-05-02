package google

import (
	"bytes"
	vision "cloud.google.com/go/vision/apiv1"
	"context"
	"fmt"
	"go-tg-stickfind/internal/ocr"
	"strings"
)

type OCR struct {
	Key string
}

func NewOCR(apiKey string) (ocr.OCR, error) {
	return ocr.OCR(&OCR{Key: apiKey}), nil
}

func (o *OCR) ParseImage(imageBytes []byte, _, _ int) (string, error) {
	ctx := context.Background()

	// Creates a client.
	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to create google ocr client: %w", err)
	}
	defer client.Close()

	image, err := vision.NewImageFromReader(bytes.NewReader(imageBytes))
	if err != nil {
		return "", fmt.Errorf("failed to read image: %w", err)
	}

	labels, err := client.DetectTexts(ctx, image, nil, 10)
	if err != nil {
		return "", fmt.Errorf("failed to detect labels: %w", err)
	}

	texts := make([]string, 0, len(labels))
	for _, label := range labels {
		texts = append(texts, label.Description)
	}

	return strings.Join(texts, "\n"), nil
}
