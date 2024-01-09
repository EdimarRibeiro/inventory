package tests

import (
	"errors"
	"strconv"
	"testing"
	"time"

	utils "github.com/EdimarRibeiro/inventory/internal/utils"

	"github.com/stretchr/testify/assert"
)

func TestValidValidateDataNoError(t *testing.T) {
	err := utils.ValidateData("|test|", 3)
	assert.NoError(t, err, nil)
}

func TestValidInsertStrNoErrorValue(t *testing.T) {
	var err error = nil
	var value string = ""
	value, err = utils.InsertStr("20231231", "####-##-##")
	assert.NoError(t, err, nil)
	if value != "2023-12-31" {
		assert.NoError(t, errors.New("value date not is iguals 2023-12-31"), nil)
	}
}

func TestValidCopyTextNoError(t *testing.T) {
	var err error = nil
	var value string = ""
	value, err = utils.CopyText("|test|", 1)
	assert.NoError(t, err, nil)
	if value != "test" {
		err = errors.New("value not is test invalid")
	}
	assert.NoError(t, err, nil)
}
func TestValidCopyTextNoErrorValue(t *testing.T) {
	var err error = nil
	var value string = ""
	value, err = utils.CopyText("|test|", 0)
	assert.NoError(t, err, nil)
	if value != "" {
		err = errors.New("value not is empty (" + value + ")")
	}
	assert.NoError(t, err, nil)
}
func TestValidCopyTextError(t *testing.T) {
	var err error = nil
	var value string = ""
	value, err = utils.CopyText("", 1)
	assert.Error(t, err, value)
}

func TestValidCopyTextDateNoErrorValue(t *testing.T) {
	var err error = nil
	var value time.Time
	d, err := time.Parse("2006-01-02", "2023-12-31")
	assert.NoError(t, err, nil)

	value, err = utils.CopyTextDate("|20231231|", 1, "####-##-##")
	assert.NoError(t, err, nil)
	if !value.Equal(d) {
		assert.NoError(t, errors.New("copyTextDate parse value date not is iguals 2023-12-31"), nil)
	}
}

func TestValidCopyTextFloatNoErrorValue(t *testing.T) {
	var err error = nil
	var value float64 = 0
	value, err = utils.CopyTextFloat("|1000|", 1, 2)
	assert.NoError(t, err, nil)
	if value != 10 {
		err = errors.New("value not is empty (" + strconv.FormatFloat(value, 'g', -1, 64) + ")")
	}
	assert.NoError(t, err, nil)
}
