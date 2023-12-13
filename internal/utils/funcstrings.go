package utils

import (
	"errors"
	"math"
	"strconv"
	"strings"
	"time"
)

func InsertStr(text string, mask string) (string, error) {
	if text == "" {
		return "", errors.New("value text is empty")
	}
	if mask == "" {
		return "", errors.New("value mask is empty")
	}
	var result string = ""
	x := 0
	for i := 0; i < len(mask); i++ {
		if string(mask[i]) != "#" {
			result += string(mask[i])
		} else {
			result += string(text[x])
			x++
		}
	}
	return result, nil
}

func ValidateData(text string, position int) error {
	if position < 0 {
		return errors.New("position is invalid")
	}
	if text == "" {
		return errors.New("text is invalid or null")
	}
	ind := strings.Index(text, "|")
	if ind == -1 {
		return errors.New("text with format is invalid not localized caracter |")
	}
	if (position) > len(strings.Split(text, "|")) {
		return errors.New("position not found in text")
	}
	return nil
}

func CopyText(text string, position int) (string, error) {
	err := ValidateData(text, position)
	if err != nil {
		return "", err
	}
	return strings.Split(text, "|")[position], nil
}

func CopyTextDate(text string, position int, mask string) (time.Time, error) {
	err := ValidateData(text, position)
	if err != nil {
		return time.Now(), err
	}

	format := "2006-01-02"

	if mask == "" {
		mask = "####-##-##"
	} else if mask == "##-##-####" {
		format = "02-01-2006"
	} else if mask == "##/##/####" {
		format = "02/01/2006"
	}

	arg := strings.Split(text, "|")[position]
	if arg == "" {
		return time.Time{}, nil
	}
	arg, err = InsertStr(arg, mask)
	if err != nil {
		return time.Time{}, err
	}
	d, err := time.Parse(format, arg)
	if err != nil {
		return time.Time{}, err
	}
	return d, nil
}

func CopyTextFloat(text string, position int, decimal int) (float64, error) {
	if decimal < 0 {
		return 0, errors.New("decimal is invalid")
	}
	text = strings.ReplaceAll(text, ".", "")
	if strings.Contains(text, ",") {
		decimal = 0
		text = strings.ReplaceAll(text, ",", ".")
	}

	err := ValidateData(text, position)
	if err != nil {
		return 0, err
	}
	arg := strings.Split(text, "|")[position]
	if arg == "" {
		return 0, nil
	}
	n, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		return 0, err
	}
	return n * float64(1/math.Pow(10, float64(decimal))), nil
}
