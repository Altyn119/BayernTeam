package api

type Player struct {
	Id           int    `json:"id"`
	Number       int    `json:"number"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Trnsfermarkt int    `json:"market_value"`
	DateOfBirth  string `json:"date_of_birth"`
	Nation       string `json:"nation"`
	Position     string `json:"position"`
}

var Squad = []Player{
	{1, 1, "Manuel", "Neuer", 5000000, "27-03-1986", "Germany", "GK"},
	{18, 2, "Daniel", "Peretz", 4000000, "10-07-2000", "Germany", "GK"},
	{26, 26, "Sven", "Ulreich", 700000, "03-08-1988", "Germany", "GK"},
	{43, 43, "Tom", "Hülsmann", 250000, "11-04-2004", "Germany", "GK"},
	{4, 4, "Matthijs", "de Ligt", 65000000, "12-08-1999", "Netherlands", "CB"},
	{2, 2, "Dayot", "Upamecano", 60000000, "27-10-1998", "France", "CB"},
	{3, 3, "Min-jae", "Kim", 60000000, "15-11-1996", "South Korea", "CB"},
	{15, 15, "Eric", "Dier", 12000000, "15-01-1994", "England", "CB"},
	{28, 28, "Tarek", "Buchmann", 500000, "28-02-2005", "Germany", "CB"},
	{19, 19, "Alphonso", "Davies", 70000000, "02-11-2000", "Canada", "LB"},
	{22, 22, "Raphaël", "Guerreiro", 15000000, "22-12-1993", "Portugal", "LB"},
	{41, 41, "Frans", "Krätzig", 1500000, "14-01-2003", "Germany", "LB"},
	{6, 45, "Joshua", "Kimmich", 75000000, "08-02-1995", "Germany", "CDM"},
	{45, 45, "Aleksandar", "Pavlovic", 2000000, "03-05-2004", "Germany", "CDM"},
	{8, 8, "Leon", "Goretzka", 40000000, "06-02-1995", "Germany", "CM"},
	{27, 27, "Konrad", "Laimer", 28000000, "27-05-1997", "Austria", "CM"},
	{42, 42, "Jamal", "Musiala", 110000000, "26-02-2003", "Germany", "CAM"},
	{11, 10, "Kingsley", "Coman", 65000000, "13-06-1996", "France", "LW"},
	{39, 39, "Mathys", "Tel", 50000000, "27-04-2005", "France", "LW"},
	{17, 17, "Bryan", "Zaragoza", 12000000, "09-09-2001", "Spain", "LW"},
	{10, 10, "Leroy", "Sané", 80000000, "11-01-1996", "Germany", "RW"},
	{7, 7, "Serge", "Gnabry", 45000000, "14-07-1995", "Germany", "RW"},
	{25, 25, "Thomas", "Müller", 10000000, "13-09-1989", "Germany", "CF"},
	{9, 9, "Harry", "Kane", 110000000, "28-07-1993", "England", "ST"},
	{13, 13, "Eric Maxim", "Choupo-Moting", 3000000, "23-03-1989", "Cameroon", "ST"},
}
