# frontend/Dockerfile
FROM node:20-alpine AS builder

WORKDIR /app

# Копіюємо тільки package.json і package-lock.json для кешування залежностей
COPY package*.json ./

RUN npm install

# Копіюємо весь код фронтенду
COPY . .

# Створюємо production build
RUN npm run build

# Stage для nginx
FROM nginx:alpine
COPY --from=builder /app/build /usr/share/nginx/html

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]