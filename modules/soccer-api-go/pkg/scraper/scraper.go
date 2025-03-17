package scraper

import (
    "fmt"

    "github.com/playwright-community/playwright-go"
)

type Match struct {
    Time          string `json:"time"`
    HomeTeamName  string `json:"homeTeamName"`
    HomeTeamScore string `json:"homeTeamScore"`
    AwayTeamName  string `json:"awayTeamName"`
    AwayTeamScore string `json:"awayTeamScore"`
}

type League struct {
    LeagueName string  `json:"leagueName"`
    Matches    []Match `json:"matches"`
}

func initPlaywright() (playwright.Browser, playwright.Page, error) {
	pw, err := playwright.Run()
	if err != nil {
		return nil, nil, fmt.Errorf("could not start Playwright: %w", err)
	}

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
        Headless: playwright.Bool(true),
    })
    if err != nil {
        return nil, nil, fmt.Errorf("could not launch the browser: %w", err)
    }

	page, err := browser.NewPage()
    if err != nil {
        return nil, nil, fmt.Errorf("could not open a new page: %w", err)
    }

	return browser, page, nil
}

func ScrapeMatches() ([]League, error) {
	browser, page, err := initPlaywright()
    if err != nil {
        return nil, fmt.Errorf("initialization error: %w", err)
    }
    defer browser.Close()

	results := []League{}

    if _, err := page.Goto("https://www.livescore.com/en/"); err != nil {
        return nil, fmt.Errorf("could not visit the desired page: %w", err)
    }

	if _, err := page.WaitForSelector("[data-test-id='virtuoso-item-list']"); err != nil {
        return nil, fmt.Errorf("could not find the item list: %w", err)
    }

	leagueEls, err := page.Locator("[data-test-id='virtuoso-item-list'] > div").All()
    if err != nil {
        return nil, fmt.Errorf("could not find league elements: %w", err)
    }

    for _, leagueEl := range leagueEls {
		leagueName, _ := leagueEl.Locator("#category-header__stage").TextContent()
		timeEls, _ := leagueEl.Locator("[data-testid*='_status-or-time']").All()
		homeTeamNameEls, _ := leagueEl.Locator("[id*='_home-team-name']").All()
		homeTeamScoreEls, _ := leagueEl.Locator("[id*='_home-team-score']").All()
		awayTeamNameEls, _ := leagueEl.Locator("[id*='_away-team-name']").All()
		awayTeamScoreEls, _ := leagueEl.Locator("[id*='_away-team-score']").All()
		
		times := []string{}
		for _, timeEl := range timeEls {
			time, _ := timeEl.TextContent()
			times = append(times, time)
		}
		homeTeamNames := []string{}
		for _, homeTeamNameEl := range homeTeamNameEls {
			homeTeamName, _ := homeTeamNameEl.TextContent()
			homeTeamNames = append(homeTeamNames, homeTeamName)
		}
		homeTeamScores := []string{}
		for _, homeTeamScoreEl := range homeTeamScoreEls {
			homeTeamScore, _ := homeTeamScoreEl.TextContent()
			homeTeamScores = append(homeTeamScores, homeTeamScore)
		}
		awayTeamNames := []string{}
		for _, awayTeamNameEl := range awayTeamNameEls {
			awayTeamName, _ := awayTeamNameEl.TextContent()
			awayTeamNames = append(awayTeamNames, awayTeamName)
		}
		awayTeamScores := []string{}
		for _, awayTeamScoreEl := range awayTeamScoreEls {
			awayTeamScore, _ := awayTeamScoreEl.TextContent()
			awayTeamScores = append(awayTeamScores, awayTeamScore)
		}
		matches := []Match{}
        for i := range homeTeamNames {
            matches = append(matches, Match{
                Time:          times[i],
                HomeTeamName:  homeTeamNames[i],
                HomeTeamScore: homeTeamScores[i],
                AwayTeamName:  awayTeamNames[i],
                AwayTeamScore: awayTeamScores[i],
            })
        }

        results = append(results, League{
            LeagueName: leagueName,
            Matches:    matches,
        })
    }

    return results, nil
}
