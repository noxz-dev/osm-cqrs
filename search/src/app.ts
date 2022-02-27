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

app.get('/searchByDistance', async (req, res) => {
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

app.get('/searchAll', async (req, res) => {
  const result = await client.search({
    index: 'osm',
    query: {
      match_all: {},
    },
  });

  res.send(result.hits.hits);
});

app.get('/searchByName', async (req, res) => {
  const { name } = req.query;

  if (!name) {
    res.status(400).send('no name specified');
    return;
  }

  const result = await client.search({
    index: 'osm',
    query: {
      match: {
        name: name as string,
      },
    },
  });

  res.send(result.hits.hits);
});

async function subscribeToEvents() {
  const sub = nc.subscribe('search');
  for await (const m of sub) {
    const data = sc.decode(m.data) as any;
    const event = JSON.parse(data);

    const modify = event.data.Modify as SearchPoint[];
    const create = event.data.Create as SearchPoint[];
    const remove = event.data.Delete as SearchPoint[];

    for await (const loc of modify) {
      console.log('inserting: ', loc);
      await insertDocument(loc);
    }

    for await (const loc of create) {
      console.log('inserting: ', loc);
      await insertDocument(loc);
    }

    for await (const loc of remove) {
      console.log('removing: ', loc);
      await removeDocument(loc);
    }

    await client.indices.refresh({ index: 'osm' });
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
}

async function removeDocument(sp: SearchPoint) {
  await client.delete({
    index: 'osm',
    id: sp.Id,
  });
}

async function demoInsert() {
  // await client.index({
  //   index: 'osm',
  //   id: '1234',
  //   document: {
  //     name: 'Hochschule Hannover',
  //     location: {
  //       lat: 52.353683,
  //       lon: 9.72422,
  //     },
  //     tags: [],
  //   },
  // });

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
