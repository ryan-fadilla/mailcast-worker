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

func WorkersServeTest() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: config.RedisAddr},
		asynq.Config{Concurrency: 10},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.TypeWelcomeEmail, tasks.HandleWelcomeEmailTask)
	mux.HandleFunc(tasks.TypeReminderEmail, tasks.HandleReminderEmailTask)

	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}

func sendWelcomeEmail(ctx context.Context, t *asynq.Task) error {
	var p tasks.EmailTaskPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	log.Printf(" [*] Send Welcome Email to User %d", p.UserID)
	return nil
}

func sendReminderEmail(ctx context.Context, t *asynq.Task) error {
	var p tasks.EmailTaskPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	log.Printf(" [*] Send Reminder Email to User %d", p.UserID)
	return nil
}

func handlers(ctx context.Context, t *asynq.Task) error {
	switch t.Type() {
	case "email:welcome":
		var p tasks.EmailTaskPayload
		if err := json.Unmarshal(t.Payload(), &p); err != nil {
			return err
		}
		log.Printf(" [*] Send Welcome Email to User %d", p.UserID)

	case "email:reminder":
		var p tasks.EmailTaskPayload
		if err := json.Unmarshal(t.Payload(), &p); err != nil {
			return err
		}
		log.Printf(" [*] Send Reminder Email to User %d", p.UserID)

	default:
		return fmt.Errorf("unexpected task type: %s", t.Type())
	}
	return nil
}
