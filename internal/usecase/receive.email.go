package usecase

import (
	"context"
	"fmt"

	"github.com/burungbangkai/fakesmtp/internal/model"
)

type ReceiveEmail func(ctx context.Context, cfg *model.UserInboxConfig, raw []byte) error

func NewEmailReceiver(
	maxEmailSize int,
) ReceiveEmail {
	return func(ctx context.Context, cfg *model.UserInboxConfig, raw []byte) error {
		// reject if email size to big
		if len(raw) > maxEmailSize {
			return EmailContentToBig
		}
		// compose RawEmail, and
		rawMsg := &model.RawEmail{
			UserID:    cfg.UserID,
			InboxName: cfg.InboxName,
			RawMsg:    raw,
		}
		fmt.Print(rawMsg) // ignore this
		// sent it to be consumed/process by other service or other process
		return nil
	}
}
