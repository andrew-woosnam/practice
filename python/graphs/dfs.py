# assumes no duplicate nodes (each node str is unique)
def traverse_dfs(graph: dict[str, list[str]], start: str) -> list[str]:
    res = []
    to_visit = [start]
    visited = set()

    if not graph:
        return res

    while to_visit:
        cur = to_visit.pop()
        if cur not in visited:
            res.append(cur)
            visited.add(cur)
            for adjacent in graph.get(cur, []):
                to_visit.append(adjacent)
    return res
