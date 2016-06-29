package main

import (
	"log"

	"github.com/PuerkitoBio/goquery"
)

func check(url string, prevInfo CheckInfo) (CheckInfo, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
		return prevInfo, err
	}

	curInfo := CheckInfo{
		OldUpdateDate: prevInfo.OldUpdateDate,
		NewUpdateDate: prevInfo.NewUpdateDate,
		IsUpdated:     prevInfo.IsUpdated,
	}

	doc.Find("div[itemprop=\"datePublished\"]").Each(func(_ int, s *goquery.Selection) {
		log.Println(s.Text())
		if prevInfo.OldUpdateDate == "" {
			curInfo.OldUpdateDate = s.Text()
		} else {
			curInfo.NewUpdateDate = s.Text()
			if curInfo.OldUpdateDate != curInfo.NewUpdateDate {
				curInfo.IsUpdated = true
			}
		}
	})

	return curInfo, nil
}
