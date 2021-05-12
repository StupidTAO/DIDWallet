package model

import (
	"encoding/json"
	"strconv"
)

type WelfareDonate struct {
	ClaimId	string
	Amount int32
	Priority float32
}

func WelfareDonateUnmarshal(WelfareDonateStr string, welfateDonate *WelfareDonate) error {
	return json.Unmarshal([]byte(WelfareDonateStr), welfateDonate)
}

func WelfareDonateMarshal(welfateDonate WelfareDonate) ([]byte, error)  {
	return json.Marshal(welfateDonate)
}

func GetWelfareDonate(ClaimId string, Amount string, Priority string) (WelfareDonate, error) {
	amount, err := strconv.Atoi(Amount)
	if err != nil {
		return WelfareDonate{}, err
	}
	priority, err :=strconv.ParseFloat(Priority, 32/64)
	wd := WelfareDonate{ClaimId, int32(amount), float32(priority)}
	return wd, nil
}
