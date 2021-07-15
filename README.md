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
curl -d '{"name": "Coronavirus key indicators","publish_date": "2020-05-05T14:58:29.317Z", "content_type": "static_landing_page", "content": "eyJjaGFydHMiOltdLCJ0YWJsZXMiOltdLCJpbWFnZXMiOltdLCJlcXVhdGlvbnMiOltdLCJkb3dubG9hZHMiOltdLCJtYXJrZG93biI6WyJPdXIgd2lkZSByYW5nZSBvZiBlY29ub21pYyBhbmQgc29jaWFsIHN0YXRpc3RpY3MgaW5jbHVkZTpcblxuLSB0aGUgVUtcdTAwMjdzIE5hdGlvbmFsIEFjY291bnRzIChzdWNoIGFzIGdyb3NzIGRvbWVzdGljIHByb2R1Y3QgKEdEUCksIG5hdGlvbmFsIGluY29tZSBhbmQgZXhwZW5kaXR1cmUpXG4tIHRoZSBVSyBCYWxhbmNlIG9mIFBheW1lbnRzXG4tIHBvcHVsYXRpb24sIGRlbW9ncmFwaHkgYW5kIG1pZ3JhdGlvbiBzdGF0aXN0aWNzXG4tIGdvdmVybm1lbnQgb3V0cHV0IGFuZCBhY3Rpdml0eVxuLSBidXNpbmVzcyBvdXRwdXQgYW5kIGFjdGl2aXR5XG4tIHByaWNlcyBzdGF0aXN0aWNzIChzdWNoIGFzIGNvbnN1bWVyIHByaWNlcyBhbmQgcHJvZHVjZXIgcHJpY2VzKSBcbi0gdGhlIGxhYm91ciBtYXJrZXQgKHN1Y2ggYXMgZW1wbG95bWVudCwgdW5lbXBsb3ltZW50IGFuZCBlYXJuaW5ncylcbi0gdml0YWwgZXZlbnRzIHN0YXRpc3RpY3MgKHN1Y2ggYXMgYmlydGhzLCBtYXJyaWFnZXMgYW5kIGRlYXRocykgXG4tIHNvY2lhbCBzdGF0aXN0aWNzIChmb3IgZXhhbXBsZSwgYWJvdXQgbmVpZ2hib3VyaG9vZHMgYW5kIGZhbWlsaWVzKVxuLSBlY29ub21pYywgc29jaWV0YWwgYW5kIHBlcnNvbmFsIHdlbGwtYmVpbmdcblxuV2UgZGVzaWduIGFuZCBydW4gdGhlIGNlbnN1cyBpbiBFbmdsYW5kIGFuZCBXYWxlcyBldmVyeSAxMCB5ZWFycy4gV2UgYWxzbyB3b3JrIHdpdGggdGhlIGRldm9sdmVkIGFkbWluaXN0cmF0aW9ucyBpbiBOb3J0aGVybiBJcmVsYW5kIGFuZCBTY290bGFuZCB0byBjYXJyeSBvdXQgdGhlIGNlbnN1cyBvZiBwb3B1bGF0aW9uLlxuXG5XZSBhbHNvIHByb3ZpZGUgYSBzZXJ2aWNlIGZvciB5b3UgdG8gW3JlcXVlc3Qgc3BlY2lmaWMgaW5mb3JtYXRpb25dWzFdLlxuXG5BbGwgb2Ygb3VyIHN0YXRpc3RpY3MgYXJlIFtvZmZpY2lhbCBzdGF0aXN0aWNzXVsyXS4gVGhleSBhcmUgdXNlZCB0bzpcblxuLSBnaXZlIGNpdGl6ZW5zIGEgdmlldyBvZiBzb2NpZXR5IGFuZCBvZiB0aGUgd29yayBhbmQgcGVyZm9ybWFuY2Ugb2YgZ292ZXJubWVudCwgYWxsb3dpbmcgdGhlbSB0byBhc3Nlc3MgdGhlIGltcGFjdCBvZiBnb3Zlcm5tZW50IHBvbGljaWVzIGFuZCBhY3Rpb25zIFxuXG4tIGluZm9ybSBwYXJsaWFtZW50cyBhbmQgcG9saXRpY2FsIGFzc2VtYmxpZXMgYWJvdXQgdGhlIHN0YXRlIG9mIHRoZSBuYXRpb24sIGdpdmluZyBhIHdpbmRvdyBvbiB0aGUgd29yayBhbmQgcGVyZm9ybWFuY2Ugb2YgZ292ZXJubWVudHMsIHRvIGFzc2VzcyB0aGVpciBwb2xpY2llc+KAmSBpbXBhY3QgXG5cbi0gYWxsb3cgZ292ZXJubWVudCBhbmQgaXRzIGFnZW5jaWVzIHRvIGNhcnJ5IG91dCB0aGVpciBidXNpbmVzcywgbWFraW5nIGluZm9ybWVkIGRlY2lzaW9ucyBiYXNlZCBvbiBldmlkZW5jZVxuXG4tIGdpdmUgbWluaXN0ZXJzIGEgcGljdHVyZSBvZiB0aGUgZWNvbm9teSBhbmQgc29jaWV0eSwgc28gdGhleSBjYW4gZGV2ZWxvcCBhbmQgZXZhbHVhdGUgZWNvbm9taWMgYW5kIHNvY2lhbCBwb2xpY2llcyBcblxuLSBwcm92aWRlIGJ1c2luZXNzZXMgd2l0aCB0aGUgaW5mb3JtYXRpb24gdG8gaGVscCB0aGVtIHJ1biBlZmZlY3RpdmVseSBhbmQgZWZmaWNpZW50bHlcblxuLSBoZWxwIGFuYWx5c3RzLCByZXNlYXJjaGVycywgc2Nob2xhcnMgYW5kIHN0dWRlbnRzIHdpdGggdGhlaXIgd29yayBcblxuLSBtZWV0IHRoZSBuZWVkcyBvZiB0aGUgRXVyb3BlYW4gVW5pb24gYW5kIG90aGVyIGludGVybmF0aW9uYWwgYm9kaWVzIHNvIHBlb3BsZSBjYW4gY29tcGFyZSBkYXRhIGFjcm9zcyBjb3VudHJpZXNcblxuXG4gIFsxXTogaHR0cHM6Ly93d3cub25zLmdvdi51ay9hYm91dHVzL3doYXR3ZWRvL3N0YXRpc3RpY3MvcmVxdWVzdGluZ3N0YXRpc3RpY3NcbiAgWzJdOiBodHRwczovL3d3dy5zdGF0aXN0aWNzYXV0aG9yaXR5Lmdvdi51ay9uYXRpb25hbC1zdGF0aXN0aWNpYW4vdHlwZXMtb2Ytb2ZmaWNpYWwtc3RhdGlzdGljcy8iXSwibGlua3MiOlt7InRpdGxlIjoiUmVxdWVzdGluZyBzdGF0aXN0aWNzIiwidXJpIjoiL2Fib3V0dXMvd2hhdHdlZG8vc3RhdGlzdGljcy9yZXF1ZXN0aW5nc3RhdGlzdGljcyJ9XSwidHlwZSI6InN0YXRpY19wYWdlIiwidXJpIjoiL2Fib3V0dXMvd2hhdHdlZG8vc3RhdGlzdGljcy9zdGF0aXN0aWNzd2Vwcm9kdWNlIiwiZGVzY3JpcHRpb24iOnsidGl0bGUiOiJTdGF0aXN0aWNzIHdlIHByb2R1Y2UiLCJzdW1tYXJ5IjoiIiwia2V5d29yZHMiOltdLCJtZXRhRGVzY3JpcHRpb24iOiIiLCJ1bml0IjoiIiwicHJlVW5pdCI6IiIsInNvdXJjZSI6IiJ9fQ=="}' "localhost:26400/cms/LMSV1/aboutus" | jq
```

Get content
```
curl "localhost:26400/cms/LMSV1/aboutus" | jq
```

### Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details.

### License

Copyright © 2021, Office for National Statistics (https://www.ons.gov.uk)

Released under MIT license, see [LICENSE](LICENSE.md) for details.

