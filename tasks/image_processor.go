package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

type ImageProcessor struct { // implements asynq.Handler interface
}

type ImageProcessingPayload struct {
	ImageUrl string
}

func NewImageProcessor() *ImageProcessor {
	return &ImageProcessor{}
}

func (processor *ImageProcessor) ProcessTask(
	ctx context.Context, t *asynq.Task) error {
	var p ImageProcessingPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	log.Printf("Resizing image: URL=%s", p.ImageUrl)
	time.Sleep(5 * time.Second)
	return nil
}

func NewImageProcessingTask(imageUrl string) (*asynq.Task, error) {
	payload, err := json.Marshal(ImageProcessingPayload{
		ImageUrl: imageUrl,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeImageProcessing, payload), nil
}
