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

## Publishing and consuming Python bindings

`Publish Python protobuf wheel` builds the exact `protobuf` source commit on
GitHub Actions and publishes a versioned wheel plus SHA256SUMS as GitHub Release
assets. It does not use PyPI or write Python artifacts to `main`. Pull requests
build/test only; publishing has a separate, minimal write-permission job.

The first Python package, `ziyixi-protos`, contains newsletter v1 messages,
`.pyi` types, `py.typed` and per-domain provenance (source commit/hash, compiler,
package version, generated files and descriptor hashes). Import with:

```python
from ziyixi_protos.newsletter import editorial_pb2
```

Python generation uses protoc 36.0, uv 0.12.10 and setuptools 84.0.0. The generated
package requires protobuf >=7.36.0,<8; consumers lock their tested runtime version.

Newsletter discovery candidates also carry optional author/affiliation/venue,
publication status, specific contribution, source-selection rationale and consulted
public URLs. These are unverified discovery context, not article approval. Empty
fields preserve old candidates and non-research events; no existing field was renumbered.
No gRPC client/server or application logic is included. Existing proto package,
Go options and source paths are unchanged; a virtual generation path gives Python
its independent namespace without patching generated imports.

Each workflow run gets `0.1.0.dev<run-number>` and release tag `python-v<version>`.
Pin both the exact version and its release wheel URL in the consumer's uv config,
commit `uv.lock`, and install with `uv sync --locked`. Do not use a `latest` URL,
an expiring Actions artifact, a sibling checkout or startup-time generation.
The consumer needs neither protoc nor GitHub credentials for these public assets.
New releases do not automatically upgrade existing consumers.

Local generation/testing uses `python/build.py --version 0.1.0.dev0
--source-commit <40-character-source-sha> --output <fresh-output-dir>`, with pinned
protoc, uv and a compatible protobuf runtime explicitly installed. Test the built
wheel in a fresh environment using `python/test_package.py`; never publish a local
build as if it were the verified Actions output.

Go and Python jobs are independent. Adding newsletter also generates the new Go
module; its initial runtime pins are copied from the already-tested Todofy module
before the existing Go verification gate. Todofy's module path and consumer pin
are not changed by publishing the Python variant.
# Newsletter workflow contracts

`proto/newsletter/editorial.proto` also defines unverified discovery candidates,
research tasks, read-only DAG progress, and provider-reported token usage. Workflow
definitions and instructions remain in the newsletter service repository; protobuf
describes data, not executable workflow logic. Missing usage is distinct from zero.
Cached input and reasoning output are subsets, not additional tokens to sum.
