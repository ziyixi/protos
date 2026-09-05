# Shared protobuf definitions

The `protobuf` branch owns `proto/`. The `main` branch owns the generated Go
modules and generation script. Do not hand-edit generated `.pb.go` files.

## Gemini model catalog

The text-generation model IDs in `proto/todofy/large_language_model.proto` were
verified against the [Gemini model catalog](https://ai.google.dev/gemini-api/docs/models)
and [deprecation schedule](https://ai.google.dev/gemini-api/docs/deprecations)
on 2026-09-04. The current stable Flash model is `gemini-3.8-flash`.

The catalog also includes stable 3.7 Flash, 3.6 Flash, 3.5 Flash, 3.5 Flash-Lite,
3.1 Flash-Lite, and the optional 3.1 Pro Preview. Image, audio, video, and embedding
models are intentionally excluded from this text-summary service.

Existing enum numbers are never changed or reused. Legacy 2.5 and 3 Flash Preview
choices are marked deprecated for new project usage, but remain available for
explicit requests and wire compatibility. This annotation is a project migration
policy, **not** a claim that Google has shut down those endpoints. Previously
deprecated 1.5, 2.0, and experimental values are also retained.

`MODEL_UNSPECIFIED` delegates model selection and fallback to the server. An
explicit model requests only that model. Defaults are server policy, not an enum
ordering: todofy uses 3.8 Flash, then 3.7 Flash, then 3.5 Flash-Lite.

## Publishing and consuming Go bindings

1. Edit and push `.proto` files on `protobuf`.
2. The `Generate Go Modules from Proto Files` Action checks out the exact source
   commit, generates bindings, verifies/tests/vets/builds the Go modules, and
   commits generated `go/` files to `main` only after validation succeeds.
3. Check that the Action succeeded and that its generated `main` commit names
   the source commit. The Action run itself belongs to `protobuf`; the published
   module belongs to `main`. A bot push does not require a second Action run.
4. In a Go consumer, pin the generated `main` commit (not the source commit):

   ```sh
   go get github.com/ziyixi/protos/go/todofy@<generated-main-commit>
   go mod tidy
   go test ./...
   ```

Validate the source schema locally with `protoc` (output stays outside the repo):

```sh
proto_check_dir=$(mktemp -d)
protoc --proto_path=. --descriptor_set_out="$proto_check_dir/todofy.pb" proto/todofy/*.proto
```

The schema is language-neutral. Currently only Go bindings are published by this
workflow; another language should generate its own bindings from a pinned
`protobuf` commit, not consume the generated Go module.
