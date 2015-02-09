package tmdb

import (
	"fmt"
)

// AccountInfo struct
type AccountInfo struct {
	ID           int
	IncludeAdult bool   `json:"include_adult"`
	Iso3166_1    string `json:"iso_3166_1"`
	Iso639_1     string `json:"iso_639_1"`
	Name         string
	Username     string
}

// AccountPagedMovieResults struct
type AccountPagedMovieResults struct {
	Page         int
	Results      []MovieShort
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

// AccountPagedTvResults struct
type AccountPagedTvResults struct {
	Page         int
	Results      []TvShort
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

// GetAccountInfo gets the basic information for an account
// http://docs.themoviedb.apiary.io/#reference/account/account/get
func (tmdb *TMDb) GetAccountInfo(sessionID string) (*AccountInfo, error) {
	var account AccountInfo
	uri := fmt.Sprintf("%s/account?api_key=%s&session_id=%s", baseURL, tmdb.apiKey, sessionID)
	result, err := getTmdb(uri, &account)
	return result.(*AccountInfo), err
}

// GetAccountLists gets the lists that you have created and marked as a favorite
// http://docs.themoviedb.apiary.io/#reference/account/accountidlists/get
func (tmdb *TMDb) GetAccountLists(id int, sessionID string, options map[string]string) (*MovieLists, error) {
	var availableOptions = map[string]struct{}{
		"page":     {},
		"language": {}}
	var lists MovieLists
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/account/%v/lists?api_key=%s&session_id=%s%s", baseURL, id, tmdb.apiKey, sessionID, optionsString)
	result, err := getTmdb(uri, &lists)
	json, _ := ToJSON(result)
	fmt.Printf("%v\n%v\n", uri, json)
	return result.(*MovieLists), err
}

// GetAccountFavoriteMovies gets the list of favorite movies for an account
// http://docs.themoviedb.apiary.io/#reference/account/accountidfavoritemovies/get
func (tmdb *TMDb) GetAccountFavoriteMovies(id int, sessionID string, options map[string]string) (*AccountPagedMovieResults, error) {
	var availableOptions = map[string]struct{}{
		"page":     {},
		"sort_by":  {},
		"language": {}}
	var favorites AccountPagedMovieResults
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/account/%v/favorite/movies?api_key=%s&session_id=%s%s", baseURL, id, tmdb.apiKey, sessionID, optionsString)
	result, err := getTmdb(uri, &favorites)
	json, _ := ToJSON(result)
	fmt.Printf("%v\n%v\n", uri, json)
	return result.(*AccountPagedMovieResults), err
}

// GetAccountFavoriteTv gets the list of favorite movies for an account
// http://docs.themoviedb.apiary.io/#reference/account/accountidfavoritetv/get
func (tmdb *TMDb) GetAccountFavoriteTv(id int, sessionID string, options map[string]string) (*AccountPagedTvResults, error) {
	var availableOptions = map[string]struct{}{
		"page":     {},
		"sort_by":  {},
		"language": {}}
	var favorites AccountPagedTvResults
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/account/%v/favorite/tv?api_key=%s&session_id=%s%s", baseURL, id, tmdb.apiKey, sessionID, optionsString)
	result, err := getTmdb(uri, &favorites)
	json, _ := ToJSON(result)
	fmt.Printf("%v\n%v\n", uri, json)
	return result.(*AccountPagedTvResults), err
}

// GetAccountRatedMovies gets the list of rated movies (and associated rating) for an account
// http://docs.themoviedb.apiary.io/#reference/account/accountidratedmovies/get
func (tmdb *TMDb) GetAccountRatedMovies(id int, sessionID string, options map[string]string) (*AccountPagedMovieResults, error) {
	var availableOptions = map[string]struct{}{
		"page":     {},
		"sort_by":  {},
		"language": {}}
	var favorites AccountPagedMovieResults
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/account/%v/rated/movies?api_key=%s&session_id=%s%s", baseURL, id, tmdb.apiKey, sessionID, optionsString)
	result, err := getTmdb(uri, &favorites)
	json, _ := ToJSON(result)
	fmt.Printf("%v\n%v\n", uri, json)
	return result.(*AccountPagedMovieResults), err
}

// GetAccountRatedTv gets the list of rated TV shows (and associated rating) for an account
// http://docs.themoviedb.apiary.io/#reference/account/accountidratedtv/get
func (tmdb *TMDb) GetAccountRatedTv(id int, sessionID string, options map[string]string) (*AccountPagedTvResults, error) {
	var availableOptions = map[string]struct{}{
		"page":     {},
		"sort_by":  {},
		"language": {}}
	var favorites AccountPagedTvResults
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/account/%v/rated/tv?api_key=%s&session_id=%s%s", baseURL, id, tmdb.apiKey, sessionID, optionsString)
	result, err := getTmdb(uri, &favorites)
	json, _ := ToJSON(result)
	fmt.Printf("%v\n%v\n", uri, json)
	return result.(*AccountPagedTvResults), err
}

// GetAccountWatchlistMovies gets the list of movies on an accounts watchlist
// http://docs.themoviedb.apiary.io/#reference/account/accountidwatchlistmovies/get
func (tmdb *TMDb) GetAccountWatchlistMovies(id int, sessionID string, options map[string]string) (*AccountPagedMovieResults, error) {
	var availableOptions = map[string]struct{}{
		"page":     {},
		"sort_by":  {},
		"language": {}}
	var favorites AccountPagedMovieResults
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/account/%v/watchlist/movies?api_key=%s&session_id=%s%s", baseURL, id, tmdb.apiKey, sessionID, optionsString)
	result, err := getTmdb(uri, &favorites)
	json, _ := ToJSON(result)
	fmt.Printf("%v\n%v\n", uri, json)
	return result.(*AccountPagedMovieResults), err
}

// GetAccountWatchlistTv gets the list of TV series on an accounts watchlist
// http://docs.themoviedb.apiary.io/#reference/account/accountidwatchlisttv/get
func (tmdb *TMDb) GetAccountWatchlistTv(id int, sessionID string, options map[string]string) (*AccountPagedTvResults, error) {
	var availableOptions = map[string]struct{}{
		"page":     {},
		"sort_by":  {},
		"language": {}}
	var favorites AccountPagedTvResults
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/account/%v/watchlist/tv?api_key=%s&session_id=%s%s", baseURL, id, tmdb.apiKey, sessionID, optionsString)
	result, err := getTmdb(uri, &favorites)
	json, _ := ToJSON(result)
	fmt.Printf("%v\n%v\n", uri, json)
	return result.(*AccountPagedTvResults), err
}