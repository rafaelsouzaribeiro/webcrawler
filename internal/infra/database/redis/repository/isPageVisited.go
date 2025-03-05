package repository

import "fmt"

func (r *Iredis) IsPageVisited(url string) (bool, error) {
	exists, err := r.Db.SIsMember(ctx, "sites", url).Result()
	if err != nil {
		return false, fmt.Errorf("failed to check if url %s is in set: %w", url, err)
	}
	return exists, nil
}
