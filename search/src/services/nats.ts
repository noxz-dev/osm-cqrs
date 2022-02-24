import { connect } from 'nats';
import { DEFAULT_HOST } from 'nats/lib/nats-base-client/types';

export const nc = await connect({ servers: DEFAULT_HOST });
