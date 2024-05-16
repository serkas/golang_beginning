package items

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
)

const (
	viewsCounterKeyPrefix = "item_views"
)

type ViewsTracker struct {
	cli *redis.Client
}

func NewViewsTracker(cli *redis.Client) *ViewsTracker {
	return &ViewsTracker{cli: cli}
}

func (lt *ViewsTracker) View(ctx context.Context, itemID int) error {
	return lt.cli.ZIncrBy(ctx, viewsCounterKeyPrefix, 1, strconv.Itoa(itemID)).Err()
}

func (lt *ViewsTracker) GetTopViewed(ctx context.Context, num int) (viewed []int, err error) {
	results, err := lt.cli.ZRevRange(ctx, viewsCounterKeyPrefix, 0, int64(num)).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, fmt.Errorf("getting top viewed items: %w", err)
	}

	for _, idStr := range results {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("parsing top result: %w", err)
		}
		viewed = append(viewed, int(id))
	}

	return viewed, nil
}
