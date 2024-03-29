package db

import (
	"context"
	"prototype_app_api/models"
	"strings"
	"time"
)

func QueryMarkers() ([]models.Marker, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := DB.QueryContext(ctx, "SELECT id, latitude, longitude FROM markers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var markers []models.Marker
	for rows.Next() {
		var marker models.Marker
		if err := rows.Scan(&marker.ID,
			&marker.Latitude,
			&marker.Longitude); err != nil {
			return nil, err
		}
		markers = append(markers, marker)
	}
	if err := rows.Err(); err != nil {
		return nil, err
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

	queryVars := strings.Join(inserts, ",")
	query = query + queryVars

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

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

func DeleteMarker(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err := DB.PrepareContext(ctx, "DELETE FROM markers WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	return err
}

func InsertMarker(latitude float64, longitude float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err := DB.PrepareContext(ctx,
		"INSERT INTO markers (latitude, longitude) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, latitude, longitude)
	return err
}
