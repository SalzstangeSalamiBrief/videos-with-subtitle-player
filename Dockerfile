FROM node:20-slim AS node_base
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable


ARG VITE_BASE_URL
ARG ROOT_PATH
ARG ALLOWED_CORS
ARG FRONTEND_PORT

# Set environment variables during the build process
ENV VITE_BASE_URL=$VITE_BASE_URL
ENV ALLOWED_CORS=$ALLOWED_CORS
ENV ROOT_PATH=$ROOT_PATH

# TODO ENV VARIABELS
FROM node_base AS frontend_build
COPY ./frontend temp/frontend
WORKDIR /temp/frontend
RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --frozen-lockfile
RUN pnpm run build

FROM nginx:alpine-slim AS videos-with-subtitle-player_frontend
COPY --from=frontend_build /videos-with-subtitle-player-production/frontend/dist /usr/share/nginx/html
EXPOSE $FRONTEND_PORT
CMD ["nginx", "-g", "daemon off;"]

FROM golang:1.23.4-bookworm AS backend_build
COPY ./backend /temp/backend
COPY ./backend/go.mod /temp/backend/go.mod
COPY ./backend/go.sum /temp/backend/go.sum
WORKDIR /temp/backend
RUN go mod download
RUN go mod verify
# COPY *.go ./temp/backend
WORKDIR /temp/backend/cmd/api
COPY --from=frontend_build /temp/frontend/dist /temp/backend/cmd/api/public
RUN go build -v -o /temp/backend/cmd/api


FROM backend_build AS videos-with-subtitle-player
RUN mkdir /usr/share/videos-with-subtitle-player-backend
COPY --from=backend_build /temp/backend/cmd/api /usr/share/video-with-subtitle-player-backend
WORKDIR /usr/share/video-with-subtitle-player-backend
# RUN ROOT_PATH=$ROOT_PATH ALLOWED_CORS=$ALLOWED_CORS go build -v 
CMD ["/api"]