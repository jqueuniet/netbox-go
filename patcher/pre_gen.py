#!/usr/bin/env python3

import yaml
import argparse
import logging
import os
import sys


MISSING_REQUIRED_FIELDS = (
    "console_port_template_count",
    "console_server_port_template_count",
    "device_bay_template_count",
    "front_port_template_count",
    "interface_template_count",
    "inventory_item_template_count",
    "module_bay_template_count",
    "power_outlet_template_count",
    "power_port_template_count",
    "rear_port_template_count",
    "devicetype_count",
    "device_count",
    "inventoryitem_count",
    "virtualmachine_count",
    "prefix_count",
    "platform_count",
    "tagged_items",
    "vlan_count",
    "rack_count",
    "vrf_count",
    "circuit_count",
    "display_url",
    "cluster_count",
    "site_count",
    "created",  # for sites
    "provider_count",  # ASNs
    "display",  # Virtual disks
)


def build_upstream_path(version: str) -> str:
    return f"schemas/upstream-{version}.yaml"


def build_patched_path(version: str) -> str:
    return f"schemas/patched-{version}.yaml"


def discover_package_version(version: str) -> str:
    gen_config = f"configs/oapi-codegen-v{version}.yaml"
    if not os.path.isfile(gen_config):
        print(
            f"Invalid version, could not find matching generator config {gen_config} file",
            file=sys.stderr,
        )
        sys.exit(1)

    with open(gen_config, "r") as file:
        gen_config_data = yaml.load(file, Loader=yaml.CSafeLoader)

    if "packageVersion" not in gen_config_data:
        print(
            f"Could not find packageVersion field in generator config {gen_config}",
            file=sys.stderr,
        )
        sys.exit(1)

    return gen_config_data["packageVersion"]


def patch_spec_file(version: str):
    input_path = build_upstream_path(version)
    if not os.path.isfile(input_path):
        print(
            f"Could not find upstream spec file {input_path} from generator config packageVersion",
            file=sys.stderr,
        )
        sys.exit(1)

    with open(input_path, "r") as file:
        data = yaml.load(file, Loader=yaml.CSafeLoader)

    patched_data = patch_spec(data)

    output_path = build_patched_path(version)
    with open(output_path, "w") as file:
        yaml.dump(patched_data, file, Dumper=yaml.CDumper, sort_keys=False)


def patch_spec(data):
    if "components" in data and "schemas" in data["components"]:
        for name, schema in data["components"]["schemas"].items():
            if "properties" in schema:
                # Fix non-nullable types
                # See: https://github.com/OpenAPITools/openapi-generator/issues/18006
                non_nullable_types = [
                    "front_image",
                    "rear_image",
                ]

                for ntype in non_nullable_types:
                    if ntype in schema["properties"]:
                        if (
                            schema["properties"][ntype]["format"] == "binary"
                            and "nullable" in schema["properties"][ntype]
                        ):
                            schema["properties"][ntype].pop("nullable")

            if "required" in schema:
                schema["required"] = [
                    prop
                    for prop in schema["required"]
                    if prop not in MISSING_REQUIRED_FIELDS
                ]

    return data


def main():
    """
    Entry method for CLI parsing
    """
    parser = argparse.ArgumentParser()
    parser.add_argument("version")
    parser.add_argument("-d", "--debug", action="store_true")
    parser.add_argument("-v", "--verbose", action="store_true")
    args = parser.parse_args()
    if args.debug:
        logging.basicConfig(level=logging.DEBUG)
    elif args.verbose:
        logging.basicConfig(level=logging.INFO)

    package_version = discover_package_version(args.version)
    patch_spec_file(package_version)


if __name__ == "__main__":
    main()
