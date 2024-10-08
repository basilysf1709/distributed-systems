# Ch1 Super concise summary
- Desiging Data-intensive applications require **reliability**, **scalability**, and **maintainability**.
- **Reliability**: System should work correctly despite hardware/software faults and human errors.
- **Scalability:** System should be able to handle increased load efficiently.
- **Maintainability:** Make systems easier to operate, understand, and evolve over time.
- **Load parameter**s and **metrics** are extremely important when it comes to dicussing scalability
- **Response time percentile**s are better metrics than averages for measuring performance. 
- For distributed systems, sometimes it's important to look at the **slowest request** to understand overall performance of the system.
- Approaches to handle load: vertical scaling (more powerful machines) and horizontal scaling (distributing across machines).
- **Good abstractions** and removing accidental complexity are key to managing system complexity.
- **Evolvability** is important as system requirements constantly change.
- **Evolvability** is the ease with which a data system can be modified and adapted to changing requirements over time
- There's no one-size-fits-all scalable architecture; solutions depend on **specific use cases.**
