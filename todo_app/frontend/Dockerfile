FROM node:22-alpine3.23 AS dependencies
WORKDIR /app
COPY package.json package-lock.json ./
RUN npm install


FROM node:22-alpine3.23 AS builder
WORKDIR /app
COPY --from=dependencies /app/node_modules ./node_modules
COPY . .
ENV NODE_ENV=production
RUN npm run build


FROM node:22-alpine3.23
WORKDIR /app
ENV HOSTNAME="0.0.0.0"

RUN mkdir .next
RUN chown node:node .next

COPY --from=builder --chown=node:node /app/.next/standalone ./
COPY --from=builder --chown=node:node /app/.next/static ./.next/static

USER node

CMD [ "node", "server.js" ]