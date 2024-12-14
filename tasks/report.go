package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hibiken/asynq"
)

type ReportPayload struct {
	UserID int
}

func NewReportTask(userID int) (*asynq.Task, error) {
	payload, err := json.Marshal(ReportPayload{
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeReport, payload), nil
}

func ReportTaskHandler(ctx context.Context, t *asynq.Task) error {
	var p ReportPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	log.Printf("Preparing report for the user ID: %d", p.UserID)
	return nil
}
