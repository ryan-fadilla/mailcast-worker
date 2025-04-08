package tasks

import (
	"context"
	"encoding/json"
	"log"
	"mailcast-worker/services"
	"time"

	"github.com/hibiken/asynq"
)

// A list of task types.
const (
	TypeNotif = "type:notif"
)

type SchedulerPayload struct {
	Payload    map[string]interface{}
	Phone      string
	ScheduleAt time.Time
}

func HandleSchedulerNotifTask(ctx context.Context, t *asynq.Task) error {
	var p SchedulerPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	log.Printf(" [*] Send Whatsapp to User %d", p.Phone)

	services.SendWaMessage(p.Payload)

	return nil
}
