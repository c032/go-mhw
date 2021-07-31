package mhw

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	largeMonstersURLStr = "https://monsterhunterworld.wiki.fextralife.com/Large+Monsters"
	smallMonstersURLStr = "https://monsterhunterworld.wiki.fextralife.com/Small+Monsters"
	endemicLifeURLStr   = "https://monsterhunterworld.wiki.fextralife.com/Endemic+Life"
	petsURLStr          = "https://monsterhunterworld.wiki.fextralife.com/Pets"

	tooltipURLStrTemplate = "https://monsterhunterworld.wiki.fextralife.com/_tooltip_%s"
)

func Monsters() ([]Creature, error) {
	var (
		err error

		currentURL *url.URL

		req  *http.Request
		resp *http.Response
	)

	currentURL, err = url.Parse(largeMonstersURLStr)
	if err != nil {
		return nil, fmt.Errorf("could not parse url: %w", err)
	}

	req, err = http.NewRequest(http.MethodGet, largeMonstersURLStr, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %w", err)
	}

	c := http.DefaultClient
	resp, err = c.Do(req)
	if err != nil {
	}
	defer resp.Body.Close()

	var doc *goquery.Document
	doc, err = goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not parse response: %w", err)
	}

	containerSelection := doc.Find(`#tagged-pages-container`)

	var monsters []Creature
	containerSelection.Find(`a.wiki_link`).Each(func(i int, link *goquery.Selection) {
		monsterName := link.Text()
		monsterName = strings.TrimSpace(monsterName)
		if monsterName == "" {
			return
		}

		href := link.AttrOr("href", "")
		href = strings.TrimSpace(href)
		if href == "" {
			return
		}

		var (
			err        error
			monsterURL *url.URL
		)

		monsterURL, err = currentURL.Parse(href)
		if err != nil {
			// TODO: Log.

			return
		}

		monsters = append(monsters, &LargeMonster{
			name:   monsterName,
			rawURL: monsterURL.String(),
		})
	})

	return monsters, nil
}
