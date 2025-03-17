# soccer-api

This project was created using go. install <https://go.dev/doc/install> first.

## How it works

Using [Playwright](https://playwright.dev/) browser testing framework for scraping the matches on [livescore.com](https://www.livescore.com/) website. Then extract the matches data to json format, and return it.

## How to run

To run:

```bash
make run
```

The application will run on `http://localhost:3000`.

## API paths

- `/health` - returns `200 OK` if the application is running
- `/matches` - returns a list of current matches
- `/premier-league/fixtures` - returns a list of next fixtures for the Premier League
- `/premier-league/match-results` - returns a list of latest match results for the Premier League
- `/premier-league/rankings` - returns a list of current team ranking table for the Premier League
