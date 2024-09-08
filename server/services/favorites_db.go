package services

import "task/system"

type Favorite struct {
	ID      int    `json:"id"`
	Created string `json:"created"`
	Updated string `json:"updated"`
	MalID   string `json:"mal_id"`
	Title   string `json:"title"`
	Image   string `json:"image"`
}

func dest(fav *Favorite) []interface{} {
	return []interface{}{
		&fav.MalID,
		&fav.Created,
		&fav.Updated,
		&fav.Title,
		&fav.Image,
	}
}

func selectAllFavorites() ([]Favorite, error) {
	rows, err := system.Db.Query("select * from favorites")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var favs []Favorite = make([]Favorite, 0)
	for rows.Next() {
		fav := Favorite{}
		err = rows.Scan(dest(&fav)...)
		if err != nil {
			return nil, err
		}
		favs = append(favs, fav)
	}
	return favs, nil
}

func insertFavorite(fav *Favorite) error {
	query := "INSERT INTO favorites (mal_id, title, image) VALUES (?, ?, ?)"
	stmt, err := system.Db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(fav.MalID, fav.Title, fav.Image)
	if err != nil {
		return err
	}

	return nil
}

func deleteFavoriteByMalID(malID int) error {
	query := "DELETE FROM favorites WHERE mal_id = ?"
	stmt, err := system.Db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(malID)
	if err != nil {
		return err
	}

	return nil
}

func malIDExists(malID int) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM favorites WHERE mal_id = ?)"
	err := system.Db.QueryRow(query, malID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func countFavorites() (int, error) {
	var count int
	err := system.Db.QueryRow("SELECT COUNT(*) FROM favorites").Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
