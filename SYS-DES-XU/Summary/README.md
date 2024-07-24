# System Design Interview - Alex Wu

Alex Xu:

- System Design questions essentially are open ended

Ch_1: Scale from Zero to Millions:

- Single Server Setup
    - Traffic goes through mobile or website to our server
- As more users come
    - Database is required
    - SQL or NoSQL
        - SQL
            - Relational Data
            - Been here for 40 years
        - NoSQL
            - Super low latency
            - Good for unstructured data
            - Good for storing massive amounts of data
- Vertical Scaling vs Horizontal Scaling
    - Vertical
        - Scale up
        - Add more power (CPU, RAM)
        - good for low traffic
        - there’s a cap to how much CPU/RAM power can be added
        - if one server goes down, you are done for (not highly available)
    - Horizontal
        - Scale out
        - Add more servers
- Load Balancer
    - evenly distributes traffic to different server
- Database replication
    - master/slave relationship
    - usually more reads are required compared to writes
    - so there are more slaves than masters
    - masters:
        - write operations
    - slaves:
        - read operations (replicated from masters)
    - Advantages of database replication
        - High performance
        - Reliability
        - High Availability
    - case 1:
        - slave db goes offline
        - all operations are redirected to master (read and write) until a new slave comes up
    - case 2:
        - master db goes offline
        - slave db is promoted to master
        - this is complicated as there is a case where data is not replicated to the slave and it could be promoted to master
        - ^ to counteract this, circular replication or multi master models are used
            
<img width="625" alt="Screenshot 1446-01-18 at 8 53 19 PM" src="https://github.com/user-attachments/assets/2343227d-d08f-4c4f-af62-5592aaf0dbd9">

- Cache
    - The cache tier is a temporary data store layer, much faster than the database
    - used for reducing database workloads, by adding a layer between db and server for frequent and similar reads
    - Considerations for cache:
        - data is read frequently but modified infrequently
        - consistency: make it consistent, especially when it scales
    - SPOF (single point of failure)
        - dont have this, multiple cache servers avoid this problem
    - Select Cache Eviction Policy, when cache is full
        - LRU
        - LFU
        - FIFO
    - depends on use case ^
- CDN (Content Delivery Network)
    - geographically dispersed servers to deliver static content
    - TTL: Time to Live in the cache
    - Considerations:
        - Cost
        - Setting Cache Expiry
        - CDN Fallback: outage should be handled by making calls to the original server
        - Invalidating and versioning
            - retrieve and update the versions through the API CDN vendors provide
- Shared data store for session data
- Data centres for geographical scaling
    - multiple data centres

- Scalable architecture: (before data centres for geo routing)
    
<img width="678" alt="Screenshot 1446-01-18 at 8 53 51 PM" src="https://github.com/user-attachments/assets/666072b5-06c0-411c-bf1f-a2a94e00136e">

    
- Technical Challenges for data centers:
    - Traffic redirection:
        - GeoDNS
    - Data Synchronization
    - Test and Deployment
        - CI/CD
- Message Queue
    - durable component that supports asynchronous communication
- Logging, metrics and automation
    - for scalable and big architectures
- Database Scaling
    - horizontal: add more databases
        - Sharding
            - Use hash function to find data across many shards
            - Sharding key is selected based on even distribution
            - There are drawbacks to sharding:
                - shard exhaustion: one shard gets exhausted
                - celebrity problem (one shard gets loading with a celebrity like justin beiber loool)
                - join and de-normalization
    - vertical: add more CPU/RAM to the one database you have

Ch_2: Back-of-the-envelope estimation:

- 2^(10n) - data layer calculation
- Memory is fast but the disk is slow
- Avoid disk seeks if possible
- Simple compression algorithms are fast
- Compress data before sending it over to the internet
- Commonly asked back-of-the-envelope estimations: QPS, peak QPS, storage, cache, number of servers, etc. You can practice these calculations when preparing for an interview. Practice makes perfect.

Ch_3: A Framework for System Design Interviews:

- Step 1: Understand the problem and establish design scope
- Step 2: Propose High level design and get buy-in
- Step 3: Design Deep Dive
- Step 4: Wrap up
    - system bottlenecks
    - recap
    - alternative solutions (possibilities)

Ch_4: Design a Rate Limiter:

- Limit the amount of API calls a user can do
- Advantages:
    - Prevents DDos attacks (flooding the server)
    - Reduce costs
    - Prevents servers from being overloaded
- Designing a Rate Limiter
    - To Ask:
<img width="576" alt="Screenshot 1446-01-18 at 8 54 12 PM" src="https://github.com/user-attachments/assets/763951c7-ba7a-47ec-95ce-fa78f2528889">



- You can have a:
    - Rate limiter in middleware
    - Rate limiter in server
    - Rate limiter in client (unreliable and not secure)
- Token Bucket Algorithm:
    - Tokens are put in a bucket
        - If there are enough tokens, we take one token out for each request, and the request goes through.
        - If there are not enough tokens, the request is dropped.
        - Takes two parameters:
            - Bucket size
            - Refill rate
- Leaky bucket algorithm
    - When a request arrives, the system checks if the queue is full. If it is not full, the request
    is added to the queue.
    - Otherwise, the request is dropped. Requests are pulled from the queue and processed at regular intervals.
        - Takes two parameters:
            - Bucket size
            - outflow rate
        - Cons:
            - A burst of traffic fills up the queue with old requests, and if they are not processed in time, recent requests will be rate limited.
- Fixed window counter algorithm
    - Within 1 minute window only n requests go through
    - Flaw
        - example, the system allows a maximum of 5 requests per minute, and the available quota resets at the human-friendly round minute. As seen, there are five requests between 2:00:00 and 2:01:00 and five more requests between 2:01:00 and 2:02:00. For the one-minute window between 2:00:30 and 2:01:30, 10 requests go through. That is twice as many as allowed requests.
- Sliding window log algorithm
    - This basically has the request time stamps of previous request, usually this data is stored in cache like redis.
    - if 2 requests per minute is allowed, only that will pass through, solves the problem with having a unit time for requests like in Fixed window counter algorithm
    - Pros:
        - Extremely accurate rate limit
    - Cons:
        - Needs a cache, memory (accurate, but requires space data)
- Sliding window counter algorithm
    - Takes the average rate of previous window and this window and checks for overlap
    - Pros:
        - smooths out traffic
        - memory efficient
    - **Cons:**
        - based on the assumption that traffic in previous request is somewhat evenly distributed
        - not as bad tho, Cloudflare said only 0.003% inaccuracy in 400 million requests
- High Level Architecture:
    - Redis is used as it is an in-memory cache
    ****
        
        <img width="654" alt="Screenshot 1446-01-18 at 8 54 33 PM" src="https://github.com/user-attachments/assets/f1521e54-c6af-43c4-9ba3-5f045314c797">

    - General design ^

- Rate Limiter in a distributed environment:
    - Challenges:
        - Race condition
            - Two threads could be carrying the same increment counter, but to make sure proper write is in place, we need locks
            - but locks slow down the system
        - Synchronization issue
            - When you scale you need multiple rate limiters, but synchronization is not achieved in rate limiters
            - Two solutions:
                - Make the rate limiters sticky
                    - Bad solution: not scalable, defeats the purpose
                - Use a centralized data store like Redis for consistent data across all rate limiters
    - One common thing is:
        - Make sure to have many servers across different places for low latency and performance optimization
    

Ch_5: Design consistent hasing:

- What is the hashing problem?
    - Major problem that occurs with rehashing is:
        - When a server goes down, you need to do the entire modular operation again and this results in uneven distribution of servers
    - To solve the above problem you do consistent hashing:
        - We use a hash ring
        - Go in clockwise direction
        - First key that is encountered gets mapped to the server
        - In removing/adding we only have to do at 1 operation for remapping the keys to servers
        - Again two issues:
            - It could be possible the multiple keys are distributed in a way that they are mapped to only server
            - To solve this you have virtual nodes:
            - s_0_0, s_0_1, ⇒ both means Server 0
            
            <img width="627" alt="Screenshot 1446-01-18 at 8 55 00 PM" src="https://github.com/user-attachments/assets/4be107bd-024c-4b9f-9f83-fa6c847e4aa8">

            
            - The standard deviation will be smaller when we increase the number of virtual nodes. However, more spaces are needed to store data about virtual nodes. This is a tradeoff, and we can tune the number of virtual nodes to fit our system requirements. (imp)
            - For redistribution you go anti clockwise
            - Pros of consistent hashing:
                - Minimized keys are redistributed when servers are added or removed.
                - It is easy to scale horizontally because data are more evenly distributed.

Ch_6: Design a key value store:

- Single server key-value store:
    - fast, but good for scalability as it reaches capacity pretty quickly
- Distributed server key-value store:
    - CAP Theorem:
<img width="709" alt="Screenshot 1446-01-18 at 8 55 18 PM" src="https://github.com/user-attachments/assets/8286737c-5922-48d4-ba0e-32218434ef42">


- CA cant be possible in real world because network failures are bound to happen
- Real world Distributed Systems:
    - CP System: (Consistency)
        - block all write operations when failure occurs
    - AP System: (Availability)
        - Keeps accepting Read

<img width="647" alt="Screenshot 1446-01-18 at 8 55 38 PM" src="https://github.com/user-attachments/assets/f9ba08dc-9a1f-4f11-ae42-5ac294c9e3aa">


- Strong consistency: any read operation returns a value corresponding to the result of the
most updated write data item. A client never sees out-of-date data.
- Weak consistency: subsequent read operations may not see the most updated value.
- Eventual consistency: this is a specific form of weak consistency. Given enough time, all
updates are propagated, and all replicas are consistent.
- Dynamo and Cassandra use Eventual Consistency. Netflix uses Dynamo
- High amounts of replication means system is highly available but might cause inconsistencies
    - To mitigate this problem you have versioning (v1 and v2)
    - vector clocks
        - used to add versions to read and writes within different servers
        - detects conflicts and has conflict resolution logic
        - downside:
            - conflict resolution logic required in clients side
            - (server, version) combo grows rapidly.
                - to fix this, we have a thresh hold for length
- Gossip Protocol Algorithm
    - each node has member IDs and heartbeat counters
    - Each node maintains a node membership list, which contains member IDs and heartbeat
    counters.
    - Each node periodically increments its heartbeat counter.
    - Each node periodically sends heartbeats to a set of random nodes, which in turn
    propagate to another set of nodes.
    - Once nodes receive heartbeats, membership list is updated to the latest info.
    - If the heartbeat has not increased for more than predefined periods, the member is
    considered as offline
- Hinted Handoff Approach for dealing with temporary failures
    - s2 goes down, s3 handles responsibility
    - s2 comes back, s3 hands off data to s2
- Permanent Failure
    - anti entropy protocol, merkle hash trees
- System Architecture Diagram:
    - Write Path
    - Read Path [Check diagrams in book]

![Diagram](https://prod-files-secure.s3.us-west-2.amazonaws.com/70c3badf-7b78-4bb0-9263-f7c8e5e66fec/985d1da1-e78a-4853-897d-222be91bf7ad/Untitled.png)

Ch_7: Design a Unique ID Generator in Distributed Systems:

- Four popular approaches:
    - Multi-master replication
        - increase by k (k is number of database servers)
        - However, this strategy has some major drawbacks:
            - Hard to scale with multiple data centers
            - IDs do not go up with time across multiple servers.
            - It does not scale well when a server is added or removed.
    - Universally unique identifier (UUID)
    
    ![Diagram](https://prod-files-secure.s3.us-west-2.amazonaws.com/70c3badf-7b78-4bb0-9263-f7c8e5e66fec/c42d84d7-8ab6-49d0-ac12-12a59512eccc/Untitled.png)
    
    - Ticket server
    
    ![Diagram](https://prod-files-secure.s3.us-west-2.amazonaws.com/70c3badf-7b78-4bb0-9263-f7c8e5e66fec/259d4201-33ed-4de2-904b-ec41d6eea548/Untitled.png)
    
    - Twitter snowflake approach
    
    ![Diagram](https://prod-files-secure.s3.us-west-2.amazonaws.com/70c3badf-7b78-4bb0-9263-f7c8e5e66fec/f5ae985e-602a-4331-a299-541a3b6d6e36/Untitled.png)
    
    - will fail after 69 years
- should look into these for further talking points:
    - clock synchronization
    - section length tuning
    - high availability

Ch_8: Design a url shortener:

- Built as an API endpoint
- 301 redirect: permanently moved to the long URL
- 302 redirect: temporarily moved to the short URL
- Hash + collision resolution:
    - Check bloom filtering ??
- Base 62 conversion:
    
    ![Diagram](https://prod-files-secure.s3.us-west-2.amazonaws.com/70c3badf-7b78-4bb0-9263-f7c8e5e66fec/45d959a7-af1e-4f1e-89ae-e1745d44d91a/Untitled.png)
    

Ch_9: Design a web crawler:

- [check book]

Ch_10: Design a notification system:

![Diagram](https://prod-files-secure.s3.us-west-2.amazonaws.com/70c3badf-7b78-4bb0-9263-f7c8e5e66fec/e828d041-13e5-4ca9-ad9c-e6c423e930ce/Untitled.png)

Ch_11 and Ch_12 in book read it
