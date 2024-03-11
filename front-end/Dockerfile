# First stage: build the front-end application
FROM node:latest AS builder

WORKDIR /app

COPY package.json package-lock.json ./

RUN npm install

COPY . .

RUN npm run build

# Second stage: deploy the application with NGINX
FROM nginx:latest

# Copy the built files from the previous stage into NGINX's html directory
COPY --from=builder /app/dist /usr/share/nginx/html

# Update the NGINX configuration to listen on port 8081
RUN sed -i -e 's/80/8081/g' /etc/nginx/conf.d/default.conf

# Expose port 8081
EXPOSE 8081

# NGINX is set to automatically run upon container startup, so no CMD or ENTRYPOINT is needed
