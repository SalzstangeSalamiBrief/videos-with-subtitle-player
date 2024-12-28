FROM node:20-slim AS node_base
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable


ARG VITE_BASE_URL

# Set environment variables during the build process
ENV VITE_BASE_URL=$VITE_BASE_URL

FROM node_base AS frontend_build
COPY ./frontend temp/frontend
WORKDIR /temp/frontend
RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --frozen-lockfile
RUN pnpm run build

FROM nginx:alpine-slim AS videos-with-subtitle-player_frontend
COPY --from=frontend_build /temp/frontend/dist /usr/share/nginx/html
CMD ["nginx", "-g", "daemon off;"]

FROM golang:1.23.4-bookworm AS backend_build
WORKDIR /temp/backend
COPY /backend .
RUN go mod download
RUN go mod verify
WORKDIR /temp/backend/cmd/api
COPY --from=frontend_build /temp/frontend/dist /temp/backend/cmd/api/public
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /backend-output


FROM gcr.io/distroless/base-debian12 AS videos-with-subtitle-player_backend
WORKDIR /
COPY --from=backend_build /backend-output /backend
EXPOSE 3000
USER nonroot:nonroot
ENTRYPOINT ["/backend"]