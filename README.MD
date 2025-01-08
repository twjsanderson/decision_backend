
## Scripts
Ensure all cli bash scripts are executable
`chmod +x <path>`

To format code base, run `./cmd/cli/format-code.sh`

## Local Development

1. Add secrets and config data to **.local.env** and change it to **.env**
2. Run `./cmd/cli/setup-local-mod.sh` to build **go.mod** local development 
3. Run `go mod tidy`
4. Run `air` to build binary locally

## Deployment

#### Build & Test Production Binary
`./cmd/cli/create-production-build.sh`

You should see **Producion binary ran successfully.** in your terminal

Double check that `replace github.com/twjsanderson/decision_backend => ./` has been removed from the go.mod file


#### 
TODO: 
1. add tests 
2. add tests script in production build cli
3. consider better req body validation
