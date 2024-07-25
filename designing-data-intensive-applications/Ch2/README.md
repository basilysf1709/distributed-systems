# Ch2: Super concise sumamry
- Data Models have a profound effect no how a system is built
- Popular Data Models include:
  - **SQL**
    - Relational
    - Hide implementation detail behind a cleaner interface
  - **NoSQL**
    - Non-relational
    - Document stores, Wide Column, Graph, etc
    - People were frustrated with the restriveness of relational model
    - There was a need for greater scalability than relational databases can easily achieve, includ‐
ing very large datasets or **very high write throughput**
- Some data requirements are better suited for JSON document stores
  - One-to-many
    - Usually requires joins or complicated queries
    - With something like MongoDB, it can be easily queried once and you get the data
  - Sometimes other relationships like many-to-one and many-to-many are better suited for SQL Databases because **joins** (SQL) are usually easier than **multiple querying** (Document Stores)
- Historical debate between the network model and the SQL model
- Query Optimizer made SQL models **general purpose** (query your data in the way you want to)
- Arguments for SQL:
  - easier to do joins on **relational** data
- Arguments for Document Stores:
  - better performance due to data **locality** (related data is usually close to each other
  - flexibility
- Document stores have **schema-on-read** whereas SQL databases have **schema-on-write**
- Stopped at **Data locality for queries** tbc
  
    
