package model

import (
	"testing"
)

func TestWelfareDonateMarshal(t *testing.T) {
	wdEntry := new(WelfareDonate)
	wdEntry.ClaimId = "48pVXCQk37Rzf9mMWbEboaaoQKJ3"
	wdEntry.Priority = 1.5
	wdEntry.Amount = 10
	bs, err := WelfareDonateMarshal(*wdEntry)
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Log(string(bs))
}

func TestWelfareDonateUnmarshal(t *testing.T) {
	wdEntry := new(WelfareDonate)
	wdEntry.ClaimId = "48pVXCQk37Rzf9mMWbEboaaoQKJ3"
	wdEntry.Priority = 1.5
	wdEntry.Amount = 10
	bs, err := WelfareDonateMarshal(*wdEntry)
	if err != nil {
		t.Error(err.Error())
		return
	}

	wdEntry1 := new(WelfareDonate)
	err = WelfareDonateUnmarshal(string(bs), wdEntry1)
	if err != nil {
		t.Error(err.Error())
		return
	}
}
