import { Client } from '@elastic/elasticsearch';
import { ES_IP } from '../config';
export const client = new Client({ node: ES_IP });
