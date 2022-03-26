FROM node:17-bullseye-slim as build-stage
WORKDIR /app
ARG PNPM_VERSION=6.23.6
RUN npm --global install pnpm@${PNPM_VERSION}

COPY pnpm-*.yaml ./
COPY . .
RUN pnpm install
RUN pnpm run build

# production stage
FROM nginx:stable-alpine as production-stage
COPY nginx.conf /etc/nginx/conf.d/default.conf 
COPY --from=build-stage /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
