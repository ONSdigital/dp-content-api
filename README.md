dp-content-api
================
An API for ONS website content

`POST  /cms/{collection_id}/{url}` - Add content for the given collection_id and URL

`GET   /cms/{collection_id}/{url}` - Get content for the given collection_id and URL 

`PATCH /cms/{collection_id}/{url}` - Update content for the given collection_id and URL

`GET   /{url}` - Get the latest published content for the given URL

### Caching

When getting the latest published content for a page, a `Cache-Control` header is returned. 
```
Cache-Control: public, max-age=600
```
The `max-age` value represents how long the page can safely be cached for.
If there is no publish due for the page, a default cache time is returned. 
The cache time will decrease as a publish time approaches, so when the publish takes place the page will not be cached.

### Getting started

* Run `make debug`

### Dependencies

* No further dependencies other than those defined in `go.mod`

### Configuration

| Environment variable           | Default     | Description
| ------------------------------ | ----------- | -----------
| BIND_ADDR                      | :26400      | The host and port to bind to
| GRACEFUL_SHUTDOWN_TIMEOUT      | 5s          | The graceful shutdown timeout in seconds (`time.Duration` format)
| HEALTHCHECK_INTERVAL           | 30s         | Time between self-healthchecks (`time.Duration` format)
| HEALTHCHECK_CRITICAL_TIMEOUT   | 90s         | Time to wait until an unhealthy dependent propagates its state to make this app unhealthy (`time.Duration` format)
| MONGODB_CONTENT_DATABASE       | content     | The MongoDB content database
| MONGODB_CONTENT_COLLECTION     | content     | The MongoDB content collection
| MONGODB_USERNAME               | test        | The MongoDB Username
| MONGODB_PASSWORD               | test        | The MongoDB Password
| MONGODB_CA_FILE_PATH           | file-path   | The MongoDB CA FilePath

### Example requests

Post content
```
curl -d '{"name": "Coronavirus key indicators","publish_date": "2020-05-05T00:00:00.000Z", "content_type": "static_landing_page", "content": "ewogICAidHlwZSI6InN0YXRpY19sYW5kaW5nX3BhZ2UiLAogICAibWFya2Rvd24iOiJWZXJzaW9uIDEiCn0="}' "localhost:26400/cms/LMSV1/aboutus" | jq
```

Patch content
```
curl -X PATCH -d '[ { "op": "replace", "path": "publish_date", "value": "2021-06-06T00:00:00.000Z" }]' "localhost:26400/cms/LMSV1/aboutus" | jq
curl -X PATCH -d '[ { "op": "replace", "path": "approved", "value": "true" }]' "localhost:26400/cms/LMSV1/aboutus" | jq
curl -X PATCH -d '[ { "op": "replace", "path": "content", "value": "ewogICAidHlwZSI6InN0YXRpY19sYW5kaW5nX3BhZ2UiLAogICAibWFya2Rvd24iOiJWZXJzaW9uIDEiCn0=" }]' "localhost:26400/cms/LMSV1/aboutus" | jq
```

Get collection content
```
curl "localhost:26400/cms/LMSV1/aboutus" | jq
```

Get latest published content

```
curl "localhost:26400/aboutus" | jq
```

##### Full process

Post V1 content to a collection with ID LMSV1. Imitate a manual publish without a publish date specified.
```
curl -d '{"name": "Coronavirus key indicators", "content_type": "static_landing_page", "content": "ewogICAidHlwZSI6InN0YXRpY19sYW5kaW5nX3BhZ2UiLAogICAibWFya2Rvd24iOiJWZXJzaW9uIDEiCn0="}' "localhost:26400/cms/LMSV1/aboutus" | jq
```

The collection content can be retrieved/previewed by using the content endpoint with the collection ID
```
curl "localhost:26400/cms/LMSV1/aboutus" | jq
```
Expected response
```
{
  "type": "static_landing_page",
  "markdown": "Version 1"
}
```

A GET request for the published page returns no content, as it's not approved / published
```
curl "localhost:26400/aboutus" | jq
```

Set the approved flag and publish date in the past to manually publish the content
```
curl -X PATCH -d '[ { "op": "replace", "path": "publish_date", "value": "2020-06-06T00:00:00.000Z" }, { "op": "replace", "path": "approved", "value": "true" }]' "localhost:26400/cms/LMSV1/aboutus" | jq
```

A GET request for the published page now returns version 1
```
curl "localhost:26400/aboutus" | jq
```

Post V2 content to a different collection. This time setting the approval flag and adding a publish date in the future
** update the publish time as required, but note that the time is in UTC, so may need to be put back an hour to account for BST **
```
curl -d '{"name": "Coronavirus key indicators", "publish_date":"2021-08-19T14:00:00.000Z", "approved":true, "content_type": "static_landing_page", "content": "ewogICAidHlwZSI6InN0YXRpY19sYW5kaW5nX3BhZ2UiLAogICAibWFya2Rvd24iOiJWZXJzaW9uIDIiCn0="}' "localhost:26400/cms/LMSV2/aboutus" | jq
```

The V2 collection content can be retrieved/previewed by using the content endpoint with the collection ID
```
curl "localhost:26400/cms/LMSV2/aboutus" | jq
```
Expected response
```
{
  "type": "static_landing_page",
  "markdown": "Version 2"
}
```

While the V2 publish date is in the future, A GET request for the published page still returns version 1.
Once the publish date for V2 is in the past, the V2 content should be returned
```
curl "localhost:26400/aboutus" | jq
```

For testing, the publish date can be updated using:
```
curl -X PATCH -d '[ { "op": "replace", "path": "publish_date", "value": "2021-07-20T13:52:00.000Z" }]' "localhost:26400/cms/LMSV2/aboutus" | jq
```

### Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details.

### License

Copyright © 2021, Office for National Statistics (https://www.ons.gov.uk)

Released under MIT license, see [LICENSE](LICENSE.md) for details.

