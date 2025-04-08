package workers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mailcast-worker/config"
	"mailcast-worker/tasks"

	"github.com/hibiken/asynq"
)

func WorkersServe() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: config.RedisAddr},
		asynq.Config{Concurrency: 10},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.TypeNotif, tasks.HandleSchedulerNotifTask)

	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}

func handler(ctx context.Context, t *asynq.Task) error {
	switch t.Type() {
	case "type:notif":
		var p tasks.SchedulerPayload
		if err := json.Unmarshal(t.Payload(), &p); err != nil {
			return err
		}
		log.Printf(" [*] Send Whatsapp to User %d", p.Phone)

	default:
		return fmt.Errorf("unexpected task type: %s", t.Type())
	}
	return nil
}
