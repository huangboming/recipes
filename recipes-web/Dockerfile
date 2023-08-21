FROM node

WORKDIR /recipes-web

# install dependencies
COPY package.json /recipes-web
RUN npm install

COPY . /recipes-web

# start app
EXPOSE 3000
CMD ["npm", "start"]

