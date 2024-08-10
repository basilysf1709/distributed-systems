# 🌐 Distributed Systems

<p align="center">
  <img src="https://img.shields.io/badge/status-in_progress-brightgreen" alt="Status: In Progress"/>
  <img src="https://img.shields.io/badge/last_update-2024--07--25-blue" alt="Last Update"/>
  <img src="https://img.shields.io/badge/contributions-welcome-orange" alt="Contributions Welcome"/>
</p>

## 📚 The Absolute Best Book: Designing Data-Intensive Applications

<img width="389" alt="Screenshot 1446-01-19 at 2 04 02 PM" src="https://github.com/user-attachments/assets/19a57547-30c1-4f41-8552-586647dbd23d">

### Chapters

1. [Reliable, Scalable, and Maintainable Applications](https://github.com/basilysf1709/distributed-systems/tree/main/designing-data-intensive-applications/Ch1)
2. [Data Models and Query Languages](https://github.com/basilysf1709/distributed-systems/tree/main/designing-data-intensive-applications/Ch2)
3. [Storage and Retrieval](https://github.com/basilysf1709/distributed-systems/tree/main/designing-data-intensive-applications/Ch3)
4. [Encoding and Evolution](https://github.com/basilysf1709/distributed-systems/tree/main/designing-data-intensive-applications/Ch4)
5. [Replication](https://github.com/basilysf1709/distributed-systems/tree/main/designing-data-intensive-applications/Ch5)
6. [Partitioning](https://github.com/basilysf1709/distributed-systems/tree/main/designing-data-intensive-applications/Ch6)
7. [Transactions](https://github.com/basilysf1709/distributed-systems/tree/main/designing-data-intensive-applications/Ch7)
8. [The Trouble with Distributed Systems](https://github.com/basilysf1709/distributed-systems/tree/main/designing-data-intensive-applications/Ch8)
9. [Consistency and Consensus](https://github.com/basilysf1709/distributed-systems/tree/main/designing-data-intensive-applications/Ch9)
10. [Batch Processing](https://github.com/basilysf1709/distributed-systems/tree/main/designing-data-intensive-applications/Ch10)
11. [Stream Processing](https://github.com/basilysf1709/distributed-systems/tree/main/designing-data-intensive-applications/Ch11)
12. [The Future of Data Systems](https://github.com/basilysf1709/distributed-systems/tree/main/designing-data-intensive-applications/Ch12)

📖 [Access the full book here](https://github.com/user-attachments/files/16344190/Designing.Data.Intensive.Applications.pdf)

---

## 🎓 MIT's Distributed Systems Course

🔗 [Watch the full course](https://www.youtube.com/watch?v=cQP8WApzIQQ&list=PLrw6a1wE39_tb2fErI4-WkMbsvGQk9_UB)

---

## 📘 High-Quality Resources

- 📑 [Dynamo: Amazon's Highly Available Key-value Store](https://www.allthingsdistributed.com/files/amazon-dynamo-sosp2007.pdf)
- 🛢 [Tabular Comparison of Key Value Stores](https://github.com/basilysf1709/distributed-systems/tree/main/databases)

---

## 💻 Practical Learnings

- [✅ Kubernetes: A Basic Overview](https://www.youtube.com/watch?v=X48VuDVv0do)
- [✅ Systems Design Interviews - Alex Xu](https://github.com/basilysf1709/distributed-systems/tree/main/system-design-interviews/Summary)

---

## 💻 Algorithms

- [Consistent Hashing in Zig](https://github.com/basilysf1709/distributed-systems/tree/main/algorithms/consistent-hashing)
- [Hinted Handoff in Zig](https://github.com/basilysf1709/distributed-systems/tree/main/algorithms/hinted-handoff)

---

## 📊 Learning Roadmap

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

---

## Star History

<picture>
  <source media="(prefers-color-scheme: dark)" srcset="https://api.star-history.com/svg?repos=basilysf1709/distributed-systems&type=Date&theme=dark" />
  <source media="(prefers-color-scheme: light)" srcset="https://api.star-history.com/svg?repos=basilysf1709/distributed-systems&type=Date" />
  <img alt="Star History Chart" src="https://api.star-history.com/svg?repos=star-history/star-history&type=Date" />
</picture>

<p align="center">
  <i>Happy learning! May your distributed systems knowledge scale infinitely! 🚀</i>
</p>
