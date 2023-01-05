## Caching HTTP responses

The [`Cache-Control`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Cache-Control)
HTTP header field holds directives (instructions) — in both requests and
responses — that control caching in browsers and shared caches (e.g. Proxies,
CDNs).

If we send back the header
```
Cache-Control: max-age=604800
```
we're instructing the HTTP client to cache for 604800 seconds (7 days).

## Does Hasura cache responses for webhook authentication?

In Hasura, you can defer auth decisions to an external endpoint (i.e., a 
webhook).

https://hasura.io/docs/latest/auth/authentication/webhook/