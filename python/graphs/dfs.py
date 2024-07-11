# assumes no duplicate nodes (each node str is unique)
def traverse_dfs(graph: dict[str, list[str]], start: str) -> list[str]:
    res = []
    to_visit = []
    visited = set()

    if len(graph) < 1:
        return res

    if start not in graph:
        return [start]

    to_visit.append(start)

    while to_visit:
        cur = to_visit.pop()
        res.append(cur)
        visited.add(cur)
        for next in graph[cur]:
            if next not in visited and next not in to_visit:
                to_visit.append(next)
    return res
