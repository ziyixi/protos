"""Run against an installed release wheel, not generated files on sys.path."""

import hashlib
import json
import os
import pickle
import unittest
from importlib.metadata import distribution
from importlib.resources import files

from google.protobuf import json_format
from ziyixi_protos.newsletter import editorial_pb2 as pb


class PublishedPackageTests(unittest.TestCase):
    def test_metadata_and_provenance(self):
        package = files("ziyixi_protos.newsletter")
        manifest = json.loads(package.joinpath("provenance.json").read_text())
        self.assertEqual(manifest["package_version"], distribution("ziyixi-protos").version)
        self.assertRegex(manifest["source_commit"], r"^[a-f0-9]{40}$")
        if "EXPECTED_SOURCE_COMMIT" in os.environ:
            self.assertEqual(manifest["source_commit"], os.environ["EXPECTED_SOURCE_COMMIT"])
        self.assertEqual(manifest["source_repository"], "https://github.com/ziyixi/protos")
        self.assertEqual(manifest["source_path"], "proto/newsletter/editorial.proto")
        self.assertEqual(manifest["protoc_version"], "libprotoc 36.0")
        for filename, key in [
            ("editorial_pb2.py", "generated_sha256"),
            ("editorial_pb2.pyi", "stubs_sha256"),
        ]:
            self.assertEqual(
                hashlib.sha256(package.joinpath(filename).read_bytes()).hexdigest(), manifest[key]
            )
        self.assertEqual(
            hashlib.sha256(pb.DESCRIPTOR.serialized_pb).hexdigest(), manifest["descriptor_sha256"]
        )
        self.assertRegex(manifest["source_sha256"], r"^[a-f0-9]{64}$")

    def test_typed_and_messages_only(self):
        package = distribution("ziyixi-protos")
        self.assertTrue(files("ziyixi_protos").joinpath("py.typed").is_file())
        self.assertEqual(package.requires, ["protobuf<8,>=7.36.0"])
        self.assertFalse(any("grpc" in str(path) for path in package.files))

    def test_binary_json_and_pickle_round_trips(self):
        request = pb.StartRunRequest(request_key="published-wheel", issue_date="2026-09-05")
        self.assertEqual(pb.StartRunRequest.FromString(request.SerializeToString()), request)
        value = json_format.MessageToDict(request, preserving_proto_field_name=True)
        self.assertEqual(value, {"request_key": "published-wheel", "issue_date": "2026-09-05"})
        self.assertEqual(json_format.ParseDict(value, pb.StartRunRequest()), request)
        self.assertEqual(pickle.loads(pickle.dumps(request)), request)
        self.assertEqual(pb.StartRunRequest.__module__, "ziyixi_protos.newsletter.editorial_pb2")

    def test_presence_and_oneof(self):
        digest = pb.PersonalDigest()
        self.assertFalse(digest.HasField("task_count"))
        digest.task_count = 0
        self.assertTrue(digest.HasField("task_count"))
        point = pb.ChartPoint(decimal_value="0")
        self.assertEqual(point.WhichOneof("observation"), "decimal_value")
        point.missing_reason = "not reported"
        self.assertFalse(point.HasField("decimal_value"))

    def test_workflow_and_usage_contracts(self):
        summary = pb.UsageSummary(invocations=2, missing_invocations=1, partial=True)
        self.assertFalse(summary.HasField("usage"))
        summary.usage.total_tokens = 0
        self.assertTrue(summary.HasField("usage"))
        summary.usage.input_tokens = 1200
        summary.usage.cached_input_tokens = 1000
        summary.usage.output_tokens = 80
        summary.usage.total_tokens = 1280
        run = pb.CollectionRun(id="run", usage=summary)
        run.workflow.definition_hash = "a" * 64
        run.workflow.nodes.add(id="research", type="research", state="succeeded")
        self.assertEqual(pb.CollectionRun.FromString(run.SerializeToString()), run)
        value = json_format.MessageToDict(run, preserving_proto_field_name=True)
        self.assertEqual(value["usage"]["usage"]["total_tokens"], "1280")
        self.assertEqual(json_format.ParseDict(value, pb.CollectionRun()), run)
        candidate = pb.Candidate(id="lead", provenance="crossref_metadata", access_scope="metadata")
        task = pb.ResearchTask(id="research-lead", candidate_ids=[candidate.id], priority=1)
        self.assertEqual(task.candidate_ids, ["lead"])

    def test_contract_namespace_and_methods(self):
        self.assertEqual(pb.DESCRIPTOR.package, "newsletter.v1")
        service = pb.DESCRIPTOR.services_by_name["NewsletterService"]
        self.assertEqual(
            [method.name for method in service.methods],
            [
                "StartRun",
                "GetRun",
                "PutPacket",
                "ReadInbox",
                "PrepareEdition",
                "GetEdition",
                "RenderEdition",
                "SendEdition",
            ],
        )


if __name__ == "__main__":
    unittest.main()
