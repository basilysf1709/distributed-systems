# Tabular comparison of key value stores

| Database | Type | Replication Technique | Consistency Protocol | Node Discovery | Partitioning |
|----------|------|----------------------|----------------------|----------------|--------------|
| Cassandra | Wide-column store | Leaderless | Quorum | Gossip Protocol | Hash-based |
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