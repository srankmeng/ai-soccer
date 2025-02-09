import {
  getMatches,
  getPremierLeagueFixtures,
  getPremierLeagueMatchResults,
  getPremierLeagueRankings,
} from './scraper';

const corsHeaders = {
  'Access-Control-Allow-Origin': '*',
  'Access-Control-Allow-Methods': 'GET, OPTIONS',
  'Access-Control-Allow-Headers': 'Content-Type',
};

const returnResponse = (name, respData) => {
  try {
    return new Response(JSON.stringify(respData), {
      headers: {
        'Content-Type': 'application/json',
        ...corsHeaders
      }
    });
  } catch (error) {
    return new Response(
      JSON.stringify({ error: `Failed to fetch ${name}` }), {
        status: 500,
        headers: {
          'Content-Type': 'application/json',
          ...corsHeaders
        }
      }
    );
  }
}

const server = Bun.serve({
  port: 3000,
  async fetch(req) {
    const url = new URL(req.url);
    const headers = {
      'Content-Type': 'application/json',
      ...corsHeaders
    };

    if (req.method === 'OPTIONS') {
      return new Response(null, {
        headers: corsHeaders
      });
    }

    if (url.pathname === '/health') {
      return new Response(JSON.stringify({ status: 'ok' }), {
        headers,
      });
    }

    if (url.pathname === '/matches') {
      const matches = await getMatches();
      return returnResponse('matches', matches);
    }

    if (url.pathname === '/premier-league/fixtures') {
      const fixtures = await getPremierLeagueFixtures();
      return returnResponse('fixtures', fixtures);
    }

    if (url.pathname === '/premier-league/match-results') {
      const results = await getPremierLeagueMatchResults();
      return returnResponse('results', results);
    }

    if (url.pathname === '/premier-league/rankings') {
      const results = await getPremierLeagueRankings();
      return returnResponse('rankings', results);
    }

    return new Response(JSON.stringify({ error: 'Not Found' }), {
      status: 404,
      headers,
    });
  },
});

console.log(`Soccer API server listening on http://localhost:${server.port}`);