# go inworld.ai client
unofficial inworld client / sdk written in go 

wip

### build proto
```bash
protoc --go_out=. --go-grpc_out=. -I proto/ proto/* && \
 mv ./github.com/zivoy/go-inworld/internal/protoBuf ./internal && \
 rm -r ./github.com
```