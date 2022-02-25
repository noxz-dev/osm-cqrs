FROM node:17-bullseye-slim

# Create app directory
WORKDIR /usr/src/app
ARG PNPM_VERSION=6.23.6
RUN npm --global install pnpm@${PNPM_VERSION}

# Install app dependencies
COPY pnpm-*.yaml ./
COPY . .
RUN pnpm install
RUN pnpm run build

EXPOSE 6000:6000

CMD ["node", "dist/app.mjs"]