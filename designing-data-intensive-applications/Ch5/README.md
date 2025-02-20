# Chapter 5: Replication

## Key Concepts

### Types of Replication
1. **Single-Leader (Active/Passive)**
   - One leader accepts writes
   - Followers only handle reads
   - Common in: MySQL, PostgreSQL, MongoDB

2. **Multi-Leader (Master/Master)**
   - Multiple nodes accept writes
   - Best for geo-distributed scenarios
   - Requires conflict resolution

3. **Leaderless**
   - All replicas accept writes
   - Uses quorum for consistency
   - Examples: Cassandra, DynamoDB

### Replication Characteristics

#### Synchronization Modes
- **Synchronous**: Guaranteed consistency but higher latency
- **Asynchronous**: Better performance but potential data loss
- **Semi-synchronous**: Practical compromise (one sync follower)

#### Consistency Challenges
- Read-after-write consistency
- Monotonic reads
- Consistent prefix reads
- Replication lag management

### Failure Handling
- **Follower Failure**: Catch-up recovery
- **Leader Failure**: Failover process
  - Leader detection
  - Leader election
  - System reconfiguration

### Trade-offs
| Type | Consistency | Availability | Latency |
|------|-------------|--------------|---------|
| Single-Leader | Strong | Medium | Medium |
| Multi-Leader | Eventually | High | Low |
| Leaderless | Eventually | High | Low |

## Best Practices
1. Choose replication strategy based on:
   - Geographic distribution
   - Consistency requirements
   - Latency tolerance
   - Failure tolerance needs

2. Consider implementing:
   - Proper monitoring
   - Automated failover
   - Conflict resolution strategies
   - Regular replication testing
