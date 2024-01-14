import { build, files } from "$service-worker";

import { precacheAndRoute } from "workbox-precaching";

precacheAndRoute([...build, ...files]);
