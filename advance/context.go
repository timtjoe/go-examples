package advance

import (
	"context"
	"time"
)


func fetchUser(ctx context.Context, id string)(*User, error){
	select {
	case <-ctx.Done(): return nil, ctx.Err()
	case result := <-db.QueryAsync(id): return result.User, result.Err
	}
}

// Creating contenxt
func UseCtx(){
	ctx := context.Background()
	ctx := context.TODO()

	// Derived:
	ctx, cancel := context.WithCancel(parent);
	ctx, cancel := context.WithTimeout(parent, 5*time.Second);
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(10*time.Minute));
	ctx, cancel := context.WithValue(parent, "user_id", 42)

	defer cancel();
}