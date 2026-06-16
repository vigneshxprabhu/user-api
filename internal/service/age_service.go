package service

import (
	"time"
)

func CalculateAge(dob string) (int, error) {

	birthDate, err := time.Parse("2006-01-02", dob)

	if err != nil {
		return 0, err
	}

	now := time.Now()

	age := now.Year() - birthDate.Year()

	if now.YearDay() < birthDate.YearDay() {
		age--
	}

	return age, nil
}
