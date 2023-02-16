package model

import (
   // "fmt"
   // "os"
)

type CinemaHall struct {
	Id            int `json:"id"`
    Name          string `json:"name"`
	City		  string `json:"city"`
	Movie_Name 	   string `json:"movie_name"`
	Facilities    string `json:"facilities"`
	Parking	string `json:"parking"`
	Time string `json:"time"`
	Type        string `json:"type"`
	Capacity   int `json:"capacity"`
}
var CinemaHalls = []CinemaHall{

	{
        Id:            1,
		Name:          "inox_cinema",
		City:          "new_delhi",
        Movie_Name:         "pathan",
		Facilities:    "dolby_digital",
        Parking:       "yes",
		Time:          "3:00pm",
        Type:          "gold",
		Capacity:       100,
},
    {
        Id:            2,
		Name:          "fun_cinema",
		City:          "up",
		Movie_Name:    "dhoom",
		Facilities:    "3d",
		Parking:       "yes",
		Time:          "4:00pm",
		Type:          "prime",
        Capacity:       300,
},
    {
        Id:            3,
		Name:          "pvr_cinema",
		City:          "noida",
		Movie_Name:         "pathan",
		Facilities:    "4d",
		Parking:       "yes",
		Time:          "8:00pm",
		Type:          "gold",
        Capacity:       200,
},
    {
        Id:            4,
		Name:          "scooter_cinema",
		City:          "noida",
		Movie_Name:         "shershah",
		Facilities:    "dolby_digital,",
		Parking:       "no",
		Time:          "7:00pm",
		Type:          "classic",
        Capacity:       80,
},	
	{
		Id:            5,
        Name:          "cinepolis_cinema",
        City:          "gurgaon",
        Movie_Name:         "shershah",
		Facilities:    ",4d",
		Parking:       "no",
		Time:          "7:00pm",
		Type:          "classic",
        Capacity:      70,
},
	}