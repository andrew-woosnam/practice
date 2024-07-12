# assumes no duplicate nodes (each node str is unique)
from collections import deque


def traverse_bfs(graph: dict[str, list[str]], start: str) -> list[str]:
    res = []
    visited = set()
    to_visit = deque()
    to_visit.append(start)

    if not graph:
        return []

    if start not in graph:
        return [start]

    while to_visit:
        cur = to_visit.popleft()
        if cur not in visited:
            res.append(cur)
            visited.add(cur)
            for adjacent in graph.get(cur, []):
                to_visit.append(adjacent)

    return res
