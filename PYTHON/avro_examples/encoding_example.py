"""Avro encoding test sample."""

import io

import avro.schema
from avro.datafile import DataFileReader, DataFileWriter
from avro.io import DatumReader, DatumWriter


class EncodingExample:  # pylint: disable=too-few-public-methods
    """Class wrapper for testing Avro Encoding."""

    @staticmethod
    def basic_encoding_example():
        """Performs a basic Avro encoding and decoding, displaying the contents in the cosole."""
        # schema = avro.schema.parse(open("avro\\new_catalog_data.avsc", "rb").read())
        schema = avro.schema.parse(
            '{"namespace": "gsm.dms","type": "record",'
            '"name": "new_catalog_data","fields": [{'
            '"name": "dataAssetType", "type": "string"},'
            '{"name": "sourceId",  "type": "string"},'
            '{"name": "fileUri", "type": "string"}]}'
        )

        print(f"########## schema:  {schema.to_json()}")

        kafka_payload_stream = io.BytesIO()
        kafka_payload_stream.mode = "b"

        writer = DataFileWriter(kafka_payload_stream, DatumWriter(), schema)
        writer.append(
            {"dataAssetType": "image", "sourceId": "source_ferg", "fileUri": "this.img"}
        )
        writer.append(
            {"dataAssetType": "text", "sourceId": "source_ferg", "fileUri": "that.txt"}
        )
        writer.flush()

        print(f"##########  (binary pre send)  {kafka_payload_stream.getvalue()}")

        reader = DataFileReader(kafka_payload_stream, DatumReader())
        for message in reader:
            print(f"##########  decoded message  {message}")
        writer.close()
        reader.close()
