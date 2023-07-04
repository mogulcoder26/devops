FROM node:15
WORKDIR /app
# Optimization Technique
######
COPY package.json .
######
RUN npm install
COPY . ./
EXPOSE 5000
CMD ["npm", "run","dev"]
