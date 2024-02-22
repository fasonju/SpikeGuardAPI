package db

import (
	"database/sql"
	"prototype_app_api/models"
)

var DB *sql.DB

func QueryMarkers() ([]models.Marker, error) {
	rows, err := DB.Query("SELECT * FROM markers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var markers []models.Marker

	//loop through rows using Scan to assign column data to struct fields
	for rows.Next() {
		var m models.Marker
		if err := rows.Scan(&m.ID, &m.Latitude, &m.Longitude); err != nil {
			return markers, err
		}
		markers = append(markers, m)
	}

	if err := rows.Err(); err != nil {
		return markers, err
	}

	return markers, nil
}
