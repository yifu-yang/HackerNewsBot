/*
HackerNews response data struct
*/
package HackerNewsBot

type User struct {
	About     string `json:"about"`
	Created   int    `json:"created"`
	Delay     int    `json:"delay"`
	ID        string `json:"id"`
	Karma     int    `json:"karma"`
	Submitted []int  `json:"submitted"`
}

type Stories []int

type ChangedItems struct {
	Items    []int    `json:"items"`
	Profiles []string `json:"profiles"`
}

type Item struct {
	ID          int    `json:"id"`
	Deleted     bool   `json:"deleted"`
	Type        string `json:"type"`
	By          string `json:"by"`
	Time        int    `json:"time"`
	Text        string `json:"text"`
	Dead        bool   `json:"dead"`
	Descendants int    `json:"descendants"`
	Kids        []int  `json:"kids"`
	Score       int    `json:"score"`
	Title       string `json:"title"`
	URL         string `json:"url"`
	Parent      int    `json:"parent"`
	Parts       []int  `json:"parts"`
	Poll        int    `json:"poll"`
}
