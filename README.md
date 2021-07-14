dp-content-api
================
An API for ONS website content

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
curl -d '{"name": "Coronavirus key indicators","publish_date": "2020-05-05T14:58:29.317Z", "content_type": "static_landing_page", "content": "ewogICJzZWN0aW9ucyI6IFsKICAgIHsKICAgICAgInN1bW1hcnkiOiAiV2UgYXJlIHRoZSBVS+KAmXMgbGFyZ2VzdCBpbmRlcGVuZGVudCBwcm9kdWNlciBvZiBvZmZpY2lhbCBzdGF0aXN0aWNzIiwKICAgICAgInRpdGxlIjogIldoYXQgd2UgZG8iLAogICAgICAidXJpIjogIi9hYm91dHVzL3doYXR3ZWRvIgogICAgfQogIF0sCiAgIm1hcmtkb3duIjogWwogICAgIldlIGFyZSB0aGUgVUvigJlzIGxhcmdlc3QgaW5kZXBlbmRlbnQgcHJvZHVjZXIgb2Ygb2ZmaWNpYWwgc3RhdGlzdGljcyIKICBdLAogICJsaW5rcyI6IFtdLAogICJ0eXBlIjogInN0YXRpY19sYW5kaW5nX3BhZ2UiLAogICJ1cmkiOiAiL2Fib3V0dXMiLAogICJkZXNjcmlwdGlvbiI6IHsKICAgICJ0aXRsZSI6ICJBYm91dCB1cyIsCiAgICAic3VtbWFyeSI6ICJXZSBhcmUgdGhlIFVL4oCZcyBsYXJnZXN0IGluZGVwZW5kZW50IHByb2R1Y2VyIG9mIG9mZmljaWFsIHN0YXRpc3RpY3MiLAogICAgImtleXdvcmRzIjogWwogICAgICAiY2FyZWVycyIsCiAgICAgICJjb250YWN0IiwKICAgICAgInByb2N1cmVtZW50IiwKICAgICAgImNvbW1lcmNpYWwiLAogICAgICAic3VwcGxpZXJzIgogICAgXSwKICAgICJtZXRhRGVzY3JpcHRpb24iOiAiVGhlIE9mZmljZSBmb3IgTmF0aW9uYWwgU3RhdGlzdGljcyAoT05TKSBpcyB0aGUgVUvigJlzIGxhcmdlc3QgaW5kZXBlbmRlbnQgcHJvZHVjZXIgb2Ygb2ZmaWNpYWwgc3RhdGlzdGljcyIsCiAgICAidW5pdCI6ICIiLAogICAgInByZVVuaXQiOiAiIiwKICAgICJzb3VyY2UiOiAiIgogIH0KfQ=="}' "localhost:26400/cms/LMSV1/aboutus" | jq
```

### Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details.

### License

Copyright © 2021, Office for National Statistics (https://www.ons.gov.uk)

Released under MIT license, see [LICENSE](LICENSE.md) for details.

