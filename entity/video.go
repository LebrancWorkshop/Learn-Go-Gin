package entity 

type Video struct {
	Title 				string	`json:"title"`
	Uploader 			string	`json:"uploader"`
	Description 	string	`json:"description"`
	URL 					string	`json:"url"`
}