# System Engineering Reference Architecture

Below is a high-level architecture diagram showing the shared components that
all demo projects will consume.

```mermaid
flowchart LR
    subgraph Shared
        A[Backend Template (Go preferred)]
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

## Local Development with Docker Compose

A `docker-compose.yml` has been added at the repository root to simplify
bringing up the backend service (and eventually others) for local testing.

```bash
# build images and start all services defined in compose
cd /home/mike/junk/my_website
docker compose up --build
```

The backend has been expanded into a minimal framework that can power all of
the example projects.  When running locally it exposes several endpoints:

- `GET /health` – simple health check
- `GET /projects` – returns a JSON array of supported project slugs
- `GET /projects/{slug}` – returns a placeholder description for the named
  project

These endpoints currently emit hard‑coded placeholder data; in a real
implementation each project would live in its own package or subservice and
provide real logic.  For example, the `aerospace` project already supports

- fetching JSON or TLE data by satellite name from the CelesTrak API
- caching both formats in `aerospace_cache.json`
- exposing `/projects/aerospace/data` and `/projects/aerospace/tle` for the
  cached values plus `/fetch` variants that refresh from the upstream API

The frontend uses the TLE endpoint with [satellite.js](https://github.com/shashwatak/satellite-js)
to draw a simple orbit path on a canvas whenever the Aerospace project is
selected.  This gives you a starting point for parsing and visualizing real
orbital elements in the browser.

The frontend has been updated to consume this API.  When you open
<http://localhost/> it will fetch the project list and render clickable
entries; clicking one pops open the dummy description.  CORS headers are
allowed on the backend so the static site (served by nginx) can call the API
on port 8080.

After startup you can also hit the backend directly at <http://localhost:8080/>.
Shutting down is as simple as `docker compose down`.

The compose file currently contains only the `backend` service but you can
extend it with a frontend, database, or any other components as the project
grows. See the comments in `docker-compose.yml` for an example of a
PostgreSQL service.


