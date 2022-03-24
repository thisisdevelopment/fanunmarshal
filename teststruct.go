package fanunmarshal

type TestData []TestObj

type TestObj struct {
	ID            string   `json:"_id"`
	Index         int64    `json:"index"`
	GUID          string   `json:"guid"`
	IsActive      bool     `json:"isActive"`
	Balance       string   `json:"balance"`
	Picture       string   `json:"picture"`
	Age           int64    `json:"age"`
	EyeColor      string   `json:"eyeColor"`
	Name          string   `json:"name"`
	Gender        string   `json:"gender"`
	Company       string   `json:"company"`
	Email         string   `json:"email"`
	Phone         string   `json:"phone"`
	Address       string   `json:"address"`
	About         string   `json:"about"`
	Registered    string   `json:"registered"`
	Latitude      float64  `json:"latitude"`
	Longitude     float64  `json:"longitude"`
	Tags          []string `json:"tags"`
	Friends       []Friend `json:"friends"`
	Greeting      string   `json:"greeting"`
	FavoriteFruit string   `json:"favoriteFruit"`
}

type Friend struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
