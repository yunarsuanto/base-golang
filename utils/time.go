package utils

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/yunarsuanto/base-go/constants"
)

// TimeCountdown struct for time count down
type TimeCountdown struct {
	Total   int
	Days    int
	Hours   int
	Minutes int
	Seconds int
}

// GetRemainingTime function to get remaining time
// Params:
// currentTime: current time
// timeout: time out
// Returns TimeCountdown
func GetRemainingTime(currentTime, timeout time.Time) TimeCountdown {
	difference := timeout.Sub(currentTime)

	total := int(difference.Seconds())
	days := int(total / (60 * 60 * 24))
	hours := int(total / (60 * 60) % 24)
	minutes := int(total/60) % 60
	seconds := int(total % 60)

	return TimeCountdown{
		Total:   total,
		Days:    days,
		Hours:   hours,
		Minutes: minutes,
		Seconds: seconds,
	}
}

// NewNullTime function to handle nullable value in SQL
// Params:
// t: time to be checked
// Returns sql.NullString
func NewNullTime(t time.Time) sql.NullTime {
	if t.IsZero() {
		return sql.NullTime{}
	}
	return sql.NullTime{
		Time:  t,
		Valid: true,
	}
}

// NullTimeDurationScan function to handle pointer duration
// Params:
// t: duration
// defaultDuration default Duration
// Returns duration
func NullTimeDurationScan(t *time.Duration, defaultDuration time.Duration) time.Duration {
	if t != nil {
		if *t != 0 {
			return *t
		}
	}

	return defaultDuration
}

// NextResendTime function to calculate duration to be able to send next mail
// Params:
// lastSend: time email last sent
// interval: interval to resend mail
func NextResendTime(lastSend time.Time, interval *time.Duration) TimeCountdown {
	now := time.Now()

	resendLimitDuration := NullTimeDurationScan(interval, constants.DefaultMailResendTimePeriod)

	timeLimit := lastSend.Add(resendLimitDuration)

	diff := GetRemainingTime(now, timeLimit)

	return diff
}

// GetHourMinuteIntFromTime function to get hour and minute as integer
func GetHourMinuteIntFromTime(t time.Time) (int, *constants.ErrorResponse) {
	result, err := strconv.Atoi(t.Format("1504"))
	if err != nil {
		return 0, ErrorInternalServer(err.Error())
	}

	return result, nil
}

func GetHourAndMinuteFromIntTime(t int) (int, int, *constants.ErrorResponse) {
	var hour int
	var minute int

	stringTime := strconv.Itoa(t)
	hour, err := strconv.Atoi(stringTime[:len(stringTime)-2])
	if err != nil {
		return hour, minute, ErrorInternalServer(err.Error())
	}
	minute, err = strconv.Atoi(stringTime[len(stringTime)-2:])
	if err != nil {
		return hour, minute, ErrorInternalServer(err.Error())
	}

	return hour, minute, nil
}

func SafetyDate(s *time.Time) string {
	if s == nil {
		return ""
	}
	var d time.Time = *s
	return d.Format(time.DateOnly)
}

func SafetyTime(s *time.Time) string {
	if s == nil {
		return ""
	}
	var d time.Time = *s
	return d.Format(constants.HourMinuteOnly)
}

func SafetyDateTime(s *time.Time) string {
	if s == nil {
		return ""
	}
	var d time.Time = *s
	return d.Format(time.RFC3339)
}

func AddTimeToDate(date time.Time, hour int, minute int) (time.Time, *constants.ErrorResponse) {
	loc, err := time.LoadLocation(constants.JakartaTimezone)
	if err != nil {
		return time.Time{}, ErrorInternalServer(err.Error())
	}
	return time.Date(date.Year(), date.Month(), date.Day(), hour, minute, 0, 0, loc), nil
}

func CheckTimeOverlap(startTime1, endTime1, startTime2, endTime2 time.Time) bool {
	if startTime1.Before(endTime2) && startTime2.Before(endTime1) {
		return true
	}

	return false
}

func NameMonth(month int) string {
	var result string
	switch month {
	case 1:
		result = "Januari"
	case 2:
		result = "Februari"
	case 3:
		result = "Maret"
	case 4:
		result = "April"
	case 5:
		result = "Mei"
	case 6:
		result = "Juni"
	case 7:
		result = "Juli"
	case 8:
		result = "Agustus"
	case 9:
		result = "September"
	case 10:
		result = "Oktober"
	case 11:
		result = "November"
	case 12:
		result = "Desember"
	}

	return result
}

func GenerateDateSeries(startDate, endDate time.Time, step string) []time.Time {
	var result []time.Time

	switch step {
	case constants.CalendarGroupByDate:
		for i := 0; startDate.AddDate(0, 0, i).Before(endDate); i++ {
			result = append(result, startDate.AddDate(0, 0, i))
		}
	case constants.CalendarGroupByMonth:
		for i := 0; startDate.AddDate(0, i, 0).Before(endDate); i++ {
			result = append(result, startDate.AddDate(0, i, 0))
		}
	}

	return result
}

func BeginningOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 0, -date.Day()+1)
}

func EndOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 1, -date.Day())
}

func CheckCronMatch(cronExpression string, startDate time.Time) (bool, *constants.ErrorResponse) {
	var result bool
	now := time.Now().Truncate(time.Minute)

	if now.Before(startDate) {
		return result, nil
	}

	schedule, err := cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow).Parse(cronExpression)
	if err != nil {
		return result, ErrorInternalServer(err.Error())
	}

	next := schedule.Next(now.Add(time.Minute * -1))
	for !next.After(now) {
		if next.Equal(now) {
			result = true
			break
		}
		next = schedule.Next(next)
	}

	return result, nil
}
