# lnkshrtnr
A lightweight link shortener built with Go and Redis.

## Table of contents
- [API](#api)
- [Disclaimer](#disclaimer)
- [License](#license)

## API

### `POST /api/links`  
Shorten a link

**Example body**
```json
{
  "url": "http://example.com"
}
```

**Example response**
```json
{
  "id": "1Dfu27E9nYr27l0Y"
}
```

### `GET /:id`  
Redirects to the url mapped to the id or returns `404 page not found`.

## Disclaimer
This project was intended as a demo-project, there may be hickups, bugs and roadblocks.

## License
MIT