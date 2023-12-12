import type {Route} from "./route";

export function normalize(path: string): string {
  path = path.startsWith("#") ? path.slice(1) : path;
  path = path.length > 1 && path.endsWith("/") ? path.slice(0, -1) : path;
  return path;
}

export function resolve(routes: Route[], path: string): Route[] {
/*  const split = normalize(path).split("/");

  function traverse(routes: Route[]): Route[] | undefined {
    for(let i = split.length; i >= 0; i--) {
      const path = split.slice(0, i).join("/");
      const route = routes.find(route => route.path == path);

      if(route) {
        if(i == split.length) {
          return [route];
        }

        if(route.children) {
          const children = traverse(route.children);

          if(children) {
            return [route, ...children];
          }
        }
      }
    }
  }

  return traverse(routes);*/
}