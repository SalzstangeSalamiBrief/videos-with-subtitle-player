# Dev folder

The content of this folder is related to optional things that can be used while working on the frontend.

## Content

The folder contains these things for development
|Name|Description|Usage|Limitations|
|-|-|-|-|
|json-server-database|A database that can be used while development as a replacement for the actual backend. If you want to develop without preparing the backend (e. g. adding folders with files) you can use this. The content is used by [json-server](https://github.com/typicode/json-server).| A command for running the server is added to the [package.json](../package.json). Run this command in a extra terminal.|You cannot display media files because this server is not able to provide them. These files can be added via the public folder but would increase the size of the repository.|
