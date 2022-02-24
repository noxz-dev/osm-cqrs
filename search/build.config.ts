export default {
  entries: ['./src/app'],
  rollup: {
    esbuild: {
      target: ['esnext'],
    },
  },
};
