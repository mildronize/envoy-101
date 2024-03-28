# Enjoy Playground

https://www.youtube.com/watch?v=40gKzHQWgP0


Terminology: https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/intro/terminology
connection_pooling: https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/connection_pooling#arch-overview-conn-pool


## 1. How to run `http-allbackend`


1. Run this to start the services
   
```sh
docker-compose -f docker-compose.http-allbackend.yml up`
```

Go to `http://localhost` to see the application running.

The result will be swapped between the two services.

```
App ID: 1111, Home Page
```

and

```
App ID: 2222, Home Page
```

## 2. How to run `https-basic`

create cert files

```sh
# Run this only once
mkcert -install

mkcert -cert-file certs/local-cert.pem -key-file certs/local-key.pem localhost
```

Run this to start the services
   
```sh
docker-compose -f docker-compose.https.yml up
```

Go to `https://localhost` to see the application running.

### How to config TLS

```yml
  transport_socket:
    name: envoy.transport_sockets.tls
    typed_config:
      "@type": type.googleapis.com/envoy.extensions.transport_sockets.tls.v3.DownstreamTlsContext
      common_tls_context: 
        ## Setups for TLS and Certificates
        tls_certificates:
          certificate_chain: {filename: "/certs/local-cert.pem"}
          private_key: {filename: "/certs/local-key.pem"}
        ## Setup support both HTTP/2 and HTTP/1.1.
        alpn_protocols: ["h2,http/1.1"]
        ## Setup minimum TLS version
        tls_params:
            tls_minimum_protocol_version: "TLSv1_2"
```