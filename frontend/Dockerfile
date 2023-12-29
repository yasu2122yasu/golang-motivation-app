# M1のみ指定する
FROM --platform=linux/x86_64 node:18.12.1-alpine
WORKDIR /usr/src/app
COPY ./app/package*.json /usr/src/app
RUN npm install
# ここから追加する
COPY . .
CMD ["npm", "start"]
