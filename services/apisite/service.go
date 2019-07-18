package apisite

import (
	"../../domains/apisite"
	"../../utils"
	"fmt"
	"sync"
)

//func GetSitesFromAPI() ([]apisite.Site, *utils.ApiError) {
//	sites, err := apisite.GetAll()
//	if err != nil {
//		return nil, err
//	}
//
//	var allSites []apisite.Site
//	for _, s := range sites {
//		ss, _ := GetSiteFromAPI(s.ID)
//		if err != nil {
//			return nil, err
//		}
//
//		allSites = append(allSites, *ss)
//	}
//
//	return allSites, nil
//}

func GetSiteFromAPI(siteID string) (*apisite.Site, *utils.ApiError) {
	site := &apisite.Site{
		ID: siteID,
	}

	if err := site.Get(); err != nil {
		return nil, err
	}

	return site, nil
}


func GetUserFromAPI(userID int64) (*apisite.User, *utils.ApiError) {
	user := &apisite.User{
		ID: userID,
	}

	if err := user.Get(); err != nil {
		return nil, err
	}

	return user, nil
}


func GetCountryFromAPI(countryID string) (*apisite.Country, *utils.ApiError) {
	country := &apisite.Country{
		ID: countryID,
	}

	if err := country.Get(); err != nil {
		return nil, err
	}

	return country, nil

}

func getCountry (countryID string, channel chan *apisite.Result) {
	country := &apisite.Country{
		ID: countryID,
	}
	if err := country.Get(); err != nil {
		fmt.Println("ERROR EN COUNTRY")
		//return nil, err
	}
	channel <- &apisite.Result{
		Country:country,
	}

}

func getSite(siteID string, channel chan *apisite.Result)  {
	site := &apisite.Site{
		ID: siteID,
	}
	if err := site.Get(); err != nil {
		fmt.Println("ERROR EN SITE!")
		//return nil, err
	}
	channel <- &apisite.Result{
		Site: site,
	}

}

func GetResultFromAPI(userID int64) (*apisite.Result, *utils.ApiError) {

	channel := make(chan *apisite.Result)
	defer close(channel)
	//errorChannel := make(chan *utils.ApiError)
	//defer close(errorChannel)

	var wg sync.WaitGroup

	/* Obtain the user from user.get() */
	user := &apisite.User{
		ID: userID,
	}
	if err := user.Get(); err != nil {
		return nil, err
	}

	result := &apisite.Result{
		User:user,
	}

	/* Obtain the user from country.get() using the user.countryID */
	wg.Add(1)
	go getCountry(user.CountryID, channel)

	/*Obtain the site from site.get() using the user.SiteID*/
	wg.Add(1)
	go getSite(user.SiteID, channel)

	go func () {
		for res := range channel {
			wg.Done()
			if res.Site != nil {
				result.Site = res.Site
				continue
			}
			if res.Country != nil {
				result.Country = res.Country
				continue
			}
		}
	} ()

	wg.Wait()
	return result, nil
}