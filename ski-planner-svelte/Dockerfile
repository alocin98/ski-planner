FROM node:18-alpine AS builder

WORKDIR /app
COPY package*.json .
RUN npm ci
COPY . .
ENV BUILDING="true"
ENV PUBLIC_FIREBASE_API_KEY="AIzaSyAewZwjaN3VvarDkmCs_N-5VRnK3dNh_qQ"
ENV PUBLIC_FIREBASE_AUTHDOMAIN="skiyeti-io.firebaseapp.com"
ENV PUBLIC_FIREBASE_PROJECTID="skiyeti-io"
ENV PUBLIC_FIREBASE_STORAGEBUCKET="skiyeti-io.appspot.com"
ENV PUBLIC_FIREBASE_MESSAGINGSENDERID="587406996488"
ENV PUBLIC_FIREBASE_APPID="1:587406996488:web:36545f98c2d1c06f35cdbd"
ENV PUBLIC_FIREBASE_MEASUREMENTID="G-V9X2GS7KVB"
RUN npm run build
RUN npm prune --production

FROM node:18-alpine
WORKDIR /app
COPY --from=builder /app/build build/
COPY --from=builder /app/node_modules node_modules/
COPY package.json .
EXPOSE 3000
ENV NODE_ENV=production
CMD [ "node", "build/index.js" ]
