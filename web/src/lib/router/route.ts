export interface Route {
  path: string;
  children?: Route[];
}