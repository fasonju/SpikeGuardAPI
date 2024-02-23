package db

import (
	"context"
	"database/sql"
	"prototype_app_api/models"
	"strings"
	"time"
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

func InsertMarkers(markers []models.Marker) (int64, error) {
	query := "INSERT INTO markers (latitude, longitude) VALUES "
	var inserts []string
	var params []interface{}
	for _, marker := range markers {
		inserts = append(inserts, "(?, ?)")
		params = append(params, marker.Latitude, marker.Longitude)
	}

	queryVals := strings.Join(inserts, ",")
	query = query + queryVals

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	stmt, err := DB.PrepareContext(ctx, query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, params...)
	if err != nil {
		return 0, err
	}

	rowsCount, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsCount, nil
}
