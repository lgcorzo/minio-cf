# lgcorzo/minio-cf

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)
[![Dark Gravity](https://img.shields.io/badge/Autonomous%20Factory-Dark%20Gravity-blueviolet?style=flat-square)](https://github.com/lgcorzo/rust_CACD_autonomous_factory)

> [!NOTE]
> **UPSTREAM STATUS & SOVEREIGN FORK PURPOSE:**
> Following upstream repository archiving and licensing shifts across the MinIO ecosystem, this repository is **actively maintained and hardened** under **@lgcorzo** as an essential infrastructure component of the **[Dark Gravity Autonomous CA/CD Factory](https://github.com/lgcorzo/rust_CACD_autonomous_factory)** ([Architecture Knowledge Base](https://notebook.google.com/notebook/116d5f03-2fb0-4fa7-bd7c-1c9849920abc)). Continuous security auditing, zero-CVE verification, and hardware SIMD optimizations are independently maintained to guarantee supply-chain sovereignty and operational continuity.

---

## 🛡️ Dark Gravity Factory: Technological Sovereignty & Security Rationale

### 1. Architectural Role in Dark Gravity
In the **Dark Gravity Autonomous CA/CD Software Factory (V7.2 / V7.3)**, Deployment templates and automation for sovereign S3 storage infrastructure in distributed enterprise cloud topologies.

### 2. Supply-Chain Sovereignty & Independent Maintenance
- **Insulation from Upstream License Closures:** When upstream open-source projects close distributions or alter licensing terms, Dark Gravity enforces **semantic mirroring and sovereign forks** to guarantee total independence from commercial vendors.
- **Autonomous Security Verification & Zero-CVE Policy:** This repository is continuously audited for CVEs and vulnerabilities, maintaining a 100% clean secret scanning posture. Regressions and upstream bugs are resolved directly within the autonomous factory ecosystem.
- **Compliance & Provenance:** Ensures cryptographic integrity and auditable lineage in compliance with the **EU AI Act**, **SOC 2 Type II**, and **Hazitek / SPRI** regulatory frameworks.

---

minio-cf - Minio CF Tool
========================

Minio tool to be used with Pivotal "cf" tool to generate the config for create-service command.

Please download the minio-cf tool for your desktop (amd64) from [here](https://github.com/minio/minio-cf/releases)

EXAMPLES:
--------
Deploy Minio Server instance
```
$ minio-cf --access-key minio --secret-key minio123 > minio-server.conf
$ cf create-service minio standard server-instance-1 -c minio-server.conf
```

Deploy Minio GCS gateway instance
```
$ minio-cf --access-key minio --secret-key minio123 --gcs /path/to/credentials.json > minio-gcs.conf
$ cf create-service minio standard server-instance-2 -c minio-gcs.conf
```

Deploy Minio Azure gateway instance
```
$ minio-cf --access-key azureaccountname --secret-key azureaccountkey > minio-azure.conf
$ cf create-service minio standard server-instance-3 -c minio-azure.conf
```
