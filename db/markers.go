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
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	rows, err := DB.QueryContext(ctx, "SELECT id, latitude, longitude FROM markers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var markers []models.Marker
	for rows.Next() {
		var marker models.Marker
		if err := rows.Scan(&marker.ID, &marker.Latitude, &marker.Longitude); err != nil {
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

func DeleteMarker(id int) error {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	stmt, err := DB.PrepareContext(ctx, "DELETE FROM markers WHERE id=?"); if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id);
	return err
}
