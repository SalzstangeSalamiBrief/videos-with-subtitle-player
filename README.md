# Videos with subtitle player

This software utilizes http streaming to stream multimedia files to a client.
Such files could be video or audio files with or without subtitle tracks.

## Contributors

- [SalzstangeSalamiBrief](https://github.com/SalzstangeSalamiBrief)

## Changelog

| Date                    | Description                                       |
| ----------------------- | ------------------------------------------------- |
| 07.01.2024 - 30.03.2024 | Initialize Project by adding a front- and backend |
| 30.03.2024              | Add support for videos                            |
| 01.04.2024              | Replace ReactRouter with TanstackRouter           |
| 20.04.2024              | Add tests                                         |
| 24.04.2024 - 27.04.2024 | Refine router and path calculation in the backend |
| 01.04.2024              | Test refinement                                   |
| 03.08.2024 - TODO       | Restructure information architecture              |

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

## How to build

## How to test

## Possible Improvements

There are some things that could be implemented in the future to improve the software.
But before implementing them, a detailed analysis of the requirements and the general context has to be performed, to decide if implementing would the idea would be a real benefit to the software.

### Use a database for the files

Currently, all multimedia files are part of an in-memory tree like structure that is constructed on start.
Because of that, IDs of the files change after each restart of the server.
To prevent this behavior, a database could be used to persist each item of the tree.
But with implementing this idea, the items in the database could become stale (e.g. a file is deleted, or another one is added).
Now another mechanism should be found and implemented to prevent this behavior.

#### Caching on the server

Using a database increases the response time, because the server has to communicate with another entity, the database, to get the data.
To reduce the newly introduced latency, a cache could be used.
Using a cache would increase the complexity of the app and the update behavior of the cache (e.g. time) has to be defined.

### Display Illustrations

## TODOs

- Display Illustrations
- Fix styling of the GUI
- Add integrations tests
- Error Handler Middleware
- Logging
- generic, chainable middlewares without an order
- Empty root folder crashes the application
- breadcrumbs: display the current folder while watching a video
- Reload of the frontend crashes the app => Maybe use the loader of the current route instead of the \_\_root => is it merged by tanstack query?
- cards: try to make the title clickable and not the whole card? Test and evaluate
