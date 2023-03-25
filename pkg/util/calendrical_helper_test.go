package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJulianDay(t *testing.T) {
	// Comparison values generated from http://aa.usno.navy.mil/data/docs/JulianDate.php
	testCases := []struct {
		year  int
		month int
		day   int
		want  float64
	}{
		{2010, 1, 2, 2455198.500000},
		{2011, 2, 4, 2455596.500000},
		{2012, 3, 6, 2455992.500000},
		{2013, 4, 8, 2456390.500000},
		{2014, 5, 10, 2456787.500000},
		{2015, 6, 12, 2457185.500000},
		{2016, 7, 14, 2457583.500000},
		{2017, 8, 16, 2457981.500000},
		{2018, 9, 18, 2458379.500000},
		{2019, 10, 20, 2458776.500000},
		{2020, 11, 22, 2459175.500000},
		{2021, 12, 24, 2459572.500000},
	}
	for _, tc := range testCases {
		got := GetJulianDay(tc.year, tc.month, tc.day, 0)
		assert.InDelta(t, tc.want, got, 1e-5)
	}
}

func TestJulianDayWithHoursAndMinutes(t *testing.T) {
	testCases := []struct {
		year    int
		month   int
		day     int
		hours   float64
		minutes float64
		want    float64
	}{
		{2015, 7, 12, 4.25, 0, 2457215.67708333},
		{2015, 7, 12, 4, 15, 2457215.67708333},
		{2015, 7, 12, 8.0, 0, 2457215.833333},
		{1992, 10, 13, 0.0, 0, 2448908.5},
	}
	for _, tc := range testCases {
		got := GetJulianDay(tc.year, tc.month, tc.day, tc.hours+tc.minutes/60.0)
		assert.InDelta(t, tc.want, got, 1e-6)
	}
}

func TestJulianHours(t *testing.T) {
	jdWithoutHours := GetJulianDay(2010, 1, 3, 0)
	jdWithHours := GetJulianDay(2010, 1, 1, 48)

	assert.InDelta(t, jdWithoutHours, jdWithHours, 1e-7)
}
