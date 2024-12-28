# Videos with subtitle player

This software utilizes http streaming to stream multimedia files to a client.
Such files could be video or audio files with or without subtitle tracks.

## Contributors

- [SalzstangeSalamiBrief](https://github.com/SalzstangeSalamiBrief)

## Changelog

| Date                     | Description                                                                                                                                                      |
| ------------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 07.01.2024 - 30.03.2024  | Initialize Project by adding a front- and backend                                                                                                                |
| 30.03.2024               | Add support for videos                                                                                                                                           |
| 01.04.2024               | Replace ReactRouter with TanstackRouter                                                                                                                          |
| 20.04.2024               | Add tests                                                                                                                                                        |
| 24.04.2024 - 27.04.2024  | Refine router and path calculation in the backend                                                                                                                |
| 01.04.2024               | Test refinement                                                                                                                                                  |
| 03.08.2024 - 230.08.2024 | Restructure information architecture                                                                                                                             |
| 23.09.2024               | Add JSONServer as alternative for the real backend while developing                                                                                              |
| 30.09.2024               | Add image (thumbnail) support to the backend and display thumbnails in the frontend                                                                              |
| 01.10.2024               | <ul><li>Create a Github project and move todos from the readme to the project</li><li>Fix an issue that created a new id for each item on page refresh</li></ul> |

## Motivation

The goal of this project is to learn go and provide a solution to my problem, that I am missing a program to display multimedia files with separate subtitle tracks.

## How the streaming works

To enable the user to navigate through content in the used media player, the idea of _partial content (206)_ is used:
If a file just gets flushed to the client, then the client has to wait to receive the whole file, before playing it and enabling the user to navigate through it.
Instead of that, the client requests a byte range of the file and gets these bytes served.
Before the buffer of the client runs out, a request with the next byte range is sent.

## Used technologies

This project ist contains a frontend and a backend with their own technology stack

## Frontend

- React
- Ant Design
- TanStack Router

## Backend

- Go

## Technological Requirements

- Node 20
- Go 1.22

## Configuration

## How to develop

### How to debug the player

```javascript
const player = document.querySelector("video");
const events = [
  "abort",
  "canplay",
  "canplaythrough",
  "durationchange",
  "emptied",
  "encrypted",
  "ended",
  "error",
  "interruptbegin",
  "interruptend",
  "loadeddata",
  "loadedmetadata",
  "loadstart",
  "mozaudioavailable",
  "pause",
  "play",
  "playing",
  "progress",
  "ratechange",
  "seeked",
  "seeking",
  "stalled",
  "suspend",
  "timeupdate",
  "volumechange",
  "waiting",
];
events.forEach((a) =>
  player.addEventListener(a, (e) => console.log(e.type, e))
);
```

## Docker

### Backend

```bash
docker build . --target videos-with-subtitle-player_backend --tag videos-with-subtitle-player_backend:latest --build-arg VITE_BASE_URL=localhost:3000

# While running in docker HOST_ADDRESS has to be 0.0.0.0 to enable the server to listen to incoming requests
docker run -v C:/Temp:/Temp -p 3000:3000 -e ALLOWED_CORS=4200 -e ROOT_PATH=/Temp -e HOST_ADDRESS=0.0.0.0 -e HOST_PORT=3000 videos-with-subtitle-player_backend
```

### Frontend

```bash
docker build . --target videos-with-subtitle-player_frontend --tag videos-with-subtitle-player_frontend:latest --build-arg VITE_BASE_URL=localhost:3000

# to use nginx the port inside the container has to be 80
docker run -p 4200:80 videos-with-subtitle-player_frontend
```
