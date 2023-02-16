package model

//"fmt"
//"movie/api/model"

func (movie *Movie) ValidateMovie() (bool , string) {
	if movie.Title == "" {
		return false, "Title is required"
	}
	if movie.ID == 0 {
		return false, "ID is required"
	}
	if movie.Year == 0 {
		return false, "Year is required"
	}
	if movie.Cast == "" {
		return false, "Cast is required"
	}
	if movie.Genre == "" {
		return false, "Genre is required"
	}
	if movie.Language == "" {
		return false, "Language is required"
	}
	if movie.BoxOffice == "" {
		return false, "BoxOffice is required"
	}
	if movie.Rating == "" {
		return false, "Rating is required"
	}
	if movie.Verdict == "" {
		return false, "Verdict is required"
	}
	return true,  "success" 
}
func (cinemaHall *CinemaHall) ValidateCinemaHall() (bool , string) {
	if cinemaHall.Id == 0 {
		return false, "ID is required"
	}
	if cinemaHall.Name == "" {
		return false, "Name is required"
	}
	if cinemaHall.City == "" {
		return false, "city is required"
	}
	if cinemaHall.Movie_Name == "" {
		return false, "movie name is required"
    }
	if cinemaHall.Facilities == "" {
		return false, "facilities is required"
	}
	if cinemaHall.Capacity == 0 {
		return false, "capacity is required"
	}
	if cinemaHall.Parking == "" {
		return false, "parking is required"
	}
	if cinemaHall.Time == "" {
		return false, "time is required"
	}
	if cinemaHall.Type == "" {
		return false, "type is required"
	}
	return true,  "success" 
}




	
