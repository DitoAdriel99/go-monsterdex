package timeutil

import (
	"log"
	"time"
)

func BeginOfDay(t time.Time) time.Time {
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	return t
}

func BeginOfNextDay(t time.Time) time.Time {
	log.Println(t)
	t = BeginOfDay(t).Add(24 * time.Hour)
	log.Println(t)
	return t
}

func SameDay(t1, t2 time.Time) bool {
	d1, m1, y1 := t1.Date()
	d2, m2, y2 := t2.Date()
	return d1 == d2 && m1 == m2 && y1 == y2
}

func SameWeek(t1, t2 time.Time) bool {
	y1, w1 := t1.ISOWeek()
	y2, w2 := t2.ISOWeek()
	return w1 == w2 && y1 == y2
}

func SameMonth(t1, t2 time.Time) bool {
	y1, m1, _ := t1.Date()
	y2, m2, _ := t2.Date()
	return m1 == m2 && y1 == y2
}

var (
	dayInID = map[int]string{
		0: "Minggu",
		1: "Senin",
		2: "Selasa",
		3: "Rabu",
		4: "Kamis",
		5: "Jum'at",
		6: "Sabtu",
	}
	monthInID = map[int]string{
		1:  "Januari",
		2:  "Februari",
		3:  "Maret",
		4:  "April",
		5:  "Mei",
		6:  "Juni",
		7:  "Juli",
		8:  "Agustus",
		9:  "September",
		10: "Oktober",
		11: "November",
		12: "Desember",
	}
)

func TranslateDay(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return dayInID[int(t.Weekday())]
}

func TransleteMonth(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return monthInID[int(t.Month())]
}
