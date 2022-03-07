import { Client } from '@elastic/elasticsearch';
import { ES_IP } from '../config';
import { logger } from './logger';
export const client = new Client({ node: ES_IP });

try {
  const resp = await client.indices.create(
    {
      index: 'osm',
      mappings: {
        properties: {
          name: {
            type: 'text',
          },
          location: {
            type: 'geo_point',
          },
        },
      },
    },
    { ignore: [400] }
  );

  logger.info(resp);
} catch (err) {
  logger.error(err);
}
