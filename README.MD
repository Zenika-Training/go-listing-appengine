Project listing the services in your current appengine project

# Go

Use the `go 1.26` version.

# Install gcloud command

Use the `gcloud components install app-engine-go` command

# Install dependencies

```
go mod vendor
```

# Documentation

If needed, please read these links:

- [Admin API](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services/list)
- [Go Doc for google.golang.org/api/appengine/v1](https://pkg.go.dev/google.golang.org/api/appengine/v1?tab=doc)

# Test

`go run listing.go`

# Deploy

Automatic: each push on `master` runs the CircleCI job `deploy` (`.circleci/config.yml`),
which deploys the `default` service of the `zen-formations` project and promotes the new version.
It uses the `training-deploy-gae` context (`GAE_KEY_FILE_CONTENT`).

Manual:
`gcloud app deploy --project zen-formations ./app.yaml --version go-listing-XX --promote`
