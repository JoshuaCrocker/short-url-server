package main

// Check the database to determine if the given shortcode is already present.
//
// Returns: true if present.
func qryIsShortcodeUnique(shortcode string) bool {
	var url shortURL
	result := db.Where("short_code = ?", shortcode).Find(&url)
	if result.Error != nil {
		panic(result.Error)
	}

	return result.RowsAffected == 0
}
