package webservice

type UINames struct {
	Name      string    `json:"name"`
	Surname  string `json:"surname"`
	Gender string `json:"gender"`
	Region string `json:"region"`
}

type Joke struct{
	Type string `json:"type"`
	Val  Value   `json:"value"`  

}
type Value struct {
	Id int `json:"id"`
	Joke string `json:"joke"`
	Categories []string `json:"categories"`


}
