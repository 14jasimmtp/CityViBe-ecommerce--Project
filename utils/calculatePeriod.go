package utils

import "time"

func CalcualtePeriodDate(period string) (time.Time, time.Time) {
	now := time.Now().Local()

	switch period {
	case "weekly":
		start := now.AddDate(0, 0, -6)
		return start, now
	case "monthly":
		start := now.AddDate(0, -1, 0)
		return start, now
	case "yearly":
		start := now.AddDate(-1, 0, 0)
		return start, now
	default:
		start := now.AddDate(0, 0, -6)
		return start, now
	}
}
