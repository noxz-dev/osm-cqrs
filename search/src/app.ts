import express, { Request } from 'express';
import 'dotenv/config';
import { PORT } from './config';
import { logger } from './services/logger';
import { client } from './services/es';
import { nc } from './services/nats';
import { StringCodec } from 'nats';
import { QueryDslQueryContainer } from '@elastic/elasticsearch/lib/api/types';

const app = express();
const sc = StringCodec();

subscribeToEvents();

app.use(express.json());

app.listen(PORT, () => {
  logger.info(`search backend is running http://localhost:${PORT}`);
});

interface SearchPoint {
  Name: string;
  Id: string;
  Location: {
    Lat: number;
    Lng: number;
  };
  Tags?: any[];
}

//just for testing purposes, production will run on events
app.post('/addData', async (req, res) => {
  await demoInsert();
  res.status(200).send();
});

async function subscribeToEvents() {
  const sub = nc.subscribe('search');
  for await (const m of sub) {
    const startTime = performance.now();

    const data = sc.decode(m.data) as any;
    const event = JSON.parse(data);

    if (!event.data) return;

    try {
      const modify = event.data.Modify as SearchPoint[];
      logger.info(`modify payload size: ${modify.length} elements`);

      await bulkInsertDocument(modify);
    } catch (err) {
      logger.error(err);
    }

    try {
      const create = event.data.Create as SearchPoint[];
      logger.info(`create payload size: ${create.length} elements`);

      await bulkInsertDocument(create);
    } catch (err) {
      logger.error(err);
    }

    try {
      const remove = event.data.Delete as SearchPoint[];
      logger.info(`delete payload size: ${remove.length} elements`);

      for await (const loc of remove) {
        logger.info('removing: ', loc.Id);
        await removeDocument(loc);
      }
    } catch (err) {
      logger.error(err);
    }

    const endTime = performance.now();
    logger.info(
      `processing the full search payload took ${(endTime - startTime).toFixed(
        3
      )} ms`
    );

    await client.indices.refresh({ index: 'osm' });
  }
}

async function insertDocument(sp: SearchPoint) {
  try {
    await client.update({
      index: 'osm',
      id: sp.Id,
      doc: {
        name: sp.Name,
        location: {
          lat: sp.Location.Lat,
          lon: sp.Location.Lng,
        },
        tags: sp.Tags || [],
      },
      doc_as_upsert: true,
    });
  } catch (err: any) {
    logger.error(err);
  }
}

async function bulkInsertDocument(points: SearchPoint[]) {
  try {
    const body = points.flatMap((doc) => [
      { update: { _index: 'osm', _id: doc.Id } },
      {
        doc: {
          name: doc.Name,
          location: {
            lat: doc.Location.Lat,
            lon: doc.Location.Lng,
          },
          tags: doc.Tags || [],
        },
        doc_as_upsert: true,
      },
    ]);

    if (body.length === 0) return;

    await client.bulk({ refresh: true, operations: body });
  } catch (err: any) {
    logger.error(err);
  }
}

async function removeDocument(sp: SearchPoint) {
  try {
    const startTime = performance.now();

    await client.delete({
      index: 'osm',
      id: sp.Id,
    });
    const endTime = performance.now();
    logger.info(`delete document took ${(endTime - startTime).toFixed(3)} ms`);
  } catch (err: any) {
    logger.error(err);
  }
}

async function demoInsert() {
  await client.update({
    index: 'osm',
    id: '1234',
    doc: {
      name: 'Hochschule Hannover',
      tags: [],
      location: {
        lat: 52.353683,
        lon: 9.72422,
      },
    },
    doc_as_upsert: true,
  });

  await client.indices.refresh({ index: 'osm' });
}

interface SearchByDistanceQuery {
  lat: number;
  lon: number;
  distance: string;
}

app.get(
  '/searchByDistance',
  async (req: Request<{}, {}, {}, SearchByDistanceQuery>, res) => {
    const { lat = 52.353683, lon = 9.72422, distance = '1000km' } = req.query;

    if (!lat || !lon) {
      res.status(400).send('no location specified');
    }

    if (!distance) {
      res.status(400).send('no distance specified');
    }

    const result = await client.search({
      index: 'osm',
      query: {
        bool: {
          filter: {
            geo_distance: {
              distance: distance,
              location: {
                lat,
                lon,
              },
            },
          },
        },
      },
    });

    res.send(result.hits.hits);
  }
);

app.get('/all', async (req, res) => {
  try {
    const result = await client.search({
      index: 'osm',
      query: {
        match_all: {},
      },
      size: 1000,
    });

    res.send(result.hits.hits);
  } catch (err) {
    logger.error(err);
    res.status(500).send();
  }
});

app.get('/searchByName', async (req, res) => {
  try {
    const { name } = req.query;
    if (!name) {
      res.status(400).send('no query specified');
    }
    const result = await client.search({
      index: 'osm',
      query: {
        match: {
          name: {
            query: name as string,
            fuzziness: 'AUTO',
            operator: 'AND',
          },
        },
      },
    });

    res.send(result.hits.hits);
  } catch (err) {
    res.status(400).send();
  }
});

interface AddressQuery {
  city: string;
  housenumber: string;
  street: string;
  amenity: string;
}

app.get(
  '/searchByAddress',
  async (req: Request<{}, {}, {}, AddressQuery>, res) => {
    try {
      const { city, housenumber, street, amenity } = req.query;

      const query = [];

      if (city && city.length > 1) {
        query.push({ match: { 'tags.K': 'addr:city' } });
        query.push({
          match: {
            'tags.V': {
              query: city,
              operator: 'AND',
            },
          },
        });
      }

      if (housenumber && housenumber.length > 1) {
        query.push({ match: { 'tags.K': 'addr:housenumber' } });
        query.push({ match: { 'tags.V': housenumber } });
      }

      if (street && street.length > 1) {
        query.push({ match: { 'tags.K': 'addr:street' } });
        query.push({
          match: {
            'tags.V': {
              query: street,
              operator: 'AND',
            },
          },
        });
      }

      if (amenity && amenity.length > 1) {
        query.push({ match: { 'tags.K': 'amenity' } });
        query.push({ match: { 'tags.V': amenity } });
      }

      if (query.length === 0) {
        res.send([]);
        return;
      }

      const result = await client.search({
        index: 'osm',
        query: {
          bool: {
            must: query as QueryDslQueryContainer[],
          },
        },
      });

      res.send(result.hits.hits);
    } catch (err) {
      logger.error(err);
      res.status(400).send();
    }
  }
);

app.get('/rawQuery', async (req, res) => {
  try {
    const { query } = req.body;
    if (!query) {
      res.status(400).send('no query specified');
    }
    const result = await client.search(query);
    res.send(result);
  } catch (err) {
    logger.error(err);
    res.status(400).send('raw query failed');
  }
});

app.get('/count', async (req, res) => {
  try {
    const result = await client.count({
      query: {
        match_all: {},
      },
    });
    res.send(result);
  } catch (err) {
    logger.error(err);
    res.status(400).send('count query failed');
  }
});
