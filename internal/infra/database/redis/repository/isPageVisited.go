package repository

import "fmt"

func (r *Iredis) IsPageVisited(url string) (bool, error) {
	val, err := r.Db.Get(ctx, url).Result()
	if err != nil {
		return false, fmt.Errorf("failed to get key %s: %w", url, err)
	}

	if val == "" {
		return false, nil
	}

	return true, nil
}
