import { chromium } from 'playwright';

const initBrowser = async () => {
  try {
    return await chromium.launch({ headless: true });
  } catch (error) {
    console.error('Failed to launch browser:', error);
    throw error;
  }
}

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

export { getMatches }
