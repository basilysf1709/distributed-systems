# Rate Limiting Algorithms

## Overview

Ground implementation of rate limiting algorithms for a simple go api
#### 1. Token Bucket Algorithm (Inspired by Alex Xu's - System Design Interview Book)
- Tokens are added at a fixed rate over time
- Implementation is in go-api/nameservice/TokenBucket.go
- If there are no tokens in the bucket, the api gets rate-limited and sends HTTP 429 response. (Too many requests)

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go (1.18 or later)
- Docker and Docker Compose (for containerized environment)

### Installing

1. Clone the repository:
   ```bash
   git clone https://github.com/basilysf1709/rate-limiting-as-a-service.git
2. Docker Compose
   ```bash
   make
3. Docker Compose Down
   ```bash
   make down
