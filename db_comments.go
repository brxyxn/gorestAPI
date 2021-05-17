package main

import (
	"database/sql"
)

// p is equal to post
// p.$key from each struct

// Get Comments by $id
// This could be avoid, since every post will get all comments
// And we wouldn't have any method to show 1 comment only.

// This will require the restaurant_id in order to get
// this value from db
// func (c *Comment) getComment(db *sql.DB) error {
// 	return db.QueryRow("SELECT * FROM public.restaurants WHERE id=$1",
// 		c.Id).Scan(&c.Body, &c.Rate, &c.Rate_avg, c.Restaurant_id)
// }

func (c *Comment) updateComment(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE public.comments SET name=$1, rate=$2, rate_avg=$3 WHERE id=$4",
			c.Body, c.Rate, c.Rate_avg, c.Id)

	return err
}

func (c *Comment) deleteComment(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM public.comments WHERE id=$1", c.Id)

	return err
}

func (c *Comment) createComment(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO public.comments(body, rate, rate_avg, restaurant_id) VALUES($1, $2, $3, $4) RETURNING id",
		c.Body, c.Rate, c.Rate_avg, c.Restaurant_id).Scan(&c.Id)

	if err != nil {
		return err
	}

	return nil
}

// Getting all comments, passing params to handle pagination
// db = connection string, id=restaurant_id, start, count = how many items this method will gather.
func getComments(db *sql.DB, id, start, count int) ([]Comment, error) {
	rows, err := db.Query("SELECT * FROM public.comments WHERE restaurant_id=$1 LIMIT $2 OFFSET $3",
		id, count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	comments := []Comment{}

	for rows.Next() {
		var c Comment
		if err := rows.Scan(&c.Id, &c.Body, &c.Rate, &c.Rate_avg, &c.Restaurant_id); err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}

	return comments, nil
}
