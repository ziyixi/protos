"""Generate and build a wheel in isolation; consumers never need protoc.

Only newsletter messages are published in the first Python variant. Shared Go
source paths/options remain untouched. Extend the explicit source list when a
real Python consumer needs another domain, not by rewriting generated imports.
"""

import argparse
import hashlib
import json
import re
import shutil
import subprocess
import sys
import tempfile
from pathlib import Path

PROTOC_VERSION = "libprotoc 36.0"
SOURCE_PATH = "proto/newsletter/editorial.proto"


def sha256(path: Path) -> str:
    return hashlib.sha256(path.read_bytes()).hexdigest()


def main() -> None:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--version", required=True)
    parser.add_argument("--source-commit", required=True)
    parser.add_argument("--output", type=Path, required=True)
    parser.add_argument("--uv", default="uv")
    args = parser.parse_args()
    if not re.fullmatch(r"0\.1\.0\.dev[0-9]+", args.version):
        parser.error("Expected release version 0.1.0.dev<workflow run number>")
    if not re.fullmatch(r"[a-f0-9]{40}", args.source_commit):
        parser.error("source-commit must be a full Git SHA")
    protoc = shutil.which("protoc")
    uv = shutil.which(args.uv)
    if not protoc or not uv:
        parser.error("Install pinned protoc and uv explicitly before building")
    if subprocess.check_output([protoc, "--version"], text=True).strip() != PROTOC_VERSION:
        parser.error("This generator requires protoc 36.0")
    root = Path(__file__).resolve().parents[1]
    output = args.output.resolve()
    output.mkdir(parents=True, exist_ok=True)
    if list(output.glob("*.whl")):
        parser.error("Output already contains a wheel; use a fresh output directory")
    with tempfile.TemporaryDirectory(prefix="ziyixi-protos-build-") as directory:
        stage = Path(directory)
        package = stage / "src/ziyixi_protos"
        domain = package / "newsletter"
        domain.mkdir(parents=True)
        virtual_root = stage / "proto-input"
        virtual = virtual_root / "ziyixi_protos/newsletter/editorial.proto"
        virtual.parent.mkdir(parents=True)
        shutil.copyfile(root / SOURCE_PATH, virtual)
        shutil.copyfile(root / "python/pyproject.toml", stage / "pyproject.toml")
        (package / "__init__.py").write_text(f'__version__ = "{args.version}"\n')
        (package / "py.typed").touch()
        (domain / "__init__.py").write_text(
            '"""Newsletter v1 messages; no HTTP or gRPC client."""\n'
        )
        subprocess.run(
            [
                protoc,
                f"--proto_path={virtual_root}",
                f"--python_out={stage / 'src'}",
                f"--pyi_out={stage / 'src'}",
                "ziyixi_protos/newsletter/editorial.proto",
            ],
            check=True,
        )
        # A separate interpreter avoids polluting the builder's descriptor pool.
        descriptor_hash = subprocess.check_output(
            [
                sys.executable,
                "-c",
                (
                    "import sys,hashlib; sys.path.insert(0,sys.argv[1]); "
                    "from ziyixi_protos.newsletter import editorial_pb2 as pb; "
                    "print(hashlib.sha256(pb.DESCRIPTOR.serialized_pb).hexdigest())"
                ),
                str(stage / "src"),
            ],
            text=True,
        ).strip()
        manifest = {
            "source_repository": "https://github.com/ziyixi/protos",
            "source_path": SOURCE_PATH,
            "source_commit": args.source_commit,
            "package_version": args.version,
            "source_sha256": sha256(root / SOURCE_PATH),
            "protoc_version": PROTOC_VERSION,
            "generated_sha256": sha256(domain / "editorial_pb2.py"),
            "stubs_sha256": sha256(domain / "editorial_pb2.pyi"),
            "descriptor_sha256": descriptor_hash,
        }
        (domain / "provenance.json").write_text(
            json.dumps(manifest, sort_keys=True, indent=2) + "\n"
        )
        subprocess.run(
            [
                uv,
                "build",
                "--python",
                sys.executable,
                "--wheel",
                "--out-dir",
                str(output),
                str(stage),
            ],
            check=True,
        )
    wheels = list(output.glob("*.whl"))
    if len(wheels) != 1:
        raise RuntimeError("Expected exactly one release wheel")
    (output / "SHA256SUMS").write_text(f"{sha256(wheels[0])}  {wheels[0].name}\n")


if __name__ == "__main__":
    main()
