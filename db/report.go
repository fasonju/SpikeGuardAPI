package db

import (
	"context"
	"prototype_app_api/models"
	"time"
)

func InsertReport(request models.ReportPutRequest) error {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	stmt, err := DB.PrepareContext(ctx,
		`INSERT INTO unvalidated_reports (
                                 victim_name,
                                 longitude,
                                 latitude,
                                 emergency_contact_phone,
                                 victim_phone,
                                 victim_country,
                                 victim_address,
                                 city,
                                 victim_last_name,
                                 time) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, request.VictimName,
		request.Longitude,
		request.Latitude,
		request.EmergencyContactPhone,
		request.VictimPhone,
		request.VictimCountry,
		request.VictimAddress,
		request.VictimCity,
		request.VictimLastName,
		request.ReportDateTime)
	return err

}
