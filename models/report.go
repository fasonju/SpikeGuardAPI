package models

type Report struct {
	ID                    int     `json:"id"`
	VictimName            string  `json:"victim_name"`
	VictimLastName        string  `json:"victim_last_name"`
	VictimPhone           string  `json:"victim_phone"`
	Longitude             float64 `json:"longitude"`
	Latitude              float64 `json:"latitude"`
	ReportDateTime        string  `json:"report_datetime"`
	EmergencyContactPhone string  `json:"emergency_contact_phone"`
	VictimCountry         string  `json:"victim_country"`
	VictimCity            string  `json:"victim_city"`
	VictimAddress         string  `json:"victim_address"`
}

type ReportPutRequest struct {
	VictimName            string  `json:"victim_name"`
	VictimLastName        string  `json:"victim_last_name"`
	VictimPhone           string  `json:"victim_phone"`
	Longitude             float64 `json:"longitude"`
	Latitude              float64 `json:"latitude"`
	ReportDateTime        string  `json:"report_datetime"`
	EmergencyContactPhone string  `json:"emergency_contact_phone"`
	VictimCountry         string  `json:"victim_country"`
	VictimCity            string  `json:"victim_city"`
	VictimAddress         string  `json:"victim_address"`
}
