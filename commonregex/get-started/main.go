package main

import (
	cregex "github.com/mingrammer/commonregex"
)

func main() {

	text := `John, please get that article on www.linkedin.com to me by 5:00PM on Jan 9th 2012. 4:00 would be ideal, actually. If you have any questions, You can reach me at (519)-236-2723x341 or get in touch with my associate at harold.smith@gmail.com`

	_ = cregex.Date(text)
	// ['Jan 9th 2012']
	_ = cregex.Time(text)
	// ['5:00PM', '4:00']
	_ = cregex.Links(text)
	// ['www.linkedin.com', 'harold.smith@gmail.com']
	_ = cregex.PhonesWithExts(text)
	// ['(519)-236-2723x341']
	_ = cregex.Emails(text)
	// ['harold.smith@gmail.com']
}
