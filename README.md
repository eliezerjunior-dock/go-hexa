# go-hexa

Install grpc-health-probe

see : https://github.com/grpc-ecosystem/grpc-health-probe/releases
see : https://cloud.google.com/blog/topics/developers-practitioners/health-checking-your-grpc-servers-gke

wget https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.4.6/grpc_health_probe-linux-386

Testing healthcheck
./grpc_health_probe-linux-386 -addr=127.0.0.1:50051 -rpc-header=jwt:cookie

https://istio.io/latest/docs/ops/configuration/mesh/app-health-check/
https://github.com/grpc/grpc/blob/master/doc/health-checking.md
