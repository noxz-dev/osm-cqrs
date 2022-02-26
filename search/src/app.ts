import express from 'express';
import 'dotenv/config';
import { PORT } from './config';
import { logger } from './services/logger';
import { client } from './services/es';
import { nc } from './services/nats';
import { StringCodec } from 'nats';

const app = express();
const sc = StringCodec();

subscribeToEvents();

app.use(express.json());

app.listen(PORT, () => {
  logger.info(`search backend is running http://localhost:${PORT}`);
});

//just for testing purposes, production will run on events
app.post('/addData', async (req, res) => {
  await demoInsert();
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

async function subscribeToEvents() {
  const sub = nc.subscribe('search');
  for await (const m of sub) {
    const data = sc.decode(m.data);
    console.log(data);
    // insertDocument(data);
  }
}

interface SearchPoint {
  Name: string;
  Id: string;
  Location: {
    Lat: number;
    Lng: number;
  };
  Tags?: any[];
}

async function insertDocument(sp: SearchPoint) {
  await client.index({
    index: 'osm',
    document: {
      name: sp.Name,
      location: {
        lat: sp.Location.Lat,
        lon: sp.Location.Lng,
      },
      tags: sp.Tags,
    },
  });

  await client.indices.refresh({ index: 'osm' });
}

async function demoInsert() {
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
}
