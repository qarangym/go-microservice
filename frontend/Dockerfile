# Use an official Node runtime as a parent image
FROM node:16 AS build

# Set the working directory
WORKDIR /app

# Copy package.json and yarn.lock to the working directory
COPY package.json yarn.lock ./

# Install dependencies
RUN yarn install

# Copy the rest of the application to the working directory
COPY . .

# Build the application
RUN yarn build

######## Start a new stage from scratch #######
# Use an official nginx image as the base image
FROM nginx:alpine

# Copy the build output to the nginx html directory
COPY --from=build /app/build /usr/share/nginx/html

# Expose port 80
EXPOSE 80

# Command to run nginx
CMD ["nginx", "-g", "daemon off;"]
