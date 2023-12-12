export type LazyPromise<T> = Promise<{ default: T }>;
export type Lazy<T> = T & { lazy: () => LazyPromise<T> };

export function lazy<T>(fn: () => LazyPromise<T>): Lazy<T> {
  const wrap: Lazy<T> = fn() as any;
  wrap.lazy = fn;
  return wrap;
}