# Readme

## Development

Use the following command to start up the development environment:
`overmind start`

## Build

The application is built into a docker container and released
under cdear/camp-forms repository.

This can be updated via the following command:

```bash
docker build -t cdear/camp-forms:latest . && docker push cdear/camp-forms:latest
```
