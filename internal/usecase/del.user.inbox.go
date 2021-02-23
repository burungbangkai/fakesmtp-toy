package usecase

import "context"

type DeleteUserInbox func(ctx context.Context, userID, inboxName string) error
