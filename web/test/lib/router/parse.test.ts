import {expect, test} from "vitest";
import {Route} from "@lib/router/route";
import {normalize, resolve} from "@lib/router/parse";

test("normalize", t => {
  expect(normalize("path")).toBe("path");
  expect(normalize("#path")).toBe("path");
  expect(normalize("path/")).toBe("path");
  expect(normalize("path//")).toBe("path");
  expect(normalize("/path")).toBe("/path");
  expect(normalize("/path/")).toBe("/path");
  expect(normalize("/path//")).toBe("/path");
  expect(normalize("#/path")).toBe("/path");
  expect(normalize("#/path/")).toBe("/path");
  expect(normalize("#/path//")).toBe("/path");
  expect(normalize("#/path/child")).toBe("/path/child");
  expect(normalize("#/path/child/")).toBe("/path/child");
  expect(normalize("#/path/child//")).toBe("/path/child");
})

test("resolve", t => {
  const routes: Route[] = [
    {
      path: "/path1"
    },
    {
      path: "/path2"
    },
    {
      path: "/path3"
    }
  ]

  expect(resolve(routes, "/path2")).toEqual([routes[1]]);
})

test("resolve_children", t => {
  const routes: Route[] = [
    {
      path: "/path1"
    },
    {
      path: "/path2",
      children: [
        {
          path: "/path2/child1"
        },
        {
          path: "/path2/child2"
        }
      ]
    },
    {
      path: "/path3"
    }
  ]

  expect(resolve(routes, "/path2/child1")).toEqual([routes[1], routes[1].children[0]]);
})