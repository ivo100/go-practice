package chrono

import (
	"log/slog"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	YMDLayout         = "2006-01-02"
	YearMonthLayout   = "2006-01"
	HourMinOnlyLayout = "15:04"
	HourMinSecLayout  = "15:04:05.00"
	ShortLayout       = "2006-01-02 15:04:05"
)

const NewYorkLocation = "America/New_York"

var NewYork *time.Location

func init() {
	var err error
	NewYork, err = time.LoadLocation(NewYorkLocation)
	if err != nil {
		panic("Error loading location " + err.Error())
	}
}

func ZeroTime() time.Time {
	return time.Unix(0, 0)
}

func ToYMD(tm time.Time) string {
	return tm.Format(YMDLayout)
}

func ToYearMonth(tm time.Time) string {
	return tm.Format(YearMonthLayout)
}

func FromYMDHm(year, month, day, hour, min int) time.Time {
	return time.Date(year, time.Month(month), day, hour, min, 0, 0, NewYork)
}
func FromYMD(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, NewYork)
}

func FromMDstr(dt string) time.Time {
	a := strings.Split(dt, "-")
	if len(a) != 2 {
		return time.Time{}
	}
	m := Int(a[0])
	d := Int(a[1])
	year := time.Now().Year()
	return FromYMD(year, m, d)
}

func FromMD(month, day int) time.Time {
	year := time.Now().Year()
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, NewYork)
}

func FromHourMin(hour, min int) time.Time {
	now := time.Now()
	tm := time.Date(now.Year(), now.Month(), now.Day(), hour, min, 0, 0, NewYork)
	return tm
}

func FromYM(year, month int) time.Time {
	return time.Date(year, time.Month(month), 1, 0, 0, 0, 0, NewYork)
}

func FromYMDstr(dt string) time.Time {
	if len(dt) == 6 {
		y := 2000 + Int(dt[0:2])
		m := Int(dt[2:4])
		d := Int(dt[4:])
		return FromYMD(y, m, d)
	}
	a := strings.Split(dt, "-")
	if len(a) != 3 {
		return time.Time{}
	}
	y := Int(a[0])
	if y < 100 {
		y += 2000
	}
	m := Int(a[1])
	d := Int(a[2])
	return FromYMD(y, m, d)
}

func FromStr(dt string) time.Time {
	a := strings.Split(dt, "-")
	if len(a) == 3 {
		y := Int(a[0])
		m := Int(a[1])
		d := Int(a[2])
		return FromYMDHm(y, m, d, 0, 0)
	}
	now := time.Now()
	if len(a) == 2 {
		y := now.Year()
		m := Int(a[0])
		d := Int(a[1])
		return FromYMDHm(y, m, d, 0, 0)
	}
	y := now.Year()
	m := int(now.Month())
	d := Int(a[1])
	return FromYMDHm(y, m, d, 0, 0)
}

func FromDateTime(date, hourMin string) time.Time {
	Assert(len(date) >= 10, "invalid date "+date)
	Assert(len(hourMin) == 5, "invalid time "+hourMin)
	a := strings.Split(date, "-")
	Assert(len(a) == 3, "invalid date "+date)
	y := Int(a[0])
	m := Int(a[1])
	d := Int(a[2])

	a = strings.Split(hourMin, ":")
	Assert(len(a) == 2, "invalid time "+hourMin)
	H := Int(a[0])
	M := Int(a[1])
	return FromYMDHm(y, m, d, H, M)
}

func ParseIsoDate(date string) (time.Time, error) {
	return time.ParseInLocation(time.RFC3339, date, NewYork)
}

func ParseDate(date string) (time.Time, error) {
	if len(date) == 10 {
		date += " 00:00:00"
	}
	return time.ParseInLocation(ShortLayout, date, NewYork)
}

func MustParseDate(date string) (tm time.Time) {
	if len(date) == 10 {
		date += " 00:00:00"
	}
	tm, err := time.ParseInLocation(ShortLayout, date, NewYork)
	if err != nil {
		slog.Error("*** Invalid date " + date)
	}
	return tm
}

func MustParseIsoDate(date string) (tm time.Time) {
	tm, err := time.ParseInLocation(time.RFC3339, date, NewYork)
	if err != nil {
		slog.Error("*** Invalid date " + date)
	}
	return tm
}

// DS returns date only string in EST zone
func DS(dt time.Time) string {
	t := dt.In(NewYork).Format(YMDLayout)
	return t
}

// DT returns date+time string in EST zone
func DT(dt time.Time) string {
	t := dt.In(NewYork).Format(time.RFC3339)[0:19]
	t = strings.ReplaceAll(t, "T", " ")
	return t
}

// TM returns time only string in EST zone
func TM(dt time.Time) string {
	t := dt.In(NewYork).Format(HourMinOnlyLayout)
	return t
}

func TMS(dt time.Time) string {
	t := dt.In(NewYork).Format(HourMinSecLayout)
	return t
}

func IsoDateTimeStr(dt time.Time) string {
	t := dt.In(NewYork).Format(time.RFC3339)
	//t = strings.ReplaceAll(t, "T", " ")
	return t
}

func YearMonthOnlyStr(dt time.Time) string {
	t := dt.In(NewYork).Format(YearMonthLayout)
	return t
}

func ActualNYtime() time.Time {
	t := time.Now().In(NewYork)
	return t
}

// DayDiff calculates the number of days between two given time.Created values (a and b).
// It does this by subtracting b from a, converting the resulting duration to hours,
// and then dividing by 24 to get the number of days.
// The result is returned as an integer.
func DayDiff(a, b time.Time) int {
	return int(a.Sub(b).Hours() / 24)
}

// RandomDelay sleeps for a random delay between min and max milliseconds
func RandomDelay(min, max int) {
	time.Sleep(time.Duration(min+rand.Intn(max)) * time.Millisecond)
}

func CombineDateTime(dt time.Time, t time.Time) time.Time {
	return time.Date(dt.Year(), dt.Month(), dt.Day(), t.Hour(), t.Minute(), t.Second(), 0, NewYork)
}

func ToMarketTime(dt time.Time) time.Time {
	return dt.In(NewYork)
}

func UnixToMarketTime(ts int64) time.Time {
	return time.Unix(ts, 0).In(NewYork)
}

func UnixToHMstr(ts int64) string {
	st := TM(UnixToMarketTime(ts))
	return st[0:5]
}

func RelativeDate(dt string, offsetDays int) string {
	d := FromYMDstr(dt)
	d2 := d.AddDate(0, 0, offsetDays)
	s := d2.Format(YMDLayout)
	return s
}

func TradingDayStart(y, m, d int) time.Time {
	return FromYMDHm(y, m, d, 9, 30)
}

func Int(s string) int {
	Assert(s != "", "invalid int")
	n, err := strconv.Atoi(s)
	Assert(err == nil, "invalid int "+s)
	return n
}

func Assert(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}

// 2024-02-12 09:29:00

func YMD(dt string) string {
	Assert(len(dt) >= 10, "expected date YYYY-MM-DD...")
	return dt[0:10]
}

func HM(dt string) string {
	Assert(len(dt) > 15, "expected date YYYY-MM-DD HH:MM...")
	return dt[11:16]
}
