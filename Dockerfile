FROM node:15
WORKDIR /app
# Optimization Technique
######
COPY package.json .
######
RUN npm install
COPY . ./
ENV PORT 5000
EXPOSE $PORT
CMD ["npm", "run","dev"]
