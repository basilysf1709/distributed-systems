# Tabular comparison of key value stores


## Table

| Database | Type | Replication Technique | Consistency Protocol | Node Discovery | Partitioning |
|----------|------|----------------------|----------------------|----------------|--------------|
| [Cassandra]() | Wide-column store | Leaderless | Quorum | Gossip Protocol | Hash-based |
| MongoDB | Document store | Single-leader | Consensus (Raft) | Config Servers | Range-based |
| Riak | Key-value store | Leaderless | Vector Clocks | Gossip Protocol | Consistent Hashing |
| CouchDB | Document store | Multi-master | MVCC | Cluster-aware | Hash-based |
| DynamoDB (AWS) | Key-value & Document | Multi-leader | Eventual Consistency | Proprietary | Consistent Hashing |
| HBase | Wide-column store | Single-leader | Consensus (ZooKeeper) | ZooKeeper | Range-based |
| Couchbase | Document store | Multi-master | MVCC | Gossip Protocol | Hash-based |
| Redis | Key-value store | Single-leader | Eventual Consistency | Sentinel | Hash slots |
| Neo4j | Graph database | Causal Clustering | Raft | Core Discovery | Graph partitioning |
| Cosmos DB (Azure) | Multi-model | Multi-master | Tunable Consistency | Proprietary | Hash-based |
| Cloud Firestore (Google) | Document store | Multi-region | Strong Consistency | Proprietary | Automatic |
| InfluxDB | Time series | Single-leader | Eventual Consistency | Meta nodes | Time-based sharding |
| ArangoDB | Multi-model | Single-leader | Eventual Consistency | Agency (Raft-based) | Hash-based |
| ScyllaDB | Wide-column store | Leaderless | Tunable Consistency | Gossip Protocol | Consistent Hashing |
| Elasticsearch | Search engine & Document store | Primary-backup | Eventual Consistency | Zen Discovery | Hash-based |

---

## History

<div class="mermaid" style="font-size: 20px;">
graph TD
    A[Relational Databases] -->|Limitations led to| B[NoSQL Movement]
    B --> C[Document Stores]
    B --> D[Key-Value Stores]
    B --> E[Column-Family Stores]
    B --> F[Graph Databases]
    C --> G[MongoDB 2009]
    C --> H[CouchDB 2005]
    D --> I[Redis 2009]
    D --> J[Riak 2009]
    E --> K[Cassandra 2008]
    E --> L[HBase 2008]
    F --> M[Neo4j 2007]
    F --> N[OrientDB 2010]
    O[Google's BigTable 2006] --> K
    O --> L
    P[Amazon's Dynamo 2007] --> J
    Q[Facebook's Cassandra 2008] --> K
    R[CAP Theorem 2000] --> B
    S[Web 2.0 Era] --> B
    T[Big Data Challenges] --> B
    U[NewSQL Movement] --> V[Google Spanner 2012]
    U --> W[CockroachDB 2014]
    B --> U
</div>

