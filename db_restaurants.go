package main

import "database/sql"

func (p *Restaurant) getRestaurant(db *sql.DB) error {
	return db.QueryRow("SELECT name, description, image_url FROM public.restaurants WHERE id=$1",
		p.Id).Scan(&p.Name, &p.Description, &p.Image_url)
}

func (p *Restaurant) updateRestaurant(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE public.restaurants SET name=$1, description=$2, image_url=$3 WHERE id=$4",
			&p.Name, &p.Description, &p.Image_url)

	return err
}

func (p *Restaurant) deleteRestaurant(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM public.restaurants WHERE id=$1", p.Id)

	return err
}

func (p *Restaurant) createRestaurant(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO public.restaurants(name, description, image_url) VALUES($1, $2, $3) RETURNING id",
		&p.Name, &p.Description, &p.Image_url).Scan(&p.Id)

	if err != nil {
		return err
	}

	return nil
}

func getRestaurants(db *sql.DB, start, count int) ([]Restaurant, error) {
	rows, err := db.Query(
		"SELECT * FROM public.restaurants LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	restaurants := []Restaurant{}

	for rows.Next() {
		var p Restaurant
		if err := rows.Scan(&p.Id, &p.Name, &p.Description, &p.Image_url); err != nil {
			return nil, err
		}
		restaurants = append(restaurants, p)
	}

	return restaurants, nil
}
