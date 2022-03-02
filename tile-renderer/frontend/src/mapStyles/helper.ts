import { FilterSpecification } from 'maplibre-gl';

export const isTunnel: FilterSpecification = ['==', 'tunnel', 1];
export const isNotTunnel: FilterSpecification = ['==', 'tunnel', 0];
export const isBridge: FilterSpecification = ['==', 'bridge', 1];
export const isNotBridge: FilterSpecification = ['==', 'bridge', 0];
