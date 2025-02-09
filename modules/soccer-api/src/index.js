import { getMatches  } from './scraper';

const server = Bun.serve({
  port: 3000,
  async fetch(req) {
    
    const corsHeaders = {
      'Access-Control-Allow-Origin': '*',
      'Access-Control-Allow-Methods': 'GET, OPTIONS',
      'Access-Control-Allow-Headers': 'Content-Type',
    };
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
      try {
        const matches = await getMatches();
        
        // // Handle query parameters for filtering
        // const competition = url.searchParams.get('competition');
        // const filteredMatches = competition
        //   ? matches.filter(match => 
        //       match.competition.toLowerCase().includes(competition.toLowerCase())
        //     )
        //   : matches;

        return new Response(JSON.stringify(matches), {
          headers: {
            'Content-Type': 'application/json',
            ...corsHeaders
          }
        });
      } catch (error) {
        return new Response(
          JSON.stringify({ error: 'Failed to fetch matches' }), {
            status: 500,
            headers: {
              'Content-Type': 'application/json',
              ...corsHeaders
            }
          }
        );
      }
    }

    return new Response(JSON.stringify({ error: 'Not Found' }), {
      status: 404,
      headers,
    });
  },
});

console.log(`Soccer API server listening on http://localhost:${server.port}`);