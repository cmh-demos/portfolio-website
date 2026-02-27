# simple nginx-based frontend that serves the static index.html and assets

FROM nginx:stable-alpine
# copy the portfolio page and any local assets
COPY index.html /usr/share/nginx/html/index.html
# build context is repo root, so satellite file is under frontend/
COPY frontend/satellite.min.js /usr/share/nginx/html/satellite.min.js
