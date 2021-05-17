package main

// Json template: `json:""`

// CREATE TABLE public."restaurants" (
// 	id bigint NOT NULL GENERATED ALWAYS AS IDENTITY,
// 	body varchar(250) NULL,
// 	rate int NOT NULL,
// 	avg_rate int NOT NULL
//  CONSTRAINT restaurants_pk PRIMARY KEY (id);
// );

// Restaurant structures
type Restaurant struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image_url   string `json:"image_url"`
}

// CREATE TABLE public."comments" (
// 	id bigint NOT NULL GENERATED ALWAYS AS IDENTITY,
// 	body varchar(250) NULL,
// 	rate int NOT NULL,
// 	avg_rate int NOT NULL,
// 	restaurant_id bigint NOT NULL,
// 	CONSTRAINT comments_pk PRIMARY KEY (id),
// 	CONSTRAINT comments_fk FOREIGN KEY (id) REFERENCES public.restaurants(id) ON DELETE CASCADE ON UPDATE CASCADE
// );

type Comment struct {
	Id            int    `json:"id"`
	Body          string `json:"body"`
	Rate          int    `json:"rate"`
	Avg_rate      int    `json:"avg_rate"`
	Restaurant_id int    `json:"restaurant_id"`
}
