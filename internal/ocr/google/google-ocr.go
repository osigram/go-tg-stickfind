package google

import (
	"bytes"
	oldvision "cloud.google.com/go/vision/apiv1"
	vision "cloud.google.com/go/vision/v2/apiv1"
	"cloud.google.com/go/vision/v2/apiv1/visionpb"
	"context"
	"fmt"
	"go-tg-stickfind/internal/ocr"
	"google.golang.org/api/option"
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
	client, err := vision.NewImageAnnotatorRESTClient(ctx, option.WithAPIKey(o.Key))
	if err != nil {
		return "", fmt.Errorf("failed to create google ocr client: %w", err)
	}
	defer client.Close()

	image, err := oldvision.NewImageFromReader(bytes.NewReader(imageBytes))
	if err != nil {
		return "", fmt.Errorf("failed to read image: %w", err)
	}

	resp, err := client.BatchAnnotateImages(ctx, &visionpb.BatchAnnotateImagesRequest{
		Requests: []*visionpb.AnnotateImageRequest{{
			Image:        image,
			Features:     []*visionpb.Feature{{Type: visionpb.Feature_TEXT_DETECTION, MaxResults: 10}},
			ImageContext: nil,
		}},
	},
	)
	if err != nil {
		return "", fmt.Errorf("failed to detect labels: %w", err)
	}

	var texts []string
	for _, response := range resp.Responses {
		if response.Error != nil {
			for _, annotations := range response.TextAnnotations {
				texts = append(texts, annotations.Description)
			}
		}
	}

	return strings.Join(texts, "\n"), nil
}
