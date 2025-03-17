package scraper

import (
    "fmt"
	"regexp"

    "github.com/playwright-community/playwright-go"
)

type Match struct {
    Date          string `json:"date"`
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

type Fixture struct {
	Date 		  string `json:"date"`
    Time          string `json:"time"`
    HomeTeamName  string `json:"homeTeamName"`
    AwayTeamName  string `json:"awayTeamName"`
}

type Ranking struct {
	Rank string `json:"rank"`
	TeamName string `json:"teamName"`
	Played string `json:"played"`
	Win string `json:"win"`
	Draw string `json:"draw"`
	Lost string `json:"lost"`
	ForGoal string `json:"forGoal"`
	AgainstGoal string `json:"againstGoal"`
	GoalDiff string `json:"goalDiff"`
	Points string `json:"points"`
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

func extractTextContents(elements []playwright.Locator) []string {
	texts := []string{}
	for _, el := range elements {
		text, _ := el.TextContent()
		texts = append(texts, text)
	}
	return texts
}

func removeChar(input string) string {
	re := regexp.MustCompile(`[^0-9]`)
	return re.ReplaceAllString(input, "")
}

// === Scraping ===

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
		
		times := extractTextContents(timeEls)
		homeTeamNames := extractTextContents(homeTeamNameEls)
		homeTeamScores := extractTextContents(homeTeamScoreEls)
		awayTeamNames := extractTextContents(awayTeamNameEls)
		awayTeamScores := extractTextContents(awayTeamScoreEls)
		
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

func ScrapePremierLeagueFixtures() ([]Fixture, error) {
	browser, page, err := initPlaywright()
    if err != nil {
        return nil, fmt.Errorf("initialization error: %w", err)
    }
    defer browser.Close()

	results := []Fixture{}

    if _, err := page.Goto("https://www.livescore.com/en/football/england/premier-league/fixtures/"); err != nil {
        return nil, fmt.Errorf("could not visit the desired page: %w", err)
    }

	if _, err := page.WaitForSelector("[data-test-id='virtuoso-item-list']"); err != nil {
        return nil, fmt.Errorf("could not find the item list: %w", err)
    }

	fixturesEls, err := page.Locator("[data-test-id='virtuoso-item-list'] [data-test-id='virtuoso-item-list'] > div").All()
    if err != nil {
        return nil, fmt.Errorf("could not find fixture elements: %w", err)
    }

    for _, fixtureEl := range fixturesEls {
		date, _ := fixtureEl.Locator("[id*='_match-row'] > a > div > span > span:nth-child(1)").TextContent()
		time, _ := fixtureEl.Locator("[data-testid*='_status-or-time']").TextContent()
		homeTeamName, _ := fixtureEl.Locator("[id*='_home-team-name']").TextContent()
		awayTeamName, _ := fixtureEl.Locator("[id*='_away-team-name']").TextContent()
		
		results = append(results, Fixture{
			Date:  date,
			Time:  time,
			HomeTeamName:  homeTeamName,
			AwayTeamName:  awayTeamName,
		})
    }

    return results, nil
}

func ScrapePremierLeagueMatchResults() ([]Match, error) {
	browser, page, err := initPlaywright()
    if err != nil {
        return nil, fmt.Errorf("initialization error: %w", err)
    }
    defer browser.Close()

	results := []Match{}

    if _, err := page.Goto("https://www.livescore.com/en/football/england/premier-league/results/"); err != nil {
        return nil, fmt.Errorf("could not visit the desired page: %w", err)
    }

	if _, err := page.WaitForSelector("[data-test-id='virtuoso-item-list']"); err != nil {
        return nil, fmt.Errorf("could not find the item list: %w", err)
    }

	fixturesEls, err := page.Locator("[data-test-id='virtuoso-item-list'] [data-test-id='virtuoso-item-list'] > div").All()
    if err != nil {
        return nil, fmt.Errorf("could not find fixture elements: %w", err)
    }

    for _, fixtureEl := range fixturesEls {
		date, _ := fixtureEl.Locator("[id*='_match-row'] > a > div > span > span:nth-child(1)").TextContent()
		time, _ := fixtureEl.Locator("[data-testid*='_status-or-time']").TextContent()
		homeTeamName, _ := fixtureEl.Locator("[id*='_home-team-name']").TextContent()
		homeTeamScore, _ := fixtureEl.Locator("[id*='_home-team-score']").TextContent()
		awayTeamName, _ := fixtureEl.Locator("[id*='_away-team-name']").TextContent()
		awayTeamScore, _ := fixtureEl.Locator("[id*='_away-team-score']").TextContent()
		
		results = append(results, Match{
			Date:  date,
			Time:  time,
			HomeTeamName:  homeTeamName,
			HomeTeamScore: homeTeamScore,
			AwayTeamName:  awayTeamName,
			AwayTeamScore: awayTeamScore,
		})
    }

    return results, nil
}

func ScrapePremierLeagueRankings() ([]Ranking, error) {
	browser, page, err := initPlaywright()
    if err != nil {
        return nil, fmt.Errorf("initialization error: %w", err)
    }
    defer browser.Close()

	results := []Ranking{}

    if _, err := page.Goto("https://www.livescore.com/en/football/england/premier-league/table/"); err != nil {
        return nil, fmt.Errorf("could not visit the desired page: %w", err)
    }

	if _, err := page.WaitForSelector("[data-testid*='table-all']"); err != nil {
        return nil, fmt.Errorf("could not find the item list: %w", err)
    }

	rankingsEls, err := page.Locator("[data-testid*='table-all'] tbody tr").All()
    if err != nil {
        return nil, fmt.Errorf("could not find ranking elements: %w", err)
    }

    for _, rankingEl := range rankingsEls {
		rankContent, _ := rankingEl.Locator("td:nth-child(1)").TextContent()
		rank := removeChar(rankContent)
		teamName, _ := rankingEl.Locator("td:nth-child(2)").TextContent()
		played, _ := rankingEl.Locator("td:nth-child(3)").TextContent()
		win, _ := rankingEl.Locator("td:nth-child(4)").TextContent()
		draw, _ := rankingEl.Locator("td:nth-child(5)").TextContent()
		lost, _ := rankingEl.Locator("td:nth-child(6)").TextContent()
		forGoal, _ := rankingEl.Locator("td:nth-child(7)").TextContent()
		againstGoal, _ := rankingEl.Locator("td:nth-child(8)").TextContent()
		goalDiff, _ := rankingEl.Locator("td:nth-child(9)").TextContent()
		points, _ := rankingEl.Locator("td:nth-child(10)").TextContent()
		
		results = append(results, Ranking{
			Rank: rank,
			TeamName: teamName,
			Played: played,
			Win: win,
			Draw: draw,
			Lost: lost,
			ForGoal: forGoal,
			AgainstGoal: againstGoal,
			GoalDiff: goalDiff,
			Points: points,
		})
    }

    return results, nil
}
