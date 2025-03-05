package repository

import "fmt"

func (r *Iredis) InsertVisitedPage(url string) error {
	key := fmt.Sprintf("visited:%s", url)
	err := r.Db.Incr(ctx, key).Err()
	if err != nil {
		return fmt.Errorf("failed to increment url %s: %w", url, err)
	}
	return nil
}
