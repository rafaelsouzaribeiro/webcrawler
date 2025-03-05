package repository

import "fmt"

func (r *Iredis) InsertVisitedPage(url string) error {
	err := r.Db.Incr(ctx, url).Err()
	if err != nil {
		return fmt.Errorf("failed to increment url %s: %w", url, err)
	}
	return nil
}
