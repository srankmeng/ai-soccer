# soccer-api

This project was created using `bun init` in bun v1.2.2. [Bun](https://bun.sh) is a fast all-in-one JavaScript runtime.

## Concepts

Using [Playwright](https://playwright.dev/) browser testing framework for scraping the matches on [livescore.com](https://www.livescore.com/) website. Then extract the matches data to json format, and return it.

## How to run

Install dependencies:

```bash
bun install
```

Install Playwright browsers:

```bash
bunx playwright install chromium
```

To run:

```bash
make run
```

The application will run on `http://localhost:3000`.
