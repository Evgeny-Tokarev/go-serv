package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "recordings"
)

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var db *sql.DB

func dbInit() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	var err error
	db, err = sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	//var albums []Album
	//albums, err = getAlbumsByArtist("John Coltrane")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("Albums found: %v\n", albums)
	//var alb Album
	//alb, err = getAlbumByID(7)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("Album found: %v\n", alb)
	//var albID int64
	//albID, err = addAlbum(Album{
	//	Title:  "The Modern Sound of Betty Carter",
	//	Artist: "Betty Carter",
	//	Price:  49.99,
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("ID of added album: %v\n", albID)
	//var idToDelete int64 = 7
	//if removeAlbumById(idToDelete) {
	//	fmt.Printf("Successfully deleted 1 row with ID %d", idToDelete)
	//}

}

func GetAlbumsByArtist(name string) ([]Album, error) {
	dbInit()
	var albums []Album
	rows, err := db.Query("SELECT * FROM album WHERE artist = $1", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()

	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}
func GetAlbumByTitle(title string) (Album, error) {
	dbInit()
	var alb Album

	row := db.QueryRow("SELECT * FROM album WHERE title = $1", title)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %v: no such album", title)
		}
		return alb, fmt.Errorf("albumsById %v: %v", title, err)
	}
	return alb, nil
}

func AddAlbum(alb Album) (int64, error) {
	dbInit()
	var id int64
	if err := db.QueryRow("INSERT INTO album (title, artist, price) VALUES ($1, $2, $3) RETURNING id", alb.Title, alb.Artist, alb.Price).Scan(&id); err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}

	return id, nil
}

func RemoveAlbumById(id int64) bool {
	dbInit()
	var err error = nil
	var res sql.Result
	res, err = db.Exec("DELETE FROM album WHERE id = $1", id)

	if err == nil {
		count, err1 := res.RowsAffected()
		if err1 == nil {
			if count == 1 {
				return true
			}
		}

	}
	return false
}
