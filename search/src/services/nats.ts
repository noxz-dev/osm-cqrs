import { NATS_IP } from '../config';
import { connect } from 'nats';

export const nc = await connect({ servers: NATS_IP });
