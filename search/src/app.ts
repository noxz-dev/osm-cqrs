import express from 'express';
import 'dotenv/config';
import { PORT } from './config';
import { logger } from 'services/logger';
import { client } from 'services/es';

const app = express();

app.use(express.json());

app.listen(PORT, () => {
  logger.info(`search backend is running http://localhost:${PORT}`);
});

app.post('/addData', async (req, res) => {
  await client.index({
    index: 'osm',
    document: {
      name: 'Hochschule Hannover',
      location: {
        lat: 52.353683,
        lon: 9.72422,
      },
    },
  });

  await client.indices.refresh({ index: 'osm' });

  res.status(200).send();
});

app.get('/search', async (req, res) => {
  const result = await client.search({
    index: 'osm',
    query: {
      bool: {
        filter: {
          geo_distance: {
            distance: '1000km',
            location: {
              lat: 52.398425,
              lon: 9.725097,
            },
          },
        },
      },
    },
  });

  res.send(result.hits.hits);
});
