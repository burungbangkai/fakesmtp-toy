package usecase

import (
	"context"
	"testing"

	"github.com/burungbangkai/fakesmtp/internal/model"
	"github.com/burungbangkai/fakesmtp/internal/port"
)

func TestAddUser(t *testing.T) {
	var inserter port.InsertUserInboxConfig = func(ctx context.Context, cfg *model.UserInboxConfig) error {
		if cfg == nil {
			t.Fatal("nill")
		}
		return nil
	}
	fn := NewUserInboxAdder(10, inserter, 10, 10)
	e, err := fn(context.Background(), "id")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(e)
}
