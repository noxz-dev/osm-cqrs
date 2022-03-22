import { Client } from '@elastic/elasticsearch';
import { ES_IP } from '../config';
import { logger } from './logger';
export const client = new Client({ node: ES_IP });

try {
  const resp = await client.indices.create(
    {
      index: 'osm',
      settings: {
        max_ngram_diff: 20,
        analysis: {
          filter: {
            ngram_filter: {
              type: 'ngram',
              min_gram: 2,
              max_gram: 20,
            },
          },
          analyzer: {
            ngram_analyzer: {
              type: 'custom',
              tokenizer: 'standard',
              filter: ['lowercase', 'ngram_filter'],
            },
          },
        },
      },
      mappings: {
        properties: {
          name: {
            type: 'text',
            term_vector: 'yes',
            analyzer: 'ngram_analyzer',
            search_analyzer: 'standard',
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
