package model

type Symptoms struct {
	UserId        string `json:"userId" db:"userId"`
	PulseRate     int    `json:"pulseRate" db:"pulseRate"`
	BloodPressure int    `json:"bloodPressure" db:"bloodPressure"`
	OtherSymptoms string `json:"otherSymptoms" db:"otherSymptoms"`
}
