
## Local Development

1. Create and populate an `.env` file
2. Ensure that your `go.mod` file includes `replace github.com/twjsanderson/decision_backend => ./` (this must be included to use imports during local development)
3. Run `go mod tidy`

## Deployment

Run all commands from project root

#### Ensure script is executable
`chmod +x cmd/cli/create-production-build.sh`

#### Build & Test Production Binary
`./cmd/cli/create-production-build.sh`

You should see **Producion binary ran successfully.** in your terminal

Double check that `replace github.com/twjsanderson/decision_backend => ./` has been removed from the go.mod file


#### 
