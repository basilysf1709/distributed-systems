# Tabular comparison of key value stores


## Table

| Database | Type | Replication Technique | Consistency Protocol | Node Discovery | Partitioning |
|----------|------|----------------------|----------------------|----------------|--------------|
| [MongoDB](https://www.mongodb.com/docs/manual/core/architecture-concepts) | Document store | Single-leader | Consensus (Raft) | Config Servers | Range-based and Hash-based |
| [Redis](https://redis.io/topics/architecture) | Key-value store | Single-leader | Eventual Consistency | Sentinel | Hash slots |
| [Valkey](https://valkey.io/) | Key-value store | Single-leader | Eventual Consistency | Sentinel | Hash slots |
| [Cassandra](https://cassandra.apache.org/doc/latest/architecture/overview.html) | Wide-column store | Leaderless | Quorum | Gossip Protocol | Hash-based |
| [DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.html) | Key-value & Document | Leaderless | Eventual Consistency (also offers strong consistency) | Proprietary | Consistent Hashing |
| [Couchbase](https://docs.couchbase.com/server/current/learn/architecture-overview.html) | Document store | Leaderless | MVCC | Gossip Protocol | Hash-based |
| [Neo4j](https://neo4j.com/docs/operations-manual/current/architecture/) | Graph database | Causal Clustering | Raft | Core Discovery | Graph partitioning |
| [Cosmos DB](https://learn.microsoft.com/en-us/azure/cosmos-db/distribute-data-globally) | Multi-model | Leaderless (and others) | Tunable Consistency | Proprietary | Hash-based |
| [HBase](https://hbase.apache.org/book.html#architecture) | Wide-column store | Single-leader | Consensus (ZooKeeper) | ZooKeeper | Range-based |
| [Cloud Firestore](https://firebase.google.com/docs/firestore/data-model) | Document store | Multi-region | Strong Consistency | Proprietary | Automatic |
| [CouchDB](https://docs.couchdb.org/en/stable/intro/overview.html) | Document store | Leaderless | MVCC | Cluster-aware | Hash-based |
| [ScyllaDB](https://docs.scylladb.com/stable/architecture/) | Wide-column store | Leaderless | Tunable Consistency | Gossip Protocol | Consistent Hashing |
| [Riak](https://docs.riak.com/riak/kv/latest/learn/concepts/index.html) | Key-value store | Leaderless | Vector Clocks | Gossip Protocol | Consistent Hashing |
| [Fauna](https://docs.fauna.com/fauna/current/get-started/overview/#distributed-service) | Document-relational store | Multi-region Consensus (Raft+Calvinesque) | Strongly consistent | Proprietary | Consistent Hashing |
| [ETCD3](https://etcd.io/docs/v3.4/learning/architecture/) | Key-value store | Multi-leader | Consensus (Raft) | Self-discovery | Hash-based |
| [ZooKeeper](https://zookeeper.apache.org/doc/r3.5.9/zookeeperOver.html) | Coordination service | Single-leader | Consensus (Zab) | Dynamic Discovery | N/A |
---

### To be Noted
1. **Redis:** Redis Cluster uses a different model than Redis with Sentinel.
2. **DynamoDB:** Offers both eventual and strong consistency options.
3. **Couchbase:** Provides tunable consistency levels, not just MVCC.
4. **Cosmos DB:** Replication method varies based on the chosen consistency level and API.
5. **Cloud Firestore:** Multi-region setup functions as leaderless replication.
6. **Neo4j:** Causal Clustering is specific to the enterprise edition.

## History

```mermaid
graph TD
    A[Start] --> B[Fundamentals]
    B --> C[Designing Data-Intensive Applications Book]
    B --> D[MIT Distributed Systems Course]
    B --> E[Martin Kleppmann's Distributed Systems Lecture Series]

    C --> F[Data Models and Storage]
    C --> G[Replication and Partitioning]
    C --> H[Transactions and Consistency]

    D --> I[RPC and Threads]
    D --> J[Fault Tolerance]
    D --> K[Consensus Algorithms]

    F --> L[Practical Learning]
    G --> L
    H --> L
    I --> L
    J --> L
    K --> L

    L --> M[Kubernetes Basics]
    L --> N[Key-Value Stores Comparison]

    M --> O[Advanced Topics]
    N --> O

    O --> P[Landmark Papers]
    O --> Q[Distributed Algorithms]
    O --> R[System Design Concepts]

    P --> S[Google File System]
    P --> T[MapReduce]
    P --> U[BigTable]
    P --> V[Dynamo]
    P --> W[Cassandra]
    P --> X[Spanner]
    P --> Y[F1]
    P --> Z[Borg]
    P --> AA[Dapper]
    P --> AB[Chubby]
    P --> AC[Zanzibar]
    P --> AD[Shard Manager]
    P --> AE[Monarch]

    Q --> AF[Paxos]
    Q --> AG[Raft]
    Q --> AH[Vector Clocks]
    Q --> AI[Two-Phase Commit]
    Q --> AJ[Three-Phase Commit]
    Q --> AK[Byzantine Fault Tolerance]

    R --> AL[SAGA Pattern]
    R --> AM[Transactional Outbox]
    R --> AN[Workflow Engines]
    R --> AO[Event Sourcing]
    R --> AP[CQRS]

    O --> AQ[Database Evolution]
    AQ --> AR[Relational Databases]
    AR -->|Limitations led to| AS[NoSQL Movement]
    AS --> AT[Document Stores]
    AS --> AU[Key-Value Stores]
    AS --> AV[Column-Family Stores]
    AS --> AW[Graph Databases]

    AT --> AX[MongoDB]
    AT --> AY[CouchDB]
    AU --> AZ[Redis]
    AU --> BA[Riak]
    AV --> BB[Cassandra]
    AV --> BC[HBase]
    AW --> BD[Neo4j]
    AW --> BE[OrientDB]

    BF[CAP Theorem] --> AS
    BG[Web 2.0 Era] --> AS
    BH[Big Data Challenges] --> AS

    AS --> BI[NewSQL Movement]
    BI --> BJ[Google Spanner]
    BI --> BK[CockroachDB]
    BI --> BL[TiDB]

    O --> BM[Streaming Systems]
    BM --> BN[Kafka]
    BM --> BO[Flink]
    BM --> BP[Spark Streaming]

    O --> BQ[Machine Learning in Distributed Systems]
    BQ --> BR[Distributed ML Model Training]
    BQ --> BS[Federated Learning]

    O --> BT[Chaos Engineering]
    O --> BU[RDMA (Remote Direct Memory Access)]

    O --> BV[Formal Methods]
    BV --> BW[TLA+]

    O --> BX[Metrics and Monitoring]
    BX --> BY[Prometheus]
    BX --> BZ[Grafana]

    O --> CA[Tracing]
    CA --> CB[Jaeger]
    CA --> CC[Zipkin]

    O --> CD[Serverless Computing]
    CD --> CE[AWS Lambda]
    CD --> CF[Azure Functions]
    CD --> CG[Google Cloud Functions]

    O --> CH[Containerization]
    CH --> CI[Docker]
    CH --> CJ[Kubernetes]

    O --> CK[Service Mesh]
    CK --> CL[Istio]
    CK --> CM[Linkerd]

    O --> CN[Blockchain and Distributed Ledgers]
    CN --> CO[Bitcoin]
    CN --> CP[Ethereum]

    P --> CQ[Attention Is All You Need]
    P --> CR[FoundationDB]
    P --> CS[Amazon Aurora]
    P --> CT[Thrift]
    P --> CU[MyRocks]
    P --> CV[WTF: Twitter's Who To Follow]

    Q --> CW[LCR Algorithm]
    Q --> CX[HS Algorithm]
    Q --> CY[TimeSlice Algorithm]
    Q --> CZ[FloodMax Algorithm]
    Q --> DA[Bellman-Ford Algorithm]
    Q --> DB[GHS Algorithm]
    Q --> DC[FloodSet Algorithm]
    Q --> DD[Exponential Information Gathering]

    R --> DE[Vector Databases]
    R --> DF[In-Memory File Systems]

    BX --> DG[Recovery Strategies]
    BX --> DH[Data Quality Checking]

    DI[System Design Interviews] --> DJ[Continued Learning and Practice]
    O --> DI
```
