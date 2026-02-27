# Remaining Systems Engineering Tasks

Below are the recommended steps to further develop the shared reference
architecture and support new project development.

1. **Populate the templates**
   - Flesh out `backend/` with a Go service starter (health checks, config
     management, logging); retain Python example only as a secondary option.  
     *Basic project routing and CORS support have been added as a placeholder
     framework.*
     - Added aerospace project endpoints (`/projects/aerospace/data` and
       `/projects/aerospace/fetch`) with caching of TLE data from CelesTrak.
     - Implemented TLE-specific endpoints (`/projects/aerospace/tle` and
       `/projects/aerospace/tle/fetch`) and frontend orbit rendering using
       satellite.js.
   - Add a basic React app in `frontend/` that consumes the backend and
     demonstrates the shared components.

2. **Build and publish reusable modules**
   - Create real Terraform modules (`storage_bucket`, `cloud_function`,
     `iam_policy`) with examples and tests.
   - Package common Python utilities (data processing, auth middleware) in a
     pip-installable library.
   - Publish the React component library to npm or keep it as a git submodule.

3. **Standardise CI/CD**
   - Implement the workflows hinted at in `ci_cd/`: a `lint-and-test.yml` and
     `terraform.yml` with reusable jobs and inputs so every project can
     `workflow_call` them.
   - Add automatic versioning, image builds, and deployment steps.

4. **Documentation & onboarding**
   - Write setup guides (how to start a new project using the architecture).
   - Create architecture diagrams for each project variant.
   - Add a “cheat sheet” for common commands and branch/release procedures.

5. **Testing & validation**
   - Add unit/integration tests to the template backend and modules.
   - Configure a local development environment (e.g. Docker Compose) so
     projects can run end-to-end.

6. **Governance & maintenance**
   - Establish a branching/merge strategy (the `sys_eng` branch is a good
     start).
   - Define update procedures for shared modules and versioning policy.
   - Add a changelog or upgrade notes to help projects track changes.

7. **Tooling & automation**
   - Create a small CLI or script that scaffolds a new project from the
     template, injecting names and basic configuration.
   - Add monitoring/alerting templates that can be attached to any new
     service.

8. **Cross-project demos**
   - Build one or two concrete projects (e.g. the Aerospace Viz and the GCP
     Auditor) using the shared architecture to validate and refine it.

9. **Repository hygiene**
   - Add a top-level `.gitignore` for Python, Node, Terraform, and IDE files.
   - Include a `LICENSE` if the templates will be open-sourced.
   - Update the root `README.md` with usage instructions.

10. **Baseline tooling**
   - Provide a `pyproject.toml`/`setup.cfg` or `requirements.txt` in the
     backend template.
   - Add linting configuration (`.flake8`, `eslint` for frontend).

11. **Tests & CI**
   - Add a trivial `pytest` test hitting `/` in `app.py`.
   - Wire up an initial CI workflow to install deps and run the test using the
     templates under `ci_cd/`.

12. **Developer convenience**
   - Create a `Makefile` or simple startup scripts (Docker Compose) for local
     dev.

13. **Project scaffold**
   - Write a small CLI or shell script to copy the template into a new project
     directory and replace placeholders.

14. **Content consolidation**
   - Consolidate the individual project markdowns into a `projects/` section or
     reference them from the architecture docs to keep them organized.

15. **Versioning policy**
   - Add a `VERSION` file in `architecture/` and note bumping procedures when
     shared modules change.
