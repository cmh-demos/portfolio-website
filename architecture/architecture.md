# System Engineering Reference Architecture

Below is a high-level architecture diagram showing the shared components that
all demo projects will consume.

```mermaid
flowchart LR
    subgraph Shared
        A[Backend Template]
        B[Frontend Components]
        C[Terraform Modules]
        D[CI/CD Workflows]
        E[Monitoring & Logging]
    end

    subgraph Project1[Aerospace Viz]
        P1A[FastAPI Service]
        P1B[React UI]
    end
    subgraph Project2[Monitoring Dashboard]
        P2A[FastAPI Service]
        P2B[React UI]
    end
    subgraph Project3[GCP Auditor]
        P3A[FastAPI Service]
        P3B[React UI]
    end

    A --> P1A
    B --> P1B
    C --> P1A
    D --> P1A
    D --> P1B

    A --> P2A
    B --> P2B
    C --> P2A
    D --> P2A
    D --> P2B

    A --> P3A
    C --> P3A
    D --> P3A

    E --> P1A
    E --> P2A
    E --> P3A
```
```

Each project imports or extends the shared modules instead of duplicating code.

