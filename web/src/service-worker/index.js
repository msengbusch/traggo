import { build, files, prerendered, version } from "$service-worker";
import { precacheAndRoute } from "workbox-precaching";

let cacheEntries = [...build, ...files, ...prerendered].map((url) => ({ url, revision: version }));
precacheAndRoute(cacheEntries);
