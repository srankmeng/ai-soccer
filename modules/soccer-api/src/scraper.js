import { chromium } from 'playwright';

const initBrowser = async () => {
  try {
    return await chromium.launch({ headless: true });
  } catch (error) {
    console.error('Failed to launch browser:', error);
    throw error;
  }
}

// === Scraping ===

const scrapeMatches = async (browserContext) => {
  const page = await browserContext.newPage();
  const results = [];

  try {
    await page.goto('https://www.livescore.com/en/');
    await page.waitForSelector('[data-test-id="virtuoso-item-list"]');
    const leagueEls = await page.locator('[data-test-id="virtuoso-item-list"] > div').elementHandles();
  
    for (const leagueEl of leagueEls) {
      const leagueName = await leagueEl.$eval('#category-header__stage', (el) => el.textContent?.trim() || '');
      const time = await leagueEl.$$eval('[data-testid*="_status-or-time"]', (els) => els.map(el => el.textContent?.trim() || ''));
      const homeTeamNames = await leagueEl.$$eval('[id*="_home-team-name"]', (els) => els.map(el => el.textContent?.trim() || ''));
      const homeTeamScores = await leagueEl.$$eval('[id*="_home-team-score"]', (els) => els.map(el => el.textContent?.trim() || ''));
      const awayTeamNames = await leagueEl.$$eval('[id*="_away-team-name"]', (els) => els.map(el => el.textContent?.trim() || ''));
      const awayTeamScores = await leagueEl.$$eval('[id*="_away-team-score"]', (els) => els.map(el => el.textContent?.trim() || ''));

      const obj = {
        leagueName,
        matches: homeTeamNames.map((homeTeamName, index) => ({
          time: time[index],
          homeTeamName,
          homeTeamScore: homeTeamScores[index],
          awayTeamName: awayTeamNames[index],
          awayTeamScore: awayTeamScores[index]
        }))
      }
      results.push(obj);
    }
  } catch (error) {
    console.error('Error scraping matches:', error);
    throw error;
  } finally {
    await page.close();
  }
  return results;
}

const scrapePremierLeagueFixtures = async (browserContext) => {
  const page = await browserContext.newPage();
  const results = [];

  try {
    await page.goto('https://www.livescore.com/en/football/england/premier-league/fixtures/');
    await page.waitForSelector('[data-test-id="virtuoso-item-list"]');
    const fixturesEls = await page.locator('[data-test-id="virtuoso-item-list"] [data-test-id="virtuoso-item-list"] > div').elementHandles();
  
    for (const fixtureEl of fixturesEls) {
      const date = await fixtureEl.$eval('[id*="_match-row"] > a > div > span > span', (el) => el.textContent?.trim() || '');
      const time = await fixtureEl.$eval('[data-testid*="_status-or-time"]', (el) => el.textContent?.trim() || '');
      const homeTeamName = await fixtureEl.$eval('[id*="_home-team-name"]', (el) => el.textContent?.trim() || '');
      const awayTeamName = await fixtureEl.$eval('[id*="_away-team-name"]', (el) => el.textContent?.trim() || '');

      const obj = {
        date,
        time,
        homeTeamName,
        awayTeamName,
      }
      results.push(obj);
    }
  } catch (error) {
    console.error('Error scraping Premier League fixtures:', error);
    throw error;
  } finally {
    await page.close();
  }
  return results;
}

const scrapePremierLeagueMatchResults = async (browserContext) => {
  const page = await browserContext.newPage();
  const results = [];

  try {
    await page.goto('https://www.livescore.com/en/football/england/premier-league/results/');
    await page.waitForSelector('[data-test-id="virtuoso-item-list"]');
    const fixturesEls = await page.locator('[data-test-id="virtuoso-item-list"] [data-test-id="virtuoso-item-list"] > div').elementHandles();
  
    for (const fixtureEl of fixturesEls) {
      const date = await fixtureEl.$eval('[id*="_match-row"] > a > div > span > span', (el) => el.textContent?.trim() || '');
      const time = await fixtureEl.$eval('[data-testid*="_status-or-time"]', (el) => el.textContent?.trim() || '');
      const homeTeamName = await fixtureEl.$eval('[id*="_home-team-name"]', (el) => el.textContent?.trim() || '');
      const homeTeamScore = await fixtureEl.$eval('[id*="_home-team-score"]', (el) => el.textContent?.trim() || '');
      const awayTeamName = await fixtureEl.$eval('[id*="_away-team-name"]', (el) => el.textContent?.trim() || '');
      const awayTeamScore = await fixtureEl.$eval('[id*="_away-team-score"]', (el) => el.textContent?.trim() || '');

      const obj = {
        date,
        time,
        homeTeamName,
        homeTeamScore,
        awayTeamName,
        awayTeamScore,
      }
      results.push(obj);
    }
  } catch (error) {
    console.error('Error scraping Premier League match results:', error);
    throw error;
  } finally {
    await page.close();
  }
  return results;
}

const scrapePremierLeagueRankings = async (browserContext) => {
  const page = await browserContext.newPage();
  const results = [];

  try {
    await page.goto('https://www.livescore.com/en/football/england/premier-league/table/');
    await page.waitForSelector('[data-testid*="table-all"]');
    const rankingsEls = await page.locator('[data-testid*="table-all"] tbody tr').elementHandles();
  
    for (const rankingEl of rankingsEls) {
      const rank = await rankingEl.$eval('td:nth-child(1)', (el) => el.textContent?.replace(/[^\d]/g, '') || '');
      const teamName = await rankingEl.$eval('td:nth-child(2)', (el) => el.textContent?.trim() || '');
      const played = await rankingEl.$eval('td:nth-child(3)', (el) => el.textContent?.trim() || '');
      const win = await rankingEl.$eval('td:nth-child(4)', (el) => el.textContent?.trim() || '');
      const draw = await rankingEl.$eval('td:nth-child(5)', (el) => el.textContent?.trim() || '');
      const lost = await rankingEl.$eval('td:nth-child(6)', (el) => el.textContent?.trim() || '');
      const forGoal = await rankingEl.$eval('td:nth-child(7)', (el) => el.textContent?.trim() || '');
      const againstGoal = await rankingEl.$eval('td:nth-child(8)', (el) => el.textContent?.trim() || '');
      const goalDiff = await rankingEl.$eval('td:nth-child(9)', (el) => el.textContent?.trim() || '');
      const points = await rankingEl.$eval('td:nth-child(10)', (el) => el.textContent?.trim() || '');
      
      const obj = {
        rank,
        teamName,
        played,
        win,
        draw,
        lost,
        forGoal,
        againstGoal,
        goalDiff,
        points,
      }
      results.push(obj);
    }
  } catch (error) {
    console.error('Error scraping Premier League rankings:', error);
    throw error;
  } finally {
    await page.close();
  }
  return results;
}

// === Get data from scraping ===

const getMatches = async () => {
  let browser = null;
  
  try {
    browser = await initBrowser();
    const context = await browser.newContext();
    const results = await scrapeMatches(context);
    console.log('Scraped matches successfully');
    return results;
  } catch (error) {
    console.error('getMatches failed:', error);
    throw error;
  } finally {
    if (browser) {
      await browser.close();
    }
  }
}

const getPremierLeagueFixtures = async () => {
  let browser = null;
  
  try {
    browser = await initBrowser();
    const context = await browser.newContext();
    const results = await scrapePremierLeagueFixtures(context);
    console.log('Scraped Premier League fixtures successfully');
    return results;
  } catch (error) {
    console.error('getPremierLeagueFixtures failed:', error);
    throw error;
  } finally {
    if (browser) {
      await browser.close();
    }
  }
}

const getPremierLeagueMatchResults = async () => {
  let browser = null;
  
  try {
    browser = await initBrowser();
    const context = await browser.newContext();
    const results = await scrapePremierLeagueMatchResults(context);
    console.log('Scraped Premier League match results successfully');
    return results;
  } catch (error) {
    console.error('getPremierLeagueMatchResults failed:', error);
    throw error;
  } finally {
    if (browser) {
      await browser.close();
    }
  }
}

const getPremierLeagueRankings = async () => {
  let browser = null;
  
  try {
    browser = await initBrowser();
    const context = await browser.newContext();
    const results = await scrapePremierLeagueRankings(context);
    console.log('Scraped Premier League rankings successfully');
    return results;
  } catch (error) {
    console.error('getPremierLeagueRankings failed:', error);
    throw error;
  } finally {
    if (browser) {
      await browser.close();
    }
  }
}

export {
  getMatches,
  getPremierLeagueFixtures,
  getPremierLeagueMatchResults,
  getPremierLeagueRankings,
}
