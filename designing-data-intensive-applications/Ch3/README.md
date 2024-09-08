# Ch3: Super concise sumamry

## Storage Engines
- Two main types: log-structured and page-oriented

### Log-Structured
- Append-only data file
- Efficient for writing, slow for reading
- Use indices to improve read performance

### Indices
- Slow down writes, but speed up specific reads
- Must be explicitly defined by developers

#### Types:
1. Hash Indices
   - Good for key-value data
   - Limited by memory, inefficient for range queries

2. Log-Structured Merge-Tree (LSM)
   - Uses in-memory 'memtable' and on-disk 'SSTable'
   - Efficient for writes

3. B-trees
   - Most popular for relational and many non-relational DBs
   - Uses fixed-size blocks/pages
   - Generally faster for reads than LSM-trees

## OLTP vs OLAP Databases
- OLTP: User-facing, many short-lasting queries
- OLAP: For analytics, fewer but more complex queries

### Data Warehouses
- Dedicated OLAP databases
- Use ETL (Extract, Transform, Load) to populate
- Often use star or snowflake schemas

### Column-Oriented Storage
- Common in OLAP databases
- Stores column data together on disk
- More efficient for analytical queries that access few columns but many rows
