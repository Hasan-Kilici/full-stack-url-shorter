package Database

import(
	"database/sql"
	_ "github.com/lib/pq"
	"urlShorter/Utils"
	"fmt"
)
//Insert
func CreateUrl(OwnerToken, Redirect string) error {
	Token := Utils.GenerateToken(31)
	query := "INSERT INTO urls (Token, OwnerToken, Redirect, Click) VALUES ($1, $2, $3, $4)"
	_, err := db.Exec(query,Token,OwnerToken,Redirect,0)
	if err != nil {
		fmt.Println(err)
		return err
	}
	
	return nil
}

func SaveUserClicked(UrlToken,IP,Country,City,Latitude,Longitude string) error {
	Token := Utils.GenerateToken(31)
	query := "INSERT INTO clicks (Token, UrlToken, IP, Country, City, Latitude, Longitude) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	_, err := db.Exec(query,Token,UrlToken,IP,Country,City,Latitude,Longitude)
	if err != nil {
		fmt.Println(err)
		return err
	}
	
	UpdateClickedCount(UrlToken)
	return nil
}
//Update
func UpdateClickedCount(UrlToken string) error {
	ClickCount := GetClickCount(UrlToken)
	query := "UPDATE urls SET click=$1 WHERE Token=$2"
	_, err := db.Exec(query,ClickCount,UrlToken)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
//Count
func GetClickCount(UrlToken string) int {
	var rowCount int
	query := "SELECT COUNT(*) FROM clicks WHERE UrlToken=$1"
	err := db.QueryRow(query, UrlToken).Scan(&rowCount)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return rowCount
}
//Find
func FindUrl(Token string) (Url, error){
    query := "SELECT * FROM urls WHERE token = $1"
    row := db.QueryRow(query, Token)

    var url Url
    err := row.Scan(&url.Token,&url.OwnerToken,&url.Redirect,&url.Click)
    if err != nil {
        if err == sql.ErrNoRows {
            return Url{}, fmt.Errorf("url bulunamadı")
        }
        return Url{}, err
    }
    return url, nil
}
//List
func ListUrls(OwnerToken string) ([]Url, error){
	query := "SELECT * FROM urls WHERE OwnerToken=$1"
	rows, err := db.Query(query, OwnerToken)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	urls := []Url{}
	for rows.Next() {
		var url Url
		err := rows.Scan(&url.Token,&url.OwnerToken,&url.Redirect,&url.Click)
		if err != nil {
			return nil, err
		}
		urls = append(urls, url)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return urls, nil	
}

func ListClicks(Token string) ([]Clicked, error){
	query := "SELECT * FROM clicks WHERE UrlToken=$1"
	rows, err := db.Query(query, Token)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	clicks := []Clicked{}
	for rows.Next() {
		var click Clicked
		err := rows.Scan(&click.Token,&click.UrlToken,&click.IP,&click.Country,&click.City,&click.Latitude,&click.Longitude)
		if err != nil {
			return nil, err
		}
		clicks = append(clicks, click)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return clicks, nil	
}
//Delete
func DeleteUrl(Token string) error {
	query := "DELETE FROM urls WHERE Token=$1"
	_, err := db.Exec(query,Token)
	if err != nil {
		return err
	}

	DeleteAllClicks(Token)
	return nil
}

func DeleteAllClicks(UrlToken string) error {
	query := "DELETE FROM clicks WHERE UrlToken=$1"
	_, err := db.Exec(query,UrlToken)
	if err != nil {
		return err
	}
	return nil
}