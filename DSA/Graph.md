# Graph

* Graph Representation
* Graph is very Easy
* What is adjacency Matrix

### Adjacency Matrix

An **adjacency matrix** is a 2D table used to represent a graph.  
If there is an edge between two vertices, the value is `1`, otherwise `0`.

### Example Graph
Vertices: A, B, C, D

### Adjacency Matrix Table

|   | A | B | C | D |
|---|---|---|---|---|
| A | 0 | 1 | 1 | 0 |
| B | 1 | 0 | 0 | 1 |
| C | 1 | 0 | 0 | 1 |
| D | 0 | 1 | 1 | 0 |

**Meaning:**
- A → B and A → C are connected
- B → D is connected
- C → D is connected
- 0 means no edge, 1 means edge exists


### Adjacency List

An **adjacency list** is a way to represent a graph where each vertex stores a list of its connected vertices.

It is more memory-efficient than an adjacency matrix for sparse graphs.

---

### Example Graph
Vertices: A, B, C, D

---

### Adjacency List Representation

- A → B, C  
- B → A, D  
- C → A, D  
- D → B, C  

---

### Adjacency List (Table Form)

| Vertex | Connected Vertices |
|--------|--------------------|
| A      | B, C               |
| B      | A, D               |
| C      | A, D               |
| D      | B, C               |

---

### Properties

- Stores only existing edges  
- Uses less memory than adjacency matrix  
- Easy to traverse neighbors of a node  
- Best for sparse graphs  

---

### Key Points

- Each node has a list of adjacent nodes  
- Undirected graph stores connections both ways  
- Directed graph stores only outgoing edges  

---

> An adjacency list represents a graph as a collection of lists, where each list describes the neighbors of a vertex.

Unordered map <int, vector of int >

ajd[1] -> {2, 5}


```C++
DFS(adj, 0, visited);

void DFS(ajd, int u, visited){
    if(visited[u] == true) return;
    visited[u] = true
    for(int &v : adj[u]){
        if(!visited[u]){
            DFS(adj, v visited);
        }
    }
}
```

### Final Code For DFS!!

```C++
class Solution {
  public:
  
  void DFS( vector<vector<int>> &adj, int u, vector<int> &result, vector<int> &visted){
      if(visted[u])return;
      
      visted[u] = true;
      result.push_back(u);
      
      for(auto t : adj[u]){
          if(!visted[t]){
              DFS(adj, t, result, visted);
          }
      }
  }
  
  
    vector<int> dfs(vector<vector<int>> &adj) {
        vector<int> result, visted(adj.size(), 0);
        DFS(adj, 0, result, visted);
        return result;
    }   
};
```

