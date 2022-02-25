import { connect } from 'nats';

export const nc = await connect({ servers: '127.0.0.1' });
