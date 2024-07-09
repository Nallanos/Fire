FROM node:16
WORKDIR /app
COPY package.json package-lock.json ./
RUN npm ci
COPY . .
RUN npm run build
ENV PORT 5173
EXPOSE 5173
CMD ["node", "build/index.js"]
