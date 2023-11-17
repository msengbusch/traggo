export const ROUTER = {}

export interface Route {
    path: String;
    callback: Function;
}

export interface Match {
    route: Route;
}

// TODO: Rewrite with svelte runes
export class Router {
    routes: Route[] = [];

    public on(path: String, callback: Function): void {
        this.routes.push({path, callback});
    }

    public navigate(path: String): boolean {
        const match = this.resolve(path);
        if(match == null) return false;
        match.route.callback();
        return true
    }

    private resolve(path: String): null | Match {
        const route = this.routes.find(route => route.path === path);
        if(route == null) return null;
        const match: Match = {route};
        return match;
    }
}