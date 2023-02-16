package model

import "net/http"

type Movie struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Year      int    `json:"year"`
	Cast      string `json:"cast"`
	Genre     string `json:"genre"`
	Language  string `json:"language"`
	BoxOffice string `json:"box_office"`
	Rating    string `json:"rating"`
	Verdict   string `json:"verdict"`
}

var Movies = []Movie{
	{
		ID:        1,
		Title:     "dilwale",
		Year:      2017,
		Cast:      "shahrukh_khan,kajol,varun_dhawan",
		Genre:     "drama,romance,comedy",
		Language:  "Hindi",
		BoxOffice: "200cr",
		Rating:    "4.5",
		Verdict:   "Hit",
	},
	{
		ID:        2,
		Title:     "PATHAN",
		Year:      2023,
		Cast:      "jhon_abraham,deepika_padukone,shahrukh_khan",
		Genre:     "crime,drama,thriller,action",
		Language:  "hindi",
		BoxOffice: "300cr",
		Rating:    "4.8",
		Verdict:   "all_time_blockbuster",
	},
	{
		ID:        3,
		Title:     " no_time_to_die",
		Year:      2021,
		Cast:      "daniel_craig,rami_malek,ana_de_armas",
		Genre:     "crime,drama,thriller,action",
		Language:  "English",
		BoxOffice: "400cr",
		Rating:    "4.6",
		Verdict:   "hit",
	},
	{
		ID:        4,
		Title:     "tiger_zinda_hai",
		Year:      2008,
		Cast:      "salman_khan,katrina_kaif",
		Genre:     "action,crime,drama",
		Language:  "hindi",
		BoxOffice: "350cr",
		Rating:    "4.7",
		Verdict:   "blockbuster",
	},
	{
		ID:        5,
		Title:     "the_dark_knight",
		Year:      2008,
		Cast:      "christian_bale,heath_ledger,aaron_eckhart",
		Genre:     "action,crime,drama",
		Language:  "english",
		BoxOffice: "650cr",
		Rating:    "4.8",
		Verdict:   "all_time_blockbuster",
	},
	{
		ID:        6,
		Title:     "theri",
		Year:      2016,
		Cast:      "vijay,amy_jackson,samantha_akkineni",
		Genre:     "action,crime",
		Language:  "hindi",
		BoxOffice: "250cr",
		Rating:    "4.3",
		Verdict:   "hit",
	},
	{
		ID:        7,
		Title:     "radhe",
		Year:      2017,
		Cast:      "salman_khan,disha_patani,randeep_hooda",
		Genre:     "comedy",
		Language:  "hindi",
		BoxOffice: "350cr",
		Rating:    "4.5",
		Verdict:   "flop",
	},
	{
		ID:        8,
		Title:     "lal_singh_chaddha",
		Year:      2021,
		Cast:      "aamir_khan,kareena_kapoor,gippy_grewal",
		Genre:     "romance",
		Language:  "hindi",
		BoxOffice: "350cr",
		Rating:    "4.2",
		Verdict:   "flop",
	},
	{
		ID:        9,
		Title:     "kgf",
		Year:      2019,
		Cast:      "yash,srinidhi_shetty,prakash_rao",
		Genre:     "action,crime,drama",
		Language:  "hindi",
		BoxOffice: "550cr",
		Rating:    "4.9",
		Verdict:   "all_time_blockbuster",
	},
	{
		ID:        10,
		Title:     "pushpa",
		Year:      2020,
		Cast:      " allu_arjun,rashmika_mandanna,prakash_rao",
		Genre:     "action,crime,drama",
		Language:  "hindi",
		BoxOffice: "450cr",
		Rating:    "5",
		Verdict:  "super_hit",
	},
}

func AttachResult(obj interface{}, err error) interface{} {
	if err != nil {
		return map[string]interface{}{
			"message": "Failed",
			"code":    http.StatusPreconditionRequired,
		}
	} else {
		return map[string]interface{}{
			"message": "Ok",
			"code":    http.StatusOK,
		}
	}

}
